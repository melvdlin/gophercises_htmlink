package utils

import "os"

func ReadFile(path string) (content []byte, err error) {
	var file *os.File
	file, err = os.Open(path)
	if err != nil {
		return nil, err
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	content = make([]byte, fileInfo.Size())
	_, err = file.Read(content)
	if err != nil {
		return nil, err
	}
	return content, err
}
