package main

type Dictionary map[string]string

type DictError string

const (
	ErrNotFound         = DictError("Could not find the key you asked for.")
	ErrKeyExists        = DictError("Key already exists. Value cannot be overwritten")
	ErrWordDoesNotExist = DictError("cannot update word because it does not exist")
)

func (d DictError) Error() string {
	return string(d)
}

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "%q", ErrNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}

func (d Dictionary) AddWithCheck(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrKeyExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, newValue string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = newValue
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
