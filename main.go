package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileIn := flag.String("in", "", "file with encoded strings")
	sep := flag.String("sep", "\n", "separator character of encoded strings in")
	base := flag.Int("base", 64, "encoding base (32 or 64)")

	flag.Parse()

	fmt.Printf("Read file %v\n", *fileIn)
	file, err := os.Open(*fileIn)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	defer file.Close()
	datas, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	fmt.Println("Searching for hidden datas")
	res, err := unhide(string(datas), *sep, *base)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	fmt.Printf("\n%v\n", res)
}
