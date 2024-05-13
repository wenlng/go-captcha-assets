package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-bindata/go-bindata"
	"github.com/wenlng/go-captcha-assets/helper"
)

func configs() []*bindata.Config {
	return []*bindata.Config{
		{
			Output:  "bindata/fonts/fzshengsksjw/fzshengsksjw_cu.go",
			Package: "fzshengsksjw",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/fonts/fzshengsksjw_cu.ttf",
				},
			},
		},
		{
			Output:  "bindata/fonts/yrdzst/yrdzst-bold.go",
			Package: "yrdzst",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/fonts/yrdzst-bold.ttf",
				},
			},
		},
		{
			Output:  "bindata/images/images.go",
			Package: "images",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/images",
				},
			},
		},
		{
			Output:  "bindata/thumbs/thumbs.go",
			Package: "thumbs",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/thumbs",
				},
			},
		},
		{
			Output:  "bindata/shapes/shapes.go",
			Package: "shapes",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/shapes",
				},
			},
		},
		{
			Output:  "bindata/tiles/tiles.go",
			Package: "tiles",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/tiles",
				},
			},
		},
	}
}

func main() {
	for _, config := range configs() {
		err := bindata.Translate(config)
		if err != nil {
			fmt.Fprintf(os.Stderr, "bindata: %v\n", err)
			os.Exit(1)
		}
	}

	err := helper.CopyFileTo("sourcedata/chars/char.go", "bindata/chars/char.go")
	if err != nil {
		fmt.Fprintf(os.Stderr, "copy resource to bindata: %v \n", err)
		os.Exit(1)
	}

	log.Println("gen bindata success!")
}
