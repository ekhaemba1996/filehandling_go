package fileparser 
import (
    "fmt"
    "io/ioutil"
    "regexp"
)

func ParseFile(path string) string{
	data, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Println("File reading error", err)
        return ""
	}
	return string(data)
}

func ParseFasta(fasta string) map[string]string{
	fastaRegexString := `>(?P<key>Rosalind_\d+)\r?\n(?P<entire_sequence>(?P<seq_line>[GATC]+\r?\n?)+)`
	reReplaceNewline := regexp.MustCompile(`(\r?\n)`)
	fastaRegex := regexp.MustCompile(fastaRegexString)
	resultMap := make(map[string]string)

	for _, submatches := range fastaRegex.FindAllStringSubmatch(fasta, -1){
		key := submatches[1]
		value := submatches[2]
		value = reReplaceNewline.ReplaceAllString(value, "")
		resultMap[key] = value
	}
	return resultMap
}