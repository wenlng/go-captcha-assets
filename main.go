package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-bindata/go-bindata"
	"github.com/wenlng/go-captcha-assets/helper"
)

func makeFonts() []*bindata.Config {
	return []*bindata.Config{
		{
			Output:  "bindata/fonts/fzshengsksjw_cu/font.go",
			Package: "fzshengsksjw_cu",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/fonts/fzshengsksjw_cu",
				},
			},
		},
		{
			Output:  "bindata/fonts/yrdzst_bold/font.go",
			Package: "yrdzst_bold",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/fonts/yrdzst-bold",
				},
			},
		},
	}
}

func makeTilesImage() []*bindata.Config {
	return []*bindata.Config{
		{
			Output:  "bindata/tiles/tile_1/image.go",
			Package: "tile_1",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/tiles/tile-1",
				},
			},
		},
		{
			Output:  "bindata/tiles/tile_2/image.go",
			Package: "tile_2",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/tiles/tile-2",
				},
			},
		},
		{
			Output:  "bindata/tiles/tile_3/image.go",
			Package: "tile_3",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/tiles/tile-3",
				},
			},
		},
		{
			Output:  "bindata/tiles/tile_4/image.go",
			Package: "tile_4",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/tiles/tile-4",
				},
			},
		},
	}
}

func makeImage() []*bindata.Config {
	return []*bindata.Config{
		{
			Output:  "bindata/images/image_1/image.go",
			Package: "image_1",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/images/image-1",
				},
			},
		},
		{
			Output:  "bindata/images/image_2/image.go",
			Package: "image_2",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/images/image-2",
				},
			},
		},
		{
			Output:  "bindata/images/image_3/image.go",
			Package: "image_3",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/images/image-3",
				},
			},
		},
		{
			Output:  "bindata/images/image_4/image.go",
			Package: "image_4",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/images/image-4",
				},
			},
		},
		{
			Output:  "bindata/images/image_5/image.go",
			Package: "image_5",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/images/image-5",
				},
			},
		},
		// image-v2
		{
			Output:  "bindata/images/image_v2_1/image.go",
			Package: "image_v2_1",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/images/image-v2-1",
				},
			},
		},
		{
			Output:  "bindata/images/image_v2_2/image.go",
			Package: "image_v2_2",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/images/image-v2-2",
				},
			},
		},
		{
			Output:  "bindata/images/image_v2_3/image.go",
			Package: "image_v2_3",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/images/image-v2-3",
				},
			},
		},
		{
			Output:  "bindata/images/image_v2_4/image.go",
			Package: "image_v2_4",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/images/image-v2-4",
				},
			},
		},
		{
			Output:  "bindata/images/image_v2_5/image.go",
			Package: "image_v2_5",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/images/image-v2-5",
				},
			},
		},
		{
			Output:  "bindata/images/image_v2_6/image.go",
			Package: "image_v2_6",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/images/image-v2-6",
				},
			},
		},
		{
			Output:  "bindata/images/image_v2_7/image.go",
			Package: "image_v2_7",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/images/image-v2-7",
				},
			},
		},
		{
			Output:  "bindata/images/image_v2_8/image.go",
			Package: "image_v2_8",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/images/image-v2-8",
				},
			},
		},
	}
}

func makeThumbImage() []*bindata.Config {
	return []*bindata.Config{
		{
			Output:  "bindata/thumbs/thumb_1/image.go",
			Package: "thumb_1",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/thumbs/thumb-1",
				},
			},
		},
		{
			Output:  "bindata/thumbs/thumb_2/image.go",
			Package: "thumb_2",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/thumbs/thumb-2",
				},
			},
		},
		{
			Output:  "bindata/thumbs/thumb_3/image.go",
			Package: "thumb_3",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/thumbs/thumb-3",
				},
			},
		},
		{
			Output:  "bindata/thumbs/thumb_4/image.go",
			Package: "thumb_4",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/thumbs/thumb-4",
				},
			},
		},
		{
			Output:  "bindata/thumbs/thumb_5/image.go",
			Package: "thumb_5",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/thumbs/thumb-5",
				},
			},
		},
	}
}

func makeShapesImage() []*bindata.Config {
	return []*bindata.Config{
		{
			Output:  "bindata/shapes/shape_1/image.go",
			Package: "shape_1",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/shapes/shape-1",
				},
			},
		},
		{
			Output:  "bindata/shapes/shape_2/image.go",
			Package: "shape_2",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/shapes/shape-2",
				},
			},
		},
		{
			Output:  "bindata/shapes/shape_3/image.go",
			Package: "shape_3",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/shapes/shape-3",
				},
			},
		},
		{
			Output:  "bindata/shapes/shape_4/image.go",
			Package: "shape_4",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/shapes/shape-4",
				},
			},
		},
		{
			Output:  "bindata/shapes/shape_5/image.go",
			Package: "shape_5",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/shapes/shape-5",
				},
			},
		},
		{
			Output:  "bindata/shapes/shape_6/image.go",
			Package: "shape_6",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/shapes/shape-6",
				},
			},
		},
		{
			Output:  "bindata/shapes/shape_7/image.go",
			Package: "shape_7",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/shapes/shape-7",
				},
			},
		},
		{
			Output:  "bindata/shapes/shape_8/image.go",
			Package: "shape_8",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/shapes/shape-8",
				},
			},
		},
		{
			Output:  "bindata/shapes/shape_9/image.go",
			Package: "shape_9",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/shapes/shape-9",
				},
			},
		},
		{
			Output:  "bindata/shapes/shape_10/image.go",
			Package: "shape_10",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/shapes/shape-10",
				},
			},
		},
		{
			Output:  "bindata/shapes/shape_11/image.go",
			Package: "shape_11",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/shapes/shape-11",
				},
			},
		},
		{
			Output:  "bindata/shapes/shape_12/image.go",
			Package: "shape_12",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/shapes/shape-12",
				},
			},
		},
		{
			Output:  "bindata/shapes/shape_13/image.go",
			Package: "shape_13",
			Input: []bindata.InputConfig{
				{
					Path: "sourcedata/shapes/shape-13",
				},
			},
		},
	}
}

func configs() []*bindata.Config {
	confs := make([]*bindata.Config, 0)
	confs = append(confs, makeFonts()...)
	confs = append(confs, makeImage()...)
	confs = append(confs, makeThumbImage()...)
	confs = append(confs, makeTilesImage()...)
	confs = append(confs, makeShapesImage()...)
	return confs
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
