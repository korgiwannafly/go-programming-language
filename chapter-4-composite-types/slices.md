### Slices in GO

- The zero value of a slice type is nil
- `var t []int        // len(s) == 0, s == nil`
- `t = nil            // len(s) == 0, s == nil`
- `t = []int(nil)     // len(s) == 0, s == nil`
- `t = []int{}        // len(s) == 0, s != nil`

- The built-in function make creates a slice of a specified element type, length, and capacity
- `make([]T, len)`
- `make([]T, len, cap) // same as make([]T, cap)[:len]`
