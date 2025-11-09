package errhandle

import "os"

func ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err // return the error value
	}
	return data, nil // return nil = no error
}
