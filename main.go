package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var (
		num = flag.Uint("n", 2, "number of line breaks to check")
	)

	flag.Parse()

	// フラグでないコマンドライン引数へのアクセス
	dirArg := flag.Arg(0)

	if dirArg == "" {
		log.Fatal("please make a folder or file.")
	}

	dir, err := filepath.Abs(dirArg)
	if err != nil {
		log.Fatal(err)
	}

	dirInfo, _ := os.Stat(dir)

	var files []os.FileInfo
	var isSingleFile bool
	if dirInfo.IsDir() {
		files, err = ioutil.ReadDir(dir)
		if err != nil {
			panic(err)
		}
	} else {
		isSingleFile = true
		files = []os.FileInfo{dirInfo}
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		var filePath string
		if isSingleFile {
			filePath = dir
		} else {
			filePath = filepath.Join(dir, file.Name())
		}

		fst := token.NewFileSet()
		f, err := parser.ParseFile(fst, filePath, nil, 0)
		if err != nil {
			log.Fatal("Error:", err)
		}

		var (
			beforeImportLine = 3
			afterImportLine  int
			lineCount        uint
			lintText         string
		)
		ast.Inspect(f, func(n ast.Node) bool {

			switch n := n.(type) {
			case *ast.ImportSpec:
				afterImportLine = fst.Position(n.Pos()).Line
				if afterImportLine-beforeImportLine >= 2 {
					fmt.Println("")
					lineCount++
					if lineCount >= *num {
						lintText += "-\n"
					} else {
						lintText += "\n"
					}
				}
				beforeImportLine = afterImportLine
				lintText += fmt.Sprintf("  %s\n", n.Path.Value)
			}

			return true
		})

		if lineCount >= *num {
			fmt.Println(file.Name())
			fmt.Printf("フォーマットが不正のようです\n")
			fmt.Printf("imports (\n" + lintText + " ) \n")
			os.Exit(1)
		}
	}

}
