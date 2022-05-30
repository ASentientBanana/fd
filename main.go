package main

import (
	"errors"
	"flag"

	"github.com/ASentientBanana/fd/util"
)

func validate(replace, copy, move, truncate, clone_into *bool) {
	if flag.NArg() == 0 {
		panic(errors.New("Error no arguments or flags provided"))
	}
	if *replace && flag.NArg() == 3 {
		util.Replace(flag.Arg(0), flag.Arg(1), flag.Arg(2))
	}
	if *copy && flag.NArg() == 2 {
		util.Copy(flag.Arg(0), flag.Arg(1))
	}
	if *move && flag.NArg() == 2 {
		util.Move(flag.Arg(0), flag.Arg(1))
	}
	if *truncate && flag.NArg() == 1 {
		util.Truncate_file(flag.Arg(0))
	}
	if *clone_into && flag.NArg() == 2 {
		util.Clone_into(flag.Arg(0), flag.Arg(1))
	}

}

func main() {

	replace := flag.Bool("r", false, "Replace all string occurrences in a given file: path_to_string;new_string;target_string")
	copy := flag.Bool("c", false, "Copy file from one path 1 to path2")
	clone_into := flag.Bool("ci", false, "Copy file contents into a different file")
	move := flag.Bool("m", false, "Move file from one path 1 to path2")
	truncate := flag.Bool("t", false, "Empty out file")
	flag.Parse()

	if flag.NFlag() > 1 {
		panic(errors.New("Only one flag supported currently."))
	}

	validate(replace, copy, move, truncate, clone_into)

}
