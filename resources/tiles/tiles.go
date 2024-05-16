package tiles

import (
	"fmt"
	"image"

	assets "github.com/wenlng/go-captcha-assets/bindata/tiles"
	"github.com/wenlng/go-captcha-assets/helper"
)

type GraphImage struct {
	OverlayImage image.Image
	ShadowImage  image.Image
	MaskImage    image.Image
}

func GetTiles() ([]*GraphImage, error) {
	var graphs []*GraphImage
	var err error
	for i := 1; i <= 4; i++ {
		var tileImg image.Image
		var tileShadowImg image.Image
		var tileMaskImg image.Image
		var asset []byte

		asset, err = assets.Asset(fmt.Sprintf("sourcedata/tiles/tile-%d.png", i))
		if err != nil {
			return graphs, err
		}
		tileImg, err = helper.DecodeByteToPng(asset)
		if err != nil {
			return graphs, err
		}

		asset, err = assets.Asset(fmt.Sprintf("sourcedata/tiles/tile-shadow-%d.png", i))
		if err != nil {
			return graphs, err
		}
		tileShadowImg, err = helper.DecodeByteToPng(asset)
		if err != nil {
			return graphs, err
		}

		asset, err = assets.Asset(fmt.Sprintf("sourcedata/tiles/tile-mask-%d.png", i))
		if err != nil {
			return graphs, err
		}
		tileMaskImg, err = helper.DecodeByteToPng(asset)
		if err != nil {
			return graphs, err
		}

		graphs = append(graphs, &GraphImage{
			OverlayImage: tileImg,
			ShadowImage:  tileShadowImg,
			MaskImage:    tileMaskImg,
		})
	}

	return graphs, nil
}
