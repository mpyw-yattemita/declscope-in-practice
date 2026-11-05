# type checker: identify a compatible type from a reference

The type checker misses a wrong argument type when the value comes from a reference:

```rego
package p

a := "foo"
s := sum([1, a])
```

This should fail, since `sum` only accepts numbers, but no error is reported. When the value is written directly, the error is reported:

```rego
package p

s := sum([1, "foo"])
```

Please fix the built-in argument check so that an argument whose type is already known (for example through a reference to a rule, or a local variable assigned in the same rule body) is checked against the expected argument type. Both programs must report `sum: invalid argument(s)`, and so must

```rego
allow := s if {
    a := "foo"
    s := sum([1, a])
}
```

Arguments that are valid must still type-check: `a := 1` followed by `s := sum([1, a])`, and `s := sum([1, input.a])` where `input.a` has no known type.

While there, fix the comparison of partial objects with static keys. Checking that an object type unifies with another must not fail just because a key of one object is missing from the other's static keys, when the object has dynamic values. These policies must compile without errors:

```rego
package p

obj["a"] := input.a
obj["b"] := input.b

test_obj2 if {
	{"a":"1"} == obj with input as {"a": "1"}
}
```

```rego
package p

obj1["a"] := input.a
obj1["b"] := input.b

obj2["x"] := input.x
obj2["y"] := input.y

test if {
	obj1 == obj2
}
```

```rego
package test

obj["a"] := true

obj["b"] := "foo" if {
    input.foo == "bar"
}

test_obj if {
    obj == {
        "a": true
    } with input as {"foo": "baz"}
}
```

Add type checker tests for these cases.
