package main

import (
	"fmt"
	"os"
	"path"

	"gitlab.com/oozie/example/fetcher"
)

func main() {

	fmt.Println(fetcher.BANNER)
	fmt.Println()

	cmdArgs := os.Args[1:]
	if len(cmdArgs) != 1 {
		panic(fmt.Sprintf("Usage: %s <elastic-url>", path.Base(os.Args[0])))
	}
	fetchTarget := cmdArgs[0]

	greeting, err := fetcher.GetGreetingFromURL(fetchTarget)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", greeting)

}
