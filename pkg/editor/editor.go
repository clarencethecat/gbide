package editor

import (
	"os"
	"template"
	"io"
	"bytes"
	"fmt"
	"gonicetrace.googlecode.com/hg/nicetrace"
)

const EditorTemplatePath = "templates/editor.template"

var T *template.Template

func init() {
	T = template.New(nil)
	T.SetDelims("{{", "}}")
	T.ParseFile(EditorTemplatePath)
}

type Source struct {
	Path, Data string
}



func OpenFile(file string) (code string) {
	defer nicetrace.Print()

	data, err := ReadFile(file)
	
	if err == nil {
		buf := bytes.NewBuffer([]byte{})
		T.Execute(buf, Source{file, data})
		code = buf.String()
	} else {
		code = err.String()
	}
	return
}

func ReadAll(r io.Reader) (res string) {
	longbuf := bytes.NewBuffer(make([]byte, 0, 1024))
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if n == 0 || err != nil {
			break
		}
		longbuf.Write(buf[0:n])
	}
	res = longbuf.String()
	return
}

func ReadFile(file string) (data string, err os.Error) {
	var fin *os.File
	fin, err = os.Open(file, os.O_RDONLY, 0)
	if err != nil {
		return
	}
	data = ReadAll(fin)
	return
}
