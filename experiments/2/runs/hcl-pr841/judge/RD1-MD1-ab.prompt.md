Two versions of the same Go package are in this directory, as `x/` and `y/`.
Both started from the same code and gained the same features, written by
different people. Compare how well each package is divided into files and types.

Judge the package as a whole, and how the new code fits into it:
- Does each file and each type have one clear job?
- Is closely related code kept together?
- Do the boundaries between files follow the boundaries between features?
- Does the new code fit the structure the package already had?

Behaviour and tests are known to be equivalent. Do not judge style, comments,
or naming conventions on their own, except where a name shows how the code is
divided.

Below is a map of each package: every file with its line count and its
top-level declarations. Files marked NEW were added and files marked CHANGED
were modified. The full text of every NEW and CHANGED file follows the maps.
You may read any other file under x/ or y/ if you need to.

### x/

| file | lines | mark | declarations |
| --- | --- | --- | --- |
| ast.go | 121 |  | File (type), NewEmptyFile (func), File.Body (method), File.WriteTo (method), File.Bytes (method), comments (type), newComments (func), comments.BuildTokens (method), identifier (type), newIdentifier (func), identifier.BuildTokens (method), identifier.hasName (method), number (type), newNumber (func), number.BuildTokens (method), quoted (type), newQuoted (func), quoted.BuildTokens (method) |
| ast_attribute.go | 65 |  | Attribute (type), newAttribute (func), Attribute.init (method), Attribute.LeadComments (method), Attribute.LineComments (method), Attribute.Expr (method), Attribute.setName (method) |
| ast_attribute_test.go | 125 |  | TestAttributeLeadComments (func), TestAttributeLineComments (func) |
| ast_block.go | 184 |  | Block (type), newBlock (func), NewBlock (func), Block.init (method), Block.Body (method), Block.LeadComments (method), Block.Type (method), Block.SetType (method), Block.Labels (method), Block.SetLabels (method), Block.labelsObj (method), blockLabels (type), newBlockLabels (func), blockLabels.Replace (method), blockLabels.Current (method) |
| ast_block_test.go | 579 |  | TestBlockType (func), TestBlockLabels (func), TestBlockSetType_buildTokens (func), TestBlockSetType (func), TestBlockSetLabels (func), TestBlockLeadComments (func), mustParseConfig (func) |
| ast_body.go | 309 |  | Body (type), newBody (func), Body.appendItem (method), Body.appendItemNode (method), Body.Clear (method), Body.AppendUnstructuredTokens (method), Body.Attributes (method), Body.Blocks (method), Body.GetAttribute (method), Body.getAttributeNode (method), Body.RenameAttribute (method), Body.FirstMatchingBlock (method), Body.RemoveNewlineBeforeBlock (method), Body.RemoveBlock (method), Body.SetAttributeRaw (method), Body.SetAttributeValue (method), Body.SetAttributeTraversal (method), Body.RemoveAttribute (method), Body.AppendBlock (method), Body.AppendNewBlock (method), Body.AppendNewline (method) |
| ast_body_test.go | 2195 |  | TestBodyGetAttribute (func), TestBodyFirstMatchingBlock (func), TestBodySetAttributeValue (func), TestBodySetAttributeTraversal (func), TestBodySetAttributeTraversal_ReturnsTheAttribute (func), TestBodySetAttributeRaw (func), TestBodySetAttributeRaw_ReturnsTheAttribute (func), TestBodySetAttributeValue_ReturnsTheAttribute (func), TestBodySetAttributeValueInBlock (func), TestBodySetAttributeValueInNestedBlock (func), TestBodyRemoveAttribute (func), TestBodyRenameAttribute (func), TestBodyAppendBlock (func), TestBodyRemoveBlock (func), TestBodyRemoveNewlineBeforeBlock (func), testFn (func) |
| ast_expression.go | 317 | CHANGED | Expression (type), newExpression (func), Expression.wrap (method), NewExpressionRaw (func), NewExpressionLiteral (func), NewExpressionAbsTraversal (func), Expression.Variables (method), Expression.RenameVariablePrefix (method), Expression.AsObjectConsExpr (method), Expression.AsFunctionCallExpr (method), Expression.AsTupleConsExpr (method), Expression.traversals (method), Expression.asQuoted (method), Expression.AsQuotedLiteral (method), Traversal (type), newTraversal (func), Traversal.renamePrefix (method), TraverseName (type), newTraverseName (func), TraverseIndex (type), newTraverseIndex (func) |
| ast_expression_test.go | 189 | NEW | TestExpressionVariablesNested (func), TestExpressionAsFunctionCallExprAndTupleConsExpr (func) |
| ast_function_call_expr.go | 61 | NEW | FunctionCallExpr (type), newFunctionCallExpr (func), FunctionCallExpr.Name (method), FunctionCallExpr.traversals (method) |
| ast_function_call_expr_test.go | 40 | NEW | TestFunctionCallExprName (func) |
| ast_object_cons_expr.go | 323 | CHANGED | ObjectConsExpr (type), newObjectConsExpr (func), ObjectConsExpr.Items (method), ObjectConsExpr.ItemFor (method), ObjectConsExpr.ValueExprFor (method), ObjectConsExpr.SetItem (method), ObjectConsExpr.SetItemRaw (method), ObjectConsExpr.SetItemValue (method), ObjectConsExpr.SetItemTraversal (method), ObjectConsExpr.traversals (method), ObjectConsExpr.appendItem (method), ObjectConsExpr.multiline (method), objectConsBreaksLine (func), ObjectConsItem (type), newObjectConsItem (func), ObjectConsItem.KeyExpr (method), ObjectConsItem.ValueExpr (method), ObjectConsItem.init (method), ObjectConsKeyExpr (type), newObjectConsKeyExpr (func), newObjectConsKeyExprNamed (func), ObjectConsKeyExpr.hasName (method), ObjectConsKeyExpr.traversals (method), ObjectConsValue (type), newObjectConsValue (func), ObjectConsValue.Expr (method) |
| ast_object_cons_expr_test.go | 501 | NEW | TestObjectConsExprItemFor (func), TestObjectConsExprItems (func), TestObjectConsExprSetItemKey (func), TestObjectConsExprSetItemExisting (func), TestObjectConsExprSetItemHelpers (func), TestObjectConsExprSetItemAttached (func), TestObjectConsExprSetItemLayout (func) |
| ast_test.go | 50 |  | TestTreeNode (type), makeTestTree (func) |
| ast_tuple_cons_expr.go | 34 | NEW | TupleConsExpr (type), newTupleConsExpr (func), TupleConsExpr.traversals (method) |
| doc.go | 14 |  |  |
| examples_test.go | 125 | CHANGED | Example_generateFromScratch (func), ExampleExpression_RenameVariablePrefix (func), ExampleObjectConsExpr_SetItemRaw (func) |
| format.go | 474 | CHANGED | format (func), formatIndent (func), formatSpaces (func), formatCells (func), spaceAfterToken (func), linesForFormat (func), tokenIsNewline (func), tokenBracketChange (func), formatLine (type) |
| format_test.go | 800 | CHANGED | TestFormat (func), TestLinesForFormat (func) |
| generate.go | 399 |  | TokensForValue (func), TokensForTraversal (func), TokensForIdentifier (func), TokensForTuple (func), TokensForObject (func), TokensForFunctionCall (func), appendTokensForValue (func), appendTokensForTraversal (func), appendTokensForTraversalStep (func), escapeQuotedStringLit (func), appendRune (func) |
| generate_test.go | 920 |  | TestTokensForValue (func), TestTokensForTraversal (func), TestTokensForTuple (func), TestTokensForObject (func), TestTokensForFunctionCall (func), TestTokenGenerateConsistency (func) |
| native_node_sorter.go | 26 |  | nativeNodeSorter (type), nativeNodeSorter.Len (method), nativeNodeSorter.Less (method), nativeNodeSorter.Swap (method) |
| node.go | 331 | CHANGED | node (type), newNode (func), node.Equal (method), node.BuildTokens (method), node.Detach (method), node.ReplaceWith (method), node.assertUnattached (method), nodeContent (type), nodes (type), nodes.BuildTokens (method), nodes.List (method), nodes.Last (method), nodes.Clear (method), nodes.Append (method), nodes.AppendNode (method), nodes.Insert (method), nodes.InsertNode (method), nodes.AppendUnstructuredTokens (method), nodes.FindNodeWithContent (method), nodeSet (type), newNodeSet (func), nodeSet.Has (method), nodeSet.Add (method), nodeSet.Remove (method), nodeSet.Clear (method), nodeSet.List (method), nodeSet.FindNodeWithContent (method), internalWalkFunc (type), inTree (type), newInTree (func), inTree.assertUnattached (method), inTree.setHolder (method), inTree.walkChildNodes (method), inTree.BuildTokens (method), leafNode (type), leafNode.walkChildNodes (method) |
| parser.go | 527 | CHANGED | parse (func), parseBody (func), parseBodyItem (func), parseAttribute (func), parseBlock (func), parseBlockLabels (func), parseExpression (func), parseAnyExpression (func), parseObjectConsExpr (func), parseObjectConsKeyExpr (func), parseFunctionCallExpr (func), parseTupleConsExpr (func), parseExpressionList (func), parseTemplateExpr (func), parseTraversal (func), parseTraversalStep (func) |
| parser_test.go | 2168 | CHANGED | TestParse (func) |
| partition.go | 253 | NEW | partitionableTokens (type), newPartitionableTokens (func), partitionableTokens.Partition (method), partitionableTokens.PartitionType (method), partitionableTokens.PartitionTypeOk (method), partitionableTokens.PartitionTypeSingle (method), partitionableTokens.PartitionIncludingComments (method), partitionableTokens.PartitionBlockItem (method), partitionableTokens.PartitionLeadComments (method), partitionableTokens.PartitionLineEndTokens (method), partitionableTokens.Slice (method), partitionableTokens.Len (method), partitionableTokens.Tokens (method), partitionableTokens.Types (method), partitionTokens (func), partitionLeadCommentTokens (func), partitionLineEndTokens (func) |
| partition_test.go | 228 | NEW | TestPartitionTokens (func), TestPartitionLeadCommentTokens (func) |
| public.go | 48 | CHANGED | NewFile (func), ParseConfig (func), Format (func) |
| round_trip_test.go | 192 | CHANGED | TestRoundTripVerbatim (func), TestRoundTripFormat (func), TestRoundTripSafeConcurrent (func) |
| tokens.go | 189 | CHANGED | Token (type), Token.asHCLSyntax (method), Tokens (type), Tokens.Bytes (method), Tokens.testValue (method), Tokens.Columns (method), Tokens.WriteTo (method), Tokens.walkChildNodes (method), Tokens.BuildTokens (method), ObjectAttrTokens (type), writerTokens (func), lexConfigTokens (func), newIdentToken (func) |
| tokens_test.go | 158 | NEW | TestLexConfig (func) |

### y/

| file | lines | mark | declarations |
| --- | --- | --- | --- |
| ast.go | 121 |  | File (type), NewEmptyFile (func), File.Body (method), File.WriteTo (method), File.Bytes (method), comments (type), newComments (func), comments.BuildTokens (method), identifier (type), newIdentifier (func), identifier.BuildTokens (method), identifier.hasName (method), number (type), newNumber (func), number.BuildTokens (method), quoted (type), newQuoted (func), quoted.BuildTokens (method) |
| ast_attribute.go | 65 |  | Attribute (type), newAttribute (func), Attribute.init (method), Attribute.LeadComments (method), Attribute.LineComments (method), Attribute.Expr (method), Attribute.setName (method) |
| ast_attribute_test.go | 125 |  | TestAttributeLeadComments (func), TestAttributeLineComments (func) |
| ast_block.go | 184 |  | Block (type), newBlock (func), NewBlock (func), Block.init (method), Block.Body (method), Block.LeadComments (method), Block.Type (method), Block.SetType (method), Block.Labels (method), Block.SetLabels (method), Block.labelsObj (method), blockLabels (type), newBlockLabels (func), blockLabels.Replace (method), blockLabels.Current (method) |
| ast_block_test.go | 579 |  | TestBlockType (func), TestBlockLabels (func), TestBlockSetType_buildTokens (func), TestBlockSetType (func), TestBlockSetLabels (func), TestBlockLeadComments (func), mustParseConfig (func) |
| ast_body.go | 309 |  | Body (type), newBody (func), Body.appendItem (method), Body.appendItemNode (method), Body.Clear (method), Body.AppendUnstructuredTokens (method), Body.Attributes (method), Body.Blocks (method), Body.GetAttribute (method), Body.getAttributeNode (method), Body.RenameAttribute (method), Body.FirstMatchingBlock (method), Body.RemoveNewlineBeforeBlock (method), Body.RemoveBlock (method), Body.SetAttributeRaw (method), Body.SetAttributeValue (method), Body.SetAttributeTraversal (method), Body.RemoveAttribute (method), Body.AppendBlock (method), Body.AppendNewBlock (method), Body.AppendNewline (method) |
| ast_body_test.go | 2195 |  | TestBodyGetAttribute (func), TestBodyFirstMatchingBlock (func), TestBodySetAttributeValue (func), TestBodySetAttributeTraversal (func), TestBodySetAttributeTraversal_ReturnsTheAttribute (func), TestBodySetAttributeRaw (func), TestBodySetAttributeRaw_ReturnsTheAttribute (func), TestBodySetAttributeValue_ReturnsTheAttribute (func), TestBodySetAttributeValueInBlock (func), TestBodySetAttributeValueInNestedBlock (func), TestBodyRemoveAttribute (func), TestBodyRenameAttribute (func), TestBodyAppendBlock (func), TestBodyRemoveBlock (func), TestBodyRemoveNewlineBeforeBlock (func), testFn (func) |
| ast_expression.go | 315 | CHANGED | Expression (type), newExpression (func), Expression.wrap (method), NewExpressionRaw (func), NewExpressionLiteral (func), NewExpressionAbsTraversal (func), Expression.Variables (method), Expression.RenameVariablePrefix (method), Expression.AsObjectConsExpr (method), Expression.AsFunctionCallExpr (method), Expression.AsTupleConsExpr (method), Expression.variableNodes (method), Expression.asQuoted (method), Expression.AsQuotedLiteral (method), Traversal (type), newTraversal (func), TraverseName (type), newTraverseName (func), TraverseIndex (type), newTraverseIndex (func) |
| ast_expression_test.go | 210 | NEW | TestExpressionVariablesObjectCons (func), TestExpressionVariablesFunctionCallAndTupleCons (func), TestExpressionAsFunctionCallExprAndTupleConsExpr (func) |
| ast_function_call_expr.go | 64 | NEW | FunctionCallExpr (type), newFunctionCallExpr (func), FunctionCallExpr.Name (method), FunctionCallExpr.variableNodes (method) |
| ast_function_call_expr_test.go | 38 | NEW | TestFunctionCallExprName (func) |
| ast_object_cons_expr.go | 373 | CHANGED | ObjectConsExpr (type), newObjectConsExpr (func), ObjectConsExpr.Items (method), ObjectConsExpr.ItemFor (method), ObjectConsExpr.ValueExprFor (method), ObjectConsExpr.SetItem (method), ObjectConsExpr.SetItemRaw (method), ObjectConsExpr.SetItemValue (method), ObjectConsExpr.SetItemTraversal (method), ObjectConsExpr.variableNodes (method), ObjectConsExpr.appendItem (method), ObjectConsExpr.isMultiline (method), ObjectConsExpr.splitClosing (method), ObjectConsExpr.tokenEndsLine (method), ObjectConsItem (type), newObjectConsItem (func), ObjectConsItem.KeyExpr (method), ObjectConsItem.ValueExpr (method), ObjectConsItem.init (method), ObjectConsItem.setValueExpr (method), ObjectConsKeyExpr (type), newObjectConsKeyExpr (func), ObjectConsKeyExpr.staticName (method), ObjectConsValue (type), newObjectConsValue (func), ObjectConsValue.Expr (method) |
| ast_object_cons_expr_test.go | 381 | NEW | TestObjectConsExprItemFor (func), TestObjectConsExprItems (func), TestObjectConsExprSetItemKey (func), TestObjectConsExprSetItemExisting (func), TestObjectConsExprSetItemKinds (func), TestObjectConsExprSetItemPanics (func), TestObjectConsExprSetItemLayout (func) |
| ast_test.go | 50 |  | TestTreeNode (type), makeTestTree (func) |
| ast_tuple_cons_expr.go | 34 | NEW | TupleConsExpr (type), newTupleConsExpr (func), TupleConsExpr.variableNodes (method) |
| doc.go | 14 |  |  |
| examples_test.go | 124 | CHANGED | Example_generateFromScratch (func), ExampleExpression_RenameVariablePrefix (func), ExampleObjectConsExpr_SetItemRaw (func) |
| format.go | 472 |  | format (func), formatIndent (func), formatSpaces (func), formatCells (func), spaceAfterToken (func), linesForFormat (func), tokenIsNewline (func), tokenBracketChange (func), formatLine (type) |
| format_test.go | 800 |  | TestFormat (func), TestLinesForFormat (func) |
| generate.go | 399 |  | TokensForValue (func), TokensForTraversal (func), TokensForIdentifier (func), TokensForTuple (func), TokensForObject (func), TokensForFunctionCall (func), appendTokensForValue (func), appendTokensForTraversal (func), appendTokensForTraversalStep (func), escapeQuotedStringLit (func), appendRune (func) |
| generate_test.go | 920 |  | TestTokensForValue (func), TestTokensForTraversal (func), TestTokensForTuple (func), TestTokensForObject (func), TestTokensForFunctionCall (func), TestTokenGenerateConsistency (func) |
| native_node_sorter.go | 26 |  | nativeNodeSorter (type), nativeNodeSorter.Len (method), nativeNodeSorter.Less (method), nativeNodeSorter.Swap (method) |
| node.go | 307 | CHANGED | node (type), newNode (func), node.Equal (method), node.BuildTokens (method), node.Detach (method), node.ReplaceWith (method), node.assertUnattached (method), nodeContent (type), nodes (type), nodes.BuildTokens (method), nodes.Clear (method), nodes.Append (method), nodes.AppendNode (method), nodes.Insert (method), nodes.InsertNode (method), nodes.AppendUnstructuredTokens (method), nodes.FindNodeWithContent (method), nodeSet (type), newNodeSet (func), nodeSet.Has (method), nodeSet.Add (method), nodeSet.Remove (method), nodeSet.Clear (method), nodeSet.List (method), nodeSet.FindNodeWithContent (method), internalWalkFunc (type), inTree (type), newInTree (func), inTree.assertUnattached (method), inTree.setParent (method), inTree.walkChildNodes (method), inTree.BuildTokens (method), leafNode (type), leafNode.walkChildNodes (method) |
| parser.go | 810 | CHANGED | parse (func), inputTokens (type), inputTokens.Partition (method), inputTokens.PartitionType (method), inputTokens.PartitionTypeOk (method), inputTokens.PartitionTypeSingle (method), inputTokens.PartitionIncludingComments (method), inputTokens.PartitionBlockItem (method), inputTokens.PartitionLeadComments (method), inputTokens.PartitionLineEndTokens (method), inputTokens.Slice (method), inputTokens.Len (method), inputTokens.Tokens (method), inputTokens.Types (method), parseBody (func), parseBodyItem (func), parseAttribute (func), parseBlock (func), parseBlockLabels (func), parseExpression (func), parseAnyExpression (func), parseObjectConsExpr (func), parseObjectConsKeyExpr (func), parseFunctionCallExpr (func), parseTupleConsExpr (func), parseTemplateExpr (func), parseTraversal (func), parseTraversalStep (func), writerTokens (func), partitionTokens (func), partitionLeadCommentTokens (func), partitionLineEndTokens (func), lexConfig (func) |
| parser_test.go | 2308 | CHANGED | TestParse (func), TestPartitionTokens (func), TestPartitionLeadCommentTokens (func), TestLexConfig (func), TestParseFunctionCallAndTupleConsRoundTrip (func) |
| public.go | 48 |  | NewFile (func), ParseConfig (func), Format (func) |
| round_trip_test.go | 181 |  | TestRoundTripVerbatim (func), TestRoundTripFormat (func), TestRoundTripSafeConcurrent (func) |
| tokens.go | 136 |  | Token (type), Token.asHCLSyntax (method), Tokens (type), Tokens.Bytes (method), Tokens.testValue (method), Tokens.Columns (method), Tokens.WriteTo (method), Tokens.walkChildNodes (method), Tokens.BuildTokens (method), ObjectAttrTokens (type), newIdentToken (func) |

#### x/ast_expression.go

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
// expression, including those inside the items of an object-construct
// expression, the arguments of a function call expression and the elements
// of a tuple-construct expression, in source order.
func (e *Expression) Variables() []*Traversal {
	return e.traversals(make([]*Traversal, 0))
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
func (e *Expression) RenameVariablePrefix(search, replacement []string) {
	if len(search) != len(replacement) {
		panic(fmt.Sprintf("search and replacement length mismatch (%d and %d)", len(search), len(replacement)))
	}
	for _, traversal := range e.traversals(nil) {
		traversal.renamePrefix(search, replacement)
	}
}

func (e *Expression) AsObjectConsExpr() *ObjectConsExpr {
	if e.wrapped == nil {
		return nil
	}

	if object, ok := e.wrapped.content.(*ObjectConsExpr); ok {
		return object
	}

	return nil
}

// AsFunctionCallExpr returns the wrapped function call expression, or nil if
// this expression does not parse as a function call expression.
func (e *Expression) AsFunctionCallExpr() *FunctionCallExpr {
	if e.wrapped == nil {
		return nil
	}

	if call, ok := e.wrapped.content.(*FunctionCallExpr); ok {
		return call
	}

	return nil
}

// AsTupleConsExpr returns the wrapped tuple-construct expression, or nil if
// this expression does not parse as a tuple-construct expression.
func (e *Expression) AsTupleConsExpr() *TupleConsExpr {
	if e.wrapped == nil {
		return nil
	}

	if tuple, ok := e.wrapped.content.(*TupleConsExpr); ok {
		return tuple
	}

	return nil
}

// traversals appends the absolute traversals within the expression to to, in
// source order, and returns the result. For an object-construct expression,
// a function call expression or a tuple-construct expression, these are the
// traversals of its items, arguments or elements.
//
func (e *Expression) traversals(to []*Traversal) []*Traversal {
	for _, node := range e.absTraversals.List() {
		to = append(to, node.content.(*Traversal))
	}
	if e.wrapped == nil {
		return to
	}
	switch wrapped := e.wrapped.content.(type) {
	case *ObjectConsExpr:
		to = wrapped.traversals(to)
	case *FunctionCallExpr:
		to = wrapped.traversals(to)
	case *TupleConsExpr:
		to = wrapped.traversals(to)
	}
	return to
}

func (e *Expression) asQuoted() *quoted {
	if e.wrapped == nil {
		return nil
	}

	if quoted, ok := e.wrapped.content.(*quoted); ok {
		return quoted
	}

	return nil
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

// renamePrefix replaces the names of the traversal's leading steps with
// replacement when they match search. Only name steps can match.
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

#### x/ast_expression_test.go

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

func TestExpressionVariablesNested(t *testing.T) {
	tests := []struct {
		src     string
		want    []string // the traversals Variables reports, in order
		renamed string   // the source after renaming old to new
	}{
		{
			src:     "a = { k = old.x }\n",
			want:    []string{"old.x"},
			renamed: "a = { k = new.x }\n",
		},
		{
			src:     "a = { k = old.x, j = old.y }\n",
			want:    []string{"old.x", "old.y"},
			renamed: "a = { k = new.x, j = new.y }\n",
		},
		{
			src:     "a = { k = { j = old.x } }\n",
			want:    []string{"old.x"},
			renamed: "a = { k = { j = new.x } }\n",
		},
		{
			src:     "a = { k = [old.x] }\n",
			want:    []string{"old.x"},
			renamed: "a = { k = [new.x] }\n",
		},
		{
			src:     "a = { k = f(old.x) }\n",
			want:    []string{"old.x"},
			renamed: "a = { k = f(new.x) }\n",
		},
		{
			// A naked identifier key is a literal name, not a variable.
			src:     "a = { old = 1 }\n",
			want:    nil,
			renamed: "a = { old = 1 }\n",
		},
		{
			src:     "a = { old = old.x }\n",
			want:    []string{"old.x"},
			renamed: "a = { old = new.x }\n",
		},
		{
			src:     "a = { (old.x) = 1 }\n",
			want:    []string{"old.x"},
			renamed: "a = { (new.x) = 1 }\n",
		},
		{
			src:     "a = { (old.x) = old.y }\n",
			want:    []string{"old.x", "old.y"},
			renamed: "a = { (new.x) = new.y }\n",
		},
		{
			src:     "a = { \"k\" = old.x }\n",
			want:    []string{"old.x"},
			renamed: "a = { \"k\" = new.x }\n",
		},
		{
			src:     "a = { k = { old = 1 } }\n",
			want:    nil,
			renamed: "a = { k = { old = 1 } }\n",
		},
		{
			src:     "a = {}\n",
			want:    nil,
			renamed: "a = {}\n",
		},
		{
			src:     "a = f(old.x)\n",
			want:    []string{"old.x"},
			renamed: "a = f(new.x)\n",
		},
		{
			src:     "a = f(old.x, g(old.y), [old.z])\n",
			want:    []string{"old.x", "old.y", "old.z"},
			renamed: "a = f(new.x, g(new.y), [new.z])\n",
		},
		{
			src:     "a = f(old.x...)\n",
			want:    []string{"old.x"},
			renamed: "a = f(new.x...)\n",
		},
		{
			src:     "a = provider::p::f(old.x)\n",
			want:    []string{"old.x"},
			renamed: "a = provider::p::f(new.x)\n",
		},
		{
			src:     "a = [old.x, { k = old.y }, [old.z]]\n",
			want:    []string{"old.x", "old.y", "old.z"},
			renamed: "a = [new.x, { k = new.y }, [new.z]]\n",
		},
		{
			src:     "a = { k = f(old.x), j = [old.y] }\n",
			want:    []string{"old.x", "old.y"},
			renamed: "a = { k = f(new.x), j = [new.y] }\n",
		},
		{
			src:     "a = [f(old.x), old.y.z]\n",
			want:    []string{"old.x", "old.y.z"},
			renamed: "a = [f(new.x), new.y.z]\n",
		},
		{
			src:     "a = f()\n",
			want:    nil,
			renamed: "a = f()\n",
		},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
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
				t.Errorf("wrong traversals (-want +got):\n%s", diff)
			}

			native, diags := hclsyntax.ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected hclsyntax diagnostics: %s", diags.Error())
			}
			nativeVars := native.Body.(*hclsyntax.Body).Attributes["a"].Expr.Variables()
			if len(got) != len(nativeVars) {
				t.Errorf("got %d traversals, but hclsyntax reports %d", len(got), len(nativeVars))
			}

			expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
			if got := string(f.Bytes()); got != test.renamed {
				t.Errorf("wrong result after renaming\ngot:  %q\nwant: %q", got, test.renamed)
			}
		})
	}
}

func TestExpressionAsFunctionCallExprAndTupleConsExpr(t *testing.T) {
	tests := []struct {
		src       string // the expression assigned to a
		wantCall  bool
		wantTuple bool
	}{
		{`f(1)`, true, false},
		{`[1, 2]`, false, true},
		{`[f(1)]`, false, true},
		{`(f(1))`, false, false},
		{`f(1)[0]`, false, false},
		{`[for x in xs : x]`, false, false},
		{`{ k = [1] }`, false, false},
		{`x`, false, false},
		{`"f(1)"`, false, false},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte("a = "+test.src+"\n"), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			expr := f.Body().GetAttribute("a").Expr()
			if got := expr.AsFunctionCallExpr() != nil; got != test.wantCall {
				t.Errorf("AsFunctionCallExpr returned a call: %v, want %v", got, test.wantCall)
			}
			if got := expr.AsTupleConsExpr() != nil; got != test.wantTuple {
				t.Errorf("AsTupleConsExpr returned a tuple: %v, want %v", got, test.wantTuple)
			}
		})
	}
}

```

#### x/ast_function_call_expr.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0


package hclwrite

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// FunctionCallExpr represents a function call expression such as
// join(" ", list). Use Expression.AsFunctionCallExpr to get one from an
// expression.
//
// Each argument is an Expression of its own. The function name, the
// parentheses, the commas and an expansion symbol (...) after the last
// argument are kept as tokens of the call.
type FunctionCallExpr struct {
	inTree

	args nodeSet
}

func newFunctionCallExpr() *FunctionCallExpr {
	return &FunctionCallExpr{
		inTree: newInTree(),
		args:   newNodeSet(),
	}
}

// Name returns the name of the called function. For a function in a
// namespace, the name includes the namespace, as in provider::aws::arn_parse.
func (call *FunctionCallExpr) Name() string {
	var name []byte
	for _, n := range call.children.List() {
		tokens, ok := n.content.(Tokens)
		if !ok {
			// The name and the opening parenthesis come before any argument.
			break
		}
		for _, tok := range tokens {
			switch tok.Type {
			case hclsyntax.TokenOParen:
				return string(name)
			case hclsyntax.TokenIdent, hclsyntax.TokenDoubleColon:
				name = append(name, tok.Bytes...)
			}
		}
	}
	return string(name)
}

// traversals appends the absolute traversals in the arguments to to, in
// source order, and returns the result.
//
func (call *FunctionCallExpr) traversals(to []*Traversal) []*Traversal {
	for _, n := range call.args.List() {
		to = n.content.(*Expression).traversals(to)
	}
	return to
}

```

#### x/ast_function_call_expr_test.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0


package hclwrite

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
)

func TestFunctionCallExprName(t *testing.T) {
	tests := []struct {
		src  string // the function call assigned to a
		want string
	}{
		{`join(" ", ["fedora", "fez"])`, "join"},
		{`f()`, "f"},
		{`f(xs...)`, "f"},
		{`provider::aws::arn_parse(x)`, "provider::aws::arn_parse"},
		{"f(\n  1,\n)", "f"},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte("a = "+test.src+"\n"), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			call := f.Body().GetAttribute("a").Expr().AsFunctionCallExpr()
			if call == nil {
				t.Fatalf("not parsed as a function call expression")
			}
			if got := call.Name(); got != test.want {
				t.Errorf("wrong name %q, want %q", got, test.want)
			}
		})
	}
}

```

#### x/ast_object_cons_expr.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0


package hclwrite

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

// ObjectConsExpr is an object-construct expression, such as
// { hat = "derby", (cat) = "calico" }. Use Expression.AsObjectConsExpr to
// get one from an expression.
//
// An item can be looked up and set by its key only when the key is written
// as an identifier or as a quoted string literal. The two forms are
// equivalent: the key "hat" names both hat = ... and "hat" = .... A key
// that is taken from a variable or another expression, such as (hat), is
// never found by its name.
type ObjectConsExpr struct {
	inTree

	items nodeSet
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
	items := make([]*ObjectConsItem, 0, len(nodes))
	for _, n := range nodes {
		items = append(items, n.content.(*ObjectConsItem))
	}
	return items
}

// ItemFor returns the first item whose key is the given string, or nil if
// there is no such item.
//
// A quoted key is compared by its value after its escape sequences are
// decoded, so ItemFor("${hat}") finds "$${hat}" = .... Keys that are not an
// identifier or a quoted string literal never match.
func (o *ObjectConsExpr) ItemFor(key string) *ObjectConsItem {
	for _, item := range o.Items() {
		if item.KeyExpr().hasName(key) {
			return item
		}
	}
	return nil
}

// ValueExprFor returns the value of the item that ItemFor finds for the
// given key, or nil if there is no such item.
func (o *ObjectConsExpr) ValueExprFor(key string) *ObjectConsValue {
	if item := o.ItemFor(key); item != nil {
		return item.ValueExpr()
	}
	return nil
}

// SetItem either replaces the value of the item that ItemFor finds for the
// given key, or adds a new item to the end of the object, with the given
// expression as its value. The key of an existing item is left as written.
//
// A new key is written as an identifier if it is a valid one, and otherwise
// as a quoted string literal that evaluates to the given key. The key "for"
// is always quoted, since HCL would read it as a for expression.
//
// SetItem panics if the expression is already part of a syntax tree, since
// two items would then share it.
//
// The return values are the key and the value of the item that was either
// modified in-place or created.
func (o *ObjectConsExpr) SetItem(key string, expr *Expression) (*ObjectConsKeyExpr, *ObjectConsValue) {
	expr.assertUnattached()

	item := o.ItemFor(key)
	if item != nil {
		value := item.ValueExpr()
		value.expr = value.expr.ReplaceWith(expr)
	} else {
		item = o.appendItem(key, expr)
	}
	return item.KeyExpr(), item.ValueExpr()
}

// SetItemRaw is like SetItem, but uses the given tokens verbatim as the
// value.
//
// The same caveats apply to this function as for NewExpressionRaw on which
// it is based. If possible, prefer to use SetItemValue or SetItemTraversal.
func (o *ObjectConsExpr) SetItemRaw(key string, tokens Tokens) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionRaw(tokens))
}

// SetItemValue is like SetItem, but writes the given value as a literal, as
// NewExpressionLiteral does. To set a variable reference or other
// traversal, use SetItemTraversal.
func (o *ObjectConsExpr) SetItemValue(key string, val cty.Value) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionLiteral(val))
}

// SetItemTraversal is like SetItem, but writes the given traversal, which
// must be absolute, as NewExpressionAbsTraversal does. To set a literal
// value, use SetItemValue.
func (o *ObjectConsExpr) SetItemTraversal(key string, traversal hcl.Traversal) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionAbsTraversal(traversal))
}

// traversals appends the absolute traversals in the items' keys and values to
// to, in source order, and returns the result.
//
func (o *ObjectConsExpr) traversals(to []*Traversal) []*Traversal {
	for _, item := range o.Items() {
		to = item.KeyExpr().traversals(to)
		to = item.ValueExpr().Expr().traversals(to)
	}
	return to
}

// appendItem adds a new item with the given key and value after the last
// item of the object, and returns it.
//
// The tokens that follow the last item on its line, such as a comma or a
// line comment, stay with that item. Everything after them stays with the
// closing brace. In an object written on a single line the new item follows
// a comma, and otherwise it goes on its own line.
func (o *ObjectConsExpr) appendItem(key string, expr *Expression) *ObjectConsItem {
	closing := o.children.Last()
	tail, ok := closing.content.(Tokens)
	if !ok {
		panic("object-construct expression does not end with its closing brace")
	}
	multiline := o.multiline()

	split := 0
	for split < len(tail) && tail[split].Type == hclsyntax.TokenComma {
		split++
	}
	if multiline {
		for i, tok := range tail {
			if tokenIsNewline(tok) {
				split = i + 1
				break
			}
		}
	}
	before := append(Tokens{}, tail[:split]...)
	rest := append(Tokens{}, tail[split:]...)

	switch {
	case multiline:
		if len(before) == 0 || !tokenIsNewline(before[len(before)-1]) {
			before = append(before, &Token{Type: hclsyntax.TokenNewline, Bytes: []byte{'\n'}})
		}
		rest = append(Tokens{{Type: hclsyntax.TokenNewline, Bytes: []byte{'\n'}}}, rest...)
	case len(before) == 0 && len(o.items) > 0:
		before = Tokens{{Type: hclsyntax.TokenComma, Bytes: []byte{','}}}
	}

	item := newObjectConsItem()
	item.init(before, key, expr)
	o.items.Add(o.children.Insert(closing, item))
	closing.ReplaceWith(rest)
	return item
}

// multiline reports whether the object is written on several lines. Only
// the line breaks between its items count, not those inside a key or a
// value.
func (o *ObjectConsExpr) multiline() bool {
	for _, n := range o.children.List() {
		switch content := n.content.(type) {
		case Tokens:
			if objectConsBreaksLine(content) {
				return true
			}
		case *ObjectConsItem:
			for _, m := range content.children.List() {
				if tokens, ok := m.content.(Tokens); ok && objectConsBreaksLine(tokens) {
					return true
				}
			}
		}
	}
	return false
}

// objectConsBreaksLine reports whether the tokens between the parts of an
// object contain a line break.
func objectConsBreaksLine(tokens Tokens) bool {
	for _, tok := range tokens {
		if tokenIsNewline(tok) {
			return true
		}
	}
	return false
}

// ObjectConsItem is one item of an object-construct expression: a key and
// the value assigned to it.
type ObjectConsItem struct {
	inTree

	key *node
	value *node
}

func newObjectConsItem() *ObjectConsItem {
	return &ObjectConsItem{
		inTree: newInTree(),
	}
}

// KeyExpr returns the key of the item.
func (item *ObjectConsItem) KeyExpr() *ObjectConsKeyExpr {
	return item.key.content.(*ObjectConsKeyExpr)
}

// ValueExpr returns the value of the item.
func (item *ObjectConsItem) ValueExpr() *ObjectConsValue {
	return item.value.content.(*ObjectConsValue)
}

// init builds a new item with the given key and value, after the given
// tokens that separate it from what comes before.
func (item *ObjectConsItem) init(before Tokens, key string, expr *Expression) {
	item.children.AppendUnstructuredTokens(before)
	item.key = item.children.Append(newObjectConsKeyExprNamed(key))
	item.children.AppendUnstructuredTokens(Tokens{
		{
			Type:  hclsyntax.TokenEqual,
			Bytes: []byte{'='},
		},
	})
	value := newObjectConsValue()
	value.expr = value.children.Append(expr)
	item.value = item.children.Append(value)
}

// ObjectConsKeyExpr is the key of an item of an object-construct
// expression. It is written as an identifier, as a quoted string, or as
// another expression, such as a variable in parentheses.
type ObjectConsKeyExpr struct {
	inTree

	wrapped *node

	// name is the string the key evaluates to, when literal is true.
	name string

	// literal is true when the key is a naked identifier or a quoted string
	// literal. Such a key names the item instead of referring to a variable.
	literal bool
}

func newObjectConsKeyExpr(wrapped *node, name string, literal bool) *ObjectConsKeyExpr {
	expr := &ObjectConsKeyExpr{
		inTree:  newInTree(),
		wrapped: wrapped,
		name:    name,
		literal: literal,
	}
	expr.children.AppendNode(wrapped)

	return expr
}

// newObjectConsKeyExprNamed builds a key that evaluates to name. It is an
// identifier if name is a valid one, and a quoted string literal otherwise.
//
// The identifier for is quoted too: as the first key of an object, HCL would
// read it as the start of a for expression.
func newObjectConsKeyExprNamed(name string) *ObjectConsKeyExpr {
	var expr *Expression
	if hclsyntax.ValidIdentifier(name) && name != "for" {
		expr = NewExpressionAbsTraversal(hcl.Traversal{hcl.TraverseRoot{Name: name}})
	} else {
		expr = NewExpressionLiteral(cty.StringVal(name))
	}
	return newObjectConsKeyExpr(newNode(expr), name, true)
}

// hasName reports whether the key is a literal name equal to name.
func (k *ObjectConsKeyExpr) hasName(name string) bool {
	return k.literal && k.name == name
}

// traversals appends the absolute traversals in the key to to, and returns
// the result. A literal name contributes none.
func (k *ObjectConsKeyExpr) traversals(to []*Traversal) []*Traversal {
	if k.literal {
		return to
	}
	return k.wrapped.content.(*Expression).traversals(to)
}

// ObjectConsValue is the value of an item of an object-construct
// expression.
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

```

#### x/ast_object_cons_expr_test.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0


package hclwrite

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

func TestObjectConsExprItemFor(t *testing.T) {
	tests := []struct {
		src   string // the object assigned to a
		key   string
		found bool
	}{
		{`{ hat = 1 }`, "hat", true},
		{`{ "hat" = 1 }`, "hat", true},
		{`{ "$${hat}" = 1 }`, "${hat}", true},
		{`{ "%%{hat}" = 1 }`, "%{hat}", true},
		{`{ "a hat" = 1 }`, "a hat", true},
		{`{ "say \"hi\"" = 1 }`, `say "hi"`, true},
		{`{ "a\\b" = 1 }`, `a\b`, true},
		{`{ "" = 1 }`, "", true},
		{`{ null = 1 }`, "null", true},
		{`{ hat: 1 }`, "hat", true},
		{`{ cat = 1, hat = 1 }`, "hat", true},
		{`{ (hat) = 1 }`, "hat", false},
		{`{ ("hat") = 1 }`, "hat", false},
		{`{ hat.band = 1 }`, "hat", false},
		{`{ hat.band = 1 }`, "hat.band", false},
		{`{ "${hat}" = 1 }`, "hat", false},
		{`{ "${hat}" = 1 }`, "${hat}", false},
		{`{ "$${hat}" = 1 }`, "$${hat}", false},
		{`{ hat = 1 }`, "Hat", false},
		{`{ cat = { hat = 1 } }`, "hat", false},
		{`{}`, "hat", false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s %q", test.src, test.key), func(t *testing.T) {
			f, diags := ParseConfig([]byte("a = "+test.src+"\n"), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

			item := object.ItemFor(test.key)
			value := object.ValueExprFor(test.key)
			if found := item != nil; found != test.found {
				t.Fatalf("ItemFor found an item: %v, want %v", found, test.found)
			}
			if !test.found {
				if value != nil {
					t.Errorf("ValueExprFor returned %q, want nil", value.BuildTokens(nil).Bytes())
				}
				return
			}
			if value != item.ValueExpr() {
				t.Errorf("ValueExprFor returned another value than ItemFor")
			}
			if got := strings.TrimSpace(string(value.Expr().BuildTokens(nil).Bytes())); got != "1" {
				t.Errorf("wrong value %q", got)
			}
		})
	}
}

func TestObjectConsExprItems(t *testing.T) {
	src := "a = { hat = \"derby\", (cat) = \"calico\", \"dog\" = 3 }\n"
	f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics: %s", diags.Error())
	}
	object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

	var got [][2]string
	for _, item := range object.Items() {
		got = append(got, [2]string{
			strings.TrimSpace(string(item.KeyExpr().BuildTokens(nil).Bytes())),
			strings.TrimSpace(string(item.ValueExpr().Expr().BuildTokens(nil).Bytes())),
		})
	}
	want := [][2]string{
		{"hat", `"derby"`},
		{"(cat)", `"calico"`},
		{`"dog"`, "3"},
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("wrong items (-want +got):\n%s", diff)
	}
}

func TestObjectConsExprSetItemKey(t *testing.T) {
	tests := []struct {
		key  string
		want string // the key as written
	}{
		{"hat", "hat"},
		{"hat_band", "hat_band"},
		{"hat-band", "hat-band"},
		{"null", "null"},
		{"for", `"for"`},
		{"if", "if"},
		{"a hat", `"a hat"`},
		{"hat.band", `"hat.band"`},
		{"1hat", `"1hat"`},
		{`say "hi"`, `"say \"hi\""`},
		{`a\b`, `"a\\b"`},
		{"${hat}", `"$${hat}"`},
		{"%{hat}", `"%%{hat}"`},
		{"$hat", `"$hat"`},
		{"", `""`},
	}

	for _, test := range tests {
		t.Run(test.key, func(t *testing.T) {
			f, diags := ParseConfig([]byte("a = { cat = 1 }\n"), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

			key, value := object.SetItemValue(test.key, cty.NumberIntVal(2))
			if got := strings.TrimSpace(string(key.BuildTokens(nil).Bytes())); got != test.want {
				t.Errorf("wrong key\ngot:  %s\nwant: %s", got, test.want)
			}
			if item := object.ItemFor(test.key); item == nil || item.KeyExpr() != key || item.ValueExpr() != value {
				t.Errorf("ItemFor does not find the item just set")
			}

			// The key must evaluate back to the original string.
			native, diags := hclsyntax.ParseConfig(f.Bytes(), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("result is not valid HCL: %s\n%s", diags.Error(), f.Bytes())
			}
			val, diags := native.Body.(*hclsyntax.Body).Attributes["a"].Expr.Value(nil)
			if diags.HasErrors() {
				t.Fatalf("result does not evaluate: %s\n%s", diags.Error(), f.Bytes())
			}
			wantVal := cty.ObjectVal(map[string]cty.Value{
				"cat":    cty.NumberIntVal(1),
				test.key: cty.NumberIntVal(2),
			})
			if !val.RawEquals(wantVal) {
				t.Errorf("wrong value\ngot:  %#v\nwant: %#v", val, wantVal)
			}
		})
	}
}

func TestObjectConsExprSetItemExisting(t *testing.T) {
	tests := []struct {
		src  string
		key  string
		want string
	}{
		{
			"a = { hat = \"derby\" }\n",
			"hat",
			"a = { hat = \"bowler\" }\n",
		},
		{
			"a = { \"hat\" = \"derby\", cat = 1 }\n",
			"hat",
			"a = { \"hat\" = \"bowler\", cat = 1 }\n",
		},
		{
			"a = { \"$${hat}\" = \"derby\" }\n",
			"${hat}",
			"a = { \"$${hat}\" = \"bowler\" }\n",
		},
		{
			"a = {\n  hat = \"derby\" # felt\n  cat = 1\n}\n",
			"hat",
			"a = {\n  hat = \"bowler\" # felt\n  cat = 1\n}\n",
		},
		{
			// A key from a variable is not the one that is set.
			"a = { (hat) = \"derby\" }\n",
			"hat",
			"a = { (hat) = \"derby\", hat = \"bowler\" }\n",
		},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

			key, value := object.SetItemValue(test.key, cty.StringVal("bowler"))
			if item := object.ItemFor(test.key); item == nil || item.KeyExpr() != key || item.ValueExpr() != value {
				t.Errorf("SetItemValue did not return the item that ItemFor finds")
			}
			if got := string(f.Bytes()); got != test.want {
				t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, test.want)
			}
		})
	}
}

func TestObjectConsExprSetItemHelpers(t *testing.T) {
	f, diags := ParseConfig([]byte("a = { hat = 1 }\n"), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics: %s", diags.Error())
	}
	object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

	object.SetItemRaw("hat", TokensForValue(cty.StringVal("derby")))
	object.SetItemValue("cat", cty.ListVal([]cty.Value{cty.StringVal("calico")}))
	_, value := object.SetItemTraversal("dog", hcl.Traversal{
		hcl.TraverseRoot{Name: "var"},
		hcl.TraverseAttr{Name: "dog"},
	})

	want := "a = { hat = \"derby\", cat = [\"calico\"], dog = var.dog }\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result\ngot:  %s\nwant: %s", got, want)
	}

	var got []string
	for _, traversal := range f.Body().GetAttribute("a").Expr().Variables() {
		got = append(got, strings.TrimSpace(string(traversal.BuildTokens(nil).Bytes())))
	}
	if diff := cmp.Diff([]string{"var.dog"}, got); diff != "" {
		t.Errorf("wrong traversals (-want +got):\n%s", diff)
	}
	if got := value.Expr().Variables(); len(got) != 1 {
		t.Errorf("the value has %d traversals, want 1", len(got))
	}
}

func TestObjectConsExprSetItemAttached(t *testing.T) {
	tests := map[string]func(object *ObjectConsExpr) *Expression{
		"attached by the parser": func(object *ObjectConsExpr) *Expression {
			return object.ValueExprFor("hat").Expr()
		},
		"attached by SetItem": func(object *ObjectConsExpr) *Expression {
			expr := NewExpressionLiteral(cty.StringVal("bowler"))
			object.SetItem("cat", expr)
			return expr
		},
		"attached to an attribute": func(object *ObjectConsExpr) *Expression {
			f, _ := ParseConfig([]byte("b = 1\n"), "", hcl.InitialPos)
			return f.Body().GetAttribute("b").Expr()
		},
	}

	for name, attached := range tests {
		t.Run(name, func(t *testing.T) {
			f, diags := ParseConfig([]byte("a = { hat = \"derby\" }\n"), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()
			expr := attached(object)
			before := string(f.Bytes())

			for _, key := range []string{"hat", "dog"} {
				func() {
					defer func() {
						if recover() == nil {
							t.Errorf("SetItem(%q) did not panic", key)
						}
					}()
					object.SetItem(key, expr)
				}()
			}
			if got := string(f.Bytes()); got != before {
				t.Errorf("SetItem changed the object before panicking\ngot:  %s\nwant: %s", got, before)
			}
		})
	}

	t.Run("replaced", func(t *testing.T) {
		f, diags := ParseConfig([]byte("a = { hat = \"derby\" }\n"), "", hcl.InitialPos)
		if diags.HasErrors() {
			t.Fatalf("unexpected diagnostics: %s", diags.Error())
		}
		object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

		// An expression that SetItem replaced is no longer in the tree, so it
		// can be used again.
		old := object.ValueExprFor("hat").Expr()
		object.SetItemValue("hat", cty.StringVal("bowler"))
		object.SetItem("dog", old)

		want := "a = { hat = \"bowler\", dog = \"derby\" }\n"
		if got := string(f.Bytes()); got != want {
			t.Errorf("wrong result\ngot:  %s\nwant: %s", got, want)
		}
	})
}

func TestObjectConsExprSetItemLayout(t *testing.T) {
	tests := []struct {
		name string
		src  string
		keys []string // the keys added, each set to "bowler"
		want string
	}{
		{
			"single line",
			`a = { hat = "derby" }`,
			[]string{"dog"},
			`a = { hat = "derby", dog = "bowler" }`,
		},
		{
			"empty",
			`a = {}`,
			[]string{"dog"},
			`a = { dog = "bowler" }`,
		},
		{
			"empty, two items",
			`a = {}`,
			[]string{"dog", "cat"},
			`a = { dog = "bowler", cat = "bowler" }`,
		},
		{
			"single line, two items",
			`a = { hat = "derby" }`,
			[]string{"dog", "cat"},
			`a = { hat = "derby", dog = "bowler", cat = "bowler" }`,
		},
		{
			"single line, trailing comma",
			`a = { hat = "derby", }`,
			[]string{"dog"},
			`a = { hat = "derby", dog = "bowler" }`,
		},
		{
			"single line, colon",
			`a = { hat: "derby" }`,
			[]string{"dog"},
			`a = { hat : "derby", dog = "bowler" }`,
		},
		{
			"single line, inline comment before the brace",
			`a = { hat = "derby" /* end of a */ }`,
			[]string{"dog"},
			`a = { hat = "derby", dog = "bowler" /* end of a */ }`,
		},
		{
			"single line, nested multi-line value",
			"a = { hat = {\n  band = \"silk\"\n} }",
			[]string{"dog"},
			"a = { hat = {\n  band = \"silk\"\n}, dog = \"bowler\" }",
		},
		{
			"multi-line",
			"a = {\n  hat = \"derby\"\n}",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}",
		},
		{
			"multi-line, two items",
			"a = {\n  hat = \"derby\"\n}",
			[]string{"dog", "cat"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n}",
		},
		{
			"multi-line, empty",
			"a = {\n}",
			[]string{"dog"},
			"a = {\n  dog = \"bowler\"\n}",
		},
		{
			"multi-line, longer key",
			"a = {\n  hat = \"derby\"\n}",
			[]string{"bowler"},
			"a = {\n  hat    = \"derby\"\n  bowler = \"bowler\"\n}",
		},
		{
			"multi-line, comma",
			"a = {\n  hat = \"derby\",\n}",
			[]string{"dog"},
			"a = {\n  hat = \"derby\",\n  dog = \"bowler\"\n}",
		},
		{
			"multi-line, line comment",
			"a = {\n  hat = \"derby\" # felt\n}",
			[]string{"dog"},
			"a = {\n  hat = \"derby\" # felt\n  dog = \"bowler\"\n}",
		},
		{
			"multi-line, comma and line comment",
			"a = {\n  hat = \"derby\", // felt\n}",
			[]string{"dog"},
			"a = {\n  hat = \"derby\", // felt\n  dog = \"bowler\"\n}",
		},
		{
			"multi-line, comments before the brace",
			"a = {\n  hat = \"derby\"\n  # more hats\n  /* to come */\n}",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  # more hats\n  /* to come */\n}",
		},
		{
			"multi-line, empty with comment",
			"a = {\n  # no hats\n}",
			[]string{"dog"},
			"a = {\n  dog = \"bowler\"\n  # no hats\n}",
		},
		{
			"multi-line, comment after the brace",
			"a = { # hats\n}",
			[]string{"dog"},
			"a = { # hats\n  dog = \"bowler\"\n}",
		},
		{
			"multi-line, brace after the last item",
			"a = {\n  hat = \"derby\" }",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}",
		},
		{
			"multi-line, brace after the last item, two items",
			"a = {\n  hat = \"derby\" }",
			[]string{"dog", "cat"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n}",
		},
		{
			"multi-line, comma and brace after the last item",
			"a = {\n  hat = \"derby\", }",
			[]string{"dog"},
			"a = {\n  hat = \"derby\",\n  dog = \"bowler\"\n}",
		},
		{
			"multi-line, first item on the brace's line",
			"a = { hat = \"derby\"\n  cat = \"calico\" }",
			[]string{"dog"},
			"a = { hat = \"derby\"\n  cat = \"calico\"\n  dog = \"bowler\"\n}",
		},
		{
			"multi-line, inline comment before the brace",
			"a = {\n  hat = \"derby\" /* end of a */ }",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n/* end of a */ }",
		},
		{
			"multi-line, nested",
			"a = {\n  hat = {\n    band = \"silk\"\n  }\n}",
			[]string{"dog"},
			"a = {\n  hat = {\n    band = \"silk\"\n  }\n  dog = \"bowler\"\n}",
		},
		{
			"empty, keyword key",
			`a = {}`,
			[]string{"for"},
			`a = { "for" = "bowler" }`,
		},
		{
			"nested object set",
			"a = { hat = {} }",
			[]string{"hat"},
			`a = { hat = "bowler" }`,
		},
	}

	bowler := TokensForValue(cty.StringVal("bowler"))
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src+"\n"), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

			for _, key := range test.keys {
				object.SetItemRaw(key, bowler)
			}

			got := string(f.Bytes())
			want := test.want + "\n"
			if got != want {
				t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, want)
			}
			if formatted := string(Format([]byte(got))); formatted != got {
				t.Errorf("result is not in canonical form\ngot:\n%s\nformatted:\n%s", got, formatted)
			}
			if _, diags := hclsyntax.ParseConfig([]byte(got), "", hcl.InitialPos); diags.HasErrors() {
				t.Errorf("result is not valid HCL: %s", diags.Error())
			}
			for _, key := range test.keys {
				if object.ItemFor(key) == nil {
					t.Errorf("ItemFor(%q) does not find the item just set", key)
				}
			}
		})
	}
}

```

#### x/ast_tuple_cons_expr.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0


package hclwrite

// TupleConsExpr represents a tuple-construct expression such as
// ["fedora", "fez"]. Use Expression.AsTupleConsExpr to get one from an
// expression.
//
// Each element is an Expression of its own. The brackets and the commas are
// kept as tokens of the tuple.
type TupleConsExpr struct {
	inTree

	elems nodeSet
}

func newTupleConsExpr() *TupleConsExpr {
	return &TupleConsExpr{
		inTree: newInTree(),
		elems:  newNodeSet(),
	}
}

// traversals appends the absolute traversals in the elements to to, in
// source order, and returns the result.
//
func (tuple *TupleConsExpr) traversals(to []*Traversal) []*Traversal {
	for _, n := range tuple.elems.List() {
		to = n.content.(*Expression).traversals(to)
	}
	return to
}

```

#### x/examples_test.go

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

func ExampleObjectConsExpr_SetItemRaw() {
	config := []byte(`a = {
    hat = "derby", (cat) = "calico" }` + "\n")

	bowler := hclwrite.TokensForValue(cty.StringVal(`bowler`))

	file, _ := hclwrite.ParseConfig(config, "<inline>", hcl.Pos{})
	attr := file.Body().GetAttribute("a")

	expr := attr.Expr().AsObjectConsExpr()
	key, value := expr.SetItemRaw("hat", bowler)
	expr.SetItemRaw("dog", bowler)

	fmt.Printf("%s\n", key.BuildTokens(nil).Bytes())
	fmt.Printf("%s\n", value.BuildTokens(nil).Bytes())
	fmt.Printf("%s", file.Bytes())
	// Output:
	// hat
	// "bowler"
	// a = {
	//   hat = "bowler", (cat) = "calico"
	//   dog = "bowler"
	// }
}

```

#### x/format.go

```go
// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// format rewrites tokens within the given sequence, in-place, to adjust the
// whitespace around their content to achieve canonical formatting.
func format(tokens Tokens) {
	// Formatting is a multi-pass process. More details on the passes below,
	// but this is the overview:
	// - adjust the leading space on each line to create appropriate
	//   indentation
	// - adjust spaces between tokens in a single cell using a set of rules
	// - adjust the leading space in the "assign" and "comment" cells on each
	//   line to vertically align with neighboring lines.
	// All of these steps operate in-place on the given tokens, so a caller
	// may collect a flat sequence of all of the tokens underlying an AST
	// and pass it here and we will then indirectly modify the AST itself.
	// Formatting must change only whitespace. Specifically, that means
	// changing the SpacesBefore attribute on a token while leaving the
	// other token attributes unchanged.

	lines := linesForFormat(tokens)
	formatIndent(lines)
	formatSpaces(lines)
	formatCells(lines)
}

func formatIndent(lines []formatLine) {
	// Our methodology for indents is to take the input one line at a time
	// and count the bracketing delimiters on each line. If a line has a net
	// increase in open brackets, we increase the indent level by one and
	// remember how many new openers we had. If the line has a net _decrease_,
	// we'll compare it to the most recent number of openers and decrease the
	// dedent level by one each time we pass an indent level remembered
	// earlier.
	// The "indent stack" used here allows for us to recognize degenerate
	// input where brackets are not symmetrical within lines and avoid
	// pushing things too far left or right, creating confusion.

	// We'll start our indent stack at a reasonable capacity to minimize the
	// chance of us needing to grow it; 10 here means 10 levels of indent,
	// which should be more than enough for reasonable HCL uses.
	indents := make([]int, 0, 10)

	for i := range lines {
		line := &lines[i]
		if len(line.lead) == 0 {
			continue
		}

		if line.lead[0].Type == hclsyntax.TokenNewline {
			// Never place spaces before a newline
			line.lead[0].SpacesBefore = 0
			continue
		}

		netBrackets := 0
		for _, token := range line.lead {
			netBrackets += tokenBracketChange(token)
			if token.Type == hclsyntax.TokenOHeredoc {
				break
			}
		}

		for _, token := range line.assign {
			netBrackets += tokenBracketChange(token)
		}

		switch {
		case netBrackets > 0:
			line.lead[0].SpacesBefore = 2 * len(indents)
			indents = append(indents, netBrackets)
		case netBrackets < 0:
			closed := -netBrackets
			for closed > 0 && len(indents) > 0 {
				switch {

				case closed > indents[len(indents)-1]:
					closed -= indents[len(indents)-1]
					indents = indents[:len(indents)-1]

				case closed < indents[len(indents)-1]:
					indents[len(indents)-1] -= closed
					closed = 0

				default:
					indents = indents[:len(indents)-1]
					closed = 0
				}
			}
			line.lead[0].SpacesBefore = 2 * len(indents)
		default:
			line.lead[0].SpacesBefore = 2 * len(indents)
		}
	}
}

func formatSpaces(lines []formatLine) {
	// placeholder token used when we don't have a token but we don't want
	// to pass a real "nil" and complicate things with nil pointer checks
	nilToken := &Token{
		Type:         hclsyntax.TokenNil,
		Bytes:        []byte{},
		SpacesBefore: 0,
	}

	for _, line := range lines {
		for i, token := range line.lead {
			var before, after *Token
			if i > 0 {
				before = line.lead[i-1]
			} else {
				before = nilToken
			}
			if i < (len(line.lead) - 1) {
				after = line.lead[i+1]
			} else {
				continue
			}
			if spaceAfterToken(token, before, after) {
				after.SpacesBefore = 1
			} else {
				after.SpacesBefore = 0
			}
		}
		for i, token := range line.assign {
			if i == 0 {
				// first token in "assign" always has one space before to
				// separate the equals sign from what it's assigning.
				token.SpacesBefore = 1
			}

			var before, after *Token
			if i > 0 {
				before = line.assign[i-1]
			} else {
				before = nilToken
			}
			if i < (len(line.assign) - 1) {
				after = line.assign[i+1]
			} else {
				continue
			}
			if spaceAfterToken(token, before, after) {
				after.SpacesBefore = 1
			} else {
				after.SpacesBefore = 0
			}
		}

	}
}

func formatCells(lines []formatLine) {
	chainStart := -1
	maxColumns := 0

	// We'll deal with the "assign" cell first, since moving that will
	// also impact the "comment" cell.
	closeAssignChain := func(i int) {
		for _, chainLine := range lines[chainStart:i] {
			columns := chainLine.lead.Columns()
			spaces := (maxColumns - columns) + 1
			chainLine.assign[0].SpacesBefore = spaces
		}
		chainStart = -1
		maxColumns = 0
	}
	for i, line := range lines {
		if line.assign == nil {
			if chainStart != -1 {
				closeAssignChain(i)
			}
		} else {
			if chainStart == -1 {
				chainStart = i
			}
			columns := line.lead.Columns()
			if columns > maxColumns {
				maxColumns = columns
			}
		}
	}
	if chainStart != -1 {
		closeAssignChain(len(lines))
	}

	// Now we'll deal with the comments
	closeCommentChain := func(i int) {
		for _, chainLine := range lines[chainStart:i] {
			columns := chainLine.lead.Columns() + chainLine.assign.Columns()
			spaces := (maxColumns - columns) + 1
			chainLine.comment[0].SpacesBefore = spaces
		}
		chainStart = -1
		maxColumns = 0
	}
	for i, line := range lines {
		if line.comment == nil {
			if chainStart != -1 {
				closeCommentChain(i)
			}
		} else {
			if chainStart == -1 {
				chainStart = i
			}
			columns := line.lead.Columns() + line.assign.Columns()
			if columns > maxColumns {
				maxColumns = columns
			}
		}
	}
	if chainStart != -1 {
		closeCommentChain(len(lines))
	}
}

// spaceAfterToken decides whether a particular subject token should have a
// space after it when surrounded by the given before and after tokens.
// "before" can be TokenNil, if the subject token is at the start of a sequence.
func spaceAfterToken(subject, before, after *Token) bool {
	switch {

	case after.Type == hclsyntax.TokenNewline || after.Type == hclsyntax.TokenNil:
		// Never add spaces before a newline
		return false

	case subject.Type == hclsyntax.TokenIdent && after.Type == hclsyntax.TokenOParen:
		// Don't split a function name from open paren in a call
		return false

	case (subject.Type == hclsyntax.TokenIdent && after.Type == hclsyntax.TokenDoubleColon) ||
		(subject.Type == hclsyntax.TokenDoubleColon && after.Type == hclsyntax.TokenIdent):
		// Don't split namespace segments in a function call
		return false

	case subject.Type == hclsyntax.TokenDot || after.Type == hclsyntax.TokenDot:
		// Don't use spaces around attribute access dots
		return false

	case after.Type == hclsyntax.TokenComma || after.Type == hclsyntax.TokenEllipsis:
		// No space right before a comma or ... in an argument list
		return false

	case subject.Type == hclsyntax.TokenComma:
		// Always a space after a comma
		return true

	case subject.Type == hclsyntax.TokenQuotedLit || subject.Type == hclsyntax.TokenStringLit || subject.Type == hclsyntax.TokenOQuote || subject.Type == hclsyntax.TokenOHeredoc || after.Type == hclsyntax.TokenQuotedLit || after.Type == hclsyntax.TokenStringLit || after.Type == hclsyntax.TokenCQuote || after.Type == hclsyntax.TokenCHeredoc:
		// No extra spaces within templates
		return false

	case hclsyntax.Keyword([]byte{'i', 'n'}).TokenMatches(subject.asHCLSyntax()) && before.Type == hclsyntax.TokenIdent:
		// This is a special case for inside for expressions where a user
		// might want to use a literal tuple constructor:
		// [for x in [foo]: x]
		// ... in that case, we would normally produce in[foo] thinking that
		// in is a reference, but we'll recognize it as a keyword here instead
		// to make the result less confusing.
		return true

	case after.Type == hclsyntax.TokenOBrack && (subject.Type == hclsyntax.TokenIdent || subject.Type == hclsyntax.TokenNumberLit || tokenBracketChange(subject) < 0):
		return false

	case subject.Type == hclsyntax.TokenBang:
		// No space after a bang
		return false

	case subject.Type == hclsyntax.TokenMinus:
		// Since a minus can either be subtraction or negation, and the latter
		// should _not_ have a space after it, we need to use some heuristics
		// to decide which case this is.
		// We guess that we have a negation if the token before doesn't look
		// like it could be the end of an expression.

		switch before.Type {

		case hclsyntax.TokenNil:
			// Minus at the start of input must be a negation
			return false

		case hclsyntax.TokenOParen, hclsyntax.TokenOBrace, hclsyntax.TokenOBrack, hclsyntax.TokenEqual, hclsyntax.TokenColon, hclsyntax.TokenComma, hclsyntax.TokenQuestion:
			// Minus immediately after an opening bracket or separator must be a negation.
			return false

		case hclsyntax.TokenPlus, hclsyntax.TokenStar, hclsyntax.TokenSlash, hclsyntax.TokenPercent, hclsyntax.TokenMinus:
			// Minus immediately after another arithmetic operator must be negation.
			return false

		case hclsyntax.TokenEqualOp, hclsyntax.TokenNotEqual, hclsyntax.TokenGreaterThan, hclsyntax.TokenGreaterThanEq, hclsyntax.TokenLessThan, hclsyntax.TokenLessThanEq:
			// Minus immediately after another comparison operator must be negation.
			return false

		case hclsyntax.TokenAnd, hclsyntax.TokenOr, hclsyntax.TokenBang:
			// Minus immediately after logical operator doesn't make sense but probably intended as negation.
			return false

		default:
			return true
		}

	case subject.Type == hclsyntax.TokenOBrace || after.Type == hclsyntax.TokenCBrace:
		// Unlike other bracket types, braces have spaces on both sides of them,
		// both in single-line nested blocks foo { bar = baz } and in object
		// constructor expressions foo = { bar = baz }.
		if subject.Type == hclsyntax.TokenOBrace && after.Type == hclsyntax.TokenCBrace {
			// An open brace followed by a close brace is an exception, however.
			// e.g. foo {} rather than foo { }
			return false
		}
		return true

	// In the unlikely event that an interpolation expression is just
	// a single object constructor, we'll put a space between the ${ and
	// the following { to make this more obvious, and then the same
	// thing for the two braces at the end.
	case (subject.Type == hclsyntax.TokenTemplateInterp || subject.Type == hclsyntax.TokenTemplateControl) && after.Type == hclsyntax.TokenOBrace:
		return true
	case subject.Type == hclsyntax.TokenCBrace && after.Type == hclsyntax.TokenTemplateSeqEnd:
		return true

	// Don't add spaces between interpolated items
	case subject.Type == hclsyntax.TokenTemplateSeqEnd && (after.Type == hclsyntax.TokenTemplateInterp || after.Type == hclsyntax.TokenTemplateControl):
		return false

	case tokenBracketChange(subject) > 0:
		// No spaces after open brackets
		return false

	case tokenBracketChange(after) < 0:
		// No spaces before close brackets
		return false

	default:
		// Most tokens are space-separated
		return true

	}
}

func linesForFormat(tokens Tokens) []formatLine {
	if len(tokens) == 0 {
		return make([]formatLine, 0)
	}

	// first we'll count our lines, so we can allocate the array for them in
	// a single block. (We want to minimize memory pressure in this codepath,
	// so it can be run somewhat-frequently by editor integrations.)
	lineCount := 1 // if there are zero newlines then there is one line
	for _, tok := range tokens {
		if tokenIsNewline(tok) {
			lineCount++
		}
	}

	// To start, we'll just put everything in the "lead" cell on each line,
	// and then do another pass over the lines afterwards to adjust.
	lines := make([]formatLine, lineCount)
	li := 0
	lineStart := 0
	for i, tok := range tokens {
		if tok.Type == hclsyntax.TokenEOF {
			// The EOF token doesn't belong to any line, and terminates the
			// token sequence.
			lines[li].lead = tokens[lineStart:i]
			break
		}

		if tokenIsNewline(tok) {
			lines[li].lead = tokens[lineStart : i+1]
			lineStart = i + 1
			li++
		}
	}

	// If a set of tokens doesn't end in TokenEOF (e.g. because it's a
	// fragment of tokens from the middle of a file) then we might fall
	// out here with a line still pending.
	if lineStart < len(tokens) {
		lines[li].lead = tokens[lineStart:]
		if lines[li].lead[len(lines[li].lead)-1].Type == hclsyntax.TokenEOF {
			lines[li].lead = lines[li].lead[:len(lines[li].lead)-1]
		}
	}

	// Now we'll pick off any trailing comments and attribute assignments
	// to shuffle off into the "comment" and "assign" cells.
	for i := range lines {
		line := &lines[i]

		if len(line.lead) == 0 {
			// if the line is empty then there's nothing for us to do
			// (this should happen only for the final line, because all other
			// lines would have a newline token of some kind)
			continue
		}

		if len(line.lead) > 1 && line.lead[len(line.lead)-1].Type == hclsyntax.TokenComment {
			line.comment = line.lead[len(line.lead)-1:]
			line.lead = line.lead[:len(line.lead)-1]
		}

		for i, tok := range line.lead {
			if i > 0 && tok.Type == hclsyntax.TokenEqual {
				// We only move the tokens into "assign" if the RHS seems to
				// be a whole expression, which we determine by counting
				// brackets. If there's a net positive number of brackets
				// then that suggests we're introducing a multi-line expression.
				netBrackets := 0
				for _, token := range line.lead[i:] {
					netBrackets += tokenBracketChange(token)
				}

				if netBrackets == 0 {
					line.assign = line.lead[i:]
					line.lead = line.lead[:i]
				}
				break
			}
		}
	}

	return lines
}

// tokenIsNewline reports whether tok ends a line.
//
func tokenIsNewline(tok *Token) bool {
	switch tok.Type {
	case hclsyntax.TokenNewline:
		return true
	case hclsyntax.TokenComment:
		// Single line tokens (# and //) consume their terminating newline,
		// so we need to treat them as newline tokens as well.
		if len(tok.Bytes) > 0 && tok.Bytes[len(tok.Bytes)-1] == '\n' {
			return true
		}
	}
	return false
}

func tokenBracketChange(tok *Token) int {
	switch tok.Type {
	case hclsyntax.TokenOBrace, hclsyntax.TokenOBrack, hclsyntax.TokenOParen, hclsyntax.TokenTemplateControl, hclsyntax.TokenTemplateInterp:
		return 1
	case hclsyntax.TokenCBrace, hclsyntax.TokenCBrack, hclsyntax.TokenCParen, hclsyntax.TokenTemplateSeqEnd:
		return -1
	default:
		return 0
	}
}

// formatLine represents a single line of source code for formatting purposes,
// splitting its tokens into up to three "cells":
//
//   - lead: always present, representing everything up to one of the others
//   - assign: if line contains an attribute assignment, represents the tokens
//     starting at (and including) the equals symbol
//   - comment: if line contains any non-comment tokens and ends with a
//     single-line comment token, represents the comment.
//
// When formatting, the leading spaces of the first tokens in each of these
// cells is adjusted to align vertically their occurences on consecutive
// rows.
type formatLine struct {
	lead    Tokens
	assign  Tokens
	comment Tokens
}

```

#### x/format_test.go

```go
// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"fmt"
	"testing"

	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			``,
			``,
		},
		{
			`a=1`,
			`a = 1`,
		},
		{
			`a=b.c`,
			`a = b.c`,
		},
		{
			`a=b[c]`,
			`a = b[c]`,
		},
		{
			`a=b()[c]`,
			`a = b()[c]`,
		},
		{
			`a=["hello"][0]`,
			`a = ["hello"][0]`,
		},
		{
			`( a+2 )`,
			`(a + 2)`,
		},
		{
			`( a*2 )`,
			`(a * 2)`,
		},
		{
			`( a+-2 )`,
			`(a + -2)`,
		},
		{
			`( a*-2 )`,
			`(a * -2)`,
		},
		{
			`(-2+1)`,
			`(-2 + 1)`,
		},
		{
			`foo(1, -2,a*b, b,c)`,
			`foo(1, -2, a * b, b, c)`,
		},
		{
			`foo(a,b...)`,
			`foo(a, b...)`,
		},
		{
			`! true`,
			`!true`,
		},
		{
			`a="hello ${ name }"`,
			`a = "hello ${name}"`,
		},
		{
			`a="hello ${~ name ~}"`,
			`a = "hello ${~name~}"`,
		},
		{
			`a="${b}${c}${ d } ${e}"`,
			`a = "${b}${c}${d} ${e}"`,
		},
		{
			`"%{if true}${var.foo}%{endif}"`,
			`"%{if true}${var.foo}%{endif}"`,
		},
		{
			`b{}`,
			`b {}`,
		},
		{
			`
"${
hello
}"
`,
			`
"${
  hello
}"
`,
		},
		{
			`
foo(
1,
- 2,
a*b,
b,
c,
)
`,
			`
foo(
  1,
  -2,
  a * b,
  b,
  c,
)
`,
		},
		{
			`a?b:c`,
			`a ? b : c`,
		},
		{
			`[ [ ] ]`,
			`[[]]`,
		},
		{
			`[for x in y : x]`,
			`[for x in y : x]`,
		},
		{
			`[for x in [y] : x]`,
			`[for x in [y] : x]`,
		},
		{
			`
[
[
a
]
]
`,
			`
[
  [
    a
  ]
]
`,
		},
		{
			`
[[
a
]]
`,
			`
[[
  a
]]
`,
		},
		{
			`
[[
[
a
]
]]
`,
			`
[[
  [
    a
  ]
]]
`,
		},
		{
			// degenerate case with asymmetrical brackets
			`
[[
[
a
]]
]
`,
			`
[[
  [
    a
  ]]
]
`,
		},
		{
			`
b {
a = 1
}
`,
			`
b {
  a = 1
}
`,
		},
		{
			`
b {a = 1}
`,
			`
b { a = 1 }
`,
		},
		{
			`
a = 1
bungle = 2
`,
			`
a      = 1
bungle = 2
`,
		},
		{
			`
a = 1

bungle = 2
`,
			`
a = 1

bungle = 2
`,
		},
		{
			`
a = 1 # foo
bungle = 2
`,
			`
a      = 1 # foo
bungle = 2
`,
		},
		{
			`
a = 1 # foo
bungle = "bonce" # baz
`,
			`
a      = 1       # foo
bungle = "bonce" # baz
`,
		},
		{
			`
# here we go
a = 1 # foo
bungle = "bonce" # baz
`,
			`
# here we go
a      = 1       # foo
bungle = "bonce" # baz
`,
		},
		{
			`
foo {} # here we go
a = 1 # foo
bungle = "bonce" # baz
`,
			`
foo {}           # here we go
a      = 1       # foo
bungle = "bonce" # baz
`,
		},
		{
			`
a = 1 # foo
bungle = "bonce" # baz
zebra = "striped" # baz
`,
			`
a      = 1         # foo
bungle = "bonce"   # baz
zebra  = "striped" # baz
`,
		},
		{
			`
a = 1 # foo
bungle = (
    "bonce"
) # baz
zebra = "striped" # baz
`,
			`
a = 1 # foo
bungle = (
  "bonce"
)                 # baz
zebra = "striped" # baz
`,
		},
		{
			`
a="apple"# foo
bungle=(# woo parens
"bonce"
)# baz
zebra="striped"# baz
`,
			`
a = "apple" # foo
bungle = (  # woo parens
  "bonce"
)                 # baz
zebra = "striped" # baz
`,
		},
		{
			`
𝒜 = 1 # foo
bungle = "🇬🇧" # baz
zebra = "striped" # baz
`,
			`
𝒜      = 1         # foo
bungle = "🇬🇧"       # baz
zebra  = "striped" # baz
`,
		},
		{
			`
foo {
# ...
}
`,
			`
foo {
  # ...
}
`,
		},
		{
			`
foo = {
# ...
}
`,
			`
foo = {
  # ...
}
`,
		},
		{
			`
foo = [
# ...
]
`,
			`
foo = [
  # ...
]
`,
		},
		{
			`
foo = [{
# ...
}]
`,
			`
foo = [{
  # ...
}]
`,
		},
		{
			`
foo {
bar {
# ...
}
}
`,
			`
foo {
  bar {
    # ...
  }
}
`,
		},
		{
			`
foo {
bar = {
# ...
}
}
`,
			`
foo {
  bar = {
    # ...
  }
}
`,
		},
		{
			`
foo {
bar = [
# ...
]
}
`,
			`
foo {
  bar = [
    # ...
  ]
}
`,
		},
		{
			`
foo {
bar = <<EOT
Foo bar baz
EOT
}
`,
			`
foo {
  bar = <<EOT
Foo bar baz
EOT
}
`,
		},
		{
			`
foo {
bar = <<-EOT
Foo bar baz
EOT
}
`,
			`
foo {
  bar = <<-EOT
Foo bar baz
EOT
}
`,
		},
		{
			`
foo {
bar = <<-EOT
  Foo bar baz
EOT
}
`,
			`
foo {
  bar = <<-EOT
  Foo bar baz
EOT
}
`,
		},
		{
			`
foo {
bar = <<-EOT
  blahblahblah = x
EOT
}
`,
			`
foo {
  bar = <<-EOT
  blahblahblah = x
EOT
}
`,
		},
		{
			`
foo {
bar = <<-EOT
  ${{ blahblahblah = x }}
EOT
}
`,
			`
foo {
  bar = <<-EOT
  ${ { blahblahblah = x } }
EOT
}
`,
		},
		{
			`
foo {
  bar = <<-EOT
  ${a}${b}${ c } ${d}
EOT
}
`,
			`
foo {
  bar = <<-EOT
  ${a}${b}${c} ${d}
EOT
}
`,
		},
		{
			`
foo {
bar = <<EOT
Foo bar baz
EOT
}

baz {
default="string"
}
`,
			`
foo {
  bar = <<EOT
Foo bar baz
EOT
}

baz {
  default = "string"
}
`,
		},
		{
			`
foo {
bar = <<EOT
Foo bar baz
EOT
baz = <<EOT
Foo bar baz
EOT
}

bar {
foo = "bar"
}
`,
			`
foo {
  bar = <<EOT
Foo bar baz
EOT
  baz = <<EOT
Foo bar baz
EOT
}

bar {
  foo = "bar"
}
`,
		},
		{
			`
module "foo" {
foo = <<EOF
5
EOF
}

module "x" {
a = "b"
abcde = "456"
}`,
			`
module "foo" {
  foo = <<EOF
5
EOF
}

module "x" {
  a     = "b"
  abcde = "456"
}`,
		},
		{
			`attr = provider::framework::example()`,
			`attr = provider::framework::example()`,
		},
		{
			`attr = provider :: framework :: example()`,
			`attr = provider::framework::example()`,
		},
		{
			`attr = provider ::framework:: example()`,
			`attr = provider::framework::example()`,
		},
		{
			// This is invalid syntax so formatting it with spaces
			// does not have any meaning other than to make the fact more visible
			`attr = provider::+example()`,
			`attr = provider:: + example()`,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%02d", i), func(t *testing.T) {
			tokens := lexConfigTokens([]byte(test.input))
			format(tokens)
			t.Logf("tokens %s\n", spew.Sdump(tokens))
			got := string(tokens.Bytes())

			if got != test.want {
				t.Errorf("wrong result\ninput:\n%s\ngot:\n%s\nwant:\n%s", test.input, got, test.want)
			}
		})
	}

}

func TestLinesForFormat(t *testing.T) {
	tests := []struct {
		tokens Tokens
		want   []formatLine
	}{
		{
			Tokens{
				&Token{Type: hclsyntax.TokenEOF},
			},
			[]formatLine{
				{
					lead: Tokens{},
				},
			},
		},
		{
			Tokens{
				&Token{Type: hclsyntax.TokenIdent},
				&Token{Type: hclsyntax.TokenEOF},
			},
			[]formatLine{
				{
					lead: Tokens{
						&Token{Type: hclsyntax.TokenIdent},
					},
				},
			},
		},
		{
			Tokens{
				&Token{Type: hclsyntax.TokenIdent},
				&Token{Type: hclsyntax.TokenNewline},
				&Token{Type: hclsyntax.TokenNumberLit},
				&Token{Type: hclsyntax.TokenEOF},
			},
			[]formatLine{
				{
					lead: Tokens{
						&Token{Type: hclsyntax.TokenIdent},
						&Token{Type: hclsyntax.TokenNewline},
					},
				},
				{
					lead: Tokens{
						&Token{Type: hclsyntax.TokenNumberLit},
					},
				},
			},
		},
		{
			Tokens{
				&Token{Type: hclsyntax.TokenIdent},
				&Token{Type: hclsyntax.TokenComment, Bytes: []byte("#foo\n")},
				&Token{Type: hclsyntax.TokenNumberLit},
				&Token{Type: hclsyntax.TokenEOF},
			},
			[]formatLine{
				{
					lead: Tokens{
						&Token{Type: hclsyntax.TokenIdent},
					},
					comment: Tokens{
						&Token{Type: hclsyntax.TokenComment, Bytes: []byte("#foo\n")},
					},
				},
				{
					lead: Tokens{
						&Token{Type: hclsyntax.TokenNumberLit},
					},
				},
			},
		},
		{
			Tokens{
				&Token{Type: hclsyntax.TokenIdent},
				&Token{Type: hclsyntax.TokenEqual},
				&Token{Type: hclsyntax.TokenNumberLit},
				&Token{Type: hclsyntax.TokenEOF},
			},
			[]formatLine{
				{
					lead: Tokens{
						&Token{Type: hclsyntax.TokenIdent},
					},
					assign: Tokens{
						&Token{Type: hclsyntax.TokenEqual},
						&Token{Type: hclsyntax.TokenNumberLit},
					},
				},
			},
		},
		{
			Tokens{
				&Token{Type: hclsyntax.TokenIdent},
				&Token{Type: hclsyntax.TokenEqual},
				&Token{Type: hclsyntax.TokenNumberLit},
				&Token{Type: hclsyntax.TokenComment, Bytes: []byte("#foo\n")},
				&Token{Type: hclsyntax.TokenEOF},
			},
			[]formatLine{
				{
					lead: Tokens{
						&Token{Type: hclsyntax.TokenIdent},
					},
					assign: Tokens{
						&Token{Type: hclsyntax.TokenEqual},
						&Token{Type: hclsyntax.TokenNumberLit},
					},
					comment: Tokens{
						&Token{Type: hclsyntax.TokenComment, Bytes: []byte("#foo\n")},
					},
				},
				{
					lead: Tokens{},
				},
			},
		},
		{
			Tokens{
				// A comment goes into a comment cell only if it is after
				// some non-comment tokens, since whole-line comments must
				// stay flush with the indent level.
				&Token{Type: hclsyntax.TokenComment, Bytes: []byte("#foo\n")},
				&Token{Type: hclsyntax.TokenEOF},
			},
			[]formatLine{
				{
					lead: Tokens{
						&Token{Type: hclsyntax.TokenComment, Bytes: []byte("#foo\n")},
					},
				},
				{
					lead: Tokens{},
				},
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%02d", i), func(t *testing.T) {
			got := linesForFormat(test.tokens)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("wrong result\ngot:  %#v\nwant: %#v", got, test.want)
			}
		})
	}
}

```

#### x/node.go

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
	if held, ok := c.(interface{ setHolder(*node) }); ok {
		held.setHolder(n)
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
	}
	if after != nil {
		after.before = nn
	}
	if list.first == n {
		list.first = nn
	}
	if list.last == n {
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

// List returns the nodes in the list, in order.
func (ns *nodes) List() []*node {
	var ret []*node
	for n := ns.first; n != nil; n = n.after {
		ret = append(ret, n)
	}
	return ret
}

// Last returns the last node in the list, or nil if the list is empty.
func (ns *nodes) Last() *node {
	return ns.last
}

func (ns *nodes) Clear() {
	ns.first = nil
	ns.last = nil
}

func (ns *nodes) Append(c nodeContent) *node {
	n := newNode(c)
	ns.AppendNode(n)
	n.list = ns
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
	n.list = ns
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

// inTree can be embedded into a content struct that has child nodes to get
// a standard implementation of the NodeContent interface and a record of
// the node that holds it.
type inTree struct {
	// holder is the node that was last constructed to hold the content.
	// The content is part of a syntax tree while holder is in a list.
	holder   *node
	children *nodes
}

func newInTree() inTree {
	return inTree{
		children: &nodes{},
	}
}

// assertUnattached panics if the content is already part of a syntax tree.
func (it *inTree) assertUnattached() {
	if it.holder != nil && it.holder.list != nil {
		panic(fmt.Sprintf("%T is already attached to a syntax tree", it.holder.content))
	}
}

// setHolder records n as the node that holds the content. newNode calls it.
func (it *inTree) setHolder(n *node) {
	it.holder = n
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

#### x/parser.go

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

```

#### x/parser_test.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0


package hclwrite

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/hcl/v2"
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
			"a = join(\" \", [\"fedora\", \"fez\"])\n",
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
										Type: "FunctionCallExpr",
										Children: []TestTreeNode{
											{
												Type: "Tokens",
												Val:  " join(",
											},
											{
												Type: "Expression",
												Children: []TestTreeNode{
													{
														Type: "quoted",
														Val:  "\" \"",
													},
												},
											},
											{
												Type: "Tokens",
												Val:  ",",
											},
											{
												Type: "Expression",
												Children: []TestTreeNode{
													{
														Type: "TupleConsExpr",
														Children: []TestTreeNode{
															{
																Type: "Tokens",
																Val:  " [",
															},
															{
																Type: "Expression",
																Children: []TestTreeNode{
																	{
																		Type: "quoted",
																		Val:  "\"fedora\"",
																	},
																},
															},
															{
																Type: "Tokens",
																Val:  ",",
															},
															{
																Type: "Expression",
																Children: []TestTreeNode{
																	{
																		Type: "quoted",
																		Val:  " \"fez\"",
																	},
																},
															},
															{
																Type: "Tokens",
																Val:  "]",
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
			"a = f()\n",
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
										Type: "FunctionCallExpr",
										Children: []TestTreeNode{
											{
												Type: "Tokens",
												Val:  " f()",
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
			"a = [\n  1, # one\n  b,\n]\n",
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
										Type: "TupleConsExpr",
										Children: []TestTreeNode{
											{
												Type: "Tokens",
												Val:  " [\n",
											},
											{
												Type: "Expression",
												Children: []TestTreeNode{
													{
														Type: "Tokens",
														Val:  "  1",
													},
												},
											},
											{
												Type: "Tokens",
												Val:  ", # one\n",
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
																		Val:  "  b",
																	},
																},
															},
														},
													},
												},
											},
											{
												Type: "Tokens",
												Val:  ",\n]",
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
			"a = [{ k = f(x) }]\n",
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
										Type: "TupleConsExpr",
										Children: []TestTreeNode{
											{
												Type: "Tokens",
												Val:  " [",
											},
											{
												Type: "Expression",
												Children: []TestTreeNode{
													{
														Type: "ObjectConsExpr",
														Children: []TestTreeNode{
															{
																Type: "Tokens",
																Val:  "{",
															},
															{
																Type: "ObjectConsItem",
																Children: []TestTreeNode{
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
																										Val:  " k",
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
																						Type: "FunctionCallExpr",
																						Children: []TestTreeNode{
																							{
																								Type: "Tokens",
																								Val:  " f(",
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
																														Val:  "x",
																													},
																												},
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
												Type: "Tokens",
												Val:  "]",
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

```

#### x/partition.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// partitionableTokens is a sequence of tokens in two correlated forms: the
// native tokens from hclsyntax, which carry the absolute source positions
// that the parser partitions by, and the equivalent writer tokens, which end
// up in the syntax tree. Both slices have the same length and indices.
//
type partitionableTokens struct {
	// nativeTokens carry the source positions to partition by.
	//
	nativeTokens hclsyntax.Tokens
	// writerTokens are the tokens that each partition yields.
	//
	writerTokens Tokens
}

// newPartitionableTokens pairs the given native tokens with writer tokens
// converted from them.
//
func newPartitionableTokens(nativeTokens hclsyntax.Tokens) partitionableTokens {
	return partitionableTokens{
		nativeTokens: nativeTokens,
		writerTokens: writerTokens(nativeTokens),
	}
}

func (it partitionableTokens) Partition(rng hcl.Range) (before, within, after partitionableTokens) {
	start, end := partitionTokens(it.nativeTokens, rng)
	before = it.Slice(0, start)
	within = it.Slice(start, end)
	after = it.Slice(end, len(it.nativeTokens))
	return
}

func (it partitionableTokens) PartitionType(ty hclsyntax.TokenType) (before, within, after partitionableTokens) {
	for i, t := range it.writerTokens {
		if t.Type == ty {
			return it.Slice(0, i), it.Slice(i, i+1), it.Slice(i+1, len(it.nativeTokens))
		}
	}
	panic(fmt.Sprintf("didn't find any token of type %s", ty))
}

func (it partitionableTokens) PartitionTypeOk(ty hclsyntax.TokenType) (before, within, after partitionableTokens, ok bool) {
	for i, t := range it.writerTokens {
		if t.Type == ty {
			return it.Slice(0, i), it.Slice(i, i+1), it.Slice(i+1, len(it.nativeTokens)), true
		}
	}

	return partitionableTokens{}, partitionableTokens{}, partitionableTokens{}, false
}

func (it partitionableTokens) PartitionTypeSingle(ty hclsyntax.TokenType) (before partitionableTokens, found *Token, after partitionableTokens) {
	before, within, after := it.PartitionType(ty)
	if within.Len() != 1 {
		panic("PartitionType found more than one token")
	}
	return before, within.Tokens()[0], after
}

// PartitionIncludeComments is like Partition except the returned "within"
// range includes any lead and line comments associated with the range.
func (it partitionableTokens) PartitionIncludingComments(rng hcl.Range) (before, within, after partitionableTokens) {
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
func (it partitionableTokens) PartitionBlockItem(rng hcl.Range) (before, leadComments, within, lineComments, newline, after partitionableTokens) {
	before, within, after = it.Partition(rng)
	before, leadComments = before.PartitionLeadComments()
	lineComments, newline, after = after.PartitionLineEndTokens()
	return
}

func (it partitionableTokens) PartitionLeadComments() (before, within partitionableTokens) {
	start := partitionLeadCommentTokens(it.nativeTokens)
	before = it.Slice(0, start)
	within = it.Slice(start, len(it.nativeTokens))
	return
}

func (it partitionableTokens) PartitionLineEndTokens() (comments, newline, after partitionableTokens) {
	afterComments, afterNewline := partitionLineEndTokens(it.nativeTokens)
	comments = it.Slice(0, afterComments)
	newline = it.Slice(afterComments, afterNewline)
	after = it.Slice(afterNewline, len(it.nativeTokens))
	return
}

func (it partitionableTokens) Slice(start, end int) partitionableTokens {
	// When we slice, we create a new slice with no additional capacity because
	// we expect that these slices will be mutated in order to insert
	// new code into the AST, and we want to ensure that a new underlying
	// array gets allocated in that case, rather than writing into some
	// following slice and corrupting it.
	return partitionableTokens{
		nativeTokens: it.nativeTokens[start:end:end],
		writerTokens: it.writerTokens[start:end:end],
	}
}

func (it partitionableTokens) Len() int {
	return len(it.nativeTokens)
}

func (it partitionableTokens) Tokens() Tokens {
	return it.writerTokens
}

func (it partitionableTokens) Types() []hclsyntax.TokenType {
	ret := make([]hclsyntax.TokenType, len(it.nativeTokens))
	for i, tok := range it.nativeTokens {
		ret[i] = tok.Type
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

```

#### x/partition_test.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

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

```

#### x/public.go

```go
// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"bytes"

	"github.com/hashicorp/hcl/v2"
)

// NewFile creates a new file object that is empty and ready to have constructs
// added t it.
func NewFile() *File {
	body := &Body{
		inTree: newInTree(),
		items:  newNodeSet(),
	}
	file := &File{
		inTree: newInTree(),
	}
	file.body = file.children.Append(body)
	return file
}

// ParseConfig interprets the given source bytes into a *hclwrite.File. The
// resulting AST can be used to perform surgical edits on the source code
// before turning it back into bytes again.
func ParseConfig(src []byte, filename string, start hcl.Pos) (*File, hcl.Diagnostics) {
	return parse(src, filename, start)
}

// Format takes source code and performs simple whitespace changes to transform
// it to a canonical layout style.
//
// Format skips constructing an AST and works directly with tokens, so it
// is less expensive than formatting via the AST for situations where no other
// changes will be made. It also ignores syntax errors and can thus be applied
// to partial source code, although the result in that case may not be
// desirable.
func Format(src []byte) []byte {
	tokens := lexConfigTokens(src)
	format(tokens)
	buf := &bytes.Buffer{}
	//nolint:errcheck // FIXME: Propogate errors upward.
	tokens.WriteTo(buf)
	return buf.Bytes()
}

```

#### x/round_trip_test.go

```go
// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	"github.com/zclconf/go-cty/cty/function/stdlib"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

func TestRoundTripVerbatim(t *testing.T) {
	tests := []string{
		``,
		`foo = 1
`,
		`
foobar = 1
baz    = 1
`,
		`
# this file is awesome

# tossed salads and scrambled eggs
foobar = 1
baz    = 1

block {
  a = "a"
  b = "b"
  c = "c"
  d = "d"

  subblock {
  }

  subblock {
    e = "e"
  }
}

# and they all lived happily ever after
`,
		`hats = join(" ", ["fedora", "fez"])
`,
		`hats = concat(
  # the formal ones
  ["top", "bowler"],
  [for h in var.hats : upper(h)], // all caps
  local.more...
)
nested = { a = [f(x), { b = g([y, z]) }], c = provider::p::f() }
empty  = [f(), [], {}]
`,
	}

	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			src := []byte(test)
			file, diags := parse(src, "", hcl.Pos{Line: 1, Column: 1})
			if len(diags) != 0 {
				for _, diag := range diags {
					t.Logf(" - %s", diag.Error())
				}
				t.Fatalf("unexpected diagnostics")
			}

			wr := &bytes.Buffer{}
			n, err := file.WriteTo(wr)
			if n != int64(len(test)) {
				t.Errorf("wrong number of bytes %d; want %d", n, len(test))
			}
			if err != nil {
				t.Fatalf("error from WriteTo")
			}

			result := wr.Bytes()
			if !bytes.Equal(result, src) {
				t.Errorf("wrong result\ndiff:\n%s", cmp.Diff(string(src), string(result)))
			}
		})
	}
}

func TestRoundTripFormat(t *testing.T) {
	// The goal of this test is to verify that the formatter doesn't change
	// the semantics of any expressions when it adds and removes whitespace.
	// String templates are the primary area of concern here, but we also
	// test some other things for completeness sake.
	//
	// The tests here must define zero or more attributes, which will be
	// extract with JustAttributes and evaluated both before and after
	// formatting.

	tests := []string{
		"",
		"\n\n\n",
		"a=1\n",
		"a=\"hello\"\n",
		"a=\"${hello} world\"\n",
		"a=upper(\"hello\")\n",
		"a=upper(hello)\n",
		"a=[1,2,3,4,five]\n",
		"a={greeting=hello}\n",
		"a={\ngreeting=hello\n}\n",
		"a={\ngreeting=hello}\n",
		"a={greeting=hello\n}\n",
		"a={greeting=hello,number=five,sarcastic=\"${upper(hello)}\"\n}\n",
		"a={\ngreeting=hello\nnumber=five\nsarcastic=\"${upper(hello)}\"\n}\n",
		"a=<<EOT\nhello\nEOT\n\n",
		"a=[<<EOT\nhello\nEOT\n]\n",
		"a=[\n<<EOT\nhello\nEOT\n]\n",
		"a=[\n]\n",
		"a=1\nb=2\nc=3\n",
		"a=\"${\n5\n}\"\n",
	}

	ctx := &hcl.EvalContext{
		Variables: map[string]cty.Value{
			"hello": cty.StringVal("hello"),
			"five":  cty.NumberIntVal(5),
		},
		Functions: map[string]function.Function{
			"upper": stdlib.UpperFunc,
		},
	}

	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			attrsAsObj := func(src []byte, phase string) cty.Value {
				t.Logf("source %s:\n%s", phase, src)
				f, diags := hclsyntax.ParseConfig(src, "", hcl.Pos{Line: 1, Column: 1})
				if len(diags) != 0 {
					for _, diag := range diags {
						t.Logf(" - %s", diag.Error())
					}
					t.Fatalf("unexpected diagnostics in parse %s", phase)
				}

				attrs, diags := f.Body.JustAttributes()
				if len(diags) != 0 {
					for _, diag := range diags {
						t.Logf(" - %s", diag.Error())
					}
					t.Fatalf("unexpected diagnostics in JustAttributes %s", phase)
				}

				vals := map[string]cty.Value{}
				for k, attr := range attrs {
					val, diags := attr.Expr.Value(ctx)
					if len(diags) != 0 {
						for _, diag := range diags {
							t.Logf(" - %s", diag.Error())
						}
						t.Fatalf("unexpected diagnostics evaluating %s", phase)
					}
					vals[k] = val
				}
				return cty.ObjectVal(vals)
			}

			src := []byte(test)
			before := attrsAsObj(src, "before")

			formatted := Format(src)
			after := attrsAsObj(formatted, "after")

			if !after.RawEquals(before) {
				t.Errorf("mismatching after format\nbefore: %#v\nafter:  %#v", before, after)
			}
		})
	}
}

// TestRoundTripSafeConcurrent concurrently generates a new file. When run with
// the race detector it will fail if a data race is detected.
func TestRoundTripSafeConcurrent(t *testing.T) {
	for i := 0; i < 2; i++ {
		go func() {
			f := NewEmptyFile()
			b := f.Body()
			b.SetAttributeValue("foo", cty.StringVal("bar"))
		}()
	}
}

```

#### x/tokens.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"bytes"
	"io"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/internal/unicodeutil"
)

// Token is a single sequence of bytes annotated with a type. It is similar
// in purpose to hclsyntax.Token, but discards the source position information
// since that is not useful in code generation.
type Token struct {
	Type  hclsyntax.TokenType
	Bytes []byte

	// We record the number of spaces before each token so that we can
	// reproduce the exact layout of the original file when we're making
	// surgical changes in-place. When _new_ code is created it will always
	// be in the canonical style, but we preserve layout of existing code.
	SpacesBefore int
}

// asHCLSyntax returns the receiver expressed as an incomplete hclsyntax.Token.
// A complete token is not possible since we don't have source location
// information here, and so this method is unexported so we can be sure it will
// only be used for internal purposes where we know the range isn't important.
//
// This is primarily intended to allow us to re-use certain functionality from
// hclsyntax rather than re-implementing it against our own token type here.
func (t *Token) asHCLSyntax() hclsyntax.Token {
	return hclsyntax.Token{
		Type:  t.Type,
		Bytes: t.Bytes,
		Range: hcl.Range{
			Filename: "<invalid>",
		},
	}
}

// Tokens is a flat list of tokens.
type Tokens []*Token

func (ts Tokens) Bytes() []byte {
	buf := &bytes.Buffer{}
	//nolint:errcheck // FIXME: Propogate errors upward.
	ts.WriteTo(buf)
	return buf.Bytes()
}

func (ts Tokens) testValue() string {
	return string(ts.Bytes())
}

// Columns returns the number of columns (grapheme clusters) the token sequence
// occupies. The result is not meaningful if there are newline or single-line
// comment tokens in the sequence.
func (ts Tokens) Columns() int {
	ret := 0
	for _, token := range ts {
		ret += token.SpacesBefore // spaces are always worth one column each
		ct, _ := unicodeutil.GraphemeCount(token.Bytes)
		ret += ct
	}
	return ret
}

// WriteTo takes an io.Writer and writes the bytes for each token to it,
// along with the spacing that separates each token. In other words, this
// allows serializing the tokens to a file or other such byte stream.
func (ts Tokens) WriteTo(wr io.Writer) (int64, error) {
	// We know we're going to be writing a lot of small chunks of repeated
	// space characters, so we'll prepare a buffer of these that we can
	// easily pass to wr.Write without any further allocation.
	spaces := make([]byte, 40)
	for i := range spaces {
		spaces[i] = ' '
	}

	var n int64
	var err error
	for _, token := range ts {
		if err != nil {
			return n, err
		}

		for spacesBefore := token.SpacesBefore; spacesBefore > 0; spacesBefore -= len(spaces) {
			thisChunk := spacesBefore
			if thisChunk > len(spaces) {
				thisChunk = len(spaces)
			}
			var thisN int
			thisN, err = wr.Write(spaces[:thisChunk])
			n += int64(thisN)
			if err != nil {
				return n, err
			}
		}

		var thisN int
		thisN, err = wr.Write(token.Bytes)
		n += int64(thisN)
	}

	return n, err
}

func (ts Tokens) walkChildNodes(w internalWalkFunc) {
	// Unstructured tokens have no child nodes
}

func (ts Tokens) BuildTokens(to Tokens) Tokens {
	return append(to, ts...)
}

// ObjectAttrTokens represents the raw tokens for the name and value of
// one attribute in an object constructor expression.
//
// This is defined primarily for use with function TokensForObject. See
// that function's documentation for more information.
type ObjectAttrTokens struct {
	Name  Tokens
	Value Tokens
}

// writerTokens takes a sequence of tokens as produced by the main hclsyntax
// package and transforms it into an equivalent sequence of tokens using
// this package's own token model.
//
// The resulting list contains the same number of tokens and uses the same
// indices as the input, allowing the two sets of tokens to be correlated
// by index.
//
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

// lexConfigTokens uses the hclsyntax scanner to get a token stream and then
// rewrites it into this package's token model.
//
// Any errors produced during scanning are ignored, so the results of this
// function should be used with care.
//
func lexConfigTokens(src []byte) Tokens {
	mainTokens, _ := hclsyntax.LexConfig(src, "", hcl.Pos{Byte: 0, Line: 1, Column: 1})
	return writerTokens(mainTokens)
}

func newIdentToken(name string) *Token {
	return &Token{
		Type:  hclsyntax.TokenIdent,
		Bytes: []byte(name),
	}
}

```

#### x/tokens_test.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/hcl/v2/hclsyntax"
)

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
			got := lexConfigTokens([]byte(test.input))

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

#### y/ast_expression.go

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
// expression.
func (e *Expression) Variables() []*Traversal {
	nodes := e.variableNodes()
	ret := make([]*Traversal, len(nodes))
	for i, node := range nodes {
		ret[i] = node.content.(*Traversal)
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
func (e *Expression) RenameVariablePrefix(search, replacement []string) {
	if len(search) != len(replacement) {
		panic(fmt.Sprintf("search and replacement length mismatch (%d and %d)", len(search), len(replacement)))
	}
Traversals:
	for _, node := range e.variableNodes() {
		traversal := node.content.(*Traversal)
		if len(traversal.steps) < len(search) {
			// If it's shorter then it can't have our prefix
			continue
		}

		stepNodes := traversal.steps.List()
		for i, name := range search {
			step, isName := stepNodes[i].content.(*TraverseName)
			if !isName {
				continue Traversals // only name nodes can match
			}
			foundNameBytes := step.name.content.(*identifier).token.Bytes
			if len(foundNameBytes) != len(name) {
				continue Traversals
			}
			if string(foundNameBytes) != name {
				continue Traversals
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
}

func (e *Expression) AsObjectConsExpr() *ObjectConsExpr {
	if e.wrapped == nil {
		return nil
	}

	if object, ok := e.wrapped.content.(*ObjectConsExpr); ok {
		return object
	}

	return nil
}

// AsFunctionCallExpr returns the wrapped function call expression, or nil if
// this expression does not parse as a function call expression.
func (e *Expression) AsFunctionCallExpr() *FunctionCallExpr {
	if e.wrapped == nil {
		return nil
	}

	if call, ok := e.wrapped.content.(*FunctionCallExpr); ok {
		return call
	}

	return nil
}

// AsTupleConsExpr returns the wrapped tuple-construct expression, or nil if
// this expression does not parse as a tuple-construct expression.
func (e *Expression) AsTupleConsExpr() *TupleConsExpr {
	if e.wrapped == nil {
		return nil
	}

	if tuple, ok := e.wrapped.content.(*TupleConsExpr); ok {
		return tuple
	}

	return nil
}

// variableNodes returns the nodes of the absolute traversals in the
// expression, in source order, including those inside a wrapped
// object-construct, function call or tuple-construct expression.
//
func (e *Expression) variableNodes() []*node {
	nodes := e.absTraversals.List()
	if object := e.AsObjectConsExpr(); object != nil {
		nodes = append(nodes, object.variableNodes()...)
	}
	if call := e.AsFunctionCallExpr(); call != nil {
		nodes = append(nodes, call.variableNodes()...)
	}
	if tuple := e.AsTupleConsExpr(); tuple != nil {
		nodes = append(nodes, tuple.variableNodes()...)
	}
	return nodes
}

func (e *Expression) asQuoted() *quoted {
	if e.wrapped == nil {
		return nil
	}

	if quoted, ok := e.wrapped.content.(*quoted); ok {
		return quoted
	}

	return nil
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

#### y/ast_expression_test.go

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

func TestExpressionVariablesObjectCons(t *testing.T) {
	tests := []struct {
		src     string
		vars    []string
		renamed string
	}{
		{
			src:     "a = { k = old.x }\n",
			vars:    []string{"old.x"},
			renamed: "a = { k = new.x }\n",
		},
		{
			src:     "a = { k = { j = old.x } }\n",
			vars:    []string{"old.x"},
			renamed: "a = { k = { j = new.x } }\n",
		},
		{
			src:     "a = { k = [old.x] }\n",
			vars:    []string{"old.x"},
			renamed: "a = { k = [new.x] }\n",
		},
		{
			src:     "a = { k = f(old.x) }\n",
			vars:    []string{"old.x"},
			renamed: "a = { k = f(new.x) }\n",
		},
		{
			src:     "a = { k = \"${old.x}\" }\n",
			vars:    []string{"old.x"},
			renamed: "a = { k = \"${new.x}\" }\n",
		},
		{
			// A naked identifier key is a literal name, not a variable.
			src:     "a = { old = 1 }\n",
			vars:    nil,
			renamed: "a = { old = 1 }\n",
		},
		{
			src:     "a = { old = old.x }\n",
			vars:    []string{"old.x"},
			renamed: "a = { old = new.x }\n",
		},
		{
			// A parenthesized key is an expression.
			src:     "a = { (old.x) = 1 }\n",
			vars:    []string{"old.x"},
			renamed: "a = { (new.x) = 1 }\n",
		},
		{
			src:     "a = { \"k\" = old.x }\n",
			vars:    []string{"old.x"},
			renamed: "a = { \"k\" = new.x }\n",
		},
		{
			src:     "a = { (old.a) = old.b, k = { j = old.c }, l = other.d }\n",
			vars:    []string{"old.a", "old.b", "old.c", "other.d"},
			renamed: "a = { (new.a) = new.b, k = { j = new.c }, l = other.d }\n",
		},
		{
			src:     "a = {}\n",
			vars:    nil,
			renamed: "a = {}\n",
		},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			expr := f.Body().GetAttribute("a").Expr()

			var got []string
			for _, traversal := range expr.Variables() {
				got = append(got, strings.TrimSpace(string(traversal.BuildTokens(nil).Bytes())))
			}
			if diff := cmp.Diff(test.vars, got); diff != "" {
				t.Errorf("wrong variables\n%s", diff)
			}

			nativeFile, diags := hclsyntax.ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected hclsyntax diagnostics: %s", diags.Error())
			}
			nativeAttrs, _ := nativeFile.Body.JustAttributes()
			if want := len(nativeAttrs["a"].Expr.Variables()); len(got) != want {
				t.Errorf("got %d variables, but hclsyntax reports %d", len(got), want)
			}

			expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
			if diff := cmp.Diff(test.renamed, string(f.Bytes())); diff != "" {
				t.Errorf("wrong result after rename\n%s", diff)
			}
		})
	}
}

func TestExpressionVariablesFunctionCallAndTupleCons(t *testing.T) {
	tests := []struct {
		src     string
		vars    []string
		renamed string
	}{
		{
			src:     "a = f(old.x, old.y)\n",
			vars:    []string{"old.x", "old.y"},
			renamed: "a = f(new.x, new.y)\n",
		},
		{
			src:     "a = [old.x, other.y]\n",
			vars:    []string{"old.x", "other.y"},
			renamed: "a = [new.x, other.y]\n",
		},
		{
			src:     "a = f(old.a, [old.b, g(old.c)], { k = old.d }, old.e...)\n",
			vars:    []string{"old.a", "old.b", "old.c", "old.d", "old.e"},
			renamed: "a = f(new.a, [new.b, g(new.c)], { k = new.d }, new.e...)\n",
		},
		{
			src:     "a = { k = f(old.x, [old.y]), l = [g(old.z)] }\n",
			vars:    []string{"old.x", "old.y", "old.z"},
			renamed: "a = { k = f(new.x, [new.y]), l = [g(new.z)] }\n",
		},
		{
			src:     "a = f(\"${old.x}\", old.y + 1)\n",
			vars:    []string{"old.x", "old.y"},
			renamed: "a = f(\"${new.x}\", new.y + 1)\n",
		},
		{
			src:     "a = f()\n",
			vars:    nil,
			renamed: "a = f()\n",
		},
		{
			src:     "a = []\n",
			vars:    nil,
			renamed: "a = []\n",
		},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			expr := f.Body().GetAttribute("a").Expr()

			var got []string
			for _, traversal := range expr.Variables() {
				got = append(got, strings.TrimSpace(string(traversal.BuildTokens(nil).Bytes())))
			}
			if diff := cmp.Diff(test.vars, got); diff != "" {
				t.Errorf("wrong variables\n%s", diff)
			}

			expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
			if diff := cmp.Diff(test.renamed, string(f.Bytes())); diff != "" {
				t.Errorf("wrong result after rename\n%s", diff)
			}
		})
	}
}

func TestExpressionAsFunctionCallExprAndTupleConsExpr(t *testing.T) {
	tests := []struct {
		src       string
		wantCall  bool
		wantTuple bool
	}{
		{src: "a = f(x)\n", wantCall: true},
		{src: "a = [x]\n", wantTuple: true},
		{src: "a = f(x) + 1\n"},
		{src: "a = [for x in y : x]\n"},
		{src: "a = { k = [x] }\n"},
		{src: "a = x\n"},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			expr := f.Body().GetAttribute("a").Expr()
			if got := expr.AsFunctionCallExpr() != nil; got != test.wantCall {
				t.Errorf("AsFunctionCallExpr returned a call: %t, want %t", got, test.wantCall)
			}
			if got := expr.AsTupleConsExpr() != nil; got != test.wantTuple {
				t.Errorf("AsTupleConsExpr returned a tuple: %t, want %t", got, test.wantTuple)
			}
		})
	}
}

```

#### y/ast_function_call_expr.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// FunctionCallExpr represents a function call expression such as
// join(" ", list). Get one from an expression with
// Expression.AsFunctionCallExpr.
//
// Each argument is an Expression of its own. The function name, the
// parentheses, the commas and a final "..." stay as tokens of the call.
type FunctionCallExpr struct {
	inTree

	args nodeSet
}

func newFunctionCallExpr() *FunctionCallExpr {
	return &FunctionCallExpr{
		inTree: newInTree(),
		args:   newNodeSet(),
	}
}

// Name returns the name of the called function. A namespaced name keeps its
// namespaces, as in "provider::example::upper".
func (call *FunctionCallExpr) Name() string {
	// The parser puts the name and the opening parenthesis in the first run
	// of tokens.
	var name []byte
	done := false
	call.walkChildNodes(func(n *node) {
		tokens, ok := n.content.(Tokens)
		if done || !ok {
			done = true
			return
		}
		for _, tok := range tokens {
			switch tok.Type {
			case hclsyntax.TokenOParen:
				done = true
				return
			case hclsyntax.TokenIdent, hclsyntax.TokenDoubleColon:
				name = append(name, tok.Bytes...)
			}
		}
	})
	return string(name)
}

// variableNodes returns the traversal nodes in the arguments, in source
// order.
//
func (call *FunctionCallExpr) variableNodes() []*node {
	var ret []*node
	for _, n := range call.args.List() {
		ret = append(ret, n.content.(*Expression).variableNodes()...)
	}
	return ret
}

```

#### y/ast_function_call_expr_test.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
)

func TestFunctionCallExprName(t *testing.T) {
	tests := []struct {
		src  string
		want string
	}{
		{src: "a = join(\" \", [\"fedora\", \"fez\"])\n", want: "join"},
		{src: "a = timestamp()\n", want: "timestamp"},
		{src: "a = provider::example::upper(\"x\")\n", want: "provider::example::upper"},
		{src: "a = f(g(x))\n", want: "f"},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			call := f.Body().GetAttribute("a").Expr().AsFunctionCallExpr()
			if call == nil {
				t.Fatal("not parsed as a function call expression")
			}
			if got := call.Name(); got != test.want {
				t.Errorf("wrong name\ngot:  %s\nwant: %s", got, test.want)
			}
		})
	}
}

```

#### y/ast_object_cons_expr.go

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

// ObjectConsExpr is an object-construct expression, such as
// { hat = "derby" }. Get one from an expression with
// Expression.AsObjectConsExpr.
//
// Its methods find an item by its key only when the key is an identifier or
// a quoted string literal. The two forms are equivalent: the key "hat" finds
// both hat = ... and "hat" = .... A key that is taken from a variable or from
// another expression, such as (hat), is never found.
type ObjectConsExpr struct {
	inTree

	items nodeSet
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

// ItemFor returns the first item whose key is the given string, or nil if
// there is none.
//
// A quoted key is compared by its value after its escape sequences are
// decoded, so ItemFor("${hat}") finds the key "$${hat}".
func (o *ObjectConsExpr) ItemFor(key string) *ObjectConsItem {
	for _, item := range o.Items() {
		if name, ok := item.KeyExpr().staticName(); ok && name == key {
			return item
		}
	}
	return nil
}

// ValueExprFor returns the value of the item that ItemFor finds for the
// given key, or nil if there is none.
func (o *ObjectConsExpr) ValueExprFor(key string) *ObjectConsValue {
	item := o.ItemFor(key)
	if item == nil {
		return nil
	}
	return item.ValueExpr()
}

// SetItem either replaces the value of the item that ItemFor finds for the
// given key, leaving its key as written, or adds a new item at the end of
// the object.
//
// A new key is written as an identifier when it is a valid identifier other
// than "for", and otherwise as a quoted string literal that evaluates to the
// given key.
//
// SetItem panics if the expression is already part of a syntax tree, such as
// one returned by Attribute.Expr or ObjectConsValue.Expr. Two items sharing
// one expression would change together.
//
// The return values are the key and the value of the item that was either
// modified in place or created.
func (o *ObjectConsExpr) SetItem(key string, expr *Expression) (*ObjectConsKeyExpr, *ObjectConsValue) {
	expr.assertUnattached()

	item := o.ItemFor(key)
	if item != nil {
		item.setValueExpr(expr)
	} else {
		item = newObjectConsItem()
		item.init(key, expr)
		o.appendItem(item)
	}
	return item.KeyExpr(), item.ValueExpr()
}

// SetItemRaw is like SetItem, but uses the given tokens verbatim as the
// value expression.
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

// SetItemTraversal is like SetItem, but sets the value to the given
// traversal, as NewExpressionAbsTraversal does. The traversal must be
// absolute. To set a literal value, use SetItemValue.
func (o *ObjectConsExpr) SetItemTraversal(key string, traversal hcl.Traversal) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionAbsTraversal(traversal))
}

// variableNodes returns the traversal nodes in the keys and values of the
// items, in source order. A key that is a literal name contributes none.
//
func (o *ObjectConsExpr) variableNodes() []*node {
	var ret []*node
	for _, n := range o.items.List() {
		item := n.content.(*ObjectConsItem)
		key := item.key.content.(*ObjectConsKeyExpr)
		if !key.literalName {
			ret = append(ret, key.wrapped.content.(*Expression).variableNodes()...)
		}
		value := item.value.content.(*ObjectConsValue)
		ret = append(ret, value.expr.content.(*Expression).variableNodes()...)
	}
	return ret
}

// appendItem adds the item after the last one, laid out as the object
// already is: on the same line after a comma when the object is written on
// one line, and otherwise on a line of its own before the closing brace.
func (o *ObjectConsExpr) appendItem(item *ObjectConsItem) {
	// The parser leaves everything after the last item, up to and including
	// the closing brace, in one run of tokens at the end of the children.
	var closing *node
	o.walkChildNodes(func(n *node) { closing = n })
	closingTokens, ok := closing.content.(Tokens)
	if !ok || len(closingTokens) == 0 || closingTokens[len(closingTokens)-1].Type != hclsyntax.TokenCBrace {
		panic(fmt.Sprintf("object-construct expression does not end with a closing brace, but with %T", closing.content))
	}
	// Each use gets its own token, since formatting changes tokens in place.
	newline := func() *Token {
		return &Token{
			Type:  hclsyntax.TokenNewline,
			Bytes: []byte{'\n'},
		}
	}

	multiline := o.isMultiline()
	kept, rest := o.splitClosing(closingTokens)
	// Copy kept, since appending to it would overwrite the start of rest.
	separator := append(Tokens(nil), kept...)
	var afterItem Tokens
	switch {
	case multiline:
		if len(kept) == 0 || !o.tokenEndsLine(kept[len(kept)-1]) {
			separator = append(separator, newline())
		}
		afterItem = Tokens{newline()}
	case len(o.items) > 0 && len(kept) == 0:
		separator = append(separator, &Token{
			Type:  hclsyntax.TokenComma,
			Bytes: []byte{','},
		})
	}
	afterItem = append(afterItem, rest...)

	closing.Detach()
	o.children.AppendUnstructuredTokens(separator)
	o.items.Add(o.children.Append(item))
	o.children.AppendUnstructuredTokens(afterItem)
}

// isMultiline reports whether the object is written on several lines. Line
// breaks inside the keys and values of its items do not count.
func (o *ObjectConsExpr) isMultiline() bool {
	multiline := false
	checkTokens := func(n *node) {
		if tokens, ok := n.content.(Tokens); ok {
			for _, tok := range tokens {
				multiline = multiline || o.tokenEndsLine(tok)
			}
		}
	}
	o.walkChildNodes(func(n *node) {
		checkTokens(n)
		if item, ok := n.content.(*ObjectConsItem); ok {
			item.walkChildNodes(checkTokens)
		}
	})
	return multiline
}

// splitClosing splits the tokens after the last item into those that stay
// with that item and the rest, which end with the closing brace. The item
// keeps a comma that follows it, and the comments up to the end of its line
// when the closing brace is not on that line.
func (o *ObjectConsExpr) splitClosing(tokens Tokens) (kept, rest Tokens) {
	afterComma := 0
	if tokens[0].Type == hclsyntax.TokenComma {
		afterComma = 1
	}
	for i := afterComma; i < len(tokens); i++ {
		if o.tokenEndsLine(tokens[i]) {
			return tokens[:i+1], tokens[i+1:]
		}
		if tokens[i].Type != hclsyntax.TokenComment {
			break
		}
	}
	return tokens[:afterComma], tokens[afterComma:]
}

// tokenEndsLine reports whether the token is a newline, or a line comment,
// which includes its newline.
func (o *ObjectConsExpr) tokenEndsLine(tok *Token) bool {
	switch tok.Type {
	case hclsyntax.TokenNewline:
		return true
	case hclsyntax.TokenComment:
		return len(tok.Bytes) > 0 && tok.Bytes[len(tok.Bytes)-1] == '\n'
	}
	return false
}

// ObjectConsItem is one item of an object-construct expression: a key and a
// value, as in hat = "derby".
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

// KeyExpr returns the key of the item.
func (item *ObjectConsItem) KeyExpr() *ObjectConsKeyExpr {
	return item.key.content.(*ObjectConsKeyExpr)
}

// ValueExpr returns the value of the item.
func (item *ObjectConsItem) ValueExpr() *ObjectConsValue {
	return item.value.content.(*ObjectConsValue)
}

// init builds a new item from the key and the value expression, which must
// not be attached yet. The key is written as an identifier when it is a
// valid one, and otherwise as a quoted string literal.
func (item *ObjectConsItem) init(key string, expr *Expression) {
	// A naked "for" as the first key would start a for expression.
	literalName := hclsyntax.ValidIdentifier(key) && key != "for"
	var keyTokens Tokens
	if literalName {
		keyTokens = TokensForIdentifier(key)
	} else {
		keyTokens = TokensForValue(cty.StringVal(key))
	}
	item.key = item.children.Append(newObjectConsKeyExpr(newNode(NewExpressionRaw(keyTokens)), literalName))
	item.children.AppendUnstructuredTokens(Tokens{
		{
			Type:  hclsyntax.TokenEqual,
			Bytes: []byte{'='},
		},
	})

	item.value = item.children.Append(newObjectConsValue())
	item.setValueExpr(expr)
}

// setValueExpr makes the given expression, which must not be attached yet,
// the expression of the item's value. The expression it replaces is
// detached.
func (item *ObjectConsItem) setValueExpr(expr *Expression) {
	value := item.ValueExpr()
	if value.expr != nil {
		value.Expr().setParent(nil)
	}
	// The value has no children other than its expression.
	value.children.Clear()
	value.expr = value.children.Append(expr)
	expr.setParent(item.value)
}

// ObjectConsKeyExpr is the key of an item in an object-construct
// expression. Get one from ObjectConsItem.KeyExpr.
type ObjectConsKeyExpr struct {
	inTree

	wrapped *node

	// literalName is true when the key is a naked name, such as "k" in
	// { k = v }. As in hclsyntax, such a key is a literal string, not a
	// reference, so it has no variables.
	literalName bool
}

func newObjectConsKeyExpr(wrapped *node, literalName bool) *ObjectConsKeyExpr {
	expr := &ObjectConsKeyExpr{
		inTree:      newInTree(),
		wrapped:     wrapped,
		literalName: literalName,
	}
	expr.children.AppendNode(wrapped)

	return expr
}

// staticName returns the string that the key stands for, and true, when the
// key is an identifier or a quoted string literal. Otherwise it returns false.
//
// It reads the key from its tokens as hclsyntax does, so that a quoted key
// is compared by its value after its escape sequences are decoded.
func (k *ObjectConsKeyExpr) staticName() (string, bool) {
	tokens := k.BuildTokens(nil)
	if len(tokens) == 0 {
		return "", false
	}
	native, diags := hclsyntax.ParseExpression(tokens.Bytes(), "", hcl.InitialPos)
	if diags.HasErrors() {
		return "", false
	}

	// As in hclsyntax, a naked identifier is a literal name. A parenthesized
	// one is a ParenthesesExpr, which this does not unwrap.
	if name := hcl.ExprAsKeyword(native); name != "" {
		return name, true
	}

	template, ok := native.(*hclsyntax.TemplateExpr)
	if !ok || tokens[0].Type != hclsyntax.TokenOQuote {
		return "", false
	}
	for _, part := range template.Parts {
		if _, ok := part.(*hclsyntax.LiteralValueExpr); !ok {
			return "", false
		}
	}
	val, diags := template.Value(nil)
	if diags.HasErrors() || !val.IsKnown() || val.IsNull() || val.Type() != cty.String {
		return "", false
	}
	return val.AsString(), true
}

// ObjectConsValue is the value of an item in an object-construct
// expression. Get one from ObjectConsItem.ValueExpr.
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

```

#### y/ast_object_cons_expr_test.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/zclconf/go-cty/cty"

	"github.com/hashicorp/hcl/v2"
)

func TestObjectConsExprItemFor(t *testing.T) {
	src := `a = {
  hat         = "derby"
  "cat"       = "calico"
  "$${dog}"   = "beagle"
  "a hat"     = "fedora"
  "say \"hi\"" = "hello"
  true        = "yes"
  ""          = "empty"
  (band)      = "variable"
  hat.band    = "traversal"
  "${hat}"    = "template"
  "x${y}"     = "mixed"
}
`
	tests := []struct {
		key  string
		want string // the value's source, or "" when no item is found
	}{
		{key: "hat", want: `"derby"`},
		{key: "cat", want: `"calico"`},
		{key: "${dog}", want: `"beagle"`},
		{key: "$${dog}", want: ""},
		{key: "a hat", want: `"fedora"`},
		{key: `say "hi"`, want: `"hello"`},
		{key: "true", want: `"yes"`},
		{key: "", want: `"empty"`},
		{key: "band", want: ""},
		{key: "(band)", want: ""},
		{key: "hat.band", want: ""},
		{key: "${hat}", want: ""},
		{key: "x${y}", want: ""},
		{key: "missing", want: ""},
	}

	f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatal(diags)
	}
	object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()
	if got := len(object.Items()); got != 11 {
		t.Fatalf("got %d items, want 11", got)
	}

	for _, test := range tests {
		t.Run(test.key, func(t *testing.T) {
			item := object.ItemFor(test.key)
			value := object.ValueExprFor(test.key)
			if test.want == "" {
				if item != nil || value != nil {
					t.Fatalf("found an item for %q, want none", test.key)
				}
				return
			}
			if item == nil || value == nil {
				t.Fatalf("found no item for %q", test.key)
			}
			if value != item.ValueExpr() {
				t.Error("ValueExprFor returned a different value than ItemFor")
			}
			got := strings.TrimSpace(string(value.Expr().BuildTokens(nil).Bytes()))
			if got != test.want {
				t.Errorf("wrong value\ngot:  %s\nwant: %s", got, test.want)
			}
		})
	}
}

func TestObjectConsExprItems(t *testing.T) {
	f, diags := ParseConfig([]byte("a = { b = 1, (c) = 2, \"d\" = 3 }\n"), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatal(diags)
	}
	object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

	var got []string
	for _, item := range object.Items() {
		got = append(got, strings.TrimSpace(string(item.KeyExpr().BuildTokens(nil).Bytes())))
	}
	want := []string{"b", "(c)", `"d"`}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("wrong keys\n%s", diff)
	}
}

func TestObjectConsExprSetItemKey(t *testing.T) {
	tests := []struct {
		key  string
		want string
	}{
		{key: "hat", want: `hat`},
		{key: "hat_band-2", want: `hat_band-2`},
		{key: "a hat", want: `"a hat"`},
		{key: "hat.band", want: `"hat.band"`},
		{key: "1hat", want: `"1hat"`},
		{key: `say "hi"`, want: `"say \"hi\""`},
		{key: `a\b`, want: `"a\\b"`},
		{key: "${hat}", want: `"$${hat}"`},
		{key: "%{hat}", want: `"%%{hat}"`},
		{key: "$hat", want: `"$hat"`},
		{key: "line\nbreak", want: `"line\nbreak"`},
		{key: "", want: `""`},
		// A naked "for" as the first key would start a for expression.
		{key: "for", want: `"for"`},
	}

	for _, test := range tests {
		t.Run(test.key, func(t *testing.T) {
			f, diags := ParseConfig([]byte("a = {}\n"), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatal(diags)
			}
			object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

			key, value := object.SetItemValue(test.key, cty.StringVal("bowler"))
			if got := string(key.BuildTokens(nil).Bytes()); got != test.want {
				t.Errorf("wrong key\ngot:  %s\nwant: %s", got, test.want)
			}
			item := object.ItemFor(test.key)
			if item == nil {
				t.Fatalf("ItemFor(%q) found nothing after SetItemValue", test.key)
			}
			if item.KeyExpr() != key || item.ValueExpr() != value {
				t.Error("ItemFor found a different item than SetItemValue returned")
			}

			// The result evaluates to an object with the original key.
			out := f.Bytes()
			reparsed, diags := ParseConfig(out, "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("invalid result:\n%s\n%s", out, diags)
			}
			if reparsed.Body().GetAttribute("a").Expr().AsObjectConsExpr().ItemFor(test.key) == nil {
				t.Errorf("ItemFor(%q) found nothing after reparsing:\n%s", test.key, out)
			}
		})
	}
}

func TestObjectConsExprSetItemExisting(t *testing.T) {
	src := "a = { \"hat\" = \"derby\", cat = \"calico\" }\n"
	f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatal(diags)
	}
	object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()
	oldItem := object.ItemFor("hat")

	key, value := object.SetItemValue("hat", cty.StringVal("bowler"))
	if key != oldItem.KeyExpr() || value != oldItem.ValueExpr() {
		t.Error("SetItemValue did not return the existing item's key and value")
	}
	if got := len(object.Items()); got != 2 {
		t.Errorf("got %d items, want 2", got)
	}

	want := "a = { \"hat\" = \"bowler\", cat = \"calico\" }\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, want)
	}
}

func TestObjectConsExprSetItemKinds(t *testing.T) {
	f, diags := ParseConfig([]byte("a = {}\n"), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatal(diags)
	}
	object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

	object.SetItemRaw("raw", TokensForFunctionCall("upper", TokensForValue(cty.StringVal("x"))))
	object.SetItemValue("value", cty.NumberIntVal(5))
	_, value := object.SetItemTraversal("traversal", hcl.Traversal{
		hcl.TraverseRoot{Name: "var"},
		hcl.TraverseAttr{Name: "hat"},
	})
	object.SetItem("expr", NewExpressionLiteral(cty.True))

	want := "a = { raw = upper(\"x\"), value = 5, traversal = var.hat, expr = true }\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, want)
	}

	// The traversal is a variable of the object, so it can be renamed.
	vars := value.Expr().Variables()
	if len(vars) != 1 {
		t.Fatalf("got %d variables in the traversal value, want 1", len(vars))
	}
	f.Body().GetAttribute("a").Expr().RenameVariablePrefix([]string{"var"}, []string{"local"})
	want = "a = { raw = upper(\"x\"), value = 5, traversal = local.hat, expr = true }\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result after renaming\ngot:\n%s\nwant:\n%s", got, want)
	}
}

func TestObjectConsExprSetItemPanics(t *testing.T) {
	src := "a = { hat = \"derby\" }\nb = \"bowler\"\n"
	f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatal(diags)
	}
	object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()
	attached := NewExpressionLiteral(cty.StringVal("bowler"))
	object.SetItem("cat", attached)

	tests := map[string]*Expression{
		"attribute expression from the parser": f.Body().GetAttribute("b").Expr(),
		"item value from the parser":           object.ValueExprFor("hat").Expr(),
		"expression attached by SetItem":       attached,
	}
	for name, expr := range tests {
		t.Run(name, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Error("SetItem did not panic")
				}
			}()
			object.SetItem("dog", expr)
		})
	}

	// A replaced expression is no longer attached, so it can be reused.
	replaced := object.ValueExprFor("cat").Expr()
	object.SetItemValue("cat", cty.StringVal("calico"))
	object.SetItem("dog", replaced)

	want := "a = { hat = \"derby\", cat = \"calico\", dog = \"bowler\" }\nb = \"bowler\"\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, want)
	}
}

func TestObjectConsExprSetItemLayout(t *testing.T) {
	tests := []struct {
		name string
		src  string
		keys []string
		want string
	}{
		{
			name: "single line",
			src:  "a = { hat = \"derby\" }\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\", dog = \"bowler\" }\n",
		},
		{
			name: "empty",
			src:  "a = {}\n",
			keys: []string{"dog"},
			want: "a = { dog = \"bowler\" }\n",
		},
		{
			name: "empty, two items",
			src:  "a = {}\n",
			keys: []string{"dog", "cat"},
			want: "a = { dog = \"bowler\", cat = \"bowler\" }\n",
		},
		{
			name: "single line, trailing comma",
			src:  "a = { hat = \"derby\", }\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\", dog = \"bowler\" }\n",
		},
		{
			name: "single line, inline comment before the brace",
			src:  "a = { hat = \"derby\" /* end of a */ }\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\", dog = \"bowler\" /* end of a */ }\n",
		},
		{
			name: "single line, multi-line value",
			src:  "a = { hat = {\n  band = \"red\"\n} }\n",
			keys: []string{"dog"},
			want: "a = { hat = {\n  band = \"red\"\n}, dog = \"bowler\" }\n",
		},
		{
			name: "single line, multi-line key",
			src:  "a = { (\n  hat\n) = \"derby\" }\n",
			keys: []string{"dog"},
			want: "a = { (\n  hat\n) = \"derby\", dog = \"bowler\" }\n",
		},
		{
			name: "multi-line",
			src:  "a = {\n  hat = \"derby\"\n}\n",
			keys: []string{"dog", "cat"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n}\n",
		},
		{
			name: "multi-line, empty",
			src:  "a = {\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  dog = \"bowler\"\n}\n",
		},
		{
			name: "multi-line, brace on the last item's line",
			src:  "a = {\n  hat = \"derby\" }\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}\n",
		},
		{
			name: "multi-line, brace on the first item's line",
			src:  "a = { hat = \"derby\"\n  cat = \"calico\" }\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\"\n  cat = \"calico\"\n  dog = \"bowler\"\n}\n",
		},
		{
			name: "multi-line, comma and line comment after the last item",
			src:  "a = {\n  hat = \"derby\", # a hat\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\", # a hat\n  dog = \"bowler\"\n}\n",
		},
		{
			name: "multi-line, inline comment after the last item",
			src:  "a = {\n  hat = \"derby\" /* a hat */\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\" /* a hat */\n  dog = \"bowler\"\n}\n",
		},
		{
			name: "multi-line, comments before the brace",
			src:  "a = {\n  hat = \"derby\"\n\n  # end of a\n  // really\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n\n  # end of a\n  // really\n}\n",
		},
		{
			name: "multi-line, inline comment directly before the brace",
			src:  "a = {\n  hat = \"derby\"\n/* end of a */ }\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n/* end of a */ }\n",
		},
		{
			name: "multi-line, line comment after the opening brace",
			src:  "a = { # no items\n}\n",
			keys: []string{"dog"},
			want: "a = { # no items\n  dog = \"bowler\"\n}\n",
		},
		{
			name: "multi-line, aligned with existing items",
			src:  "a = {\n  hat = \"derby\"\n}\n",
			keys: []string{"a hat"},
			want: "a = {\n  hat     = \"derby\"\n  \"a hat\" = \"bowler\"\n}\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatal(diags)
			}
			object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()
			for _, key := range test.keys {
				object.SetItemValue(key, cty.StringVal("bowler"))
			}

			got := f.Bytes()
			if string(got) != test.want {
				t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, test.want)
			}
			if formatted := Format(got); string(formatted) != string(got) {
				t.Errorf("result is not in canonical form\ngot:\n%s\nformatted:\n%s", got, formatted)
			}
			if _, diags := ParseConfig(got, "", hcl.InitialPos); diags.HasErrors() {
				t.Errorf("result is invalid: %s", diags)
			}
		})
	}
}

```

#### y/ast_tuple_cons_expr.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

// TupleConsExpr represents a tuple-construct expression such as
// ["fedora", "fez"]. Get one from an expression with
// Expression.AsTupleConsExpr.
//
// Each element is an Expression of its own. The brackets and the commas stay
// as tokens of the tuple.
type TupleConsExpr struct {
	inTree

	elems nodeSet
}

func newTupleConsExpr() *TupleConsExpr {
	return &TupleConsExpr{
		inTree: newInTree(),
		elems:  newNodeSet(),
	}
}

// variableNodes returns the traversal nodes in the elements, in source
// order.
//
func (tuple *TupleConsExpr) variableNodes() []*node {
	var ret []*node
	for _, n := range tuple.elems.List() {
		ret = append(ret, n.content.(*Expression).variableNodes()...)
	}
	return ret
}

```

#### y/examples_test.go

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

func ExampleObjectConsExpr_SetItemRaw() {
	src := []byte("a = {\n    hat = \"derby\", (cat) = \"calico\" }\n")
	f, diags := hclwrite.ParseConfig(src, "", hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		fmt.Printf("errors: %s", diags)
		return
	}

	object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()
	bowler := hclwrite.TokensForValue(cty.StringVal("bowler"))

	// "hat" is found and replaced. A key from an expression, such as (cat),
	// is never found, so "cat" is added as a new item.
	object.SetItemRaw("hat", bowler)
	object.SetItemRaw("cat", bowler)

	fmt.Printf("%s", f.Bytes())
	// Output:
	// a = {
	//   hat = "bowler", (cat) = "calico"
	//   cat = "bowler"
	// }
}

```

#### y/node.go

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
	return &node{
		content: c,
	}
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
	}
	if after != nil {
		after.before = nn
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
	n := &node{
		content: c,
	}
	ns.AppendNode(n)
	n.list = ns
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
	n := &node{
		content: c,
	}
	ns.InsertNode(pos, n)
	n.list = ns
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
		pos.before.after = n
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

// inTree can be embedded into a content struct that has child nodes to get
// a standard implementation of the NodeContent interface and a record of
// a potential parent node.
type inTree struct {
	parent   *node
	children *nodes
}

func newInTree() inTree {
	return inTree{
		children: &nodes{},
	}
}

func (it *inTree) assertUnattached() {
	if it.parent != nil {
		panic(fmt.Sprintf("node is already attached to %T", it.parent.content))
	}
}

// setParent records the node that the content is attached under, so that a
// later assertUnattached can tell that it is already part of a syntax tree.
// A nil parent records that the content is no longer attached.
//
func (it *inTree) setParent(parent *node) {
	it.parent = parent
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

#### y/parser.go

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

	attrNode := newNode(attr)
	attr.Expr().setParent(attrNode)
	return attrNode
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

	// Function call expression
	case *hclsyntax.FunctionCallExpr:
		return parseFunctionCallExpr(tNativeExpr, from)

	// Tuple-construct expression
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
		value.Expr().setParent(item.value)

		node := expr.children.Append(item)
		expr.items.Add(node)
	}

	_, from, _ = from.Partition(nativeExpr.Range())
	children.AppendUnstructuredTokens(from.writerTokens)

	// Wrap in an Expression
	wrapExpr := newExpression()
	wrapExpr.wrap(expr)
	return newNode(wrapExpr)
}

// parseObjectConsKeyExpr parses an object-construct key expression.
// An object key is an Identifier or an Expression. In this context, a quoted
// literal is functionally equivalent to an Identifier.
func parseObjectConsKeyExpr(nativeExpr *hclsyntax.ObjectConsKeyExpr, from inputTokens) *node {
	expr := parseExpression(nativeExpr.Wrapped, from)
	// hclsyntax reads the key as a literal name by the same test.
	literalName := hcl.ExprAsKeyword(nativeExpr.Wrapped) != ""
	wrapExpr := newObjectConsKeyExpr(expr, literalName)

	return newNode(wrapExpr)
}

// parseFunctionCallExpr parses a function call expression, defined as:
//
//	FunctionCall = Identifier ("::" Identifier)* "(" Arguments ")";
//	Arguments = (
//	    () ||
//	    (Expression ("," Expression)* ("," | "...")?)
//	);
//
// Each argument becomes an Expression of its own. The tokens around the
// arguments, comments included, stay as unstructured tokens of the call.
//
func parseFunctionCallExpr(nativeExpr *hclsyntax.FunctionCallExpr, from inputTokens) *node {
	expr := newFunctionCallExpr()
	children := expr.children

	var before, argTokens inputTokens
	for _, nativeArg := range nativeExpr.Args {
		before, argTokens, from = from.Partition(nativeArg.Range())
		children.AppendUnstructuredTokens(before.writerTokens)
		arg := parseExpression(nativeArg, argTokens)
		children.AppendNode(arg)
		expr.args.Add(arg)
	}
	children.AppendUnstructuredTokens(from.writerTokens)

	// Wrap in an Expression
	wrapExpr := newExpression()
	wrapExpr.wrap(expr)
	for _, arg := range expr.args.List() {
		arg.content.(*Expression).setParent(wrapExpr.wrapped)
	}
	return newNode(wrapExpr)
}

// parseTupleConsExpr parses a tuple-construct expression, defined as:
//
//	tuple = "[" (
//	    (Expression ("," Expression)* ","?)?
//	) "]";
//
// Each element becomes an Expression of its own. The tokens around the
// elements, comments included, stay as unstructured tokens of the tuple.
//
func parseTupleConsExpr(nativeExpr *hclsyntax.TupleConsExpr, from inputTokens) *node {
	expr := newTupleConsExpr()
	children := expr.children

	var before, elemTokens inputTokens
	for _, nativeElem := range nativeExpr.Exprs {
		before, elemTokens, from = from.Partition(nativeElem.Range())
		children.AppendUnstructuredTokens(before.writerTokens)
		elem := parseExpression(nativeElem, elemTokens)
		children.AppendNode(elem)
		expr.elems.Add(elem)
	}
	children.AppendUnstructuredTokens(from.writerTokens)

	// Wrap in an Expression
	wrapExpr := newExpression()
	wrapExpr.wrap(expr)
	for _, elem := range expr.elems.List() {
		elem.content.(*Expression).setParent(wrapExpr.wrapped)
	}
	return newNode(wrapExpr)
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

#### y/parser_test.go

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
			`a = join(" ", ["fedora", "fez"])`,
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
													Type:     "Expression",
													Children: []TestTreeNode{{Type: "quoted", Val: ` "fez"`}},
												},
												{Type: "Tokens", Val: "]"},
											},
										}},
									},
									{Type: "Tokens", Val: ")"},
								},
							}},
						},
						{Type: "comments"},
					},
				}},
			},
		},
		{
			`a = [f(), {}, [], provider::ex::g(x...),]`,
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
									{Type: "Tokens", Val: " ["},
									{
										Type: "Expression",
										Children: []TestTreeNode{{
											Type:     "FunctionCallExpr",
											Children: []TestTreeNode{{Type: "Tokens", Val: "f()"}},
										}},
									},
									{Type: "Tokens", Val: ","},
									{
										Type: "Expression",
										Children: []TestTreeNode{{
											Type: "ObjectConsExpr",
											Children: []TestTreeNode{
												{Type: "Tokens", Val: " {"},
												{Type: "Tokens", Val: "}"},
											},
										}},
									},
									{Type: "Tokens", Val: ","},
									{
										Type: "Expression",
										Children: []TestTreeNode{{
											Type:     "TupleConsExpr",
											Children: []TestTreeNode{{Type: "Tokens", Val: " []"}},
										}},
									},
									{Type: "Tokens", Val: ","},
									{
										Type: "Expression",
										Children: []TestTreeNode{{
											Type: "FunctionCallExpr",
											Children: []TestTreeNode{
												{Type: "Tokens", Val: " provider::ex::g("},
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
									{Type: "Tokens", Val: ",]"},
								},
							}},
						},
						{Type: "comments"},
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

func TestParseFunctionCallAndTupleConsRoundTrip(t *testing.T) {
	tests := []string{
		"a = f()\n",
		"a = f(x...)\n",
		"a = f(x, y,)\n",
		"a = f( # c\n  x, // d\n  /* e */ y\n)\n",
		"a = f(<<EOT\nhello ${x}\nEOT\n, y)\n",
		"a = f(\"${x}\", (y), z ? 1 : 2, -1, !b)\n",
		"a = [<<EOT\nhello\nEOT\n]\n",
		"a = [\n  1, # one\n  [2, {k = f(old.x)}],\n]\n",
		"a = [for x in y : x]\n",
		"a = { k = f(old.x, [old.y]) }\n",
	}

	for _, src := range tests {
		t.Run(src, func(t *testing.T) {
			file, diags := parse([]byte(src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			// Bytes formats its result, so read the tokens as parsed.
			if got := string(file.BuildTokens(nil).Bytes()); got != src {
				t.Errorf("wrong result\ngot:  %q\nwant: %q", got, src)
			}
		})
	}
}

```

Answer with one JSON object and nothing else:
{"better": "x" | "y" | "tie", "reasons": ["...", "..."]}
Give two to five short reasons, each pointing at concrete files or types.
