package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	clear    = flag.Bool("clear", false, "make it so that no space is active")
	global   = flag.String("global", "", "make a binary accessible from every workspace")
	unglobal = flag.String("unglobal", "", "undo what -global does")
)

func main() {
	handleFlags()
	switch len(os.Args) {
	case 1:
		dst, err := readLink(spacefile())
		if err != nil {
			fmt.Println(err)
			return
		}
		if dst != "" {
			fmt.Println(dst)
		} else {
			fmt.Println("no active workspace")
		}
	case 2:
		dst := os.Args[1]
		link := spacefile()
		err := setLink(link, dst)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s -> %s\n", link, dst)
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Printf(`gospace is a Go workspace switcher.

More info: http://github.com/rynlbrwn/gospace

Usage: %s [-clear|-global|-unglobal] [<path>]`, os.Args[0])
	fmt.Println()
	flag.PrintDefaults()
	fmt.Println()
}

func handleFlags() {
	flag.Usage = printUsage
	flag.Parse()
	// Remove flags and binary name from args.
	args := []string{}
	for _, arg := range os.Args[1:] {
		if !strings.HasPrefix(arg, "-") {
			args = append(args, arg)
		}
	}
	if *clear && *global == "" && *unglobal == "" {
		doClear(args)
	} else if *global != "" && !(*clear) && *unglobal == "" {
		doGlobal(args)
	} else if *unglobal != "" && !(*clear) && *global == "" {
		doUnglobal(args)
	} else if *clear || *global != "" || *unglobal != "" {
		fmt.Println("one flag at a time please")
		os.Exit(1)
	} else {
		return
	}
	os.Exit(0)
}

func doClear(args []string) {
	if len(args) != 0 {
		fmt.Println("-clear doesn't take any args")
		return
	}
	dst := "/dev/null"
	link := spacefile()
	err := setLink(link, dst)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s -> %s\n", link, dst)
}

func doGlobal(args []string) {
	if len(args) != 1 {
		fmt.Println("-global takes exactly 1 argument")
		return
	}
	dst := args[0]
	base := filepath.Base(dst)
	link := "/usr/local/bin/" + base
	err := setLink(link, dst)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s -> %s\n", link, dst)
}

func doUnglobal(args []string) {
	if len(args) != 1 {
		fmt.Println("-unglobal takes exactly 1 argument")
		return
	}
	dst := args[0]
	base := filepath.Base(dst)
	link := "/usr/local/bin/" + base
	err := rmLink(link)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("removed %s\n", link)
}

func readLink(path string) (string, error) {
	dst, err := os.Readlink(path)
	if err != nil && os.IsNotExist(err) {
		return "", nil
	}
	return dst, err
}

func setLink(name, dst string) error {
	err := rmLink(name)
	if err != nil {
		return err
	}
	err = os.Symlink(dst, name)
	return err
}

func rmLink(path string) error {
	err := os.Remove(path)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

func spacefile() string {
	return os.Getenv("HOME") + "/.gospace"
}
