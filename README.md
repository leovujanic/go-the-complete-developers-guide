# Hands-on Go lang






To format code use: `gofmt -w .`


## Method declaration

You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).


## Type Switches

```go
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```


## `toString()` or `Stringer` in go

```go

func (t T) String() string {
	return fmt.Sprintf("string representation %s", t.prop)
}

```