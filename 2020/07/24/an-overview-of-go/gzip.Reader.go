package gzip

func NewReader(r io.Reader) (*Reader, error) { ... }

type Reader struct {...}

func (z *Reader) Read(p []byte) (n int, err error) { ... }