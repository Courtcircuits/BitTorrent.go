package main

import (
	"fmt"

	"github.com/Courtcircuits/BitTorrent.go/bencode"
)

func main() {
	integer:= "l4:spam4:eggsi-32el4:spami32eee"
	fmt.Println(bencode.Parse([]byte(integer)))
}
