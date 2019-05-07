# argparse

Super simple library for parsing strings into a slice of string arguments.

```go
input := `this "is a" test`
output := argparse.Parse(input)
fmt.Printf("%#v", output)
// []string{"this", "is a", "test"}
```

Supports escaping quotes and spaces with backslash:

```go
input := `this \"is\" a\ test`
output := argparse.Parse(input)
fmt.Printf("%#v", output)
// []string{"this", "\"is\"", "a test"}
```
