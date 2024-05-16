package sourcedata

import (
	"fmt"
	"image"

	"github.com/golang/freetype/truetype"
	"github.com/wenlng/go-captcha-assets/helper"
	"github.com/wenlng/go-captcha-assets/sourcedata/chars"
)

type GraphImage struct {
	OverlayImage image.Image
	ShadowImage  image.Image
	MaskImage    image.Image
}

// ResourcesFile .
type ResourcesFile struct {
}

func NewResourcesFile() *ResourcesFile {
	return &ResourcesFile{}
}

func (sf *ResourcesFile) GetChineseChars() []string {
	return chars.GetChineseChars()
}

func (sf *ResourcesFile) GetAlphaChars() []string {
	return chars.GetAlphaChars()
}

func (sf *ResourcesFile) GetFonts() (fonts []*truetype.Font, err error) {
	font, err := helper.LoadFont("./resources/fonts/fzshengsksjw_cu.ttf")
	if err != nil {
		return fonts, err
	}

	return []*truetype.Font{
		font,
	}, nil
}

func (sf *ResourcesFile) GetImages() ([]image.Image, error) {
	var images []image.Image
	var err error
	for i := 1; i <= 5; i++ {
		var img image.Image
		img, err = helper.LoadJpeg(fmt.Sprintf("resources/images/%d.jpg", i))
		if err != nil {
			return images, err
		}
		images = append(images, img)
	}

	return images, nil
}

func (sf *ResourcesFile) GetThumbImages() ([]image.Image, error) {
	var images []image.Image
	var err error
	for i := 1; i <= 5; i++ {
		var img image.Image
		img, err = helper.LoadJpeg(fmt.Sprintf("resources/images/thumb/r%d.jpg", i))
		if err != nil {
			return images, err
		}
		images = append(images, img)
	}

	return images, nil
}

func (sf *ResourcesFile) GetShapeMaps() (maps map[string]image.Image, err error) {
	shapeImage1, err := helper.LoadPNG("resources/shapes/shape-1.png")
	if err != nil {
		return maps, err
	}

	shapeImage2, err := helper.LoadPNG("resources/shapes/shape-2.png")
	if err != nil {
		return maps, err
	}

	shapeImage3, err := helper.LoadPNG("resources/shapes/shape-3.png")
	if err != nil {
		return maps, err
	}

	shapeImage4, err := helper.LoadPNG("resources/shapes/shape-4.png")
	if err != nil {
		return maps, err
	}

	shapeImage5, err := helper.LoadPNG("resources/shapes/shape-5.png")
	if err != nil {
		return maps, err
	}

	shapeImage6, err := helper.LoadPNG("resources/shapes/shape-6.png")
	if err != nil {
		return maps, err
	}

	shapeImage7, err := helper.LoadPNG("resources/shapes/shape-7.png")
	if err != nil {
		return maps, err
	}

	shapeImage8, err := helper.LoadPNG("resources/shapes/shape-8.png")
	if err != nil {
		return maps, err
	}

	shapeImage9, err := helper.LoadPNG("resources/shapes/shape-9.png")
	if err != nil {
		return maps, err
	}

	shapeImage10, err := helper.LoadPNG("resources/shapes/shape-10.png")
	if err != nil {
		return maps, err
	}

	shapeImage11, err := helper.LoadPNG("resources/shapes/shape-11.png")
	if err != nil {
		return maps, err
	}

	shapeImage12, err := helper.LoadPNG("resources/shapes/shape-12.png")
	if err != nil {
		return maps, err
	}

	shapeImage13, err := helper.LoadPNG("resources/shapes/shape-13.png")
	if err != nil {
		return maps, err
	}

	return map[string]image.Image{
		"shape1":  shapeImage1,
		"shape2":  shapeImage2,
		"shape3":  shapeImage3,
		"shape4":  shapeImage4,
		"shape5":  shapeImage5,
		"shape6":  shapeImage6,
		"shape7":  shapeImage7,
		"shape8":  shapeImage8,
		"shape9":  shapeImage9,
		"shape10": shapeImage10,
		"shape11": shapeImage11,
		"shape12": shapeImage12,
		"shape13": shapeImage13,
	}, nil
}

func (sf *ResourcesFile) GetTileImages() (images []*GraphImage, err error) {
	tileImage1, err := helper.LoadPNG("resources/tiles/tile-1.png")
	if err != nil {
		return images, nil
	}
	tileShadowImage1, err := helper.LoadPNG("resources/tiles/tile-shadow-1.png")
	if err != nil {
		return images, nil
	}
	tileMaskImage1, err := helper.LoadPNG("resources/tiles/tile-mask-1.png")
	if err != nil {
		return images, nil
	}

	tileImage2, err := helper.LoadPNG("resources/tiles/tile-2.png")
	if err != nil {
		return images, nil
	}
	tileShadowImage2, err := helper.LoadPNG("resources/tiles/tile-shadow-2.png")
	if err != nil {
		return images, nil
	}
	tileMaskImage2, err := helper.LoadPNG("resources/tiles/tile-mask-2.png")
	if err != nil {
		return images, nil
	}

	tileImage3, err := helper.LoadPNG("resources/tiles/tile-3.png")
	if err != nil {
		return images, nil
	}
	tileShadowImage3, err := helper.LoadPNG("resources/tiles/tile-shadow-3.png")
	if err != nil {
		return images, nil
	}
	tileMaskImage3, err := helper.LoadPNG("resources/tiles/tile-mask-3.png")
	if err != nil {
		return images, nil
	}

	tileImage4, err := helper.LoadPNG("resources/tiles/tile-4.png")
	if err != nil {
		return images, nil
	}
	tileShadowImage4, err := helper.LoadPNG("resources/tiles/tile-shadow-4.png")
	if err != nil {
		return images, nil
	}
	tileMaskImage4, err := helper.LoadPNG("resources/tiles/tile-mask-4.png")
	if err != nil {
		return images, nil
	}

	return []*GraphImage{
		{
			OverlayImage: tileImage1,
			ShadowImage:  tileShadowImage1,
			MaskImage:    tileMaskImage1,
		},
		{
			OverlayImage: tileImage2,
			ShadowImage:  tileShadowImage2,
			MaskImage:    tileMaskImage2,
		},
		{
			OverlayImage: tileImage3,
			ShadowImage:  tileShadowImage3,
			MaskImage:    tileMaskImage3,
		},
		{
			OverlayImage: tileImage4,
			ShadowImage:  tileShadowImage4,
			MaskImage:    tileMaskImage4,
		},
	}, nil
}
