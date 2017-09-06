package main

import (
	"bytes"
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
	verboseLog := flag.Bool("verbose", false, "verbose log")
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
		log(verboseLog, "start processing file "+filePath+" with name "+functionName)

		// read file
		fileContent, err := ioutil.ReadFile(*path + filePath)
		if err != nil {
			panic(err)
		}
		log(verboseLog, "- file successfully loaded")

		// convert binary data for output
		var buffer bytes.Buffer
		for i, byt := range fileContent {
			if i != 0 {
				buffer.WriteString(",")
			}
			buffer.WriteString(fmt.Sprintf("%#x", byt))
		}
		fileContentString := buffer.String()
		log(verboseLog, "- file content converted into byte stream")

		// build output object
		outFile := gofileFile{FunctionName: functionName, Content: fileContentString}
		outData.Files = append(outData.Files, outFile)
		log(verboseLog, "- file successfully processed")
	}

	// write output file
	tmpl := template.Must(template.New("out").Parse(outTemplate))
	tmpl.Execute(os.Stdout, outData)
	log(verboseLog, "gofile successfully saved")
}

func log(verboseLog *bool, text string) {
	if verboseLog == nil || !*verboseLog {
		return
	}

	fmt.Println("// " + text)
}
