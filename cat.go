package gosh

import (
	"fmt"
	"io"
	"os"
)

func Cat(filenames ...string) {
	if len(filenames) == 0 {
		///TODO read from Stdin
	} else {
		fmt.Fprintln(debugStream, "branch catfile")
		//TODO move to an other function, in order to go cat
		files := make([]*os.File, len(filenames))
		for i, filename := range filenames {
			fmt.Fprintln(debugStream, "opening"+filename)
			file, err := os.Open(filename)
			if err != nil {
				panic("Error while opening " + filename)
			}
			files[i] = file
		}
		pipe := NewPipe()
		if len(curPipes) == 0 {
			curPipes = make([]Pipe, 1)
		}
		curPipes[0] = pipe
		fmt.Fprintln(debugStream, "calling fileCat")
		wg.Add(1)
		go fileCat(files, pipe)
	}
}

func fileCat(files []*os.File, pipe Pipe) {
	defer wg.Done()
	defer close(pipe)
	fmt.Fprintln(debugStream, "fileCat called")
	for _, file := range files {
		fmt.Fprintln(debugStream, "catting files")
		//TODO write in curPipe[0] i files uno dopo l'altro, blocco per blocco
		var err error
		for err == nil { //TODO ErrUnexpectedEOF EOF
			block := NewBlock()
			_, err = io.ReadFull(file, block)
			pipe <- block
		}
	}
}

/*
func readLine(path string) {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}*/
