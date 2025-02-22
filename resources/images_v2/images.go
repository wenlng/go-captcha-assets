package images

import (
	"image"

	assets_1 "github.com/wenlng/go-captcha-assets/bindata/images/image_v2_1"
	assets_2 "github.com/wenlng/go-captcha-assets/bindata/images/image_v2_2"
	assets_3 "github.com/wenlng/go-captcha-assets/bindata/images/image_v2_3"
	assets_4 "github.com/wenlng/go-captcha-assets/bindata/images/image_v2_4"
	assets_5 "github.com/wenlng/go-captcha-assets/bindata/images/image_v2_5"
	assets_6 "github.com/wenlng/go-captcha-assets/bindata/images/image_v2_6"
	assets_7 "github.com/wenlng/go-captcha-assets/bindata/images/image_v2_7"
	assets_8 "github.com/wenlng/go-captcha-assets/bindata/images/image_v2_8"
	"github.com/wenlng/go-captcha-assets/helper"
)

func GetImages() ([]image.Image, error) {
	var images []image.Image
	var asset []byte
	var img image.Image
	var err error

	//
	asset, err = assets_1.Asset("sourcedata/images/image-v2-1/image.jpg")
	if err != nil {
		return images, err
	}
	img, err = helper.DecodeByteToJpeg(asset)
	if err != nil {
		return images, err
	}
	images = append(images, img)

	//
	asset, err = assets_2.Asset("sourcedata/images/image-v2-2/image.jpg")
	if err != nil {
		return images, err
	}
	img, err = helper.DecodeByteToJpeg(asset)
	if err != nil {
		return images, err
	}
	images = append(images, img)

	//
	asset, err = assets_3.Asset("sourcedata/images/image-v2-3/image.jpg")
	if err != nil {
		return images, err
	}
	img, err = helper.DecodeByteToJpeg(asset)
	if err != nil {
		return images, err
	}
	images = append(images, img)

	//
	asset, err = assets_4.Asset("sourcedata/images/image-v2-4/image.jpg")
	if err != nil {
		return images, err
	}
	img, err = helper.DecodeByteToJpeg(asset)
	if err != nil {
		return images, err
	}
	images = append(images, img)

	//
	asset, err = assets_5.Asset("sourcedata/images/image-v2-5/image.jpg")
	if err != nil {
		return images, err
	}
	img, err = helper.DecodeByteToJpeg(asset)
	if err != nil {
		return images, err
	}
	images = append(images, img)

	//
	asset, err = assets_6.Asset("sourcedata/images/image-v2-6/image.jpg")
	if err != nil {
		return images, err
	}
	img, err = helper.DecodeByteToJpeg(asset)
	if err != nil {
		return images, err
	}
	images = append(images, img)

	//
	asset, err = assets_7.Asset("sourcedata/images/image-v2-7/image.jpg")
	if err != nil {
		return images, err
	}
	img, err = helper.DecodeByteToJpeg(asset)
	if err != nil {
		return images, err
	}
	images = append(images, img)

	//
	asset, err = assets_8.Asset("sourcedata/images/image-v2-8/image.jpg")
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
