package dictionary

import "errors"

type Dictionary map[string]string

var ErrorWordNotFound = errors.New("Word not found")

func (d Dictionary) Search(s string) (string, error) {
	value, success := d[s]
	if !success {
		return "", ErrorWordNotFound
	}
	return value, nil
}

func (d Dictionary) AddWord(key, value string) {
	d[key] = value
}
