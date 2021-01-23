package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
)

type File struct {
	HashValue [32]byte
	Name      string
	Known     bool
}

func main() {
	osfiles, _ := ioutil.ReadDir("./files/")
	Files := make([]File, 0)
	for _, f := range osfiles {
		data, _ := ioutil.ReadFile("./files/" + f.Name())
		var File File = File{
			Name:      f.Name(),
			HashValue: sha256.Sum256(data),
			Known:     false,
		}
		Files = append(Files, File)
	}
	for _, f := range Files {
		if f.Known {
			continue
		}
		RepeatFiles := make([]File, 0)
		RepeatFiles = append(RepeatFiles, f)
		f.Known = true
		for _, f2 := range Files {
			if f.HashValue == f2.HashValue && f2.Name != f.Name {
				f2.Known = true
				RepeatFiles = append(RepeatFiles, f2)
			}
		}
		if len(RepeatFiles) != 1 {
			for _, f3 := range RepeatFiles {
				fmt.Println(f3.Name)
			}
			fmt.Println()
		}
	}
}
