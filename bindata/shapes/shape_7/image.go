// Code generated for package shape_7 by go-bindata DO NOT EDIT. (@generated)
// sources:
// sourcedata/shapes/shape-7/shape.png
package shape_7

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

var _sourcedataShapesShape7ShapePng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x0c\x0a\xf3\xf5\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52\x00\x00\x00\x78\x00\x00\x00\x78\x08\x06\x00\x00\x00\x39\x64\x36\xd2\x00\x00\x09\xd3\x49\x44\x41\x54\x78\x9c\xed\x9d\x7f\xb0\x55\x55\x15\xc7\x3f\x1a\x08\xfe\x42\x22\xbc\x31\x59\x86\xa4\x98\xb6\xd0\x34\xcc\x54\x1c\x11\x25\x85\xd4\xe2\x87\x96\x92\xd5\x34\x8e\xfd\xf0\x07\x38\xc8\x0c\xe3\x34\xe5\xd4\xa4\xd6\x1f\xfd\x30\xc9\x9a\x2c\xc1\x42\x69\xc4\x5f\xe3\x98\x64\xc5\x18\xc1\x30\x58\x43\x4d\x2d\x49\x48\xb0\x0c\x50\x2e\x42\xbe\x27\xf9\x78\x3e\x88\xfe\xd8\xe7\xbd\xb9\xdc\x77\xdf\x7b\xe7\x9e\xbb\xf6\xd9\xf7\xdc\xb7\x3f\x33\x77\x86\x77\xef\x39\x6b\x2f\xee\xf7\xee\x7d\xce\xd9\x7b\xad\xb5\x21\x12\x89\x14\x97\x43\x42\x3b\x90\x17\x22\x32\x12\x18\x85\xfb\x3f\xb7\xa9\xea\x6b\x81\x5d\xca\x85\x96\x15\x58\x44\x8e\x02\x66\x01\x33\x80\xb3\x80\x77\x55\x1d\xb2\x0b\x58\x0d\x3c\x0a\x3c\xac\xaa\x1d\xf9\x7a\x98\x0f\x2d\x27\xb0\x88\x0c\x07\x16\x00\xf3\x70\x3d\x36\x0d\x65\xe0\x9b\xc0\x0f\x55\x75\x9f\x2f\xdf\x42\xd0\x52\x02\x8b\xc8\x44\x60\x29\x30\x3e\xa3\x89\x75\xc0\x35\xaa\xba\xc5\xce\xab\xb0\xb4\x8c\xc0\x22\x72\x15\xb0\x04\x18\xde\xa0\xa9\xdd\xc0\x0c\x55\x5d\xd5\xb8\x57\xe1\x69\x09\x81\x45\x64\x0e\xf0\x00\x70\xa8\x91\xc9\x0e\x9c\xc8\xbf\x36\xb2\x17\x8c\xc2\x0b\x2c\x22\x53\x81\xa7\x81\xb7\x19\x9b\xee\x00\xa6\xaa\xea\x1a\x63\xbb\xb9\x52\x68\x81\x45\x64\x1c\xb0\x1e\x38\xc6\x53\x13\x3b\x81\x33\x55\x75\xab\x27\xfb\xde\xb1\x1a\xd2\x72\x47\x44\x86\x02\x0f\xe1\x4f\x5c\x80\x63\x81\x9f\x8b\x88\xf5\xe8\x90\x1b\x85\x15\x18\x98\x0f\x7c\x38\x87\x76\x26\x03\x73\x73\x68\xc7\x0b\x85\x1c\xa2\x45\xe4\x04\xe0\x79\xe0\xf0\x9c\x9a\x7c\x03\x38\x49\x55\x77\xe4\xd4\x9e\x19\x45\xed\xc1\x77\x90\x9f\xb8\x00\x47\x03\x5f\xcf\xb1\x3d\x33\x0a\xd7\x83\x45\xe4\x74\xe0\xcf\xe4\xef\x7b\x17\x70\x82\xaa\x6e\xcb\xb9\xdd\x86\x28\x62\x0f\xbe\x95\x30\x3f\xcc\xa1\xc0\x2d\x01\xda\x6d\x88\x42\xf5\x60\x11\x19\x03\xbc\x8c\xfb\xb2\x43\xf0\x1a\x70\x9c\xaa\xbe\x15\xa8\xfd\xba\x29\x5a\x0f\xbe\x96\x70\xe2\x02\x8c\x06\x2e\x0b\xd8\x7e\xdd\x14\x4d\xe0\xab\x43\x3b\x00\x7c\x32\xb4\x03\xf5\x50\x98\x21\x5a\x44\xc6\x02\x2f\x85\xf6\x03\x68\x07\x8e\x2d\xca\x30\x5d\xa4\x1e\x7c\x49\x68\x07\x12\x46\x00\xe7\x87\x76\x22\x2d\x45\x12\xf8\xa2\xd0\x0e\x54\x70\x61\x68\x07\xd2\x52\x24\x81\xf3\x98\x96\x4c\xcb\xe4\xd0\x0e\xa4\xa5\x10\xd7\x60\x11\x19\x8d\x5b\xd9\x69\x16\xf6\x02\x47\x17\x21\xbc\xa7\x28\x3d\xf8\x94\xd0\x0e\x54\x31\x1c\x38\x39\xb4\x13\x69\x28\x8a\xc0\x59\x63\xac\x7c\xd2\x6c\x3f\xba\x9a\x14\x45\xe0\xe3\x43\x3b\x50\x83\x93\x42\x3b\x90\x86\xa2\x08\x3c\x26\xb4\x03\x35\x28\x85\x76\x20\x0d\x45\x11\xb8\x19\xbf\xcc\xd1\xa1\x1d\x48\x43\x51\x04\x3e\x22\xb4\x03\x35\x78\x47\x68\x07\xd2\x50\x14\x81\x87\x85\x76\xa0\x06\xcd\xf8\xa3\xeb\x45\x51\x04\x8e\x64\x24\x0a\xdc\xe2\x0c\xb1\x34\x26\x22\x25\x60\x1c\x2e\xdc\xb4\xf2\x1a\xb5\x07\xd8\x0a\x6c\x03\x5e\x29\xc2\x0c\x50\x0a\x3a\xb3\x9c\x24\x22\xef\x04\x4e\xc4\x7d\x47\x23\x80\x23\x81\xb7\x70\x81\x7d\xed\xc0\x66\x60\x8b\xaa\xee\xb7\x70\x32\xb3\xc0\x22\x72\x08\x70\x06\x30\x0d\x98\x84\x9b\x2b\x4e\x93\xcd\xb7\x5f\x44\x9e\x07\x9e\x03\xfe\x08\x3c\xab\xaa\x9b\xb2\xfa\x11\x90\xf6\x81\x0e\x48\x7e\xf0\xe7\x03\xe7\x26\x2f\x01\x8e\x4a\x61\xbb\x53\x44\x36\x00\xab\x70\x59\x1b\x2b\x55\xb5\x2b\x8b\x93\x75\xcf\x45\x8b\xc8\xf1\xc0\xf5\xc0\x1c\x60\x6c\x96\x46\x6b\xf0\x0f\xe0\x49\xe0\x31\x60\x8d\xaa\x1e\xa8\x6a\xf3\x59\xe0\x02\xa3\xb6\xac\x58\xa4\xaa\x37\x56\xbf\x29\x22\x13\x70\x79\xc9\xd3\x81\x89\xd8\xcc\xf7\xef\x04\x16\x03\xf7\xa8\xea\xcb\xf5\x9c\x98\xba\x71\x11\x39\x15\xf8\x1a\xce\x79\x9f\x91\xfe\x9b\x81\xfb\x81\xc5\xdd\x11\x8c\x4d\x2a\xf0\x6d\xaa\x7a\x27\xf4\x88\x7a\x55\xf2\xf2\x39\xad\xda\x05\xfc\x0c\xb8\x5d\x55\x5f\x4d\x73\xc2\x80\x02\x8b\xc8\x11\xb8\x38\xe4\x1b\xf1\x2b\x6c\x35\xff\xc3\x0d\x4f\xf7\xe1\x46\x8c\x69\x39\xb6\x9d\x86\x6f\xe1\xbe\xf0\x2b\xc9\x7f\xe1\xe1\x0d\x60\x21\x70\x6f\xf5\x68\x57\x4d\xbf\x02\x8b\xc8\x89\xc0\x13\xc0\xa9\x76\xbe\x45\x0c\x59\x01\x7c\x5a\x55\x77\xf5\x75\x40\x9f\x02\x27\xc3\xce\x4a\x0a\x32\x25\x37\x88\xd9\x08\x5c\xdc\x57\x06\x64\x4d\x81\x45\xe4\xbd\xb8\xbb\xdc\x66\x9c\x03\x8e\xf4\x66\x13\x70\x8e\xaa\xee\xae\xfe\xa0\xd7\x44\x87\x88\x0c\x01\x1e\x26\x8a\x5b\x24\xc6\x03\x4b\x93\x47\xd7\x83\xa8\x35\x93\x75\x33\xae\xec\x50\xa4\x58\x5c\x0a\x7c\xb6\xfa\xcd\x83\x14\x4f\x6a\x4b\xfd\x93\x82\xac\x94\x44\x7a\xf1\x2a\x30\xae\xb2\xe6\x57\x75\x0f\xfe\x14\x51\xdc\x22\x33\x06\x98\x5d\xf9\x46\x2d\x81\x23\xc5\xe6\x9a\xca\x3f\x7a\x86\xe8\xa4\xe6\x45\x3b\x8d\xd7\x99\x8a\x84\xa5\x03\x18\xd1\xbd\xa0\x53\xd9\x83\xdf\x47\x14\xb7\x15\x38\x1c\xb7\xa2\x07\x1c\x2c\x70\x33\x06\xb6\x45\xb2\xf1\x9e\xee\x7f\x54\x0a\x3c\x22\x80\x23\x11\x3f\xf4\x84\x13\x55\x0a\xdc\x16\xc0\x91\x88\x1f\x7a\xb4\xac\x14\xf8\x95\x00\x8e\x44\xfc\xd0\xa3\x65\xa5\xc0\x5b\xc8\x18\x86\x12\x69\x2a\xde\xc4\x69\x09\x54\x08\x9c\xdc\x56\xaf\x0d\xe1\x51\xc4\x94\x35\x95\xf1\x5c\xd5\x13\x1d\xbf\xca\xd9\x99\x88\x3d\x4f\x54\xfe\x51\x2d\xf0\x52\x5c\x24\x45\xa4\x98\x74\xe2\x0a\xb4\xf6\x70\x90\xc0\xaa\xba\x1d\xb7\x54\x18\x29\x26\x8b\xab\xd7\x84\x6b\x2d\x17\xde\x01\xf4\x1b\xe7\x13\x69\x4a\xf6\xe2\xb4\x3b\x88\x5e\x02\xab\xea\x5f\x71\x91\x7b\x91\x62\x71\x67\xad\x90\xda\xbe\x52\x57\x16\xe2\xca\xf6\x45\x8a\xc1\x46\xe0\xdb\xb5\x3e\xa8\x29\x70\xb2\x2b\xd8\x97\x7c\x7a\x14\x31\x63\x1f\xf0\x19\x55\xdd\x5b\xeb\xc3\x3e\xe3\x9c\xcb\xe5\xf2\x86\x52\xa9\x34\x06\x17\x9d\x1f\x69\x5e\xe6\xab\xea\x23\x7d\x7d\x38\x50\x76\xe1\x5c\xe0\x0f\xb6\xfe\x44\x0c\x59\xa2\xaa\xdf\xeb\xef\x80\x7e\x05\x4e\xea\x31\x5e\x01\xa8\xa5\x57\x11\x13\x56\xe0\x32\x3e\xfa\x65\xc0\xfc\x60\x55\x7d\x1d\x98\x42\x14\xb9\x99\xf8\x1d\x30\x2b\x4d\x41\xd4\x54\x09\xe0\xaa\xba\x13\x97\x22\xfa\xfb\x06\x1d\x8b\x34\xce\x72\x60\xba\xaa\xbe\x99\xe6\xe0\xd4\xc9\x64\xe5\x72\xb9\xb3\x54\x2a\x3d\x88\xcb\x01\x6e\xa6\xba\x91\x83\x85\x03\xc0\x37\x80\x1b\xea\x49\x0e\xcf\x94\xbb\x2a\x22\x97\xe3\x26\x43\x62\xde\x52\x3e\xec\xc0\x3d\x0a\x3d\x53\xef\x89\x99\x6a\x74\xa8\xea\x93\xb8\x6c\xf5\xa5\x59\xce\x8f\xd4\xc5\x03\xc0\x07\xb2\x88\x0b\x06\xd9\xe7\x22\x72\x2e\xf0\x7d\xe2\xf3\xb2\x35\x6b\x81\x05\x8d\x6e\x8e\x69\x56\x4e\x58\x44\x2e\x05\x6e\x07\xce\xb6\xb2\x39\x48\x59\x87\xcb\xe0\x5f\x61\x61\xcc\xbc\x5e\xb4\x88\x9c\x07\x7c\x11\x57\xce\xe0\x30\x6b\xfb\x2d\x4a\x27\xb0\x0c\xf8\xb1\xaa\x9a\x46\xd5\x78\x2b\x08\x9e\x14\xf1\x9e\x0d\xcc\xc4\x95\xc0\x37\x2d\xd9\xd4\x02\xec\xc3\x3d\xcf\x2e\x07\x1e\xad\x95\xdb\x6b\x41\x2e\x15\xdf\x45\x64\x14\x6e\x46\xec\x63\xc0\x54\xfc\x6e\x09\xdb\xcc\x6c\x07\x9e\x49\x5e\x4f\x27\x93\x48\x5e\xc9\xbd\xa4\x7f\x92\x60\x3e\x05\x57\x32\xa9\x10\xf5\x1e\x8d\x98\xa9\xaa\x8f\xe5\xdd\x68\x88\x52\x86\x07\x80\x1b\x18\x5c\xe2\x02\x7c\x55\x44\x72\x2f\xaa\x1a\x42\xe0\x7b\x70\xc3\xf5\x60\xe3\x83\xb8\xd2\x4b\xb9\x92\xeb\x10\x2d\x22\x37\x01\x77\xe7\xd9\x66\x13\x72\x99\xaa\x3e\x95\x57\x63\xb9\x09\x2c\x22\x67\xe3\xd6\x96\x43\x6e\x2e\xd9\x0c\x94\x01\x49\x16\x70\xbc\x93\xcb\x10\x9d\xd4\xfe\x58\x46\x14\x17\x5c\xf5\xa2\x9f\xe6\xd5\x58\x5e\xd7\xe0\xef\x60\x57\xb8\xb4\x15\xb8\x5c\x44\xae\xcb\xa3\x21\xef\x43\xb4\x88\x4c\x22\x86\xfd\xd4\xa2\x1d\x78\xbf\xaa\x7a\xcd\xea\xf4\xda\x83\x45\xe4\x50\xe0\x5e\x9f\x6d\x14\x98\x11\xb8\x45\x1a\xaf\xf8\x1e\xa2\x3f\x87\x5b\x56\x8c\xd4\xe6\x4a\x11\xb9\xd8\x67\x03\x3e\xe7\xa2\x0f\x03\x5e\xa4\xa2\x5e\x44\xa4\x26\xeb\x81\x89\x03\x95\x05\xce\x8a\xcf\x1e\x3c\x87\x28\x6e\x1a\xce\xc4\xd5\x9c\xf6\x82\x17\x81\x93\xa2\x98\xb7\xfa\xb0\xdd\xa2\x7c\xc5\x97\x61\x5f\x3d\xf8\x42\x62\x11\xf1\x7a\x98\xe0\xeb\x5a\xec\x4b\xe0\x5c\x9e\xf1\x5a\x8c\x79\x3e\x8c\xfa\x88\xe8\x38\x06\x57\xf5\x34\x56\xcd\xab\x8f\x7d\xc0\x71\xaa\x5a\xb6\x34\xea\xa3\x07\x5f\x41\x14\x37\x0b\x43\xf0\x70\xb3\xe5\x43\xe0\x19\x1e\x6c\x0e\x16\x66\x5a\x1b\x34\x1d\xa2\x93\x68\x8d\xd7\x71\xdb\xb5\x45\xea\xa7\x13\x18\xd9\x57\xae\x6f\x16\xac\x7b\xf0\x44\xa2\xb8\x8d\x30\x0c\xf8\x88\xa5\x41\x6b\x81\xcf\x33\xb6\x37\x18\x31\x4d\x20\xb0\x16\xf8\x34\x63\x7b\x83\x11\xd3\xf9\x03\x6b\x81\x4f\x31\xb6\x37\x18\x31\xdd\xfb\xd0\x5a\xe0\x77\x1b\xdb\xb3\x66\x11\xf0\x78\x68\x27\x06\xc0\xb4\x30\xbb\x75\xb6\x41\x33\x6f\xa6\xb5\x15\xb8\x0d\x17\xae\x3b\x19\x18\x19\xd4\x9b\xbe\x31\x2d\xcc\x6e\xdd\x83\xf3\xdc\x9d\xb4\x5e\xbe\xac\xaa\xed\xc9\xb6\xac\xb7\x84\x76\xa6\x1f\x4c\xb3\x3e\x42\xc4\x45\x87\xe0\xc1\x24\xa7\xb9\x9b\x25\xc0\x6f\x42\x39\x33\x00\x1d\x03\x1f\x92\x1e\x6b\x81\xcd\x1e\xd0\x0d\xd9\x0e\xdc\x54\xf9\x46\xb2\xb8\x7e\x1d\xcd\xb9\x8d\x41\xaa\xda\x1b\x69\xb1\x16\x38\xd5\xae\xd4\x39\x73\x7d\xad\xcc\xbd\xa4\xae\xe3\x17\x02\xf8\x33\x10\x4d\x2d\xf0\xbf\x8c\xed\x35\xca\x0f\xfa\xcb\x22\x50\xd5\x5f\xd2\x7c\x41\x81\xa6\xdf\xa1\xb5\xc0\xeb\x8d\xed\x35\xc2\x1a\x60\x7e\x8a\xe3\xe6\xd2\x5c\xe5\xa1\x5e\xb0\x34\x66\x2d\xf0\x9f\x8c\xed\x65\x65\x1b\x30\x5b\x55\xbb\x06\x3a\x30\x39\x66\x16\x15\x1b\x59\x04\xc6\x34\xc3\xdf\x5a\xe0\xdf\x12\xbe\x98\x78\x1b\x30\x2d\x79\x1c\x4a\x85\xaa\xee\x02\x2e\xc1\x95\x2b\x0a\xc9\x01\x60\xa5\xa5\x41\x53\x81\x93\x68\x84\x55\x96\x36\xeb\xa4\x03\x97\xbd\xf7\xb7\x7a\x4f\x54\xd5\x17\x71\xd5\x07\xda\xcd\xbd\x4a\xcf\xea\x64\x5b\x05\x33\x7c\x3c\x07\xdf\xe7\xc1\x66\x1a\xfe\x8b\x13\x77\x75\x56\x03\xc9\x0f\xe3\x02\xc2\x15\x43\x5f\x64\x6d\xd0\x87\xc0\xcb\xc8\xff\x7a\xb6\x1b\xb8\x48\x55\x1b\x1e\xde\x54\xf5\x2f\xb8\xba\x9c\x5b\x1b\xf6\xaa\x3e\x5e\x00\xfa\xac\xfb\x9c\x15\x73\x81\x93\x0d\xb6\x16\x58\xdb\xed\x87\xcd\xc0\x24\x55\x5d\x67\x65\x50\x55\x37\x02\x67\xe1\xee\xc4\xf3\x62\x6e\xf7\x9e\xbf\x96\x78\x99\x3b\x2e\x97\xcb\x7f\x2f\x95\x4a\x27\x03\x13\x7c\xd8\xaf\xe0\x71\xdc\xb0\xfc\x6f\x6b\xc3\xe5\x72\x79\x4f\xa9\x54\xfa\x05\xf9\x14\x5f\xbd\x5b\x55\xcd\x87\x67\xf0\x9b\x9b\x74\x24\xee\xf9\xf2\x43\x1e\xcc\xb7\x01\xf3\x54\x75\xb1\x07\xdb\xbd\x10\x91\x29\xb8\xa4\xed\xb1\x1e\xcc\x3f\x05\x7c\xbc\x9e\x0a\xb2\xf5\xe0\x6d\xf5\xa7\x5c\x2e\x77\x95\x4a\xa5\xe5\xb8\x9b\x16\xab\x75\xe2\xfd\xc0\xfd\xc0\x27\x1a\xad\xe1\x58\x0f\xe5\x72\xf9\xa5\x52\xa9\xf4\x13\x5c\xec\xf2\x19\xd8\x85\x05\x2f\x03\xae\xf6\x31\x34\x77\x93\x47\x02\xf8\x30\xe0\x2e\xe0\x66\xb2\x5f\xf3\xf7\xe2\x2a\xdb\xde\x95\x3c\xce\x04\x23\xa9\xe0\x37\x1f\xf8\x3c\xd9\xd7\xbf\xdb\x80\x85\xaa\xfa\x23\x33\xc7\xfa\x20\xcf\x22\x2c\xa7\xe3\x16\xdc\x67\x92\x2e\xd0\x60\x1f\xee\x99\xfa\x11\xe0\x21\x55\xfd\x8f\x47\xf7\xea\x46\x44\x86\x02\xd3\x81\x6b\x71\x85\xdd\xde\x9e\xe2\xb4\x1d\xb8\xa1\xfe\xbb\xc9\xd6\x45\xde\x09\x51\xe9\x6e\x14\xf0\x51\xe0\x1c\x5c\xfc\xd1\x48\xdc\xa5\xa2\x0d\xb7\xb4\xb7\x01\x37\xa7\xbd\x56\x55\xf7\xe4\xed\x5f\x16\x92\x4a\x06\xa7\xe1\xa2\x4a\xc7\xe3\x2e\x49\xc3\x81\x2e\x60\x27\xb0\x09\x58\x0d\x3c\xe7\xeb\x5a\x1b\x89\x44\x22\x91\xc2\xf1\x7f\xea\x4e\x84\x68\x59\x26\x16\x04\x00\x00\x00\x00\x49\x45\x4e\x44\xae\x42\x60\x82\x01\x00\x00\xff\xff\xe9\x79\xa2\xe8\x0c\x0a\x00\x00")

func sourcedataShapesShape7ShapePngBytes() ([]byte, error) {
	return bindataRead(
		_sourcedataShapesShape7ShapePng,
		"sourcedata/shapes/shape-7/shape.png",
	)
}

func sourcedataShapesShape7ShapePng() (*asset, error) {
	bytes, err := sourcedataShapesShape7ShapePngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sourcedata/shapes/shape-7/shape.png", size: 2572, mode: os.FileMode(420), modTime: time.Unix(1715607050, 0)}
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
	"sourcedata/shapes/shape-7/shape.png": sourcedataShapesShape7ShapePng,
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
			"shape-7": &bintree{nil, map[string]*bintree{
				"shape.png": &bintree{sourcedataShapesShape7ShapePng, map[string]*bintree{}},
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
