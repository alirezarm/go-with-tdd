package maps

// Dictionary type is a thin wrapper around the map type
type Dictionary map[string]string

// Create DictionaryErr type which implements the error interface
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

var (
	ErrWordNotFound = DictionaryErr("could not find the word")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrWordNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrWordNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	// In case Search() returns an error other than ErrWordNotFound
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrWordNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	// In case Search() returns an error other than ErrWordNotFound
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrWordNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	// In case Search() returns an error other than ErrWordNotFound
	default:
		return err
	}
	return nil
}