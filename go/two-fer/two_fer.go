package twofer

import "fmt"

// ShareWith prints the sentence "One for $name, one for me"
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
