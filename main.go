package main

import (
	"fmt"
	"flag"
	"filehandling"
)

func main() {  
    fptr := flag.String("fpath", "test.txt", "file path to read from")
    flag.Parse()
    
	fmt.Println("Contents of file:")
	fmt.Println(filehandling.ParseFile(*fptr))
}