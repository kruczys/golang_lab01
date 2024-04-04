package main

import (
    "fmt"
    "math/big"
    "strings"
)

func factorial(number int) *big.Int {
    result := big.NewInt(1)
    for i := 1; i < number; i++ {
        result.Mul(result, big.NewInt(int64(i)))
    }
    return result
}

func fib(number uint64) uint64 {
    if number <= 1 {
        return number
    }
    return fib(number - 1) + fib(number - 2)
}

func containsSubstring(container, containee string) bool {
    return strings.Contains(container, containee)
}

func generateNick(name, surname string) string {
    return strings.ToLower(name[:3] + surname[1:4])
}

func generateAscii(nick string) [6]int {
    var result [6]int
    for i := 0; i <= 5; i++ {
        result[i] = int(nick[i])
    }
    return result
}

func main() {
    test := factorial(123)
    fmt.Printf("%s\n", test.String())
    fmt.Println(containsSubstring(test.String(), `000`))
    nick := generateNick(`Konrad`, `Kreczko`)
    fmt.Println(nick)
    fmt.Println(generateAscii(nick))
    fmt.Println(fib(3))
    fmt.Println(fib(4))
    fmt.Println(fib(5))
    fmt.Println(fib(40))
}
