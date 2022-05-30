package main

import (
	"flag"
	"fmt"

	"github.com/ASentientBanana/fd/util"
)

func define_flags() {
	flag.Func("r", "Replace a string in the target file with a given string.", func(s string) error {
		util.Replace(s, flag.Arg(0), flag.Arg(1))
		return nil
	})
	flag.Func("c", "Copy the target file to a given path.", func(s string) error {
		util.Copy(s, flag.Arg(0))
		return nil
	})
	flag.Func("m", "Move the target file to a given path.", func(s string) error {
		util.Move(s, flag.Arg(0))
		return nil
	})
	flag.Func("t", "Move the target file to a given path.", func(s string) error {
		util.Truncate_file(s)
		fmt.Println("THIS IS HAPPENING")
		return nil
	})
	flag.Func("ci", "Move the target file to a given path.", func(s string) error {
		util.Clone_into(s, flag.Arg(0))
		return nil
	})
	flag.Func("d", "Delete file or empty dir", func(s string) error {
		util.Delete_file(s)
		return nil
	})
	flag.Func("dr", "Run recursive delete on a path", func(s string) error {
		util.Recursive_delete(s)
		return nil
	})
	flag.Parse()
}

func main() {
	define_flags()
}
