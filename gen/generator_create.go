package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const preamble = `
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

`

func (g *Generator) Create(path string, sym string) (io.WriteCloser, error) {
	if len(sym) > 0 {
		symWords := splitName(sym)
		numWords := 1
		if strings.HasPrefix(path+sym, "RequestWorkbookFunctions") {
			numWords = 3
		}
		if numWords > len(symWords) {
			numWords = len(symWords)
		}
		path += strings.Join(symWords[:numWords], "")
	}
	if !strings.HasSuffix(path, ".go") {
		path += ".go"
	}
	path = filepath.Join(g.Out, path)
	if g.Created[path] {
		return os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0644)
	}
	log.Printf("Creating %s", path)
	g.Created[path] = true
	out, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	out.WriteString(preamble)
	return out, nil
}