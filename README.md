# GoFiles
[![Build Status](https://travis-ci.org/martinr92/gofiles.svg?branch=master)](https://travis-ci.org/martinr92/gofiles)

This application creates a .go file, that contains binary data (like templates, images, ...).
Within your application, you can access to this data. This step is used to compile this data into the output binary file.

## How to use GoFiles?
### Step 1: Install GoFiles application
```
go get github.com/martinr92/gofiles
```

### Step 2: Create a gofiles.txt
To use GoFiles, simply create a text file called ```gofiles.txt``` within your project.
The content of this file must have the following format:
```
<filename>;<function name>
template/index.html;TemplateIndex
template/home.html;TemplateHome
```

### Step 3: Execute GoFiles
```
gofiles -path=$GOPATH/src/github.com/myapplication/ > $GOPATH/src/github.com/myapplication/gofiles.go
```

### Step 4: Use generated file
```
// GoFile<function name>
myFileContent := GoFileTemplateIndex()
```
