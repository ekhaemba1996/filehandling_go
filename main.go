package main

import (
	"fmt"
	"flag"
	"github.com/ekhaemba1996/filehandling_go/fileparser"
)

func main() {  
    fptr := flag.String("fpath", "test.txt", "file path to read from")
    flag.Parse()
    
	fmt.Println("Contents of file:")
	fmt.Println(fileparser.ParseFile(*fptr))
}