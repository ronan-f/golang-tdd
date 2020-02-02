package maps

import "errors"

// Dictionary is a custom type of map
type Dictionary map[string]string

// Search checks a dictionary for words
func (d Dictionary) Search(word string) (string, error) {

	definition, found := d[word]

	if !found {
		return "", errors.New("Could not find the word you were looking for")
	}

	return definition, nil
}
