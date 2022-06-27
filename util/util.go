package util

import (
	"bufio"
	"io"
	"os"
)

func ErrorCleanup(err error) {
	panic(err)
}

func file_text_to_slice(file io.Reader) []string {
	scan := bufio.NewScanner(file)
	new_file_data := []string{}
	for scan.Scan() {
		new_file_data = append(new_file_data, scan.Text()+"\n")
	}
	return new_file_data
}
func write_file_from_array(file_contents_array []string, file *os.File) error {
	for i := 0; i < len(file_contents_array); i++ {
		_, wErr := file.WriteString(file_contents_array[i])
		if wErr != nil {
			return wErr
		}
	}
	return nil
}
