package fileparser 
import (
    "fmt"
    "io/ioutil"
)

func ParseFile(path string) string{
	data, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Println("File reading error", err)
        return ""
	}
	return string(data)
}