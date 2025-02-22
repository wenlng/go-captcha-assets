package thumbs

import (
	"image"

	assets_1 "github.com/wenlng/go-captcha-assets/bindata/thumbs/thumb_1"
	assets_2 "github.com/wenlng/go-captcha-assets/bindata/thumbs/thumb_2"
	assets_3 "github.com/wenlng/go-captcha-assets/bindata/thumbs/thumb_3"
	assets_4 "github.com/wenlng/go-captcha-assets/bindata/thumbs/thumb_4"
	assets_5 "github.com/wenlng/go-captcha-assets/bindata/thumbs/thumb_5"
	"github.com/wenlng/go-captcha-assets/helper"
)

func GetThumbs() ([]image.Image, error) {
	var images []image.Image
	var asset []byte
	var img image.Image
	var err error

	//
	asset, err = assets_1.Asset("sourcedata/thumbs/thumb-1/thumb.jpg")
	if err != nil {
		return images, err
	}
	img, err = helper.DecodeByteToJpeg(asset)
	if err != nil {
		return images, err
	}
	images = append(images, img)

	//
	asset, err = assets_2.Asset("sourcedata/thumbs/thumb-2/thumb.jpg")
	if err != nil {
		return images, err
	}
	img, err = helper.DecodeByteToJpeg(asset)
	if err != nil {
		return images, err
	}
	images = append(images, img)

	//
	asset, err = assets_3.Asset("sourcedata/thumbs/thumb-3/thumb.jpg")
	if err != nil {
		return images, err
	}
	img, err = helper.DecodeByteToJpeg(asset)
	if err != nil {
		return images, err
	}
	images = append(images, img)

	//
	asset, err = assets_4.Asset("sourcedata/thumbs/thumb-4/thumb.jpg")
	if err != nil {
		return images, err
	}
	img, err = helper.DecodeByteToJpeg(asset)
	if err != nil {
		return images, err
	}
	images = append(images, img)

	//
	asset, err = assets_5.Asset("sourcedata/thumbs/thumb-5/thumb.jpg")
	if err != nil {
		return images, err
	}
	img, err = helper.DecodeByteToJpeg(asset)
	if err != nil {
		return images, err
	}
	images = append(images, img)

	return images, nil
}
