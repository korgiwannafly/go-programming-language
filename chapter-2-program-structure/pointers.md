### Pointers
A pointer value is the address of a variable

- `x := 1`
- `p := &x         // p, of type *int, points to x`
- `fmt.Println(*p) // "1"`
- `*p = 2          // equivalent to x = 2`
- `fmt.Println(x)  // "2"`

Variables are sometimes described as addressable values. Expressions that denote variables are the only expressions to which the address-of operator & may be applied.
The zero value for a pointer of any type is nil. The test p != nil is true if p points to a vari- able. Pointers are comparable; two pointers are equal if and only if they point to the same variable or both are nil.

- `var x, y int`
- `fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"`
