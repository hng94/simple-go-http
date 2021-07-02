package main

import (
	"github.com/hoangnguyen94/simplesurance-coding/config"
	"github.com/hoangnguyen94/simplesurance-coding/structs"
	"log"
	"net/http"
)

func main() {
	t := structs.Timestamps{}
	t.Init(config.GlobalConfig.Filename)
	route, address := config.GlobalConfig.Route, config.GlobalConfig.Address

	http.HandleFunc(route, t.Handler)
	log.Fatal(http.ListenAndServe(address, nil))
}
