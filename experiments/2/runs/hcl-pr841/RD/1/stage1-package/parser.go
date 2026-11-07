// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0


package hclwrite

import (
	"fmt"
	"sort"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

// Our "parser" here is actually not doing any parsing of its own. Instead,
// it leans on the native parser in hclsyntax, and then uses the source ranges
// from the AST to partition the raw token sequence to match the raw tokens
// up to AST nodes.
//
// This strategy feels somewhat counter-intuitive, since most of the work the
// parser does is thrown away here, but this strategy is chosen because the
// normal parsing work done by hclsyntax is considered to be the "main case",
// while modifying and re-printing source is more of an edge case, used only
// in ancillary tools, and so it's good to keep all the main parsing logic
// with the main case but keep all of the extra complexity of token wrangling
// out of the main parser, which is already rather complex just serving the
// use-cases it already serves.
//
// If the parsing step produces any errors, the returned File is nil because
// we can't reliably extract tokens from the partial AST produced by an
// erroneous parse.
//
func parse(src []byte, filename string, start hcl.Pos) (*File, hcl.Diagnostics) {
	file, diags := hclsyntax.ParseConfig(src, filename, start)
	if diags.HasErrors() {
		return nil, diags
	}

	// To do our work here, we use the "native" tokens (those from hclsyntax)
	// to match against source ranges in the AST, but ultimately produce
	// slices from our sequence of "writer" tokens, which contain only
	// *relative* position information that is more appropriate for
	// transformation/writing use-cases.
	nativeTokens, diags := hclsyntax.LexConfig(src, filename, start)
	if diags.HasErrors() {
		// should never happen, since we would've caught these diags in
		// the first call above.
		return nil, diags
	}
	from := newPartitionableTokens(nativeTokens)

	before, root, after := parseBody(file.Body.(*hclsyntax.Body), from)
	ret := &File{
		inTree: newInTree(),

		srcBytes: src,
		body:     root,
	}

	nodes := ret.children
	nodes.Append(before.Tokens())
	nodes.AppendNode(root)
	nodes.Append(after.Tokens())

	return ret, diags
}

// parseBody locates the given body within the given input tokens and returns
// the resulting *Body object as well as the tokens that appeared before and
// after it.
func parseBody(nativeBody *hclsyntax.Body, from partitionableTokens) (partitionableTokens, *node, partitionableTokens) {
	before, within, after := from.PartitionIncludingComments(nativeBody.SrcRange)

	// The main AST doesn't retain the original source ordering of the
	// body items, so we need to reconstruct that ordering by inspecting
	// their source ranges.
	nativeItems := make([]hclsyntax.Node, 0, len(nativeBody.Attributes)+len(nativeBody.Blocks))
	for _, nativeAttr := range nativeBody.Attributes {
		nativeItems = append(nativeItems, nativeAttr)
	}
	for _, nativeBlock := range nativeBody.Blocks {
		nativeItems = append(nativeItems, nativeBlock)
	}
	sort.Sort(nativeNodeSorter{nativeItems})

	body := &Body{
		inTree: newInTree(),
		items:  newNodeSet(),
	}

	remain := within
	for _, nativeItem := range nativeItems {
		beforeItem, item, afterItem := parseBodyItem(nativeItem, remain)

		if beforeItem.Len() > 0 {
			body.AppendUnstructuredTokens(beforeItem.Tokens())
		}
		body.appendItemNode(item)

		remain = afterItem
	}

	if remain.Len() > 0 {
		body.AppendUnstructuredTokens(remain.Tokens())
	}

	return before, newNode(body), after
}

func parseBodyItem(nativeItem hclsyntax.Node, from partitionableTokens) (partitionableTokens, *node, partitionableTokens) {
	before, leadComments, within, lineComments, newline, after := from.PartitionBlockItem(nativeItem.Range())

	var item *node

	switch tItem := nativeItem.(type) {
	case *hclsyntax.Attribute:
		item = parseAttribute(tItem, within, leadComments, lineComments, newline)
	case *hclsyntax.Block:
		item = parseBlock(tItem, within, leadComments, lineComments, newline)
	default:
		// should never happen if caller is behaving
		panic("unsupported native item type")
	}

	return before, item, after
}

func parseAttribute(nativeAttr *hclsyntax.Attribute, from, leadComments, lineComments, newline partitionableTokens) *node {
	attr := &Attribute{
		inTree: newInTree(),
	}
	children := attr.children

	{
		cn := newNode(newComments(leadComments.Tokens()))
		attr.leadComments = cn
		children.AppendNode(cn)
	}

	before, nameTokens, from := from.Partition(nativeAttr.NameRange)
	{
		children.AppendUnstructuredTokens(before.Tokens())
		if nameTokens.Len() != 1 {
			// Should never happen with valid input
			panic("attribute name is not exactly one token")
		}
		token := nameTokens.Tokens()[0]
		in := newNode(newIdentifier(token))
		attr.name = in
		children.AppendNode(in)
	}

	before, equalsTokens, from := from.Partition(nativeAttr.EqualsRange)
	children.AppendUnstructuredTokens(before.Tokens())
	children.AppendUnstructuredTokens(equalsTokens.Tokens())

	before, exprTokens, from := from.Partition(nativeAttr.Expr.Range())
	{
		children.AppendUnstructuredTokens(before.Tokens())
		exprNode := parseExpression(nativeAttr.Expr, exprTokens)
		attr.expr = exprNode
		children.AppendNode(exprNode)
	}

	{
		cn := newNode(newComments(lineComments.Tokens()))
		attr.lineComments = cn
		children.AppendNode(cn)
	}

	children.AppendUnstructuredTokens(newline.Tokens())

	// Collect any stragglers, though there shouldn't be any
	children.AppendUnstructuredTokens(from.Tokens())

	return newNode(attr)
}

func parseBlock(nativeBlock *hclsyntax.Block, from, leadComments, lineComments, newline partitionableTokens) *node {
	block := &Block{
		inTree: newInTree(),
	}
	children := block.children

	{
		cn := newNode(newComments(leadComments.Tokens()))
		block.leadComments = cn
		children.AppendNode(cn)
	}

	before, typeTokens, from := from.Partition(nativeBlock.TypeRange)
	{
		children.AppendUnstructuredTokens(before.Tokens())
		if typeTokens.Len() != 1 {
			// Should never happen with valid input
			panic("block type name is not exactly one token")
		}
		token := typeTokens.Tokens()[0]
		in := newNode(newIdentifier(token))
		block.typeName = in
		children.AppendNode(in)
	}

	_, labelsNode, from := parseBlockLabels(nativeBlock, from)
	block.labels = labelsNode
	children.AppendNode(labelsNode)

	before, oBrace, from := from.Partition(nativeBlock.OpenBraceRange)
	children.AppendUnstructuredTokens(before.Tokens())
	block.open = children.AppendUnstructuredTokens(oBrace.Tokens())

	// We go a bit out of order here: we go hunting for the closing brace
	// so that we have a delimited body, but then we'll deal with the body
	// before we actually append the closing brace and any straggling tokens
	// that appear after it.
	bodyTokens, cBrace, from := from.Partition(nativeBlock.CloseBraceRange)
	before, body, after := parseBody(nativeBlock.Body, bodyTokens)
	children.AppendUnstructuredTokens(before.Tokens())
	block.body = body
	children.AppendNode(body)
	children.AppendUnstructuredTokens(after.Tokens())

	block.close = children.AppendUnstructuredTokens(cBrace.Tokens())

	// stragglers
	children.AppendUnstructuredTokens(from.Tokens())
	if lineComments.Len() > 0 {
		// blocks don't actually have line comments, so we'll just treat
		// them as extra stragglers
		children.AppendUnstructuredTokens(lineComments.Tokens())
	}
	children.AppendUnstructuredTokens(newline.Tokens())

	return newNode(block)
}

func parseBlockLabels(nativeBlock *hclsyntax.Block, from partitionableTokens) (partitionableTokens, *node, partitionableTokens) {
	labelsObj := newBlockLabels(nil)
	children := labelsObj.children

	var beforeAll partitionableTokens
	for i, rng := range nativeBlock.LabelRanges {
		var before, labelTokens partitionableTokens
		before, labelTokens, from = from.Partition(rng)
		if i == 0 {
			beforeAll = before
		} else {
			children.AppendUnstructuredTokens(before.Tokens())
		}
		tokens := labelTokens.Tokens()
		var ln *node
		if len(tokens) == 1 && tokens[0].Type == hclsyntax.TokenIdent {
			ln = newNode(newIdentifier(tokens[0]))
		} else {
			ln = newNode(newQuoted(tokens))
		}
		labelsObj.items.Add(ln)
		children.AppendNode(ln)
	}

	after := from
	return beforeAll, newNode(labelsObj), after
}

func parseExpression(nativeExpr hclsyntax.Expression, from partitionableTokens) *node {
	switch tNativeExpr := nativeExpr.(type) {

	// Object-construct expression
	case *hclsyntax.ObjectConsExpr:
		return parseObjectConsExpr(tNativeExpr, from)

	case *hclsyntax.ObjectConsKeyExpr:
		return parseObjectConsKeyExpr(tNativeExpr, from)

	case *hclsyntax.FunctionCallExpr:
		return parseFunctionCallExpr(tNativeExpr, from)

	case *hclsyntax.TupleConsExpr:
		return parseTupleConsExpr(tNativeExpr, from)

	case *hclsyntax.TemplateExpr:
		return parseTemplateExpr(tNativeExpr, from)

	default:
		return parseAnyExpression(nativeExpr, from)
	}
}

func parseAnyExpression(nativeExpr hclsyntax.Expression, from partitionableTokens) *node {
	expr := newExpression()
	children := expr.children

	nativeVars := nativeExpr.Variables()

	for _, nativeTraversal := range nativeVars {
		before, traversal, after := parseTraversal(nativeTraversal, from)
		children.AppendUnstructuredTokens(before.Tokens())
		children.AppendNode(traversal)
		expr.absTraversals.Add(traversal)
		from = after
	}
	// Attach any stragglers that don't belong to a traversal to the expression
	// itself. In an expression with no traversals at all, this is just the
	// entirety of "from".
	children.AppendUnstructuredTokens(from.Tokens())

	return newNode(expr)
}

// parseObjectConsExpr parses an object-construct expression, defined as:
//
//	object = "{" (
//	    (objectelem (( "," | Newline) objectelem)* ","?)?
//	) "}";
//	objectelem = (Identifier | Expression) ("=" | ":") Expression;
//
// It leaves any lead comments and line comments as unstructured tokens.
func parseObjectConsExpr(nativeExpr *hclsyntax.ObjectConsExpr, from partitionableTokens) *node {
	expr := newObjectConsExpr()
	children := expr.children

	var before, open, keyTokens, valueTokens partitionableTokens

	before, open, from = from.Partition(nativeExpr.StartRange())
	children.AppendUnstructuredTokens(before.Tokens())
	children.AppendUnstructuredTokens(open.Tokens())

	for _, nativeItem := range nativeExpr.Items {
		item := newObjectConsItem()
		value := newObjectConsValue()

		nativeKeyExpr := nativeItem.KeyExpr.(*hclsyntax.ObjectConsKeyExpr)

		before, keyTokens, from = from.Partition(nativeKeyExpr.Range())
		item.children.AppendUnstructuredTokens(before.Tokens())
		item.key = parseObjectConsKeyExpr(nativeKeyExpr, keyTokens)
		item.children.AppendNode(item.key)

		before, valueTokens, from = from.Partition(nativeItem.ValueExpr.Range())
		item.children.AppendUnstructuredTokens(before.Tokens())
		value.expr = parseExpression(nativeItem.ValueExpr, valueTokens)
		value.children.AppendNode(value.expr)
		item.value = item.children.Append(value)

		node := expr.children.Append(item)
		expr.items.Add(node)
	}

	_, from, _ = from.Partition(nativeExpr.Range())
	children.AppendUnstructuredTokens(from.Tokens())

	// Wrap in an Expression
	wrapExpr := newExpression()
	wrapExpr.wrap(expr)
	return newNode(wrapExpr)
}

// parseObjectConsKeyExpr parses an object-construct key expression.
// An object key is an Identifier or an Expression. In this context, a quoted
// literal is functionally equivalent to an Identifier.
//
// A key that is a naked identifier or a quoted string literal is recorded
// with the name it evaluates to, so that it can be looked up by that name
// and its traversal is not taken as a variable.
func parseObjectConsKeyExpr(nativeExpr *hclsyntax.ObjectConsKeyExpr, from partitionableTokens) *node {
	expr := parseExpression(nativeExpr.Wrapped, from)

	if !nativeExpr.ForceNonLiteral {
		if name := hcl.ExprAsKeyword(nativeExpr.Wrapped); name != "" {
			return newNode(newObjectConsKeyExpr(expr, name, true))
		}
		if template, ok := nativeExpr.Wrapped.(*hclsyntax.TemplateExpr); ok && template.IsStringLiteral() {
			val, diags := template.Value(nil)
			if !diags.HasErrors() {
				return newNode(newObjectConsKeyExpr(expr, val.AsString(), true))
			}
		}
	}

	return newNode(newObjectConsKeyExpr(expr, "", false))
}

// parseFunctionCallExpr parses a function call expression, defined as:
//
//	FunctionCall = Identifier ("::" Identifier)* "(" Arguments ")";
//	Arguments = (
//	    () ||
//	    (Expression ("," Expression)* ("," | "...")?)
//	);
//
// Each argument becomes an Expression of its own. The tokens around them,
// including any comments, are kept as unstructured tokens of the call.
func parseFunctionCallExpr(nativeExpr *hclsyntax.FunctionCallExpr, from partitionableTokens) *node {
	call := newFunctionCallExpr()
	parseExpressionList(nativeExpr.Args, from, call.children, call.args)

	wrapExpr := newExpression()
	wrapExpr.wrap(call)
	return newNode(wrapExpr)
}

// parseTupleConsExpr parses a tuple-construct expression, defined as:
//
//	tuple = "[" (
//	    (Expression ("," Expression)* ","?)?
//	) "]";
//
// Each element becomes an Expression of its own. The tokens around them,
// including any comments, are kept as unstructured tokens of the tuple.
func parseTupleConsExpr(nativeExpr *hclsyntax.TupleConsExpr, from partitionableTokens) *node {
	tuple := newTupleConsExpr()
	parseExpressionList(nativeExpr.Exprs, from, tuple.children, tuple.elems)

	wrapExpr := newExpression()
	wrapExpr.wrap(tuple)
	return newNode(wrapExpr)
}

// parseExpressionList parses each of the given expressions from the given
// tokens, appends it to children and adds it to items. The tokens before,
// between and after the expressions are appended to children as
// unstructured tokens.
func parseExpressionList(nativeExprs []hclsyntax.Expression, from partitionableTokens, children *nodes, items nodeSet) {
	for _, nativeExpr := range nativeExprs {
		before, exprTokens, after := from.Partition(nativeExpr.Range())
		children.AppendUnstructuredTokens(before.Tokens())
		expr := parseExpression(nativeExpr, exprTokens)
		children.AppendNode(expr)
		items.Add(expr)
		from = after
	}
	children.AppendUnstructuredTokens(from.Tokens())
}

// parseTemplateExpr is specifically interested in string literal expressions
// at this time.
func parseTemplateExpr(nativeExpr *hclsyntax.TemplateExpr, from partitionableTokens) *node {
	if nativeExpr.IsStringLiteral() {
		quoted := newQuoted(from.Tokens())

		expr := newExpression()
		expr.wrapped = expr.children.Append(quoted)
		return newNode(expr)
	} else {
		return parseAnyExpression(nativeExpr, from)
	}
}

func parseTraversal(nativeTraversal hcl.Traversal, from partitionableTokens) (before partitionableTokens, n *node, after partitionableTokens) {
	traversal := newTraversal()
	children := traversal.children
	before, from, after = from.Partition(nativeTraversal.SourceRange())

	stepAfter := from
	for _, nativeStep := range nativeTraversal {
		before, step, after := parseTraversalStep(nativeStep, stepAfter)
		children.AppendUnstructuredTokens(before.Tokens())
		children.AppendNode(step)
		traversal.steps.Add(step)
		stepAfter = after
	}

	return before, newNode(traversal), after
}

func parseTraversalStep(nativeStep hcl.Traverser, from partitionableTokens) (before partitionableTokens, n *node, after partitionableTokens) {
	var children *nodes
	switch tNativeStep := nativeStep.(type) {

	case hcl.TraverseRoot, hcl.TraverseAttr:
		step := newTraverseName()
		children = step.children
		before, from, after = from.Partition(nativeStep.SourceRange())
		inBefore, token, inAfter := from.PartitionTypeSingle(hclsyntax.TokenIdent)
		name := newIdentifier(token)
		children.AppendUnstructuredTokens(inBefore.Tokens())
		step.name = children.Append(name)
		children.AppendUnstructuredTokens(inAfter.Tokens())
		return before, newNode(step), after

	case hcl.TraverseIndex:
		step := newTraverseIndex()
		children = step.children
		before, from, after = from.Partition(nativeStep.SourceRange())

		if inBefore, dot, from, ok := from.PartitionTypeOk(hclsyntax.TokenDot); ok {
			children.AppendUnstructuredTokens(inBefore.Tokens())
			children.AppendUnstructuredTokens(dot.Tokens())

			valBefore, valToken, valAfter := from.PartitionTypeSingle(hclsyntax.TokenNumberLit)
			children.AppendUnstructuredTokens(valBefore.Tokens())
			key := newNumber(valToken)
			step.key = children.Append(key)
			children.AppendUnstructuredTokens(valAfter.Tokens())

			return before, newNode(step), after
		}

		var inBefore, oBrack, keyTokens, cBrack partitionableTokens
		inBefore, oBrack, from = from.PartitionType(hclsyntax.TokenOBrack)
		children.AppendUnstructuredTokens(inBefore.Tokens())
		children.AppendUnstructuredTokens(oBrack.Tokens())
		keyTokens, cBrack, from = from.PartitionType(hclsyntax.TokenCBrack)

		keyVal := tNativeStep.Key
		switch keyVal.Type() {
		case cty.String:
			key := newQuoted(keyTokens.Tokens())
			step.key = children.Append(key)
		case cty.Number:
			valBefore, valToken, valAfter := keyTokens.PartitionTypeSingle(hclsyntax.TokenNumberLit)
			children.AppendUnstructuredTokens(valBefore.Tokens())
			key := newNumber(valToken)
			step.key = children.Append(key)
			children.AppendUnstructuredTokens(valAfter.Tokens())
		}

		children.AppendUnstructuredTokens(cBrack.Tokens())
		children.AppendUnstructuredTokens(from.Tokens())

		return before, newNode(step), after
	default:
		panic(fmt.Sprintf("unsupported traversal step type %T", nativeStep))
	}

}
