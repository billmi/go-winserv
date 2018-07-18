package helper

import "errors"

// InArray is
func InArray(needle string, haystack []string) (int, error) {
	for k, v := range haystack {
		if needle == v {
			return k, nil
		}
	}
	return -1, errors.New("notFound")
}
