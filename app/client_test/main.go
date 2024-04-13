package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

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
	cl, err := v17localresident.ConnectClient(addr)
	if err != nil {
		log.Fatal(err.Error())
	}

	cl.MouseMove(20, 20, false)
	cl.MouseClick("right", false)
	scr, err := cl.ScreenshotRect(image.Rect(10, 10, 50, 50))
	if err != nil {
		log.Fatal(err.Error())
	}

	f, _ := os.Create("t.jpg")
	defer f.Close()
	jpeg.Encode(f, scr, nil)

	log.Println(cl.MouseLocation())
}
