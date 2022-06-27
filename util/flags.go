package util

import (
	"flag"
)

func Define_flags() {
	flag.Func("r", "Replace a string in the target file with a given string.", func(s string) error {
		Replace(s, flag.Arg(0), flag.Arg(1))
		return nil
	})
	flag.Func("c", "Copy the target file to a given path.", func(s string) error {
		Copy(s, flag.Arg(0))
		return nil
	})
	flag.Func("m", "Move the target file to a given path.", func(s string) error {
		Move(s, flag.Arg(0))
		return nil
	})
	flag.Func("t", "Move the target file to a given path.", func(s string) error {
		Truncate_file(s)
		return nil
	})
	flag.Func("ci", "Move the target file to a given path.", func(s string) error {
		Clone_into(s, flag.Arg(0))
		return nil
	})
	flag.Func("d", "Delete file or empty dir", func(s string) error {
		Delete_file(s)
		return nil
	})
	flag.Func("dr", "Run recursive delete on a path", func(s string) error {
		Recursive_delete(s)
		return nil
	})
	flag.Parse()
}
