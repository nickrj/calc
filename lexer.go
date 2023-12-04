package main

import "fmt"


type TokenType int

const (
    FLOAT TokenType = iota
    op_start
    ADD         // +
    SUB         // -
    MUL         // *
    DIV         // /
    MOD         // %
    POS         // +
    NEG         // -
    POW         // ^
    op_end
    LPAREN      // (
    RPAREN      // )
)


type Token struct {
    Type  TokenType
    Value string
}


var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}


var precs = [...]int{
    ADD: 1,
    SUB: 1,
    MUL: 2,
    DIV: 2,
    MOD: 2,
    POS: 3,
    NEG: 3,
    POW: 4,
}


func isOp(t TokenType) bool {
    return op_start < t && t < op_end
}



func eatSpaces(b []byte) []byte {
    i := 0
    for ; i < len(b); i++ {
        if asciiSpace[b[i]] == 0 {
            break
        }
    }
    return b[i:]
}



func parseFloatToken(b []byte) (Token, []byte, error) {
    i := 0
    for ; i < len(b); i++ {
        c := b[i]
        if (c < '0' || c > '9') && c != '.' {
            break
        }
    }
    if i == 0 {
        return Token{}, nil, fmt.Errorf("invalid float")
    }
    return Token{FLOAT, string(b[:i])}, b[i:], nil
}



func lexer(b []byte) (tokens []Token, err error) {
    b = eatSpaces(b)

    for len(b) != 0 {
        var token Token

        c := b[0]
        if ('0' <= c && c <= '9') || c == '.' {
            token, b, err = parseFloatToken(b)
            if err != nil {
                return nil, err
            }
        } else {
            switch c {
            case '+':
                token = Token{ADD, "+"}
            case '-':
                token = Token{SUB, "-"}
            case '*':
                token = Token{MUL, "*"}
            case '/':
                token = Token{DIV, "/"}
            case '%':
                token = Token{MOD, "%"}
            case '^':
                token = Token{POW, "^"}
            case '(':
                token = Token{LPAREN, "("}
            case ')':
                token = Token{RPAREN, ")"}
            default:
                return nil, fmt.Errorf("unknown character '%c'", c)
            }
            b = b[1:]
        }
        tokens = append(tokens, token)
        b = eatSpaces(b)
    }
    return tokens, nil
}
