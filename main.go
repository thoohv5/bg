package main

import (
	"log"

	"github.com/thoohv5/bg/boot"
)

func main() {
	// http
	if err := boot.Http(); nil != err {
		log.Fatalf("boot Http err:%v", err)
	}
}
