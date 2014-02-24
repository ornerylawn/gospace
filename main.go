package main

import (
	"flag"
	"fmt"
	"os"
)

var clearFlag = flag.Bool("clear", false, "make it so that no space is active")

func main() {
	flag.Usage = func() {
		fmt.Printf(`gospace is a tool for changing between Go workspaces. Just set your
$GOPATH to ~/.gospace, which will be a symlink to your current
workspace. gospace simply points the symlink to the path that you
specifiy as an argument. Leave out the path to see the current one.

Usage: %s [<path>]`, os.Args[0])
		fmt.Println()
		flag.PrintDefaults()
		fmt.Println()
	}
	flag.Parse()
	if *clearFlag {
		clear()
		get()
		return
	}
	switch len(os.Args) {
	case 1:
		get()
	case 2:
		set(os.Args[1])
	default:
		fmt.Printf("Usage: %s [<path>]", os.Args[0])
	}
}

func get() {
	path, err := os.Readlink(spacefile())
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println(err)
		}
		return
	}
	fmt.Println(path)
}

func set(path string) {
	clear()
	err := os.Symlink(path, spacefile())
	if err != nil {
		fmt.Println(err)
		return
	}
	get()
}

func clear() {
	err := os.Remove(spacefile())
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println(err)
		}
	}
}

func spacefile() string {
	return os.Getenv("HOME") + "/.gospace"
}
