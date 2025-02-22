package fzshengsksjw

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	assets "github.com/wenlng/go-captcha-assets/bindata/fonts/fzshengsksjw_cu"
)

func GetFont() (font *truetype.Font, err error) {
	asset, err := assets.Asset("sourcedata/fonts/fzshengsksjw_cu/font.ttf")
	if err != nil {
		return font, err
	}

	font, err = freetype.ParseFont(asset)
	if err != nil {
		return nil, err
	}

	return font, nil
}
