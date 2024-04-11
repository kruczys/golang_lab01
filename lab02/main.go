package main

import (
    "fmt"
    "math/big"
    "strings"
    "strconv"
)

func factorial(number int) *big.Int {
    result := big.NewInt(1)
    for i := 1; i <= number; i++ {
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

func containsSubstring(container, containee string) int {
    if strings.Contains(container, containee) {
        return 1
    }
    return 0
}

func generateNick(name, surname string) string {
    return strings.ToLower(name[:3] + surname[1:4])
}

func generateAsciiCodesAsStrings(nick string) [6]string {
    var result [6]string
    for i := 0; i <= 5; i++ {
        result[i] = strconv.Itoa(int(nick[i]))
    }
    return result
}

func isStrongNumber(asciiCodes [6]string, number *big.Int) bool {
    controlSum := 0
    numberAsStr := number.String()
    
    for i := 0; i < 6; i++ {
        controlSum += containsSubstring(numberAsStr, asciiCodes[i])
    }
    
    return controlSum == 6
}

func findStrongNumber(name, surname string) int {
    nick := generateNick(name, surname)
    asciiCodes := generateAsciiCodesAsStrings(nick)
    for i := 0 ; ; i++ {
        if isStrongNumber(asciiCodes, factorial(i)) {
            return i
        }
    }
    return 0
}

func main() {
    fmt.Println(findStrongNumber("Piotr", "Karl"))
}
