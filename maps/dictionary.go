package maps

// Dictionary is a custom type of map
type Dictionary map[string]string

// Search checks a dictionary for words
func (d Dictionary) Search(word string) string {
	return d[word]
}
