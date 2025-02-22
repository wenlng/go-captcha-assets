package shapes

import (
	"image"

	assets_1 "github.com/wenlng/go-captcha-assets/bindata/shapes/shape_1"
	assets_10 "github.com/wenlng/go-captcha-assets/bindata/shapes/shape_10"
	assets_11 "github.com/wenlng/go-captcha-assets/bindata/shapes/shape_11"
	assets_12 "github.com/wenlng/go-captcha-assets/bindata/shapes/shape_12"
	assets_13 "github.com/wenlng/go-captcha-assets/bindata/shapes/shape_13"
	assets_2 "github.com/wenlng/go-captcha-assets/bindata/shapes/shape_2"
	assets_3 "github.com/wenlng/go-captcha-assets/bindata/shapes/shape_3"
	assets_4 "github.com/wenlng/go-captcha-assets/bindata/shapes/shape_4"
	assets_5 "github.com/wenlng/go-captcha-assets/bindata/shapes/shape_5"
	assets_6 "github.com/wenlng/go-captcha-assets/bindata/shapes/shape_6"
	assets_7 "github.com/wenlng/go-captcha-assets/bindata/shapes/shape_7"
	assets_8 "github.com/wenlng/go-captcha-assets/bindata/shapes/shape_8"
	assets_9 "github.com/wenlng/go-captcha-assets/bindata/shapes/shape_9"
	"github.com/wenlng/go-captcha-assets/helper"
)

func GetShapes() (map[string]image.Image, error) {
	var imgMaps = make(map[string]image.Image, 13)
	var asset []byte
	var img image.Image
	var err error

	//
	asset, err = assets_1.Asset("sourcedata/shapes/shape-1/shape.png")
	if err != nil {
		return imgMaps, err
	}
	img, err = helper.DecodeByteToPng(asset)
	if err != nil {
		return imgMaps, err
	}
	imgMaps["shapes-1"] = img

	//
	asset, err = assets_2.Asset("sourcedata/shapes/shape-2/shape.png")
	if err != nil {
		return imgMaps, err
	}
	img, err = helper.DecodeByteToPng(asset)
	if err != nil {
		return imgMaps, err
	}
	imgMaps["shapes-2"] = img

	//
	asset, err = assets_3.Asset("sourcedata/shapes/shape-3/shape.png")
	if err != nil {
		return imgMaps, err
	}
	img, err = helper.DecodeByteToPng(asset)
	if err != nil {
		return imgMaps, err
	}
	imgMaps["shapes-3"] = img

	//
	asset, err = assets_4.Asset("sourcedata/shapes/shape-4/shape.png")
	if err != nil {
		return imgMaps, err
	}
	img, err = helper.DecodeByteToPng(asset)
	if err != nil {
		return imgMaps, err
	}
	imgMaps["shapes-4"] = img

	//
	asset, err = assets_5.Asset("sourcedata/shapes/shape-5/shape.png")
	if err != nil {
		return imgMaps, err
	}
	img, err = helper.DecodeByteToPng(asset)
	if err != nil {
		return imgMaps, err
	}
	imgMaps["shapes-5"] = img

	//
	asset, err = assets_6.Asset("sourcedata/shapes/shape-6/shape.png")
	if err != nil {
		return imgMaps, err
	}
	img, err = helper.DecodeByteToPng(asset)
	if err != nil {
		return imgMaps, err
	}
	imgMaps["shapes-6"] = img

	//
	asset, err = assets_7.Asset("sourcedata/shapes/shape-7/shape.png")
	if err != nil {
		return imgMaps, err
	}
	img, err = helper.DecodeByteToPng(asset)
	if err != nil {
		return imgMaps, err
	}
	imgMaps["shapes-7"] = img

	//
	asset, err = assets_8.Asset("sourcedata/shapes/shape-8/shape.png")
	if err != nil {
		return imgMaps, err
	}
	img, err = helper.DecodeByteToPng(asset)
	if err != nil {
		return imgMaps, err
	}
	imgMaps["shapes-8"] = img

	//
	asset, err = assets_9.Asset("sourcedata/shapes/shape-9/shape.png")
	if err != nil {
		return imgMaps, err
	}
	img, err = helper.DecodeByteToPng(asset)
	if err != nil {
		return imgMaps, err
	}
	imgMaps["shapes-9"] = img

	//
	asset, err = assets_10.Asset("sourcedata/shapes/shape-10/shape.png")
	if err != nil {
		return imgMaps, err
	}
	img, err = helper.DecodeByteToPng(asset)
	if err != nil {
		return imgMaps, err
	}
	imgMaps["shapes-10"] = img

	//
	asset, err = assets_11.Asset("sourcedata/shapes/shape-11/shape.png")
	if err != nil {
		return imgMaps, err
	}
	img, err = helper.DecodeByteToPng(asset)
	if err != nil {
		return imgMaps, err
	}
	imgMaps["shapes-11"] = img

	//
	asset, err = assets_12.Asset("sourcedata/shapes/shape-12/shape.png")
	if err != nil {
		return imgMaps, err
	}
	img, err = helper.DecodeByteToPng(asset)
	if err != nil {
		return imgMaps, err
	}
	imgMaps["shapes-12"] = img

	//
	asset, err = assets_13.Asset("sourcedata/shapes/shape-13/shape.png")
	if err != nil {
		return imgMaps, err
	}
	img, err = helper.DecodeByteToPng(asset)
	if err != nil {
		return imgMaps, err
	}
	imgMaps["shapes-13"] = img

	return imgMaps, nil
}
