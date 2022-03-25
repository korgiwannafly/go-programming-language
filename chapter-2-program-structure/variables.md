### Variables in GO

- `var name type = expression`
- `var s string`
- `fmt.Printf(s) // ""`
- `var i, j, k int // int int int`
- `var b, f, s = true, 0.2, 'four' // boolean float64 string`

A set of variables can also be initialized by calling a function that returns multiple values:     
- `var f, err = os.Open(name) // os.Open returns a file and an error`
