package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string


var (
	errNotFound = errors.New("not found")
	errWordExists = errors.New("that word already exists")
	errCantUpdate = errors.New("cant update non-existing word")
)

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

//update word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

//delete word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}