// Code generated for package shape_4 by go-bindata DO NOT EDIT. (@generated)
// sources:
// sourcedata/shapes/shape-4/shape.png
package shape_4

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _sourcedataShapesShape4ShapePng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xea\x0c\xf0\x73\xe7\xe5\x92\xe2\x62\x60\x60\xe0\xf5\xf4\x70\x09\x62\x60\x60\xa8\x00\x61\x0e\x36\x06\x06\x06\xcb\x14\xb3\x4b\x0c\x0c\x4c\x2f\x3d\x5d\x1c\x43\x2a\xe6\xbc\xbd\xb1\x71\x95\x6c\xa8\x48\xfb\x47\xfd\x9d\x81\xcd\x81\xdc\xdd\x99\x6b\x84\x1c\x75\xca\x54\x25\xbc\xbb\x39\xbc\x1c\x04\xd7\xb0\xc8\x1f\x8c\x53\xc9\x32\x98\xe4\x24\x50\xc3\xe9\xa8\xab\x97\xd4\x92\x30\xd1\xa6\xc1\x53\xdc\xac\xec\xd8\x2d\x4d\x97\x45\xe2\xc9\x79\x7a\x11\x1d\x4b\x0d\xb2\xe7\x5b\x5a\xfe\x38\xff\x7d\xf5\xf5\x5b\xe9\x0f\x92\x0f\x9f\x61\xf4\x9c\xa4\x02\x42\xb7\x76\x26\x67\xcb\x5d\x8b\xaa\x96\xb2\xe0\x9f\xc3\xea\xc7\xec\xbf\xae\xb2\x6e\x07\xf3\x03\xb3\x02\xe1\x1f\xfb\xbe\xdb\xad\x2f\x7c\xa1\xd5\x18\x9d\x7d\x9e\x59\xbc\x71\x7b\x83\xfb\x35\x8b\xcf\x8a\xe5\xf2\x3f\xc4\x76\x5c\xa9\xde\xbf\xf5\xa4\xa3\x34\x6b\xde\xda\xb7\x8c\xe2\x0d\xfc\x06\x6f\xcc\x64\x18\xe7\x33\xca\x7f\xbb\xc7\xef\xce\x27\xc1\x77\x27\xaa\x5a\xe6\x07\x57\x45\xed\x17\x1e\xdd\x32\xde\x8a\xcd\xdf\xeb\x02\x62\x65\xb8\xee\x44\x59\xf3\xd5\x30\xef\x31\xdc\xb4\x5e\xf1\x7a\xc3\xe7\xab\xb3\x6b\x6a\x3e\x1c\xbf\x77\x7c\x07\xff\x2d\xed\x03\xea\x57\x97\x7e\xe2\xcd\x5b\x6b\xcb\xfd\x87\xc3\x26\xea\xee\xd5\x63\x7d\x0d\xe2\x5a\xfa\x71\xca\xe6\x3f\x98\x8e\xdd\x9d\x6a\xfb\x60\xf1\x29\xa7\x97\x6c\xdb\xaf\xcd\x63\xaa\x63\xf4\x0b\xfb\xb6\x65\xd5\xef\x03\xc6\xde\xfb\x27\xaf\x36\xff\x21\x73\xec\xee\x56\xe1\x03\xd9\xde\xed\xab\x1b\xcf\x8b\x3d\xdf\x2a\xab\xaf\x27\xbb\x86\xcd\x2e\x78\xce\xf5\x32\x9d\x77\x01\xbc\x7f\x22\x0a\xe3\x6e\x17\x66\xbd\x7c\xf0\xa2\x7c\xc3\x3f\x75\x1e\x99\x73\x07\x1e\x0e\x39\x62\x9d\xd4\x8f\x77\x61\xf3\x9c\x42\xf7\xeb\xcb\xae\x61\xdb\x17\x3b\xe7\x7a\xd9\x9f\x3d\x25\x1b\xa4\x7e\xb4\x3f\xda\xb3\xed\xf1\xb1\x4b\x07\xb3\x1e\xbc\x3e\xe0\x58\xe8\xfe\xda\x8a\x65\x9d\x9e\xac\x1e\xeb\xbb\xbd\xff\x19\x6b\x02\x6b\x95\xd6\x66\xed\xff\xfe\xef\xf5\xa5\x87\x13\xb7\xbf\xf0\xff\x98\x9a\x1d\xce\xd7\xd8\xaf\xcd\x5d\x7d\x35\x7e\xc6\xf3\x9f\xdb\x33\x67\xfd\x63\x89\x5b\xfb\x2e\x2b\x79\x7a\xd5\xfb\xdf\x53\x5b\x2f\xbe\x7f\x70\xbe\x71\x3d\x7b\xb9\xe9\x33\xeb\x07\x45\xa5\x2b\x6b\x9a\x19\x9f\x55\xa6\xcd\x5a\xa6\x7d\xa0\xd8\x41\xda\xd2\x59\x35\xbc\x61\x7b\xf5\xfb\xef\xab\xba\x9d\x1b\x19\x9f\x55\xda\xcd\x5a\x96\x29\xbf\x83\x79\x0f\xcf\x1c\x09\x5b\xef\xff\x2f\xe5\x0a\x04\x2b\xca\x55\xf4\x0e\x6c\x34\x9d\xb7\x68\x6b\x99\x54\x05\x7f\x05\xbb\xc5\xd6\x63\x33\xf9\x77\xee\x3e\x30\xfd\xc0\xd3\x7b\xbf\xfd\x67\xcb\x6c\xf1\x7a\x50\xbf\xf5\xda\x3c\x86\x73\x8c\xfa\x8c\xe9\x0d\x93\xb3\xe7\xc9\xbf\xd3\xff\xf7\x73\xf9\x07\xb6\x82\xbd\xdf\xeb\xd2\x8f\x9f\x57\x86\xa4\x41\x5b\x87\x7a\x66\x66\xf3\x92\xf2\xdf\x31\xcb\x57\x31\x30\x30\x30\x78\xba\xfa\xb9\xac\x73\x4a\x68\x02\x04\x00\x00\xff\xff\x07\x35\xb0\x4b\x22\x03\x00\x00")

func sourcedataShapesShape4ShapePngBytes() ([]byte, error) {
	return bindataRead(
		_sourcedataShapesShape4ShapePng,
		"sourcedata/shapes/shape-4/shape.png",
	)
}

func sourcedataShapesShape4ShapePng() (*asset, error) {
	bytes, err := sourcedataShapesShape4ShapePngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sourcedata/shapes/shape-4/shape.png", size: 802, mode: os.FileMode(420), modTime: time.Unix(1715607091, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"sourcedata/shapes/shape-4/shape.png": sourcedataShapesShape4ShapePng,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"sourcedata": &bintree{nil, map[string]*bintree{
		"shapes": &bintree{nil, map[string]*bintree{
			"shape-4": &bintree{nil, map[string]*bintree{
				"shape.png": &bintree{sourcedataShapesShape4ShapePng, map[string]*bintree{}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
