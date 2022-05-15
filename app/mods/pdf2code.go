package mod

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"code.sajari.com/docconv"
)

func Pdf2Code(input string, output string) (string, error) {
	testCode()
}

func testCode() {
	for _, v := range []string{"2022ST42AndroidSample01.pdf", "2022ST42AndroidSample02.pdf", "2022ST42AndroidSample03.pdf"} {
		res, err := docconv.ConvertPath(v)
		if err != nil {
			log.Println(err)
			continue
		}

		// コードを分割
		splitCodeFiles(res.Body)

		fmt.Println(v + ".txt")
		outputCode(res.Body, v+".txt")
	}
}

func splitCodeFiles(body string) [][]string {
	bodyLines := regexp.MustCompile("\r\n|\n").Split(body, -1)
	for i, v := range bodyLines {
		if regexp.MustCompile("").MatchString(v) {

		}
		else if regexp.MustCompile("").MatchString(v) {

		}
		else if regexp.MustCompile("").MatchString(v) {
			
		}
	}

	return [][]string{}
}

func outputCode(body string, path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err) //ファイルが開けなかったときエラー出力
	}

	defer file.Close()
	file.Write(([]byte)(body))
}
