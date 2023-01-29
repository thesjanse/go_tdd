package dictionary

import "errors"

type Dictionary map[string]string

var ErrorWordNotFound = errors.New("Word not found")
var ErrorWordAlreadyExists = errors.New("Word already exists")

func (d Dictionary) Search(s string) (string, error) {
	value, success := d[s]
	if !success {
		return "", ErrorWordNotFound
	}
	return value, nil
}

func (d Dictionary) AddWord(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrorWordNotFound:
		d[key] = value
	case nil:
		return ErrorWordAlreadyExists
	default:
		return err
	}
	return nil
}
