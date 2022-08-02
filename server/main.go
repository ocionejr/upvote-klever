package main

import (
	"log"

	"github.com/ocionejr/upvote-klever/util"
)

func main(){
	_, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
}