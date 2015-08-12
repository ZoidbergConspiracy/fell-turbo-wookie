package main

import (
	"fmt"
	"github.com/speps/go-hashids"
	"os"
)

func main() {

	if ( len(os.Args[0]) > 8 && os.Args[0][len(os.Args[0])-8:] == "unhashid" ) {
		for _, n := range os.Args[1:] {
			fmt.Println(hashToId(n, "", 0))
		}
	} else {
		for _, n := range os.Args[1:] {
			fmt.Println(idToHash(n, "", 0))

		}
	}

}

func idToHash(s string, salt string, min int) (id string) {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = min
	h := hashids.NewWithData(hd)

	b := []byte(s)
	d := make([]int, len(b))
	for i := range b {
		d[i] = int(b[i])
	}

	id, _ = h.Encode(d)
	return

}

func hashToId(id string, salt string, min int) (s string) {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = min
	h := hashids.NewWithData(hd)

	d := h.Decode(id)
	b := make([]byte, len(d))
	for i := range b {
		b[i] = byte(d[i])
	}

	s = string(b)
	return
}
