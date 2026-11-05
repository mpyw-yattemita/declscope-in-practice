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
| ast_expression.go | 308 | CHANGED | Expression (type), newExpression (func), Expression.wrap (method), NewExpressionRaw (func), NewExpressionLiteral (func), NewExpressionAbsTraversal (func), Expression.Variables (method), Expression.RenameVariablePrefix (method), Expression.AsObjectConsExpr (method), Expression.AsFunctionCallExpr (method), Expression.AsTupleConsExpr (method), Expression.asQuoted (method), Expression.AsQuotedLiteral (method), Traversal (type), newTraversal (func), TraverseName (type), newTraverseName (func), TraverseIndex (type), newTraverseIndex (func) |
| ast_expression_test.go | 176 | NEW | TestExpressionVariablesObjectCons (func), TestExpressionVariablesCallAndTuple (func), astExpressionVariablesCase (type), checkASTExpressionVariables (func) |
| ast_object_cons_expr.go | 379 | CHANGED | ObjectConsExpr (type), newObjectConsExpr (func), ObjectConsExpr.Items (method), ObjectConsExpr.ItemFor (method), ObjectConsExpr.ValueExprFor (method), ObjectConsExpr.SetItem (method), ObjectConsExpr.SetItemRaw (method), ObjectConsExpr.SetItemValue (method), ObjectConsExpr.SetItemTraversal (method), ObjectConsExpr.variables (method), ObjectConsExpr.appendItem (method), ObjectConsExpr.isLineEnd (method), ObjectConsExpr.hasComma (method), ObjectConsExpr.newlineToken (method), ObjectConsItem (type), newObjectConsItem (func), ObjectConsItem.KeyExpr (method), ObjectConsItem.ValueExpr (method), ObjectConsKeyExpr (type), newObjectConsKeyExpr (func), ObjectConsKeyExpr.name (method), ObjectConsValue (type), newObjectConsValue (func), ObjectConsValue.Expr (method) |
| ast_object_cons_expr_test.go | 461 | NEW | TestObjectConsExprItemFor (func), TestObjectConsExprItems (func), TestObjectConsExprSetItemKey (func), TestObjectConsExprSetItemLayout (func), TestObjectConsExprSetItemHelpers (func), TestObjectConsExprSetItemPanics (func), TestObjectConsExprSetItemReplacedValue (func) |
| ast_test.go | 50 |  | TestTreeNode (type), makeTestTree (func) |
| doc.go | 14 |  |  |
| examples_test.go | 129 | CHANGED | Example_generateFromScratch (func), ExampleExpression_RenameVariablePrefix (func), ExampleObjectConsExpr_SetItemRaw (func) |
| format.go | 472 |  | format (func), formatIndent (func), formatSpaces (func), formatCells (func), spaceAfterToken (func), linesForFormat (func), tokenIsNewline (func), tokenBracketChange (func), formatLine (type) |
| format_test.go | 800 |  | TestFormat (func), TestLinesForFormat (func) |
| function_call_expr.go | 80 | NEW | FunctionCallExpr (type), FunctionCallExpr.Name (method), FunctionCallExpr.variables (method), parseFunctionCallExpr (func) |
| function_call_expr_test.go | 54 | NEW | TestFunctionCallExprName (func) |
| generate.go | 399 |  | TokensForValue (func), TokensForTraversal (func), TokensForIdentifier (func), TokensForTuple (func), TokensForObject (func), TokensForFunctionCall (func), appendTokensForValue (func), appendTokensForTraversal (func), appendTokensForTraversalStep (func), escapeQuotedStringLit (func), appendRune (func) |
| generate_test.go | 920 |  | TestTokensForValue (func), TestTokensForTraversal (func), TestTokensForTuple (func), TestTokensForObject (func), TestTokensForFunctionCall (func), TestTokenGenerateConsistency (func) |
| native_node_sorter.go | 26 |  | nativeNodeSorter (type), nativeNodeSorter.Len (method), nativeNodeSorter.Less (method), nativeNodeSorter.Swap (method) |
| node.go | 333 | CHANGED | node (type), newNode (func), node.Equal (method), node.BuildTokens (method), node.Detach (method), node.ReplaceWith (method), node.assertUnattached (method), node.setContentHolder (method), nodeContent (type), nodes (type), nodes.BuildTokens (method), nodes.Clear (method), nodes.Append (method), nodes.AppendNode (method), nodes.Insert (method), nodes.InsertNode (method), nodes.AppendUnstructuredTokens (method), nodes.FindNodeWithContent (method), nodeSet (type), newNodeSet (func), nodeSet.Has (method), nodeSet.Add (method), nodeSet.Remove (method), nodeSet.Clear (method), nodeSet.List (method), nodeSet.FindNodeWithContent (method), internalWalkFunc (type), inTree (type), newInTree (func), inTree.assertUnattached (method), inTree.setHolder (method), inTree.walkChildNodes (method), inTree.BuildTokens (method), leafNode (type), leafNode.walkChildNodes (method) |
| parser.go | 745 | CHANGED | parse (func), inputTokens (type), inputTokens.Partition (method), inputTokens.PartitionType (method), inputTokens.PartitionTypeOk (method), inputTokens.PartitionTypeSingle (method), inputTokens.PartitionIncludingComments (method), inputTokens.PartitionBlockItem (method), inputTokens.PartitionLeadComments (method), inputTokens.PartitionLineEndTokens (method), inputTokens.Slice (method), inputTokens.Len (method), inputTokens.Tokens (method), inputTokens.Types (method), parseBody (func), parseBodyItem (func), parseAttribute (func), parseBlock (func), parseBlockLabels (func), parseExpression (func), parseAnyExpression (func), parseObjectConsExpr (func), parseObjectConsKeyExpr (func), parseTemplateExpr (func), parseTraversal (func), parseTraversalStep (func), writerTokens (func), partitionTokens (func), partitionLeadCommentTokens (func), partitionLineEndTokens (func), lexConfig (func) |
| parser_test.go | 2404 | CHANGED | TestParse (func), TestPartitionTokens (func), TestPartitionLeadCommentTokens (func), TestLexConfig (func) |
| public.go | 48 |  | NewFile (func), ParseConfig (func), Format (func) |
| round_trip_test.go | 181 |  | TestRoundTripVerbatim (func), TestRoundTripFormat (func), TestRoundTripSafeConcurrent (func) |
| tokens.go | 136 |  | Token (type), Token.asHCLSyntax (method), Tokens (type), Tokens.Bytes (method), Tokens.testValue (method), Tokens.Columns (method), Tokens.WriteTo (method), Tokens.walkChildNodes (method), Tokens.BuildTokens (method), ObjectAttrTokens (type), newIdentToken (func) |
| tuple_cons_expr.go | 61 | NEW | TupleConsExpr (type), TupleConsExpr.variables (method), parseTupleConsExpr (func) |
| tuple_cons_expr_test.go | 43 | NEW | TestTupleConsExprAsTupleConsExpr (func) |

### y/

| file | lines | mark | declarations |
| --- | --- | --- | --- |
| ast.go | 121 |  | File (type), NewEmptyFile (func), File.Body (method), File.WriteTo (method), File.Bytes (method), comments (type), newComments (func), comments.BuildTokens (method), identifier (type), newIdentifier (func), identifier.BuildTokens (method), identifier.hasName (method), number (type), newNumber (func), number.BuildTokens (method), quoted (type), newQuoted (func), quoted.BuildTokens (method) |
| ast_attribute.go | 64 | CHANGED | Attribute (type), newAttribute (func), Attribute.init (method), Attribute.LeadComments (method), Attribute.LineComments (method), Attribute.Expr (method), Attribute.setName (method) |
| ast_attribute_test.go | 125 |  | TestAttributeLeadComments (func), TestAttributeLineComments (func) |
| ast_block.go | 184 |  | Block (type), newBlock (func), NewBlock (func), Block.init (method), Block.Body (method), Block.LeadComments (method), Block.Type (method), Block.SetType (method), Block.Labels (method), Block.SetLabels (method), Block.labelsObj (method), blockLabels (type), newBlockLabels (func), blockLabels.Replace (method), blockLabels.Current (method) |
| ast_block_test.go | 579 |  | TestBlockType (func), TestBlockLabels (func), TestBlockSetType_buildTokens (func), TestBlockSetType (func), TestBlockSetLabels (func), TestBlockLeadComments (func), mustParseConfig (func) |
| ast_body.go | 309 |  | Body (type), newBody (func), Body.appendItem (method), Body.appendItemNode (method), Body.Clear (method), Body.AppendUnstructuredTokens (method), Body.Attributes (method), Body.Blocks (method), Body.GetAttribute (method), Body.getAttributeNode (method), Body.RenameAttribute (method), Body.FirstMatchingBlock (method), Body.RemoveNewlineBeforeBlock (method), Body.RemoveBlock (method), Body.SetAttributeRaw (method), Body.SetAttributeValue (method), Body.SetAttributeTraversal (method), Body.RemoveAttribute (method), Body.AppendBlock (method), Body.AppendNewBlock (method), Body.AppendNewline (method) |
| ast_body_test.go | 2195 |  | TestBodyGetAttribute (func), TestBodyFirstMatchingBlock (func), TestBodySetAttributeValue (func), TestBodySetAttributeTraversal (func), TestBodySetAttributeTraversal_ReturnsTheAttribute (func), TestBodySetAttributeRaw (func), TestBodySetAttributeRaw_ReturnsTheAttribute (func), TestBodySetAttributeValue_ReturnsTheAttribute (func), TestBodySetAttributeValueInBlock (func), TestBodySetAttributeValueInNestedBlock (func), TestBodyRemoveAttribute (func), TestBodyRenameAttribute (func), TestBodyAppendBlock (func), TestBodyRemoveBlock (func), TestBodyRemoveNewlineBeforeBlock (func), testFn (func) |
| ast_expression.go | 307 | CHANGED | Expression (type), newExpression (func), Expression.wrap (method), NewExpressionRaw (func), NewExpressionLiteral (func), NewExpressionAbsTraversal (func), Expression.Variables (method), Expression.RenameVariablePrefix (method), Expression.AsObjectConsExpr (method), Expression.AsFunctionCallExpr (method), Expression.AsTupleConsExpr (method), Expression.asQuoted (method), Expression.AsQuotedLiteral (method), Traversal (type), newTraversal (func), TraverseName (type), newTraverseName (func), TraverseIndex (type), newTraverseIndex (func) |
| ast_function_call_expr.go | 50 | NEW | FunctionCallExpr (type), newFunctionCallExpr (func), FunctionCallExpr.Name (method), FunctionCallExpr.variables (method), expressionsVariables (func) |
| ast_function_call_expr_test.go | 197 | NEW | parseAttrExprForTest (func), TestFunctionCallExprName (func), TestExpressionAsFunctionCallAndTupleCons (func), TestFunctionCallAndTupleConsVariables (func), TestFunctionCallAndTupleConsRoundTrip (func) |
| ast_object_cons_expr.go | 372 | CHANGED | ObjectConsExpr (type), newObjectConsExpr (func), ObjectConsExpr.Items (method), ObjectConsExpr.ItemFor (method), ObjectConsExpr.ValueExprFor (method), ObjectConsExpr.SetItem (method), ObjectConsExpr.SetItemRaw (method), ObjectConsExpr.SetItemValue (method), ObjectConsExpr.SetItemTraversal (method), ObjectConsExpr.appendItem (method), splitObjectTail (func), ObjectConsExpr.isMultiline (method), hasLineEnd (func), isLineEnd (func), newNewlineToken (func), ObjectConsExpr.variables (method), ObjectConsItem (type), newObjectConsItem (func), ObjectConsItem.init (method), ObjectConsItem.KeyExpr (method), ObjectConsItem.ValueExpr (method), ObjectConsKeyExpr (type), newObjectConsKeyExpr (func), newObjectConsKeyExprFor (func), ObjectConsKeyExpr.literalName (method), ObjectConsKeyExpr.variables (method), ObjectConsValue (type), newObjectConsValue (func), ObjectConsValue.Expr (method) |
| ast_object_cons_expr_test.go | 636 | NEW | TestObjectConsExprVariables (func), parseObjectConsExprForTest (func), tokensSource (func), TestObjectConsExprItems (func), TestObjectConsExprItemFor (func), TestObjectConsExprSetItemKey (func), TestObjectConsExprSetItemReplace (func), TestObjectConsExprSetItemHelpers (func), TestObjectConsExprSetItemPanics (func), TestObjectConsExprSetItemLayout (func) |
| ast_test.go | 50 |  | TestTreeNode (type), makeTestTree (func) |
| ast_tuple_cons_expr.go | 27 | NEW | TupleConsExpr (type), newTupleConsExpr (func), TupleConsExpr.variables (method) |
| doc.go | 14 |  |  |
| examples_test.go | 129 | CHANGED | Example_generateFromScratch (func), ExampleExpression_RenameVariablePrefix (func), ExampleObjectConsExpr_SetItemRaw (func) |
| format.go | 472 |  | format (func), formatIndent (func), formatSpaces (func), formatCells (func), spaceAfterToken (func), linesForFormat (func), tokenIsNewline (func), tokenBracketChange (func), formatLine (type) |
| format_test.go | 800 |  | TestFormat (func), TestLinesForFormat (func) |
| generate.go | 399 |  | TokensForValue (func), TokensForTraversal (func), TokensForIdentifier (func), TokensForTuple (func), TokensForObject (func), TokensForFunctionCall (func), appendTokensForValue (func), appendTokensForTraversal (func), appendTokensForTraversalStep (func), escapeQuotedStringLit (func), appendRune (func) |
| generate_test.go | 920 |  | TestTokensForValue (func), TestTokensForTraversal (func), TestTokensForTuple (func), TestTokensForObject (func), TestTokensForFunctionCall (func), TestTokenGenerateConsistency (func) |
| native_node_sorter.go | 26 |  | nativeNodeSorter (type), nativeNodeSorter.Len (method), nativeNodeSorter.Less (method), nativeNodeSorter.Swap (method) |
| node.go | 325 | CHANGED | node (type), newNode (func), node.Equal (method), node.BuildTokens (method), node.Detach (method), node.ReplaceWith (method), node.assertUnattached (method), nodeContent (type), nodes (type), nodes.BuildTokens (method), nodes.Clear (method), nodes.Append (method), nodes.AppendNode (method), nodes.Insert (method), nodes.InsertNode (method), nodes.AppendUnstructuredTokens (method), nodes.FindNodeWithContent (method), nodeSet (type), newNodeSet (func), nodeSet.Has (method), nodeSet.Add (method), nodeSet.Remove (method), nodeSet.Clear (method), nodeSet.List (method), nodeSet.FindNodeWithContent (method), internalWalkFunc (type), inTree (type), newInTree (func), inTree.assertUnattached (method), inTree.setAttached (method), setAttached (func), inTree.walkChildNodes (method), inTree.BuildTokens (method), leafNode (type), leafNode.walkChildNodes (method) |
| parser.go | 802 | CHANGED | parse (func), inputTokens (type), inputTokens.Partition (method), inputTokens.PartitionType (method), inputTokens.PartitionTypeOk (method), inputTokens.PartitionTypeSingle (method), inputTokens.PartitionIncludingComments (method), inputTokens.PartitionBlockItem (method), inputTokens.PartitionLeadComments (method), inputTokens.PartitionLineEndTokens (method), inputTokens.Slice (method), inputTokens.Len (method), inputTokens.Tokens (method), inputTokens.Types (method), parseBody (func), parseBodyItem (func), parseAttribute (func), parseBlock (func), parseBlockLabels (func), parseExpression (func), parseAnyExpression (func), parseObjectConsExpr (func), parseFunctionCallExpr (func), parseTupleConsExpr (func), parseExpressionList (func), parseObjectConsKeyExpr (func), parseTemplateExpr (func), parseTraversal (func), parseTraversalStep (func), writerTokens (func), partitionTokens (func), partitionLeadCommentTokens (func), partitionLineEndTokens (func), lexConfig (func) |
| parser_test.go | 2422 | CHANGED | TestParse (func), TestPartitionTokens (func), TestPartitionLeadCommentTokens (func), TestLexConfig (func) |
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
// expression. For an object-construct expression, these are the traversals
// in the keys and values of its items. For a function call or a
// tuple-construct expression, these are the traversals in its arguments or
// elements.
func (e *Expression) Variables() []*Traversal {
	nodes := e.absTraversals.List()
	ret := make([]*Traversal, len(nodes))
	for i, node := range nodes {
		ret[i] = node.content.(*Traversal)
	}
	if object := e.AsObjectConsExpr(); object != nil {
		ret = append(ret, object.variables()...)
	}
	if call := e.AsFunctionCallExpr(); call != nil {
		ret = append(ret, call.variables()...)
	}
	if tuple := e.AsTupleConsExpr(); tuple != nil {
		ret = append(ret, tuple.variables()...)
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
	for _, traversal := range e.Variables() {
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

#### x/ast_expression_test.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

func TestExpressionVariablesObjectCons(t *testing.T) {
	checkASTExpressionVariables(t, []astExpressionVariablesCase{
		{
			"a = { k = old.x }\n",
			[]string{"old.x"},
			"a = { k = new.x }\n",
		},
		{
			"a = { k = old.x, j = old.y }\n",
			[]string{"old.x", "old.y"},
			"a = { k = new.x, j = new.y }\n",
		},
		{
			"a = { k = { j = old.x } }\n",
			[]string{"old.x"},
			"a = { k = { j = new.x } }\n",
		},
		{
			"a = { k = [old.x] }\n",
			[]string{"old.x"},
			"a = { k = [new.x] }\n",
		},
		{
			"a = { k = f(old.x) }\n",
			[]string{"old.x"},
			"a = { k = f(new.x) }\n",
		},
		{
			"a = { old = 1 }\n",
			nil,
			"a = { old = 1 }\n",
		},
		{
			"a = { old = old.x }\n",
			[]string{"old.x"},
			"a = { old = new.x }\n",
		},
		{
			"a = { (old.x) = 1 }\n",
			[]string{"old.x"},
			"a = { (new.x) = 1 }\n",
		},
		{
			"a = { (old) = 1 }\n",
			[]string{"old"},
			"a = { (new) = 1 }\n",
		},
		{
			"a = { \"k\" = old.x }\n",
			[]string{"old.x"},
			"a = { \"k\" = new.x }\n",
		},
		{
			"a = { (old.x) = old.y, k = { old = [f(old.z)] } }\n",
			[]string{"old.x", "old.y", "old.z"},
			"a = { (new.x) = new.y, k = { old = [f(new.z)] } }\n",
		},
	})
}

func TestExpressionVariablesCallAndTuple(t *testing.T) {
	checkASTExpressionVariables(t, []astExpressionVariablesCase{
		{
			"a = f(old.x)\n",
			[]string{"old.x"},
			"a = f(new.x)\n",
		},
		{
			"a = f(old.x, g(old.y), [old.z])\n",
			[]string{"old.x", "old.y", "old.z"},
			"a = f(new.x, g(new.y), [new.z])\n",
		},
		{
			"a = f(old.x...)\n",
			[]string{"old.x"},
			"a = f(new.x...)\n",
		},
		{
			"a = [old.x, [old.y], f(old.z)]\n",
			[]string{"old.x", "old.y", "old.z"},
			"a = [new.x, [new.y], f(new.z)]\n",
		},
		{
			"a = [old.x, { k = old.y }]\n",
			[]string{"old.x", "old.y"},
			"a = [new.x, { k = new.y }]\n",
		},
		{
			"a = [\n  old.x, # old.w\n  f(\n    old.y,\n  ),\n]\n",
			[]string{"old.x", "old.y"},
			"a = [\n  new.x, # old.w\n  f(\n    new.y,\n  ),\n]\n",
		},
		{
			"a = f(old.x + old.y, old.z[0])\n",
			[]string{"old.x", "old.y", "old.z[0]"},
			"a = f(new.x + new.y, new.z[0])\n",
		},
		{
			"a = f()\n",
			nil,
			"a = f()\n",
		},
		{
			"a = []\n",
			nil,
			"a = []\n",
		},
	})
}

// astExpressionVariablesCase is a source whose expression a has its
// variables checked by checkASTExpressionVariables.
type astExpressionVariablesCase struct {
	src     string
	want    []string // the traversals Variables reports, as source text
	renamed string   // the source after renaming old to new
}

// checkASTExpressionVariables checks that Variables reports the traversals of
// each case in source order, as many as hclsyntax does, and that
// RenameVariablePrefix renames them.
func checkASTExpressionVariables(t *testing.T, tests []astExpressionVariablesCase) {
	t.Helper()
	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			if got := string(f.Bytes()); got != test.src {
				t.Errorf("source does not round-trip\ngot:  %q\nwant: %q", got, test.src)
			}
			expr := f.Body().GetAttribute("a").Expr()

			var got []string
			for _, traversal := range expr.Variables() {
				got = append(got, strings.TrimSpace(string(traversal.BuildTokens(nil).Bytes())))
			}
			if len(got) != len(test.want) {
				t.Fatalf("wrong variables\ngot:  %q\nwant: %q", got, test.want)
			}
			for i := range got {
				if got[i] != test.want[i] {
					t.Fatalf("wrong variables\ngot:  %q\nwant: %q", got, test.want)
				}
			}

			native, diags := hclsyntax.ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected hclsyntax diagnostics: %s", diags)
			}
			nativeExpr := native.Body.(*hclsyntax.Body).Attributes["a"].Expr
			if want := len(nativeExpr.Variables()); len(got) != want {
				t.Errorf("got %d variables, but hclsyntax reports %d", len(got), want)
			}

			expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
			if got := string(f.Bytes()); got != test.renamed {
				t.Errorf("wrong result after renaming\ngot:  %q\nwant: %q", got, test.renamed)
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
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

// ObjectConsExpr is an object-construct expression, such as
// { hat = "derby" }. Use Expression.AsObjectConsExpr to get one from a parsed
// expression.
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
	ret := make([]*ObjectConsItem, 0, len(nodes))
	for _, n := range nodes {
		ret = append(ret, n.content.(*ObjectConsItem))
	}
	return ret
}

// ItemFor returns the first item whose key is the given string, or nil if
// there is no such item.
//
// Only a key written as an identifier or as a quoted string literal can
// match, and the two forms are equivalent: "hat" finds both hat = ... and
// "hat" = .... A quoted key is compared by its value, after its escape
// sequences are decoded. A key taken from a variable or another expression,
// such as (hat), hat.band or "${hat}", never matches.
func (o *ObjectConsExpr) ItemFor(key string) *ObjectConsItem {
	for _, item := range o.Items() {
		if name, ok := item.KeyExpr().name(); ok && name == key {
			return item
		}
	}
	return nil
}

// ValueExprFor returns the value of the item that ItemFor finds for the given
// key, or nil if there is no such item.
func (o *ObjectConsExpr) ValueExprFor(key string) *ObjectConsValue {
	item := o.ItemFor(key)
	if item == nil {
		return nil
	}
	return item.ValueExpr()
}

// SetItem either replaces the value of the item that ItemFor finds for the
// given key, or adds a new item with the given key to the end of the object.
// The key of an existing item is left as written. A new key is written as an
// identifier if it is a valid one, and as a quoted string literal otherwise,
// escaped so that it evaluates to the given string. The key "for" is always
// quoted, since { for = ... } would read as a for expression.
//
// The given expression must not already be part of a syntax tree, or SetItem
// panics. Reusing an attached expression would make two items share it.
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
// NewExpressionLiteral does. To set a variable reference or other traversal,
// use SetItemTraversal.
func (o *ObjectConsExpr) SetItemValue(key string, val cty.Value) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionLiteral(val))
}

// SetItemTraversal is like SetItem, but writes the given traversal, which
// must be absolute, as NewExpressionAbsTraversal does. To set a literal value,
// use SetItemValue.
func (o *ObjectConsExpr) SetItemTraversal(key string, traversal hcl.Traversal) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionAbsTraversal(traversal))
}

// variables returns the absolute traversals in the keys and values of the
// object's items, in source order. A naked identifier key is a literal name,
// not a variable, so it contributes nothing.
//
func (o *ObjectConsExpr) variables() []*Traversal {
	var ret []*Traversal
	for _, n := range o.items.List() {
		item, ok := n.content.(*ObjectConsItem)
		if !ok {
			continue
		}
		if key := item.key.content.(*ObjectConsKeyExpr); !key.literal {
			ret = append(ret, key.wrapped.content.(*Expression).Variables()...)
		}
		ret = append(ret, item.value.content.(*ObjectConsValue).Expr().Variables()...)
	}
	return ret
}

// appendItem adds a new item with the given key and value after the last
// item of the object, and returns it.
//
// In an object written on one line, the item goes on the same line, after a
// comma. Otherwise it goes on its own line. A comma or a line comment that
// follows the last item on its line stays with that item. Whatever follows
// on later lines, and an inline comment directly before the closing brace,
// stays with the brace.
func (o *ObjectConsExpr) appendItem(key string, expr *Expression) *ObjectConsItem {
	itemNodes := o.items.List()
	stream := o.BuildTokens(nil)

	// Line breaks inside the keys and values of the items do not make the
	// object multi-line.
	inner := make(map[*Token]bool)
	for _, n := range itemNodes {
		item := n.content.(*ObjectConsItem)
		for _, tok := range item.KeyExpr().BuildTokens(nil) {
			inner[tok] = true
		}
		for _, tok := range item.ValueExpr().BuildTokens(nil) {
			inner[tok] = true
		}
	}
	multiline := false
	for _, tok := range stream {
		if !inner[tok] && o.isLineEnd(tok) {
			multiline = true
			break
		}
	}

	// Split the tokens of the object into the gaps before each item and the
	// tail after the last item, or after the opening brace if there are no
	// items. The tail ends with the closing brace.
	gaps := make([]Tokens, 0, len(itemNodes)+1)
	pos := 0
	for _, n := range itemNodes {
		first := n.BuildTokens(nil)[0]
		start := pos
		for stream[start] != first {
			start++
		}
		gaps = append(gaps, stream[pos:start:start])
		pos = start + len(n.BuildTokens(nil))
	}
	tail := stream[pos:]
	if len(itemNodes) == 0 {
		for i, tok := range tail {
			if tok.Type == hclsyntax.TokenOBrace {
				gaps = append(gaps, tail[:i+1:i+1])
				tail = tail[i+1:]
				break
			}
		}
	}

	// Rebuild the children from the gaps and the items, so that the tail
	// can be rearranged around the new item.
	for _, n := range itemNodes {
		n.Detach()
	}
	o.children.Clear()
	for i, gap := range gaps {
		o.children.AppendUnstructuredTokens(gap)
		if i < len(itemNodes) {
			o.children.AppendNode(itemNodes[i])
		}
	}

	// The tokens before split stay with the last item, and the rest with
	// the closing brace.
	split := len(tail) - 1
	for split > 0 && tail[split-1].Type == hclsyntax.TokenComment && !o.isLineEnd(tail[split-1]) {
		split--
	}
	if multiline {
		for i, tok := range tail[:split] {
			if o.isLineEnd(tok) {
				split = i + 1
				break
			}
		}
	}
	before, after := tail[:split], tail[split:]

	item := newObjectConsItem()
	switch {
	case multiline && (len(before) == 0 || !o.isLineEnd(before[len(before)-1])):
		// The closing brace shared a line with the last item.
		item.children.AppendUnstructuredTokens(Tokens{o.newlineToken()})
	case !multiline && len(itemNodes) > 0 && !o.hasComma(before):
		item.children.AppendUnstructuredTokens(Tokens{{
			Type:  hclsyntax.TokenComma,
			Bytes: []byte{','},
		}})
	}

	// An object that starts with the identifier for is read as a for
	// expression, so that key is quoted too.
	var keyExpr *Expression
	literal := hclsyntax.ValidIdentifier(key) && key != "for"
	if literal {
		keyExpr = NewExpressionAbsTraversal(hcl.Traversal{hcl.TraverseRoot{Name: key}})
	} else {
		keyExpr = NewExpressionRaw(TokensForValue(cty.StringVal(key)))
	}
	if !multiline {
		keyExpr.BuildTokens(nil)[0].SpacesBefore = 1
	}
	item.key = item.children.Append(newObjectConsKeyExpr(newNode(keyExpr), literal))
	item.children.AppendUnstructuredTokens(Tokens{{
		Type:         hclsyntax.TokenEqual,
		Bytes:        []byte{'='},
		SpacesBefore: 1,
	}})
	value := newObjectConsValue()
	value.expr = value.children.Append(expr)
	item.value = item.children.Append(value)

	if multiline {
		after = append(Tokens{o.newlineToken()}, after...)
	} else {
		after[0].SpacesBefore = 1
	}

	o.children.AppendUnstructuredTokens(before)
	o.items.Add(o.children.Append(item))
	o.children.AppendUnstructuredTokens(after)
	return item
}

// isLineEnd reports whether tok ends a line. A line comment ends the line
// it is on, since it includes its newline.
func (o *ObjectConsExpr) isLineEnd(tok *Token) bool {
	switch tok.Type {
	case hclsyntax.TokenNewline:
		return true
	case hclsyntax.TokenComment:
		return len(tok.Bytes) > 0 && tok.Bytes[len(tok.Bytes)-1] == '\n'
	}
	return false
}

func (o *ObjectConsExpr) hasComma(tokens Tokens) bool {
	for _, tok := range tokens {
		if tok.Type == hclsyntax.TokenComma {
			return true
		}
	}
	return false
}

func (o *ObjectConsExpr) newlineToken() *Token {
	return &Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte{'\n'},
	}
}

// ObjectConsItem is one item of an object-construct expression: a key and
// its value.
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

// ObjectConsKeyExpr is the key of an item of an object-construct expression.
// It is an identifier, a quoted string, or another expression whose value
// gives the key.
type ObjectConsKeyExpr struct {
	inTree

	wrapped *node
	// literal is set when the key is a naked identifier, which names the
	// attribute literally instead of referring to a variable.
	literal bool
}

func newObjectConsKeyExpr(wrapped *node, literal bool) *ObjectConsKeyExpr {
	expr := &ObjectConsKeyExpr{
		inTree:  newInTree(),
		wrapped: wrapped,
		literal: literal,
	}
	expr.children.AppendNode(wrapped)

	return expr
}

// name returns the string the key gives, if the key is an identifier or a
// quoted string literal. Otherwise it returns false.
func (k *ObjectConsKeyExpr) name() (string, bool) {
	tokens := k.BuildTokens(nil)
	if len(tokens) == 1 && tokens[0].Type == hclsyntax.TokenIdent {
		return string(tokens[0].Bytes), true
	}
	if len(tokens) < 2 || tokens[0].Type != hclsyntax.TokenOQuote || tokens[len(tokens)-1].Type != hclsyntax.TokenCQuote {
		return "", false
	}
	var b strings.Builder
	for _, tok := range tokens[1 : len(tokens)-1] {
		if tok.Type != hclsyntax.TokenQuotedLit {
			return "", false
		}
		s, diags := hclsyntax.ParseStringLiteralToken(hclsyntax.Token{Type: tok.Type, Bytes: tok.Bytes})
		if diags.HasErrors() {
			return "", false
		}
		b.WriteString(s)
	}
	return b.String(), true
}

// ObjectConsValue is the value of an item of an object-construct expression.
type ObjectConsValue struct {
	inTree
	expr *node
}

func newObjectConsValue() *ObjectConsValue {
	return &ObjectConsValue{
		inTree: newInTree(),
	}
}

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
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

func TestObjectConsExprItemFor(t *testing.T) {
	tests := []struct {
		src  string
		key  string
		want string // the value found, as source text, or "" for none
	}{
		{`a = { hat = "derby" }`, "hat", `"derby"`},
		{`a = { "hat" = "derby" }`, "hat", `"derby"`},
		{`a = { "$${hat}" = "derby" }`, "${hat}", `"derby"`},
		{`a = { "%%{hat}" = "derby" }`, "%{hat}", `"derby"`},
		{`a = { "a hat" = "derby" }`, "a hat", `"derby"`},
		{`a = { "say \"hi\"" = "derby" }`, `say "hi"`, `"derby"`},
		{`a = { "a\\b" = "derby" }`, `a\b`, `"derby"`},
		{`a = { "" = "derby" }`, "", `"derby"`},
		{`a = { true = "derby" }`, "true", `"derby"`},
		{`a = { hat: "derby" }`, "hat", `"derby"`},
		{`a = { cat = "calico", hat = "derby" }`, "hat", `"derby"`},
		{`a = { (hat) = "derby" }`, "hat", ""},
		{`a = { hat.band = "derby" }`, "hat.band", ""},
		{`a = { hat.band = "derby" }`, "hat", ""},
		{`a = { "${hat}" = "derby" }`, "hat", ""},
		{`a = { "${hat}" = "derby" }`, "${hat}", ""},
		{`a = { "$${hat}" = "derby" }`, "$${hat}", ""},
		{`a = { hat = "derby" }`, "cat", ""},
		{`a = { hat = "derby" }`, "Hat", ""},
		{`a = {}`, "hat", ""},
	}

	for _, test := range tests {
		t.Run(test.src+" "+test.key, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src+"\n"), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

			item := object.ItemFor(test.key)
			value := object.ValueExprFor(test.key)
			if test.want == "" {
				if item != nil || value != nil {
					t.Fatalf("found an item, but want none")
				}
				return
			}
			if item == nil || value == nil {
				t.Fatalf("found no item")
			}
			if value != item.ValueExpr() {
				t.Errorf("ValueExprFor returned a value other than that of ItemFor")
			}
			if got := strings.TrimSpace(string(value.BuildTokens(nil).Bytes())); got != test.want {
				t.Errorf("wrong value\ngot:  %s\nwant: %s", got, test.want)
			}
		})
	}
}

func TestObjectConsExprItems(t *testing.T) {
	src := "a = { hat = 1, (cat) = 2, \"dog\" = 3 }\n"
	f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics: %s", diags)
	}
	items := f.Body().GetAttribute("a").Expr().AsObjectConsExpr().Items()

	var got []string
	for _, item := range items {
		key := strings.TrimSpace(string(item.KeyExpr().BuildTokens(nil).Bytes()))
		value := strings.TrimSpace(string(item.ValueExpr().BuildTokens(nil).Bytes()))
		got = append(got, key+"="+value)
	}
	want := []string{"hat=1", "(cat)=2", `"dog"=3`}
	if strings.Join(got, " ") != strings.Join(want, " ") {
		t.Errorf("wrong items\ngot:  %q\nwant: %q", got, want)
	}
}

func TestObjectConsExprSetItemKey(t *testing.T) {
	tests := []struct {
		key  string
		want string // the key as written
	}{
		{"dog", `dog`},
		{"dog-1", `dog-1`},
		{"true", `true`},
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
		{"line\nbreak", `"line\nbreak"`},
	}

	for _, test := range tests {
		t.Run(test.key, func(t *testing.T) {
			f, diags := ParseConfig([]byte("a = {}\n"), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

			key, _ := object.SetItemValue(test.key, cty.True)
			if got := strings.TrimSpace(string(key.BuildTokens(nil).Bytes())); got != test.want {
				t.Errorf("wrong key\ngot:  %s\nwant: %s", got, test.want)
			}
			if item := object.ItemFor(test.key); item == nil || item.KeyExpr() != key {
				t.Errorf("ItemFor does not find the key that was set")
			}

			// The key must evaluate back to the original string.
			src := f.Bytes()
			native, diags := hclsyntax.ParseConfig(src, "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("result is not valid HCL: %s\n%s", diags, src)
			}
			val, diags := native.Body.(*hclsyntax.Body).Attributes["a"].Expr.Value(nil)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			if !val.Type().IsObjectType() || !val.Type().HasAttribute(test.key) {
				t.Errorf("result has no key %q: %#v", test.key, val)
			}
		})
	}
}

func TestObjectConsExprSetItemLayout(t *testing.T) {
	bowler := TokensForValue(cty.StringVal("bowler"))

	tests := []struct {
		name string
		src  string
		keys []string // the keys set, in order, each to "bowler"
		want string
	}{
		{
			"single line",
			"a = { hat = \"derby\" }\n",
			[]string{"dog"},
			"a = { hat = \"derby\", dog = \"bowler\" }\n",
		},
		{
			"empty",
			"a = {}\n",
			[]string{"dog"},
			"a = { dog = \"bowler\" }\n",
		},
		{
			"empty, two items",
			"a = {}\n",
			[]string{"dog", "cat"},
			"a = { dog = \"bowler\", cat = \"bowler\" }\n",
		},
		{
			"single line with trailing comma",
			"a = { hat = \"derby\", }\n",
			[]string{"dog"},
			"a = { hat = \"derby\", dog = \"bowler\" }\n",
		},
		{
			"single line with trailing comma, two items",
			"a = { hat = \"derby\", }\n",
			[]string{"dog", "cat"},
			"a = { hat = \"derby\", dog = \"bowler\", cat = \"bowler\" }\n",
		},
		{
			"single line with comment before the brace",
			"a = { hat = \"derby\" /* end of a */ }\n",
			[]string{"dog"},
			"a = { hat = \"derby\", dog = \"bowler\" /* end of a */ }\n",
		},
		{
			"single line, replace",
			"a = { hat = \"derby\", cat = \"calico\" }\n",
			[]string{"hat"},
			"a = { hat = \"bowler\", cat = \"calico\" }\n",
		},
		{
			"single line, replace quoted",
			"a = { \"hat\" = \"derby\" }\n",
			[]string{"hat"},
			"a = { \"hat\" = \"bowler\" }\n",
		},
		{
			"single line with nested multi-line value",
			"a = { hat = {\n  band = \"red\"\n} }\n",
			[]string{"dog"},
			"a = { hat = {\n  band = \"red\"\n}, dog = \"bowler\" }\n",
		},
		{
			"multi-line",
			"a = {\n  hat = \"derby\"\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}\n",
		},
		{
			"multi-line, two items, aligned",
			"a = {\n  hat = \"derby\"\n}\n",
			[]string{"dog", "color"},
			"a = {\n  hat   = \"derby\"\n  dog   = \"bowler\"\n  color = \"bowler\"\n}\n",
		},
		{
			"multi-line, replace",
			"a = {\n  hat = \"derby\"\n  cat = \"calico\"\n}\n",
			[]string{"hat"},
			"a = {\n  hat = \"bowler\"\n  cat = \"calico\"\n}\n",
		},
		{
			"multi-line empty",
			"a = {\n}\n",
			[]string{"dog"},
			"a = {\n  dog = \"bowler\"\n}\n",
		},
		{
			"multi-line with commas",
			"a = {\n  hat = \"derby\",\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\",\n  dog = \"bowler\"\n}\n",
		},
		{
			"brace on the line of the last item",
			"a = {\n  hat = \"derby\" }\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}\n",
		},
		{
			"brace on the line of the last item, two items",
			"a = {\n  hat = \"derby\" }\n",
			[]string{"dog", "cat"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n}\n",
		},
		{
			"brace on the line of the last item, after a comma",
			"a = {\n  hat = \"derby\", }\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\",\n  dog = \"bowler\"\n}\n",
		},
		{
			"brace on the line of the first item",
			"a = { hat = \"derby\",\n  cat = \"calico\" }\n",
			[]string{"dog"},
			"a = { hat = \"derby\",\n  cat = \"calico\"\n  dog = \"bowler\"\n}\n",
		},
		{
			"line comment after the last item",
			"a = {\n  hat = \"derby\" # felt\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\" # felt\n  dog = \"bowler\"\n}\n",
		},
		{
			"comma and line comment after the last item",
			"a = {\n  hat = \"derby\", # felt\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\", # felt\n  dog = \"bowler\"\n}\n",
		},
		{
			"comments before the brace",
			"a = {\n  hat = \"derby\"\n\n  # more hats\n  // later\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n\n  # more hats\n  // later\n}\n",
		},
		{
			"comment before the brace in an empty object",
			"a = {\n  # no hats\n}\n",
			[]string{"dog"},
			"a = {\n  dog = \"bowler\"\n  # no hats\n}\n",
		},
		{
			"inline comment directly before the brace",
			"a = {\n  hat = \"derby\"\n/* end of a */ }\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n/* end of a */ }\n",
		},
		{
			"inline comment directly before the brace, on the line of the last item",
			"a = {\n  hat = \"derby\" /* end of a */ }\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n/* end of a */ }\n",
		},
		{
			"nested multi-line value",
			"a = {\n  hat = {\n    band = \"red\"\n  }\n}\n",
			[]string{"dog"},
			"a = {\n  hat = {\n    band = \"red\"\n  }\n  dog = \"bowler\"\n}\n",
		},
		{
			"nested multi-line value with the brace on its line",
			"a = {\n  hat = {\n    band = \"red\"\n  } }\n",
			[]string{"dog"},
			"a = {\n  hat = {\n    band = \"red\"\n  }\n  dog = \"bowler\"\n}\n",
		},
		{
			"inside a block",
			"b {\n  a = {\n    hat = \"derby\"\n  }\n}\n",
			[]string{"dog"},
			"b {\n  a = {\n    hat = \"derby\"\n    dog = \"bowler\"\n  }\n}\n",
		},
		{
			"key from an expression",
			"a = { (hat) = \"derby\" }\n",
			[]string{"hat"},
			"a = { (hat) = \"derby\", hat = \"bowler\" }\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			body := f.Body()
			if body.GetAttribute("a") == nil {
				body = body.Blocks()[0].Body()
			}
			object := body.GetAttribute("a").Expr().AsObjectConsExpr()

			for _, key := range test.keys {
				gotKey, gotValue := object.SetItemRaw(key, bowler)
				item := object.ItemFor(key)
				if item == nil || item.KeyExpr() != gotKey || item.ValueExpr() != gotValue {
					t.Fatalf("SetItemRaw(%q) returned other than the item ItemFor finds", key)
				}
			}

			got := f.Bytes()
			if string(got) != test.want {
				t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, test.want)
			}
			if formatted := Format(got); string(formatted) != string(got) {
				t.Errorf("result is not in canonical form\ngot:\n%s\nformatted:\n%s", got, formatted)
			}
			if _, diags := hclsyntax.ParseConfig(got, "", hcl.InitialPos); diags.HasErrors() {
				t.Errorf("result is not valid HCL: %s", diags)
			}
		})
	}
}

func TestObjectConsExprSetItemHelpers(t *testing.T) {
	f, diags := ParseConfig([]byte("a = { hat = \"derby\" }\n"), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics: %s", diags)
	}
	expr := f.Body().GetAttribute("a").Expr()
	object := expr.AsObjectConsExpr()

	object.SetItemValue("hat", cty.StringVal("bowler"))
	object.SetItemValue("size", cty.NumberIntVal(7))
	object.SetItemTraversal("owner", hcl.Traversal{
		hcl.TraverseRoot{Name: "old"},
		hcl.TraverseAttr{Name: "name"},
	})
	object.SetItem("cat", NewExpressionLiteral(cty.True))

	want := "a = { hat = \"bowler\", size = 7, owner = old.name, cat = true }\n"
	if got := string(f.Bytes()); got != want {
		t.Fatalf("wrong result\ngot:  %s\nwant: %s", got, want)
	}

	// A traversal that was set is a variable of the object.
	expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
	want = "a = { hat = \"bowler\", size = 7, owner = new.name, cat = true }\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result after renaming\ngot:  %s\nwant: %s", got, want)
	}
}

func TestObjectConsExprSetItemPanics(t *testing.T) {
	src := "a = { hat = \"derby\" }\nb = \"bowler\"\n"

	tests := []struct {
		name   string
		attach func(f *File, object *ObjectConsExpr) *Expression
	}{
		{
			"expression of an attribute",
			func(f *File, object *ObjectConsExpr) *Expression {
				return f.Body().GetAttribute("b").Expr()
			},
		},
		{
			"value of an item",
			func(f *File, object *ObjectConsExpr) *Expression {
				return object.ValueExprFor("hat").Expr()
			},
		},
		{
			"value set by an earlier SetItem",
			func(f *File, object *ObjectConsExpr) *Expression {
				expr := NewExpressionLiteral(cty.True)
				object.SetItem("dog", expr)
				return expr
			},
		},
		{
			"value replaced by an earlier SetItem",
			func(f *File, object *ObjectConsExpr) *Expression {
				expr := NewExpressionLiteral(cty.True)
				object.SetItem("hat", expr)
				return expr
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()
			expr := test.attach(f, object)

			defer func() {
				if recover() == nil {
					t.Errorf("SetItem did not panic")
				}
			}()
			object.SetItem("cat", expr)
		})
	}
}

func TestObjectConsExprSetItemReplacedValue(t *testing.T) {
	// The value an earlier SetItem replaced is no longer in the tree, so it
	// may be set again.
	f, diags := ParseConfig([]byte("a = { hat = \"derby\" }\n"), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics: %s", diags)
	}
	object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

	old := object.ValueExprFor("hat").Expr()
	object.SetItemValue("hat", cty.StringVal("bowler"))
	object.SetItem("cat", old)

	want := "a = { hat = \"bowler\", cat = \"derby\" }\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result\ngot:  %s\nwant: %s", got, want)
	}
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

	file, diags := hclwrite.ParseConfig(config, "<inline>", hcl.Pos{})
	if diags.HasErrors() {
		fmt.Printf("errors: %s", diags)
		return
	}
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

#### x/function_call_expr.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// FunctionCallExpr represents a function call expression such as
// join(" ", list). Use Expression.AsFunctionCallExpr to get one from a parsed
// expression.
//
// Each argument is an Expression of its own. The function name, the
// parentheses, the commas and a final ellipsis stay as tokens of the call.
type FunctionCallExpr struct {
	inTree

	args nodeSet
}

// Name returns the name of the called function. A namespaced name, such as
// provider::time::rfc3339_parse, is returned whole.
func (call *FunctionCallExpr) Name() string {
	var b strings.Builder
	for _, tok := range call.BuildTokens(nil) {
		if tok.Type == hclsyntax.TokenOParen {
			break
		}
		if tok.Type == hclsyntax.TokenIdent || tok.Type == hclsyntax.TokenDoubleColon {
			b.Write(tok.Bytes)
		}
	}
	return b.String()
}

// variables returns the absolute traversals in the arguments of the call, in
// source order.
//
func (call *FunctionCallExpr) variables() []*Traversal {
	var ret []*Traversal
	for _, n := range call.args.List() {
		ret = append(ret, n.content.(*Expression).Variables()...)
	}
	return ret
}

// parseFunctionCallExpr parses a function call expression, defined as:
//
//	FunctionCall = Identifier ("::" Identifier)* "(" Arguments ")";
//	Arguments = (
//	    () ||
//	    (Expression ("," Expression)* ("," | "...")?)
//	);
//
// Each argument becomes an Expression of its own. The tokens between them,
// and any comments, stay as unstructured tokens of the call.
//
func parseFunctionCallExpr(nativeExpr *hclsyntax.FunctionCallExpr, from inputTokens) *node {
	call := &FunctionCallExpr{
		inTree: newInTree(),
		args:   newNodeSet(),
	}

	var before, argTokens inputTokens
	for _, nativeArg := range nativeExpr.Args {
		before, argTokens, from = from.Partition(nativeArg.Range())
		call.children.AppendUnstructuredTokens(before.Tokens())
		arg := parseExpression(nativeArg, argTokens)
		call.children.AppendNode(arg)
		call.args.Add(arg)
	}
	call.children.AppendUnstructuredTokens(from.Tokens())

	wrapExpr := newExpression()
	wrapExpr.wrap(call)
	return newNode(wrapExpr)
}

```

#### x/function_call_expr_test.go

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
		want string // the name, or "" when the expression is not a call
	}{
		{`a = join(" ", ["fedora", "fez"])`, "join"},
		{`a = f()`, "f"},
		{`a = f(x...)`, "f"},
		{`a = provider::time::rfc3339_parse("2026-11-05T00:00:00Z")`, "provider::time::rfc3339_parse"},
		{"a = f(\n  1,\n  2,\n)", "f"},
		{`a = f(1)[0]`, ""},
		{`a = (f(1))`, ""},
		{`a = ["fedora"]`, ""},
		{`a = "fez"`, ""},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			src := test.src + "\n"
			f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			if got := string(f.Bytes()); got != src {
				t.Errorf("source does not round-trip\ngot:  %q\nwant: %q", got, src)
			}

			call := f.Body().GetAttribute("a").Expr().AsFunctionCallExpr()
			if test.want == "" {
				if call != nil {
					t.Fatalf("got a call named %q, but want none", call.Name())
				}
				return
			}
			if call == nil {
				t.Fatalf("got no call")
			}
			if got := call.Name(); got != test.want {
				t.Errorf("wrong name\ngot:  %q\nwant: %q", got, test.want)
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
	n.setContentHolder(nil)
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
	n.setContentHolder(nil)

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
	nn.setContentHolder(nn)
	return nn
}

func (n *node) assertUnattached() {
	if n.list != nil {
		panic(fmt.Sprintf("attempt to attach already-attached node %#v", n))
	}
}

// setContentHolder records holder as the node that holds the content of n,
// for content that keeps such a record. A nil holder records that the
// content is no longer in a list.
func (n *node) setContentHolder(holder *node) {
	if c, ok := n.content.(interface{ setHolder(*node) }); ok {
		c.setHolder(holder)
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
	for n := ns.first; n != nil; n = n.after {
		n.setContentHolder(nil)
	}
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
	n.setContentHolder(n)
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
	n.setContentHolder(n)
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
	// holder is the node that holds this content while the content is in a
	// list of nodes, and nil otherwise. The list operations keep it.
	holder *node
}

func newInTree() inTree {
	return inTree{
		children: &nodes{},
	}
}

// assertUnattached panics if the content is already part of a syntax tree.
func (it *inTree) assertUnattached() {
	if it.parent != nil {
		panic(fmt.Sprintf("node is already attached to %T", it.parent.content))
	}
	if it.holder != nil {
		panic("node is already attached to a syntax tree")
	}
}

func (it *inTree) setHolder(holder *node) {
	it.holder = holder
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

// inputTokens is a run of input tokens, both as hclsyntax lexed them and as
// hclwrite keeps them, which the parse functions partition.
//
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

// parseExpression parses the given expression from the given tokens, which
// hold its source.
//
func parseExpression(nativeExpr hclsyntax.Expression, from inputTokens) *node {
	switch tNativeExpr := nativeExpr.(type) {

	// Object-construct expression
	case *hclsyntax.ObjectConsExpr:
		return parseObjectConsExpr(tNativeExpr, from)

	case *hclsyntax.ObjectConsKeyExpr:
		return parseObjectConsKeyExpr(tNativeExpr, from)

	case *hclsyntax.TemplateExpr:
		return parseTemplateExpr(tNativeExpr, from)

	case *hclsyntax.FunctionCallExpr:
		return parseFunctionCallExpr(tNativeExpr, from)

	case *hclsyntax.TupleConsExpr:
		return parseTupleConsExpr(tNativeExpr, from)

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
	// hclsyntax treats a key that reads as a keyword as a literal name, and
	// reports no variables for it. We follow the same rule.
	wrapExpr := newObjectConsKeyExpr(expr, hcl.ExprAsKeyword(nativeExpr.Wrapped) != "")

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

#### x/parser_test.go

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
			`a = join(" ", ["fedora", "fez"])` + "\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{Type: "comments"},
							{Type: "identifier", Val: "a"},
							{Type: "Tokens", Val: " ="},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
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
												Children: []TestTreeNode{
													{
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
													},
												},
											},
											{Type: "Tokens", Val: ")"},
										},
									},
								},
							},
							{Type: "comments"},
							{Type: "Tokens", Val: "\n"},
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
							{Type: "comments"},
							{Type: "identifier", Val: "a"},
							{Type: "Tokens", Val: " ="},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type:     "FunctionCallExpr",
										Children: []TestTreeNode{{Type: "Tokens", Val: " f()"}},
									},
								},
							},
							{Type: "comments"},
							{Type: "Tokens", Val: "\n"},
						},
					},
				},
			},
		},
		{
			"a = f(old.x, { k = [] }, rest...)\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{Type: "comments"},
							{Type: "identifier", Val: "a"},
							{Type: "Tokens", Val: " ="},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "FunctionCallExpr",
										Children: []TestTreeNode{
											{Type: "Tokens", Val: " f("},
											{
												Type: "Expression",
												Children: []TestTreeNode{
													{
														Type: "Traversal",
														Children: []TestTreeNode{
															{
																Type:     "TraverseName",
																Children: []TestTreeNode{{Type: "identifier", Val: "old"}},
															},
															{
																Type: "TraverseName",
																Children: []TestTreeNode{
																	{Type: "Tokens", Val: "."},
																	{Type: "identifier", Val: "x"},
																},
															},
														},
													},
												},
											},
											{Type: "Tokens", Val: ","},
											{
												Type: "Expression",
												Children: []TestTreeNode{
													{
														Type: "ObjectConsExpr",
														Children: []TestTreeNode{
															{Type: "Tokens", Val: " {"},
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
																								Type:     "TraverseName",
																								Children: []TestTreeNode{{Type: "identifier", Val: " k"}},
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
																				Type: "Expression",
																				Children: []TestTreeNode{
																					{
																						Type:     "TupleConsExpr",
																						Children: []TestTreeNode{{Type: "Tokens", Val: " []"}},
																					},
																				},
																			},
																		},
																	},
																},
															},
															{Type: "Tokens", Val: " }"},
														},
													},
												},
											},
											{Type: "Tokens", Val: ","},
											{
												Type: "Expression",
												Children: []TestTreeNode{
													{
														Type: "Traversal",
														Children: []TestTreeNode{
															{
																Type:     "TraverseName",
																Children: []TestTreeNode{{Type: "identifier", Val: " rest"}},
															},
														},
													},
												},
											},
											{Type: "Tokens", Val: "...)"},
										},
									},
								},
							},
							{Type: "comments"},
							{Type: "Tokens", Val: "\n"},
						},
					},
				},
			},
		},
		{
			"a = [\n  \"fedora\", # hat\n  f(),\n]\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{Type: "comments"},
							{Type: "identifier", Val: "a"},
							{Type: "Tokens", Val: " ="},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "TupleConsExpr",
										Children: []TestTreeNode{
											{Type: "Tokens", Val: " [\n"},
											{
												Type:     "Expression",
												Children: []TestTreeNode{{Type: "quoted", Val: `  "fedora"`}},
											},
											{Type: "Tokens", Val: ", # hat\n"},
											{
												Type: "Expression",
												Children: []TestTreeNode{
													{
														Type:     "FunctionCallExpr",
														Children: []TestTreeNode{{Type: "Tokens", Val: "  f()"}},
													},
												},
											},
											{Type: "Tokens", Val: ",\n]"},
										},
									},
								},
							},
							{Type: "comments"},
							{Type: "Tokens", Val: "\n"},
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

#### x/tuple_cons_expr.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// TupleConsExpr represents a tuple-construct expression such as
// ["fedora", "fez"]. Use Expression.AsTupleConsExpr to get one from a parsed
// expression.
//
// Each element is an Expression of its own. The brackets and the commas stay
// as tokens of the tuple.
type TupleConsExpr struct {
	inTree

	elems nodeSet
}

// variables returns the absolute traversals in the elements of the tuple, in
// source order.
//
func (tuple *TupleConsExpr) variables() []*Traversal {
	var ret []*Traversal
	for _, n := range tuple.elems.List() {
		ret = append(ret, n.content.(*Expression).Variables()...)
	}
	return ret
}

// parseTupleConsExpr parses a tuple-construct expression, defined as:
//
//	tuple = "[" (
//	    (Expression (("," | Newline) Expression)* ","?)?
//	) "]";
//
// Each element becomes an Expression of its own. The tokens between them,
// and any comments, stay as unstructured tokens of the tuple.
//
func parseTupleConsExpr(nativeExpr *hclsyntax.TupleConsExpr, from inputTokens) *node {
	tuple := &TupleConsExpr{
		inTree: newInTree(),
		elems:  newNodeSet(),
	}

	var before, elemTokens inputTokens
	for _, nativeElem := range nativeExpr.Exprs {
		before, elemTokens, from = from.Partition(nativeElem.Range())
		tuple.children.AppendUnstructuredTokens(before.Tokens())
		elem := parseExpression(nativeElem, elemTokens)
		tuple.children.AppendNode(elem)
		tuple.elems.Add(elem)
	}
	tuple.children.AppendUnstructuredTokens(from.Tokens())

	wrapExpr := newExpression()
	wrapExpr.wrap(tuple)
	return newNode(wrapExpr)
}

```

#### x/tuple_cons_expr_test.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
)

func TestTupleConsExprAsTupleConsExpr(t *testing.T) {
	tests := []struct {
		src  string
		want bool // whether the expression is a tuple
	}{
		{`a = ["fedora", "fez"]`, true},
		{`a = []`, true},
		{"a = [\n  \"fedora\", # hat\n  [\"fez\"],\n]", true},
		{`a = [for hat in hats : hat]`, false},
		{`a = ["fedora"][0]`, false},
		{`a = join(" ", ["fedora"])`, false},
		{`a = "fez"`, false},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			src := test.src + "\n"
			f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			if got := string(f.Bytes()); got != src {
				t.Errorf("source does not round-trip\ngot:  %q\nwant: %q", got, src)
			}

			tuple := f.Body().GetAttribute("a").Expr().AsTupleConsExpr()
			if got := tuple != nil; got != test.want {
				t.Errorf("got a tuple: %t, want %t", got, test.want)
			}
		})
	}
}

```

#### y/ast_attribute.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

type Attribute struct {
	inTree

	leadComments *node
	name         *node
	expr         *node
	lineComments *node
}

func newAttribute() *Attribute {
	return &Attribute{
		inTree: newInTree(),
	}
}

func (a *Attribute) init(name string, expr *Expression) {
	expr.assertUnattached()

	nameTok := newIdentToken(name)
	nameObj := newIdentifier(nameTok)
	a.leadComments = a.children.Append(newComments(nil))
	a.name = a.children.Append(nameObj)
	a.children.AppendUnstructuredTokens(Tokens{
		{
			Type:  hclsyntax.TokenEqual,
			Bytes: []byte{'='},
		},
	})
	a.expr = a.children.Append(expr)
	a.lineComments = a.children.Append(newComments(nil))
	a.children.AppendUnstructuredTokens(Tokens{
		{
			Type:  hclsyntax.TokenNewline,
			Bytes: []byte{'\n'},
		},
	})
}

func (a *Attribute) LeadComments() Tokens {
	return a.leadComments.content.BuildTokens(nil)
}

func (a *Attribute) LineComments() Tokens {
	return a.lineComments.content.BuildTokens(nil)
}

func (a *Attribute) Expr() *Expression {
	return a.expr.content.(*Expression)
}

// setName updates the name of the attribute.
func (a *Attribute) setName(name string) {
	nameObj := newIdentifier(newIdentToken(name))
	a.name = a.name.ReplaceWith(nameObj)
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
// expression, in source order.
//
// For an object-construct expression, this includes the traversals in the
// keys and values of its items. A naked identifier key, such as "foo" in
// { foo = 1 }, is a literal name and not a traversal. For a function call or
// a tuple-construct expression, this includes the traversals in its
// arguments or elements.
func (e *Expression) Variables() []*Traversal {
	nodes := e.absTraversals.List()
	ret := make([]*Traversal, len(nodes))
	for i, node := range nodes {
		ret[i] = node.content.(*Traversal)
	}
	if e.wrapped != nil {
		if nested, ok := e.wrapped.content.(interface{ variables() []*Traversal }); ok {
			ret = append(ret, nested.variables()...)
		}
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
	for _, traversal := range e.Variables() {
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

#### y/ast_function_call_expr.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

// FunctionCallExpr represents a function call expression such as
// join(" ", list). Use Expression.AsFunctionCallExpr to get one.
//
// Each argument is its own Expression. The function name, the parentheses,
// the commas and an expansion symbol after the last argument are tokens of
// the FunctionCallExpr itself.
type FunctionCallExpr struct {
	inTree

	// name holds the tokens of the function name, which can be namespaced
	// such as provider::example::func.
	name *node
	args nodeSet
}

func newFunctionCallExpr() *FunctionCallExpr {
	return &FunctionCallExpr{
		inTree: newInTree(),
		args:   newNodeSet(),
	}
}

// Name returns the name of the called function.
func (call *FunctionCallExpr) Name() string {
	var name []byte
	for _, tok := range call.name.content.(Tokens) {
		name = append(name, tok.Bytes...)
	}
	return string(name)
}

// variables returns the traversals in the arguments, in source order.
func (call *FunctionCallExpr) variables() []*Traversal {
	return expressionsVariables(call.args)
}

// expressionsVariables returns the traversals in each of the given
// expression nodes, in source order.
func expressionsVariables(exprs nodeSet) []*Traversal {
	var ret []*Traversal
	for _, n := range exprs.List() {
		ret = append(ret, n.content.(*Expression).Variables()...)
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
	"slices"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// parseAttrExprForTest parses src and returns the expression of its
// attribute "a".
func parseAttrExprForTest(t *testing.T, src string) (*File, *Expression) {
	t.Helper()
	f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics: %s", diags)
	}
	return f, f.Body().GetAttribute("a").Expr()
}

func TestFunctionCallExprName(t *testing.T) {
	tests := map[string]string{
		"a = join(\" \", [\"fedora\", \"fez\"])\n": "join",
		"a = f()\n":               "f",
		"a = provider::p::f(1)\n": "provider::p::f",
		"a = f(\n  1,\n  2,\n)\n": "f",
	}
	for src, want := range tests {
		t.Run(src, func(t *testing.T) {
			_, expr := parseAttrExprForTest(t, src)
			call := expr.AsFunctionCallExpr()
			if call == nil {
				t.Fatalf("attribute a is not a function call expression")
			}
			if got := call.Name(); got != want {
				t.Errorf("wrong name\ngot:  %q\nwant: %q", got, want)
			}
		})
	}
}

func TestExpressionAsFunctionCallAndTupleCons(t *testing.T) {
	tests := map[string]struct {
		src       string
		wantCall  bool
		wantTuple bool
	}{
		"function call":          {src: "a = f(1)\n", wantCall: true},
		"tuple":                  {src: "a = [1]\n", wantTuple: true},
		"empty tuple":            {src: "a = []\n", wantTuple: true},
		"object":                 {src: "a = { b = f(1) }\n"},
		"traversal":              {src: "a = b.c\n"},
		"call in binary op":      {src: "a = f(1) + 1\n"},
		"tuple in parentheses":   {src: "a = ([1])\n"},
		"for expression":         {src: "a = [for x in y : x]\n"},
		"index on function call": {src: "a = f(1)[0]\n"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, expr := parseAttrExprForTest(t, test.src)
			if got := expr.AsFunctionCallExpr() != nil; got != test.wantCall {
				t.Errorf("AsFunctionCallExpr() != nil is %t, want %t", got, test.wantCall)
			}
			if got := expr.AsTupleConsExpr() != nil; got != test.wantTuple {
				t.Errorf("AsTupleConsExpr() != nil is %t, want %t", got, test.wantTuple)
			}
		})
	}
}

func TestFunctionCallAndTupleConsVariables(t *testing.T) {
	tests := map[string]struct {
		src string
		// want is the source of each traversal that Variables returns,
		// in order.
		want []string
		// renamed is the source after renaming the prefix "old" to "new".
		renamed string
	}{
		"function arguments in source order": {
			src:     "a = f(old.x, other.y, old.z)\n",
			want:    []string{"old.x", "other.y", "old.z"},
			renamed: "a = f(new.x, other.y, new.z)\n",
		},
		"tuple elements in source order": {
			src:     "a = [old.x, other.y, old.z]\n",
			want:    []string{"old.x", "other.y", "old.z"},
			renamed: "a = [new.x, other.y, new.z]\n",
		},
		"nested calls and tuples": {
			src:     "a = f([old.x, g(old.y)], old.z)\n",
			want:    []string{"old.x", "old.y", "old.z"},
			renamed: "a = f([new.x, g(new.y)], new.z)\n",
		},
		"object in function argument": {
			src:     "a = f({ k = old.x }, old.y)\n",
			want:    []string{"old.x", "old.y"},
			renamed: "a = f({ k = new.x }, new.y)\n",
		},
		"function call in object item": {
			src:     "a = { k = f(old.x) }\n",
			want:    []string{"old.x"},
			renamed: "a = { k = f(new.x) }\n",
		},
		"tuple in object item": {
			src:     "a = { k = [old.x] }\n",
			want:    []string{"old.x"},
			renamed: "a = { k = [new.x] }\n",
		},
		"expanded final argument": {
			src:     "a = f(old.x, old.y...)\n",
			want:    []string{"old.x", "old.y"},
			renamed: "a = f(new.x, new.y...)\n",
		},
		"multi-line tuple with comments": {
			src:     "a = [\n  old.x, # c\n  # d\n  old.y,\n]\n",
			want:    []string{"old.x", "old.y"},
			renamed: "a = [\n  new.x, # c\n  # d\n  new.y,\n]\n",
		},
		"argument with operators": {
			src:     "a = f(old.x + old.y)\n",
			want:    []string{"old.x", "old.y"},
			renamed: "a = f(new.x + new.y)\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// Variables must agree with hclsyntax on the same source.
			nf, diags := hclsyntax.ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected hclsyntax diagnostics: %s", diags)
			}
			nativeAttr := nf.Body.(*hclsyntax.Body).Attributes["a"]
			if nativeCount := len(nativeAttr.Expr.Variables()); nativeCount != len(test.want) {
				t.Fatalf("test case is wrong: hclsyntax reports %d traversals, but want has %d", nativeCount, len(test.want))
			}

			f, expr := parseAttrExprForTest(t, test.src)

			var got []string
			for _, traversal := range expr.Variables() {
				got = append(got, strings.TrimSpace(string(traversal.BuildTokens(nil).Bytes())))
			}
			if !slices.Equal(got, test.want) {
				t.Errorf("wrong variables\ngot:  %q\nwant: %q", got, test.want)
			}

			expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
			if got := string(f.Bytes()); got != test.renamed {
				t.Errorf("wrong result after rename\ngot:  %q\nwant: %q", got, test.renamed)
			}
		})
	}
}

func TestFunctionCallAndTupleConsRoundTrip(t *testing.T) {
	tests := []string{
		"a = join(\" \", [\"fedora\", \"fez\"])\n",
		"a = f()\n",
		"a = f( )\n",
		"a = [ ]\n",
		"a = provider::p::f(1)\n",
		"a = f(1, 2,)\n",
		"a = f(x...)\n",
		"a = f(\n  1, # one\n  /* two */ 2,\n)\n",
		"a = [\n  1,\n  # comment\n  [2, [3]],\n  { b = f(4) },\n]\n",
		"a = f(<<EOT\nhello\nEOT\n)\n",
		"a = [for x in y : f(x)]\n",
		"block \"label\" {\n  a = f([1], { b = [2] })\n}\n",
	}
	for _, src := range tests {
		t.Run(src, func(t *testing.T) {
			f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			// The tree must hold every token unchanged.
			if got := string(f.BuildTokens(nil).Bytes()); got != src {
				t.Errorf("wrong tokens\ngot:  %q\nwant: %q", got, src)
			}
			// File.Bytes formats its result, so it can only reproduce
			// source that is already formatted.
			if string(Format([]byte(src))) != src {
				return
			}
			if got := string(f.Bytes()); got != src {
				t.Errorf("wrong result\ngot:  %q\nwant: %q", got, src)
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
	"bytes"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

// ObjectConsExpr is an object-construct expression, such as
// { name = "value" }. Use Expression.AsObjectConsExpr to get one.
//
// Its items can be looked up by key and changed with the SetItem methods,
// in the same way as the attributes of a Body.
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
// there is no such item.
//
// Only a key written as an identifier, such as hat, or as a quoted string
// literal, such as "hat", can match. A quoted key is compared by its value
// after escape sequences are decoded. A key that is computed from another
// expression, such as (hat), hat.band or "${hat}", never matches.
func (o *ObjectConsExpr) ItemFor(key string) *ObjectConsItem {
	for _, item := range o.Items() {
		if name, ok := item.KeyExpr().literalName(); ok && name == key {
			return item
		}
	}
	return nil
}

// ValueExprFor returns the value of the item that ItemFor finds for the
// given key, or nil if there is no such item.
func (o *ObjectConsExpr) ValueExprFor(key string) *ObjectConsValue {
	item := o.ItemFor(key)
	if item == nil {
		return nil
	}
	return item.ValueExpr()
}

// SetItem either replaces the value of the item that ItemFor finds for the
// given key, or adds a new item with that key to the end of the object.
//
// The key of an existing item is left as it is written. A new key is written
// as an identifier if it is a valid identifier, or as a quoted string literal
// otherwise.
//
// The given expression must not already be part of a syntax tree, or this
// method panics.
//
// The return values are the key and the value of the item that was either
// modified in-place or created.
func (o *ObjectConsExpr) SetItem(key string, expr *Expression) (*ObjectConsKeyExpr, *ObjectConsValue) {
	expr.assertUnattached()

	if item := o.ItemFor(key); item != nil {
		value := item.ValueExpr()
		value.expr = value.expr.ReplaceWith(expr)
		return item.KeyExpr(), value
	}

	item := newObjectConsItem()
	item.init(key, expr)
	o.appendItem(item)
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

// SetItemValue is like SetItem, but sets the value to a literal that
// represents the given value. It is based on NewExpressionLiteral.
func (o *ObjectConsExpr) SetItemValue(key string, val cty.Value) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionLiteral(val))
}

// SetItemTraversal is like SetItem, but sets the value to the given
// traversal, which must be absolute. It is based on NewExpressionAbsTraversal.
func (o *ObjectConsExpr) SetItemTraversal(key string, traversal hcl.Traversal) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionAbsTraversal(traversal))
}

// appendItem adds the given new item to the end of the object, laid out
// as described in appendItemTokens.
func (o *ObjectConsExpr) appendItem(item *ObjectConsItem) {
	multiline := o.isMultiline()

	// Collect the tokens after the last item, or after the opening brace
	// if there are no items, so that they can be split around the new item.
	var rest Tokens
	var start *node
	if items := o.items.List(); len(items) > 0 {
		start = items[len(items)-1].after
	} else {
		start = o.children.first
	}
	for n := start; n != nil; {
		next := n.after
		rest = append(rest, n.content.(Tokens)...)
		n.Detach()
		n = next
	}
	if len(o.items) == 0 {
		open := 0
		for open < len(rest) && rest[open].Type != hclsyntax.TokenOBrace {
			open++
		}
		o.children.AppendUnstructuredTokens(rest[:open+1])
		rest = rest[open+1:]
	}

	before, after := splitObjectTail(rest, multiline, len(o.items) > 0)
	o.children.AppendUnstructuredTokens(before)
	o.items.Add(o.children.Append(item))
	o.children.AppendUnstructuredTokens(after)
}

// splitObjectTail splits the tokens that follow the last item of an object,
// up to and including the closing brace, into the tokens to put before and
// after a new item that is added at the end.
//
// A comma or a line comment that follows the last item on its line stays
// before the new item. Everything else, including comments on the following
// lines and an inline comment directly before the brace, stays with the
// closing brace after the new item. Separators are added so that the new
// item goes on the same line in a single-line object, or on its own line in
// a multi-line object.
func splitObjectTail(tail Tokens, multiline, hasItems bool) (before, after Tokens) {
	split := 0
	for i, tok := range tail {
		if multiline && isLineEnd(tok) {
			split = i + 1
			break
		}
		if tok.Type == hclsyntax.TokenComma {
			split = i + 1
			if !multiline {
				break
			}
		}
	}
	before = append(Tokens{}, tail[:split]...)
	after = tail[split:]

	switch {
	case multiline:
		if len(before) == 0 || !isLineEnd(before[len(before)-1]) {
			before = append(before, newNewlineToken())
		}
		after = append(Tokens{newNewlineToken()}, after...)
	case hasItems && split == 0:
		before = append(before, &Token{
			Type:  hclsyntax.TokenComma,
			Bytes: []byte{','},
		})
	}
	return before, after
}

// isMultiline reports whether the object's own tokens contain a line break.
// Line breaks inside the keys and values of its items are not counted.
func (o *ObjectConsExpr) isMultiline() bool {
	for n := o.children.first; n != nil; n = n.after {
		switch c := n.content.(type) {
		case Tokens:
			if hasLineEnd(c) {
				return true
			}
		case *ObjectConsItem:
			for in := c.children.first; in != nil; in = in.after {
				if tokens, ok := in.content.(Tokens); ok && hasLineEnd(tokens) {
					return true
				}
			}
		}
	}
	return false
}

func hasLineEnd(tokens Tokens) bool {
	for _, tok := range tokens {
		if isLineEnd(tok) {
			return true
		}
	}
	return false
}

// isLineEnd reports whether the token ends a line. A line comment includes
// the newline that ends it.
func isLineEnd(tok *Token) bool {
	switch tok.Type {
	case hclsyntax.TokenNewline:
		return true
	case hclsyntax.TokenComment:
		return bytes.HasSuffix(tok.Bytes, []byte{'\n'})
	}
	return false
}

func newNewlineToken() *Token {
	return &Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte{'\n'},
	}
}

// variables returns the traversals in the keys and values of the items,
// in source order.
func (o *ObjectConsExpr) variables() []*Traversal {
	var ret []*Traversal
	for _, item := range o.Items() {
		ret = append(ret, item.KeyExpr().variables()...)
		ret = append(ret, item.ValueExpr().Expr().Variables()...)
	}
	return ret
}

// ObjectConsItem is a single key = value item of an ObjectConsExpr.
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

func (item *ObjectConsItem) init(key string, expr *Expression) {
	item.key = item.children.Append(newObjectConsKeyExprFor(key))
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

// KeyExpr returns the key of the item.
func (item *ObjectConsItem) KeyExpr() *ObjectConsKeyExpr {
	return item.key.content.(*ObjectConsKeyExpr)
}

// ValueExpr returns the value of the item.
func (item *ObjectConsItem) ValueExpr() *ObjectConsValue {
	return item.value.content.(*ObjectConsValue)
}

// ObjectConsKeyExpr is the key of an ObjectConsItem. It is either a literal
// name, written as an identifier or a quoted string, or an expression whose
// result is the key.
type ObjectConsKeyExpr struct {
	inTree

	wrapped *node
}

func newObjectConsKeyExpr(wrapped *node) *ObjectConsKeyExpr {
	expr := &ObjectConsKeyExpr{
		inTree:  newInTree(),
		wrapped: wrapped,
	}
	expr.children.AppendNode(wrapped)

	return expr
}

// newObjectConsKeyExprFor returns a key that evaluates to the given string.
// It is an identifier if the string is a valid identifier, or a quoted string
// literal otherwise.
func newObjectConsKeyExprFor(key string) *ObjectConsKeyExpr {
	expr := newExpression()
	if hclsyntax.ValidIdentifier(key) {
		expr.children.AppendUnstructuredTokens(TokensForIdentifier(key))
	} else {
		expr.wrapped = expr.children.Append(newQuoted(TokensForValue(cty.StringVal(key))))
	}
	return newObjectConsKeyExpr(newNode(expr))
}

// literalName returns the name that the key stands for if it is an
// identifier or a quoted string literal. Otherwise, ok is false.
//
// The rules match hclsyntax.ObjectConsKeyExpr: the keywords null, true and
// false are names too, and a key in parentheses is always an expression.
func (k *ObjectConsKeyExpr) literalName() (name string, ok bool) {
	src := k.BuildTokens(nil).Bytes()
	expr, diags := hclsyntax.ParseExpression(src, "", hcl.InitialPos)
	if diags.HasErrors() {
		return "", false
	}

	switch expr := expr.(type) {
	case *hclsyntax.ScopeTraversalExpr, *hclsyntax.LiteralValueExpr:
		name := hcl.ExprAsKeyword(expr)
		return name, name != ""
	case *hclsyntax.TemplateExpr:
		if !expr.IsStringLiteral() {
			return "", false
		}
		val, diags := expr.Value(nil)
		if diags.HasErrors() {
			return "", false
		}
		return val.AsString(), true
	default:
		return "", false
	}
}

func (k *ObjectConsKeyExpr) variables() []*Traversal {
	if k.wrapped == nil {
		return nil
	}
	return k.wrapped.content.(*Expression).Variables()
}

// ObjectConsValue is the value of an ObjectConsItem.
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
	"slices"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

func TestObjectConsExprVariables(t *testing.T) {
	tests := map[string]struct {
		src string
		// want is the source of each traversal that Variables returns,
		// in order.
		want []string
		// renamed is the source after renaming the prefix "old" to "new".
		renamed string
	}{
		"value": {
			src:     "a = { k = old.x }\n",
			want:    []string{"old.x"},
			renamed: "a = { k = new.x }\n",
		},
		"multiple items in source order": {
			src:     "a = { k = old.x, j = other.y, (old.z) = old.w }\n",
			want:    []string{"old.x", "other.y", "old.z", "old.w"},
			renamed: "a = { k = new.x, j = other.y, (new.z) = new.w }\n",
		},
		"nested object": {
			src:     "a = { k = { j = old.x } }\n",
			want:    []string{"old.x"},
			renamed: "a = { k = { j = new.x } }\n",
		},
		"tuple value": {
			src:     "a = { k = [old.x] }\n",
			want:    []string{"old.x"},
			renamed: "a = { k = [new.x] }\n",
		},
		"function call value": {
			src:     "a = { k = f(old.x) }\n",
			want:    []string{"old.x"},
			renamed: "a = { k = f(new.x) }\n",
		},
		"object in tuple value": {
			src:     "a = { k = [{ j = old.x }] }\n",
			want:    []string{"old.x"},
			renamed: "a = { k = [{ j = new.x }] }\n",
		},
		"naked identifier key": {
			src:     "a = { old = 1 }\n",
			want:    nil,
			renamed: "a = { old = 1 }\n",
		},
		"naked identifier key in nested object": {
			src:     "a = { k = { old = old.x } }\n",
			want:    []string{"old.x"},
			renamed: "a = { k = { old = new.x } }\n",
		},
		"parenthesized key": {
			src:     "a = { (old.x) = 1 }\n",
			want:    []string{"old.x"},
			renamed: "a = { (new.x) = 1 }\n",
		},
		"parenthesized single name key": {
			src:     "a = { (old) = 1 }\n",
			want:    []string{"old"},
			renamed: "a = { (new) = 1 }\n",
		},
		"naked traversal key": {
			// hclsyntax treats only a single name as a literal key.
			src:     "a = { old.x = 1 }\n",
			want:    []string{"old.x"},
			renamed: "a = { new.x = 1 }\n",
		},
		"quoted key": {
			src:     "a = { \"k\" = old.x }\n",
			want:    []string{"old.x"},
			renamed: "a = { \"k\" = new.x }\n",
		},
		"empty object": {
			src:     "a = {}\n",
			want:    nil,
			renamed: "a = {}\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// Variables must agree with hclsyntax on the same source.
			nf, diags := hclsyntax.ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected hclsyntax diagnostics: %s", diags)
			}
			nativeAttr := nf.Body.(*hclsyntax.Body).Attributes["a"]
			nativeCount := len(nativeAttr.Expr.Variables())
			if nativeCount != len(test.want) {
				t.Fatalf("test case is wrong: hclsyntax reports %d traversals, but want has %d", nativeCount, len(test.want))
			}

			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			expr := f.Body().GetAttribute("a").Expr()

			var got []string
			for _, traversal := range expr.Variables() {
				got = append(got, strings.TrimSpace(string(traversal.BuildTokens(nil).Bytes())))
			}
			if !slices.Equal(got, test.want) {
				t.Errorf("wrong variables\ngot:  %q\nwant: %q", got, test.want)
			}

			expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
			if got := string(f.Bytes()); got != test.renamed {
				t.Errorf("wrong result after rename\ngot:  %q\nwant: %q", got, test.renamed)
			}
		})
	}
}

// parseObjectConsExprForTest parses src and returns the object-construct expression
// of its attribute "a".
func parseObjectConsExprForTest(t *testing.T, src string) (*File, *ObjectConsExpr) {
	t.Helper()
	f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics: %s", diags)
	}
	obj := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()
	if obj == nil {
		t.Fatalf("attribute a is not an object-construct expression")
	}
	return f, obj
}

func tokensSource(tokens Tokens) string {
	return strings.TrimSpace(string(tokens.Bytes()))
}

func TestObjectConsExprItems(t *testing.T) {
	_, obj := parseObjectConsExprForTest(t, "a = { hat = 1, (cat) = 2, \"dog\" = 3 }\n")

	var got []string
	for _, item := range obj.Items() {
		got = append(got, tokensSource(item.KeyExpr().BuildTokens(nil))+"="+tokensSource(item.ValueExpr().BuildTokens(nil)))
	}
	want := []string{"hat=1", "(cat)=2", `"dog"=3`}
	if !slices.Equal(got, want) {
		t.Errorf("wrong items\ngot:  %q\nwant: %q", got, want)
	}

	_, empty := parseObjectConsExprForTest(t, "a = {}\n")
	if items := empty.Items(); len(items) != 0 {
		t.Errorf("empty object has %d items", len(items))
	}
}

func TestObjectConsExprItemFor(t *testing.T) {
	tests := map[string]struct {
		src string
		key string
		// want is the source of the value that is found, or "" if no
		// item is found.
		want string
	}{
		"identifier":                     {`a = { hat = "derby" }`, "hat", `"derby"`},
		"quoted":                         {`a = { "hat" = "derby" }`, "hat", `"derby"`},
		"colon":                          {`a = { hat: "derby" }`, "hat", `"derby"`},
		"quoted with space":              {`a = { "a hat" = "derby" }`, "a hat", `"derby"`},
		"quoted with escaped template":   {`a = { "$${hat}" = "derby" }`, "${hat}", `"derby"`},
		"quoted with escaped quote":      {`a = { "say \"hi\"" = "derby" }`, `say "hi"`, `"derby"`},
		"keyword":                        {`a = { null = "derby" }`, "null", `"derby"`},
		"second item":                    {`a = { cat = "calico", hat = "derby" }`, "hat", `"derby"`},
		"first of duplicates":            {`a = { hat = 1, "hat" = 2 }`, "hat", `1`},
		"after expression key":           {`a = { (hat) = 1, hat = 2 }`, "hat", `2`},
		"missing":                        {`a = { cat = "calico" }`, "hat", ``},
		"empty object":                   {`a = {}`, "hat", ``},
		"parenthesized":                  {`a = { (hat) = "derby" }`, "hat", ``},
		"traversal":                      {`a = { hat.band = "derby" }`, "hat.band", ``},
		"traversal prefix":               {`a = { hat.band = "derby" }`, "hat", ``},
		"template":                       {`a = { "${hat}" = "derby" }`, "hat", ``},
		"template source":                {`a = { "${hat}" = "derby" }`, "${hat}", ``},
		"escaped template is not raw":    {`a = { "$${hat}" = "derby" }`, "$${hat}", ``},
		"identifier is case sensitive":   {`a = { Hat = "derby" }`, "hat", ``},
		"nested object is not searched":  {`a = { cat = { hat = "derby" } }`, "hat", ``},
		"nested object value is found":   {`a = { cat = { hat = "derby" } }`, "cat", `{ hat = "derby" }`},
		"function call key":              {`a = { upper("hat") = "derby" }`, "HAT", ``},
		"number key":                     {`a = { 1 = "derby" }`, "1", ``},
		"quoted key with only a newline": {`a = { "\n" = "derby" }`, "\n", `"derby"`},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, obj := parseObjectConsExprForTest(t, test.src+"\n")

			item := obj.ItemFor(test.key)
			value := obj.ValueExprFor(test.key)
			if test.want == "" {
				if item != nil || value != nil {
					t.Fatalf("found an item, want none")
				}
				return
			}
			if item == nil || value == nil {
				t.Fatalf("found no item")
			}
			if value != item.ValueExpr() {
				t.Errorf("ValueExprFor returned a different value than ItemFor")
			}
			if got := tokensSource(value.BuildTokens(nil)); got != test.want {
				t.Errorf("wrong value\ngot:  %s\nwant: %s", got, test.want)
			}
		})
	}
}

func TestObjectConsExprSetItemKey(t *testing.T) {
	tests := map[string]struct {
		key  string
		want string
	}{
		"identifier":       {"hat", `hat`},
		"identifier dash":  {"top-hat", `top-hat`},
		"keyword":          {"null", `null`},
		"space":            {"a hat", `"a hat"`},
		"dot":              {"hat.band", `"hat.band"`},
		"leading digit":    {"1hat", `"1hat"`},
		"quote":            {`say "hi"`, `"say \"hi\""`},
		"backslash":        {`a\b`, `"a\\b"`},
		"interpolation":    {"${hat}", `"$${hat}"`},
		"directive":        {"%{hat}", `"%%{hat}"`},
		"escaped template": {"$${hat}", `"$$${hat}"`},
		"newline":          {"a\nb", `"a\nb"`},
		"empty":            {"", `""`},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			f, obj := parseObjectConsExprForTest(t, "a = {}\n")

			key, value := obj.SetItemValue(test.key, cty.True)
			if got := tokensSource(key.BuildTokens(nil)); got != test.want {
				t.Errorf("wrong key\ngot:  %s\nwant: %s", got, test.want)
			}
			if got := obj.ItemFor(test.key); got == nil || got.KeyExpr() != key || got.ValueExpr() != value {
				t.Errorf("ItemFor does not find the item that was set")
			}

			// The key must evaluate back to the original string.
			nf, diags := hclsyntax.ParseConfig(f.Bytes(), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("result is not valid HCL: %s\n%s", diags, f.Bytes())
			}
			val, diags := nf.Body.(*hclsyntax.Body).Attributes["a"].Expr.Value(nil)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			want := cty.ObjectVal(map[string]cty.Value{test.key: cty.True})
			if !val.RawEquals(want) {
				t.Errorf("wrong value\ngot:  %#v\nwant: %#v", val, want)
			}
		})
	}
}

func TestObjectConsExprSetItemReplace(t *testing.T) {
	tests := map[string]struct {
		src  string
		key  string
		want string
		// added is true if SetItemRaw adds a new item instead of
		// replacing a value.
		added bool
	}{
		"identifier": {
			src:  "a = { hat = \"derby\", cat = \"calico\" }\n",
			key:  "hat",
			want: "a = { hat = \"bowler\", cat = \"calico\" }\n",
		},
		"quoted key is left as written": {
			src:  "a = { \"hat\" = \"derby\" }\n",
			key:  "hat",
			want: "a = { \"hat\" = \"bowler\" }\n",
		},
		"escaped key is left as written": {
			src:  "a = { \"$${hat}\" = \"derby\" }\n",
			key:  "${hat}",
			want: "a = { \"$${hat}\" = \"bowler\" }\n",
		},
		"colon is kept": {
			src:  "a = { hat: \"derby\" }\n",
			key:  "hat",
			want: "a = { hat : \"bowler\" }\n",
		},
		"nested object value": {
			src:  "a = {\n  hat = {\n    b = 1\n  }\n  cat = 2\n}\n",
			key:  "hat",
			want: "a = {\n  hat = \"bowler\"\n  cat = 2\n}\n",
		},
		"comments are kept": {
			src:  "a = {\n  # lead\n  hat = \"derby\" # line\n}\n",
			key:  "hat",
			want: "a = {\n  # lead\n  hat = \"bowler\" # line\n}\n",
		},
		"expression key is not replaced": {
			src:   "a = { (hat) = \"derby\" }\n",
			key:   "hat",
			want:  "a = { (hat) = \"derby\", hat = \"bowler\" }\n",
			added: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			f, obj := parseObjectConsExprForTest(t, test.src)
			itemsBefore := len(obj.Items())

			key, value := obj.SetItemRaw(test.key, TokensForValue(cty.StringVal("bowler")))
			if got := string(f.Bytes()); got != test.want {
				t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, test.want)
			}
			item := obj.ItemFor(test.key)
			if item == nil || item.KeyExpr() != key || item.ValueExpr() != value {
				t.Errorf("SetItemRaw did not return the item that ItemFor finds")
			}
			wantItems := itemsBefore
			if test.added {
				wantItems++
			}
			if got := len(obj.Items()); got != wantItems {
				t.Errorf("wrong number of items %d, want %d", got, wantItems)
			}
		})
	}
}

func TestObjectConsExprSetItemHelpers(t *testing.T) {
	f, obj := parseObjectConsExprForTest(t, "a = { hat = \"derby\" }\n")

	obj.SetItemValue("hat", cty.StringVal("bowler"))
	obj.SetItemValue("size", cty.NumberIntVal(7))
	obj.SetItemTraversal("cat", hcl.Traversal{
		hcl.TraverseRoot{Name: "var"},
		hcl.TraverseAttr{Name: "cat"},
	})
	obj.SetItemRaw("dog", Tokens{
		{Type: hclsyntax.TokenIdent, Bytes: []byte("upper")},
		{Type: hclsyntax.TokenOParen, Bytes: []byte("(")},
		{Type: hclsyntax.TokenIdent, Bytes: []byte("var")},
		{Type: hclsyntax.TokenDot, Bytes: []byte(".")},
		{Type: hclsyntax.TokenIdent, Bytes: []byte("dog")},
		{Type: hclsyntax.TokenCParen, Bytes: []byte(")")},
	})

	want := "a = { hat = \"bowler\", size = 7, cat = var.cat, dog = upper(var.dog) }\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result\ngot:  %s\nwant: %s", got, want)
	}

	// The traversal is a variable of the object, but raw tokens are not
	// interpreted, as with NewExpressionRaw.
	expr := f.Body().GetAttribute("a").Expr()
	expr.RenameVariablePrefix([]string{"var"}, []string{"local"})
	want = "a = { hat = \"bowler\", size = 7, cat = local.cat, dog = upper(var.dog) }\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result after rename\ngot:  %s\nwant: %s", got, want)
	}
}

func TestObjectConsExprSetItemPanics(t *testing.T) {
	const src = "a = { hat = \"derby\", cat = { b = 1 } }\nb = 1\n"

	tests := map[string]func(f *File, obj *ObjectConsExpr) *Expression{
		"parsed item value": func(f *File, obj *ObjectConsExpr) *Expression {
			return obj.ValueExprFor("hat").Expr()
		},
		"parsed nested item value": func(f *File, obj *ObjectConsExpr) *Expression {
			return obj.ValueExprFor("cat").Expr().AsObjectConsExpr().ValueExprFor("b").Expr()
		},
		"parsed attribute": func(f *File, obj *ObjectConsExpr) *Expression {
			return f.Body().GetAttribute("b").Expr()
		},
		"object itself": func(f *File, obj *ObjectConsExpr) *Expression {
			return f.Body().GetAttribute("a").Expr()
		},
		"new item value": func(f *File, obj *ObjectConsExpr) *Expression {
			_, value := obj.SetItemValue("dog", cty.True)
			return value.Expr()
		},
		"replaced item value": func(f *File, obj *ObjectConsExpr) *Expression {
			_, value := obj.SetItemValue("hat", cty.True)
			return value.Expr()
		},
		"expression passed to SetItem": func(f *File, obj *ObjectConsExpr) *Expression {
			expr := NewExpressionLiteral(cty.True)
			obj.SetItem("dog", expr)
			return expr
		},
		"attribute set on the body": func(f *File, obj *ObjectConsExpr) *Expression {
			return f.Body().SetAttributeValue("c", cty.True).Expr()
		},
	}

	for name, get := range tests {
		t.Run(name, func(t *testing.T) {
			f, obj := parseObjectConsExprForTest(t, src)
			expr := get(f, obj)
			before := string(f.Bytes())

			func() {
				defer func() {
					if recover() == nil {
						t.Errorf("SetItem did not panic")
					}
				}()
				obj.SetItem("other", expr)
			}()

			if got := string(f.Bytes()); got != before {
				t.Errorf("the tree changed\ngot:\n%s\nwant:\n%s", got, before)
			}
		})
	}

	t.Run("replaced expression can be reused", func(t *testing.T) {
		f, obj := parseObjectConsExprForTest(t, src)
		old := obj.ValueExprFor("hat").Expr()
		obj.SetItemValue("hat", cty.True)
		obj.SetItem("dog", old)

		want := "a = { hat = true, cat = { b = 1 }, dog = \"derby\" }\nb = 1\n"
		if got := string(f.Bytes()); got != want {
			t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, want)
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
			src:  "a = {  }\n",
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
		"single line inline comment before brace": {
			src:  "a = { hat = \"derby\" /* end of a */ }\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\", dog = \"bowler\" /* end of a */ }\n",
		},
		"single line comma and inline comment": {
			src:  "a = { hat = \"derby\", /* end of a */ }\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\", dog = \"bowler\" /* end of a */ }\n",
		},
		"single line comment after attribute": {
			src:  "a = { hat = \"derby\" } # a\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\", dog = \"bowler\" } # a\n",
		},
		"single line nested multi-line value": {
			src:  "a = { hat = {\n  b = 1\n} }\n",
			keys: []string{"dog"},
			want: "a = { hat = {\n  b = 1\n}, dog = \"bowler\" }\n",
		},
		"heredoc value ends a line": {
			src:  "a = { hat = <<EOT\nderby\nEOT\n}\n",
			keys: []string{"dog"},
			want: "a = { hat = <<EOT\nderby\nEOT\n  dog = \"bowler\"\n}\n",
		},
		"multi line": {
			src:  "a = {\n  hat = \"derby\"\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}\n",
		},
		"multi line two items": {
			src:  "a = {\n  hat = \"derby\"\n}\n",
			keys: []string{"dog", "cat"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n}\n",
		},
		"multi line empty": {
			src:  "a = {\n}\n",
			keys: []string{"dog", "cat"},
			want: "a = {\n  dog = \"bowler\"\n  cat = \"bowler\"\n}\n",
		},
		"multi line empty with comment": {
			src:  "a = {\n  # nothing yet\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  dog = \"bowler\"\n  # nothing yet\n}\n",
		},
		"multi line empty with comment after brace": {
			src:  "a = { # nothing yet\n}\n",
			keys: []string{"dog"},
			want: "a = { # nothing yet\n  dog = \"bowler\"\n}\n",
		},
		"multi line commas": {
			src:  "a = {\n  hat = \"derby\",\n  cat = \"calico\",\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\",\n  cat = \"calico\",\n  dog = \"bowler\"\n}\n",
		},
		"brace on the line of the last item": {
			src:  "a = {\n  hat = \"derby\" }\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}\n",
		},
		"brace on the line of the last item with comma": {
			src:  "a = {\n  hat = \"derby\", }\n",
			keys: []string{"dog", "cat"},
			want: "a = {\n  hat = \"derby\",\n  dog = \"bowler\"\n  cat = \"bowler\"\n}\n",
		},
		"items on the line of the opening brace": {
			src:  "a = { hat = \"derby\",\n  cat = \"calico\" }\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\",\n  cat = \"calico\"\n  dog = \"bowler\"\n}\n",
		},
		"line comment after the last item": {
			src:  "a = {\n  hat = \"derby\" # a hat\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\" # a hat\n  dog = \"bowler\"\n}\n",
		},
		"comma and line comment after the last item": {
			src:  "a = {\n  hat = \"derby\", // a hat\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\", // a hat\n  dog = \"bowler\"\n}\n",
		},
		"comments before the closing brace": {
			src:  "a = {\n  hat = \"derby\"\n  # more hats\n\n  # later\n}\n",
			keys: []string{"dog", "cat"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n  # more hats\n\n  # later\n}\n",
		},
		"inline comment before the closing brace": {
			src:  "a = {\n  hat = \"derby\" /* end of a */ }\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n/* end of a */ }\n",
		},
		"inline comment on the line of the closing brace": {
			src:  "a = {\n  hat = \"derby\"\n/* end of a */ }\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n/* end of a */ }\n",
		},
		"inline comment after the last item": {
			src:  "a = {\n  hat = \"derby\" /* a hat */\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\" /* a hat */\n  dog = \"bowler\"\n}\n",
		},
		"multi line nested object": {
			src:  "a = {\n  hat = {\n    b = 1\n  }\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = {\n    b = 1\n  }\n  dog = \"bowler\"\n}\n",
		},
		"nested object": {
			src:  "a = {\n  hat = { b = 1 }\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = { b = 1 }\n  dog = \"bowler\"\n}\n",
		},
		"in a block": {
			src:  "b {\n  a = {\n    hat = \"derby\"\n  }\n}\n",
			keys: []string{"dog"},
			want: "b {\n  a = {\n    hat = \"derby\"\n    dog = \"bowler\"\n  }\n}\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			body := f.Body()
			if blocks := body.Blocks(); len(blocks) > 0 {
				body = blocks[0].Body()
			}
			obj := body.GetAttribute("a").Expr().AsObjectConsExpr()

			for _, key := range test.keys {
				obj.SetItemRaw(key, TokensForValue(cty.StringVal("bowler")))
			}

			got := f.Bytes()
			if string(got) != test.want {
				t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, test.want)
			}
			if formatted := Format(got); string(formatted) != string(got) {
				t.Errorf("result is not in canonical form\ngot:\n%s\nformatted:\n%s", got, formatted)
			}
			if _, diags := hclsyntax.ParseConfig(got, "", hcl.InitialPos); diags.HasErrors() {
				t.Errorf("result is not valid HCL: %s", diags)
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

#### y/ast_tuple_cons_expr.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

// TupleConsExpr represents a tuple-construct expression such as
// ["fedora", "fez"]. Use Expression.AsTupleConsExpr to get one.
//
// Each element is its own Expression. The brackets and the commas are tokens
// of the TupleConsExpr itself.
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

// variables returns the traversals in the elements, in source order.
func (tuple *TupleConsExpr) variables() []*Traversal {
	return expressionsVariables(tuple.elems)
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
	config := []byte(`a = {
    hat = "derby", (cat) = "calico" }` + "\n")

	bowler := hclwrite.TokensForValue(cty.StringVal(`bowler`))

	file, diags := hclwrite.ParseConfig(config, "<inline>", hcl.Pos{})
	if diags.HasErrors() {
		fmt.Printf("errors: %s", diags)
		return
	}
	expr := file.Body().GetAttribute("a").Expr().AsObjectConsExpr()

	// Replaces the value of the existing item "hat".
	key, value := expr.SetItemRaw("hat", bowler)
	fmt.Printf("%s = %s\n", key.BuildTokens(nil).Bytes(), value.BuildTokens(nil).Bytes())

	// (cat) is an expression, not the name "cat", so this adds a new item.
	expr.SetItemRaw("cat", bowler)

	fmt.Printf("%s", file.Bytes())
	// Output:
	// hat = "bowler"
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
	setAttached(n.content, false)
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
	setAttached(n.content, false)
	setAttached(c, true)
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
	for n := ns.first; n != nil; n = n.after {
		setAttached(n.content, false)
	}
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
	setAttached(n.content, true)
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
	setAttached(n.content, true)
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
// whether the content is part of a syntax tree.
type inTree struct {
	// attached is true while the content belongs to a node in a list.
	// The methods of nodes and node keep it up to date.
	attached bool
	children *nodes
}

func newInTree() inTree {
	return inTree{
		children: &nodes{},
	}
}

func (it *inTree) assertUnattached() {
	if it.attached {
		panic("node is already attached to a syntax tree")
	}
}

func (it *inTree) setAttached(attached bool) {
	it.attached = attached
}

// setAttached records on the given content whether it belongs to a node in
// a list. Content that does not embed inTree is not tracked.
func setAttached(c nodeContent, attached bool) {
	if a, ok := c.(interface{ setAttached(bool) }); ok {
		a.setAttached(attached)
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
	children.AppendUnstructuredTokens(from.writerTokens)

	// Wrap in an Expression
	wrapExpr := newExpression()
	wrapExpr.wrap(expr)
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
// Each argument becomes its own Expression. All other tokens, including any
// comments between the arguments, stay as tokens of the FunctionCallExpr.
func parseFunctionCallExpr(nativeExpr *hclsyntax.FunctionCallExpr, from inputTokens) *node {
	call := newFunctionCallExpr()
	children := call.children

	before, name, from := from.Partition(nativeExpr.NameRange)
	children.AppendUnstructuredTokens(before.writerTokens)
	call.name = children.AppendUnstructuredTokens(name.writerTokens)

	parseExpressionList(children, call.args, nativeExpr.Args, from)

	wrapExpr := newExpression()
	wrapExpr.wrap(call)
	return newNode(wrapExpr)
}

// parseTupleConsExpr parses a tuple-construct expression, defined as:
//
//	tuple = "[" (
//	    (Expression (("," | Newline) Expression)* ","?)?
//	) "]";
//
// Each element becomes its own Expression. All other tokens, including any
// comments between the elements, stay as tokens of the TupleConsExpr.
func parseTupleConsExpr(nativeExpr *hclsyntax.TupleConsExpr, from inputTokens) *node {
	tuple := newTupleConsExpr()

	parseExpressionList(tuple.children, tuple.elems, nativeExpr.Exprs, from)

	wrapExpr := newExpression()
	wrapExpr.wrap(tuple)
	return newNode(wrapExpr)
}

// parseExpressionList appends each of the given native expressions to
// children as its own Expression node, and records it in exprs. The tokens
// around the expressions are appended to children as unstructured tokens.
func parseExpressionList(children *nodes, exprs nodeSet, nativeExprs []hclsyntax.Expression, from inputTokens) {
	for _, nativeExpr := range nativeExprs {
		var before, exprTokens inputTokens
		before, exprTokens, from = from.Partition(nativeExpr.Range())
		children.AppendUnstructuredTokens(before.writerTokens)
		expr := parseExpression(nativeExpr, exprTokens)
		children.AppendNode(expr)
		exprs.Add(expr)
	}
	children.AppendUnstructuredTokens(from.writerTokens)
}

// parseObjectConsKeyExpr parses an object-construct key expression.
// An object key is an Identifier or an Expression. In this context, a quoted
// literal is functionally equivalent to an Identifier.
//
// A naked identifier key is a literal name, not a variable, so its traversal
// stays in the tree but is not recorded as an absolute traversal. This uses
// the same rule as hclsyntax.ObjectConsKeyExpr.
func parseObjectConsKeyExpr(nativeExpr *hclsyntax.ObjectConsKeyExpr, from inputTokens) *node {
	expr := parseExpression(nativeExpr.Wrapped, from)
	if hcl.ExprAsKeyword(nativeExpr.Wrapped) != "" {
		expr.content.(*Expression).absTraversals.Clear()
	}
	wrapExpr := newObjectConsKeyExpr(expr)

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
			`a = join(" ", ["fedora", "fez"])` + "\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{Type: "comments"},
							{Type: "identifier", Val: "a"},
							{Type: "Tokens", Val: " ="},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "FunctionCallExpr",
										Children: []TestTreeNode{
											{Type: "Tokens", Val: " join"},
											{Type: "Tokens", Val: "("},
											{
												Type: "Expression",
												Children: []TestTreeNode{
													{Type: "quoted", Val: `" "`},
												},
											},
											{Type: "Tokens", Val: ","},
											{
												Type: "Expression",
												Children: []TestTreeNode{
													{
														Type: "TupleConsExpr",
														Children: []TestTreeNode{
															{Type: "Tokens", Val: " ["},
															{
																Type: "Expression",
																Children: []TestTreeNode{
																	{Type: "quoted", Val: `"fedora"`},
																},
															},
															{Type: "Tokens", Val: ","},
															{
																Type: "Expression",
																Children: []TestTreeNode{
																	{Type: "quoted", Val: ` "fez"`},
																},
															},
															{Type: "Tokens", Val: "]"},
														},
													},
												},
											},
											{Type: "Tokens", Val: ")"},
										},
									},
								},
							},
							{Type: "comments"},
							{Type: "Tokens", Val: "\n"},
						},
					},
				},
			},
		},
		{
			"a = provider::p::f()\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{Type: "comments"},
							{Type: "identifier", Val: "a"},
							{Type: "Tokens", Val: " ="},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "FunctionCallExpr",
										Children: []TestTreeNode{
											{Type: "Tokens", Val: " provider::p::f"},
											{Type: "Tokens", Val: "()"},
										},
									},
								},
							},
							{Type: "comments"},
							{Type: "Tokens", Val: "\n"},
						},
					},
				},
			},
		},
		{
			"a = f(b.c, [d]...)\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{Type: "comments"},
							{Type: "identifier", Val: "a"},
							{Type: "Tokens", Val: " ="},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "FunctionCallExpr",
										Children: []TestTreeNode{
											{Type: "Tokens", Val: " f"},
											{Type: "Tokens", Val: "("},
											{
												Type: "Expression",
												Children: []TestTreeNode{
													{
														Type: "Traversal",
														Children: []TestTreeNode{
															{
																Type: "TraverseName",
																Children: []TestTreeNode{
																	{Type: "identifier", Val: "b"},
																},
															},
															{
																Type: "TraverseName",
																Children: []TestTreeNode{
																	{Type: "Tokens", Val: "."},
																	{Type: "identifier", Val: "c"},
																},
															},
														},
													},
												},
											},
											{Type: "Tokens", Val: ","},
											{
												Type: "Expression",
												Children: []TestTreeNode{
													{
														Type: "TupleConsExpr",
														Children: []TestTreeNode{
															{Type: "Tokens", Val: " ["},
															{
																Type: "Expression",
																Children: []TestTreeNode{
																	{
																		Type: "Traversal",
																		Children: []TestTreeNode{
																			{
																				Type: "TraverseName",
																				Children: []TestTreeNode{
																					{Type: "identifier", Val: "d"},
																				},
																			},
																		},
																	},
																},
															},
															{Type: "Tokens", Val: "]"},
														},
													},
												},
											},
											{Type: "Tokens", Val: "...)"},
										},
									},
								},
							},
							{Type: "comments"},
							{Type: "Tokens", Val: "\n"},
						},
					},
				},
			},
		},
		{
			"a = [\n  1, # one\n  { b = 2 },\n]\n",
			TestTreeNode{
				Type: "Body",
				Children: []TestTreeNode{
					{
						Type: "Attribute",
						Children: []TestTreeNode{
							{Type: "comments"},
							{Type: "identifier", Val: "a"},
							{Type: "Tokens", Val: " ="},
							{
								Type: "Expression",
								Children: []TestTreeNode{
									{
										Type: "TupleConsExpr",
										Children: []TestTreeNode{
											{Type: "Tokens", Val: " [\n"},
											{
												Type: "Expression",
												Children: []TestTreeNode{
													{Type: "Tokens", Val: "  1"},
												},
											},
											{Type: "Tokens", Val: ", # one\n"},
											{
												Type: "Expression",
												Children: []TestTreeNode{
													{
														Type: "ObjectConsExpr",
														Children: []TestTreeNode{
															{Type: "Tokens", Val: "  {"},
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
																									{Type: "identifier", Val: " b"},
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
																				Type: "Expression",
																				Children: []TestTreeNode{
																					{Type: "Tokens", Val: " 2"},
																				},
																			},
																		},
																	},
																},
															},
															{Type: "Tokens", Val: " }"},
														},
													},
												},
											},
											{Type: "Tokens", Val: ",\n]"},
										},
									},
								},
							},
							{Type: "comments"},
							{Type: "Tokens", Val: "\n"},
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
{"better": "x" | "y" | "tie", "reasons": ["...", "..."]}
Give two to five short reasons, each pointing at concrete files or types.
