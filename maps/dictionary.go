package maps

// Search checks a dictionary for words
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
