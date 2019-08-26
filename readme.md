# Goflat

**Goflat** is a small utility library for flattening nested integer arrays.

## Example

**Nested Array**

[0, [2], [[5], 0]]

**Flattened Array**

[0, 2, 5, 0]

## Usage

```
input := []interface{}{0, []interface{}{2}, []interface{}{[]interface{}{5}, 0}}
output := goflat.Flatten(input)
    
// output [0, 2, 5, 0]
println(output)
```
