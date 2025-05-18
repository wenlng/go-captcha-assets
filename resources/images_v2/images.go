package images

import (
	"image"

	"github.com/wenlng/go-captcha-assets/resources/imagesv2"
)

// Deprecated: As of 1.1.0, it will be removed, please use [imagesv2.GetImages()]
func GetImages() ([]image.Image, error) {
	return imagesv2.GetImages()
}
