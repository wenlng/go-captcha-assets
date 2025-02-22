# Captcha Assets
Go Captcha presets some default embedded resources and stores them in the Go file format. In addition, you can also configure the captcha according to your own needs.

## Install
```shell
$ go get -u github.com/wenlng/go-captcha-assets@latest
```

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
import "github.com/wenlng/go-captcha-assets/resources/fonts/yrdzst"

func Demo() {
    fonts, err := fzshengsksjw.GetFont()
    if err != nil {
        log.Fatalln(err)
    }
    
    // OR
    
    fonts, err := yrdzst.GetFont()
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