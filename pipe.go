package gosh

import (
	"bufio"
	"io"
)

type BufferedWritePipe interface {
	io.WriteCloser
	Flush() error
}

type BufferedReadPipe interface {
	io.ReadCloser
}

type readPipeWrapper struct {
	pr *io.PipeReader
	br *bufio.Reader
}

func (r *readPipeWrapper) Read(p []byte) (n int, err error) {
	return r.br.Read(p)
}

func (r *readPipeWrapper) Close() error {
	return r.pr.Close()
}

type writePipeWrapper struct {
	pw *io.PipeWriter
	bw *bufio.Writer
}

func (w *writePipeWrapper) Write(p []byte) (n int, err error) {
	return w.bw.Write(p)
}

func (w *writePipeWrapper) Close() error {
	ferr := w.bw.Flush()
	cerr := w.pw.Close()
	if cerr != nil {
		return cerr
	}
	return ferr
}

func (w *writePipeWrapper) Flush() error {
	return w.Flush()
}

func NewBufferedPipe() (BufferedReadPipe, BufferedWritePipe) {
	return BufferPipe(io.Pipe())
}

func BufferPipe(pr *io.PipeReader, pw *io.PipeWriter) (BufferedReadPipe, BufferedWritePipe) {
	return &readPipeWrapper{
			pr: pr,
			br: bufio.NewReader(pr),
		},
		&writePipeWrapper{
			pw: pw,
			bw: bufio.NewWriter(pw),
		}
}
