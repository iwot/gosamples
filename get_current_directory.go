package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// 実行ファイルの場所
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	// 実行時ディレクトリ
	current, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(current)
}
