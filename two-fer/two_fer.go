// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

// Iteration number
const version = 3

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	if name == "" {
		name = "me"
	}

	return "One for " + name + ", one for you"
}
