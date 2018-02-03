// Code generated by go-bindata.
// sources:
// templates/views/ses.html
// templates/views/sns.html
// DO NOT EDIT!

package internal

import (
	"github.com/elazarl/go-bindata-assetfs"
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesViewsSesHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x5b\x6f\xdb\xc8\x15\x7e\xf7\xaf\x38\x60\x1d\x50\x2e\x2c\x51\x96\xe3\x5d\x40\xa1\xd4\x7a\xd3\x2c\x36\x40\xd3\xa6\x91\xd1\x87\x0d\x0c\x63\x44\x1e\x89\xe3\x0e\x67\x98\x99\xa1\x2c\xaf\xc0\xff\x5e\x0c\x6f\xa2\xa4\x11\x2d\x6d\x9d\x22\x0e\x96\x2f\x1a\x72\xce\xf5\x3b\x97\xb9\x68\xb5\x82\x10\x67\x94\x23\x38\x81\xe0\x1a\xb9\x76\x20\xcb\x4e\xfc\x90\x2e\x20\x60\x44\xa9\x91\x93\x90\x39\x76\x35\xd5\x0c\x9d\xf1\x09\x00\x40\x73\x32\xff\x7e\xc7\x70\xa6\xcb\xc9\x9c\x20\xba\x1c\x5f\xc7\xe4\x37\xc1\x61\x42\xe3\x84\x21\xbc\x8b\x09\x65\x30\x41\xb9\xa0\x01\xfa\x5e\x74\x59\x4a\xf2\x42\xba\x18\x9f\x94\x3f\x1b\x5a\x03\x86\x44\xce\xe8\xd2\x19\x57\xb3\xab\x15\xd0\x19\xcc\x35\xf4\x24\xc6\x84\x72\xca\xe7\xd0\xef\xf5\xb7\xcd\x95\xe2\xc1\x62\x67\x20\x58\x37\x0e\xbb\x17\x03\x30\x23\x15\x57\xa3\xa5\xea\x5e\x0c\x9a\xa6\x37\x78\x96\x77\x09\xe1\xc8\x1a\xb3\xbb\x14\x4d\x5c\x36\xa8\xa2\xc1\x78\x82\x3c\x84\xc9\x87\x89\xef\x45\x03\x0b\x45\xca\x2a\x31\x9c\x2c\x80\x93\xc5\x94\xc8\xae\xa4\xf3\x48\x43\xae\xf7\x4e\x0b\xc1\xa6\x62\x69\x91\x9e\xf3\x33\x3a\xf6\x49\xc3\x3f\x46\x12\x85\x5d\x46\xf9\x7f\x9c\xb1\x4f\xab\x89\x19\x81\x19\xe9\x06\x11\x2e\xa4\xe0\xdd\x34\x31\x78\xd2\xb1\xef\x91\xb1\xef\x31\x6a\x31\xcb\x4b\x99\xe5\x6b\x5b\x64\x36\xd9\x77\x3f\x6d\x00\x56\x65\x99\x45\xc7\x4c\xc8\xb8\x36\x5b\xc8\xb8\x1b\x09\x49\x7f\x13\x5c\x13\x06\xf9\x3b\x23\x53\x64\xdd\x3c\xd9\x40\x0a\x86\x05\x99\x03\x31\xea\x48\x84\x23\x27\x11\x4a\x3b\x40\xc3\x91\xa3\x90\x87\x5d\x93\x73\x0e\x90\x40\x53\xc1\x47\x8e\x47\x1e\x94\xa7\x50\x39\xc0\xc5\x82\x30\x1a\x12\x8d\x7b\x90\x6d\x18\x4c\x35\xc6\x85\xf2\xb9\x14\x06\x3d\x2b\x47\x11\x0f\x63\x9d\xa1\x1d\x39\x33\x29\x62\x67\x1d\x1a\xae\xa5\x60\x85\xf5\x50\x26\xe2\x65\x95\x87\x97\xd6\x34\xb4\x3d\x3f\x4b\x11\x83\xaf\x12\xc2\xeb\x54\xc7\x2f\x29\x95\x18\x3a\xe3\x3f\xfb\x9e\x99\x68\xb1\xce\xcb\xd5\xb7\x10\xec\x16\xcb\x0f\x95\x8d\x3f\x1c\x6c\xa3\x4f\x79\x92\x6a\xd0\x8f\x09\x8e\x1c\x8d\x4b\xed\x6c\x04\xb4\x84\xa2\x88\x51\x01\x12\x27\x31\x56\xe3\x84\x91\x00\x23\xc1\x42\x94\x23\x07\x97\xc4\xb4\x8e\x8b\xbf\x86\xc2\x54\x7b\x2f\x30\x14\x0b\xc2\x52\x1c\x39\xab\x15\xf4\x4c\x88\x73\x48\xb2\xcc\x81\x0a\x89\x26\x26\x2d\x58\xec\xe4\xe8\x21\x53\xff\x73\x56\x68\xf1\xec\x39\x71\x23\xbe\xa7\x8c\x30\x00\x15\xf9\x60\x46\x4f\x65\xc3\x39\x94\x1f\x07\x6d\x29\x72\x23\x5e\x4e\x82\xa8\x74\x7a\x8f\x81\x7e\xf6\x2c\x99\x14\x72\xbf\xa7\x54\xa9\xa1\x2a\xf2\xa5\x7e\xdd\x8a\x7e\xe5\xf9\x8b\x49\x01\xe3\xfc\xb3\xc7\xff\x6d\xb1\xe6\xe6\xc8\x7e\xf3\x49\xa0\x90\x99\x88\x59\x22\x0f\xc5\xd4\xc0\x01\xa5\x1f\xcd\xea\xff\x40\x43\x1d\x0d\x2f\xfa\xfd\x57\x75\xdf\xc8\xe1\xcb\x7b\x49\x3e\x3a\x2a\xe6\xb5\x09\x22\x31\x5b\x86\x2a\x97\x12\x46\x28\x77\x8a\xdd\x27\x7e\x29\xbb\x8a\x41\xb2\x9c\x81\x2c\x2b\x2d\xc3\x70\xb5\x02\xb3\xe9\xcb\xb2\xf1\x47\x33\x07\x26\x8b\x7d\xaf\x90\x77\xb4\xe2\x48\xc7\xcc\xaa\x37\x9f\xb0\xab\xfd\xe5\xe6\xc3\xdf\x0f\x53\xe8\x7b\x05\xf7\xb7\x57\x03\x31\x2a\x45\xe6\xcf\x5f\x06\x1f\x0a\xb9\xdf\x7c\x05\x98\xa4\x21\x12\x89\xbd\xfb\x49\xf1\xa0\x46\xce\xeb\x2a\xe3\x6b\xb0\x4c\xd2\xd7\x2f\xb6\xbc\xaf\x7a\x62\x05\x43\x96\xf9\x5e\xa5\xea\x6b\x26\x01\xe3\x77\x4a\x30\x1a\x5a\xcf\x0b\x36\x86\xc3\x12\x66\x1f\xd6\x71\xd8\x15\xb3\x99\x42\xdd\xbd\x7c\x0a\xe8\x69\xaa\xb5\xe0\xe5\x82\x23\x51\xe1\x7a\xc5\x99\x6a\x0e\x53\xcd\xbb\x94\xcf\x44\xbd\xaa\x14\x24\xe3\x4f\xe6\xc7\xf7\x0a\xee\x63\x54\xa8\x74\x1a\xd3\x5d\x1d\x2a\x0d\x02\x54\xca\xc9\x0f\x8c\x4f\xcb\x3d\x3e\x22\xbe\x67\x20\x6d\x3d\xa8\x35\x5e\x5b\x0e\xe5\xad\x67\xeb\x27\x52\xfd\x39\x4f\xd6\xff\x4a\x85\x26\x7f\x1c\xab\x0f\x3b\x56\x1b\x3a\xd3\x1c\xbe\x18\xd0\xea\xb5\x33\x42\x03\xc9\xf0\xb2\xdf\x4f\x96\x6f\xd6\x97\x2c\x36\xee\x3d\x40\xad\x81\x66\x54\xe9\x2e\xe5\x8c\x72\x84\x07\x1a\xce\x51\xdf\x69\xc2\xd8\x63\x6b\xbf\xb7\x80\xb4\x41\x90\x1c\xb0\x5e\x36\x1b\x79\x2c\xb8\x8e\x8a\x1a\xa2\x7c\x0e\xb9\xb7\x90\xa0\x84\xc1\x6b\x88\x44\x2a\xcd\x98\x8a\xf0\xa9\x26\x6f\x95\x1d\x88\xd4\x60\xeb\x63\x3c\xf6\x95\x96\x82\xcf\xeb\xfb\x2f\x5c\xea\x6e\x48\xf8\x1c\x65\xd1\x61\x63\xb2\x1c\xbc\xfe\x45\xa4\x72\x52\xac\xc8\xbe\x57\x30\x8c\x01\x63\x42\x99\xf2\x3d\x23\xe4\x00\x23\x7c\xaf\x05\x01\x7b\x92\xd5\xb3\x5f\x0b\xdc\x0f\x64\x09\x66\x09\x01\x49\x34\xfe\x1f\x90\x34\x18\x7e\x22\x1a\x2d\x38\x7a\x0a\x03\x61\xda\xe5\xd7\x85\x73\x4f\xdd\x3e\x59\xa1\xbb\xfd\xf4\x1b\x68\x9b\x13\x4d\x34\x55\x9a\x06\xea\x8f\xde\x79\x5c\xef\x54\x9a\x68\xb5\xd3\x3b\x7f\x1c\x34\x7b\x67\x8b\xce\xbd\xab\x6b\xbd\x71\x6f\x8c\x4e\x1a\x77\xf2\xf7\xaa\xb8\x8e\x57\x81\xa4\x89\x06\x25\x83\x91\xe3\x85\x44\x45\x53\x41\x64\xe8\x11\xa5\x50\x2b\x6f\x81\x3c\x14\x52\x79\xe5\xb5\xa6\x90\xde\x7d\xe3\xa5\x17\x53\xde\xbb\x57\x7f\x59\x8c\x4c\x55\x5d\x27\x09\xa3\x01\x31\xe7\x83\xde\x34\xa5\xcc\xa8\x34\x2e\x14\x1a\xc6\x87\xaa\xc2\x20\x22\x52\x2b\xa3\xa8\x1c\x1e\xa7\xa6\xd6\x53\xec\x8c\xc8\x9a\xdc\xbb\x27\x0b\x52\x4c\x96\x11\x39\xed\x84\x22\x48\x63\xe4\xfa\xac\x27\x91\x84\x8f\x9d\x59\xca\xf3\x5b\x5d\xe8\x9c\xc1\xaa\x46\xb9\x38\x2a\xf5\x50\x4a\x21\x95\x81\xad\x31\x21\x4d\x63\x81\x53\xca\x43\x5c\x9e\xc3\x69\x4e\x03\xc3\x91\x8d\x9a\xe3\x03\x7c\xfc\x87\xd0\x74\xf6\xd8\x59\x6d\x44\x35\xaf\xad\x21\xb8\xef\x0c\x8f\x7b\xbe\x39\x87\x4b\x3d\x04\x77\xb5\xaa\x84\x67\xd9\x36\xc5\x63\x62\x98\xd1\xc2\x1c\xd1\x10\x87\x30\x23\x4c\xe1\xe6\x84\x49\x38\xca\xe7\x43\x70\xa7\x42\x68\xa5\x25\x49\x2e\xdd\x9a\x22\x3b\x7b\xd3\xf0\x11\x79\xb8\xe9\x73\x95\x4f\x5b\xf8\xc4\xf5\xfe\xff\x08\x97\xdf\xf3\x99\xd8\xeb\x71\x43\xa4\xdd\xe7\x72\x87\xfb\x35\xbc\xde\xf1\xf2\xb4\xe3\xfe\xa9\xbe\xfe\x77\xcf\x7a\xc5\xa6\x7b\x4f\xca\x98\x67\x41\x24\x20\x83\x11\x9c\x76\x74\x44\xd5\xd9\x9b\xcd\x5d\x10\x9d\x41\x07\x59\x2f\x24\x9a\x74\xdc\xbc\xa8\xdc\xb3\x6d\x11\xe6\x41\x56\xa9\x6a\xd8\x97\xdb\xbb\x29\x4f\xa2\x4e\x25\x2f\xdc\x7e\xb3\xe1\xd4\x49\xd3\xa2\x7c\xff\xf2\xd6\x54\x16\x8c\xa0\x2a\x31\xca\xa9\xae\x8b\xa1\x37\x47\xfd\x8e\xa1\x19\xfe\xf4\xf8\x3e\xec\xb8\x39\x8b\x7b\x76\x5e\x91\xdf\x44\x18\x63\xc3\x98\xb5\xc8\x9e\x42\xfd\xcf\xfc\x8e\x60\x3b\xdc\x42\x30\x4d\x93\xa1\xc5\x3f\x2d\xe9\x7c\x8e\x72\x08\xae\x39\xd5\x6f\x85\xd2\x3c\xe6\x8c\x41\xb4\x36\x24\xce\x6a\x9a\x0d\x61\x15\x64\xd0\x59\x85\xd9\xab\x33\x67\x13\x90\xf3\x1d\x9d\x53\xb1\xb4\xe9\x54\x91\x78\x18\x82\x96\x29\x5a\xd4\x21\xd1\xa9\x44\x1b\x5b\xce\x4a\x16\x78\xad\xde\xc7\x64\xbe\x97\xe4\x29\x0d\xb5\x81\x65\x0d\x4c\xc8\x02\x2d\x7e\x57\x8f\x39\x8f\x0f\xa1\x8c\x82\x95\x2a\xdb\xf9\x9a\xb5\x01\xa3\x50\x52\x54\x43\xf8\x6c\x09\x46\x51\x57\x09\xb5\x19\x24\x49\x48\x53\xc3\xe7\x5e\x5e\xbd\x72\xcf\xc1\xbd\xba\x7a\xe5\xde\xee\xd2\xe5\xb7\x19\xfb\xc0\xe1\x26\x9a\x7b\x67\x61\x03\xba\x43\x9d\xdd\x35\xc1\x54\xd5\x10\x3e\x5b\x05\xac\x8a\x63\xf7\x10\xca\x8b\x0b\x0d\x59\x76\x5e\xa1\x3c\x41\xae\x5d\x9b\xc4\x1d\xce\xf5\x9f\xc5\x0d\xf6\x4f\xd5\x47\xd7\x62\xe7\xed\x66\x50\x6e\xf7\x17\x69\xbe\x2d\x38\xae\x48\x73\x96\x96\x22\x5d\x8b\x6c\x2d\xd2\xef\xa8\x60\x0a\x44\x9e\xa3\x60\xca\xee\x05\xed\xed\x8b\x2c\xe9\xf6\x4a\x64\x1e\xf3\xf9\xa3\xa0\x5c\xa3\xb4\x0a\x80\xb2\xec\xc0\x98\x1c\x91\x50\x3c\xec\xda\xdc\x6a\x1d\xc3\x39\xf2\xd0\x26\x3a\xaf\x82\xcf\xee\xdf\x90\xd1\x45\x5e\xf3\xee\xb9\xfb\x93\x48\x79\x90\x8f\xde\x8a\x38\xbf\x68\xd6\xe6\xe5\x13\xde\x63\xa0\x95\x7b\xdb\xa6\xc9\xc8\xfb\x55\x88\xf8\xd8\x0c\x51\x9a\x48\xfd\xef\xa2\x74\xf2\x75\x3d\x0f\xcd\xc4\x7c\xed\xfd\x9c\x77\x77\x70\x06\xfd\xfe\x0f\xdd\xfe\x45\xb7\x3f\xb8\xb9\xb8\x1a\xf6\x5f\x0f\xfb\x57\xbf\xf6\x7f\x1c\xf6\xfb\x66\xa3\xea\xb6\x59\xb5\xbc\x5e\x52\x65\x0d\x4d\xd1\xcc\x34\x8d\xb1\x55\xc0\xa3\x4d\xc0\xe1\x3d\xb3\x4c\xb7\x26\xca\xfb\x4c\x99\x92\xed\x1d\x1a\xac\x9b\xd5\xee\x6e\xd2\xc0\x94\x6f\x26\x73\xbc\x20\xcb\x8a\x9d\x56\x31\x6f\xda\x4e\xbd\x4d\xf9\x9c\x6f\x10\x0d\x59\xef\x86\xc6\xa8\x34\x89\x93\x03\xa1\x3d\xaf\x39\x4b\x0f\x1e\xaf\xb5\xc6\x38\xc9\x15\xde\xd6\x1a\xb6\x13\x63\x2f\x0e\x75\x82\xbd\x4c\x10\x4a\xf3\x7f\x9f\xef\xcd\x92\x7a\x99\xee\xaf\x3d\xf8\x7d\x08\x54\x7d\xe4\x85\xba\x5f\x9a\xdf\xe6\xfb\xed\xce\x81\xc1\xfc\xae\x4f\xa0\x35\xdf\x7f\x03\x00\x00\xff\xff\x99\xe5\x39\x8a\x03\x27\x00\x00")

func templatesViewsSesHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsSesHtml,
		"templates/views/ses.html",
	)
}

func templatesViewsSesHtml() (*asset, error) {
	bytes, err := templatesViewsSesHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/ses.html", size: 9987, mode: os.FileMode(420), modTime: time.Unix(1504173411, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesViewsSnsHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x5f\x6f\xe3\xb8\x11\x7f\xcf\xa7\x60\x95\x14\x92\x01\x5b\xba\x66\xef\x80\xd6\x6b\xab\x08\x76\xd3\x76\x81\xde\xe1\xd0\xdd\x3e\x1d\x0e\x0b\x5a\x1c\x5b\xcc\xd2\xa4\x4a\x52\x8a\x53\xc1\xdf\xbd\x20\x29\xc9\x92\x25\x39\x71\x92\x3b\xf4\x0e\xeb\x87\x5d\x89\x1c\xce\xfc\xe6\x0f\x7f\x1c\x2a\x65\x89\x08\xac\x29\x07\xe4\x25\x82\x6b\xe0\xda\x43\xfb\xfd\xc5\x82\xd0\x02\x25\x0c\x2b\xb5\xf4\x32\xbc\x81\x99\xa6\x9a\x81\x17\x5f\x20\x84\x50\x7b\xd2\x8e\x7f\x66\xb0\xd6\xd5\xa4\x15\x48\xdf\xc4\x37\x5b\xfc\x5f\xc1\xd1\x47\xba\xcd\x18\xa0\x1f\x84\xa6\x6b\x9a\x60\x4d\xcd\x18\xc8\x82\x26\xb0\x88\xd2\x37\x95\xc2\x88\xd0\x22\xbe\xa8\xfe\xeb\x18\x4f\x18\x60\xb9\xa6\x3b\x2f\x1e\x9a\xdd\x7d\xce\x30\x07\x36\x80\x6b\xf7\xb9\x8d\xd8\x81\xba\x8e\x17\xb4\x9e\x5f\x63\xb4\xc6\xb3\x5c\x81\x54\x46\x35\x8d\x6b\x54\x6a\x11\xa5\xd7\xad\x55\x39\xab\x97\x70\x5c\x20\x8e\x8b\x15\x96\x33\x49\x37\xa9\x46\xd6\xf6\x67\x2d\x04\x5b\x89\x5d\xcb\x92\x5d\xc7\x68\xbc\xc0\x8d\x17\x82\x31\x9c\x29\x98\x31\xca\xbf\x78\x3d\x18\x49\x0a\x85\x14\x7c\x96\x67\x0e\xcb\x22\xc2\xf1\x22\x62\xb4\x05\x23\xca\x59\xeb\xed\x54\x80\x5a\x01\xed\x07\xa5\xce\x71\x4b\x97\xc6\x2b\x06\x4d\x3a\xed\x8b\xfd\x77\xa6\xb4\xa4\x19\x10\x44\xb0\xc6\x76\xe4\xd8\x45\x9d\x02\x26\xc7\x63\xb2\x3b\x50\x09\xc6\x37\x89\xa6\x05\x20\xd5\x04\x59\xa7\x47\x2b\xa3\xe3\xa5\x46\x66\xc0\xc0\x4a\x90\x87\xee\x58\x59\x22\x89\xf9\x06\xd0\x15\xc7\x5b\x98\xa2\xab\xca\x0a\x9a\x2f\x51\x58\x5b\x34\x55\xfd\x04\xa4\x24\x2e\x4b\xa7\x07\xed\xf7\x8b\x48\x1f\x9b\xef\xa1\x2c\x4b\x04\x9c\xf4\xb4\x47\x47\x38\x17\x91\x0d\xe1\xa3\xf5\xfe\x9a\x15\x7d\x93\x65\xac\xda\x73\xff\x9f\x55\x5d\xa9\xac\x97\x11\x29\x32\x22\xee\xb9\x37\x90\x18\x8c\x52\x09\xeb\xa5\x77\xe9\x1d\x4b\xcf\xb4\xd8\x6c\x18\x78\xb6\x4e\xab\x97\x96\x2e\x24\x85\x79\x5f\xe5\x5a\x0b\xee\x21\x2c\x29\x9e\xc1\x2e\xc3\x9c\x00\x31\x48\x99\x82\x3e\xf4\x7b\x09\x3c\x49\x0f\xb0\xfb\x78\x0e\x01\x6c\x70\x6c\x81\xe7\xb5\x39\xfb\xdc\x5f\x56\xc7\x70\x70\xa2\xe3\xe7\x1d\x2e\xb0\x4a\x24\xcd\xf4\xbc\x10\x94\x04\xdf\x4c\x3c\x44\xc9\xd2\xc3\xad\x9c\xce\xf2\x8c\x60\x5d\x7b\xee\x5e\x64\x57\x64\x04\x43\x63\xae\xf1\x7b\xc3\x1e\xb2\x94\x26\x82\xa3\xe6\x69\x26\x61\x2d\x41\x55\x61\x40\xff\xb6\xfa\xc7\x91\x0f\x85\xc9\x4d\x0c\x39\xdc\x25\xb4\xbe\xdc\xeb\x13\xde\x13\x18\x2f\x15\x05\xc8\x2e\xfb\xf5\xa3\xfe\x04\x16\xac\xc6\xe3\x0f\x37\xdf\xf7\x89\xae\x9e\x6c\xed\x4e\x74\xf3\xaf\x1f\xc6\x05\xdf\x81\xac\x0e\x4f\x40\xb0\xcb\xa8\xb4\x6b\xc6\xe5\x6f\x39\xc9\x04\xe5\x7a\x80\x64\x6b\x91\x7f\x62\xa5\x91\x2b\x98\x41\x26\x3e\x97\x77\xe9\x14\x5d\xb5\x62\x64\x89\xb7\x1d\xb3\x01\xf2\x45\x65\x49\xd7\x88\x0b\xdd\x59\x19\xde\x72\x13\x7d\xb2\xdf\xd7\xbb\xcb\x18\x90\x5e\x59\x02\x27\xfb\xfd\x20\x63\x37\x99\x84\x9d\x9e\x25\xc0\x35\xc8\x91\xb2\x2f\x4b\x44\xd7\x5d\x7b\x1f\xd4\x87\x9b\xef\x8f\xe1\x35\xca\x55\x86\xf9\xa9\x1d\x22\xbe\x78\x48\xe9\x07\xb3\xe3\x13\xc1\x84\x9c\xa3\xcb\xef\x92\xd5\x9f\xbf\x4b\x4c\x71\x9a\xc5\xa3\x38\x80\x29\x78\xb6\x59\x09\x5b\x51\x40\xcf\x34\xf9\xcb\x77\x6f\xbe\x5d\x3f\x6e\xba\x7f\x5c\x21\x97\xf6\xc1\x32\x76\x27\x62\x3b\x66\x37\x92\x0f\x1e\x8e\xb5\xfc\x53\x63\xdf\x2a\xec\xdb\xa6\xae\xdf\x9b\x32\x1f\x09\xcc\x31\x90\xe7\x28\x38\xdb\xfb\xe1\x14\xb5\xfb\x63\x29\x36\x12\xd4\x29\xae\x1d\x12\x9f\xad\xb0\x44\xed\x97\x99\xca\x93\xc4\xe8\xa9\xf3\xba\xa5\x7c\x76\x4f\x89\x4e\xe7\xd7\xb0\x7d\xeb\x9e\x8e\x43\xf0\x77\xd0\xd5\x96\xf9\x11\xa4\x29\x7e\xb4\xdf\xff\xf1\x11\xd6\x1f\xd7\xf1\x4e\xe4\x56\xc3\x09\x96\x6f\x78\xf6\xf9\x7e\xde\x63\xc9\x29\xdf\x9c\xeb\xe7\x7b\xaa\x5e\xee\x68\xad\xe4\x25\x9e\x8e\x4c\x3d\x7d\x03\x19\xfa\x75\xc7\xe9\x6f\xbf\xc9\xfc\x24\x32\x9a\x7c\x6d\x2f\x7f\x27\xed\xa5\xb6\xd9\x1c\x69\x2c\xdd\xe4\xd7\x96\xf2\x25\x77\xe8\xc3\xd3\x2c\x31\x14\xd4\xeb\x55\xce\xba\x55\xdb\xcd\x37\xdc\x3d\xbe\xca\x7d\xda\xf4\x75\x36\xeb\xb6\xa3\x73\xf9\x3f\xeb\x22\x6d\x97\x84\x16\xe6\x58\xd7\xf0\x9b\x62\xbb\x8f\xf9\xca\x6d\x9b\xaf\x77\xea\xdf\x0f\xe9\xa9\x76\x52\x47\xb8\xaf\x23\xf3\x95\x02\x7f\x75\x0a\xac\xc6\xe3\xf6\x06\x3c\x7d\x6d\xae\xaf\xc1\xaf\x76\xc1\x6d\x97\x80\xfb\xb4\xd8\xae\x89\xb3\x68\xb1\xbd\x32\x6c\xbb\xf4\xd8\xd5\xaa\xb7\xb8\xf6\xf2\x57\xa3\xd6\x83\x86\x8b\xd6\xdf\x0e\x4c\x3c\xdd\x1f\x0e\x0c\x6f\x55\xfb\x2d\x22\x58\xa5\x2b\x81\x25\x89\xb0\x52\xa0\x55\x54\x00\x27\x42\xaa\xa8\xa9\x00\x15\x72\xd0\xb3\x95\x8a\x12\xe5\x46\x3f\xb9\xd1\x95\x10\x5a\x69\x89\xb3\x70\x4b\x79\x98\x28\xf5\xd7\x62\x59\x96\x28\x6c\x7d\x34\x09\x57\x39\x65\x06\x88\x87\x24\xb0\xa5\x67\x2f\x13\x2a\x05\x30\x65\xf9\x2c\x18\x6b\xba\x03\x62\x3c\x01\x59\x43\xb2\x43\xff\xb0\x43\x2f\xc6\x34\x1c\xb9\x3b\xe5\xe2\xe6\x52\x8a\x94\x4c\x9e\x8e\x38\xba\x53\xd1\xdd\x7f\x72\x90\x0f\x61\x2b\x76\x06\xdd\xdd\x49\x70\xf1\x22\x72\xe6\xe2\xe7\xd9\x35\xd1\xb9\x3b\x91\xaf\x5f\xd8\x7a\x2b\x4f\x47\x30\xda\xe9\x3a\x0b\x48\x83\x44\x3f\x64\xd0\xf9\xde\x17\x1d\x4e\x8d\x8a\xa6\xae\x02\x22\x92\x7c\x0b\x5c\x4f\x42\x09\x98\x3c\x04\xeb\x9c\x27\x96\x15\x82\x09\x2a\x9b\x4d\x74\x15\xf8\x97\x03\x9f\x6b\xa7\x97\x9d\x26\x7b\x7a\x39\x74\xfc\xf8\x93\x30\x61\x34\xf9\xd2\xa8\x0e\x26\x65\x67\xcb\x5e\x85\x99\x50\x3a\xf0\x2f\xfd\x29\x2a\xfd\xea\x98\xf2\xe7\xe8\x2a\xd0\x29\x55\x13\x5b\x0f\x41\x33\x3e\xd9\x4f\xd1\x41\x13\x2a\x7b\xdc\x72\x4f\x39\x11\xf7\x21\x13\x55\x8c\x24\x30\x81\x49\x30\x79\xdb\x91\xdc\x4f\xde\x5e\x74\x06\x24\xe8\x5c\x72\x64\x4f\xfc\x83\x68\x47\xec\x2a\x5c\xf3\x43\x75\xde\xee\x74\x88\xd7\xfc\x6f\x94\x69\x90\x94\x6f\xc2\x2c\x57\x69\xd0\xd1\xd9\xe0\x54\xa0\x35\xe5\x1b\x35\xb5\x07\xc6\xa4\x0f\x9a\xae\x51\x23\x14\x2a\xab\xff\x03\x41\xcb\xe5\x12\xf9\xed\xb8\xfb\x43\x0e\x9b\x5f\x81\x25\x2a\xd0\xd2\x26\x2a\x01\xa9\x67\x0a\x18\x24\x5a\x48\x7f\x12\x16\x98\x05\xc7\xde\xb6\xed\x16\xce\x8e\xfd\x1e\x0b\x64\xd4\x44\x2d\xfe\x07\xe3\xc2\x4f\xd7\x3f\x87\x0c\xf8\x46\xa7\xa7\xc4\xc7\xe3\x7a\xfc\xdb\x0f\xc3\xab\x5d\x23\x68\x89\xde\x63\x0d\x61\x86\xa5\x82\xa0\x42\x30\xe6\xd4\x01\xe9\x2f\x0e\xae\x52\xc0\xe1\xde\xe2\x0b\x26\x28\x5e\x22\x32\xac\xaa\xff\x81\x66\x40\x71\xa5\x50\xcb\x1c\xc6\xeb\xf5\x78\x3b\xfa\x93\xf0\x7d\x5d\x96\x41\xd7\xe3\x44\xb0\x7c\xcb\xd5\x1c\xfd\xd4\x33\x35\x1c\x1a\x13\xdb\x39\xfa\xa6\x8f\x75\x7a\x96\x82\x3f\xbd\x54\xc1\x75\x5f\x1c\xd9\xf8\x70\x02\x72\x8e\x0e\x3c\x65\x68\xe1\xb1\x9a\x35\x32\xe7\xd5\xab\x59\xf1\x49\x7c\xd4\x66\x67\x3b\x13\x2f\xa9\x0f\xdf\x7f\x6a\x45\x9c\x17\xa4\x37\xc3\x41\x12\x92\x80\x34\xc5\x30\x77\xa5\xfd\x52\x33\xdf\xbe\x56\x2e\xce\x0c\xee\xd0\x8e\x69\xbf\xfd\xdc\xc5\x45\x39\xd5\xef\xc4\x36\x63\xa0\xa1\x8d\x6a\x08\x91\x39\x5a\x42\x9c\xd1\x60\x12\x56\x5b\x24\x30\x8c\x12\x42\x01\xf2\x21\x38\x79\xc0\x20\x47\x4a\xa3\x4e\x3a\x7d\x68\x69\x6d\x0c\x47\xce\xfc\x1c\x43\x3b\xce\x5e\x54\x2f\xd5\x57\xe4\x35\x13\x58\xcf\xed\x7d\xdb\x5d\xeb\x3a\x94\x6e\x0f\x7b\xfb\x1c\xfb\x93\x93\xb5\x1c\xe2\x2c\x03\x4e\x3e\x89\xe0\x2a\x70\xa8\x42\xd7\x6b\x04\x93\xc9\x23\x2b\x05\x0f\xfc\x24\x35\x37\x06\xff\x91\xf3\x76\xd8\xfd\x50\x01\x96\x49\x1a\xf8\xfe\xd4\x92\xd9\xd4\x55\xe2\x24\x24\x12\xdf\x1f\x1f\xc6\xc7\xbf\xfd\x53\xc0\x99\x9e\xa2\x8d\xed\x64\xe1\xd5\x3f\x08\x95\x16\xd9\x8f\x52\x64\x78\x83\x9d\x4b\x8f\x41\x19\x39\x64\x5c\x06\xaa\x08\x07\xfe\x42\xb8\xcb\x54\x81\x59\x6e\xba\x2e\xc6\xbc\xf8\x63\x2a\xee\x11\x66\x6c\x11\xb9\xc9\xd8\x1f\x31\x76\x5a\x57\x75\x30\x7b\xf1\xad\x7b\x38\xa9\x6e\x7f\xdc\xe7\x5c\x1c\xcf\x98\xff\x0f\xdd\x62\xd3\xc9\xff\x2f\x00\x00\xff\xff\xe8\x8f\x42\x5a\x3e\x25\x00\x00")

func templatesViewsSnsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsSnsHtml,
		"templates/views/sns.html",
	)
}

func templatesViewsSnsHtml() (*asset, error) {
	bytes, err := templatesViewsSnsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/sns.html", size: 9534, mode: os.FileMode(420), modTime: time.Unix(1513856468, 0)}
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
	"templates/views/ses.html": templatesViewsSesHtml,
	"templates/views/sns.html": templatesViewsSnsHtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"views": &bintree{nil, map[string]*bintree{
			"ses.html": &bintree{templatesViewsSesHtml, map[string]*bintree{}},
			"sns.html": &bintree{templatesViewsSnsHtml, map[string]*bintree{}},
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


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
