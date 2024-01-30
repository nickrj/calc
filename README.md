# calc
calc is a tool that lets you evaluate floating-point expressions on the command line.

The following operators are supported, from highest precedence to lowest:

| Operator | Description                         | Associativity
| ---------|-------------------------------------|--------------
| ^        | Exponentiation                      | Right to left
| + -      | Positive, negative                  | Right to left
| * / %    | Multiplication, division, remainder | Left to right
| + -      | Addition, subtraction               | Left to right

Parenthesis can be used to group subexpressions as usual.

## Examples
```bash
$ calc '1 + 2 * 3'
7

$ calc '(1+2) * 3'
9

$ calc '0.5 + 0.25'
0.75

$ calc '-2^2'
-4

$ calc '-5 % 3'
-2
```

## Installation
Make sure you have [Go](https://go.dev/) installed on your computer.
```bash
go install github.com/nickrj/calc@latest
```

## References
- [Parsing Expressions by Recursive Descent](https://www.engr.mun.ca/~theo/Misc/exp_parsing.htm#climbing)
- [Parsing expressions by precedence climbing](https://eli.thegreenplace.net/2012/08/02/parsing-expressions-by-precedence-climbing)
- [Operator-precedence parser](https://en.wikipedia.org/wiki/Operator-precedence_parser#Precedence_climbing_method)
