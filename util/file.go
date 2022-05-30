package util

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func Copy(src, dest string) (int64, error) {
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

func Move(src, dest string) (int64, error) {
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
	os.Remove(src)
	return nBytes, err
}

func Truncate_file(filePath string) error {
	if err := os.Truncate(filePath, 0); err != nil {
		return err
	}
	return nil
}

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

func Replace(path, target, new_word string) {
	split_string := strings.Split(path, "/")
	tmp_file_path := strings.Join(split_string[:len(split_string)-1], "/") + ".tmp_file"
	_, cErr := Copy(path, tmp_file_path)
	if cErr != nil {
		panic(cErr)
	}
	tmp_file, tErr := os.OpenFile(tmp_file_path, os.O_RDWR, 0644)
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil || tErr != nil {
		panic(err)
	}

	defer file.Close()
	defer tmp_file.Close()

	scan := bufio.NewScanner(tmp_file)
	new_file_data := []string{}
	for scan.Scan() {
		line := strings.ReplaceAll(scan.Text(), target, new_word)
		new_file_data = append(new_file_data, line+"\n")
	}
	Truncate_file(path)

	write_err := write_file_from_array(new_file_data, file)
	if write_err != nil {
		ErrorCleanup(write_err)
	}
	defer os.Remove(tmp_file_path)
}

func Clone_into(source_path, target_path string) {
	Truncate_file(target_path)
	source_file, source_err := os.OpenFile(source_path, os.O_RDWR, 0644)
	target_file, target_err := os.OpenFile(target_path, os.O_RDWR, 0644)
	if source_err != nil || target_err != nil {
		ErrorCleanup(source_err)
	}
	file_contents := file_text_to_slice(source_file)
	write_err := write_file_from_array(file_contents, target_file)
	if write_err != nil {
		ErrorCleanup(write_err)
	}
}

func Delete_file(path string) {
	os.Remove(path)
}

func Recursive_delete(path string) {
	os.RemoveAll(path)
	// dir_contents, readDirErr := os.ReadDir(path)
	// if readDirErr != nil {
	// 	ErrorCleanup(readDirErr)
	// }
	// for i := 0; i < len(dir_contents); i++ {
	// 	if dir_contents[i].IsDir() {
	// 		Recursive_delete(path + "/" + dir_contents[i].Name())
	// 	} else {
	// 		fmt.Println(dir_contents[i].Name())
	// 		os.Remove(path)
	// 	}
	// }
}
