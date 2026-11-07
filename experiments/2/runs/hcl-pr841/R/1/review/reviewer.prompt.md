This directory holds one Go package, `pkg/`. Some of its files were just added
or changed; they are marked in the map below, and their full text follows it.
You may read any other file under pkg/.

Review the package as a code reviewer would. List the refactoring that should
be done now. Focus on the new and changed code and how it fits the package.

### pkg/

| file | lines | mark | declarations |
| --- | --- | --- | --- |
| ast.go | 121 |  | File (type), NewEmptyFile (func), File.Body (method), File.WriteTo (method), File.Bytes (method), comments (type), newComments (func), comments.BuildTokens (method), identifier (type), newIdentifier (func), identifier.BuildTokens (method), identifier.hasName (method), number (type), newNumber (func), number.BuildTokens (method), quoted (type), newQuoted (func), quoted.BuildTokens (method) |
| ast_attribute.go | 65 |  | Attribute (type), newAttribute (func), Attribute.init (method), Attribute.LeadComments (method), Attribute.LineComments (method), Attribute.Expr (method), Attribute.setName (method) |
| ast_attribute_test.go | 125 |  | TestAttributeLeadComments (func), TestAttributeLineComments (func) |
| ast_block.go | 184 |  | Block (type), newBlock (func), NewBlock (func), Block.init (method), Block.Body (method), Block.LeadComments (method), Block.Type (method), Block.SetType (method), Block.Labels (method), Block.SetLabels (method), Block.labelsObj (method), blockLabels (type), newBlockLabels (func), blockLabels.Replace (method), blockLabels.Current (method) |
| ast_block_test.go | 579 |  | TestBlockType (func), TestBlockLabels (func), TestBlockSetType_buildTokens (func), TestBlockSetType (func), TestBlockSetLabels (func), TestBlockLeadComments (func), mustParseConfig (func) |
| ast_body.go | 297 | CHANGED | Body (type), newBody (func), Body.appendItem (method), Body.appendItemNode (method), Body.Clear (method), Body.AppendUnstructuredTokens (method), Body.Attributes (method), Body.Blocks (method), Body.GetAttribute (method), Body.getAttributeNode (method), Body.RenameAttribute (method), Body.FirstMatchingBlock (method), Body.RemoveNewlineBeforeBlock (method), Body.RemoveBlock (method), Body.SetAttributeRaw (method), Body.SetAttributeValue (method), Body.SetAttributeTraversal (method), Body.setAttribute (method), Body.RemoveAttribute (method), Body.AppendBlock (method), Body.AppendNewBlock (method), Body.AppendNewline (method) |
| ast_body_test.go | 2195 |  | TestBodyGetAttribute (func), TestBodyFirstMatchingBlock (func), TestBodySetAttributeValue (func), TestBodySetAttributeTraversal (func), TestBodySetAttributeTraversal_ReturnsTheAttribute (func), TestBodySetAttributeRaw (func), TestBodySetAttributeRaw_ReturnsTheAttribute (func), TestBodySetAttributeValue_ReturnsTheAttribute (func), TestBodySetAttributeValueInBlock (func), TestBodySetAttributeValueInNestedBlock (func), TestBodyRemoveAttribute (func), TestBodyRenameAttribute (func), TestBodyAppendBlock (func), TestBodyRemoveBlock (func), TestBodyRemoveNewlineBeforeBlock (func), testFn (func) |
| ast_expression.go | 323 | CHANGED | Expression (type), newExpression (func), Expression.wrap (method), NewExpressionRaw (func), NewExpressionLiteral (func), NewExpressionAbsTraversal (func), Expression.Variables (method), nestedTraversals (type), Expression.traversalNodes (method), exprsTraversalNodes (func), Expression.RenameVariablePrefix (method), Expression.AsObjectConsExpr (method), Expression.AsFunctionCallExpr (method), Expression.AsTupleConsExpr (method), Expression.asQuoted (method), wrappedAs (func), Expression.AsQuotedLiteral (method), Traversal (type), newTraversal (func), Traversal.renamePrefix (method), TraverseName (type), newTraverseName (func), TraverseIndex (type), newTraverseIndex (func) |
| ast_expression_test.go | 166 | NEW | TestExpressionVariablesAndRename (func), TestExpressionAsFunctionCallAndTupleConsExpr (func) |
| ast_function_call_expr.go | 41 | NEW | FunctionCallExpr (type), newFunctionCallExpr (func), FunctionCallExpr.Name (method), FunctionCallExpr.traversalNodes (method) |
| ast_function_call_expr_test.go | 88 | NEW | parseTestExpr (func), checkRoundTrip (func), exprsStrings (func), TestFunctionCallExprParse (func) |
| ast_object_cons_expr.go | 397 | CHANGED | ObjectConsExpr (type), newObjectConsExpr (func), ObjectConsExpr.Items (method), ObjectConsExpr.ItemFor (method), ObjectConsExpr.ValueExprFor (method), ObjectConsExpr.SetItem (method), ObjectConsExpr.SetItemRaw (method), ObjectConsExpr.SetItemValue (method), ObjectConsExpr.SetItemTraversal (method), ObjectConsExpr.appendItem (method), ObjectConsExpr.isMultiline (method), ObjectConsExpr.traversalNodes (method), ObjectConsItem (type), newObjectConsItem (func), ObjectConsItem.init (method), ObjectConsItem.KeyExpr (method), ObjectConsItem.ValueExpr (method), ObjectConsItem.hasLineBreak (method), objectConsKeyKind (type), objectConsKeyExpression (const), objectConsKeyLiteralName (const), objectConsKeyQuotedLiteral (const), ObjectConsKeyExpr (type), newObjectConsKeyExpr (func), newObjectConsKeyExprForName (func), ObjectConsKeyExpr.hasName (method), ObjectConsKeyExpr.traversalNodes (method), ObjectConsValue (type), newObjectConsValue (func), ObjectConsValue.Expr (method), ObjectConsValue.setExpr (method), splitLineEnd (func), isLineBreak (func), hasLineBreak (func), hasComma (func), newNewlineToken (func), newCommaToken (func) |
| ast_object_cons_expr_test.go | 604 | NEW | parseTestObject (func), evalAttribute (func), tokensString (func), TestObjectConsExprItemFor (func), TestObjectConsExprItems (func), TestObjectConsExprSetItemKey (func), TestObjectConsExprSetItemExisting (func), TestObjectConsExprSetItemHelpers (func), TestObjectConsExprSetItemPanics (func), TestObjectConsExprSetItemLayout (func) |
| ast_test.go | 50 |  | TestTreeNode (type), makeTestTree (func) |
| ast_tuple_cons_expr.go | 31 | NEW | TupleConsExpr (type), newTupleConsExpr (func), TupleConsExpr.traversalNodes (method) |
| ast_tuple_cons_expr_test.go | 44 | NEW | TestTupleConsExprParse (func) |
| doc.go | 14 |  |  |
| examples_test.go | 125 | CHANGED | Example_generateFromScratch (func), ExampleExpression_RenameVariablePrefix (func), ExampleObjectConsExpr_SetItemValue (func) |
| format.go | 472 |  | format (func), formatIndent (func), formatSpaces (func), formatCells (func), spaceAfterToken (func), linesForFormat (func), tokenIsNewline (func), tokenBracketChange (func), formatLine (type) |
| format_test.go | 800 |  | TestFormat (func), TestLinesForFormat (func) |
| generate.go | 399 |  | TokensForValue (func), TokensForTraversal (func), TokensForIdentifier (func), TokensForTuple (func), TokensForObject (func), TokensForFunctionCall (func), appendTokensForValue (func), appendTokensForTraversal (func), appendTokensForTraversalStep (func), escapeQuotedStringLit (func), appendRune (func) |
| generate_test.go | 920 |  | TestTokensForValue (func), TestTokensForTraversal (func), TestTokensForTuple (func), TestTokensForObject (func), TestTokensForFunctionCall (func), TestTokenGenerateConsistency (func) |
| native_node_sorter.go | 26 |  | nativeNodeSorter (type), nativeNodeSorter.Len (method), nativeNodeSorter.Less (method), nativeNodeSorter.Swap (method) |
| node.go | 319 | CHANGED | node (type), newNode (func), node.Equal (method), node.BuildTokens (method), node.Detach (method), node.ReplaceWith (method), node.assertUnattached (method), nodeContent (type), nodes (type), nodes.BuildTokens (method), nodes.Clear (method), nodes.Append (method), nodes.AppendNode (method), nodes.Insert (method), nodes.InsertNode (method), nodes.AppendUnstructuredTokens (method), nodes.FindNodeWithContent (method), nodeSet (type), newNodeSet (func), nodeSet.Has (method), nodeSet.Add (method), nodeSet.Remove (method), nodeSet.Clear (method), nodeSet.List (method), nodeSet.FindNodeWithContent (method), internalWalkFunc (type), heldContent (type), inTree (type), newInTree (func), inTree.setHolder (method), inTree.assertUnattached (method), inTree.walkChildNodes (method), inTree.BuildTokens (method), leafNode (type), leafNode.walkChildNodes (method) |
| parser.go | 816 | CHANGED | parse (func), inputTokens (type), inputTokens.Partition (method), inputTokens.PartitionType (method), inputTokens.PartitionTypeOk (method), inputTokens.PartitionTypeSingle (method), inputTokens.PartitionIncludingComments (method), inputTokens.PartitionBlockItem (method), inputTokens.PartitionLeadComments (method), inputTokens.PartitionLineEndTokens (method), inputTokens.Slice (method), inputTokens.Len (method), inputTokens.Tokens (method), inputTokens.Types (method), parseBody (func), parseBodyItem (func), parseAttribute (func), parseBlock (func), parseBlockLabels (func), parseExpression (func), parseAnyExpression (func), parseObjectConsExpr (func), parseFunctionCallExpr (func), parseTupleConsExpr (func), parseExpressionList (func), wrapExpressionNode (func), parseObjectConsKeyExpr (func), objectConsKeyKindOf (func), parseTemplateExpr (func), parseTraversal (func), parseTraversalStep (func), writerTokens (func), partitionTokens (func), partitionLeadCommentTokens (func), partitionLineEndTokens (func), lexConfig (func) |
| parser_test.go | 2285 | CHANGED | TestParse (func), TestPartitionTokens (func), TestPartitionLeadCommentTokens (func), TestLexConfig (func) |
| public.go | 48 |  | NewFile (func), ParseConfig (func), Format (func) |
| round_trip_test.go | 181 |  | TestRoundTripVerbatim (func), TestRoundTripFormat (func), TestRoundTripSafeConcurrent (func) |
| tokens.go | 136 |  | Token (type), Token.asHCLSyntax (method), Tokens (type), Tokens.Bytes (method), Tokens.testValue (method), Tokens.Columns (method), Tokens.WriteTo (method), Tokens.walkChildNodes (method), Tokens.BuildTokens (method), ObjectAttrTokens (type), newIdentToken (func) |

#### pkg/ast_body.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"reflect"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

type Body struct {
	inTree

	items nodeSet
}

func newBody() *Body {
	return &Body{
		inTree: newInTree(),
		items:  newNodeSet(),
	}
}

func (b *Body) appendItem(c nodeContent) *node {
	nn := b.children.Append(c)
	b.items.Add(nn)
	return nn
}

func (b *Body) appendItemNode(nn *node) *node {
	nn.assertUnattached()
	b.children.AppendNode(nn)
	b.items.Add(nn)
	return nn
}

// Clear removes all of the items from the body, making it empty.
func (b *Body) Clear() {
	b.children.Clear()
}

func (b *Body) AppendUnstructuredTokens(ts Tokens) {
	b.children.Append(ts)
}

// Attributes returns a new map of all of the attributes in the body, with
// the attribute names as the keys.
func (b *Body) Attributes() map[string]*Attribute {
	ret := make(map[string]*Attribute)
	for n := range b.items {
		if attr, isAttr := n.content.(*Attribute); isAttr {
			nameObj := attr.name.content.(*identifier)
			name := string(nameObj.token.Bytes)
			ret[name] = attr
		}
	}
	return ret
}

// Blocks returns a new slice of all the blocks in the body.
func (b *Body) Blocks() []*Block {
	ret := make([]*Block, 0, len(b.items))
	for _, n := range b.items.List() {
		if block, isBlock := n.content.(*Block); isBlock {
			ret = append(ret, block)
		}
	}
	return ret
}

// GetAttribute returns the attribute from the body that has the given name,
// or returns nil if there is currently no matching attribute.
func (b *Body) GetAttribute(name string) *Attribute {
	for n := range b.items {
		if attr, isAttr := n.content.(*Attribute); isAttr {
			nameObj := attr.name.content.(*identifier)
			if nameObj.hasName(name) {
				// We've found it!
				return attr
			}
		}
	}

	return nil
}

// getAttributeNode is like GetAttribute but it returns the node containing
// the selected attribute (if one is found) rather than the attribute itself.
func (b *Body) getAttributeNode(name string) *node {
	for n := range b.items {
		if attr, isAttr := n.content.(*Attribute); isAttr {
			nameObj := attr.name.content.(*identifier)
			if nameObj.hasName(name) {
				// We've found it!
				return n
			}
		}
	}

	return nil
}

// RenameAttribute changes the attribute named fromName to toName.
// Takes no action if fromName is missing or there is already a
// conflicting attribute called toName.
//
// Returns true if the rename succeeded.
func (b *Body) RenameAttribute(fromName, toName string) bool {
	attr := b.GetAttribute(fromName)
	conflictingAttr := b.GetAttribute(toName)
	if attr == nil || conflictingAttr != nil {
		return false
	}
	attr.setName(toName)
	return true
}

// FirstMatchingBlock returns a first matching block from the body that has the
// given name and labels or returns nil if there is currently no matching
// block.
func (b *Body) FirstMatchingBlock(typeName string, labels []string) *Block {
	for _, block := range b.Blocks() {
		if typeName == block.Type() {
			labelNames := block.Labels()
			if len(labels) == 0 && len(labelNames) == 0 {
				return block
			}
			if reflect.DeepEqual(labels, labelNames) {
				return block
			}
		}
	}

	return nil
}

// RemoveNewlineBeforeBlock makes a best effort to remove one newline of
// vertical separation before a block.
//
// It will not remove the trailing newline of a preceding structural element,
// such as an attribute or a block.

// The pair of RemoveNewlineBeforeBlock + RemoveBlock is an "undo" move for the
// pair of AppendNewline + AppendNewBlock.
func (b *Body) RemoveNewlineBeforeBlock(block *Block) bool {
	for n := range b.items {
		if n.content != block {
			continue
		}
		if n.before == nil {
			return false
		}

		// Detect "unstructured token"-only nodes. Structured content, such as
		// an attribute, includes its trailing newline, so we will not remove
		// that.
		if tokens, ok := n.before.content.(Tokens); ok {
			count := len(tokens)
			switch count {
			case 0:
				break

			case 1:
				// It seems unlikely for this to be anything other than a
				// newline. And yet, we verify.
				if tokens[0].Type == hclsyntax.TokenNewline {
					b.items.Remove(n.before)
					n.before.Detach()
					return true
				}

			default:
				// It seems unlikely for this to be anything other than a
				// newline. And yet, we verify.
				if tokens[count-1].Type == hclsyntax.TokenNewline {
					newTokens := tokens[:count-1]
					n.before.ReplaceWith(newTokens)
					return true
				}
			}
		}
	}
	return false
}

// RemoveBlock removes the given block from the body, if it's in that body.
// If it isn't present, this is a no-op.
//
// Returns true if it removed something, or false otherwise.
func (b *Body) RemoveBlock(block *Block) bool {
	for n := range b.items {
		if n.content == block {
			n.Detach()
			b.items.Remove(n)
			return true
		}
	}
	return false
}

// SetAttributeRaw either replaces the expression of an existing attribute
// of the given name or adds a new attribute definition to the end of the block,
// using the given tokens verbatim as the expression.
//
// The same caveats apply to this function as for NewExpressionRaw on which
// it is based. If possible, prefer to use SetAttributeValue or
// SetAttributeTraversal.
//
// The return value is the attribute that was either modified in-place or
// created.
func (b *Body) SetAttributeRaw(name string, tokens Tokens) *Attribute {
	return b.setAttribute(name, NewExpressionRaw(tokens))
}

// SetAttributeValue either replaces the expression of an existing attribute
// of the given name or adds a new attribute definition to the end of the block.
//
// The value is given as a cty.Value, and must therefore be a literal. To set
// a variable reference or other traversal, use SetAttributeTraversal.
//
// The return value is the attribute that was either modified in-place or
// created.
func (b *Body) SetAttributeValue(name string, val cty.Value) *Attribute {
	return b.setAttribute(name, NewExpressionLiteral(val))
}

// SetAttributeTraversal either replaces the expression of an existing attribute
// of the given name or adds a new attribute definition to the end of the body.
//
// The new expression is given as a hcl.Traversal, which must be an absolute
// traversal. To set a literal value, use SetAttributeValue.
//
// The return value is the attribute that was either modified in-place or
// created.
func (b *Body) SetAttributeTraversal(name string, traversal hcl.Traversal) *Attribute {
	return b.setAttribute(name, NewExpressionAbsTraversal(traversal))
}

// setAttribute either replaces the expression of an existing attribute of the
// given name or adds a new attribute definition to the end of the body. It
// returns the attribute that was either modified in-place or created.
func (b *Body) setAttribute(name string, expr *Expression) *Attribute {
	attr := b.GetAttribute(name)
	if attr != nil {
		attr.expr = attr.expr.ReplaceWith(expr)
	} else {
		attr = newAttribute()
		attr.init(name, expr)
		b.appendItem(attr)
	}
	return attr
}

// RemoveAttribute removes the attribute with the given name from the body.
//
// The return value is the attribute that was removed, or nil if there was
// no such attribute (in which case the call was a no-op).
func (b *Body) RemoveAttribute(name string) *Attribute {
	node := b.getAttributeNode(name)
	if node == nil {
		return nil
	}
	node.Detach()
	b.items.Remove(node)
	return node.content.(*Attribute)
}

// AppendBlock appends an existing block (which must not be already attached
// to a body) to the end of the receiving body.
func (b *Body) AppendBlock(block *Block) *Block {
	b.appendItem(block)
	return block
}

// AppendNewBlock appends a new nested block to the end of the receiving body
// with the given type name and labels.
func (b *Body) AppendNewBlock(typeName string, labels []string) *Block {
	block := newBlock()
	block.init(typeName, labels)
	b.appendItem(block)
	return block
}

// AppendNewline appends a newline token to the end of the receiving body,
// which generally serves as a separator between different sets of body
// contents.
func (b *Body) AppendNewline() {
	b.AppendUnstructuredTokens(Tokens{
		{
			Type:  hclsyntax.TokenNewline,
			Bytes: []byte{'\n'},
		},
	})
}

```

#### pkg/ast_expression.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

type Expression struct {
	inTree

	absTraversals nodeSet
	wrapped       *node
}

func newExpression() *Expression {
	return &Expression{
		inTree:        newInTree(),
		absTraversals: newNodeSet(),
	}
}

func (e *Expression) wrap(c nodeContent) {
	if e.wrapped != nil {
		panic("wrap called on already-wrapped Expression")
	}

	e.wrapped = newNode(c)
	e.children.AppendNode(e.wrapped)
}

// NewExpressionRaw constructs an expression containing the given raw tokens.
//
// There is no automatic validation that the given tokens produce a valid
// expression. Callers of thus function must take care to produce invalid
// expression tokens. Where possible, use the higher-level functions
// NewExpressionLiteral or NewExpressionAbsTraversal instead.
//
// Because NewExpressionRaw does not interpret the given tokens in any way,
// an expression created by NewExpressionRaw will produce an empty result
// for calls to its method Variables, even if the given token sequence
// contains a subslice that would normally be interpreted as a traversal under
// parsing.
func NewExpressionRaw(tokens Tokens) *Expression {
	expr := newExpression()
	// We copy the tokens here in order to make sure that later mutations
	// by the caller don't inadvertently cause our expression to become
	// invalid.
	copyTokens := make(Tokens, len(tokens))
	copy(copyTokens, tokens)
	expr.children.AppendUnstructuredTokens(copyTokens)
	return expr
}

// NewExpressionLiteral constructs an an expression that represents the given
// literal value.
//
// Since an unknown value cannot be represented in source code, this function
// will panic if the given value is unknown or contains a nested unknown value.
// Use val.IsWhollyKnown before calling to be sure.
//
// HCL native syntax does not directly represent lists, maps, and sets, and
// instead relies on the automatic conversions to those collection types from
// either list or tuple constructor syntax. Therefore converting collection
// values to source code and re-reading them will lose type information, and
// the reader must provide a suitable type at decode time to recover the
// original value.
func NewExpressionLiteral(val cty.Value) *Expression {
	toks := TokensForValue(val)
	expr := newExpression()
	expr.children.AppendUnstructuredTokens(toks)
	return expr
}

// NewExpressionAbsTraversal constructs an expression that represents the
// given traversal, which must be absolute or this function will panic.
func NewExpressionAbsTraversal(traversal hcl.Traversal) *Expression {
	if traversal.IsRelative() {
		panic("can't construct expression from relative traversal")
	}

	physT := newTraversal()
	rootName := traversal.RootName()
	steps := traversal[1:]

	{
		tn := newTraverseName()
		tn.name = tn.children.Append(newIdentifier(&Token{
			Type:  hclsyntax.TokenIdent,
			Bytes: []byte(rootName),
		}))
		physT.steps.Add(physT.children.Append(tn))
	}

	for _, step := range steps {
		switch ts := step.(type) {
		case hcl.TraverseAttr:
			tn := newTraverseName()
			tn.children.AppendUnstructuredTokens(Tokens{
				{
					Type:  hclsyntax.TokenDot,
					Bytes: []byte{'.'},
				},
			})
			tn.name = tn.children.Append(newIdentifier(&Token{
				Type:  hclsyntax.TokenIdent,
				Bytes: []byte(ts.Name),
			}))
			physT.steps.Add(physT.children.Append(tn))
		case hcl.TraverseIndex:
			ti := newTraverseIndex()
			ti.children.AppendUnstructuredTokens(Tokens{
				{
					Type:  hclsyntax.TokenOBrack,
					Bytes: []byte{'['},
				},
			})
			indexExpr := NewExpressionLiteral(ts.Key)
			ti.key = ti.children.Append(indexExpr)
			ti.children.AppendUnstructuredTokens(Tokens{
				{
					Type:  hclsyntax.TokenCBrack,
					Bytes: []byte{']'},
				},
			})
			physT.steps.Add(physT.children.Append(ti))
		}
	}

	expr := newExpression()
	expr.absTraversals.Add(expr.children.Append(physT))
	return expr
}

// Variables returns the absolute traversals that exist within the receiving
// expression, in source order.
//
// This includes the traversals nested in the arguments of a function call,
// the elements of a tuple-construct expression, and the keys and values of
// the items of an object-construct expression. A naked identifier used as an
// object key is a literal name rather than a variable, so it is not included.
func (e *Expression) Variables() []*Traversal {
	nodes := e.traversalNodes()
	ret := make([]*Traversal, len(nodes))
	for i, node := range nodes {
		ret[i] = node.content.(*Traversal)
	}
	return ret
}

// nestedTraversals is implemented by the structured content that an
// Expression can wrap and that contains nested expressions, such as an
// object-construct expression.
type nestedTraversals interface {
	// traversalNodes returns the nodes of the absolute traversals in the
	// nested expressions, in source order.
	traversalNodes() []*node
}

// traversalNodes returns the nodes of the absolute traversals within the
// receiving expression, including those in nested expressions, in source
// order. Each returned node has content of type *Traversal.
func (e *Expression) traversalNodes() []*node {
	ret := e.absTraversals.List()
	if e.wrapped != nil {
		if nested, ok := e.wrapped.content.(nestedTraversals); ok {
			ret = append(ret, nested.traversalNodes()...)
		}
	}
	return ret
}

// exprsTraversalNodes returns the nodes of the absolute traversals within
// the given expressions, in source order. Each node in the set must have
// content of type *Expression.
func exprsTraversalNodes(exprs nodeSet) []*node {
	var ret []*node
	for _, n := range exprs.List() {
		ret = append(ret, n.content.(*Expression).traversalNodes()...)
	}
	return ret
}

// RenameVariablePrefix examines each of the absolute traversals in the
// receiving expression to see if they have the given sequence of names as
// a prefix prefix. If so, they are updated in place to have the given
// replacement names instead of that prefix.
//
// This can be used to implement symbol renaming. The calling application can
// visit all relevant expressions in its input and apply the same renaming
// to implement a global symbol rename.
//
// The search and replacement traversals must be the same length, or this
// method will panic. Only attribute access operations can be matched and
// replaced. Index steps never match the prefix.
//
// The traversals examined are the same ones that Variables returns.
func (e *Expression) RenameVariablePrefix(search, replacement []string) {
	if len(search) != len(replacement) {
		panic(fmt.Sprintf("search and replacement length mismatch (%d and %d)", len(search), len(replacement)))
	}
	for _, node := range e.traversalNodes() {
		node.content.(*Traversal).renamePrefix(search, replacement)
	}
}

// AsObjectConsExpr returns the object constructor expression that the
// receiver consists of, or nil if the receiver is any other kind of
// expression.
func (e *Expression) AsObjectConsExpr() *ObjectConsExpr {
	return wrappedAs[*ObjectConsExpr](e)
}

// AsFunctionCallExpr returns the wrapped function call expression, or nil if
// this expression does not parse as a function call expression.
func (e *Expression) AsFunctionCallExpr() *FunctionCallExpr {
	return wrappedAs[*FunctionCallExpr](e)
}

// AsTupleConsExpr returns the wrapped tuple-construct expression, or nil if
// this expression does not parse as a tuple-construct expression.
func (e *Expression) AsTupleConsExpr() *TupleConsExpr {
	return wrappedAs[*TupleConsExpr](e)
}

func (e *Expression) asQuoted() *quoted {
	return wrappedAs[*quoted](e)
}

// wrappedAs returns the structured content that the expression wraps if it
// has type T, or the zero value of T otherwise.
func wrappedAs[T nodeContent](e *Expression) T {
	var zero T
	if e.wrapped == nil {
		return zero
	}
	if c, ok := e.wrapped.content.(T); ok {
		return c
	}
	return zero
}

func (e *Expression) AsQuotedLiteral() Tokens {
	if quoted := e.asQuoted(); quoted == nil {
		return Tokens{}
	} else {
		return quoted.tokens
	}
}

// Traversal represents a sequence of variable, attribute, and/or index
// operations.
type Traversal struct {
	inTree

	steps nodeSet
}

func newTraversal() *Traversal {
	return &Traversal{
		inTree: newInTree(),
		steps:  newNodeSet(),
	}
}

// renamePrefix replaces the leading names of the traversal with the names in
// replacement, if those leading names equal the names in search. Only name
// steps can match. The caller must make sure that search and replacement have
// the same length.
func (t *Traversal) renamePrefix(search, replacement []string) {
	if len(t.steps) < len(search) {
		// If it's shorter then it can't have our prefix
		return
	}

	stepNodes := t.steps.List()
	for i, name := range search {
		step, isName := stepNodes[i].content.(*TraverseName)
		if !isName {
			return // only name nodes can match
		}
		if string(step.name.content.(*identifier).token.Bytes) != name {
			return
		}
	}

	// If we get here then the prefix matched, so now we'll swap in
	// the replacement strings.
	for i, name := range replacement {
		step := stepNodes[i].content.(*TraverseName)
		token := step.name.content.(*identifier).token
		token.Bytes = []byte(name)
	}
}

type TraverseName struct {
	inTree

	name *node
}

func newTraverseName() *TraverseName {
	return &TraverseName{
		inTree: newInTree(),
	}
}

type TraverseIndex struct {
	inTree

	key *node
}

func newTraverseIndex() *TraverseIndex {
	return &TraverseIndex{
		inTree: newInTree(),
	}
}

```

#### pkg/ast_expression_test.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

func TestExpressionVariablesAndRename(t *testing.T) {
	tests := map[string]struct {
		src string
		// want lists the traversals that Variables returns, in order.
		want []string
		// wantRenamed is src after renaming the prefix "old" to "new".
		wantRenamed string
	}{
		"plain traversal": {
			src:         "a = old.x\n",
			want:        []string{"old.x"},
			wantRenamed: "a = new.x\n",
		},
		"object item value": {
			src:         "a = { k = old.x }\n",
			want:        []string{"old.x"},
			wantRenamed: "a = { k = new.x }\n",
		},
		"object items in source order": {
			src:         "a = { k = old.x, j = other.y, (old.z) = old.w }\n",
			want:        []string{"old.x", "other.y", "old.z", "old.w"},
			wantRenamed: "a = { k = new.x, j = other.y, (new.z) = new.w }\n",
		},
		"nested object": {
			src:         "a = { k = { j = old.x } }\n",
			want:        []string{"old.x"},
			wantRenamed: "a = { k = { j = new.x } }\n",
		},
		"tuple item value": {
			src:         "a = { k = [old.x] }\n",
			want:        []string{"old.x"},
			wantRenamed: "a = { k = [new.x] }\n",
		},
		"function call item value": {
			src:         "a = { k = f(old.x) }\n",
			want:        []string{"old.x"},
			wantRenamed: "a = { k = f(new.x) }\n",
		},
		"function call arguments": {
			src:         "a = f(old.x, other.y, old.z)\n",
			want:        []string{"old.x", "other.y", "old.z"},
			wantRenamed: "a = f(new.x, other.y, new.z)\n",
		},
		"tuple elements": {
			src:         "a = [old.x, other.y, old.z]\n",
			want:        []string{"old.x", "other.y", "old.z"},
			wantRenamed: "a = [new.x, other.y, new.z]\n",
		},
		"nested calls, tuples and objects in source order": {
			src:         "a = f(old.a, [old.b, g(old.c)], { k = [old.d] }, old.e...)\n",
			want:        []string{"old.a", "old.b", "old.c", "old.d", "old.e"},
			wantRenamed: "a = f(new.a, [new.b, g(new.c)], { k = [new.d] }, new.e...)\n",
		},
		"function call in object key": {
			src:         "a = { (f(old.x)) = [old.y] }\n",
			want:        []string{"old.x", "old.y"},
			wantRenamed: "a = { (f(new.x)) = [new.y] }\n",
		},
		"operation in function argument": {
			src:         "a = f(old.x + old.y)\n",
			want:        []string{"old.x", "old.y"},
			wantRenamed: "a = f(new.x + new.y)\n",
		},
		"naked identifier key": {
			src:         "a = { old = 1 }\n",
			want:        nil,
			wantRenamed: "a = { old = 1 }\n",
		},
		"naked identifier key with traversal value": {
			src:         "a = { old = old.x }\n",
			want:        []string{"old.x"},
			wantRenamed: "a = { old = new.x }\n",
		},
		"parenthesized key": {
			src:         "a = { (old.x) = 1 }\n",
			want:        []string{"old.x"},
			wantRenamed: "a = { (new.x) = 1 }\n",
		},
		"parenthesized single identifier key": {
			src:         "a = { (old) = 1 }\n",
			want:        []string{"old"},
			wantRenamed: "a = { (new) = 1 }\n",
		},
		"quoted key": {
			src:         "a = { \"k\" = old.x }\n",
			want:        []string{"old.x"},
			wantRenamed: "a = { \"k\" = new.x }\n",
		},
		"quoted key without traversal": {
			src:         "a = { \"old\" = 1 }\n",
			want:        nil,
			wantRenamed: "a = { \"old\" = 1 }\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			expr := f.Body().GetAttribute("a").Expr()

			var got []string
			for _, traversal := range expr.Variables() {
				got = append(got, strings.TrimSpace(string(traversal.BuildTokens(nil).Bytes())))
			}
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("wrong variables (-want +got):\n%s", diff)
			}

			nativeFile, diags := hclsyntax.ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected hclsyntax diagnostics: %s", diags.Error())
			}
			nativeAttrs, _ := nativeFile.Body.JustAttributes()
			if got, want := len(got), len(nativeAttrs["a"].Expr.Variables()); got != want {
				t.Errorf("got %d variables, but hclsyntax reports %d", got, want)
			}

			expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
			if got := string(f.Bytes()); got != test.wantRenamed {
				t.Errorf("wrong result after rename\ngot:  %q\nwant: %q", got, test.wantRenamed)
			}
		})
	}
}

func TestExpressionAsFunctionCallAndTupleConsExpr(t *testing.T) {
	for _, src := range []string{
		`a = x`,
		`a = "fez"`,
		`a = { k = f(1) }`,
		`a = f(1)[0]`,
		`a = (f(1))`,
		`a = [for x in y : x]`,
		`a = f(1) + 1`,
	} {
		t.Run(src, func(t *testing.T) {
			_, expr := parseTestExpr(t, src+"\n")
			if expr.AsFunctionCallExpr() != nil {
				t.Errorf("AsFunctionCallExpr returned non-nil")
			}
			if expr.AsTupleConsExpr() != nil {
				t.Errorf("AsTupleConsExpr returned non-nil")
			}
		})
	}
	if NewExpressionRaw(Tokens{}).AsFunctionCallExpr() != nil {
		t.Errorf("AsFunctionCallExpr returned non-nil for a raw expression")
	}
}

```

#### pkg/ast_function_call_expr.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

// FunctionCallExpr represents a function call expression such as
// join(" ", list). Use Expression.AsFunctionCallExpr to get the
// FunctionCallExpr of an expression.
//
// Each argument is its own Expression. The function name, the parentheses,
// the commas between the arguments and any expansion symbol after the last
// argument are tokens of the FunctionCallExpr itself.
type FunctionCallExpr struct {
	inTree

	name string

	// args holds the nodes of the arguments, in source order. Each node has
	// content of type *Expression.
	args nodeSet
}

func newFunctionCallExpr(name string) *FunctionCallExpr {
	return &FunctionCallExpr{
		inTree: newInTree(),
		name:   name,
		args:   newNodeSet(),
	}
}

// Name returns the name of the called function. A namespaced name, such as
// provider::example::upper, is returned in full.
func (call *FunctionCallExpr) Name() string {
	return call.name
}

// traversalNodes returns the nodes of the absolute traversals in the
// arguments of the call, in source order.
func (call *FunctionCallExpr) traversalNodes() []*node {
	return exprsTraversalNodes(call.args)
}

```

#### pkg/ast_function_call_expr_test.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/hcl/v2"
)

// parseTestExpr parses src and returns the expression of its attribute "a".
func parseTestExpr(t *testing.T, src string) (*File, *Expression) {
	t.Helper()
	f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics: %s", diags.Error())
	}
	return f, f.Body().GetAttribute("a").Expr()
}

// checkRoundTrip checks that the tokens of f are exactly src, and that
// File.Bytes returns src too if src is already formatted.
func checkRoundTrip(t *testing.T, f *File, src string) {
	t.Helper()
	if got := string(f.children.BuildTokens(nil).Bytes()); got != src {
		t.Errorf("wrong tokens\ngot:  %q\nwant: %q", got, src)
	}
	if string(Format([]byte(src))) != src {
		return
	}
	if got := string(f.Bytes()); got != src {
		t.Errorf("wrong File.Bytes result\ngot:  %q\nwant: %q", got, src)
	}
}

// exprsStrings returns the source of each expression in the given set,
// without surrounding spaces.
func exprsStrings(exprs nodeSet) []string {
	var ret []string
	for _, n := range exprs.List() {
		ret = append(ret, tokensString(n.content.(*Expression).BuildTokens(nil)))
	}
	return ret
}

func TestFunctionCallExprParse(t *testing.T) {
	tests := map[string]struct {
		src      string
		wantName string
		wantArgs []string
	}{
		"no arguments":      {`a = f()`, "f", nil},
		"one argument":      {`a = upper("fez")`, "upper", []string{`"fez"`}},
		"nested tuple":      {`a = join(" ", ["fedora", "fez"])`, "join", []string{`" "`, `["fedora", "fez"]`}},
		"trailing comma":    {`a = f(1, 2,)`, "f", []string{"1", "2"}},
		"expand final":      {`a = f(1, list...)`, "f", []string{"1", "list"}},
		"namespaced":        {`a = provider::example::upper("fez")`, "provider::example::upper", []string{`"fez"`}},
		"operation":         {`a = f(1 + 2, x)`, "f", []string{"1 + 2", "x"}},
		"nested object":     {`a = f({ k = 1 })`, "f", []string{"{ k = 1 }"}},
		"multi-line":        {"a = f(\n  1, # one\n  2,\n)", "f", []string{"1", "2"}},
		"unformatted":       {"a = f (  1,2 , )", "f", []string{"1", "2"}},
		"template argument": {`a = f("${x}")`, "f", []string{`"${x}"`}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			src := test.src + "\n"
			f, expr := parseTestExpr(t, src)

			call := expr.AsFunctionCallExpr()
			if call == nil {
				t.Fatalf("expression is not a function call")
			}
			if expr.AsTupleConsExpr() != nil || expr.AsObjectConsExpr() != nil {
				t.Errorf("function call is also reported as another kind of expression")
			}
			if got := call.Name(); got != test.wantName {
				t.Errorf("wrong name %q; want %q", got, test.wantName)
			}
			if diff := cmp.Diff(test.wantArgs, exprsStrings(call.args)); diff != "" {
				t.Errorf("wrong arguments (-want +got):\n%s", diff)
			}
			checkRoundTrip(t, f, src)
		})
	}
}

```

#### pkg/ast_object_cons_expr.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

// ObjectConsExpr is an object constructor expression, such as
// { hat = "derby", (cat) = "calico" }. Use Expression.AsObjectConsExpr to get
// the ObjectConsExpr of an expression.
//
// Its methods look up items by key and set item values, in a similar way to
// how the methods of Body look up and set attributes.
type ObjectConsExpr struct {
	inTree

	items nodeSet

	// close is the node holding the tokens after the last item, up to and
	// including the closing brace.
	close *node
}

func newObjectConsExpr() *ObjectConsExpr {
	return &ObjectConsExpr{
		inTree: newInTree(),
		items:  newNodeSet(),
	}
}

// Items returns the items of the object, in source order.
func (o *ObjectConsExpr) Items() []*ObjectConsItem {
	nodes := o.items.List()
	ret := make([]*ObjectConsItem, len(nodes))
	for i, n := range nodes {
		ret[i] = n.content.(*ObjectConsItem)
	}
	return ret
}

// ItemFor returns the first item whose key evaluates to the given string, or
// nil if there is no such item.
//
// Only keys whose value is known without evaluating an expression are
// considered: a naked identifier such as hat, or a quoted string literal such
// as "hat". The two forms are equivalent, and a quoted key is compared after
// its escape sequences are decoded. Keys that are expressions, such as (hat),
// hat.band or "${hat}", never match.
func (o *ObjectConsExpr) ItemFor(key string) *ObjectConsItem {
	for _, item := range o.Items() {
		if item.KeyExpr().hasName(key) {
			return item
		}
	}
	return nil
}

// ValueExprFor returns the value of the item that ItemFor returns for the
// given key, or nil if there is no such item.
func (o *ObjectConsExpr) ValueExprFor(key string) *ObjectConsValue {
	item := o.ItemFor(key)
	if item == nil {
		return nil
	}
	return item.ValueExpr()
}

// SetItem either replaces the value of the item that ItemFor returns for the
// given key, or adds a new item with the given key to the end of the object.
// The key of an existing item is left as written.
//
// A new key is written as a naked identifier if it is a valid identifier, and
// otherwise as a quoted string literal that evaluates to the given key.
//
// The given expression must not already be part of a syntax tree, or this
// method will panic.
//
// The return values are the key and the value of the item that was either
// modified in-place or created.
func (o *ObjectConsExpr) SetItem(key string, expr *Expression) (*ObjectConsKeyExpr, *ObjectConsValue) {
	expr.assertUnattached()

	item := o.ItemFor(key)
	if item != nil {
		item.ValueExpr().setExpr(expr)
	} else {
		item = o.appendItem(newObjectConsKeyExprForName(key), expr)
	}
	return item.KeyExpr(), item.ValueExpr()
}

// SetItemRaw is like SetItem, but uses the given tokens verbatim as the value.
//
// The same caveats apply to this function as for NewExpressionRaw on which
// it is based. If possible, prefer to use SetItemValue or SetItemTraversal.
func (o *ObjectConsExpr) SetItemRaw(key string, tokens Tokens) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionRaw(tokens))
}

// SetItemValue is like SetItem, but sets the value to the given literal, as
// NewExpressionLiteral does. To set a variable reference or other traversal,
// use SetItemTraversal.
func (o *ObjectConsExpr) SetItemValue(key string, val cty.Value) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionLiteral(val))
}

// SetItemTraversal is like SetItem, but sets the value to the given traversal,
// which must be absolute, as NewExpressionAbsTraversal does. To set a literal
// value, use SetItemValue.
func (o *ObjectConsExpr) SetItemTraversal(key string, traversal hcl.Traversal) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionAbsTraversal(traversal))
}

// appendItem adds a new item at the end of the object, laid out in the style
// of the object.
//
// In an object written on a single line, the new item goes on the same line,
// separated from the previous item by a comma. Otherwise, it goes on its own
// line. A comma or line comment that ends the line of the previous item stays
// with that item. Everything after that stays with the closing brace, which
// is moved to its own line if needed.
func (o *ObjectConsExpr) appendItem(key *ObjectConsKeyExpr, expr *Expression) *ObjectConsItem {
	lead, rest := splitLineEnd(o.close.content.(Tokens))
	if o.isMultiline() {
		if len(lead) == 0 || !isLineBreak(lead[len(lead)-1]) {
			lead = append(lead, newNewlineToken())
		}
		rest = append(Tokens{newNewlineToken()}, rest...)
	} else if len(o.items) > 0 && !hasComma(lead) {
		lead = append(lead, newCommaToken())
	}

	item := newObjectConsItem()
	item.init(lead, key, expr)
	o.items.Add(o.children.Insert(o.close, item))
	o.close = o.close.ReplaceWith(rest)
	return item
}

// isMultiline reports whether the object is written on more than one line.
// Only the tokens that belong to the object itself count, so line breaks
// inside a key or value, such as in a nested object, do not.
func (o *ObjectConsExpr) isMultiline() bool {
	for n := o.children.first; n != nil; n = n.after {
		switch c := n.content.(type) {
		case Tokens:
			if hasLineBreak(c) {
				return true
			}
		case *ObjectConsItem:
			if c.hasLineBreak() {
				return true
			}
		}
	}
	return false
}

// traversalNodes returns the nodes of the absolute traversals in the keys and
// values of the object's items, in source order.
func (o *ObjectConsExpr) traversalNodes() []*node {
	var ret []*node
	for _, item := range o.Items() {
		ret = append(ret, item.KeyExpr().traversalNodes()...)
		ret = append(ret, item.ValueExpr().Expr().traversalNodes()...)
	}
	return ret
}

// ObjectConsItem is one item of an ObjectConsExpr: a key and a value, such as
// hat = "derby".
type ObjectConsItem struct {
	inTree
	key   *node
	value *node
}

func newObjectConsItem() *ObjectConsItem {
	return &ObjectConsItem{
		inTree: newInTree(),
	}
}

// init fills a new item with the given key and value. The lead tokens are
// the separators that come before the key, such as a comma or a newline.
func (item *ObjectConsItem) init(lead Tokens, key *ObjectConsKeyExpr, expr *Expression) {
	item.children.AppendUnstructuredTokens(lead)
	item.key = item.children.Append(key)
	item.children.AppendUnstructuredTokens(Tokens{
		{
			Type:         hclsyntax.TokenEqual,
			Bytes:        []byte{'='},
			SpacesBefore: 1,
		},
	})
	value := newObjectConsValue()
	value.expr = value.children.Append(expr)
	item.value = item.children.Append(value)
}

// KeyExpr returns the key of the item.
func (item *ObjectConsItem) KeyExpr() *ObjectConsKeyExpr {
	return item.key.content.(*ObjectConsKeyExpr)
}

// ValueExpr returns the value of the item.
func (item *ObjectConsItem) ValueExpr() *ObjectConsValue {
	return item.value.content.(*ObjectConsValue)
}

// hasLineBreak reports whether the tokens of the item itself, outside of its
// key and value, contain a line break.
func (item *ObjectConsItem) hasLineBreak() bool {
	for n := item.children.first; n != nil; n = n.after {
		if toks, ok := n.content.(Tokens); ok && hasLineBreak(toks) {
			return true
		}
	}
	return false
}

// objectConsKeyKind is the syntactic form of an object-construct key.
type objectConsKeyKind int

const (
	// objectConsKeyExpression is a key that is evaluated as an expression,
	// such as (hat), hat.band or "${hat}".
	objectConsKeyExpression objectConsKeyKind = iota

	// objectConsKeyLiteralName is a naked identifier, such as hat. It is a
	// literal name, not a variable.
	objectConsKeyLiteralName

	// objectConsKeyQuotedLiteral is a quoted string literal, such as "hat".
	objectConsKeyQuotedLiteral
)

// ObjectConsKeyExpr is the key of an ObjectConsItem, such as hat in
// hat = "derby".
type ObjectConsKeyExpr struct {
	inTree

	wrapped *node

	kind objectConsKeyKind

	// name is the string that the key evaluates to, with any escape
	// sequences decoded. It is set only if kind is not
	// objectConsKeyExpression.
	name string
}

// newObjectConsKeyExpr wraps the given node, which must have content of type
// *Expression.
func newObjectConsKeyExpr(wrapped *node, kind objectConsKeyKind, name string) *ObjectConsKeyExpr {
	expr := &ObjectConsKeyExpr{
		inTree:  newInTree(),
		wrapped: wrapped,
		kind:    kind,
		name:    name,
	}
	expr.children.AppendNode(wrapped)

	return expr
}

// newObjectConsKeyExprForName returns a new key that evaluates to the given
// name. The key is a naked identifier if the name is a valid identifier, and
// a quoted string literal otherwise.
func newObjectConsKeyExprForName(name string) *ObjectConsKeyExpr {
	var tokens Tokens
	var kind objectConsKeyKind
	// A naked "for" as the first key would start a for expression instead of
	// an object, so it is always quoted.
	if hclsyntax.ValidIdentifier(name) && name != "for" {
		tokens = TokensForIdentifier(name)
		kind = objectConsKeyLiteralName
	} else {
		tokens = TokensForValue(cty.StringVal(name))
		kind = objectConsKeyQuotedLiteral
	}
	tokens[0].SpacesBefore = 1

	return newObjectConsKeyExpr(newNode(NewExpressionRaw(tokens)), kind, name)
}

// hasName reports whether the key is a literal name or a quoted string
// literal that evaluates to the given name.
func (k *ObjectConsKeyExpr) hasName(name string) bool {
	return k.kind != objectConsKeyExpression && k.name == name
}

// traversalNodes returns the nodes of the absolute traversals in the key, in
// source order. A literal name key has none.
func (k *ObjectConsKeyExpr) traversalNodes() []*node {
	if k.kind == objectConsKeyLiteralName {
		return nil
	}
	return k.wrapped.content.(*Expression).traversalNodes()
}

// ObjectConsValue is the value of an ObjectConsItem, such as "derby" in
// hat = "derby".
type ObjectConsValue struct {
	inTree
	expr *node
}

func newObjectConsValue() *ObjectConsValue {
	return &ObjectConsValue{
		inTree: newInTree(),
	}
}

// Expr returns the expression of the value.
func (o *ObjectConsValue) Expr() *Expression {
	return o.expr.content.(*Expression)
}

// setExpr replaces the expression of the value with the given one, which must
// not be part of a syntax tree.
func (o *ObjectConsValue) setExpr(expr *Expression) {
	o.expr = o.expr.ReplaceWith(expr)
}

// splitLineEnd splits the tokens that follow the last item of an object into
// the tokens that end the line of that item, and the rest.
//
// The line end is an optional comma and any inline comments, up to and
// including a newline or a line comment. If there is no line break before
// the next token that is not a comma or a comment, the line end is just a
// leading comma, if any.
//
// The returned line end can be appended to without changing the rest.
func splitLineEnd(tokens Tokens) (lineEnd, rest Tokens) {
	for i, tok := range tokens {
		if isLineBreak(tok) {
			return tokens[: i+1 : i+1], tokens[i+1:]
		}
		if tok.Type != hclsyntax.TokenComma && tok.Type != hclsyntax.TokenComment {
			break
		}
	}
	if len(tokens) > 0 && tokens[0].Type == hclsyntax.TokenComma {
		return tokens[:1:1], tokens[1:]
	}
	return nil, tokens
}

// isLineBreak reports whether the token ends a line: a newline, or a line
// comment, which includes its trailing newline.
func isLineBreak(tok *Token) bool {
	switch tok.Type {
	case hclsyntax.TokenNewline:
		return true
	case hclsyntax.TokenComment:
		return len(tok.Bytes) > 0 && tok.Bytes[len(tok.Bytes)-1] == '\n'
	default:
		return false
	}
}

func hasLineBreak(tokens Tokens) bool {
	for _, tok := range tokens {
		if isLineBreak(tok) {
			return true
		}
	}
	return false
}

func hasComma(tokens Tokens) bool {
	for _, tok := range tokens {
		if tok.Type == hclsyntax.TokenComma {
			return true
		}
	}
	return false
}

func newNewlineToken() *Token {
	return &Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte{'\n'},
	}
}

func newCommaToken() *Token {
	return &Token{
		Type:  hclsyntax.TokenComma,
		Bytes: []byte{','},
	}
}

```

#### pkg/ast_object_cons_expr_test.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

// parseTestObject parses src and returns the object constructor that is
// the value of its attribute "a".
func parseTestObject(t *testing.T, src string) (*File, *ObjectConsExpr) {
	t.Helper()
	f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics: %s", diags.Error())
	}
	obj := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()
	if obj == nil {
		t.Fatalf("attribute a is not an object constructor")
	}
	return f, obj
}

// evalAttribute parses src with hclsyntax and evaluates its attribute "a".
func evalAttribute(t *testing.T, src []byte) cty.Value {
	t.Helper()
	f, diags := hclsyntax.ParseConfig(src, "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("result is not valid HCL: %s\n%s", diags.Error(), src)
	}
	attrs, diags := f.Body.JustAttributes()
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics: %s", diags.Error())
	}
	val, diags := attrs["a"].Expr.Value(nil)
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics evaluating a: %s\n%s", diags.Error(), src)
	}
	return val
}

func tokensString(toks Tokens) string {
	return strings.TrimSpace(string(toks.Bytes()))
}

func TestObjectConsExprItemFor(t *testing.T) {
	tests := map[string]struct {
		src string
		key string
		// want is the value of the item found, or "" if none is found.
		want string
	}{
		"identifier":                    {`a = { hat = 1 }`, "hat", "1"},
		"quoted":                        {`a = { "hat" = 1 }`, "hat", "1"},
		"colon":                         {`a = { hat: 1 }`, "hat", "1"},
		"second item":                   {`a = { cat = 1, hat = 2 }`, "hat", "2"},
		"first of duplicates":           {`a = { hat = 1, "hat" = 2 }`, "hat", "1"},
		"quoted with space":             {`a = { "a hat" = 1 }`, "a hat", "1"},
		"quoted with escaped template":  {`a = { "$${hat}" = 1 }`, "${hat}", "1"},
		"quoted with escaped directive": {`a = { "%%{hat}" = 1 }`, "%{hat}", "1"},
		"quoted with escaped quote":     {`a = { "say \"hi\"" = 1 }`, `say "hi"`, "1"},
		"quoted with escaped backslash": {`a = { "a\\b" = 1 }`, `a\b`, "1"},
		"quoted with unicode escape":    {`a = { "h\u0061t" = 1 }`, "hat", "1"},
		"quoted empty":                  {`a = { "" = 1 }`, "", "1"},
		"keyword":                       {`a = { null = 1 }`, "null", "1"},
		"multi-line":                    {"a = {\n  cat = 1\n  hat = 2\n}", "hat", "2"},
		"nested object value":           {`a = { hat = { b = 1 } }`, "hat", "{ b = 1 }"},

		"missing":                   {`a = { cat = 1 }`, "hat", ""},
		"empty object":              {`a = {}`, "hat", ""},
		"escaped template raw":      {`a = { "$${hat}" = 1 }`, "$${hat}", ""},
		"parenthesized":             {`a = { (hat) = 1 }`, "hat", ""},
		"parenthesized quoted":      {`a = { ("hat") = 1 }`, "hat", ""},
		"traversal":                 {`a = { hat.band = 1 }`, "hat.band", ""},
		"traversal root":            {`a = { hat.band = 1 }`, "hat", ""},
		"template":                  {`a = { "${hat}" = 1 }`, "hat", ""},
		"template source":           {`a = { "${hat}" = 1 }`, "${hat}", ""},
		"function call":             {`a = { (upper("hat")) = 1 }`, "HAT", ""},
		"key of nested object":      {`a = { cat = { hat = 1 } }`, "hat", ""},
		"case sensitive":            {`a = { Hat = 1 }`, "hat", ""},
		"quoted without decoding":   {`a = { "a\\b" = 1 }`, `a\\b`, ""},
		"value is not a key":        {`a = { cat = "hat" }`, "hat", ""},
		"identifier is not a value": {`a = { cat = hat }`, "hat", ""},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, obj := parseTestObject(t, test.src+"\n")

			item := obj.ItemFor(test.key)
			value := obj.ValueExprFor(test.key)
			if test.want == "" {
				if item != nil {
					t.Errorf("ItemFor found an item with key %q", tokensString(item.KeyExpr().BuildTokens(nil)))
				}
				if value != nil {
					t.Errorf("ValueExprFor found value %q", tokensString(value.BuildTokens(nil)))
				}
				return
			}

			if item == nil {
				t.Fatalf("ItemFor found no item")
			}
			if got := tokensString(item.ValueExpr().Expr().BuildTokens(nil)); got != test.want {
				t.Errorf("ItemFor found wrong item with value %q; want %q", got, test.want)
			}
			if value != item.ValueExpr() {
				t.Errorf("ValueExprFor returned a different value than ItemFor")
			}
		})
	}
}

func TestObjectConsExprItems(t *testing.T) {
	_, obj := parseTestObject(t, "a = { hat = 1, (cat) = 2, \"dog\" = 3 }\n")

	var got []string
	for _, item := range obj.Items() {
		got = append(got, tokensString(item.KeyExpr().BuildTokens(nil))+"="+tokensString(item.ValueExpr().BuildTokens(nil)))
	}
	want := []string{"hat=1", "(cat)=2", `"dog"=3`}
	if strings.Join(got, " ") != strings.Join(want, " ") {
		t.Errorf("wrong items\ngot:  %q\nwant: %q", got, want)
	}

	_, empty := parseTestObject(t, "a = {}\n")
	if got := empty.Items(); len(got) != 0 {
		t.Errorf("got %d items for an empty object", len(got))
	}
}

func TestObjectConsExprSetItemKey(t *testing.T) {
	tests := []struct {
		key  string
		want string
	}{
		{"hat", "hat"},
		{"hat-band", "hat-band"},
		{"hat_band", "hat_band"},
		{"null", "null"},
		{"true", "true"},
		{"if", "if"},
		{"for", `"for"`},
		{`say "hi"`, `"say \"hi\""`},
		{`a\b`, `"a\\b"`},
		{"${hat}", `"$${hat}"`},
		{"%{hat}", `"%%{hat}"`},
		{"$hat", `"$hat"`},
		{"", `""`},
		{"a hat", `"a hat"`},
		{"hat.band", `"hat.band"`},
		{"1hat", `"1hat"`},
		{"a\nb", `"a\nb"`},
		{"hat=", `"hat="`},
		{"=", `"="`},
		{"(hat)", `"(hat)"`},
		{"#", `"#"`},
		{"ハット", "ハット"},
	}

	// Each key is set on an empty object both on a single line and on
	// several lines, and in the first and in a later position.
	srcs := []string{
		"a = {}\n",
		"a = {\n}\n",
		"a = { cat = 1 }\n",
		"a = {\n  cat = 1\n}\n",
	}

	for _, test := range tests {
		for _, src := range srcs {
			t.Run(test.key+"/"+src, func(t *testing.T) {
				f, obj := parseTestObject(t, src)

				key, value := obj.SetItemValue(test.key, cty.StringVal("x"))
				if got := tokensString(key.BuildTokens(nil)); got != test.want {
					t.Errorf("wrong key\ngot:  %s\nwant: %s", got, test.want)
				}
				if got := tokensString(value.BuildTokens(nil)); got != `"x"` {
					t.Errorf("wrong value %s", got)
				}

				item := obj.ItemFor(test.key)
				if item == nil {
					t.Fatalf("ItemFor did not find the new item")
				}
				if item.KeyExpr() != key || item.ValueExpr() != value {
					t.Errorf("ItemFor found a different item than the one returned by SetItemValue")
				}

				result := f.Bytes()
				val := evalAttribute(t, result)
				if !val.Type().HasAttribute(test.key) {
					t.Fatalf("result has no key %q:\n%s", test.key, result)
				}
				if got := val.GetAttr(test.key); !got.RawEquals(cty.StringVal("x")) {
					t.Errorf("wrong value for key %q: %#v", test.key, got)
				}

				// The key must also be found after a round trip.
				_, reparsed := parseTestObject(t, string(result))
				if reparsed.ItemFor(test.key) == nil {
					t.Errorf("ItemFor did not find the key after parsing the result:\n%s", result)
				}
			})
		}
	}
}

func TestObjectConsExprSetItemExisting(t *testing.T) {
	tests := map[string]struct {
		src  string
		key  string
		want string
	}{
		"identifier": {
			"a = { hat = \"derby\", cat = \"calico\" }\n",
			"hat",
			"a = { hat = \"bowler\", cat = \"calico\" }\n",
		},
		"quoted key left as written": {
			"a = { \"hat\" = \"derby\" }\n",
			"hat",
			"a = { \"hat\" = \"bowler\" }\n",
		},
		"escaped key left as written": {
			"a = { \"$${hat}\" = \"derby\" }\n",
			"${hat}",
			"a = { \"$${hat}\" = \"bowler\" }\n",
		},
		"last item": {
			"a = { cat = \"calico\", hat = \"derby\" }\n",
			"hat",
			"a = { cat = \"calico\", hat = \"bowler\" }\n",
		},
		"multi-line with comments": {
			"a = {\n  # lead\n  hat = \"derby\" # line\n  cat = \"calico\"\n}\n",
			"hat",
			"a = {\n  # lead\n  hat = \"bowler\" # line\n  cat = \"calico\"\n}\n",
		},
		"nested object value": {
			"a = {\n  hat = {\n    style = \"derby\"\n  }\n}\n",
			"hat",
			"a = {\n  hat = \"bowler\"\n}\n",
		},
		"expression key not replaced": {
			"a = { (hat) = \"derby\" }\n",
			"hat",
			"a = { (hat) = \"derby\", hat = \"bowler\" }\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			f, obj := parseTestObject(t, test.src)
			before := len(obj.Items())
			existing := obj.ItemFor(test.key)

			key, value := obj.SetItemValue(test.key, cty.StringVal("bowler"))

			if existing != nil {
				if key != existing.KeyExpr() || value != existing.ValueExpr() {
					t.Errorf("SetItemValue did not return the existing item")
				}
				if got := len(obj.Items()); got != before {
					t.Errorf("got %d items; want %d", got, before)
				}
			}
			if got := tokensString(value.Expr().BuildTokens(nil)); got != `"bowler"` {
				t.Errorf("wrong value %s", got)
			}
			if got := string(f.Bytes()); got != test.want {
				t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, test.want)
			}
		})
	}
}

func TestObjectConsExprSetItemHelpers(t *testing.T) {
	f, obj := parseTestObject(t, "a = { hat = 1 }\n")

	obj.SetItemRaw("raw", Tokens{
		{Type: hclsyntax.TokenIdent, Bytes: []byte("upper")},
		{Type: hclsyntax.TokenOParen, Bytes: []byte("(")},
		{Type: hclsyntax.TokenOQuote, Bytes: []byte(`"`)},
		{Type: hclsyntax.TokenQuotedLit, Bytes: []byte("x")},
		{Type: hclsyntax.TokenCQuote, Bytes: []byte(`"`)},
		{Type: hclsyntax.TokenCParen, Bytes: []byte(")")},
	})
	obj.SetItemValue("value", cty.ListVal([]cty.Value{cty.NumberIntVal(1)}))
	_, traversal := obj.SetItemTraversal("traversal", hcl.Traversal{
		hcl.TraverseRoot{Name: "var"},
		hcl.TraverseAttr{Name: "hat"},
	})
	obj.SetItemTraversal("hat", hcl.Traversal{hcl.TraverseRoot{Name: "local"}})

	want := "a = { hat = local, raw = upper(\"x\"), value = [1], traversal = var.hat }\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result\ngot:  %s\nwant: %s", got, want)
	}

	// Traversals set by SetItemTraversal are variables of the object.
	expr := f.Body().GetAttribute("a").Expr()
	var vars []string
	for _, v := range expr.Variables() {
		vars = append(vars, tokensString(v.BuildTokens(nil)))
	}
	if got, want := strings.Join(vars, " "), "local var.hat"; got != want {
		t.Errorf("wrong variables %q; want %q", got, want)
	}
	if got := len(traversal.Expr().Variables()); got != 1 {
		t.Errorf("got %d variables in the new value; want 1", got)
	}

	expr.RenameVariablePrefix([]string{"var"}, []string{"input"})
	want = "a = { hat = local, raw = upper(\"x\"), value = [1], traversal = input.hat }\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result after rename\ngot:  %s\nwant: %s", got, want)
	}
}

func TestObjectConsExprSetItemPanics(t *testing.T) {
	assertPanics := func(t *testing.T, f func()) {
		t.Helper()
		defer func() {
			t.Helper()
			if recover() == nil {
				t.Errorf("did not panic")
			}
		}()
		f()
	}

	const src = "a = { hat = \"derby\" }\nb = \"calico\"\n"

	t.Run("parsed attribute expression", func(t *testing.T) {
		f, obj := parseTestObject(t, src)
		expr := f.Body().GetAttribute("b").Expr()
		assertPanics(t, func() { obj.SetItem("cat", expr) })
		assertPanics(t, func() { obj.SetItem("hat", expr) })
		if got, want := string(f.Bytes()), src; got != want {
			t.Errorf("result changed after panic\ngot:  %s\nwant: %s", got, want)
		}
	})

	t.Run("parsed item value", func(t *testing.T) {
		f, obj := parseTestObject(t, src)
		expr := obj.ValueExprFor("hat").Expr()
		assertPanics(t, func() { obj.SetItem("cat", expr) })
		assertPanics(t, func() { obj.SetItem("hat", expr) })
		if got, want := string(f.Bytes()), src; got != want {
			t.Errorf("result changed after panic\ngot:  %s\nwant: %s", got, want)
		}
	})

	t.Run("the object itself", func(t *testing.T) {
		f, obj := parseTestObject(t, src)
		expr := f.Body().GetAttribute("a").Expr()
		assertPanics(t, func() { obj.SetItem("cat", expr) })
	})

	t.Run("expression attached by an earlier call", func(t *testing.T) {
		_, obj := parseTestObject(t, src)
		expr := NewExpressionAbsTraversal(hcl.Traversal{hcl.TraverseRoot{Name: "hat"}})
		obj.SetItem("cat", expr)
		assertPanics(t, func() { obj.SetItem("dog", expr) })
		assertPanics(t, func() { obj.SetItem("hat", expr) })
		assertPanics(t, func() { obj.SetItem("cat", expr) })
	})

	t.Run("expression attached by an earlier replace", func(t *testing.T) {
		_, obj := parseTestObject(t, src)
		expr := NewExpressionAbsTraversal(hcl.Traversal{hcl.TraverseRoot{Name: "hat"}})
		obj.SetItem("hat", expr)
		assertPanics(t, func() { obj.SetItem("dog", expr) })
	})

	t.Run("replaced expression can be reused", func(t *testing.T) {
		f, obj := parseTestObject(t, src)
		old := obj.ValueExprFor("hat").Expr()
		obj.SetItemValue("hat", cty.StringVal("bowler"))
		obj.SetItem("cat", old)
		want := "a = { hat = \"bowler\", cat = \"derby\" }\nb = \"calico\"\n"
		if got := string(f.Bytes()); got != want {
			t.Errorf("wrong result\ngot:  %s\nwant: %s", got, want)
		}
	})
}

func TestObjectConsExprSetItemLayout(t *testing.T) {
	tests := map[string]struct {
		src  string
		keys []string
		want string
	}{
		"single line": {
			src:  "a = { hat = \"derby\" }\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\", dog = \"bowler\" }\n",
		},
		"single line two items": {
			src:  "a = { hat = \"derby\" }\n",
			keys: []string{"dog", "cat"},
			want: "a = { hat = \"derby\", dog = \"bowler\", cat = \"bowler\" }\n",
		},
		"empty": {
			src:  "a = {}\n",
			keys: []string{"dog"},
			want: "a = { dog = \"bowler\" }\n",
		},
		"empty two items": {
			src:  "a = {}\n",
			keys: []string{"dog", "cat"},
			want: "a = { dog = \"bowler\", cat = \"bowler\" }\n",
		},
		"empty with spaces": {
			src:  "a = { }\n",
			keys: []string{"dog"},
			want: "a = { dog = \"bowler\" }\n",
		},
		"single line trailing comma": {
			src:  "a = { hat = \"derby\", }\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\", dog = \"bowler\" }\n",
		},
		"single line trailing comma two items": {
			src:  "a = { hat = \"derby\", }\n",
			keys: []string{"dog", "cat"},
			want: "a = { hat = \"derby\", dog = \"bowler\", cat = \"bowler\" }\n",
		},
		"single line colon": {
			src:  "a = { hat: \"derby\" }\n",
			keys: []string{"dog"},
			want: "a = { hat : \"derby\", dog = \"bowler\" }\n",
		},
		"single line inline comment before brace": {
			src:  "a = { hat = \"derby\" /* end of a */ }\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\", dog = \"bowler\" /* end of a */ }\n",
		},
		"single line with nested multi-line object": {
			src:  "a = { hat = {\n  style = \"derby\"\n} }\n",
			keys: []string{"dog"},
			want: "a = { hat = {\n  style = \"derby\"\n}, dog = \"bowler\" }\n",
		},
		"single line with nested multi-line object last": {
			src:  "a = { cat = \"calico\", hat = {\n  style = \"derby\"\n  } }\n",
			keys: []string{"dog"},
			want: "a = { cat = \"calico\", hat = {\n  style = \"derby\"\n}, dog = \"bowler\" }\n",
		},
		"single line with multi-line tuple": {
			src:  "a = { hat = [\n  \"derby\",\n] }\n",
			keys: []string{"dog"},
			want: "a = { hat = [\n  \"derby\",\n], dog = \"bowler\" }\n",
		},
		"single line in block": {
			src:  "b {\n  a = { hat = \"derby\" }\n}\n",
			keys: []string{"dog"},
			want: "b {\n  a = { hat = \"derby\", dog = \"bowler\" }\n}\n",
		},
		"multi-line": {
			src:  "a = {\n  hat = \"derby\"\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}\n",
		},
		"multi-line two items": {
			src:  "a = {\n  hat = \"derby\"\n}\n",
			keys: []string{"dog", "cat"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n}\n",
		},
		"multi-line empty": {
			src:  "a = {\n}\n",
			keys: []string{"dog", "cat"},
			want: "a = {\n  dog = \"bowler\"\n  cat = \"bowler\"\n}\n",
		},
		"multi-line empty with comment": {
			src:  "a = {\n  # nothing yet\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  dog = \"bowler\"\n  # nothing yet\n}\n",
		},
		"multi-line brace on line of last item": {
			src:  "a = {\n    hat = \"derby\", (cat) = \"calico\" }\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\", (cat) = \"calico\"\n  dog = \"bowler\"\n}\n",
		},
		"multi-line brace on line of last item two items": {
			src:  "a = {\n  hat = \"derby\" }\n",
			keys: []string{"dog", "cat"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n}\n",
		},
		"multi-line first item on line of open brace": {
			src:  "a = { hat = \"derby\"\n}\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\"\n  dog = \"bowler\"\n}\n",
		},
		"multi-line trailing comma": {
			src:  "a = {\n  hat = \"derby\",\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\",\n  dog = \"bowler\"\n}\n",
		},
		"multi-line trailing comma brace on line of last item": {
			src:  "a = {\n  hat = \"derby\", }\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\",\n  dog = \"bowler\"\n}\n",
		},
		"multi-line line comment": {
			src:  "a = {\n  hat = \"derby\" # a fancy hat\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\" # a fancy hat\n  dog = \"bowler\"\n}\n",
		},
		"multi-line comma and line comment": {
			src:  "a = {\n  hat = \"derby\", // a fancy hat\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\", // a fancy hat\n  dog = \"bowler\"\n}\n",
		},
		"multi-line inline and line comment": {
			src:  "a = {\n  hat = \"derby\" /* fancy */ # a hat\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\" /* fancy */ # a hat\n  dog = \"bowler\"\n}\n",
		},
		"multi-line comments before brace": {
			src:  "a = {\n  hat = \"derby\"\n  # more hats\n  # later\n}\n",
			keys: []string{"dog", "cat"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n  # more hats\n  # later\n}\n",
		},
		"multi-line blank line and comment before brace": {
			src:  "a = {\n  hat = \"derby\"\n\n  # more hats\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n\n  # more hats\n}\n",
		},
		"multi-line inline comment before brace": {
			src:  "a = {\n  hat = \"derby\"\n  /* end of a */ }\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n/* end of a */ }\n",
		},
		"multi-line inline comment before brace on line of last item": {
			src:  "a = {\n  hat = \"derby\" /* end of a */ }\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n/* end of a */ }\n",
		},
		"multi-line with nested multi-line object last": {
			src:  "a = {\n  hat = {\n    style = \"derby\"\n  }\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = {\n    style = \"derby\"\n  }\n  dog = \"bowler\"\n}\n",
		},
		"multi-line in block": {
			src:  "b {\n  a = {\n    hat = \"derby\"\n  }\n}\n",
			keys: []string{"dog"},
			want: "b {\n  a = {\n    hat = \"derby\"\n    dog = \"bowler\"\n  }\n}\n",
		},
		"multi-line followed by attribute": {
			src:  "a = {\n  hat = \"derby\"\n}\nb = 1\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}\nb = 1\n",
		},
		"existing and new": {
			src:  "a = {\n  hat = \"derby\"\n}\n",
			keys: []string{"hat", "dog"},
			want: "a = {\n  hat = \"bowler\"\n  dog = \"bowler\"\n}\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			var obj *ObjectConsExpr
			if attr := f.Body().GetAttribute("a"); attr != nil {
				obj = attr.Expr().AsObjectConsExpr()
			} else {
				obj = f.Body().Blocks()[0].Body().GetAttribute("a").Expr().AsObjectConsExpr()
			}

			for _, key := range test.keys {
				obj.SetItemValue(key, cty.StringVal("bowler"))
			}

			got := f.Bytes()
			if string(got) != test.want {
				t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, test.want)
			}
			if formatted := Format(got); string(formatted) != string(got) {
				t.Errorf("result is not in canonical form\ngot:\n%s\nformatted:\n%s", got, formatted)
			}
			if _, diags := hclsyntax.ParseConfig(got, "", hcl.InitialPos); diags.HasErrors() {
				t.Errorf("result is not valid HCL: %s", diags.Error())
			}
			for _, key := range test.keys {
				if obj.ItemFor(key) == nil {
					t.Errorf("ItemFor(%q) found nothing", key)
				}
			}
		})
	}
}

```

#### pkg/ast_tuple_cons_expr.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

// TupleConsExpr represents a tuple-construct expression such as
// ["fedora", "fez"]. Use Expression.AsTupleConsExpr to get the TupleConsExpr
// of an expression.
//
// Each element is its own Expression. The brackets and the commas between
// the elements are tokens of the TupleConsExpr itself.
type TupleConsExpr struct {
	inTree

	// exprs holds the nodes of the elements, in source order. Each node has
	// content of type *Expression.
	exprs nodeSet
}

func newTupleConsExpr() *TupleConsExpr {
	return &TupleConsExpr{
		inTree: newInTree(),
		exprs:  newNodeSet(),
	}
}

// traversalNodes returns the nodes of the absolute traversals in the
// elements of the tuple, in source order.
func (tuple *TupleConsExpr) traversalNodes() []*node {
	return exprsTraversalNodes(tuple.exprs)
}

```

#### pkg/ast_tuple_cons_expr_test.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTupleConsExprParse(t *testing.T) {
	tests := map[string]struct {
		src       string
		wantExprs []string
	}{
		"empty":          {`a = []`, nil},
		"strings":        {`a = ["fedora", "fez"]`, []string{`"fedora"`, `"fez"`}},
		"trailing comma": {`a = [1, 2,]`, []string{"1", "2"}},
		"nested":         {`a = [[1], f(2), { k = 3 }]`, []string{"[1]", "f(2)", "{ k = 3 }"}},
		"multi-line":     {"a = [\n  # first\n  1,\n  2, // two\n]", []string{"1", "2"}},
		"traversals":     {`a = [x.y, z[0]]`, []string{"x.y", "z[0]"}},
		"unformatted":    {"a = [  1,2 , ]", []string{"1", "2"}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			src := test.src + "\n"
			f, expr := parseTestExpr(t, src)

			tuple := expr.AsTupleConsExpr()
			if tuple == nil {
				t.Fatalf("expression is not a tuple-construct expression")
			}
			if expr.AsFunctionCallExpr() != nil || expr.AsObjectConsExpr() != nil {
				t.Errorf("tuple is also reported as another kind of expression")
			}
			if diff := cmp.Diff(test.wantExprs, exprsStrings(tuple.exprs)); diff != "" {
				t.Errorf("wrong elements (-want +got):\n%s", diff)
			}
			checkRoundTrip(t, f, src)
		})
	}
}

```

#### pkg/examples_test.go

```go
// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package hclwrite_test

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func Example_generateFromScratch() {
	f := hclwrite.NewEmptyFile()
	rootBody := f.Body()
	rootBody.SetAttributeValue("string", cty.StringVal("bar")) // this is overwritten later
	rootBody.AppendNewline()
	rootBody.SetAttributeValue("object", cty.ObjectVal(map[string]cty.Value{
		"foo": cty.StringVal("foo"),
		"bar": cty.NumberIntVal(5),
		"baz": cty.True,
	}))
	rootBody.SetAttributeValue("string", cty.StringVal("foo"))
	rootBody.SetAttributeValue("bool", cty.False)
	rootBody.SetAttributeTraversal("path", hcl.Traversal{
		hcl.TraverseRoot{
			Name: "env",
		},
		hcl.TraverseAttr{
			Name: "PATH",
		},
	})
	rootBody.AppendNewline()
	fooBlock := rootBody.AppendNewBlock("foo", nil)
	fooBody := fooBlock.Body()
	rootBody.AppendNewBlock("empty", nil)
	rootBody.AppendNewline()
	barBlock := rootBody.AppendNewBlock("bar", []string{"a", "b"})
	barBody := barBlock.Body()

	fooBody.SetAttributeValue("hello", cty.StringVal("world"))

	bazBlock := barBody.AppendNewBlock("baz", nil)
	bazBody := bazBlock.Body()
	bazBody.SetAttributeValue("foo", cty.NumberIntVal(10))
	bazBody.SetAttributeValue("beep", cty.StringVal("boop"))
	bazBody.SetAttributeValue("baz", cty.ListValEmpty(cty.String))

	fmt.Printf("%s", f.Bytes())
	// Output:
	// string = "foo"
	//
	// object = {
	//   bar = 5
	//   baz = true
	//   foo = "foo"
	// }
	// bool = false
	// path = env.PATH
	//
	// foo {
	//   hello = "world"
	// }
	// empty {
	// }
	//
	// bar "a" "b" {
	//   baz {
	//     foo  = 10
	//     beep = "boop"
	//     baz  = []
	//   }
	// }
}

func ExampleExpression_RenameVariablePrefix() {
	src := []byte(
		"foo = a.x + a.y * b.c\n" +
			"bar = max(a.z, b.c)\n",
	)
	f, diags := hclwrite.ParseConfig(src, "", hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		fmt.Printf("errors: %s", diags)
		return
	}

	// Rename references of variable "a" to "z"
	for _, attr := range f.Body().Attributes() {
		attr.Expr().RenameVariablePrefix(
			[]string{"a"},
			[]string{"z"},
		)
	}

	fmt.Printf("%s", f.Bytes())
	// Output:
	// foo = z.x + z.y * b.c
	// bar = max(z.z, b.c)
}

func ExampleObjectConsExpr_SetItemValue() {
	src := []byte(`a = {
  hat = "derby", (cat) = "calico" }
`)

	f, diags := hclwrite.ParseConfig(src, "", hcl.InitialPos)
	if diags.HasErrors() {
		fmt.Printf("errors: %s", diags)
		return
	}
	obj := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

	// Replaces the value of the existing item "hat".
	obj.SetItemValue("hat", cty.StringVal("bowler"))
	// Adds a new item at the end, on its own line.
	obj.SetItemValue("dog", cty.StringVal("beagle"))

	fmt.Printf("%s", f.Bytes())
	// Output:
	// a = {
	//   hat = "bowler", (cat) = "calico"
	//   dog = "beagle"
	// }
}

```

#### pkg/node.go

```go
// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

// node represents a node in the AST.
type node struct {
	content nodeContent

	list          *nodes
	before, after *node
}

func newNode(c nodeContent) *node {
	n := &node{
		content: c,
	}
	if h, ok := c.(heldContent); ok {
		h.setHolder(n)
	}
	return n
}

func (n *node) Equal(other *node) bool {
	return cmp.Equal(n.content, other.content)
}

func (n *node) BuildTokens(to Tokens) Tokens {
	return n.content.BuildTokens(to)
}

// Detach removes the receiver from the list it currently belongs to. If the
// node is not currently in a list, this is a no-op.
func (n *node) Detach() {
	if n.list == nil {
		return
	}
	if n.before != nil {
		n.before.after = n.after
	}
	if n.after != nil {
		n.after.before = n.before
	}
	if n.list.first == n {
		n.list.first = n.after
	}
	if n.list.last == n {
		n.list.last = n.before
	}
	n.list = nil
	n.before = nil
	n.after = nil
}

// ReplaceWith removes the receiver from the list it currently belongs to and
// inserts a new node with the given content in its place. If the node is not
// currently in a list, this function will panic.
//
// The return value is the newly-constructed node, containing the given content.
// After this function returns, the reciever is no longer attached to a list.
func (n *node) ReplaceWith(c nodeContent) *node {
	if n.list == nil {
		panic("can't replace node that is not in a list")
	}

	before := n.before
	after := n.after
	list := n.list
	n.before, n.after, n.list = nil, nil, nil

	nn := newNode(c)
	nn.before = before
	nn.after = after
	nn.list = list
	if before != nil {
		before.after = nn
	} else {
		list.first = nn
	}
	if after != nil {
		after.before = nn
	} else {
		list.last = nn
	}
	return nn
}

func (n *node) assertUnattached() {
	if n.list != nil {
		panic(fmt.Sprintf("attempt to attach already-attached node %#v", n))
	}
}

// nodeContent is the interface type implemented by all AST content types.
type nodeContent interface {
	walkChildNodes(w internalWalkFunc)
	BuildTokens(to Tokens) Tokens
}

// nodes is a list of nodes.
type nodes struct {
	first, last *node
}

func (ns *nodes) BuildTokens(to Tokens) Tokens {
	for n := ns.first; n != nil; n = n.after {
		to = n.BuildTokens(to)
	}
	return to
}

func (ns *nodes) Clear() {
	ns.first = nil
	ns.last = nil
}

func (ns *nodes) Append(c nodeContent) *node {
	n := newNode(c)
	ns.AppendNode(n)
	return n
}

func (ns *nodes) AppendNode(n *node) {
	if ns.last != nil {
		n.before = ns.last
		ns.last.after = n
	}
	n.list = ns
	ns.last = n
	if ns.first == nil {
		ns.first = n
	}
}

// Insert inserts a nodeContent at a given position.
// This is just a wrapper for InsertNode. See InsertNode for details.
func (ns *nodes) Insert(pos *node, c nodeContent) *node {
	n := newNode(c)
	ns.InsertNode(pos, n)
	return n
}

// InsertNode inserts a node at a given position.
// The first argument is a node reference before which to insert.
// To insert it to an empty list, set position to nil.
func (ns *nodes) InsertNode(pos *node, n *node) {
	if pos == nil {
		// inserts n to empty list.
		ns.first = n
		ns.last = n
	} else {
		// inserts n before pos.
		if pos.before != nil {
			pos.before.after = n
		} else {
			ns.first = n
		}
		n.before = pos.before
		pos.before = n
		n.after = pos
	}

	n.list = ns
}

func (ns *nodes) AppendUnstructuredTokens(tokens Tokens) *node {
	if len(tokens) == 0 {
		return nil
	}
	n := newNode(tokens)
	ns.AppendNode(n)
	n.list = ns
	return n
}

// FindNodeWithContent searches the nodes for a node whose content equals
// the given content. If it finds one then it returns it. Otherwise it returns
// nil.
func (ns *nodes) FindNodeWithContent(content nodeContent) *node {
	for n := ns.first; n != nil; n = n.after {
		if n.content == content {
			return n
		}
	}
	return nil
}

// nodeSet is an unordered set of nodes. It is used to describe a set of nodes
// that all belong to the same list that have some role or characteristic
// in common.
type nodeSet map[*node]struct{}

func newNodeSet() nodeSet {
	return make(nodeSet)
}

func (ns nodeSet) Has(n *node) bool {
	if ns == nil {
		return false
	}
	_, exists := ns[n]
	return exists
}

func (ns nodeSet) Add(n *node) {
	ns[n] = struct{}{}
}

func (ns nodeSet) Remove(n *node) {
	delete(ns, n)
}

func (ns nodeSet) Clear() {
	for n := range ns {
		delete(ns, n)
	}
}

func (ns nodeSet) List() []*node {
	if len(ns) == 0 {
		return nil
	}

	ret := make([]*node, 0, len(ns))

	// Determine which list we are working with. We assume here that all of
	// the nodes belong to the same list, since that is part of the contract
	// for nodeSet.
	var list *nodes
	for n := range ns {
		list = n.list
		break
	}

	// We recover the order by iterating over the whole list. This is not
	// the most efficient way to do it, but our node lists should always be
	// small so not worth making things more complex.
	for n := list.first; n != nil; n = n.after {
		if ns.Has(n) {
			ret = append(ret, n)
		}
	}
	return ret
}

// FindNodeWithContent searches the nodes for a node whose content equals
// the given content. If it finds one then it returns it. Otherwise it returns
// nil.
func (ns nodeSet) FindNodeWithContent(content nodeContent) *node {
	for n := range ns {
		if n.content == content {
			return n
		}
	}
	return nil
}

type internalWalkFunc func(*node)

// heldContent is implemented by content that records the node holding it.
// newNode calls setHolder for such content.
type heldContent interface {
	setHolder(n *node)
}

// inTree can be embedded into a content struct that has child nodes to get
// a standard implementation of the NodeContent interface and a record of
// the node that holds the content.
type inTree struct {
	// holder is the node whose content is the receiver, or nil if no node
	// holds it yet. The content is part of a syntax tree while its holder
	// is in a list.
	holder   *node
	children *nodes
}

func newInTree() inTree {
	return inTree{
		children: &nodes{},
	}
}

func (it *inTree) setHolder(n *node) {
	it.holder = n
}

// assertUnattached panics if the receiver is already part of a syntax tree.
func (it *inTree) assertUnattached() {
	if it.holder != nil && it.holder.list != nil {
		panic(fmt.Sprintf("%T is already attached to a syntax tree", it.holder.content))
	}
}

func (it *inTree) walkChildNodes(w internalWalkFunc) {
	for n := it.children.first; n != nil; n = n.after {
		w(n)
	}
}

func (it *inTree) BuildTokens(to Tokens) Tokens {
	for n := it.children.first; n != nil; n = n.after {
		to = n.BuildTokens(to)
	}
	return to
}

// leafNode can be embedded into a content struct to give it a do-nothing
// implementation of walkChildNodes
type leafNode struct {
}

func (n *leafNode) walkChildNodes(w internalWalkFunc) {
}

```

#### pkg/parser.go

```go
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
	writerTokens := writerTokens(nativeTokens)

	from := inputTokens{
		nativeTokens: nativeTokens,
		writerTokens: writerTokens,
	}

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

type inputTokens struct {
	nativeTokens hclsyntax.Tokens
	writerTokens Tokens
}

func (it inputTokens) Partition(rng hcl.Range) (before, within, after inputTokens) {
	start, end := partitionTokens(it.nativeTokens, rng)
	before = it.Slice(0, start)
	within = it.Slice(start, end)
	after = it.Slice(end, len(it.nativeTokens))
	return
}

func (it inputTokens) PartitionType(ty hclsyntax.TokenType) (before, within, after inputTokens) {
	for i, t := range it.writerTokens {
		if t.Type == ty {
			return it.Slice(0, i), it.Slice(i, i+1), it.Slice(i+1, len(it.nativeTokens))
		}
	}
	panic(fmt.Sprintf("didn't find any token of type %s", ty))
}

func (it inputTokens) PartitionTypeOk(ty hclsyntax.TokenType) (before, within, after inputTokens, ok bool) {
	for i, t := range it.writerTokens {
		if t.Type == ty {
			return it.Slice(0, i), it.Slice(i, i+1), it.Slice(i+1, len(it.nativeTokens)), true
		}
	}

	return inputTokens{}, inputTokens{}, inputTokens{}, false
}

func (it inputTokens) PartitionTypeSingle(ty hclsyntax.TokenType) (before inputTokens, found *Token, after inputTokens) {
	before, within, after := it.PartitionType(ty)
	if within.Len() != 1 {
		panic("PartitionType found more than one token")
	}
	return before, within.Tokens()[0], after
}

// PartitionIncludeComments is like Partition except the returned "within"
// range includes any lead and line comments associated with the range.
func (it inputTokens) PartitionIncludingComments(rng hcl.Range) (before, within, after inputTokens) {
	start, end := partitionTokens(it.nativeTokens, rng)
	start = partitionLeadCommentTokens(it.nativeTokens[:start])
	_, afterNewline := partitionLineEndTokens(it.nativeTokens[end:])
	end += afterNewline

	before = it.Slice(0, start)
	within = it.Slice(start, end)
	after = it.Slice(end, len(it.nativeTokens))
	return

}

// PartitionBlockItem is similar to PartitionIncludeComments but it returns
// the comments as separate token sequences so that they can be captured into
// AST attributes. It makes assumptions that apply only to block items, so
// should not be used for other constructs.
func (it inputTokens) PartitionBlockItem(rng hcl.Range) (before, leadComments, within, lineComments, newline, after inputTokens) {
	before, within, after = it.Partition(rng)
	before, leadComments = before.PartitionLeadComments()
	lineComments, newline, after = after.PartitionLineEndTokens()
	return
}

func (it inputTokens) PartitionLeadComments() (before, within inputTokens) {
	start := partitionLeadCommentTokens(it.nativeTokens)
	before = it.Slice(0, start)
	within = it.Slice(start, len(it.nativeTokens))
	return
}

func (it inputTokens) PartitionLineEndTokens() (comments, newline, after inputTokens) {
	afterComments, afterNewline := partitionLineEndTokens(it.nativeTokens)
	comments = it.Slice(0, afterComments)
	newline = it.Slice(afterComments, afterNewline)
	after = it.Slice(afterNewline, len(it.nativeTokens))
	return
}

func (it inputTokens) Slice(start, end int) inputTokens {
	// When we slice, we create a new slice with no additional capacity because
	// we expect that these slices will be mutated in order to insert
	// new code into the AST, and we want to ensure that a new underlying
	// array gets allocated in that case, rather than writing into some
	// following slice and corrupting it.
	return inputTokens{
		nativeTokens: it.nativeTokens[start:end:end],
		writerTokens: it.writerTokens[start:end:end],
	}
}

func (it inputTokens) Len() int {
	return len(it.nativeTokens)
}

func (it inputTokens) Tokens() Tokens {
	return it.writerTokens
}

func (it inputTokens) Types() []hclsyntax.TokenType {
	ret := make([]hclsyntax.TokenType, len(it.nativeTokens))
	for i, tok := range it.nativeTokens {
		ret[i] = tok.Type
	}
	return ret
}

// parseBody locates the given body within the given input tokens and returns
// the resulting *Body object as well as the tokens that appeared before and
// after it.
func parseBody(nativeBody *hclsyntax.Body, from inputTokens) (inputTokens, *node, inputTokens) {
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

func parseBodyItem(nativeItem hclsyntax.Node, from inputTokens) (inputTokens, *node, inputTokens) {
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

func parseAttribute(nativeAttr *hclsyntax.Attribute, from, leadComments, lineComments, newline inputTokens) *node {
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

func parseBlock(nativeBlock *hclsyntax.Block, from, leadComments, lineComments, newline inputTokens) *node {
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

func parseBlockLabels(nativeBlock *hclsyntax.Block, from inputTokens) (inputTokens, *node, inputTokens) {
	labelsObj := newBlockLabels(nil)
	children := labelsObj.children

	var beforeAll inputTokens
	for i, rng := range nativeBlock.LabelRanges {
		var before, labelTokens inputTokens
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

func parseExpression(nativeExpr hclsyntax.Expression, from inputTokens) *node {
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

func parseAnyExpression(nativeExpr hclsyntax.Expression, from inputTokens) *node {
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
func parseObjectConsExpr(nativeExpr *hclsyntax.ObjectConsExpr, from inputTokens) *node {
	expr := newObjectConsExpr()
	children := expr.children

	var before, open, keyTokens, valueTokens inputTokens

	before, open, from = from.Partition(nativeExpr.StartRange())
	children.AppendUnstructuredTokens(before.writerTokens)
	children.AppendUnstructuredTokens(open.writerTokens)

	for _, nativeItem := range nativeExpr.Items {
		item := newObjectConsItem()
		value := newObjectConsValue()

		nativeKeyExpr := nativeItem.KeyExpr.(*hclsyntax.ObjectConsKeyExpr)

		before, keyTokens, from = from.Partition(nativeKeyExpr.Range())
		item.children.AppendUnstructuredTokens(before.writerTokens)
		item.key = parseObjectConsKeyExpr(nativeKeyExpr, keyTokens)
		item.children.AppendNode(item.key)

		before, valueTokens, from = from.Partition(nativeItem.ValueExpr.Range())
		item.children.AppendUnstructuredTokens(before.writerTokens)
		value.expr = parseExpression(nativeItem.ValueExpr, valueTokens)
		value.children.AppendNode(value.expr)
		item.value = item.children.Append(value)

		node := expr.children.Append(item)
		expr.items.Add(node)
	}

	_, from, _ = from.Partition(nativeExpr.Range())
	expr.close = children.AppendUnstructuredTokens(from.writerTokens)

	return wrapExpressionNode(expr)
}

// parseFunctionCallExpr parses a function call expression, defined as:
//
//	FunctionCall = Identifier ("::" Identifier)* "(" Arguments ")";
//	Arguments = (
//	    () ||
//	    (Expression ("," Expression)* ("," | "...")?)
//	);
//
// Each argument becomes its own Expression. All other tokens, including the
// name, the parentheses, the commas and any "...", stay with the call.
func parseFunctionCallExpr(nativeExpr *hclsyntax.FunctionCallExpr, from inputTokens) *node {
	call := newFunctionCallExpr(nativeExpr.Name)
	from = parseExpressionList(call.children, call.args, nativeExpr.Args, from)
	call.children.AppendUnstructuredTokens(from.writerTokens)
	return wrapExpressionNode(call)
}

// parseTupleConsExpr parses a tuple-construct expression, defined as:
//
//	tuple = "[" (
//	    (Expression ("," Expression)* ","?)?
//	) "]";
//
// Each element becomes its own Expression. All other tokens, including the
// brackets and the commas, stay with the tuple.
func parseTupleConsExpr(nativeExpr *hclsyntax.TupleConsExpr, from inputTokens) *node {
	tuple := newTupleConsExpr()
	from = parseExpressionList(tuple.children, tuple.exprs, nativeExpr.Exprs, from)
	tuple.children.AppendUnstructuredTokens(from.writerTokens)
	return wrapExpressionNode(tuple)
}

// parseExpressionList parses each of the given expressions, which must be in
// source order, from the given tokens. It appends the tokens before each
// expression and the node of the expression itself to children, and adds the
// node of the expression to exprNodes. It returns the tokens after the last
// expression.
func parseExpressionList(children *nodes, exprNodes nodeSet, nativeExprs []hclsyntax.Expression, from inputTokens) inputTokens {
	for _, nativeExpr := range nativeExprs {
		before, exprTokens, after := from.Partition(nativeExpr.Range())
		children.AppendUnstructuredTokens(before.writerTokens)
		exprNode := parseExpression(nativeExpr, exprTokens)
		children.AppendNode(exprNode)
		exprNodes.Add(exprNode)
		from = after
	}
	return from
}

// wrapExpressionNode returns a new node holding an Expression that wraps the
// given structured expression content.
func wrapExpressionNode(c nodeContent) *node {
	expr := newExpression()
	expr.wrap(c)
	return newNode(expr)
}

// parseObjectConsKeyExpr parses an object-construct key expression.
// An object key is an Identifier or an Expression. In this context, a quoted
// literal is functionally equivalent to an Identifier.
//
// A naked identifier key is a literal name, as in hclsyntax. A parenthesized
// key is always an expression, even when it contains a single identifier.
func parseObjectConsKeyExpr(nativeExpr *hclsyntax.ObjectConsKeyExpr, from inputTokens) *node {
	expr := parseExpression(nativeExpr.Wrapped, from)
	kind, name := objectConsKeyKindOf(nativeExpr)
	wrapExpr := newObjectConsKeyExpr(expr, kind, name)

	return newNode(wrapExpr)
}

// objectConsKeyKindOf classifies an object-construct key. For a literal name
// or a quoted string literal, it also returns the string the key evaluates
// to, with any escape sequences decoded.
func objectConsKeyKindOf(nativeExpr *hclsyntax.ObjectConsKeyExpr) (objectConsKeyKind, string) {
	if name := hcl.ExprAsKeyword(nativeExpr); name != "" {
		return objectConsKeyLiteralName, name
	}
	if nativeExpr.ForceNonLiteral {
		return objectConsKeyExpression, ""
	}
	tmpl, ok := nativeExpr.Wrapped.(*hclsyntax.TemplateExpr)
	if !ok || !tmpl.IsStringLiteral() {
		return objectConsKeyExpression, ""
	}
	val, diags := tmpl.Value(nil)
	if diags.HasErrors() || !val.IsKnown() || val.IsNull() || val.Type() != cty.String {
		return objectConsKeyExpression, ""
	}
	return objectConsKeyQuotedLiteral, val.AsString()
}

// parseTemplateExpr is specifically interested in string literal expressions
// at this time.
func parseTemplateExpr(nativeExpr *hclsyntax.TemplateExpr, from inputTokens) *node {
	if nativeExpr.IsStringLiteral() {
		quoted := newQuoted(from.writerTokens)

		expr := newExpression()
		expr.wrapped = expr.children.Append(quoted)
		return newNode(expr)
	} else {
		return parseAnyExpression(nativeExpr, from)
	}
}

func parseTraversal(nativeTraversal hcl.Traversal, from inputTokens) (before inputTokens, n *node, after inputTokens) {
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

func parseTraversalStep(nativeStep hcl.Traverser, from inputTokens) (before inputTokens, n *node, after inputTokens) {
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

		var inBefore, oBrack, keyTokens, cBrack inputTokens
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

// writerTokens takes a sequence of tokens as produced by the main hclsyntax
// package and transforms it into an equivalent sequence of tokens using
// this package's own token model.
//
// The resulting list contains the same number of tokens and uses the same
// indices as the input, allowing the two sets of tokens to be correlated
// by index.
func writerTokens(nativeTokens hclsyntax.Tokens) Tokens {
	// Ultimately we want a slice of token _pointers_, but since we can
	// predict how much memory we're going to devote to tokens we'll allocate
	// it all as a single flat buffer and thus give the GC less work to do.
	tokBuf := make([]Token, len(nativeTokens))
	var lastByteOffset int
	for i, mainToken := range nativeTokens {
		// Create a copy of the bytes so that we can mutate without
		// corrupting the original token stream.
		bytes := make([]byte, len(mainToken.Bytes))
		copy(bytes, mainToken.Bytes)

		tokBuf[i] = Token{
			Type:  mainToken.Type,
			Bytes: bytes,

			// We assume here that spaces are always ASCII spaces, since
			// that's what the scanner also assumes, and thus the number
			// of bytes skipped is also the number of space characters.
			SpacesBefore: mainToken.Range.Start.Byte - lastByteOffset,
		}

		lastByteOffset = mainToken.Range.End.Byte
	}

	// Now make a slice of pointers into the previous slice.
	ret := make(Tokens, len(tokBuf))
	for i := range ret {
		ret[i] = &tokBuf[i]
	}

	return ret
}

// partitionTokens takes a sequence of tokens and a hcl.Range and returns
// two indices within the token sequence that correspond with the range
// boundaries, such that the slice operator could be used to produce
// three token sequences for before, within, and after respectively:
//
//	start, end := partitionTokens(toks, rng)
//	before := toks[:start]
//	within := toks[start:end]
//	after := toks[end:]
//
// This works best when the range is aligned with token boundaries (e.g.
// because it was produced in terms of the scanner's result) but if that isn't
// true then it will make a best effort that may produce strange results at
// the boundaries.
//
// Native hclsyntax tokens are used here, because they contain the necessary
// absolute position information. However, since writerTokens produces a
// correlatable sequence of writer tokens, the resulting indices can be
// used also to index into its result, allowing the partitioning of writer
// tokens to be driven by the partitioning of native tokens.
//
// The tokens are assumed to be in source order and non-overlapping, which
// will be true if the token sequence from the scanner is used directly.
func partitionTokens(toks hclsyntax.Tokens, rng hcl.Range) (start, end int) {
	// We use a linear search here because we assume that in most cases our
	// target range is close to the beginning of the sequence, and the sequences
	// are generally small for most reasonable files anyway.
	for i := 0; ; i++ {
		if i >= len(toks) {
			// No tokens for the given range at all!
			return len(toks), len(toks)
		}

		if toks[i].Range.Start.Byte >= rng.Start.Byte {
			start = i
			break
		}
	}

	for i := start; ; i++ {
		if i >= len(toks) {
			// The range "hangs off" the end of the token sequence
			return start, len(toks)
		}

		if toks[i].Range.Start.Byte >= rng.End.Byte {
			end = i // end marker is exclusive
			break
		}
	}

	return start, end
}

// partitionLeadCommentTokens takes a sequence of tokens that is assumed
// to immediately precede a construct that can have lead comment tokens,
// and returns the index into that sequence where the lead comments begin.
//
// Lead comments are defined as whole lines containing only comment tokens
// with no blank lines between. If no such lines are found, the returned
// index will be len(toks).
func partitionLeadCommentTokens(toks hclsyntax.Tokens) int {
	// single-line comments (which is what we're interested in here)
	// consume their trailing newline, so we can just walk backwards
	// until we stop seeing comment tokens.
	for i := len(toks) - 1; i >= 0; i-- {
		if toks[i].Type != hclsyntax.TokenComment {
			return i + 1
		}
	}
	return 0
}

// partitionLineEndTokens takes a sequence of tokens that is assumed
// to immediately follow a construct that can have a line comment, and
// returns first the index where any line comments end and then second
// the index immediately after the trailing newline.
//
// Line comments are defined as comments that appear immediately after
// a construct on the same line where its significant tokens ended.
//
// Since single-line comment tokens (# and //) include the newline that
// terminates them, in the presence of these the two returned indices
// will be the same since the comment itself serves as the line end.
func partitionLineEndTokens(toks hclsyntax.Tokens) (afterComment, afterNewline int) {
	for i := 0; i < len(toks); i++ {
		tok := toks[i]
		if tok.Type != hclsyntax.TokenComment {
			switch tok.Type {
			case hclsyntax.TokenNewline:
				return i, i + 1
			case hclsyntax.TokenEOF:
				// Although this is valid, we mustn't include the EOF
				// itself as our "newline" or else strange things will
				// happen when we try to append new items.
				return i, i
			default:
				// If we have well-formed input here then nothing else should be
				// possible. This path should never happen, because we only try
				// to extract tokens from the sequence if the parser succeeded,
				// and it should catch this problem itself.
				panic("malformed line trailers: expected only comments and newlines")
			}
		}

		if len(tok.Bytes) > 0 && tok.Bytes[len(tok.Bytes)-1] == '\n' {
			// Newline at the end of a single-line comment serves both as
			// the end of comments *and* the end of the line.
			return i + 1, i + 1
		}
	}
	return len(toks), len(toks)
}

// lexConfig uses the hclsyntax scanner to get a token stream and then
// rewrites it into this package's token model.
//
// Any errors produced during scanning are ignored, so the results of this
// function should be used with care.
func lexConfig(src []byte) Tokens {
	mainTokens, _ := hclsyntax.LexConfig(src, "", hcl.Pos{Byte: 0, Line: 1, Column: 1})
	return writerTokens(mainTokens)
}

```

#### pkg/parser_test.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

func TestParse(t *testing.T) {
	tests := []struct {
		src  string
		want TestTreeNode
	}{
		{
			"",
			TestTreeNode{
				Type: "Body",
			},
		},
		{
			"a = 1\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{
								Type: "comments",
							},
							{
								Type: "identifier",
								Val:  "a",
							},
							{
								Type: "Tokens",
								Val:  " =",
							},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "Tokens",
										Val:  " 1",
									},
								},
							},
							{
								Type: "comments",
							},
							{
								Type: "Tokens",
								Val:  "\n",
							},
						},
					},
				},
			},
		},
		{
			"# aye aye aye\na = 1\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{
								Type: "comments",
								Val:  "# aye aye aye\n",
							},
							{
								Type: "identifier",
								Val:  "a",
							},
							{
								Type: "Tokens",
								Val:  " =",
							},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "Tokens",
										Val:  " 1",
									},
								},
							},
							{
								Type: "comments",
							},
							{
								Type: "Tokens",
								Val:  "\n",
							},
						},
					},
				},
			},
		},
		{
			"a = 1 # because it is\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{
								Type: "comments",
							},
							{
								Type: "identifier",
								Val:  "a",
							},
							{
								Type: "Tokens",
								Val:  " =",
							},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "Tokens",
										Val:  " 1",
									},
								},
							},
							{
								Type: "comments",
								Val:  " # because it is\n",
							},
						},
					},
				},
			},
		},
		{
			"# bee bee bee\n\nb = 1\n", // two newlines separate the comment from the attribute
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Tokens", // Only lead/line comments attached to an object have type "comments"
						Val:  "# bee bee bee\n\n",
					},
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{
								Type: "comments",
							},
							{
								Type: "identifier",
								Val:  "b",
							},
							{
								Type: "Tokens",
								Val:  " =",
							},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "Tokens",
										Val:  " 1",
									},
								},
							},
							{
								Type: "comments",
								Val:  "",
							},
							{
								Type: "Tokens",
								Val:  "\n",
							},
						},
					},
				},
			},
		},
		{
			"a = (\n  1 + 2\n)\nb = 3\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{Type: "comments"},
							{
								Type: "identifier",
								Val:  "a",
							},
							{
								Type: "Tokens",
								Val:  " =",
							},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "Tokens",
										Val:  " (\n  1 + 2\n)",
									},
								},
							},
							{Type: "comments"},
							{
								Type: "Tokens",
								Val:  "\n",
							},
						},
					},
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{Type: "comments"},
							{
								Type: "identifier",
								Val:  "b",
							},
							{
								Type: "Tokens",
								Val:  " =",
							},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "Tokens",
										Val:  " 3",
									},
								},
							},
							{Type: "comments"},
							{
								Type: "Tokens",
								Val:  "\n",
							},
						},
					},
				},
			},
		},
		{
			"b {}\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Block",
						Children: []TestTreeNode{
							{
								Type: "comments",
							},
							{
								Type: "identifier",
								Val:  "b",
							},
							{
								Type: "blockLabels",
							},
							{
								Type: "Tokens",
								Val:  " {",
							},
							{
								Type: "Body",
							},
							{
								Type: "Tokens",
								Val:  "}",
							},
							{
								Type: "Tokens",
								Val:  "\n",
							},
						},
					},
				},
			},
		},
		{
			"b label {}\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Block",
						Children: []TestTreeNode{
							{
								Type: "comments",
							},
							{
								Type: "identifier",
								Val:  "b",
							},
							{
								Type: "blockLabels",
								Children: []TestTreeNode{
									{
										Type: "identifier",
										Val:  " label",
									},
								},
							},
							{
								Type: "Tokens",
								Val:  " {",
							},
							{
								Type: "Body",
							},
							{
								Type: "Tokens",
								Val:  "}",
							},
							{
								Type: "Tokens",
								Val:  "\n",
							},
						},
					},
				},
			},
		},
		{
			"b \"label\" {}\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Block",
						Children: []TestTreeNode{
							{
								Type: "comments",
							},
							{
								Type: "identifier",
								Val:  "b",
							},
							{
								Type: "blockLabels",
								Children: []TestTreeNode{
									{
										Type: "quoted",
										Val:  ` "label"`,
									},
								},
							},
							{
								Type: "Tokens",
								Val:  " {",
							},
							{
								Type: "Body",
							},
							{
								Type: "Tokens",
								Val:  "}",
							},
							{
								Type: "Tokens",
								Val:  "\n",
							},
						},
					},
				},
			},
		},
		{
			"b \"label1\" /* foo */ \"label2\" {}\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Block",
						Children: []TestTreeNode{
							{
								Type: "comments",
							},
							{
								Type: "identifier",
								Val:  "b",
							},
							{
								Type: "blockLabels",
								Children: []TestTreeNode{
									{
										Type: "quoted",
										Val:  ` "label1"`,
									},
									{
										// The comment between the labels just
										// becomes an "unstructured tokens"
										// node, because this isn't a place
										// where we expect comments to attach
										// to a particular object as
										// documentation.
										Type: "Tokens",
										Val:  ` /* foo */`,
									},
									{
										Type: "quoted",
										Val:  ` "label2"`,
									},
								},
							},
							{
								Type: "Tokens",
								Val:  " {",
							},
							{
								Type: "Body",
							},
							{
								Type: "Tokens",
								Val:  "}",
							},
							{
								Type: "Tokens",
								Val:  "\n",
							},
						},
					},
				},
			},
		},
		{
			"b {\n  a = 1\n}\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Block",
						Children: []TestTreeNode{
							{
								Type: "comments",
							},
							{
								Type: "identifier",
								Val:  "b",
							},
							{
								Type: "blockLabels",
							},
							{
								Type: "Tokens",
								Val:  " {",
							},
							{
								Type: "Body",
								Children: []TestTreeNode{
									{
										Type: "Tokens",
										Val:  "\n",
									},
									{
										Type: "Attribute",
										Children: []TestTreeNode{
											{
												Type: "comments",
											},
											{
												Type: "identifier",
												Val:  "  a",
											},
											{
												Type: "Tokens",
												Val:  " =",
											},
											{
												Type: "Expression",
												Children: []TestTreeNode{
													{
														Type: "Tokens",
														Val:  " 1",
													},
												},
											},
											{
												Type: "comments",
											},
											{
												Type: "Tokens",
												Val:  "\n",
											},
										},
									},
								},
							},
							{
								Type: "Tokens",
								Val:  "}",
							},
							{
								Type: "Tokens",
								Val:  "\n",
							},
						},
					},
				},
			},
		},
		{
			"a = foo\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{
								Type: "comments",
							},
							{
								Type: "identifier",
								Val:  "a",
							},
							{
								Type: "Tokens",
								Val:  " =",
							},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "Traversal",
										Children: []TestTreeNode{
											{
												Type: "TraverseName",
												Children: []TestTreeNode{
													{
														Type: "identifier",
														Val:  " foo",
													},
												},
											},
										},
									},
								},
							},
							{
								Type: "comments",
							},
							{
								Type: "Tokens",
								Val:  "\n",
							},
						},
					},
				},
			},
		},
		{
			"a = foo.bar\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{
								Type: "comments",
							},
							{
								Type: "identifier",
								Val:  "a",
							},
							{
								Type: "Tokens",
								Val:  " =",
							},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "Traversal",
										Children: []TestTreeNode{
											{
												Type: "TraverseName",
												Children: []TestTreeNode{
													{
														Type: "identifier",
														Val:  " foo",
													},
												},
											},
											{
												Type: "TraverseName",
												Children: []TestTreeNode{
													{
														Type: "Tokens",
														Val:  ".",
													},
													{
														Type: "identifier",
														Val:  "bar",
													},
												},
											},
										},
									},
								},
							},
							{
								Type: "comments",
							},
							{
								Type: "Tokens",
								Val:  "\n",
							},
						},
					},
				},
			},
		},
		{
			"a = foo[0]\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{
								Type: "comments",
							},
							{
								Type: "identifier",
								Val:  "a",
							},
							{
								Type: "Tokens",
								Val:  " =",
							},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "Traversal",
										Children: []TestTreeNode{
											{
												Type: "TraverseName",
												Children: []TestTreeNode{
													{
														Type: "identifier",
														Val:  " foo",
													},
												},
											},
											{
												Type: "TraverseIndex",
												Children: []TestTreeNode{
													{
														Type: "Tokens",
														Val:  "[",
													},
													{
														Type: "number",
														Val:  "0",
													},
													{
														Type: "Tokens",
														Val:  "]",
													},
												},
											},
										},
									},
								},
							},
							{
								Type: "comments",
							},
							{
								Type: "Tokens",
								Val:  "\n",
							},
						},
					},
				},
			},
		},
		{
			"a = foo.0\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{
								Type: "comments",
							},
							{
								Type: "identifier",
								Val:  "a",
							},
							{
								Type: "Tokens",
								Val:  " =",
							},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "Traversal",
										Children: []TestTreeNode{
											{
												Type: "TraverseName",
												Children: []TestTreeNode{
													{
														Type: "identifier",
														Val:  " foo",
													},
												},
											},
											{
												Type: "TraverseIndex",
												Children: []TestTreeNode{
													{
														Type: "Tokens",
														Val:  ".",
													},
													{
														Type: "number",
														Val:  "0",
													},
												},
											},
										},
									},
								},
							},
							{
								Type: "comments",
							},
							{
								Type: "Tokens",
								Val:  "\n",
							},
						},
					},
				},
			},
		},
		{
			"a = foo.*\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{
								Type: "comments",
							},
							{
								Type: "identifier",
								Val:  "a",
							},
							{
								Type: "Tokens",
								Val:  " =",
							},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "Traversal",
										Children: []TestTreeNode{
											{
												Type: "TraverseName",
												Children: []TestTreeNode{
													{
														Type: "identifier",
														Val:  " foo",
													},
												},
											},
										},
									},
									{
										Type: "Tokens",
										Val:  ".*",
									},
								},
							},
							{
								Type: "comments",
							},
							{
								Type: "Tokens",
								Val:  "\n",
							},
						},
					},
				},
			},
		},
		{
			"a = foo.*.bar\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{
								Type: "comments",
							},
							{
								Type: "identifier",
								Val:  "a",
							},
							{
								Type: "Tokens",
								Val:  " =",
							},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "Traversal",
										Children: []TestTreeNode{
											{
												Type: "TraverseName",
												Children: []TestTreeNode{
													{
														Type: "identifier",
														Val:  " foo",
													},
												},
											},
										},
									},
									{
										Type: "Tokens",
										Val:  ".*.bar",
									},
								},
							},
							{
								Type: "comments",
							},
							{
								Type: "Tokens",
								Val:  "\n",
							},
						},
					},
				},
			},
		},
		{
			"a = foo[*]\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{
								Type: "comments",
							},
							{
								Type: "identifier",
								Val:  "a",
							},
							{
								Type: "Tokens",
								Val:  " =",
							},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "Traversal",
										Children: []TestTreeNode{
											{
												Type: "TraverseName",
												Children: []TestTreeNode{
													{
														Type: "identifier",
														Val:  " foo",
													},
												},
											},
										},
									},
									{
										Type: "Tokens",
										Val:  "[*]",
									},
								},
							},
							{
								Type: "comments",
							},
							{
								Type: "Tokens",
								Val:  "\n",
							},
						},
					},
				},
			},
		},
		{
			"a = foo[*].bar\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{
								Type: "comments",
							},
							{
								Type: "identifier",
								Val:  "a",
							},
							{
								Type: "Tokens",
								Val:  " =",
							},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "Traversal",
										Children: []TestTreeNode{
											{
												Type: "TraverseName",
												Children: []TestTreeNode{
													{
														Type: "identifier",
														Val:  " foo",
													},
												},
											},
										},
									},
									{
										Type: "Tokens",
										Val:  "[*].bar",
									},
								},
							},
							{
								Type: "comments",
							},
							{
								Type: "Tokens",
								Val:  "\n",
							},
						},
					},
				},
			},
		},
		{
			"a = foo[bar]\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{
								Type: "comments",
							},
							{
								Type: "identifier",
								Val:  "a",
							},
							{
								Type: "Tokens",
								Val:  " =",
							},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "Traversal",
										Children: []TestTreeNode{
											{
												Type: "TraverseName",
												Children: []TestTreeNode{
													{
														Type: "identifier",
														Val:  " foo",
													},
												},
											},
										},
									},
									{
										Type: "Tokens",
										Val:  "[",
									},
									{
										Type: "Traversal",
										Children: []TestTreeNode{
											{
												Type: "TraverseName",
												Children: []TestTreeNode{
													{
														Type: "identifier",
														Val:  "bar",
													},
												},
											},
										},
									},
									{
										Type: "Tokens",
										Val:  "]",
									},
								},
							},
							{
								Type: "comments",
							},
							{
								Type: "Tokens",
								Val:  "\n",
							},
						},
					},
				},
			},
		},
		{
			"a = foo[bar.baz]\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{
								Type: "comments",
							},
							{
								Type: "identifier",
								Val:  "a",
							},
							{
								Type: "Tokens",
								Val:  " =",
							},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "Traversal",
										Children: []TestTreeNode{
											{
												Type: "TraverseName",
												Children: []TestTreeNode{
													{
														Type: "identifier",
														Val:  " foo",
													},
												},
											},
										},
									},
									{
										Type: "Tokens",
										Val:  "[",
									},
									{
										Type: "Traversal",
										Children: []TestTreeNode{
											{
												Type: "TraverseName",
												Children: []TestTreeNode{
													{
														Type: "identifier",
														Val:  "bar",
													},
												},
											},
											{
												Type: "TraverseName",
												Children: []TestTreeNode{
													{
														Type: "Tokens",
														Val:  ".",
													},
													{
														Type: "identifier",
														Val:  "baz",
													},
												},
											},
										},
									},
									{
										Type: "Tokens",
										Val:  "]",
									},
								},
							},
							{
								Type: "comments",
							},
							{
								Type: "Tokens",
								Val:  "\n",
							},
						},
					},
				},
			},
		},
		{
			"a = foo[bar].baz\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{
								Type: "comments",
							},
							{
								Type: "identifier",
								Val:  "a",
							},
							{
								Type: "Tokens",
								Val:  " =",
							},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "Traversal",
										Children: []TestTreeNode{
											{
												Type: "TraverseName",
												Children: []TestTreeNode{
													{
														Type: "identifier",
														Val:  " foo",
													},
												},
											},
										},
									},
									{
										Type: "Tokens",
										Val:  "[",
									},
									{
										Type: "Traversal",
										Children: []TestTreeNode{
											{
												Type: "TraverseName",
												Children: []TestTreeNode{
													{
														Type: "identifier",
														Val:  "bar",
													},
												},
											},
										},
									},
									{
										Type: "Tokens",
										Val:  "].baz",
									},
								},
							},
							{
								Type: "comments",
							},
							{
								Type: "Tokens",
								Val:  "\n",
							},
						},
					},
				},
			},
		},
		{
			`a = {
				hat = "derby", (cat) = "calico" }`,
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{
								Type: "comments",
							},
							{
								Type: "identifier",
								Val:  "a",
							},
							{
								Type: "Tokens",
								Val:  " =",
							},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "ObjectConsExpr",
										Children: []TestTreeNode{
											{
												Type: "Tokens",
												Val:  " {",
											},
											{
												Type: "ObjectConsItem",
												Children: []TestTreeNode{
													{
														Type: "Tokens",
														Val:  "\n",
													},
													{
														Type: "ObjectConsKeyExpr",
														Children: []TestTreeNode{
															{
																Type: "Expression",
																Children: []TestTreeNode{

																	{
																		Type: "Traversal",
																		Children: []TestTreeNode{
																			{
																				Type: "TraverseName",
																				Children: []TestTreeNode{
																					{
																						Type: "identifier",
																						Val:  "    hat",
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
													{
														Type: "Tokens",
														Val:  " =",
													},
													{
														Type: "ObjectConsValue",
														Children: []TestTreeNode{
															{
																Type: "Expression",
																Children: []TestTreeNode{
																	{
																		Type: "quoted",
																		Val:  ` "derby"`,
																	},
																},
															},
														},
													},
												},
											},
											{
												Type: "ObjectConsItem",
												Children: []TestTreeNode{
													{
														Type: "Tokens",
														Val:  ",",
													},
													{
														Type: "ObjectConsKeyExpr",
														Children: []TestTreeNode{
															{
																Type: "Expression",
																Children: []TestTreeNode{
																	{
																		Type: "Tokens",
																		Val:  " (",
																	},
																	{
																		Type: "Traversal",
																		Children: []TestTreeNode{
																			{
																				Type: "TraverseName",
																				Children: []TestTreeNode{
																					{
																						Type: "identifier",
																						Val:  "cat",
																					},
																				},
																			},
																		},
																	},
																	{
																		Type: "Tokens",
																		Val:  ")",
																	},
																},
															},
														},
													},
													{
														Type: "Tokens",
														Val:  " =",
													},
													{
														Type: "ObjectConsValue",
														Children: []TestTreeNode{
															{
																Type: "Expression",
																Children: []TestTreeNode{
																	{
																		Type: "quoted",
																		Val:  ` "calico"`,
																	},
																},
															},
														},
													},
												},
											},
											{
												Type: "Tokens",
												Val:  " }",
											},
										},
									},
								},
							},
							{
								Type: "comments",
							},
						},
					},
				},
			},
		},
		{
			`a = {
				hat = "derby" // a fancy hat
		   	    (cat) = "calico"
			}`,
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{
								Type: "comments",
							},
							{
								Type: "identifier",
								Val:  "a",
							},
							{
								Type: "Tokens",
								Val:  " =",
							},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "ObjectConsExpr",
										Children: []TestTreeNode{
											{
												Type: "Tokens",
												Val:  " {",
											},
											{
												Type: "ObjectConsItem",
												Children: []TestTreeNode{
													{
														Type: "Tokens",
														Val:  "\n",
													},
													{
														Type: "ObjectConsKeyExpr",
														Children: []TestTreeNode{
															{
																Type: "Expression",
																Children: []TestTreeNode{

																	{
																		Type: "Traversal",
																		Children: []TestTreeNode{
																			{
																				Type: "TraverseName",
																				Children: []TestTreeNode{
																					{
																						Type: "identifier",
																						Val:  "    hat",
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
													{
														Type: "Tokens",
														Val:  " =",
													},
													{
														Type: "ObjectConsValue",
														Children: []TestTreeNode{
															{
																Type: "Expression",
																Children: []TestTreeNode{
																	{
																		Type: "quoted",
																		Val:  ` "derby"`,
																	},
																},
															},
														},
													},
												},
											},
											{
												Type: "ObjectConsItem",
												Children: []TestTreeNode{
													{
														Type: "Tokens",
														Val:  " // a fancy hat\n",
													},
													{
														Type: "ObjectConsKeyExpr",
														Children: []TestTreeNode{
															{
																Type: "Expression",
																Children: []TestTreeNode{
																	{
																		Type: "Tokens",
																		Val:  "          (",
																	},
																	{
																		Type: "Traversal",
																		Children: []TestTreeNode{
																			{
																				Type:     "TraverseName",
																				Children: []TestTreeNode{{Type: "identifier", Val: "cat"}},
																			},
																		},
																	},
																	{
																		Type: "Tokens",
																		Val:  ")",
																	},
																},
															},
														},
													},
													{
														Type: "Tokens",
														Val:  " =",
													},
													{
														Type: "ObjectConsValue",
														Children: []TestTreeNode{
															{
																Type: "Expression",
																Children: []TestTreeNode{
																	{
																		Type: "quoted",
																		Val:  ` "calico"`,
																	},
																},
															},
														},
													},
												},
											},
											{
												Type: "Tokens",
												Val:  "\n   }",
											},
										},
									},
								},
							},
							{
								Type: "comments",
							},
						},
					},
				},
			},
		},
		{
			`a = {
  hat = {
	style = "derby",
	color = (variable)
  },
  cat = "calico"
}`,
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{{
					Type: "Attribute",
					Children: []TestTreeNode{
						{Type: "comments"},
						{Type: "identifier", Val: "a"},
						{Type: "Tokens", Val: " ="},
						{
							Type: "Expression",
							Children: []TestTreeNode{{
								Type: "ObjectConsExpr",
								Children: []TestTreeNode{
									{Type: "Tokens", Val: " {"},
									{
										Type: "ObjectConsItem",
										Children: []TestTreeNode{
											{Type: "Tokens", Val: "\n"},
											{
												Type: "ObjectConsKeyExpr",
												Children: []TestTreeNode{{
													Type: "Expression",
													Children: []TestTreeNode{{
														Type: "Traversal",
														Children: []TestTreeNode{
															{
																Type:     "TraverseName",
																Children: []TestTreeNode{{Type: "identifier", Val: "  hat"}},
															},
														},
													}},
												}},
											},
											{Type: "Tokens", Val: " ="},
											{
												Type: "ObjectConsValue",
												Children: []TestTreeNode{{
													Type: "Expression",
													Children: []TestTreeNode{{
														Type: "ObjectConsExpr",
														Children: []TestTreeNode{
															{Type: "Tokens", Val: " {"},
															{
																Type: "ObjectConsItem",
																Children: []TestTreeNode{
																	{Type: "Tokens", Val: "\n"},
																	{
																		Type: "ObjectConsKeyExpr",
																		Children: []TestTreeNode{
																			{
																				Type: "Expression",
																				Children: []TestTreeNode{
																					{
																						Type: "Traversal",
																						Children: []TestTreeNode{
																							{
																								Type: "TraverseName",
																								Children: []TestTreeNode{
																									{Type: "identifier", Val: " style"},
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																	{Type: "Tokens", Val: " ="},
																	{
																		Type: "ObjectConsValue",
																		Children: []TestTreeNode{
																			{
																				Type:     "Expression",
																				Children: []TestTreeNode{{Type: "quoted", Val: ` "derby"`}},
																			},
																		},
																	},
																},
															},
															{
																Type: "ObjectConsItem",
																Children: []TestTreeNode{
																	{Type: "Tokens", Val: ",\n"},
																	{
																		Type: "ObjectConsKeyExpr",
																		Children: []TestTreeNode{{
																			Type: "Expression",
																			Children: []TestTreeNode{{
																				Type: "Traversal", Children: []TestTreeNode{{
																					Type:     "TraverseName",
																					Children: []TestTreeNode{{Type: "identifier", Val: " color"}},
																				}},
																			}},
																		}},
																	},
																	{Type: "Tokens", Val: " ="},
																	{
																		Type: "ObjectConsValue",
																		Children: []TestTreeNode{{
																			Type: "Expression",
																			Children: []TestTreeNode{
																				{Type: "Tokens", Val: " ("},
																				{
																					Type: "Traversal",
																					Children: []TestTreeNode{{
																						Type:     "TraverseName",
																						Children: []TestTreeNode{{Type: "identifier", Val: "variable"}},
																					}},
																				},
																				{Type: "Tokens", Val: ")"},
																			},
																		}},
																	},
																},
															},
															{Type: "Tokens", Val: "\n  }"},
														},
													}},
												}},
											},
										},
									},
									{
										Type: "ObjectConsItem",
										Children: []TestTreeNode{
											{Type: "Tokens", Val: ",\n"},
											{
												Type: "ObjectConsKeyExpr",
												Children: []TestTreeNode{{
													Type: "Expression",
													Children: []TestTreeNode{
														{
															Type: "Traversal",
															Children: []TestTreeNode{{
																Type: "TraverseName",
																Children: []TestTreeNode{
																	{Type: "identifier", Val: "  cat"},
																},
															}},
														},
													},
												}},
											},
											{Type: "Tokens", Val: " ="},
											{
												Type: "ObjectConsValue",
												Children: []TestTreeNode{{
													Type: "Expression",
													Children: []TestTreeNode{
														{Type: "quoted", Val: ` "calico"`},
													},
												}},
											},
										},
									},
									{Type: "Tokens", Val: "\n}"},
								},
							}},
						},
						{Type: "comments"},
					},
				}},
			},
		},
		{
			`zzz = {
				"hat" = "derby" }`,
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{
								Type: "comments",
							},
							{
								Type: "identifier",
								Val:  "zzz",
							},
							{
								Type: "Tokens",
								Val:  " =",
							},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "ObjectConsExpr",
										Children: []TestTreeNode{
											{
												Type: "Tokens",
												Val:  " {",
											},
											{
												Type: "ObjectConsItem",
												Children: []TestTreeNode{
													{
														Type: "Tokens",
														Val:  "\n",
													},
													{
														Type: "ObjectConsKeyExpr",
														Children: []TestTreeNode{
															{
																Type: "Expression",
																Children: []TestTreeNode{
																	{
																		Type: "quoted",
																		Val:  `    "hat"`,
																	},
																},
															},
														},
													},
													{
														Type: "Tokens",
														Val:  " =",
													},
													{
														Type: "ObjectConsValue",
														Children: []TestTreeNode{
															{
																Type: "Expression",
																Children: []TestTreeNode{
																	{
																		Type: "quoted",
																		Val:  ` "derby"`,
																	},
																},
															},
														},
													},
												},
											},
											{
												Type: "Tokens",
												Val:  " }",
											},
										},
									},
								},
							},
							{
								Type: "comments",
							},
						},
					},
				},
			},
		},
		{
			`a = join(" ", ["fedora", x.fez], f())` + "\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{{
					Type: "Attribute",
					Children: []TestTreeNode{
						{Type: "comments"},
						{Type: "identifier", Val: "a"},
						{Type: "Tokens", Val: " ="},
						{
							Type: "Expression",
							Children: []TestTreeNode{{
								Type: "FunctionCallExpr",
								Children: []TestTreeNode{
									{Type: "Tokens", Val: " join("},
									{
										Type:     "Expression",
										Children: []TestTreeNode{{Type: "quoted", Val: `" "`}},
									},
									{Type: "Tokens", Val: ","},
									{
										Type: "Expression",
										Children: []TestTreeNode{{
											Type: "TupleConsExpr",
											Children: []TestTreeNode{
												{Type: "Tokens", Val: " ["},
												{
													Type:     "Expression",
													Children: []TestTreeNode{{Type: "quoted", Val: `"fedora"`}},
												},
												{Type: "Tokens", Val: ","},
												{
													Type: "Expression",
													Children: []TestTreeNode{{
														Type: "Traversal",
														Children: []TestTreeNode{
															{
																Type:     "TraverseName",
																Children: []TestTreeNode{{Type: "identifier", Val: " x"}},
															},
															{
																Type: "TraverseName",
																Children: []TestTreeNode{
																	{Type: "Tokens", Val: "."},
																	{Type: "identifier", Val: "fez"},
																},
															},
														},
													}},
												},
												{Type: "Tokens", Val: "]"},
											},
										}},
									},
									{Type: "Tokens", Val: ","},
									{
										Type: "Expression",
										Children: []TestTreeNode{{
											Type: "FunctionCallExpr",
											Children: []TestTreeNode{
												{Type: "Tokens", Val: " f()"},
											},
										}},
									},
									{Type: "Tokens", Val: ")"},
								},
							}},
						},
						{Type: "comments"},
						{Type: "Tokens", Val: "\n"},
					},
				}},
			},
		},
		{
			"a = [\n  1, # one\n  f(x...),\n]\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{{
					Type: "Attribute",
					Children: []TestTreeNode{
						{Type: "comments"},
						{Type: "identifier", Val: "a"},
						{Type: "Tokens", Val: " ="},
						{
							Type: "Expression",
							Children: []TestTreeNode{{
								Type: "TupleConsExpr",
								Children: []TestTreeNode{
									{Type: "Tokens", Val: " [\n"},
									{
										Type:     "Expression",
										Children: []TestTreeNode{{Type: "Tokens", Val: "  1"}},
									},
									{Type: "Tokens", Val: ", # one\n"},
									{
										Type: "Expression",
										Children: []TestTreeNode{{
											Type: "FunctionCallExpr",
											Children: []TestTreeNode{
												{Type: "Tokens", Val: "  f("},
												{
													Type: "Expression",
													Children: []TestTreeNode{{
														Type: "Traversal",
														Children: []TestTreeNode{{
															Type:     "TraverseName",
															Children: []TestTreeNode{{Type: "identifier", Val: "x"}},
														}},
													}},
												},
												{Type: "Tokens", Val: "...)"},
											},
										}},
									},
									{Type: "Tokens", Val: ",\n]"},
								},
							}},
						},
						{Type: "comments"},
						{Type: "Tokens", Val: "\n"},
					},
				}},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			file, diags := parse([]byte(test.src), "", hcl.Pos{Line: 1, Column: 1})
			if len(diags) > 0 {
				for _, diag := range diags {
					t.Logf(" - %s", diag.Error())
				}
				t.Fatalf("unexpected diagnostics")
			}

			got := makeTestTree(file.body)

			if !cmp.Equal(got, test.want) {
				diff := cmp.Diff(got, test.want)
				t.Errorf(
					"wrong result\ninput:\n%s\n\ngot:\n%s\nwant:%s\n\ndiff:\n%s",
					test.src,
					spew.Sdump(got),
					spew.Sdump(test.want),
					diff,
				)
			}
		})
	}
}

func TestPartitionTokens(t *testing.T) {
	tests := []struct {
		tokens    hclsyntax.Tokens
		rng       hcl.Range
		wantStart int
		wantEnd   int
	}{
		{
			hclsyntax.Tokens{},
			hcl.Range{
				Start: hcl.Pos{Byte: 0},
				End:   hcl.Pos{Byte: 0},
			},
			0,
			0,
		},
		{
			hclsyntax.Tokens{
				{
					Type: hclsyntax.TokenIdent,
					Range: hcl.Range{
						Start: hcl.Pos{Byte: 0},
						End:   hcl.Pos{Byte: 4},
					},
				},
			},
			hcl.Range{
				Start: hcl.Pos{Byte: 0},
				End:   hcl.Pos{Byte: 4},
			},
			0,
			1,
		},
		{
			hclsyntax.Tokens{
				{
					Type: hclsyntax.TokenIdent,
					Range: hcl.Range{
						Start: hcl.Pos{Byte: 0},
						End:   hcl.Pos{Byte: 4},
					},
				},
				{
					Type: hclsyntax.TokenIdent,
					Range: hcl.Range{
						Start: hcl.Pos{Byte: 4},
						End:   hcl.Pos{Byte: 8},
					},
				},
				{
					Type: hclsyntax.TokenIdent,
					Range: hcl.Range{
						Start: hcl.Pos{Byte: 8},
						End:   hcl.Pos{Byte: 12},
					},
				},
			},
			hcl.Range{
				Start: hcl.Pos{Byte: 4},
				End:   hcl.Pos{Byte: 8},
			},
			1,
			2,
		},
		{
			hclsyntax.Tokens{
				{
					Type: hclsyntax.TokenIdent,
					Range: hcl.Range{
						Start: hcl.Pos{Byte: 0},
						End:   hcl.Pos{Byte: 4},
					},
				},
				{
					Type: hclsyntax.TokenIdent,
					Range: hcl.Range{
						Start: hcl.Pos{Byte: 4},
						End:   hcl.Pos{Byte: 8},
					},
				},
				{
					Type: hclsyntax.TokenIdent,
					Range: hcl.Range{
						Start: hcl.Pos{Byte: 8},
						End:   hcl.Pos{Byte: 12},
					},
				},
			},
			hcl.Range{
				Start: hcl.Pos{Byte: 0},
				End:   hcl.Pos{Byte: 8},
			},
			0,
			2,
		},
		{
			hclsyntax.Tokens{
				{
					Type: hclsyntax.TokenIdent,
					Range: hcl.Range{
						Start: hcl.Pos{Byte: 0},
						End:   hcl.Pos{Byte: 4},
					},
				},
				{
					Type: hclsyntax.TokenIdent,
					Range: hcl.Range{
						Start: hcl.Pos{Byte: 4},
						End:   hcl.Pos{Byte: 8},
					},
				},
				{
					Type: hclsyntax.TokenIdent,
					Range: hcl.Range{
						Start: hcl.Pos{Byte: 8},
						End:   hcl.Pos{Byte: 12},
					},
				},
			},
			hcl.Range{
				Start: hcl.Pos{Byte: 4},
				End:   hcl.Pos{Byte: 12},
			},
			1,
			3,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%02d", i), func(t *testing.T) {
			gotStart, gotEnd := partitionTokens(test.tokens, test.rng)

			if gotStart != test.wantStart || gotEnd != test.wantEnd {
				t.Errorf(
					"wrong result\ntokens: %s\nrange: %#v\ngot:   %d, %d\nwant:  %d, %d",
					spew.Sdump(test.tokens), test.rng,
					gotStart, test.wantStart,
					gotEnd, test.wantEnd,
				)
			}
		})
	}
}

func TestPartitionLeadCommentTokens(t *testing.T) {
	tests := []struct {
		tokens    hclsyntax.Tokens
		wantStart int
	}{
		{
			hclsyntax.Tokens{},
			0,
		},
		{
			hclsyntax.Tokens{
				{
					Type: hclsyntax.TokenComment,
				},
			},
			0,
		},
		{
			hclsyntax.Tokens{
				{
					Type: hclsyntax.TokenComment,
				},
				{
					Type: hclsyntax.TokenComment,
				},
			},
			0,
		},
		{
			hclsyntax.Tokens{
				{
					Type: hclsyntax.TokenComment,
				},
				{
					Type: hclsyntax.TokenNewline,
				},
			},
			2,
		},
		{
			hclsyntax.Tokens{
				{
					Type: hclsyntax.TokenComment,
				},
				{
					Type: hclsyntax.TokenNewline,
				},
				{
					Type: hclsyntax.TokenComment,
				},
			},
			2,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%02d", i), func(t *testing.T) {
			gotStart := partitionLeadCommentTokens(test.tokens)

			if gotStart != test.wantStart {
				t.Errorf(
					"wrong result\ntokens: %s\ngot:   %d\nwant:  %d",
					spew.Sdump(test.tokens),
					gotStart, test.wantStart,
				)
			}
		})
	}
}

func TestLexConfig(t *testing.T) {
	tests := []struct {
		input string
		want  Tokens
	}{
		{
			`a  b `,
			Tokens{
				{
					Type:         hclsyntax.TokenIdent,
					Bytes:        []byte(`a`),
					SpacesBefore: 0,
				},
				{
					Type:         hclsyntax.TokenIdent,
					Bytes:        []byte(`b`),
					SpacesBefore: 2,
				},
				{
					Type:         hclsyntax.TokenEOF,
					Bytes:        []byte{},
					SpacesBefore: 1,
				},
			},
		},
		{
			`
foo "bar" "baz" {
    pizza = " cheese "
}
`,
			Tokens{
				{
					Type:         hclsyntax.TokenNewline,
					Bytes:        []byte{'\n'},
					SpacesBefore: 0,
				},
				{
					Type:         hclsyntax.TokenIdent,
					Bytes:        []byte(`foo`),
					SpacesBefore: 0,
				},
				{
					Type:         hclsyntax.TokenOQuote,
					Bytes:        []byte(`"`),
					SpacesBefore: 1,
				},
				{
					Type:         hclsyntax.TokenQuotedLit,
					Bytes:        []byte(`bar`),
					SpacesBefore: 0,
				},
				{
					Type:         hclsyntax.TokenCQuote,
					Bytes:        []byte(`"`),
					SpacesBefore: 0,
				},
				{
					Type:         hclsyntax.TokenOQuote,
					Bytes:        []byte(`"`),
					SpacesBefore: 1,
				},
				{
					Type:         hclsyntax.TokenQuotedLit,
					Bytes:        []byte(`baz`),
					SpacesBefore: 0,
				},
				{
					Type:         hclsyntax.TokenCQuote,
					Bytes:        []byte(`"`),
					SpacesBefore: 0,
				},
				{
					Type:         hclsyntax.TokenOBrace,
					Bytes:        []byte(`{`),
					SpacesBefore: 1,
				},
				{
					Type:         hclsyntax.TokenNewline,
					Bytes:        []byte("\n"),
					SpacesBefore: 0,
				},
				{
					Type:         hclsyntax.TokenIdent,
					Bytes:        []byte(`pizza`),
					SpacesBefore: 4,
				},
				{
					Type:         hclsyntax.TokenEqual,
					Bytes:        []byte(`=`),
					SpacesBefore: 1,
				},
				{
					Type:         hclsyntax.TokenOQuote,
					Bytes:        []byte(`"`),
					SpacesBefore: 1,
				},
				{
					Type:         hclsyntax.TokenQuotedLit,
					Bytes:        []byte(` cheese `),
					SpacesBefore: 0,
				},
				{
					Type:         hclsyntax.TokenCQuote,
					Bytes:        []byte(`"`),
					SpacesBefore: 0,
				},
				{
					Type:         hclsyntax.TokenNewline,
					Bytes:        []byte("\n"),
					SpacesBefore: 0,
				},
				{
					Type:         hclsyntax.TokenCBrace,
					Bytes:        []byte(`}`),
					SpacesBefore: 0,
				},
				{
					Type:         hclsyntax.TokenNewline,
					Bytes:        []byte("\n"),
					SpacesBefore: 0,
				},
				{
					Type:         hclsyntax.TokenEOF,
					Bytes:        []byte{},
					SpacesBefore: 0,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := lexConfig([]byte(test.input))

			diff := cmp.Diff(test.want, got)
			if diff != "" {
				t.Errorf(
					"wrong result\ninput: %s\ndiff:\n%s",
					test.input, diff,
				)
			}
		})
	}
}

```

Answer with one JSON object and nothing else:
{"findings": [{"file": "...", "line": 0, "problem": "...", "fix": "..."}]}
Give each finding a concrete file, line and fix. Leave out vague remarks.
An empty list is a valid answer.
