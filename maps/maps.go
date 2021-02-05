package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("Could not find the key you asked for.")

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "%q", ErrNotFound
	}
	return value, nil
}
