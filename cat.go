package gosh

import (
	"bufio"
	"os"
)

func Cat(filenames ...string) (pipe Pipe) {
	addWait()
	defer removeWait()
	pipe = NewPipe()
	if len(filenames) == 0 {
		///TODO read from Stdin
	} else {
		files := make([]*os.File, len(filenames))
		for i, filename := range filenames {
			file, err := os.Open(filename)
			if err != nil {
				panic("Error while opening " + filename + ":" + err.Error())
			}
			files[i] = file
		}
		addWait()
		go fileCat(files, pipe)
	}
	return
}

func fileCat(files []*os.File, pipe Pipe) {
	defer removeWait()
	defer close(pipe)
	for _, file := range files {
		defer func() { _ = file.Close() }()
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			pipe <- scanner.Text()
		}
	}
}
