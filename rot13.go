package rot13

import (
	"io"
)

type rot13Reader struct {
	r io.Reader
}

// NewRot13Reader returns an io.Reader that read rot13 characters from r.
func NewRot13Reader(reader io.Reader) io.Reader {
	return rot13Reader{
		r: reader,
	}
}

func (r rot13Reader) Read(p []byte) (int, error) {
	if l, err := r.r.Read(p); err != nil && err != io.EOF {
		return l, err
	}
	for i, v := range p {
		switch {
		case byte('A') <= v && v <= byte('M'):
			p[i] = v + 13
		case byte('N') <= v && v <= byte('Z'):
			p[i] = v - 13
		case byte('a') <= v && v <= byte('m'):
			p[i] = v + 13
		case byte('n') <= v && v <= byte('z'):
			p[i] = v - 13
		}
	}
	return len(p), io.EOF
}
