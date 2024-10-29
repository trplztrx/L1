package main

import (
    "fmt"
    "strings"
)

func HasUniqueChars(s string) bool {
    s = strings.ToLower(s)
    chars := make(map[rune]struct{})

    for _, char := range s {
        if _, exists := chars[char]; exists {
            return false
        }
        chars[char] = struct{}{}
    }
    return true
}

func main() {
    tests := []string{"abcd", "abCdefAaf", "aabcd"}
    for _, test := range tests {
        fmt.Printf("%s — %v\n", test, HasUniqueChars(test))
    }
}
