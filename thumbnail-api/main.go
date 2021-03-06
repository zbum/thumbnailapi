package main

import (
	"github.com/go-martini/martini"
	"gopkg.in/gographics/imagick.v3/imagick"
	"thumbnailapi"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	classic := martini.Classic()
	classic.Post("/v1/thumbnail/convert", thumbnailapi.Resize)
	classic.Get("/v1/server-ip", thumbnailapi.ServerIp)
	classic.RunOnAddr(":10400")
}
