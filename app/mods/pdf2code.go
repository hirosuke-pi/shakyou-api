package mod

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"code.sajari.com/docconv"
)

func Pdf2Code(input string, output string) (string, error) {
	testCode()

	return "", nil
}

type CodeFiles struct {
	codes []string
	path  string
	nest  bool
}

func testCode() {
	for _, v := range []string{"2022ST42AndroidSample01.pdf"} { //, "2022ST42AndroidSample02.pdf", "2022ST42AndroidSample03.pdf"} {
		res, err := docconv.ConvertPath(v)
		if err != nil {
			log.Println(err)
			continue
		}

		bodyLines := regexp.MustCompile("\r\n|\n").Split(res.Body, -1)
		files := []CodeFiles{}

		// コードを分割
		extractCode(&files, &bodyLines)

		fmt.Println(v + ".txt")
		outputCode(strings.Join(bodyLines, "\n"), v+"_.txt")
		outputCode(res.Body, v+".txt")
	}
}

func extractCode(files *[]CodeFiles, bodyLines *[]string) {
	codes := []string{}
	expCodes := []string{}
	tmpLines := *bodyLines
	codeFlag := false
	newlineFlag := false

	for i, v := range tmpLines {
		if regexp.MustCompile(".*?¬").MatchString(v) {
			line := strings.ReplaceAll(v, "¬", "")

			if newlineFlag {
				codes[len(codes)-1] += line
				newlineFlag = false
			} else {
				codes = append(codes, line)
			}
			codeFlag = true
		} else if codeFlag && strings.TrimSpace(v) != "" {
			codes[len(codes)-1] += v
			newlineFlag = true
		} else if codeFlag {
			*files = append(*files, CodeFiles{
				codes: codes,
			})
			outputCode(strings.Join(codes, "\n"), strconv.Itoa(i)+".txt")

			codeFlag = false
			codes = []string{}
		} else {
			expCodes = append(expCodes, v)
		}
	}

	if codeFlag {
		*files = append(*files, CodeFiles{
			codes: codes,
		})
		outputCode(strings.Join(codes, "\n"), "999999999.txt")
	}
	*bodyLines = expCodes
}

func outputCode(body string, path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err) //ファイルが開けなかったときエラー出力
	}

	defer file.Close()
	file.Write(([]byte)(body))
}
