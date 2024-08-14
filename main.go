package gowhich

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Provide executable name")
		os.Exit(1)
	}
	argsWithoutProg := os.Args[1:]

	executableName := argsWithoutProg[0]

	path := os.Getenv("PATH")

	if len(path) == 0 {
		fmt.Println("Can't read PATH env variable")
		os.Exit(1)
	}

	for _, value := range filepath.SplitList(path) {
		fullPath := filepath.Join(value, executableName)
		fileStats, err := os.Stat(fullPath)
		if err != nil {
			continue
		}

		mode := fileStats.Mode()
		if mode.IsRegular() {
			if mode&0111 != 0 {
				fmt.Println(fullPath)
				os.Exit(0)
			}
		}

	}

	fmt.Println("Not found")
	os.Exit(1)
}
