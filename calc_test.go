package main

import (
    "testing"
    "math"
)



func TestCalc(t *testing.T) {
    var tests = []struct{
        input string
        want  float64
    }{
        {"1 + 2 * 3", 7},
        {"7 - 9 * (2 - 3)", 16},
        {"2 * 3 * 4", 24},
        {"4^3^2^1", 262144},
        {"-4^3^2^1", -262144},
        {"(2 ^ 3) ^ 4", 4096},
        {"2^(1+2)^2", 512},
        {"2^(1+2)^(1+2)^-1", 2.7174426159984844},
        {"2^-1+3", 3.5},
        {"5", 5},
        {"2^-2", 0.25},
        {"-2^-2", -0.25},
        {"-2^2", -4},
        {"--2^2", 4},
        {"-3^2^1", -9},
        {"5%3", 2},
        {"-5%3", -2},
        {"4 + 2", 6},
        {"9 - 8 - 7", -6},
        {"9 - (8 - 7)", 8},
        {"(9 - 8) - 7", -6},
        {"2 + 3 ^ 2 * 3 + 4", 33},
        {"(1 * 2^32) - 1", 4294967295},
        {"(1*2^-1)+(1*2^-2)+(1*2^-3)", 0.875},
        {"((2^53)-1) * (2^(1023-52))", math.MaxFloat64},
        {"0.1+0.2", 0.30000000000000004},
        {"+-1", -1},
        {"-+1", -1},
        {"(1+2)*3", 9},
        {"1-2*3^4", -161},
        {"1+2^(3+4)", 129},
        {"1+2^3+4", 13},
        {"(1+3)^0.5+5", 7},
        {"-+3*+8+9/7+3/---6^-7%+4%--8/1-5*++1^--+7^-+-4*--4/++-5/+-3+--3^--+4%-9", -24.047619047619047},
        {"((((((((((1))))))))))", 1},
    }
    for _, test := range tests {
        got, err := calc(test.input)
        if err != nil {
            t.Errorf(`calc("%s") failed: %v`, test.input, err)
        }
        if got != test.want {
            t.Errorf(`calc("%s") = %g, want %g`, test.input, got, test.want)
        }
    }
}


func TestCalcFail(t *testing.T) {
    var tests = []string {
        "1+2)",
        "(1+2",
        "1+",
        "1+)",
        "1+(",
        "1+.",
        "-1-",
        "1+a",
        ".",
        "..",
        "1..2",
        "0.f",
        "127.0.0.1",
        "test",
    }
    for _, test := range tests {
        if _, err := calc(test); err == nil {
            t.Errorf(`calc("%s") returned nil error`, test)
        }
    }
}
