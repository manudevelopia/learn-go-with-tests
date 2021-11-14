package maps

import "errors"

type Dictionary map[string]string

var ErrorNotFound = errors.New("could not find the word you were looking for")

func (dictionary Dictionary) Search(word string) (string, error) {
	definition, ok := dictionary[word]
	if ok {
		return definition, nil
	}
	return "", ErrorNotFound
}

func (dictionary Dictionary) Add(word, definition string) {
	dictionary[word] = definition
}
