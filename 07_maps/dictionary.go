package dictionary

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrorWordNotFound      = DictionaryErr("Word not found")
	ErrorWordDoesNotExist  = DictionaryErr("Word doesn't exist")
	ErrorWordAlreadyExists = DictionaryErr("Word already exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(s string) (string, error) {
	value, success := d[s]
	if !success {
		return "", ErrorWordNotFound
	}
	return value, nil
}

// add replace definition functionality
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

func (d Dictionary) UpdateWord(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrorWordNotFound:
		return ErrorWordDoesNotExist
	case nil:
		d[key] = value
	default:
		return err
	}
	return nil
}
