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
| ast_body.go | 309 |  | Body (type), newBody (func), Body.appendItem (method), Body.appendItemNode (method), Body.Clear (method), Body.AppendUnstructuredTokens (method), Body.Attributes (method), Body.Blocks (method), Body.GetAttribute (method), Body.getAttributeNode (method), Body.RenameAttribute (method), Body.FirstMatchingBlock (method), Body.RemoveNewlineBeforeBlock (method), Body.RemoveBlock (method), Body.SetAttributeRaw (method), Body.SetAttributeValue (method), Body.SetAttributeTraversal (method), Body.RemoveAttribute (method), Body.AppendBlock (method), Body.AppendNewBlock (method), Body.AppendNewline (method) |
| ast_body_test.go | 2195 |  | TestBodyGetAttribute (func), TestBodyFirstMatchingBlock (func), TestBodySetAttributeValue (func), TestBodySetAttributeTraversal (func), TestBodySetAttributeTraversal_ReturnsTheAttribute (func), TestBodySetAttributeRaw (func), TestBodySetAttributeRaw_ReturnsTheAttribute (func), TestBodySetAttributeValue_ReturnsTheAttribute (func), TestBodySetAttributeValueInBlock (func), TestBodySetAttributeValueInNestedBlock (func), TestBodyRemoveAttribute (func), TestBodyRenameAttribute (func), TestBodyAppendBlock (func), TestBodyRemoveBlock (func), TestBodyRemoveNewlineBeforeBlock (func), testFn (func) |
| ast_expression.go | 318 | CHANGED | Expression (type), newExpression (func), Expression.wrap (method), NewExpressionRaw (func), NewExpressionLiteral (func), NewExpressionAbsTraversal (func), Expression.Variables (method), variablesOf (func), Expression.RenameVariablePrefix (method), Expression.AsObjectConsExpr (method), Expression.AsFunctionCallExpr (method), Expression.AsTupleConsExpr (method), Expression.asQuoted (method), Expression.AsQuotedLiteral (method), Traversal (type), newTraversal (func), TraverseName (type), newTraverseName (func), TraverseIndex (type), newTraverseIndex (func) |
| ast_expression_test.go | 236 | NEW | TestExpressionVariablesObjectCons (func), TestExpressionVariablesFunctionCallAndTupleCons (func), TestExpressionAsFunctionCallAndTupleCons (func) |
| ast_function_call_expr.go | 36 | NEW | FunctionCallExpr (type), newFunctionCallExpr (func), FunctionCallExpr.Name (method), FunctionCallExpr.variables (method) |
| ast_object_cons_expr.go | 340 | CHANGED | ObjectConsExpr (type), newObjectConsExpr (func), ObjectConsExpr.Items (method), ObjectConsExpr.ItemFor (method), ObjectConsExpr.ValueExprFor (method), ObjectConsExpr.SetItem (method), ObjectConsExpr.SetItemRaw (method), ObjectConsExpr.SetItemValue (method), ObjectConsExpr.SetItemTraversal (method), ObjectConsExpr.appendItem (method), ObjectConsExpr.isMultiLine (method), hasLineEnd (func), isLineEnd (func), newNewlineToken (func), ObjectConsItem (type), newObjectConsItem (func), ObjectConsItem.KeyExpr (method), ObjectConsItem.ValueExpr (method), ObjectConsKeyExpr (type), newObjectConsKeyExpr (func), newObjectConsKeyExprForName (func), ObjectConsKeyExpr.name (method), ObjectConsExpr.variables (method), ObjectConsValue (type), newObjectConsValue (func), ObjectConsValue.Expr (method) |
| ast_object_cons_expr_test.go | 510 | NEW | parseObjectAttr (func), tokensText (func), evalObjectAttr (func), TestObjectConsExprItemFor (func), TestObjectConsExprItems (func), TestObjectConsExprSetItemKey (func), TestObjectConsExprSetItemReplace (func), TestObjectConsExprSetItemHelpers (func), TestObjectConsExprSetItemPanics (func), TestObjectConsExprSetItemLayout (func) |
| ast_test.go | 50 |  | TestTreeNode (type), makeTestTree (func) |
| ast_tuple_cons_expr.go | 28 | NEW | TupleConsExpr (type), newTupleConsExpr (func), TupleConsExpr.variables (method) |
| doc.go | 14 |  |  |
| examples_test.go | 125 | CHANGED | Example_generateFromScratch (func), ExampleExpression_RenameVariablePrefix (func), ExampleObjectConsExpr_SetItemRaw (func) |
| format.go | 472 |  | format (func), formatIndent (func), formatSpaces (func), formatCells (func), spaceAfterToken (func), linesForFormat (func), tokenIsNewline (func), tokenBracketChange (func), formatLine (type) |
| format_test.go | 800 |  | TestFormat (func), TestLinesForFormat (func) |
| generate.go | 399 |  | TokensForValue (func), TokensForTraversal (func), TokensForIdentifier (func), TokensForTuple (func), TokensForObject (func), TokensForFunctionCall (func), appendTokensForValue (func), appendTokensForTraversal (func), appendTokensForTraversalStep (func), escapeQuotedStringLit (func), appendRune (func) |
| generate_test.go | 920 |  | TestTokensForValue (func), TestTokensForTraversal (func), TestTokensForTuple (func), TestTokensForObject (func), TestTokensForFunctionCall (func), TestTokenGenerateConsistency (func) |
| native_node_sorter.go | 26 |  | nativeNodeSorter (type), nativeNodeSorter.Len (method), nativeNodeSorter.Less (method), nativeNodeSorter.Swap (method) |
| node.go | 299 |  | node (type), newNode (func), node.Equal (method), node.BuildTokens (method), node.Detach (method), node.ReplaceWith (method), node.assertUnattached (method), nodeContent (type), nodes (type), nodes.BuildTokens (method), nodes.Clear (method), nodes.Append (method), nodes.AppendNode (method), nodes.Insert (method), nodes.InsertNode (method), nodes.AppendUnstructuredTokens (method), nodes.FindNodeWithContent (method), nodeSet (type), newNodeSet (func), nodeSet.Has (method), nodeSet.Add (method), nodeSet.Remove (method), nodeSet.Clear (method), nodeSet.List (method), nodeSet.FindNodeWithContent (method), internalWalkFunc (type), inTree (type), newInTree (func), inTree.assertUnattached (method), inTree.walkChildNodes (method), inTree.BuildTokens (method), leafNode (type), leafNode.walkChildNodes (method) |
| parser.go | 794 | CHANGED | parse (func), inputTokens (type), inputTokens.Partition (method), inputTokens.PartitionType (method), inputTokens.PartitionTypeOk (method), inputTokens.PartitionTypeSingle (method), inputTokens.PartitionIncludingComments (method), inputTokens.PartitionBlockItem (method), inputTokens.PartitionLeadComments (method), inputTokens.PartitionLineEndTokens (method), inputTokens.Slice (method), inputTokens.Len (method), inputTokens.Tokens (method), inputTokens.Types (method), parseBody (func), parseBodyItem (func), parseAttribute (func), parseBlock (func), parseBlockLabels (func), parseExpression (func), parseAnyExpression (func), parseObjectConsExpr (func), parseObjectConsKeyExpr (func), parseFunctionCallExpr (func), parseTupleConsExpr (func), parseExprList (func), parseTemplateExpr (func), parseTraversal (func), parseTraversalStep (func), writerTokens (func), partitionTokens (func), partitionLeadCommentTokens (func), partitionLineEndTokens (func), lexConfig (func) |
| parser_test.go | 2467 | CHANGED | TestParse (func), TestPartitionTokens (func), TestPartitionLeadCommentTokens (func), TestLexConfig (func) |
| public.go | 48 |  | NewFile (func), ParseConfig (func), Format (func) |
| round_trip_test.go | 189 | CHANGED | TestRoundTripVerbatim (func), TestRoundTripFormat (func), TestRoundTripSafeConcurrent (func) |
| tokens.go | 136 |  | Token (type), Token.asHCLSyntax (method), Tokens (type), Tokens.Bytes (method), Tokens.testValue (method), Tokens.Columns (method), Tokens.WriteTo (method), Tokens.walkChildNodes (method), Tokens.BuildTokens (method), ObjectAttrTokens (type), newIdentToken (func) |

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
// expression. For an object-construct expression, this includes the
// traversals in the keys and values of its items. For a function call or a
// tuple-construct expression, this includes the traversals in its arguments
// or elements.
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

// variablesOf returns the absolute traversals in the expressions of the
// given set, in source order.
func variablesOf(exprs nodeSet) []*Traversal {
	var ret []*Traversal
	for _, n := range exprs.List() {
		ret = append(ret, n.content.(*Expression).Variables()...)
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

func TestExpressionVariablesObjectCons(t *testing.T) {
	tests := []struct {
		src     string
		want    []string // Variables, in order, as source text
		renamed string   // source after renaming "old" to "new"
	}{
		{
			"a = { k = old.x }\n",
			[]string{"old.x"},
			"a = { k = new.x }\n",
		},
		{
			"a = { k = old.x, j = other.y, l = old.z }\n",
			[]string{"old.x", "other.y", "old.z"},
			"a = { k = new.x, j = other.y, l = new.z }\n",
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
			"a = { (old.x) = old.y }\n",
			[]string{"old.x", "old.y"},
			"a = { (new.x) = new.y }\n",
		},
		{
			"a = { \"k\" = old.x }\n",
			[]string{"old.x"},
			"a = { \"k\" = new.x }\n",
		},
		{
			"a = {\n  k = old.x\n  j = {\n    l = [f(old.y)]\n  }\n}\n",
			[]string{"old.x", "old.y"},
			"a = {\n  k = new.x\n  j = {\n    l = [f(new.y)]\n  }\n}\n",
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
				t.Errorf("wrong variables\n%s", diff)
			}

			nativeFile, diags := hclsyntax.ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected hclsyntax diagnostics: %s", diags.Error())
			}
			nativeExpr := nativeFile.Body.(*hclsyntax.Body).Attributes["a"].Expr
			if got, want := len(expr.Variables()), len(nativeExpr.Variables()); got != want {
				t.Errorf("got %d variables, but hclsyntax reports %d", got, want)
			}

			expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
			if got := string(f.Bytes()); got != test.renamed {
				t.Errorf("wrong result after rename\ngot:  %q\nwant: %q", got, test.renamed)
			}
		})
	}
}

func TestExpressionVariablesFunctionCallAndTupleCons(t *testing.T) {
	tests := []struct {
		src     string
		want    []string // Variables, in order, as source text
		renamed string   // source after renaming "old" to "new"
	}{
		{
			"a = f(old.x)\n",
			[]string{"old.x"},
			"a = f(new.x)\n",
		},
		{
			"a = f(old.x, other.y, old.z...)\n",
			[]string{"old.x", "other.y", "old.z"},
			"a = f(new.x, other.y, new.z...)\n",
		},
		{
			"a = [old.x, other.y, old.z]\n",
			[]string{"old.x", "other.y", "old.z"},
			"a = [new.x, other.y, new.z]\n",
		},
		{
			"a = f(old.w, [old.x, g(old.y)], { k = old.z })\n",
			[]string{"old.w", "old.x", "old.y", "old.z"},
			"a = f(new.w, [new.x, g(new.y)], { k = new.z })\n",
		},
		{
			"a = [\n  old.x,\n  f(old.y + 1),\n]\n",
			[]string{"old.x", "old.y"},
			"a = [\n  new.x,\n  f(new.y + 1),\n]\n",
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
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			if got := string(f.Bytes()); got != test.src {
				t.Errorf("wrong result before rename\ngot:  %q\nwant: %q", got, test.src)
			}
			expr := f.Body().GetAttribute("a").Expr()

			var got []string
			for _, traversal := range expr.Variables() {
				got = append(got, strings.TrimSpace(string(traversal.BuildTokens(nil).Bytes())))
			}
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("wrong variables\n%s", diff)
			}

			expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
			if got := string(f.Bytes()); got != test.renamed {
				t.Errorf("wrong result after rename\ngot:  %q\nwant: %q", got, test.renamed)
			}
		})
	}
}

func TestExpressionAsFunctionCallAndTupleCons(t *testing.T) {
	tests := []struct {
		src      string
		wantCall string // name of the function, or "" if not a call
		wantArgs int    // number of arguments or elements
		isTuple  bool
	}{
		{`join(" ", ["fedora", "fez"])`, "join", 2, false},
		{`f()`, "f", 0, false},
		{`provider::local::upper("hat")`, "provider::local::upper", 1, false},
		{`["fedora", "fez"]`, "", 2, true},
		{`[]`, "", 0, true},
		{`f(1) + 1`, "", 0, false},
		{`[1][0]`, "", 0, false},
		{`[for x in y : x]`, "", 0, false},
		{`{ k = [1] }`, "", 0, false},
		{`"fedora"`, "", 0, false},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte("a = "+test.src+"\n"), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			expr := f.Body().GetAttribute("a").Expr()

			call := expr.AsFunctionCallExpr()
			switch {
			case test.wantCall == "" && call != nil:
				t.Errorf("unexpected function call %q", call.Name())
			case test.wantCall != "" && call == nil:
				t.Errorf("not a function call")
			case call != nil:
				if got := call.Name(); got != test.wantCall {
					t.Errorf("wrong name %q; want %q", got, test.wantCall)
				}
				if got := len(call.args); got != test.wantArgs {
					t.Errorf("wrong number of arguments %d; want %d", got, test.wantArgs)
				}
			}

			tuple := expr.AsTupleConsExpr()
			switch {
			case !test.isTuple && tuple != nil:
				t.Errorf("unexpected tuple")
			case test.isTuple && tuple == nil:
				t.Errorf("not a tuple")
			case tuple != nil:
				if got := len(tuple.elems); got != test.wantArgs {
					t.Errorf("wrong number of elements %d; want %d", got, test.wantArgs)
				}
			}
		})
	}
}

```

#### pkg/ast_function_call_expr.go

```go
// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

// FunctionCallExpr is a function call expression, such as join(" ", list).
// Use Expression.AsFunctionCallExpr to get one from an expression.
//
// Each argument is its own Expression. The function name, the parentheses,
// the commas and a final "..." are unstructured tokens of the call.
type FunctionCallExpr struct {
	inTree

	name string
	args nodeSet
}

func newFunctionCallExpr(name string) *FunctionCallExpr {
	return &FunctionCallExpr{
		inTree: newInTree(),
		name:   name,
		args:   newNodeSet(),
	}
}

// Name returns the name of the called function. A namespaced function name
// includes its namespaces, such as provider::local::func.
func (call *FunctionCallExpr) Name() string {
	return call.name
}

// variables returns the absolute traversals in the arguments of the call,
// in source order.
func (call *FunctionCallExpr) variables() []*Traversal {
	return variablesOf(call.args)
}

```

#### pkg/ast_object_cons_expr.go

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
// { hat = "derby" }. Use Expression.AsObjectConsExpr to get one from an
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
	ret := make([]*ObjectConsItem, len(nodes))
	for i, n := range nodes {
		ret[i] = n.content.(*ObjectConsItem)
	}
	return ret
}

// ItemFor returns the item of the object that has the given key, or nil if
// there is no such item.
//
// Only keys that are an identifier or a quoted string literal can match. The
// two forms are equivalent, so ItemFor("hat") finds both hat = ... and
// "hat" = .... A quoted key is compared by its value after escape sequences
// are decoded. Keys that are taken from a variable or another expression,
// such as (hat) or "${hat}", never match.
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
	if item := o.ItemFor(key); item != nil {
		return item.ValueExpr()
	}
	return nil
}

// SetItem either replaces the value of the item that ItemFor finds for the
// given key, or adds a new item to the end of the object. An existing key is
// left as written. A new key is written as an identifier if it is a valid
// identifier, and as a quoted string literal otherwise.
//
// The expression must not already be part of a syntax tree, or SetItem
// panics. Where possible, use SetItemValue or SetItemTraversal instead.
//
// The return values are the key and the value of the item that was either
// modified in-place or created.
func (o *ObjectConsExpr) SetItem(key string, expr *Expression) (*ObjectConsKeyExpr, *ObjectConsValue) {
	expr.assertUnattached()

	item := o.ItemFor(key)
	if item == nil {
		item = o.appendItem(newObjectConsKeyExprForName(key))
	} else {
		old := item.ValueExpr()
		old.Expr().parent = nil
		old.children.Clear()
	}

	value := item.ValueExpr()
	value.expr = value.children.Append(expr)
	expr.parent = item.value
	return item.KeyExpr(), value
}

// SetItemRaw is like SetItem, but uses the given tokens verbatim as the
// value.
//
// The same caveats apply to this function as for NewExpressionRaw on which
// it is based. If possible, prefer to use SetItemValue or SetItemTraversal.
func (o *ObjectConsExpr) SetItemRaw(key string, tokens Tokens) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionRaw(tokens))
}

// SetItemValue is like SetItem, but the value is given as a cty.Value, and
// must therefore be a literal. To set a variable reference or other
// traversal, use SetItemTraversal.
func (o *ObjectConsExpr) SetItemValue(key string, val cty.Value) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionLiteral(val))
}

// SetItemTraversal is like SetItem, but the value is given as a
// hcl.Traversal, which must be an absolute traversal. To set a literal value,
// use SetItemValue.
func (o *ObjectConsExpr) SetItemTraversal(key string, traversal hcl.Traversal) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionAbsTraversal(traversal))
}

// appendItem adds a new item with the given key and an empty value to the
// end of the object, and returns it.
//
// In an object written on a single line, the item goes on the same line
// after a comma. Otherwise it goes on its own line after the last item. A
// comma or a line comment that follows the last item on its line stays with
// that item. Anything after that, such as comments on the following lines,
// stays with the closing brace.
func (o *ObjectConsExpr) appendItem(key *ObjectConsKeyExpr) *ObjectConsItem {
	// The parser always puts the tokens after the last item, up to and
	// including the closing brace, in the last child node.
	closing := o.children.last
	rest := closing.content.(Tokens)

	var lead, tail Tokens
	if o.isMultiLine() {
		// Keep everything up to the end of the last item's line before the
		// new item. If the closing brace is on that line, keep only a comma.
		end := 0
		for i, t := range rest {
			if isLineEnd(t) {
				end = i + 1
				break
			}
		}
		if end == 0 && len(rest) > 0 && rest[0].Type == hclsyntax.TokenComma {
			end = 1
		}
		lead = append(lead, rest[:end]...)
		if end == 0 || !isLineEnd(rest[end-1]) {
			lead = append(lead, newNewlineToken())
		}
		tail = append(Tokens{newNewlineToken()}, rest[end:]...)
	} else {
		tail = rest
		if len(tail) > 0 && tail[0].Type == hclsyntax.TokenComma {
			tail = tail[1:]
		}
		if len(o.items) > 0 {
			lead = Tokens{{Type: hclsyntax.TokenComma, Bytes: []byte{','}}}
		}
	}

	item := newObjectConsItem()
	item.children.AppendUnstructuredTokens(lead)
	item.key = item.children.Append(key)
	item.children.AppendUnstructuredTokens(Tokens{
		{Type: hclsyntax.TokenEqual, Bytes: []byte{'='}, SpacesBefore: 1},
	})
	item.value = item.children.Append(newObjectConsValue())

	closing.Detach()
	o.items.Add(o.children.Append(item))
	o.children.AppendUnstructuredTokens(tail)
	return item
}

// isMultiLine returns true if the object is written on several lines. Only
// the tokens between the items and the braces count, so line breaks inside
// a key or a value do not make the object multi-line.
func (o *ObjectConsExpr) isMultiLine() bool {
	if hasLineEnd(o.children) {
		return true
	}
	for _, item := range o.Items() {
		if hasLineEnd(item.children) {
			return true
		}
	}
	return false
}

// hasLineEnd returns true if any of the unstructured tokens directly in ns
// ends a line.
func hasLineEnd(ns *nodes) bool {
	for n := ns.first; n != nil; n = n.after {
		if tokens, ok := n.content.(Tokens); ok {
			for _, t := range tokens {
				if isLineEnd(t) {
					return true
				}
			}
		}
	}
	return false
}

// isLineEnd returns true if the token ends a line. A line comment includes
// the newline that ends it.
func isLineEnd(t *Token) bool {
	return t.Type == hclsyntax.TokenNewline ||
		(t.Type == hclsyntax.TokenComment && bytes.HasSuffix(t.Bytes, []byte{'\n'}))
}

func newNewlineToken() *Token {
	return &Token{Type: hclsyntax.TokenNewline, Bytes: []byte{'\n'}}
}

// ObjectConsItem is one key and value pair of an object-construct
// expression.
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

// ObjectConsKeyExpr is the key of an item in an object-construct expression.
// It is an identifier, such as hat, or an expression, such as "hat" or
// (hat).
type ObjectConsKeyExpr struct {
	inTree

	wrapped *node

	// literal is true when the key is a naked identifier, which hclsyntax
	// treats as a literal name rather than as a variable reference.
	literal bool
}

func newObjectConsKeyExpr(wrapped *node) *ObjectConsKeyExpr {
	expr := &ObjectConsKeyExpr{
		inTree:  newInTree(),
		wrapped: wrapped,
	}
	expr.children.AppendNode(wrapped)

	return expr
}

// newObjectConsKeyExprForName returns a key that evaluates to the given name.
// It is an identifier if name is a valid identifier, and a quoted string
// literal otherwise. "for" is always quoted, because { for = ... } starts a
// for expression.
func newObjectConsKeyExprForName(name string) *ObjectConsKeyExpr {
	if hclsyntax.ValidIdentifier(name) && name != "for" {
		expr := NewExpressionAbsTraversal(hcl.Traversal{hcl.TraverseRoot{Name: name}})
		key := newObjectConsKeyExpr(newNode(expr))
		key.literal = true
		return key
	}

	// TokensForValue escapes the string so that it evaluates back to name.
	expr := newExpression()
	expr.wrapped = expr.children.Append(newQuoted(TokensForValue(cty.StringVal(name))))
	return newObjectConsKeyExpr(newNode(expr))
}

// name returns the string that the key evaluates to, if the key is an
// identifier or a quoted string literal. Otherwise it returns false.
func (k *ObjectConsKeyExpr) name() (string, bool) {
	expr := k.wrapped.content.(*Expression)

	if k.literal {
		tokens := expr.BuildTokens(nil)
		if len(tokens) == 1 && tokens[0].Type == hclsyntax.TokenIdent {
			return string(tokens[0].Bytes), true
		}
		return "", false
	}

	quoted := expr.asQuoted()
	if quoted == nil {
		return "", false
	}
	// The parser only makes a quoted node for a string literal, so it is
	// safe to evaluate it without a context to decode its escape sequences.
	native, diags := hclsyntax.ParseExpression(quoted.tokens.Bytes(), "", hcl.InitialPos)
	if diags.HasErrors() {
		return "", false
	}
	val, diags := native.Value(nil)
	if diags.HasErrors() || !val.IsKnown() || val.IsNull() || !val.Type().Equals(cty.String) {
		return "", false
	}
	return val.AsString(), true
}

// variables returns the absolute traversals in the keys and values of the
// object's items, in source order. Literal keys contribute no traversals.
func (o *ObjectConsExpr) variables() []*Traversal {
	var ret []*Traversal
	for _, n := range o.items.List() {
		item := n.content.(*ObjectConsItem)
		key := item.key.content.(*ObjectConsKeyExpr)
		if !key.literal {
			ret = append(ret, key.wrapped.content.(*Expression).Variables()...)
		}
		ret = append(ret, item.value.content.(*ObjectConsValue).Expr().Variables()...)
	}
	return ret
}

// ObjectConsValue is the value of an item in an object-construct expression.
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

// parseObjectAttr parses src and returns the file and the object-construct
// expression of its attribute "a".
func parseObjectAttr(t *testing.T, src string) (*File, *ObjectConsExpr) {
	t.Helper()
	f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics: %s", diags.Error())
	}
	object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()
	if object == nil {
		t.Fatalf("attribute a is not an object")
	}
	return f, object
}

// tokensText returns the source text of tokens without surrounding spaces.
func tokensText(tokens Tokens) string {
	return strings.TrimSpace(string(tokens.Bytes()))
}

// evalObjectAttr evaluates attribute "a" of src with hclsyntax.
func evalObjectAttr(t *testing.T, src []byte) cty.Value {
	t.Helper()
	f, diags := hclsyntax.ParseConfig(src, "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("result is not valid HCL: %s\n%s", diags.Error(), src)
	}
	val, diags := f.Body.(*hclsyntax.Body).Attributes["a"].Expr.Value(nil)
	if diags.HasErrors() {
		t.Fatalf("can't evaluate result: %s\n%s", diags.Error(), src)
	}
	return val
}

func TestObjectConsExprItemFor(t *testing.T) {
	tests := []struct {
		src  string
		key  string
		want string // value as source text, or "" if not found
	}{
		{`a = { hat = 1 }`, "hat", "1"},
		{`a = { "hat" = 1 }`, "hat", "1"},
		{`a = { hat: 1 }`, "hat", "1"},
		{`a = { "a hat" = 1 }`, "a hat", "1"},
		{`a = { "$${hat}" = 1 }`, "${hat}", "1"},
		{`a = { "%%{hat}" = 1 }`, "%{hat}", "1"},
		{`a = { "say \"hi\"" = 1 }`, `say "hi"`, "1"},
		{`a = { "a\\b" = 1 }`, `a\b`, "1"},
		{`a = { "" = 1 }`, "", "1"},
		{`a = { true = 1 }`, "true", "1"},
		{`a = { cat = 1, hat = 2 }`, "hat", "2"},
		{`a = { hat = { cat = 1 } }`, "hat", "{ cat = 1 }"},
		{`a = { hat = 1 }`, "cat", ""},
		{`a = { hat = 1 }`, `"hat"`, ""},
		{`a = { "$${hat}" = 1 }`, "$${hat}", ""},
		{`a = { (hat) = 1 }`, "hat", ""},
		{`a = { (hat) = 1 }`, "(hat)", ""},
		{`a = { hat.band = 1 }`, "hat", ""},
		{`a = { hat.band = 1 }`, "hat.band", ""},
		{`a = { "${hat}" = 1 }`, "hat", ""},
		{`a = { "${hat}" = 1 }`, "${hat}", ""},
		{`a = { "${"hat"}" = 1 }`, "hat", ""},
		{`a = {}`, "hat", ""},
	}

	for _, test := range tests {
		t.Run(test.src+" "+test.key, func(t *testing.T) {
			_, object := parseObjectAttr(t, test.src+"\n")

			item := object.ItemFor(test.key)
			value := object.ValueExprFor(test.key)
			if test.want == "" {
				if item != nil {
					t.Errorf("ItemFor found an item, but want nil")
				}
				if value != nil {
					t.Errorf("ValueExprFor found a value, but want nil")
				}
				return
			}

			if item == nil {
				t.Fatalf("ItemFor found no item")
			}
			if value != item.ValueExpr() {
				t.Errorf("ValueExprFor returned a different value than ItemFor")
			}
			if got := tokensText(value.Expr().BuildTokens(nil)); got != test.want {
				t.Errorf("wrong value\ngot:  %s\nwant: %s", got, test.want)
			}
		})
	}
}

func TestObjectConsExprItems(t *testing.T) {
	_, object := parseObjectAttr(t, "a = {\n  hat = 1\n  \"cat\" = 2\n  (dog) = 3\n}\n")

	var keys, values []string
	for _, item := range object.Items() {
		keys = append(keys, tokensText(item.KeyExpr().BuildTokens(nil)))
		values = append(values, tokensText(item.ValueExpr().BuildTokens(nil)))
	}
	if got, want := strings.Join(keys, " "), `hat "cat" (dog)`; got != want {
		t.Errorf("wrong keys\ngot:  %s\nwant: %s", got, want)
	}
	if got, want := strings.Join(values, " "), "1 2 3"; got != want {
		t.Errorf("wrong values\ngot:  %s\nwant: %s", got, want)
	}

	_, empty := parseObjectAttr(t, "a = {}\n")
	if got := empty.Items(); len(got) != 0 {
		t.Errorf("got %d items in an empty object", len(got))
	}
}

func TestObjectConsExprSetItemKey(t *testing.T) {
	tests := []struct {
		key  string
		want string // key as source text
	}{
		{"dog", `dog`},
		{"dog_1", `dog_1`},
		{"dog-1", `dog-1`},
		{"true", `true`},
		{"for", `"for"`},
		{"a hat", `"a hat"`},
		{"hat.band", `"hat.band"`},
		{"1hat", `"1hat"`},
		{"", `""`},
		{`say "hi"`, `"say \"hi\""`},
		{`a\b`, `"a\\b"`},
		{"${hat}", `"$${hat}"`},
		{"%{hat}", `"%%{hat}"`},
		{"a\nb", `"a\nb"`},
	}

	for _, test := range tests {
		t.Run(test.key, func(t *testing.T) {
			f, object := parseObjectAttr(t, "a = { hat = \"derby\" }\n")

			key, value := object.SetItemValue(test.key, cty.StringVal("bowler"))
			if got := tokensText(key.BuildTokens(nil)); got != test.want {
				t.Errorf("wrong key\ngot:  %s\nwant: %s", got, test.want)
			}

			item := object.ItemFor(test.key)
			if item == nil {
				t.Fatalf("ItemFor can't find the key that was just set")
			}
			if item.KeyExpr() != key || item.ValueExpr() != value {
				t.Errorf("ItemFor found a different item than the one that was set")
			}

			got := evalObjectAttr(t, f.Bytes())
			want := cty.ObjectVal(map[string]cty.Value{
				"hat":    cty.StringVal("derby"),
				test.key: cty.StringVal("bowler"),
			})
			if !got.RawEquals(want) {
				t.Errorf("wrong value\ngot:  %#v\nwant: %#v", got, want)
			}

			// Setting the key again replaces the value instead of adding an item.
			key2, value2 := object.SetItemValue(test.key, cty.StringVal("fedora"))
			if key2 != key || value2 != value {
				t.Errorf("setting the key again did not return the same item")
			}
			if got := len(object.Items()); got != 2 {
				t.Errorf("got %d items after setting the key again, want 2", got)
			}
		})
	}
}

func TestObjectConsExprSetItemReplace(t *testing.T) {
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
			"a = { \"hat\" = \"derby\" }\n",
			"hat",
			"a = { \"hat\" = \"bowler\" }\n",
		},
		{
			"a = { hat: \"derby\" }\n",
			"hat",
			"a = { hat : \"bowler\" }\n",
		},
		{
			"a = { \"$${hat}\" = \"derby\" }\n",
			"${hat}",
			"a = { \"$${hat}\" = \"bowler\" }\n",
		},
		{
			"a = {\n  hat = \"derby\" # comment\n  cat = \"calico\"\n}\n",
			"hat",
			"a = {\n  hat = \"bowler\" # comment\n  cat = \"calico\"\n}\n",
		},
		{
			// A key from a variable is not replaced, so a new item is added.
			"a = { (hat) = \"derby\" }\n",
			"hat",
			"a = { (hat) = \"derby\", hat = \"bowler\" }\n",
		},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, object := parseObjectAttr(t, test.src)
			key, value := object.SetItemValue(test.key, cty.StringVal("bowler"))

			if got := string(f.Bytes()); got != test.want {
				t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, test.want)
			}
			item := object.ItemFor(test.key)
			if item == nil || item.KeyExpr() != key || item.ValueExpr() != value {
				t.Errorf("SetItemValue did not return the item that ItemFor finds")
			}
		})
	}
}

func TestObjectConsExprSetItemHelpers(t *testing.T) {
	f, object := parseObjectAttr(t, "a = { hat = \"derby\" }\n")

	object.SetItemRaw("hat", Tokens{
		{Type: hclsyntax.TokenIdent, Bytes: []byte("upper")},
		{Type: hclsyntax.TokenOParen, Bytes: []byte("(")},
		{Type: hclsyntax.TokenIdent, Bytes: []byte("var")},
		{Type: hclsyntax.TokenDot, Bytes: []byte(".")},
		{Type: hclsyntax.TokenIdent, Bytes: []byte("hat")},
		{Type: hclsyntax.TokenCParen, Bytes: []byte(")")},
	})
	object.SetItemTraversal("cat", hcl.Traversal{
		hcl.TraverseRoot{Name: "var"},
		hcl.TraverseAttr{Name: "cat"},
	})
	object.SetItemValue("dog", cty.NumberIntVal(2))

	want := "a = { hat = upper(var.hat), cat = var.cat, dog = 2 }\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result\ngot:  %s\nwant: %s", got, want)
	}

	// A traversal set with SetItemTraversal is a variable of the object.
	var vars []string
	for _, traversal := range f.Body().GetAttribute("a").Expr().Variables() {
		vars = append(vars, tokensText(traversal.BuildTokens(nil)))
	}
	if got, want := strings.Join(vars, " "), "var.cat"; got != want {
		t.Errorf("wrong variables\ngot:  %s\nwant: %s", got, want)
	}
}

func TestObjectConsExprSetItemPanics(t *testing.T) {
	assertPanics := func(t *testing.T, f func()) {
		t.Helper()
		defer func() {
			t.Helper()
			if recover() == nil {
				t.Errorf("SetItem did not panic")
			}
		}()
		f()
	}

	src := "a = { hat = \"derby\", cat = { dog = 1 } }\nb = var.b\n"

	t.Run("attribute expression from the parser", func(t *testing.T) {
		f, object := parseObjectAttr(t, src)
		expr := f.Body().GetAttribute("b").Expr()
		assertPanics(t, func() { object.SetItem("hat", expr) })
		assertPanics(t, func() { object.SetItem("new", expr) })
	})

	t.Run("item value from the parser", func(t *testing.T) {
		_, object := parseObjectAttr(t, src)
		expr := object.ValueExprFor("hat").Expr()
		assertPanics(t, func() { object.SetItem("new", expr) })
	})

	t.Run("nested item value from the parser", func(t *testing.T) {
		_, object := parseObjectAttr(t, src)
		nested := object.ValueExprFor("cat").Expr().AsObjectConsExpr()
		expr := nested.ValueExprFor("dog").Expr()
		assertPanics(t, func() { object.SetItem("new", expr) })
	})

	t.Run("expression from an earlier SetItem", func(t *testing.T) {
		_, object := parseObjectAttr(t, src)
		expr := NewExpressionAbsTraversal(hcl.Traversal{hcl.TraverseRoot{Name: "var"}})
		object.SetItem("hat", expr)
		assertPanics(t, func() { object.SetItem("new", expr) })
		assertPanics(t, func() { object.SetItem("hat", expr) })

		expr = NewExpressionAbsTraversal(hcl.Traversal{hcl.TraverseRoot{Name: "var"}})
		object.SetItem("new", expr)
		assertPanics(t, func() { object.SetItem("other", expr) })
	})

	t.Run("replaced expression can be reused", func(t *testing.T) {
		f, object := parseObjectAttr(t, src)
		old := object.ValueExprFor("hat").Expr()
		object.SetItemValue("hat", cty.StringVal("bowler"))
		object.SetItem("new", old)

		want := "a = { hat = \"bowler\", cat = { dog = 1 }, new = \"derby\" }\nb = var.b\n"
		if got := string(f.Bytes()); got != want {
			t.Errorf("wrong result\ngot:  %s\nwant: %s", got, want)
		}
	})
}

func TestObjectConsExprSetItemLayout(t *testing.T) {
	tests := []struct {
		name string
		src  string
		keys []string
		want string
	}{
		{
			"single line",
			"a = { hat = \"derby\" }\n",
			[]string{"dog"},
			"a = { hat = \"derby\", dog = \"bowler\" }\n",
		},
		{
			"single line, empty",
			"a = {}\n",
			[]string{"dog"},
			"a = { dog = \"bowler\" }\n",
		},
		{
			"single line, empty, two items",
			"a = {}\n",
			[]string{"dog", "cat"},
			"a = { dog = \"bowler\", cat = \"bowler\" }\n",
		},
		{
			"single line, two items",
			"a = { hat = \"derby\" }\n",
			[]string{"dog", "cat"},
			"a = { hat = \"derby\", dog = \"bowler\", cat = \"bowler\" }\n",
		},
		{
			"single line, trailing comma",
			"a = { hat = \"derby\", }\n",
			[]string{"dog"},
			"a = { hat = \"derby\", dog = \"bowler\" }\n",
		},
		{
			"single line, inline comment before brace",
			"a = { hat = \"derby\" /* end of a */ }\n",
			[]string{"dog"},
			"a = { hat = \"derby\", dog = \"bowler\" /* end of a */ }\n",
		},
		{
			"single line, nested multi-line value",
			"a = { hat = {\n  b = 1\n} }\n",
			[]string{"dog"},
			"a = { hat = {\n  b = 1\n}, dog = \"bowler\" }\n",
		},
		{
			"single line, multi-line key",
			"a = { (\n  hat\n  ) = 1 }\n",
			[]string{"dog"},
			"a = { (\n  hat\n) = 1, dog = \"bowler\" }\n",
		},
		{
			"multi-line",
			"a = {\n  hat = \"derby\"\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}\n",
		},
		{
			"multi-line, two items",
			"a = {\n  hat = \"derby\"\n}\n",
			[]string{"dog", "cat"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n}\n",
		},
		{
			"multi-line, empty",
			"a = {\n}\n",
			[]string{"dog"},
			"a = {\n  dog = \"bowler\"\n}\n",
		},
		{
			"multi-line, key alignment",
			"a = {\n  hat = \"derby\"\n}\n",
			[]string{"a_long_key"},
			"a = {\n  hat        = \"derby\"\n  a_long_key = \"bowler\"\n}\n",
		},
		{
			"multi-line, brace on the last item's line",
			"a = {\n  hat = \"derby\" }\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}\n",
		},
		{
			"multi-line, brace on the last item's line, after a comma",
			"a = {\n  hat = \"derby\", }\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\",\n  dog = \"bowler\"\n}\n",
		},
		{
			"multi-line, from the example",
			"a = {\n    hat = \"derby\", (cat) = \"calico\" }\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\", (cat) = \"calico\"\n  dog = \"bowler\"\n}\n",
		},
		{
			"multi-line, trailing comma",
			"a = {\n  hat = \"derby\",\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\",\n  dog = \"bowler\"\n}\n",
		},
		{
			"multi-line, line comment",
			"a = {\n  hat = \"derby\" # the hat\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\" # the hat\n  dog = \"bowler\"\n}\n",
		},
		{
			"multi-line, comma and line comment",
			"a = {\n  hat = \"derby\", // the hat\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\", // the hat\n  dog = \"bowler\"\n}\n",
		},
		{
			"multi-line, comments before the brace",
			"a = {\n  hat = \"derby\"\n  # end of a\n  # really\n}\n",
			[]string{"dog", "cat"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n  # end of a\n  # really\n}\n",
		},
		{
			"multi-line, inline comment before the brace on its own line",
			"a = {\n  hat = \"derby\"\n  /* end of a */ }\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n/* end of a */ }\n",
		},
		{
			"multi-line, inline comment before the brace on the last item's line",
			"a = {\n  hat = \"derby\" /* end of a */ }\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n/* end of a */ }\n",
		},
		{
			"multi-line, nested multi-line value",
			"a = {\n  hat = {\n    b = 1\n  }\n}\n",
			[]string{"dog"},
			"a = {\n  hat = {\n    b = 1\n  }\n  dog = \"bowler\"\n}\n",
		},
		{
			"multi-line, in a block",
			"b {\n  a = {\n    hat = \"derby\"\n  }\n}\n",
			[]string{"dog"},
			"b {\n  a = {\n    hat = \"derby\"\n    dog = \"bowler\"\n  }\n}\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			body := f.Body()
			if blocks := body.Blocks(); len(blocks) > 0 {
				body = blocks[0].Body()
			}
			object := body.GetAttribute("a").Expr().AsObjectConsExpr()

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
			if _, diags := hclsyntax.ParseConfig(got, "", hcl.InitialPos); diags.HasErrors() {
				t.Errorf("result is not valid HCL: %s", diags.Error())
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

// TupleConsExpr is a tuple-construct expression, such as ["fedora", "fez"].
// Use Expression.AsTupleConsExpr to get one from an expression.
//
// Each element is its own Expression. The brackets and the commas are
// unstructured tokens of the tuple.
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

// variables returns the absolute traversals in the elements of the tuple,
// in source order.
func (tuple *TupleConsExpr) variables() []*Traversal {
	return variablesOf(tuple.elems)
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

func ExampleObjectConsExpr_SetItemRaw() {
	src := []byte("a = {\n    hat = \"derby\", (cat) = \"calico\" }\n")
	f, diags := hclwrite.ParseConfig(src, "", hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		fmt.Printf("errors: %s", diags)
		return
	}

	bowler := hclwrite.TokensForValue(cty.StringVal("bowler"))
	object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

	// Replace the value of an existing item, then add a new one.
	key, value := object.SetItemRaw("hat", bowler)
	fmt.Printf("%s = %s\n", key.BuildTokens(nil).Bytes(), value.BuildTokens(nil).Bytes())
	object.SetItemRaw("dog", bowler)

	fmt.Printf("%s", f.Bytes())
	// Output:
	// hat = "bowler"
	// a = {
	//   hat = "bowler", (cat) = "calico"
	//   dog = "bowler"
	// }
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

	n := newNode(attr)
	attr.Expr().parent = n
	return n
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
		value.Expr().parent = item.value

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
	wrapExpr := newObjectConsKeyExpr(expr)
	// Use the same rule as hclsyntax.ObjectConsKeyExpr to decide whether
	// the key is a literal name instead of an expression.
	wrapExpr.literal = hcl.ExprAsKeyword(nativeExpr.Wrapped) != ""

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
// Each argument becomes its own Expression. The tokens around the arguments
// stay as unstructured tokens of the call.
func parseFunctionCallExpr(nativeExpr *hclsyntax.FunctionCallExpr, from inputTokens) *node {
	expr := newFunctionCallExpr(nativeExpr.Name)
	return parseExprList(expr, expr.children, expr.args, nativeExpr.Args, from)
}

// parseTupleConsExpr parses a tuple-construct expression, defined as:
//
//	tuple = "[" (
//	    (Expression ("," Expression)* ","?)?
//	) "]";
//
// Each element becomes its own Expression. The tokens around the elements
// stay as unstructured tokens of the tuple.
func parseTupleConsExpr(nativeExpr *hclsyntax.TupleConsExpr, from inputTokens) *node {
	expr := newTupleConsExpr()
	return parseExprList(expr, expr.children, expr.elems, nativeExpr.Exprs, from)
}

// parseExprList parses each of nativeExprs into an Expression, appends it to
// children and adds it to set. The tokens before, between and after them are
// appended to children as unstructured tokens. It returns a node with an
// Expression that wraps content, which must own children.
func parseExprList(content nodeContent, children *nodes, set nodeSet, nativeExprs []hclsyntax.Expression, from inputTokens) *node {
	var before, exprTokens inputTokens
	for _, nativeExpr := range nativeExprs {
		before, exprTokens, from = from.Partition(nativeExpr.Range())
		children.AppendUnstructuredTokens(before.writerTokens)
		n := parseExpression(nativeExpr, exprTokens)
		children.AppendNode(n)
		set.Add(n)
	}
	children.AppendUnstructuredTokens(from.writerTokens)

	wrapExpr := newExpression()
	wrapExpr.wrap(content)
	for n := range set {
		n.content.(*Expression).parent = wrapExpr.wrapped
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
			`a = join(" ", ["fedora", "fez"])`,
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
														Val:  `" "`,
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
																		Val:  `"fedora"`,
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
																		Val:  ` "fez"`,
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
						},
					},
				},
			},
		},
		{
			`a = [
  hat.brim,
  f(),
  { b = 1 },
  g(c...),
]`,
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
														Type: "Traversal",
														Children: []TestTreeNode{
															{
																Type: "TraverseName",
																Children: []TestTreeNode{
																	{
																		Type: "identifier",
																		Val:  "  hat",
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
																		Val:  "brim",
																	},
																},
															},
														},
													},
												},
											},
											{
												Type: "Tokens",
												Val:  ",\n",
											},
											{
												Type: "Expression",
												Children: []TestTreeNode{
													{
														Type: "FunctionCallExpr",
														Children: []TestTreeNode{
															{
																Type: "Tokens",
																Val:  "  f()",
															},
														},
													},
												},
											},
											{
												Type: "Tokens",
												Val:  ",\n",
											},
											{
												Type: "Expression",
												Children: []TestTreeNode{
													{
														Type: "ObjectConsExpr",
														Children: []TestTreeNode{
															{
																Type: "Tokens",
																Val:  "  {",
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
																										Val:  " b",
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
																						Type: "Tokens",
																						Val:  " 1",
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
												Val:  ",\n",
											},
											{
												Type: "Expression",
												Children: []TestTreeNode{
													{
														Type: "FunctionCallExpr",
														Children: []TestTreeNode{
															{
																Type: "Tokens",
																Val:  "  g(",
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
																						Val:  "c",
																					},
																				},
																			},
																		},
																	},
																},
															},
															{
																Type: "Tokens",
																Val:  "...)",
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

#### pkg/round_trip_test.go

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
		`a = join(" ", ["fedora", "fez"])
b = [
  f(x, [1, 2]...),
  { k = g() },
  [], # empty
]
c = provider::local::upper(lower("HAT"), ) # trailing comma
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

Answer with one JSON object and nothing else:
{"findings": [{"file": "...", "line": 0, "problem": "...", "fix": "..."}]}
Give each finding a concrete file, line and fix. Leave out vague remarks.
An empty list is a valid answer.
