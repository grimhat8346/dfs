package main

import (
	"github/grimhat8346/dfs/p2p"
	"log"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")
	log.Fatal(tr.ListenAndAccept())
	select {}

}
