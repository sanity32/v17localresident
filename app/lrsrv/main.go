package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sanity32/v17localresident"
)

var port int = v17localresident.DEFAULT_RPC_SERVER_PORT
var host string = "localhost"

func parseFlags() {
	flag.IntVar(&port, "port", port, "RPC server port")
	flag.StringVar(&host, "host", host, "RPC server host")
	flag.Parse()
}

func main() {
	parseFlags()
	addr := fmt.Sprintf("%v:%v", host, port)
	ch, err := v17localresident.NewServer(addr).Setup().Run()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print("Done:", <-ch)
}
