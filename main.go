package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

const outTemplate = `package main
{{range .Files}}
// {{.FunctionName}} returns binary file content.
func {{.FunctionName}}() []byte {
    return []byte{ {{.Content}} }
}
{{end}}`

type gofile struct {
	Files []gofileFile
}

type gofileFile struct {
	FunctionName string
	Content      string
}

func main() {
	// parse arguments
	path := flag.String("path", "", "path to gofiles.txt (excluding file name)")
	flag.Parse()
	if path == nil || *path == "" {
		panic("missing parameter 'path'")
	}

	// read gofiles.txt
	fileBinary, err := ioutil.ReadFile(*path + "gofiles.txt")
	if err != nil {
		panic(err)
	}

	// parse content of gofile.txt
	fileString := string(fileBinary)
	lines := strings.Split(fileString, "\n")

	// load each file
	outData := gofile{}
	for _, line := range lines {
		elm := strings.Split(line, ";")
		filePath := elm[0]
		functionName := "GoFile" + elm[1]

		// read file
		fileContent, err := ioutil.ReadFile(*path + filePath)
		if err != nil {
			panic(err)
		}

		// convert binary data for output
		fileContentString := ""
		for i, byt := range fileContent {
			if i != 0 {
				fileContentString += ","
			}
			fileContentString += fmt.Sprintf("%#x", byt)
		}

		// build output object
		outFile := gofileFile{FunctionName: functionName, Content: fileContentString}
		outData.Files = append(outData.Files, outFile)
	}

	// write output file
	tmpl := template.Must(template.New("out").Parse(outTemplate))
	tmpl.Execute(os.Stdout, outData)
}
