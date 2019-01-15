package effective

import (
	"os"
	"io"
)

// Deferred functions are executed in LIFO order
func Content(filename string) (string, error) {
	file, err := os.Open(filename)
	if nil != err {
		return "", err
	}
	defer file.Close()
	
	var result []byte
	buffer := make([]byte, 100)
	for {
		n, err := file.Read(buffer[0:])
		result = append(result, buffer[0:n]...)
		if nil != err {
			if err == io.EOF {
				break
			}
			return "", err
		}
	}
	return string(result), nil
}
