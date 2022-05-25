package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
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

func truncate_file(filePath string) error {
	if err := os.Truncate(filePath, 0); err != nil {
		return err
	}
	return nil
}

func ErrorCleanup() {
	fmt.Println("Something went wrong")
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
	truncate_file(path)

	for i := 0; i < len(new_file_data); i++ {
		_, wErr := file.WriteString(new_file_data[i])
		if wErr != nil {
			panic(wErr)
		}
	}
	defer os.Remove(tmp_file_path)
}

func validate(replace, copy, move *bool) {
	if flag.NArg() == 0 {
		panic(errors.New("Error no arguments or flags provided"))
	}
	if *replace && flag.NArg() == 3 {
		Replace(flag.Arg(0), flag.Arg(1), flag.Arg(2))
	}
	if *copy && flag.NArg() == 2 {
		Copy(flag.Arg(0), flag.Arg(1))
	}
	if *move && flag.NArg() == 2 {
		Move(flag.Arg(0), flag.Arg(1))
	}
}

func main() {

	replace := flag.Bool("r", false, "Replace all string occurrences in a given file: path_to_string;new_string;target_string")
	copy := flag.Bool("c", false, "Copy file from one path 1 to path2")
	move := flag.Bool("m", false, "Move file from one path 1 to path2")
	flag.Parse()

	if flag.NFlag() > 1 {
		panic(errors.New("Only one flag supported currently."))
	}

	validate(replace, copy, move)

}
