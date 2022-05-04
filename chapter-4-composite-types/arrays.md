### Arrays in GO

array of 3 integers
- `var a [3]int`

In an array literal, if an ellipsis ‘‘...’’ appears in place of the length, the array length is deter- mined by the number of initializers. The definition of q can be simplified to
- `var b = [...]int{1, 2, 3}`

n this form, indices can appear in any order and some may be omitted; as before, unspecified values take on the zero value for the element type. For instance
- `var c = [...]int{1: 1, 2: 2}`
- `var d = [...]int{1, 2, 4: 5, 6}`
- `var e = [...]int{99: -1}`
