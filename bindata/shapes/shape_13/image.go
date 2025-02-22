// Code generated for package shape_13 by go-bindata DO NOT EDIT. (@generated)
// sources:
// sourcedata/shapes/shape-13/shape.png
package shape_13

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

var _sourcedataShapesShape13ShapePng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\xd4\x07\x2b\xf8\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52\x00\x00\x00\x78\x00\x00\x00\x78\x08\x06\x00\x00\x00\x39\x64\x36\xd2\x00\x00\x07\x9b\x49\x44\x41\x54\x78\x9c\xed\xdd\x4b\x8b\x15\x47\x18\xc6\xf1\xbf\x89\x0e\x22\x2a\xc1\xc0\x11\x25\x12\x83\x8b\x71\xf1\x9a\x11\xc1\x0b\x2a\xe2\x4e\xfd\x00\xba\xf0\x82\x1b\x51\x04\x89\x20\x2e\xe2\xc6\x85\x0b\x73\x41\x04\x45\x10\x83\x90\x85\xc6\x84\xf1\x03\xe8\xec\x42\x18\x5d\x28\x0c\x1a\x6b\xe1\x40\x44\x45\x19\xb1\x61\x5c\x18\x11\xe3\x4c\x48\x16\x55\xca\x41\xfb\xf4\xb5\xaa\x2f\x33\xef\x6f\xa9\x33\xd5\xef\xe9\x67\xaa\x4f\x5f\xaa\xab\x40\x29\xa5\x94\x52\x4a\x29\x95\xcf\x8c\xba\x0b\xf0\x4d\x44\x3e\x05\x96\x00\x4b\x81\x45\xc0\xe7\xc0\x02\x60\x2e\xd0\xd7\xf5\xa3\x6f\x81\x57\xc0\x0b\x60\x1c\x78\x06\x3c\x02\x9e\x18\x63\xfe\xad\xae\xe2\xb0\x5a\x1d\xb0\x88\xcc\x06\x56\x03\xeb\x81\x95\xc0\x0a\xa0\x1f\x98\x59\xa2\xd9\x49\x60\x14\xb8\x07\xdc\x01\x6e\x02\xb7\x8d\x31\x6f\xca\x55\x5b\x8f\xd6\x05\x2c\x22\x03\xc0\x56\x60\x1b\x36\xd8\x59\x15\x6c\x76\x02\x1b\xf4\x35\xe0\xba\x31\xe6\x6e\x05\xdb\xf4\xa2\x15\x01\x8b\x48\x3f\xb0\x1b\xd8\x8e\xed\xa1\x75\x1b\x05\xae\x02\x97\x8d\x31\xa3\x75\x17\x93\xa4\xb1\x01\x8b\x48\x1f\xb0\x03\x38\x00\x6c\xac\xb9\x9c\x24\xc3\xc0\x05\x60\xd0\x18\xf3\xb6\xee\x62\x3e\xd4\xb8\x80\x45\x64\x1e\x70\x10\x38\x0c\x2c\xae\xb9\x9c\x3c\xc6\x80\x33\xc0\x79\x63\xcc\xdf\x75\x17\xf3\x4e\x63\x02\x16\x91\xb9\xc0\x21\xe0\x28\xf6\xcc\xb7\xad\xc6\x81\x53\xc0\x39\x63\xcc\xab\xba\x8b\xa9\x3d\x60\x11\x99\x01\xec\x01\xbe\xa3\x5d\x3d\x36\xcd\x18\x70\x0c\xb8\x64\x8c\xf9\xaf\xae\x22\x6a\x0d\x58\x44\x56\x00\x17\x81\x35\x75\xd6\x11\xd8\x2d\x60\x9f\x31\xe6\x5e\x1d\x1b\xaf\x25\x60\x11\x99\x05\x1c\x07\xbe\xa5\xdc\x35\x6b\x5b\x4c\x02\xdf\x03\x27\x8c\x31\x13\x55\x6e\xb8\xf2\x80\xdd\x25\xcf\x15\x60\x55\xd5\xdb\x6e\x80\x11\x60\x67\x95\x97\x56\x9f\x54\xb5\x21\x00\x11\xd9\x85\xfd\x90\xd3\x31\x5c\xb0\x9f\x7b\xc4\xed\x87\x4a\x54\xd2\x83\xdd\xfd\xe1\xd3\xc0\x37\x55\x6c\xaf\x25\xce\x02\x47\x42\xdf\xf7\x0e\x1e\xb0\x88\xcc\x07\x06\x81\x2d\xa1\xb7\xd5\x42\x43\xc0\x0e\x63\xcc\xcb\x50\x1b\x08\x1a\xb0\x88\x2c\xc4\x7e\x88\x81\x90\xdb\x69\xb9\xbb\xc0\x16\x63\xcc\xf3\x10\x8d\x07\x0b\x58\x44\xbe\x00\x7e\x07\x96\x85\xda\xc6\x14\xf2\x00\xd8\x6c\x8c\x79\xea\xbb\xe1\x20\x01\xbb\x70\x87\x81\x2f\x43\xb4\x3f\x45\x3d\x06\x36\xfa\x0e\xd9\x7b\xc0\xee\xb0\x7c\x03\xed\xb9\x45\x3c\x00\x36\xf8\x3c\x5c\x7b\xbd\x4c\x72\x27\x54\x43\x68\xb8\x45\x2d\x03\x86\xdc\x7e\xf4\xc2\x5b\xc0\xee\x52\x68\x10\x3d\xa1\x2a\x6b\x00\x18\x74\xfb\xb3\x34\x9f\x3d\xf8\x34\x7a\x29\xe4\xcb\x16\xec\xfe\x2c\xcd\xcb\x5f\x89\xbb\x33\xf3\x83\x8f\xb6\xd4\x7b\x6b\x3b\x9d\xce\x5f\x51\x14\x95\x7a\x48\x51\xfa\x24\xcb\xdd\x5b\x1e\x01\xe6\x94\x6d\x4b\x7d\xe4\x35\xb0\xaa\xcc\xbd\xeb\x52\x87\x68\xf7\x54\xe8\x0a\x1a\x6e\x28\x73\x80\x2b\x6e\x3f\x17\x52\xf6\x3b\xf8\x38\xd3\xf7\xc1\x41\x55\x56\x61\xf7\x73\x21\x85\x0f\xd1\xee\x61\xfd\x08\xd3\xe3\x79\x6e\xdd\x26\xb1\x87\xea\xdc\xdf\xc7\x85\x7a\xb0\x1b\x66\x73\x11\x0d\xb7\x2a\x33\x81\x8b\x6e\xbf\xe7\x52\xf4\x10\xbd\x87\xa9\x3d\xcc\xa6\x89\xd6\x60\xf7\x7b\x2e\xb9\xff\x22\xdc\xe8\xc7\x51\xa6\xd6\x00\xb9\xb6\x18\x03\xfa\xf3\x8c\xd6\x2c\xd2\x83\x0f\xa1\xe1\xd6\x65\x31\x76\xff\x67\x96\xab\x07\xbb\x41\xe9\x0f\x69\xf7\xb8\xe5\xb6\x1b\x07\xbe\xca\x3a\xb8\x3e\x6f\x0f\x3e\x88\x86\x5b\xb7\xcf\xb1\x39\x64\x92\xb9\x07\xbb\x77\x85\x1e\xa2\x87\xe7\x26\x18\xc3\xf6\xe2\xd4\x77\xa1\xf2\xf4\xe0\x1d\x68\xb8\x4d\xb1\x18\x9b\x47\xaa\x3c\x01\x1f\x28\x56\x8b\x0a\x24\x53\x1e\x99\x0e\xd1\xee\x81\xc2\xfd\x52\xe5\xa8\x10\x96\xa7\x3d\x88\xc8\xda\x83\x77\x7b\x28\x46\xf9\x97\x9a\x4b\xd6\x80\xb7\x97\x2c\x44\x85\x91\x9a\x4b\x6a\xc0\x6e\x4e\x8c\x26\x4c\x9b\xa0\x3e\xd6\x2f\x22\x5f\x27\xfd\x40\x96\x1e\xbc\xd5\x53\x31\x2a\x8c\x6d\x49\xff\x99\x25\xe0\xc4\x06\x54\xed\x12\xf3\x49\x3c\x8b\x76\xf3\x50\xbd\xa4\x9a\xa9\x8a\x54\x31\x13\xc0\xfc\x5e\xf3\x78\xa5\xf5\xe0\xd5\x68\xb8\x4d\x37\x0b\x9b\x53\xac\xb4\x80\xd7\xfb\xad\x45\x05\xd2\x33\xa7\xb4\x80\x57\x7a\x2e\x44\x85\xd1\x33\xa7\xb4\x80\x57\x78\x2e\x44\x85\xd1\x33\xa7\x9e\x27\x59\xee\xd5\x89\x37\xe8\xb8\xab\x36\x98\x04\x66\xc7\xcd\x16\x90\xd4\x83\x97\xa0\xe1\xb6\xc5\x4c\x6c\x5e\x1f\x49\x0a\x78\x69\x90\x52\x54\x28\x4b\xe3\xfe\x31\x29\xe0\x45\x61\xea\x50\x81\xc4\xe6\x95\x14\xb0\x0e\xcd\x69\x97\xd8\xbc\x92\x02\x5e\x10\xa8\x10\x15\x46\x6c\x5e\x49\x01\xcf\x0d\x54\x88\x0a\x23\x36\xaf\xa4\x80\xfb\x12\xfe\x4f\x35\x4f\x6c\x5e\x95\x4e\x65\xa8\xaa\xa7\x01\x4f\x71\x49\x01\x37\x6e\xfd\x01\x95\x28\x36\xaf\xa4\x80\x6b\x9f\x8e\x5e\xe5\x12\x9b\x57\x52\xc0\x2f\x02\x15\xa2\xc2\x88\xcd\x2b\x29\xe0\xf1\x40\x85\xa8\x30\x62\xf3\x4a\x0a\xf8\x59\xa0\x42\x54\x18\xb1\x79\x25\x05\xfc\x28\x4c\x1d\x2a\x90\x47\x71\xff\x98\x14\xf0\x13\xec\x73\x46\xd5\x7c\x93\xd8\xbc\x3e\xd2\x33\x60\xf7\xf0\xb8\xd1\xeb\xf2\xa9\xf7\x46\x7b\x2d\x0d\x90\x76\xa3\xa3\x96\xb5\x7e\x54\x6e\x3d\x73\x4a\x0b\xf8\x8e\xe7\x42\x54\x18\x3d\x73\x4a\x0b\xf8\xa6\xe7\x42\x54\x18\x3d\x73\x4a\x0b\xf8\x36\x76\xe4\xbc\x6a\xae\x09\x6c\x4e\xb1\x12\x03\x76\xaf\x43\x68\x2f\x6e\xb6\x9b\x49\xcb\xcf\x67\x79\x9a\x74\xcd\x63\x31\xca\xbf\xc4\x7c\xb2\x04\x7c\xdd\x53\x21\x2a\x8c\xc4\x80\xb3\xce\xd1\x71\x1f\x7d\x09\xbc\x89\x46\x8d\x31\xcb\x93\x7e\x20\xeb\x03\xff\xab\x1e\x8a\x51\xfe\xa5\xe6\x92\x35\xe0\xcb\x25\x0b\x51\x61\xa4\xe6\x92\x29\x60\x37\x55\xcf\x70\xe9\x72\x94\x4f\xc3\x59\xd6\x72\xc8\x33\x26\xeb\x42\x89\x62\x94\x7f\x99\xf2\xc8\x13\xf0\x20\x76\x8e\x44\x55\xbf\x31\x6c\x1e\xa9\x32\x07\xec\x26\xbe\x3c\x53\xb4\x22\xe5\xd5\x99\x2c\x13\x91\x42\xfe\x61\xb3\xe7\xd1\xa1\x3c\x75\x1b\xc7\xe6\x90\x49\xae\x80\xdd\x24\xd4\xa7\xf2\x56\xa4\xbc\x3a\x95\x75\x32\x70\x28\x36\xf0\xfd\x1c\xfa\x5d\x5c\x97\x31\xec\xfe\xcf\x2c\x77\xc0\x6e\x41\x88\x63\x79\x7f\x4f\x79\x71\x2c\xcf\x82\x1c\x50\xfc\xd5\x95\x4b\xc0\xad\x82\xbf\xab\x8a\xb9\x85\xdd\xef\xb9\xe8\xca\x67\xed\x50\x78\xe5\xb3\xc2\xcb\xcb\x46\x51\x14\x75\x3a\x9d\x3e\x60\x53\xd1\x36\x54\x66\x27\x8d\x31\xbf\x15\xf9\xc5\xb2\x6f\x17\x9e\xc0\xf6\x62\x15\xce\x08\x76\x3f\x17\xa2\xeb\x07\x37\x5b\xbd\xeb\x07\xc3\xfb\x07\x11\xfb\xcb\xb6\xa3\x62\xed\x2f\x13\x2e\x78\x5a\xe2\x3d\x8a\xa2\x7b\x9d\x4e\x67\x01\xb0\xd6\x47\x7b\x0a\x80\xb3\xc6\x98\x1f\xcb\x36\xe2\xf3\x0d\xff\x23\xc0\x90\xc7\xf6\xa6\xb3\x21\xec\xfe\x2c\xad\xf4\x77\x70\x37\x11\x99\x0f\xfc\x01\x0c\xf8\x6c\x77\x9a\xb9\x0b\x6c\x32\xc6\xbc\xf4\xd1\x98\xd7\x80\x01\x44\x64\x21\x70\x03\x58\xe6\xbb\xed\x69\xe0\x01\xb0\xc1\x18\xf3\xdc\x57\x83\xde\x27\x61\x71\xc5\x6d\x06\x1e\xfb\x6e\x7b\x8a\x7b\x0c\x6c\xf6\x19\x2e\x04\x9a\x65\xc7\x18\xf3\x14\xd8\x88\xfd\x8b\x54\xe9\x1e\x00\x1b\xdd\x7e\xf3\x2a\xd8\x34\x4a\xae\xd8\x0d\xd8\xef\x14\xd5\xdb\x5d\xec\x61\xd9\x7b\xb8\x10\x78\x9e\x2c\x77\xb8\xd9\x84\x9e\x5d\xf7\x32\x84\x3d\xa1\xf2\x7a\x58\xee\xe6\xe5\x3a\x38\x49\x14\x45\xff\x74\x3a\x9d\x5f\x81\xcf\xd0\xeb\xe4\x6e\x67\x81\xbd\x49\xef\x15\xf9\xe0\xfd\x2c\x3a\x89\x88\xec\x02\x7e\x62\x7a\xdf\xd6\x7c\x8d\xbd\x43\xf5\x4b\x15\x1b\xab\x34\x60\x78\x7f\xef\xfa\x0a\xb0\xaa\xea\x6d\x37\xc0\x08\xb0\xb3\xec\xed\xc7\x3c\x82\x1f\xa2\x3f\x14\x45\xd1\x78\xa7\xd3\xf9\x19\xfb\xc7\xb5\x9e\xe9\x31\x5f\xe6\x24\x70\x12\xd8\x63\x8c\x89\xaa\xdc\x70\xe5\x3d\xb8\x9b\x1b\x34\x70\x11\x58\x53\x67\x1d\x81\xdd\x02\xf6\x15\x79\x58\xef\x43\xad\xbd\xc7\x7d\xe8\x75\xc0\x5e\xa6\xde\x40\xbe\x31\xec\xe7\x5a\x57\x57\xb8\x50\x73\x0f\xee\x26\x22\x73\x81\x43\xc0\x51\xda\xbd\x5e\xc4\x38\x76\x68\xf1\xb9\xbc\x03\xe4\x42\x68\x4c\xc0\xef\x88\xc8\x3c\xe0\x20\x70\x18\x58\x5c\x73\x39\x79\x8c\x61\xdf\xfc\x38\x9f\x67\xdc\x72\x68\x8d\x0b\xf8\x1d\x11\xe9\x03\x76\x00\x07\xb0\xb7\x3d\x9b\x6a\x18\xfb\x22\xd8\x60\xd6\xd7\x49\xaa\xd4\xd8\x80\xbb\xb9\x4b\xab\xdd\xd8\x35\xeb\x9b\x30\xd3\xc0\x28\xf6\xe5\xeb\xcb\x55\x5e\xf2\x14\xd1\x8a\x80\xbb\x89\xc8\x00\x76\xd9\xf9\x6d\xd8\xcb\xac\x2a\xd6\x37\x9e\xc0\xce\x36\x74\x0d\xb8\x66\x8c\xf9\xb3\x82\x6d\x7a\xd1\xba\x80\xbb\xb9\x15\xca\x57\x63\x83\x5e\x89\x5d\x85\xb3\x9f\x72\x63\xb5\x27\xb1\x3d\xf4\x1e\x76\x06\xb9\x9b\xc0\xed\xd0\xb7\x14\x43\x69\x75\xc0\x71\xdc\xaa\xa9\x4b\xb0\x6b\xf9\x2d\xc2\x9e\x91\x2f\xc0\xae\x2b\xd4\xbd\xf4\xcc\x5b\xec\x34\xf8\x2f\xb0\x67\xbe\xcf\xb0\x53\xf2\x3e\xe9\x35\xb1\xa7\x52\x4a\x29\xa5\x94\x52\x00\xfc\x0f\xb1\x72\xe9\x80\x1e\xcf\xe6\x21\x00\x00\x00\x00\x49\x45\x4e\x44\xae\x42\x60\x82\x01\x00\x00\xff\xff\xb1\x80\x99\x77\xd4\x07\x00\x00")

func sourcedataShapesShape13ShapePngBytes() ([]byte, error) {
	return bindataRead(
		_sourcedataShapesShape13ShapePng,
		"sourcedata/shapes/shape-13/shape.png",
	)
}

func sourcedataShapesShape13ShapePng() (*asset, error) {
	bytes, err := sourcedataShapesShape13ShapePngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sourcedata/shapes/shape-13/shape.png", size: 2004, mode: os.FileMode(420), modTime: time.Unix(1715606989, 0)}
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
	"sourcedata/shapes/shape-13/shape.png": sourcedataShapesShape13ShapePng,
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
			"shape-13": &bintree{nil, map[string]*bintree{
				"shape.png": &bintree{sourcedataShapesShape13ShapePng, map[string]*bintree{}},
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
