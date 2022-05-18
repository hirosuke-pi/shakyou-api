package mod

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"code.sajari.com/docconv"
)

type FileConfig struct {
	codes []string
	path  string
	name  string
	nest  bool
}

func Pdf2Code(input string, output string) ([]FileConfig, error) {
	input = "2022ST42AndroidSample01.pdf"

	// pdfから文字を抽出
	res, err := docconv.ConvertPath(input)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// 初期設定
	outputFile(res.Body, input+".txt")
	files := extractPdf(res.Body)

	fmt.Println(input + ".txt")

	return files, nil
}

func extractPdf(body string) []FileConfig {
	bodyLines := regexp.MustCompile("\r\n|\n").Split(body, -1)

	files := extractCode(&bodyLines) // コード部分を抽出
	extractName(&files, &bodyLines)  // ファイル名を抽出
	extractNest(&files, &bodyLines)  // ネストされたファイルかどうか
	joinNestedFile(&files)           // ネストされていた場合、結合する

	for _, v := range files {
		outputFileLines(v.codes, `./archives/`+v.name)
	}
	return files
}

func extractCode(bodyLines *[]string) []FileConfig {
	files := []FileConfig{}
	codes := []string{}
	expCodes := []string{}
	codeFlag := false
	newlineFlag := false

	for _, v := range *bodyLines {
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
			files = append(files, FileConfig{codes: codes})
			codeFlag = false
			codes = []string{}
		} else {
			expCodes = append(expCodes, v)
		}
	}

	if codeFlag {
		files = append(files, FileConfig{codes: codes})
	}
	*bodyLines = expCodes

	return files
}

func extractName(files *[]FileConfig, bodyLines *[]string) {
	index := 0
	expNames := []string{}

	for i, v := range *bodyLines {
		if regexp.MustCompile("^Printed:.*?").MatchString(v) {
			(*files)[index].name = (*bodyLines)[i-1]
			expNames = expNames[:len(expNames)-1]
			index++
		} else {
			expNames = append(expNames, v)
		}
	}
	*bodyLines = expNames
}

func extractNest(files *[]FileConfig, bodyLines *[]string) {
	index := 0
	expNames := []string{}

	for _, v := range *bodyLines {
		if regexp.MustCompile(`^Page 1/\d+?$`).MatchString(v) {
			(*files)[index].nest = false
			index++
		} else if regexp.MustCompile(`^Page \d+?/\d+?$`).MatchString(v) {
			(*files)[index].nest = true
			index++
		} else {
			expNames = append(expNames, v)
		}
	}
	*bodyLines = expNames
}

func joinNestedFile(files *[]FileConfig) {
	count := 0
	joinedFiles := []FileConfig{}

	for _, v := range *files {
		if v.nest {
			joinedFiles[count-1].codes = append(joinedFiles[count-1].codes, v.codes...)
		} else {
			joinedFiles = append(joinedFiles, v)
			count++
		}
	}
	*files = joinedFiles
}

func outputFile(body string, path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err) //ファイルが開けなかったときエラー出力
	}

	defer file.Close()
	file.Write(([]byte)(body))
}

func outputFileLines(bodyLines []string, path string) {
	outputFile(strings.Join(bodyLines, "\n"), path)
}
