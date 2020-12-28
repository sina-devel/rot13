package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/sina-devel/rot13"
)

func main() {
	r := rot13.NewRot13Reader(strings.NewReader("Lbh penpxrq gur pbqr!"))
	res, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", res)
}
