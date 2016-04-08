package gosh

import (
	"os"
	"sync"
)

type Block []byte
type Pipe chan Block

//NOTE string or array of runes? (otherwise buffersize has no meaning whatsoever
var BufferSize int = 4096
var BlockSize int = 128
var curPipes []Pipe
var debugStream *os.File
var wg sync.WaitGroup

func init() {
	debugStream = os.Stderr
	//TODO start the reader on stdin NOTE: rimettere gli a capi che readline toglie
	//TODO initialize Pipe
}

func NewPipe() Pipe {
	return make(Pipe, BufferSize)
}
func NewBlock() Block {
	return make(Block, BlockSize)
}

func closePipes() {
	//	for i, pipe := range curPipes {
	//		close(pipe)
	//	}
}
func Start() {}
func End() {
	defer wg.Wait()
}

//TODO create a reader che wrappa il chan con i bytes e permetta di leggerlo come stream o line-by-line
