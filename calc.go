package main

import (
    "fmt"
    "os"
    "strconv"
    "math"
)


// Evaluate float constant, unary expression, or paren expression
func evalUnaryExpr(tokens []Token) (float64, []Token, error) {
    if len(tokens) == 0 {
        return 0, nil, fmt.Errorf("expecting number")
    }

    switch tokens[0].Type {
    case FLOAT:
        n, err := strconv.ParseFloat(tokens[0].Value, 64)
        if err != nil {
            return 0, nil, fmt.Errorf("invalid number '%s'", tokens[0].Value)
        }
        return n, tokens[1:], nil
    case ADD:
        n, tokens, err := evalExpr(tokens[1:], precs[POS])
        if err != nil {
            return 0, nil, err
        }
        return n, tokens, nil
    case SUB:
        n, tokens, err := evalExpr(tokens[1:], precs[NEG])
        if err != nil {
            return 0, nil, err
        }
        return -n, tokens, nil
    case LPAREN:
        n, tokens, err := evalExpr(tokens[1:], 0)
        if err != nil {
            return 0, nil, err
        }
        if len(tokens) == 0 || tokens[0].Type != RPAREN {
            return 0, nil, fmt.Errorf("missing right parenthesis")
        }
        return n, tokens[1:], nil
    }
    return 0, nil, fmt.Errorf("expecting number, got '%s'", tokens[0].Value)
}



func evalOp(op TokenType, left, right float64) (float64, error) {
    switch op {
    case ADD:
        return left + right, nil
    case SUB:
        return left - right, nil
    case MUL:
        return left * right, nil
    case DIV:
        return left / right, nil
    case MOD:
        return math.Mod(left, right), nil
    case POW:
        return math.Pow(left, right), nil
    }
    return 0, fmt.Errorf("unknown operator")
}



// Evaluate expression using precedence climbing algorithm
func evalExpr(tokens []Token, minPrec int) (float64, []Token, error) {
    left, tokens, err := evalUnaryExpr(tokens)
    if err != nil {
        return 0, nil, err
    }

    for len(tokens) != 0 {
        op := tokens[0].Type
        if !isOp(op) {
            break
        }
        prec := precs[op]
        if prec < minPrec {
            break
        }

        // if operator is right associative, don't add 1 to the current precedence
        nextMinPrec := prec + 1
        if op == POW {
            nextMinPrec = prec
        }

        // consume operator token and get right subtree
        right, nextTokens, err := evalExpr(tokens[1:], nextMinPrec)
        if err != nil {
            return 0, nil, err
        }

        left, err = evalOp(op, left, right)
        if err != nil {
            return 0, nil, err
        }
        tokens = nextTokens
    }
    return left, tokens, nil
}



func calc(s string) (float64, error) {
    tokens, err := lexer([]byte(s))
    if err != nil {
        return 0, err
    }
    n, tokens, err := evalExpr(tokens, 0)
    if err != nil {
        return 0, err
    }
    if len(tokens) != 0 {
        return 0, fmt.Errorf("invalid expression")
    }
    return n, nil
}



var usage string = `Usage of calc:
  calc expression

Description:
  calc is a tool that evaluates float64 expressions on the command line.

  The following operators are supported, from highest precedence to lowest:

  Operator | Description                         | Associativity
  ---------+-------------------------------------+--------------
  ^        | Exponentiation                      | Right to left
  + -      | Positive, negative                  | Right to left
  * / %    | Multiplication, division, remainder | Left to right
  + -      | Addition, subtraction               | Left to right

  Parenthesis can be used to group subexpressions as usual.

Examples:
  Basic expression:
    calc '1+2*3'
    7
  Parentheses:
    calc '(1+2) * 3'
    9
  Floating-point:
    calc '0.5 + 0.25'
    0.75
  Exponentiation:
    calc '-2^2'
    -4
  Remainder:
    calc '-5 % 3'
    -2`



func main() {
    args := os.Args

    if len(args) != 2 {
        fmt.Fprintln(os.Stderr, "error: expecting 1 argument")
        os.Exit(1)
    }
    if args[1] == "-h" || args[1] == "--help" {
        fmt.Fprintln(os.Stderr, usage)
        os.Exit(1)
    }

    result, err := calc(args[1])
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }

    fmt.Println(result)
}
