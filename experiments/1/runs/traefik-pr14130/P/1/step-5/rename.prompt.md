The test files listed below do not compile against the code in this directory.
They were written against an earlier implementation of the same behaviour, which
used different names.

Edit ONLY these test files so that they compile against the current code:

- pkg/provider/kubernetes/gateway/annotations_test.go
- pkg/provider/kubernetes/gateway/client_test.go
- pkg/provider/kubernetes/gateway/grpcroute_test.go
- pkg/provider/kubernetes/gateway/httproute_test.go
- pkg/provider/kubernetes/gateway/kubernetes_test.go
- pkg/provider/kubernetes/gateway/status_test.go
- pkg/provider/kubernetes/gateway/tlsroute_test.go

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

Run `go test -c -o /dev/null ./pkg/provider/kubernetes/gateway` until it compiles. You do not need the tests to pass.
