// Code generated for package shape_1 by go-bindata DO NOT EDIT. (@generated)
// sources:
// sourcedata/shapes/shape-1/shape.png
package shape_1

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

var _sourcedataShapesShape1ShapePng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\xcd\x0b\x32\xf4\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52\x00\x00\x00\x72\x00\x00\x00\x72\x08\x06\x00\x00\x00\x8f\xdd\x85\x7d\x00\x00\x0b\x94\x49\x44\x41\x54\x78\x9c\xed\x9d\x79\xb4\xd5\x55\x15\xc7\x3f\xef\xf9\x28\x07\x06\x87\x78\x2a\x38\x90\xc8\xa0\x6e\x05\x0d\x25\x5c\x8a\x53\xa4\x89\xf3\x94\x2b\x4d\x97\x26\x2b\x4d\x5d\xe5\x80\xe9\xaa\x94\x6c\x65\x85\xf6\x47\x80\x99\x39\x52\xe4\x98\x88\x26\xa9\x15\x89\x81\x2b\x35\x15\x91\x2d\x1a\x26\x8b\x44\xe9\xf5\x78\xa8\x3d\x41\x14\x14\xfa\x63\x9f\x0b\x97\xc7\x1d\x7e\xc3\xf9\x0d\xf7\xbe\xfb\x59\xeb\x2d\xe0\x77\x7f\xe7\x9c\xcd\xfb\xde\x73\x7e\xe7\x77\xf6\x3e\xfb\x34\xd1\x8d\x11\x91\xad\x80\xe3\x81\x23\x80\x91\xc0\x50\x60\x2b\xf7\x71\x1b\xf0\x0a\xf0\x38\x30\x4d\x55\x3b\x32\x31\x32\x20\x4d\x59\x1b\x90\x05\x22\xb2\x27\x70\x25\x70\x16\xd0\x33\x40\x91\x8f\x80\xc9\xc0\xf5\xaa\xba\x32\x49\xdb\xa2\xd2\xad\x84\x14\x91\x1d\x81\x9f\x00\xe7\x02\xcd\x11\xaa\x78\x13\x38\x4d\x55\x5f\xf6\x6a\x98\x07\xba\x8d\x90\x22\x32\x0e\xb8\x09\xe8\x1d\xb3\xaa\x0f\x80\x13\x54\x75\x76\x6c\xa3\x3c\x52\xf7\x42\x8a\x48\x6f\xe0\x2e\xe0\x14\x8f\xd5\x76\x02\x47\xa8\xea\x4b\x1e\xeb\x8c\x45\x94\xe1\xa5\xd6\x78\x0c\xbf\x22\x82\xf5\xea\xe9\x22\xb2\xbd\xe7\x7a\x23\x53\xd7\x42\x8a\xc8\x40\xe0\xd0\x84\xaa\xdf\x1d\xf8\x65\x42\x75\x87\xa6\xae\x85\x04\x4e\x4f\xb8\xfe\xaf\x8a\xc8\xd8\x84\xdb\x08\x44\xbd\x0b\x79\x5a\x0a\x6d\xfc\x5c\x44\x5a\x52\x68\xa7\x22\x75\x2b\xa4\x7b\x7e\x1d\x90\x42\x53\x43\x80\xaf\xa7\xd0\x4e\x45\xea\x56\x48\x60\x34\xe9\xcd\xca\xaf\x12\x91\x4c\x7f\x97\x89\x0e\x09\x22\xb2\x25\x70\x08\xf6\xad\xed\x01\xfc\x0f\x98\x07\xcc\x57\xd5\xf5\x49\xb6\x0d\x1c\x9e\x70\xfd\xc5\x0c\x05\x8e\xc6\x96\xf3\x32\x21\x11\x21\x45\xa4\x0f\xf0\x7d\xe0\x42\x4a\x2f\x81\x2d\x11\x91\x5f\x01\x53\x54\x75\x55\x12\x36\x00\x87\x25\x54\x6f\x39\xc6\x91\xa1\x90\xde\x87\x1e\x11\x39\x08\x98\x0e\xf4\x0f\x70\xfb\xdb\xc0\x85\xaa\x3a\xd3\xb3\x0d\x5b\x61\x2b\x30\x5b\xf8\xac\xb7\x0a\x9f\x00\xfd\x55\xb5\x3d\xc5\x36\x37\xe0\x75\x5c\x17\x91\xc3\x80\xd9\x04\x13\x11\x60\x17\xe0\x31\x11\x99\x28\x22\x3e\x7f\xe9\xc3\x49\x57\x44\xb0\xd1\xed\xe4\x94\xdb\xdc\x80\x37\x21\x45\x64\x00\xf0\x30\x1b\xdd\x40\x61\x18\x0f\xdc\x23\x22\x9f\xf1\x64\xce\x81\x9e\xea\x09\x4b\x66\x42\xfa\x7c\x46\xde\x05\x6c\x17\xa3\xfc\x19\x40\xb3\x88\x9c\xa9\xaa\x9f\xc6\xb4\x65\x44\xcc\xf2\x51\x19\x23\x22\x2b\xd8\x38\x1a\xac\x06\x56\x00\xcb\x80\x85\xc0\x7c\xe0\x19\x55\x5d\xe4\xbb\x61\x2f\xcf\x48\x11\x39\x01\x78\xc4\x47\x5d\xc0\xcd\xaa\x7a\x49\x4c\x7b\x5e\x06\x86\x79\xb2\x27\x09\xde\x04\x66\x00\xbf\xf6\x25\xaa\x2f\x21\x5f\x01\xf6\xf5\x51\x97\xe3\x52\x55\x9d\x12\xd1\x96\x66\x60\x15\xb0\xa5\x47\x7b\x92\xe4\x09\xe0\x9a\xb8\x3e\xce\xd8\x42\x8a\xc8\x08\xe0\x1f\x71\xeb\xe9\xc2\x3a\x60\xac\xaa\x3e\x11\xc1\x9e\x3d\xb0\x6f\x7c\x2d\xb1\x1e\xb8\x15\x18\x1f\x35\x02\xc1\xc7\x64\x27\x89\xe5\xa9\x66\xe0\x7e\x11\xd9\x3b\x42\x59\xf1\x6d\x4c\x0a\x34\x61\xef\xdc\xcf\xb9\x49\x63\x68\x7c\x08\x79\x94\x87\x3a\x4a\xd1\x1b\x98\x21\x22\xdb\x86\x2c\x37\x24\x09\x63\x52\x62\x6f\x60\x8e\x88\xf4\x0b\x5b\x30\x96\x90\x22\xf2\x39\x60\x9f\x38\x75\x54\x61\x10\x30\x2d\xe4\x3a\xe6\xa0\xa4\x8c\x49\x89\x5d\x80\xdf\x87\x7d\xaf\x8e\xdb\x23\xf7\x8a\x59\x3e\x08\x63\x81\x09\x21\xee\x1f\x90\x8c\x19\xa9\x32\x0a\x0b\x10\x0b\x4c\x5c\x21\x3f\x1f\xb3\x7c\x50\x7e\x20\x22\xc7\x06\xbc\x77\x8f\x44\x2d\x49\x8f\x4b\xc3\xdc\x1c\x57\xc8\x34\x63\x56\xa6\x55\x9b\x08\xb8\x21\x78\xb7\x74\xcc\x49\x9c\xe1\x61\x62\x82\xe2\x0a\xd9\x2b\x66\xf9\x30\x6c\x07\x3c\x50\x65\x19\x6f\x67\xcc\x5d\x56\x2f\x04\xfe\x52\xc6\x15\x32\xed\x85\xe9\x03\x81\x1b\x2a\x7c\x1e\x7a\xb6\x97\x73\x3e\x0a\x7a\x63\x5c\x21\xb3\xd8\x0f\x71\x85\x88\x7c\xa5\xcc\x67\x7d\x53\xb5\x24\x59\x56\x13\x62\x61\x23\xae\x90\x99\xf8\xde\x80\xa9\x22\x52\x4a\xb4\xa0\xee\xb3\x5a\x60\xa6\xaa\xae\x0d\x7a\x73\x64\x21\x45\x64\x1f\x2c\xbc\x21\x0b\xfa\x02\xb7\x97\xb9\x5e\x2f\xdc\x14\xe6\xe6\x50\x6e\x2c\x17\x7e\x7f\x0e\x70\x01\xd9\x7b\x17\x4e\x10\x91\xf3\x54\xf5\xae\xa2\x6b\x71\xdc\x68\x79\xe2\x36\x55\x7d\x2e\x4c\x81\x40\x3d\x52\x44\x76\x13\x91\x29\x98\x5f\x6d\x32\xd9\x8b\x58\xe0\x17\x5d\x5e\x49\xb6\xc9\xca\x10\x8f\xcc\x07\x2e\x0b\x5b\xa8\xa2\x90\x22\xd2\x4f\x44\x6e\x07\x16\x03\x17\x93\xbf\x5f\x54\x2f\xec\x79\x59\xf8\x7f\x04\xd9\xeb\x98\x67\xde\x00\x8e\x89\x12\x90\x56\x72\x68\x75\x91\xd3\x57\x00\xd7\x02\x5b\xc7\xb3\x2d\x71\x46\x63\x5f\xb2\xc9\xc4\xdf\x32\x97\x25\xcf\x03\xc7\xa9\xea\xf2\x28\x85\x37\xf3\x47\xba\x8d\x2f\xf7\x91\x5d\xb8\x44\x14\x56\x62\x8b\xf7\x93\x80\x13\x33\xb6\x25\x0a\xb7\x61\xce\xf4\x8f\xa3\x56\xb0\x49\x8f\x14\x91\x23\xb0\x50\xc6\xb0\xae\xa3\xac\xe9\x09\xdc\x02\xbc\x97\xb5\x21\x21\x69\x03\xbe\xa9\xaa\x8f\xc6\xad\x68\x43\x8f\x14\x91\xa3\x81\x47\x01\x5f\x91\x6c\x59\xf0\x5f\x60\xc7\xac\x8d\x08\xc8\x1d\xc0\x95\xaa\xfa\xbe\x8f\xca\x9a\x00\x44\x64\x28\x16\xae\x51\xeb\x93\x85\x5a\x60\x31\x30\x4e\x55\xff\xea\xb3\xd2\x66\x11\x69\x02\x7e\x43\x43\xc4\xa4\x59\x83\xad\x13\x8b\x6f\x11\xc1\x9e\x91\x27\x92\x5d\x40\x6f\x77\x61\x16\x70\xb1\xaa\xfe\x33\xa9\x06\x5a\x08\xe9\x89\x6e\x10\x8a\x77\xb0\xe7\xe0\x7d\x49\x37\xd4\x02\x1c\x9c\x74\x23\xdd\x90\xd5\xc0\x44\x60\xa2\xaa\x7e\x98\x46\x83\x2d\x40\x6b\x1a\x0d\x75\x23\xee\x01\xbe\xab\xaa\x6f\xa7\xd9\x68\x0b\xb0\x96\xfa\xf2\xaa\x67\xc5\x2c\x4c\xc0\x17\xb3\x68\xbc\x05\x5b\xdf\x8b\x12\x08\xdc\xc0\x98\x87\x09\xf8\xe7\x2c\x8d\x68\x01\x9e\xa4\x21\x64\x1c\x06\x02\x0f\x8a\xc8\x3a\x2c\x23\x56\x31\x9d\xd8\x06\xd8\x4e\xf7\xb3\x0a\x5b\x4e\x5c\x89\x6d\xc3\xff\xc0\xfd\xb9\x1c\x8b\xb6\xe8\xc0\x76\x6f\xad\x08\xbb\x35\xbf\xc9\x85\xe5\xbf\x1a\xfd\xff\xd1\x20\x01\xd6\x63\xa2\x2e\x07\xde\x72\x3f\x4b\xdd\xcf\x62\x55\x9d\xd3\xb5\x40\x61\x65\x67\x06\xb5\xb9\xd8\xdc\x5d\x99\x06\x7c\x43\x55\xd7\x14\x2e\x14\xfc\x78\x97\x13\x22\x62\xab\x41\xe6\x9c\x8d\x6d\x2b\xd8\xb0\x2e\xde\x0c\xa0\xaa\x8b\x09\x19\xd9\xdc\x20\x73\x8e\xc7\x3c\x3e\x40\x51\x5c\x6a\x7b\x7b\xfb\x4b\xad\xad\xad\xfd\x80\x2f\x64\x61\x55\x83\x48\xec\xdf\xda\xda\xba\xb0\xbd\xbd\x7d\xe1\x26\x8e\x65\xb7\x03\xe8\x7e\xe0\xd4\x6c\xec\x6a\x10\x81\x36\x60\xf7\x4d\x22\xc5\xdb\xdb\xdb\xd7\xb7\xb6\xb6\x3e\x84\xed\x68\xca\x4b\x80\x55\x83\xca\xf4\x04\x96\x6c\x16\xf2\xef\xc4\x7c\x04\xf8\x14\xcb\x75\x5a\xcf\xf9\xea\xea\x85\xe6\x8a\x39\x04\x44\xa4\x27\x16\xdc\x74\xa4\xfb\x73\x7f\x12\xce\x5f\xd7\x20\x12\xef\x86\x4a\x06\xe1\x52\x83\x8d\xc0\x3c\x26\x23\x30\x61\x07\x26\x60\x58\x83\x90\xf8\xc8\xea\xd1\x07\xcb\x8b\xba\x1f\x96\x88\x61\x5f\x2c\xa2\xad\x11\x71\x90\x1e\x1f\x25\x96\xcf\x54\x44\xfa\x03\x83\xdd\xcf\x10\x6c\x02\x55\xf8\xa9\x97\xd0\xfe\xbc\x30\x2f\x93\xe3\x22\x44\xa4\x17\xb6\x97\xb1\x15\xdb\x9c\xda\x17\x13\xb7\x27\xd0\x07\x4b\xf9\x92\xb7\xa8\xf6\x3c\xf3\xb3\x5c\x9d\xfb\xe1\xa2\xf9\xee\xc6\xce\xa9\x6a\x10\x8c\xf5\xc0\xe0\x5c\xcc\x40\xdd\x42\xc4\xe5\xc0\x8f\x80\xcf\x66\x6c\x4e\xad\x71\x87\xaa\xfe\x2b\x73\x21\xdd\x6e\xaa\xdf\xd1\x88\x1d\x8a\xc2\x32\xe0\x6a\xc8\xf8\x65\x5f\x44\xbe\x86\x6d\x23\x6b\x88\x18\x9e\x35\xc0\x19\xaa\xba\x02\x32\x7a\xb9\x77\x93\x9d\x5b\xb0\x63\xff\x1a\x84\x67\x1d\x70\xae\xaa\x3e\x53\xb8\x90\xfa\x64\xc7\xe5\x3c\xbf\x97\xfa\x49\x6c\x94\x36\x6b\x80\xb3\x55\xf5\xc1\xe2\x8b\xa9\x0e\xad\x22\x72\x11\x30\x87\x86\x88\x51\x59\x0e\x7c\xa9\xab\x88\x90\x52\x8f\x14\x91\x6d\xb0\xa1\x34\xf3\x13\x6b\x6a\x98\xbf\x63\xcf\xc4\x92\xf1\xb2\x89\xf7\x48\x11\x19\x04\x3c\x4b\x43\xc4\xa8\x7c\x82\x9d\xa1\x72\x68\xa5\xa0\xe7\x44\x7b\xa4\x88\x9c\x0c\x4c\x25\xdd\x54\x67\xf5\xc4\x5c\x6c\x23\xec\xc2\x6a\x37\x26\x22\xa4\xdb\xaa\x77\x2d\xe1\xd2\x73\x36\xd8\xc8\x52\xe0\x1a\xe0\x9e\xa0\xf1\xad\x49\x9c\xc4\xb3\x35\xd6\x0b\xd3\x38\xf2\xaf\xde\x58\x01\xdc\x08\x4c\x52\xd5\xd5\x61\x0a\x7a\x15\x52\x44\x76\x06\x66\x62\x7e\xca\x06\xe1\x98\x04\x7c\x2f\xcb\xe4\xf4\x00\x88\xc8\x7e\xc0\x73\x34\x44\x8c\xca\xd2\xa8\x22\x82\x27\x21\x45\xe4\xcb\xd8\x83\x79\x57\x1f\xf5\x75\x53\x46\xc7\x29\x1c\x5b\x48\x77\x0a\xcf\x63\x34\x66\xa6\x71\x89\x15\x4f\x1c\xeb\x19\xe9\xf2\x8c\xcf\xa0\x3e\xf7\x57\xae\x23\x7d\xa7\xc2\x96\x51\x93\x26\xc5\x49\xf3\xb9\x17\x96\x21\xab\x1e\x44\x7c\xb7\xc4\xb5\x95\x6c\xbe\x4d\x2e\x69\x22\x27\xfb\x8f\x24\xa4\xdb\x3c\x72\x2f\xf5\x33\x9c\xfe\xb8\xc4\xb5\xde\x98\x8b\x2d\x4d\x22\x9f\xe7\x15\xb5\x47\x5e\x42\x7d\x45\xa2\xff\x16\x58\x50\xe2\xfa\x20\xe0\x81\x94\x6d\x89\x44\x68\x21\xdd\x01\xd8\x57\x27\x60\x4b\x56\x74\xb8\x8c\x8c\x53\x4b\x7c\xb6\x13\x76\xee\x63\x62\xf9\x71\xba\x10\xf9\xdc\xcc\x28\x3d\xf2\x58\xea\x2b\xe5\x74\xa1\x27\x6e\xe6\x1a\x72\x5c\x84\x65\x8d\x8e\x9c\xb9\x31\x04\x91\x73\xc4\x47\x11\x32\xd6\xfb\x4e\x0e\x79\x01\x40\x55\xdf\xc2\x16\x34\xba\xb2\x23\x96\x63\xe1\x5b\x09\xdb\xf1\x31\x29\x0b\x59\x2f\x27\xdd\x14\x28\x3e\xfb\xf2\xa1\x32\xf7\x7c\x47\x55\xef\xa4\x74\x42\x7c\x5f\xbc\x18\x36\x01\x44\x31\x51\x84\x5c\x17\xb5\xb1\x9c\x52\x2c\xe4\x93\x65\xee\x19\xe6\x0e\x3c\xbd\x04\x78\xa6\xcc\x3d\x71\xf9\x5b\x9c\xc2\x51\x84\x7c\x3d\x4e\x83\x39\x63\xb9\xaa\x2e\x29\xfa\xf7\x02\xa0\x5c\xfe\xd4\x71\xee\x65\xfd\x24\x92\x39\x31\xb6\xdc\x68\x10\x88\x28\x42\x3e\x1c\xa7\xc1\x9c\xb1\x49\xef\x72\x43\xdb\x0b\x65\xee\x3d\x5d\x44\x7a\xa8\x6a\x07\x76\x14\xa2\xcf\x53\x88\x16\xa8\x6a\xb9\x76\x03\x11\x5a\x48\x97\xa2\xab\xdc\x10\x54\x6b\x3c\x5d\xe2\x5a\xb9\x9c\x43\xdb\x01\xc7\x00\xb8\x74\x9d\x47\xe2\x4f\xcc\xeb\xe2\x56\x10\x75\x41\xe0\x02\x2c\xa2\xab\xd6\x99\x5d\xe2\x5a\xa5\x61\xf3\x94\xc2\x5f\x54\x75\x01\x70\x08\x10\xf7\xf8\xf9\xa7\x54\x35\xf6\x28\x17\x49\x48\x17\x04\x74\x14\x96\x43\xbc\x56\xe9\xa4\xf4\x6a\xce\x7f\x2a\x94\x19\x5b\x7c\x4c\xb0\xeb\x99\x5f\x04\x1e\x89\x68\xc3\x32\x3c\x05\xa5\x45\x5e\x34\x77\xdf\xc8\x11\xc0\x9f\x7c\x18\x92\x01\x73\x54\xb5\xd4\x4a\x4a\xa9\x05\xf4\x02\x7d\xb1\x4d\xbd\x1b\x50\xd5\xf7\x54\xf5\x24\xe0\x3c\xc2\x0d\xb5\x6f\x61\x31\xaa\xef\x84\x28\x53\x96\x58\x6e\x1a\x55\x7d\x5b\x55\x8f\xc6\xd2\x9f\xcd\xf3\x61\x50\x8a\x94\x7a\x3e\x42\xf5\xd7\xab\xc3\x4b\x5d\x54\xd5\xbb\xb1\xb5\xd9\xeb\xa8\xfc\x62\xff\x29\xb6\x75\xf0\x00\x55\x7d\xad\x4a\x5b\x81\xf1\x1d\xb3\x73\x10\xb6\x9f\x63\x0c\xe9\x1c\xa4\x1d\x87\x51\xaa\xfa\x6c\xd7\x8b\x22\x72\x38\xf0\x54\x85\x72\x33\x55\xf5\xb8\x4a\x15\x8b\x48\x0f\xe0\x30\x6c\x15\x6c\x30\xe6\xd5\x78\x1f\xfb\xb2\x4f\x57\xd5\xa5\x51\x8d\x2e\x47\x92\x5b\xcf\x5b\xb1\xa5\xad\x21\xd8\xae\xe4\x5e\xd8\x99\x22\x6b\xb1\xe7\x53\x07\x36\x51\x58\x84\xf9\x35\xd3\xdc\xdc\xba\x0a\xd8\x56\x55\x3f\xe9\xfa\x81\x88\x8c\xa1\xf2\xe3\xa2\x13\xd8\xbe\xcc\xb0\x9c\x19\x89\xed\xc6\x52\xd5\x76\x6c\x88\x99\x5d\xed\x5e\x11\xf9\x21\xf0\xc7\xa4\x6c\x29\xc1\xdc\x52\x22\x3a\x76\xa8\x52\xb6\x37\x96\xf4\x22\x6d\x5f\x65\x45\xf2\x92\x0c\xe9\x09\x20\xcd\x14\xd2\x95\x86\xce\x20\x39\xde\x73\x17\x29\x98\x0b\x21\xdd\x8a\xca\x15\x29\x36\x39\xab\xc2\x67\x83\x02\x94\x1f\xee\xcb\x10\x5f\xe4\x42\x48\x00\x55\x7d\x1a\x0b\xe4\x4a\x9a\x4e\x2a\xcf\xb0\x07\x07\xa8\xa3\xd1\x23\xab\xf0\x6d\x6c\x22\x92\x24\xe5\xde\x1f\x0b\x04\x09\x61\xd9\xd3\x97\x31\xbe\xc8\x95\x90\xce\xb9\x9b\x74\x18\xc9\x5f\xca\x7d\xe0\x12\x53\x04\x39\xed\xae\x9f\xdb\xe3\x92\x1b\x72\x25\xa4\xe3\x66\x92\x9b\xc1\xb6\x61\xaf\x3a\xe5\x18\x15\xa2\xae\x01\xf1\x4c\xf1\x4b\xee\x84\x74\x13\x9f\x73\xf0\xef\xf3\xeb\xc4\xce\x2f\x6e\xab\x70\xcf\x98\x10\xf5\xed\x14\xd3\x1e\xaf\xe4\x4e\x48\x00\x97\x72\xe4\x38\xfc\xb9\x89\x3a\x80\x31\xaa\x5a\xf6\xdd\xcf\xed\xe9\x3c\x26\x44\x9d\xb9\xca\xa7\x97\x4b\x21\x01\x54\xf5\x75\xcc\xe7\x17\xd7\x5d\xb6\x08\x38\x58\x55\x9f\xaf\x72\xdf\x30\x6c\x05\xaa\x1a\x6b\x80\xab\xc8\x99\x83\x3d\xb7\x42\xc2\x06\x0f\xcb\x48\x40\x23\x56\x31\x15\x38\x50\x55\xdf\x08\x70\xef\x99\x01\xee\x79\x0d\x18\xa9\xaa\x37\xaa\x6a\xae\x62\x97\x72\x95\x54\xb0\x1c\x2e\x28\x7a\x02\x70\x19\xc1\xce\x80\x9e\x8f\x9d\xdb\x58\x76\x86\xda\xa5\xfe\x26\x60\x09\x95\x23\x04\xa7\x00\xe3\x55\x35\x97\xe7\xa3\xd4\x84\x90\x05\x44\x64\x57\x2c\x60\xf8\x54\x36\x7f\x71\xef\xc0\x42\x50\xee\xc4\xbc\xee\x81\x43\x0b\x45\xe4\x50\xca\x47\xb1\xb5\x01\xe7\xab\xea\xe3\xe1\x2d\x4e\x8f\x9a\x12\xb2\x18\x11\xd9\x16\xdb\x58\xbb\x05\xd0\x56\x65\x36\x5a\xad\xae\xbb\x29\x7d\xb2\xed\x1f\xb0\xa3\x8b\x72\x1f\xd6\x52\xb3\x42\xfa\x42\x44\x76\xc0\x8e\xd0\x2d\x4e\x2f\xfa\x21\x70\xb9\xaa\xde\x9a\x8d\x55\xe1\xc9\x3c\xcd\x67\x0e\x38\x9f\x4d\x45\x7c\x01\x38\x4b\x55\xe3\x06\x55\xa5\x4a\xb7\xee\x91\x22\x52\x38\x08\x75\x00\x16\xe2\xf1\x53\x60\x82\xaa\xae\xcd\xd2\xae\x28\x74\xf7\x1e\x79\x26\x26\xe2\xbf\xb1\x8c\x8b\x73\xb3\x35\x27\x3a\xdd\xb6\x47\xba\x57\x8e\x57\x31\x87\xf6\xc5\xaa\x9a\xf6\x36\x73\xaf\x74\xe7\x1e\x39\x1a\xb8\x5e\x55\x2b\x2d\xa2\xd7\x0c\xff\x07\xd4\x06\x0a\x49\x06\xf0\xc8\x6d\x00\x00\x00\x00\x49\x45\x4e\x44\xae\x42\x60\x82\x01\x00\x00\xff\xff\x25\x09\x80\xa2\xcd\x0b\x00\x00")

func sourcedataShapesShape1ShapePngBytes() ([]byte, error) {
	return bindataRead(
		_sourcedataShapesShape1ShapePng,
		"sourcedata/shapes/shape-1/shape.png",
	)
}

func sourcedataShapesShape1ShapePng() (*asset, error) {
	bytes, err := sourcedataShapesShape1ShapePngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sourcedata/shapes/shape-1/shape.png", size: 3021, mode: os.FileMode(420), modTime: time.Unix(1715607131, 0)}
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
	"sourcedata/shapes/shape-1/shape.png": sourcedataShapesShape1ShapePng,
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
			"shape-1": &bintree{nil, map[string]*bintree{
				"shape.png": &bintree{sourcedataShapesShape1ShapePng, map[string]*bintree{}},
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
