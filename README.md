# kv-options

Simple, typed key-value parser for Go. Currently supports options of type `bool` and `string`.

## Usage

Pull the library into your Go modules:

```bash
go get github.com/jannfis/kv-options
```

Import in your Go code:

```go
import "github.com/jannfis/kv-options/pkg/options"
```

## License

BSD zero-clause (0BSD). Refer to [LICENSE](LICENSE).

## Examples

### Parsing arbitrary, single key/value pairs of known types

You can parse arbitrary key value pairs if you impose their types. Simply create a new options handler. By default, all options will be valid.

```go 
// Create a new option set
opts := options.New()

// Parse a boolean option
err := opts.ParseBool("Foo=true")
if err != nil {
    // handle error
}

// Parse a string option
err = opts.ParseString("Bar=baz")
if err != nil {
    // handle error
}
```

### Parsing a slice of key/value pairs of different types

The parser can be fed a slice of key/value pairs, each of which having a differenty type. However, the valid options and their types need to be declared before:

```go
// Create a new option set
opts := options.New()

// Declare option Foo of type Bool
opts = opts.WithOption("Foo", options.OptionTypeBool)

// Declare option Bar of type String
opts = opts.WithOption("Bar", options.OptionTypeString)

// Parse a slice of options
err := opts.Parse([]string{"Foo=true", "Bar=whatever"})
if err != nil {
    // Handle error
}
```

### Accessing values

Once the key/value pairs are parsed, you can access their values.

You can specify a default value, which will be returned if the requested option has not been set or is of a different type. The default needs to be passed as a pointer to the value. The library provides the `options.Ptr()` helper function to get a pointer for a literal such as a `"string"` or a symbol such as `true`.

If the default is given as `nil`, the option is expected to be set (i.e. mandatory):

```go
bval, err := opts.Bool("Foo", nil)
if err != nil {
    // handle error
}

sval, err := opts.String("Bar", nil)
if err != nil {
    // handle error
}
```

In the previous example, `Bool` will return an error if the requested option is either not set or of a different type.

If you provide a default value and `Foo` was never defined, `bval` will be `true` and `err` will be `nil`:

```go
bval, err := opts.Bool("Bar", options.Ptr(true))
if err != nil {
    // handle error
}

sval, err := opts.String("Foo", options.Ptr("bar"))
if err != nil {
    // handle error
}
```

Note that the default value will only be returned by `Bool` if the option has not been defined. If 

If you're not interested in error handling, you can use the getters with a `Must` prefix and provide a mandatory default value to be returned when the option is not set, or is of a different type:

```go
bval := opts.MustBool("Foo", true)
sval := opts.MustString("Bar", "bar")
```
