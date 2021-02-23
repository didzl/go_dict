package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string


var errNotFound = errors.New("not found")
var errWordExists = errors.New("that word already exists")

// search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists:= d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}


// add a word
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err ==nil {
		return errWordExists
	}

	return nil
}