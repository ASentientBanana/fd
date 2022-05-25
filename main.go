package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func copy(src, dest string) (int64, error) {
	origin, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer origin.Close()

	destinationFile, destErr := os.Create(dest)
	if destErr != nil {
		return 0, err
	}
	defer destinationFile.Close()

	nBytes, err := io.Copy(destinationFile, origin)
	return nBytes, err

}

func truncateFile(filePath string) error {
	if err := os.Truncate(filePath, 0); err != nil {
		return err
	}
	return nil
}

func ErrorCleanup() {
	fmt.Println("Something went wrong")
}

func main() {
	// 1: path 2:target 3: new word
	argsWithoutProg := os.Args[1:]
	split_string := strings.Split(argsWithoutProg[0], "/")
	tmp_file_path := strings.Join(split_string[:len(split_string)-1], "/") + ".tmp_file"
	_, cErr := copy(argsWithoutProg[0], tmp_file_path)
	if cErr != nil {
		panic(cErr)
	}
	tmp_file, tErr := os.OpenFile(tmp_file_path, os.O_RDWR, 0644)
	file, err := os.OpenFile(argsWithoutProg[0], os.O_RDWR, 0644)
	if err != nil || tErr != nil {
		panic(err)
	}

	defer file.Close()
	defer tmp_file.Close()

	scan := bufio.NewScanner(tmp_file)
	new_file_data := []string{}
	for scan.Scan() {
		line := strings.ReplaceAll(scan.Text(), argsWithoutProg[1], argsWithoutProg[2])
		new_file_data = append(new_file_data, line+"\n")
	}
	truncateFile(argsWithoutProg[0])

	for i := 0; i < len(new_file_data); i++ {
		_, wErr := file.WriteString(new_file_data[i])
		if wErr != nil {
			panic(wErr)
		}
	}
	defer os.Remove(tmp_file_path)
}
