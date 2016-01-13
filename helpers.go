package gorma

import "strings"

// deModel removes the word "Model" from the string.
func deModel(s string) string {
	return strings.Replace(s, "Model", "", -1)
}
