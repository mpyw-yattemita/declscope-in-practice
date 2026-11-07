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
