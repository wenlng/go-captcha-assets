# Golang Assets File of Captcha
Go Captcha presets some default embedded resources and stores them in the Go file format. In addition, you can also configure the captcha according to your own needs.

Source Resources File: https://github.com/wenlng/go-captcha-resources


## Install
```shell
$ go get -u github.com/wenlng/go-captcha-assets@latest
```

## Use Assets

### Chinese Text Assets
```go
import "github.com/wenlng/go-captcha-assets/bindata/chars"

func Demo() {
    chars := chars.GetChineseChars()
}
```

### Alpha Text Assets
```go
import "github.com/wenlng/go-captcha-assets/bindata/chars"

func Demo() {
    chars := chars.GetAlphaChars()
}
```

### Font Assets
```go
import "github.com/wenlng/go-captcha-assets/resources/fonts/fzshengsksjw"

func Demo() {
    fonts, err := fzshengsksjw.GetFont()
    if err != nil {
        log.Fatalln(err)
    }
}
```

### Image Assets
```go
//import "github.com/wenlng/go-captcha-assets/resources/images"
import "github.com/wenlng/go-captcha-assets/resources/images_v2"

func Demo() {
    imgs, err := images.GetImages()
    if err != nil {
        log.Fatalln(err)
    }
}
```

### Shape Assets
```go
import "github.com/wenlng/go-captcha-assets/resources/shapes"

func Demo() {
    shapeMaps, err := shapes.GetShapes()
    if err != nil {
        log.Fatalln(err)
    }
}
```

### Thumbnail Assets
```go
import "github.com/wenlng/go-captcha-assets/resources/thumb"

func Demo() {
    imgs, err := thumb.GetImages()
    if err != nil {
        log.Fatalln(err)
    }
}
```

### Tile Assets
```go
import "github.com/wenlng/go-captcha-assets/resources/tiles"

func Demo() {
    graphs, err := tiles.GetTiles()
    if err != nil {
        log.Fatalln(err)
    }
    
    // slide
    var newGraphs = make([]*slide.GraphImage, 0, len(graphs))
    for i := 0; i < len(graphs); i++ {
        graph := graphs[i]
        newGraphs = append(newGraphs, &slide.GraphImage{
            OverlayImage: graph.OverlayImage,
            MaskImage:    graph.MaskImage,
            ShadowImage:  graph.ShadowImage,
        })
    }
}
```

## Load Asset As Needed
```go
// Example
import assets "github.com/wenlng/go-captcha-assets/bindata/images/image_v2_1"

func Demo() error {
    asset, err = assets.Asset("sourcedata/images/image-v2-1/image.jpg")
    if err != nil {
    return err
    }
    img, err = helper.DecodeByteToJpeg(asset)
    if err != nil {
    return err
    }
    fmt.Println(img)
}
```
| Type   | Package Path                   | Asset Path                                | Image                                                                                                                                                        |
|--------|--------------------------------|-------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Image  | bindata/images/image_v2_1      | sourcedata/images/image-v2-1/image.jpg    | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/images/image-v2-1/image.jpg?raw=true" style="width: 130px; height: auto;"/>  |
| Image  | bindata/images/image_v2_2      | sourcedata/images/image-v2-2/image.jpg    | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/images/image-v2-2/image.jpg?raw=true" style="width: 130px; height: auto;"/>  |
| Image  | bindata/images/image_v2_3      | sourcedata/images/image-v2-3/image.jpg    | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/images/image-v2-3/image.jpg?raw=true" style="width: 130px; height: auto;"/>  |
| Image  | bindata/images/image_v2_4      | sourcedata/images/image-v2-4/image.jpg    | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/images/image-v2-4/image.jpg?raw=true" style="width: 130px; height: auto;"/>  |
| Image  | bindata/images/image_v2_5      | sourcedata/images/image-v2-5/image.jpg    | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/images/image-v2-5/image.jpg?raw=true" style="width: 130px; height: auto;"/>  |
| Image  | bindata/images/image_v2_6      | sourcedata/images/image-v2-6/image.jpg    | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/images/image-v2-6/image.jpg?raw=true" style="width: 130px; height: auto;"/>  |
| Image  | bindata/images/image_v2_7      | sourcedata/images/image-v2-7/image.jpg    | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/images/image-v2-7/image.jpg?raw=true" style="width: 130px; height: auto;"/>  |
| Image  | bindata/images/image_v2_8      | sourcedata/images/image-v2-8/image.jpg    | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/images/image-v2-8/image.jpg?raw=true" style="width: 130px; height: auto;"/>  |
| Image  | bindata/images/image_v2_9      | sourcedata/images/image-v2-9/image.jpg    | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/images/image-v2-9/image.jpg?raw=true" style="width: 130px; height: auto;"/>  |
| Image  | bindata/images/image_v2_10     | sourcedata/images/image-v2-10/image.jpg   | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/images/image-v2-10/image.jpg?raw=true" style="width: 130px; height: auto;"/> |
| Image  | bindata/images/image_v2_11     | sourcedata/images/image-v2-11/image.jpg   | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/images/image-v2-11/image.jpg?raw=true" style="width: 130px; height: auto;"/> |
| Image  | bindata/images/image_v2_12     | sourcedata/images/image-v2-12/image.jpg   | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/images/image-v2-12/image.jpg?raw=true" style="width: 130px; height: auto;"/> |
| -      |                                |                                           |                                                                                                                                                              |
| Thumb  | bindata/thumbs/thumb_1         | sourcedata/thumbs/thumb-1/thumb.jpg       | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/thumbs/thumb-1/thumb.jpg?raw=true" style="width: 130px; height: auto;"/>     |
| Thumb  | bindata/thumbs/thumb_2         | sourcedata/thumbs/thumb-2/thumb.jpg       | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/thumbs/thumb-2/thumb.jpg?raw=true" style="width: 130px; height: auto;"/>     |
| Thumb  | bindata/thumbs/thumb_3         | sourcedata/thumbs/thumb-3/thumb.jpg       | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/thumbs/thumb-3/thumb.jpg?raw=true" style="width: 130px; height: auto;"/>     |
| Thumb  | bindata/thumbs/thumb_4         | sourcedata/thumbs/thumb-4/thumb.jpg       | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/thumbs/thumb-4/thumb.jpg?raw=true" style="width: 130px; height: auto;"/>     |
| Thumb  | bindata/thumbs/thumb_5         | sourcedata/thumbs/thumb-5/thumb.jpg       | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/thumbs/thumb-5/thumb.jpg?raw=true" style="width: 130px; height: auto;"/>     |
| -      |                                |                                           |                                                                                                                                                              |
| Tile   | bindata/tiles/tile_1           | sourcedata/tiles/tile-1/tile.png          | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/tiles/tile-1/tile.png?raw=true" style="width: 80px; height: auto;"/>         |
| Tile   | bindata/tiles/tile_1           | sourcedata/tiles/tile-1/tile-shadow.png   | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/tiles/tile-1/tile-shadow.png?raw=true" style="width: 80px; height: auto;"/>  |
| Tile   | bindata/tiles/tile_1           | sourcedata/tiles/tile-1/tile-mask.png     | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/tiles/tile-1/tile-mask.png?raw=true" style="width: 80px; height: auto;"/>    |
| Tile   | bindata/tiles/tile_2           | sourcedata/tiles/tile-2/tile.png          | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/tiles/tile-2/tile.png?raw=true" style="width: 80px; height: auto;"/>         |
| Tile   | bindata/tiles/tile_2           | sourcedata/tiles/tile-2/tile-shadow.png   | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/tiles/tile-2/tile-shadow.png?raw=true" style="width: 80px; height: auto;"/>  |
| Tile   | bindata/tiles/tile_2           | sourcedata/tiles/tile-2/tile-mask.png     | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/tiles/tile-2/tile-mask.png?raw=true" style="width: 80px; height: auto;"/>    |
| Tile   | bindata/tiles/tile_3           | sourcedata/tiles/tile-3/tile.png          | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/tiles/tile-3/tile.png?raw=true" style="width: 80px; height: auto;"/>         |
| Tile   | bindata/tiles/tile_3           | sourcedata/tiles/tile-3/tile-shadow.png   | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/tiles/tile-3/tile-shadow.png?raw=true" style="width: 80px; height: auto;"/>  |
| Tile   | bindata/tiles/tile_3           | sourcedata/tiles/tile-3/tile-mask.png     | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/tiles/tile-3/tile-mask.png?raw=true" style="width: 80px; height: auto;"/>    |
| Tile   | bindata/tiles/tile_4           | sourcedata/tiles/tile-4/tile.png          | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/tiles/tile-4/tile.png?raw=true" style="width: 80px; height: auto;"/>         |
| Tile   | bindata/tiles/tile_4           | sourcedata/tiles/tile-4/tile-shadow.png   | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/tiles/tile-4/tile-shadow.png?raw=true" style="width: 80px; height: auto;"/>  |
| Tile   | bindata/tiles/tile_4           | sourcedata/tiles/tile-4/tile-mask.png     | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/tiles/tile-4/tile-mask.png?raw=true" style="width: 80px; height: auto;"/>    |
| -      |                                |                                           |                                                                                                                                                              |
| Shape  | bindata/shapes/shape_1         | sourcedata/shapes/shape-1/shape.png       | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/shapes/shape-1/shape.png?raw=true" style="width: 80px; height: auto;"/>      |
| Shape  | bindata/shapes/shape_2         | sourcedata/shapes/shape-2/shape.png       | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/shapes/shape-2/shape.png?raw=true" style="width: 80px; height: auto;"/>      |
| Shape  | bindata/shapes/shape_3         | sourcedata/shapes/shape-3/shape.png       | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/shapes/shape-3/shape.png?raw=true" style="width: 80px; height: auto;"/>      |
| Shape  | bindata/shapes/shape_4         | sourcedata/shapes/shape-4/shape.png       | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/shapes/shape-4/shape.png?raw=true" style="width: 80px; height: auto;"/>      |
| Shape  | bindata/shapes/shape_5         | sourcedata/shapes/shape-5/shape.png       | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/shapes/shape-5/shape.png?raw=true" style="width: 80px; height: auto;"/>      |
| Shape  | bindata/shapes/shape_6         | sourcedata/shapes/shape-6/shape.png       | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/shapes/shape-6/shape.png?raw=true" style="width: 80px; height: auto;"/>      |
| Shape  | bindata/shapes/shape_7         | sourcedata/shapes/shape-7/shape.png       | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/shapes/shape-7/shape.png?raw=true" style="width: 80px; height: auto;"/>      |
| Shape  | bindata/shapes/shape_8         | sourcedata/shapes/shape-8/shape.png       | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/shapes/shape-8/shape.png?raw=true" style="width: 80px; height: auto;"/>      |
| Shape  | bindata/shapes/shape_9         | sourcedata/shapes/shape-9/shape.png       | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/shapes/shape-9/shape.png?raw=true" style="width: 80px; height: auto;"/>      |
| Shape  | bindata/shapes/shape_10        | sourcedata/shapes/shape-10/shape.png      | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/shapes/shape-10/shape.png?raw=true" style="width: 80px; height: auto;"/>     |
| Shape  | bindata/shapes/shape_11        | sourcedata/shapes/shape-11/shape.png      | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/shapes/shape-11/shape.png?raw=true" style="width: 80px; height: auto;"/>     |
| Shape  | bindata/shapes/shape_12        | sourcedata/shapes/shape-12/shape.png      | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/shapes/shape-12/shape.png?raw=true" style="width: 80px; height: auto;"/>     |
| Shape  | bindata/shapes/shape_13        | sourcedata/shapes/shape-13/shape.png      | <img src="https://github.com/wenlng/go-captcha-resources/blob/master/sourcedata/shapes/shape-13/shape.png?raw=true" style="width: 80px; height: auto;"/>     |
| -      |                                |                                           |                                                                                                                                                              |
| Font   | bindata/fonts/fzshengsksjw_cu  | sourcedata/fonts/fzshengsksjw_cu/font.ttf | <p style="width: 80px; height: 50px;"></p>                                                                                                                   |




