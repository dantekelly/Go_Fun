package main

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("could not update the word you were looking for as it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(s string) (string, error) {
	definition, ok := d[s]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(p string, s string) error {
	_, err := d.Search(p)

	switch err {
	case ErrNotFound:
		d[p] = s
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(p string, s string) error {
	_, err := d.Search(p)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[p] = s
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(p string) {
	delete(d, p)
}
