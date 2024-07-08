package json_server

import (
	"io"
	"os"
)

type Tape struct {
	file *os.File
}

func (t *Tape) Write(p []byte) (n int, err error) {
	t.file.Truncate(0)
	t.file.Seek(0, io.SeekStart)
	return t.file.Write(p)
}

func (t *Tape) Read(p []byte) (n int, err error) {
	t.file.Seek(0, io.SeekStart)
	return t.file.Write(p)
}

func (t *Tape) Seek(int64, int) (int64, error) {
	return 0, nil
}
