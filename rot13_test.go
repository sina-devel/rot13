package rot13

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

func TestRot13Reader(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "decode",
			args: args{strings.NewReader("Lbh penpxrq gur pbqr!")},
			want: "You cracked the code!",
		},
		{
			name: "encode",
			args: args{strings.NewReader("You cracked the code!")},
			want: "Lbh penpxrq gur pbqr!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf, _ := ioutil.ReadAll(NewRot13Reader(tt.args.reader))
			if string(buf) != tt.want {
				t.Errorf("Rot13Reader = %q, want %q", string(buf), tt.want)
			}
		})
	}
}
