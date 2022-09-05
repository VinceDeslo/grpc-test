package utils

import (
	"io/ioutil"
	"os"
)

func ReadJsonFile(fp string) ([]byte, error) {
	jsonFile, err := os.Open(fp)

	if err != nil {
		return []byte{}, err
	}
	defer jsonFile.Close()

	bytes, _ := ioutil.ReadAll(jsonFile)
	return bytes, nil
}
