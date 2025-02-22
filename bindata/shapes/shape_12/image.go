// Code generated for package shape_12 by go-bindata DO NOT EDIT. (@generated)
// sources:
// sourcedata/shapes/shape-12/shape.png
package shape_12

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

var _sourcedataShapesShape12ShapePng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xea\x0c\xf0\x73\xe7\xe5\x92\xe2\x62\x60\x60\xe0\xf5\xf4\x70\x09\x62\x60\x60\xc8\x01\x61\x0e\x36\x06\x06\x86\xfe\xb4\xf0\xb3\x0c\x0c\x4c\xb5\x9e\x2e\x8e\x21\x15\x73\xde\xde\xdd\x7d\xf1\x76\x00\x8f\xcb\x43\x65\xd3\x07\x6e\x22\xbb\x19\x58\xf5\x6e\x35\x72\x17\x9c\xd1\xbb\xd3\xa2\x7c\xe2\x9f\xbb\xbe\x1d\x43\x8c\xee\x4e\xf5\x94\x43\xd6\x9f\x9f\xde\x10\x6e\xde\x15\xa3\x77\xbf\x61\xfe\x81\xef\x81\xf9\x11\x37\x59\x62\x96\x67\xfe\x32\x7e\xf4\x7c\xb9\xa0\xbd\xe8\xb1\xae\xb6\x85\x5b\x9b\xdf\x32\x68\x4b\x35\x5d\xe3\xda\xa6\xbc\x66\x85\xec\x5a\xde\x9c\xf6\x88\x40\xed\x28\xee\x1b\x7d\x1b\x17\x99\x6d\x29\xda\xc0\x9c\xe0\xaa\x9b\xb1\x34\x81\xfd\x50\xd7\xa9\x8b\xa7\x0e\xf0\x75\x88\x4e\x5d\x34\xb5\x41\x56\x42\x5b\xcb\xd9\x9b\xdd\xf8\xda\x5a\xa1\x2d\x21\x06\xc9\x77\xa7\xf6\x4d\x63\x4b\x5b\xb3\x7c\xe9\xa2\xa5\x0d\x92\xaf\x73\x33\x33\x66\x1f\xec\x31\xdd\x2d\xb2\xc5\x37\xe1\x58\xd4\x9d\xb4\x2d\x57\x13\x8e\x45\xdd\x55\xdb\x12\x9a\x70\xec\xd6\x9d\xac\x2d\x47\x13\x8e\x79\x65\x8b\xde\x52\x32\x48\x5e\xba\x5d\x22\xd7\x6c\xc6\xc6\x53\x51\x2d\x66\xdc\x67\x22\x9e\x9e\x75\xab\xe6\x4e\x5b\xb3\x2d\x6d\xe1\x5e\x16\xe3\xb0\xea\xfc\x8b\xf3\x9a\x24\xb3\x23\xf4\x2f\xca\x35\x49\x66\x57\xf8\x5f\x7c\x77\xa0\xe7\xab\xa1\xad\xcb\x6f\xa6\xb4\x35\xfd\xaf\xf9\xde\xb2\xa5\xad\xe9\x7f\xcb\x7d\x8b\xb3\x47\xf5\x02\xef\xdb\x14\x9d\x99\x6a\x0f\xaa\x8e\x45\x68\xc5\x36\xc4\xa7\xf7\x2c\xbd\xf2\xd4\xe6\x17\x6b\x5e\x24\x9b\xf0\xe9\x84\x24\x49\x1f\x9d\xe6\xce\x54\x83\x8d\x3c\x6a\xd3\x0e\x1e\x31\x9a\x11\x41\x33\xc1\xb9\x8d\xf7\x17\xcb\x3f\x4c\x33\x9f\xdb\x28\x2f\xe1\xbb\x90\x73\xc7\xaf\x03\x4b\xdd\xfe\xf3\xc5\x1c\x13\x5f\x5c\x9b\xf3\x5d\x3d\x45\x79\x71\x6d\xcc\xf9\x7e\x13\x77\xb7\xdf\x33\xee\xcd\x77\xec\xef\x31\xdb\xf3\xd3\x99\x79\xe2\xc5\xb6\xad\x35\xa7\x94\x8d\xa2\xaa\x5f\x1f\x2f\x52\x5a\xbc\x75\x6b\x6d\x9f\x96\x51\x54\x65\x65\x73\xfe\xc4\x8b\xa7\x4e\xcd\x0f\xe9\xe9\x99\x16\x77\x37\x23\x33\x65\xcb\x96\x57\xe9\x8d\xca\x8b\x73\xd7\xe5\x57\x58\x1b\xdd\x7a\xf5\x6b\xf9\x5b\xe5\xc5\xba\x79\xc2\x62\x39\x2c\xcc\x8b\x75\x33\x8c\xc5\x54\x1a\x1a\x2f\xa6\x4d\x69\x62\xf5\x4c\x48\xd8\x22\xbe\x34\xd1\x51\x49\x42\xe2\xad\xe3\xca\x88\x09\x3c\x6c\x6e\xbf\xd7\x97\xe8\x68\x30\x30\x2f\x0e\x6c\x61\x75\x4d\x28\xc8\xd8\x62\xa4\xa3\xc1\xd2\xec\xbc\x68\x6a\x92\xa3\xa1\x84\x0c\xd7\x5a\x19\x6e\xa9\x03\x0f\x04\x32\x74\x23\x2e\xb0\xb0\xed\x76\xcb\x3e\xdc\xe4\x66\x70\x43\x35\x4a\x5b\xcc\xa4\xf9\xe0\xba\x0b\xb7\x36\x26\x30\x48\xbc\x61\xca\xe5\x10\xe3\x61\x62\xce\xdc\x92\x18\xc1\x50\xcf\xc8\xe7\x5b\xe8\x72\xf2\xf4\x7f\x4e\x06\x06\x06\x06\x4f\x57\x3f\x97\x75\x4e\x09\x4d\x80\x00\x00\x00\xff\xff\x75\x2f\xf8\x15\xb6\x02\x00\x00")

func sourcedataShapesShape12ShapePngBytes() ([]byte, error) {
	return bindataRead(
		_sourcedataShapesShape12ShapePng,
		"sourcedata/shapes/shape-12/shape.png",
	)
}

func sourcedataShapesShape12ShapePng() (*asset, error) {
	bytes, err := sourcedataShapesShape12ShapePngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sourcedata/shapes/shape-12/shape.png", size: 694, mode: os.FileMode(420), modTime: time.Unix(1715607003, 0)}
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
	"sourcedata/shapes/shape-12/shape.png": sourcedataShapesShape12ShapePng,
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
			"shape-12": &bintree{nil, map[string]*bintree{
				"shape.png": &bintree{sourcedataShapesShape12ShapePng, map[string]*bintree{}},
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
