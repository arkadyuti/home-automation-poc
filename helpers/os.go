package helpers

import (
	"io/ioutil"
	"os"
)

// WriteStringToFile Write sting to file, create file if not exist
func WriteStringToFile(filename, data string) error {
	// Open the file with write and create permissions, set file mode to 0644
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the string data to the file
	_, err = file.WriteString(data)
	return err
}

func ReadFile(filename string) ([]byte, error) {
	// Open the file with read-only permissions
	file, err := os.Open(filename)
	if err != nil {
		return []byte(""), err
	}
	defer file.Close()

	// Read the file contents into a byte slice
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return []byte(""), err
	}

	// Convert the byte slice to a string and return
	return bytes, nil
}
