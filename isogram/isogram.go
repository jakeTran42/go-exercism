package isogram

import "strings"

func IsIsogram(word string) bool {
    charSet := make(map[string]bool)

    for _, char := range word {
        // Change the character to an uppercase string
        character := strings.ToUpper(string(char))
        if character == " " || character == "-" {
            // Catch spaces and hyphens
        } else if _, inMap := charSet[character]; inMap {
            return false
        } else {
            charSet[character] = true
        }
    }

    return true
}
