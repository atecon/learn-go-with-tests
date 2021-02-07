package main

import (
	"errors"
)

type Dictionary map[string]string

var (
	ErrNotFound  = errors.New("Could not find the key you asked for.")
	ErrKeyExists = errors.New("Key already exists. Value cannot be overwritten")
)

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
