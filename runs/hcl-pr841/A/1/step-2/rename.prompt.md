The test files listed below do not compile against the code in this directory.
They were written against an earlier implementation of the same behaviour, which
used different names.

Edit ONLY these test files so that they compile against the current code:

- hclwrite/ast_attribute_test.go
- hclwrite/ast_block_test.go
- hclwrite/ast_body_test.go
- hclwrite/ast_expression_test.go
- hclwrite/ast_object_cons_expr_test.go
- hclwrite/ast_test.go
- hclwrite/examples_test.go
- hclwrite/format_test.go
- hclwrite/generate_test.go
- hclwrite/parser_test.go
- hclwrite/round_trip_test.go

Rules:
- You may change identifiers: function, method, type, field and constant names,
  and how arguments are passed, so that each call reaches the code that does the
  same job.
- Never change a literal value, an expected result, an assertion, or what a test
  checks. Never add or remove checks.
- Do not edit any other file.
- If a test function has no counterpart in the current code, for example because
  the helper it calls does not exist as a separate function, delete that test
  function and list its name, one per line, in a file named NO_COUNTERPART.txt
  in this directory.

Run `go test -c -o /dev/null ./hclwrite` until it compiles. You do not need the tests to pass.
