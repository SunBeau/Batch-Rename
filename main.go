package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var exitCode = 0

func main() {
	delstr := flag.String("delstr", "", "String to delete.")
	dir := flag.String("dir", "", "Dir to use.")
	preview := flag.Bool("preview", true, "only preview.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s -dir=<DIR>\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	count := 0

	if len(*dir) > 0 && len(*delstr) > 0 {
		filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				// fmt.Println(path + "/")
				/* do nothing */
			} else {
				newpath := strings.Replace(path, *delstr, "", -1)

				if !strings.EqualFold(path, newpath) {
					fmt.Printf("%s -> %s\n", path, newpath)

					if !(*preview) {
						os.Rename(path, newpath)
					}

					count++
				}
			}

			return nil
		})

		fmt.Printf("Count = %d\n", count)
	} else {
		flag.Usage()
	}

	os.Exit(exitCode)
}
