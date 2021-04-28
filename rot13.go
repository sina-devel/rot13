package rot13

import (
	"io"
)

type Rot13Reader struct {
	r io.Reader
}

// NewRot13Reader returns an Rot13Reader that read rot13 characters from r.
func NewRot13Reader(reader io.Reader) *Rot13Reader {
	return &Rot13Reader{
		r: reader,
	}
}

func (rot *Rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)
	for i, v := range p {
		switch {
		case (byte('A') <= v && v <= byte('M')) || (byte('a') <= v && v <= byte('m')):
			p[i] = v + 13
		case (byte('N') <= v && v <= byte('Z')) || (byte('n') <= v && v <= byte('z')):
			p[i] = v - 13
		}

	}
	return
}
