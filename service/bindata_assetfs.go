// Code generated by go-bindata.
// sources:
// templates/index.tpl.html
// DO NOT EDIT!

package service

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

var _templatesIndexTplHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x58\x6b\x6b\xdb\x3c\x14\xfe\x9e\x5f\x21\xdc\xbe\xdf\x5e\xc7\xbc\x97\xc2\x96\x39\x86\xd0\x96\x51\x76\x61\xac\xdb\xbe\x2b\x96\x92\x88\x29\x92\x91\xe4\xb4\x21\xe4\xbf\xef\x48\xb6\x53\x39\xf5\x25\xe9\x65\x74\x0d\x04\x6c\xe9\x5c\x9e\xa3\xf3\x9c\x23\xc9\x9b\x0d\x62\x22\xe5\x39\xa1\x28\x98\x29\x29\x0c\x15\x24\x40\x01\xc7\x6b\x99\x9b\x68\x41\x31\xa1\x6a\x68\x32\x3e\x5c\x98\x25\x0f\xd0\x76\x3b\x18\xc4\x84\xad\x90\x92\x9c\x8e\x03\x83\xa7\x19\x16\x94\x07\xc9\x00\xc1\x2f\xce\x39\x4a\x39\xd6\x7a\x1c\x08\xbc\x42\xf0\x0f\x41\x42\x07\x77\xd2\x9c\x69\x53\x0a\x3b\x05\xce\xca\xb9\x4c\x51\x4d\x85\xc1\x86\x49\x11\x54\x46\x70\x6a\xd8\x8a\x06\x49\x8c\xd1\x42\xd1\xd9\x38\x38\xd1\x54\xad\x58\x4a\xc1\x24\x56\x0c\x87\x29\x00\x06\x7d\x10\xbd\x9b\xd8\xf9\x0a\x10\xc1\x06\x87\x46\xce\xe7\xd5\x48\x72\x5d\x8a\xc5\x11\x4e\xe2\x88\xb3\x5e\x28\x9e\x6f\x9c\x65\x9c\xa5\x6e\xf8\xbe\xff\xfa\x64\x27\x86\x89\x27\xfa\x10\x1c\x46\x66\x2c\xbd\x8f\xa0\x1a\xee\xf4\xfd\xcd\x09\x3d\xc4\xab\xce\xa7\x3a\x55\x2c\x6b\x0e\x7f\x6f\xb6\x3b\x07\xbe\x6c\x1d\x4a\x1c\xe5\x3c\x19\x14\x8f\x96\x65\x25\x0f\x40\xcd\xf9\x02\x5c\x3e\x79\x1a\x78\xe8\x6b\xd8\x11\x54\x52\x08\x31\xe2\x71\xe4\xce\x46\x83\x27\x4e\x43\x58\x82\x0c\xb0\x39\xee\xd5\x44\x9d\xb8\x93\xa9\x29\xa0\x42\x6d\x21\x57\x54\x95\xcf\xda\x40\x84\x94\x34\xe8\x17\x36\x6c\x65\xb5\xce\x25\x13\x87\x1a\xe9\x1d\x5b\x61\xac\x59\x38\xea\xb4\x34\x95\x64\xdd\x3c\x07\x75\xaf\xb0\x98\x53\x74\x2a\xf0\x92\xfe\x8d\x4e\x4b\x5f\x68\x34\x46\xc3\xaa\x4a\x6c\xb9\x37\x1b\x56\xcd\x56\x8b\x49\x92\x80\x75\x67\x17\x0c\x00\xc2\x36\x78\x51\x9b\x19\x50\x87\x36\xd4\xea\x3d\x6a\x89\x0b\x26\xec\xda\xef\x65\x37\x82\xf4\x7a\xa4\xd9\x7b\x3d\x80\x43\x05\x79\x6a\x05\xfe\xa2\x09\x54\xce\x27\x27\xed\xb4\xa9\x44\xae\x26\x9f\xfa\x85\xbc\x7e\x85\x26\x5f\x3f\xf7\x2b\x9c\x53\x65\xd8\xcc\xaa\x50\x44\x6f\x33\xa6\x9c\x6e\xbf\xde\xa5\x20\x99\x64\xc2\x74\xf0\xbd\x12\xfd\x88\xb5\x41\x79\x06\xcd\x85\x3e\x6b\x71\x30\xa8\x0c\x2f\xf3\xae\x3a\xfc\xfe\xdd\x51\x21\x60\x85\xcd\x90\x90\xa6\x66\x61\x78\x29\x6c\x7a\xc9\x76\x5b\xe6\x9f\x58\x47\x2a\xd8\x6c\x80\xf1\xdb\x6d\x6f\x59\x61\x42\x00\x15\xfa\xa7\xb3\xb2\x4a\xf9\x1d\xc5\xe8\xad\x09\x53\x68\x9e\xe0\xa7\x5d\xa1\x08\x1c\x20\xd7\xe0\x5e\x69\xe0\x48\x4b\x90\x3b\x57\x1a\xca\xa4\x72\x36\xe7\xeb\x6c\xc1\xa0\x5b\xa3\xdd\x53\x28\x7f\x06\x48\x9b\xb5\xad\xb2\x54\x72\xa9\x46\xe8\xe4\x2c\x9d\xbe\x39\x4b\x61\x83\x89\xac\x72\x1f\x2a\xca\x35\x7d\x2c\x08\x45\x97\xd2\xee\x04\x7b\x40\xc8\xdb\xb3\xff\xfe\x9f\x1d\x0c\xc4\x26\xa9\x7d\xcd\xfb\x32\xe2\x1a\xa3\xbf\xbc\x13\x25\x0e\xc9\xe4\x3d\x3d\xaf\xc4\x2e\x77\x15\x76\x61\x0b\xee\x10\x6b\xdd\x0b\xe9\x75\xb3\x4c\xc9\x39\xf4\xb2\xfd\x86\x77\xa8\x5a\x38\xc5\x0a\xf9\x2f\xa1\xce\xd3\xd4\xda\xab\xb2\xb0\x64\x22\xbc\x61\xc4\x2c\x46\xff\xd2\xe5\xbb\xe2\x69\x3f\xd6\xf7\xd4\x94\x45\xf3\x85\x2a\x4b\x63\x88\xf1\xaf\x03\x10\xd9\x5f\xbb\xad\x73\x99\x3b\x4b\xfd\x81\xd5\x77\x8c\xc7\xc5\x7f\x83\x95\x60\x62\x7e\x6c\xfc\x17\x4c\x3f\xdd\x02\x54\xc6\x9e\x72\x05\x7a\x44\x8e\x27\xb8\x6d\xf0\xdf\x5d\x7f\x7f\xb1\x67\x88\x07\x1c\x22\xca\x33\xfa\x6b\x39\x3e\xb8\xdb\x44\xf7\x99\xe0\x89\xf6\x60\xb7\x70\x6e\xf7\x2d\x6e\x30\x8f\x39\x99\x1e\xb3\x85\x3a\x5e\x3a\xe7\x85\xe3\xbe\x7e\xfd\x87\xf1\xb1\x7e\x6d\x7b\x2d\xb4\xf4\x2f\x98\x87\x9d\x58\xab\x93\xe7\xb3\xf3\xd8\x5f\xf0\xe2\xaa\xe5\x67\xe0\xb7\xb2\xda\x87\x52\x83\x71\xcc\x99\xa4\x66\xa4\x5a\xc5\x17\x5b\x21\xde\x63\x55\x33\x71\x56\x3b\x25\x2b\x36\x5f\x54\xdf\x18\xae\x2f\x3e\x20\x60\xae\x86\xd0\x46\x16\xdc\x10\x06\x7e\x14\xef\x16\x64\x1c\x65\xa0\xdf\xfd\xfd\x6c\x26\xa5\xd9\xfb\x7e\xf6\x2b\x00\x00\xff\xff\x5c\x1b\x8e\x28\x71\x13\x00\x00")

func templatesIndexTplHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesIndexTplHtml,
		"templates/index.tpl.html",
	)
}

func templatesIndexTplHtml() (*asset, error) {
	bytes, err := templatesIndexTplHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/index.tpl.html", size: 4977, mode: os.FileMode(420), modTime: time.Unix(1480267224, 0)}
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
	"templates/index.tpl.html": templatesIndexTplHtml,
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
		"index.tpl.html": &bintree{templatesIndexTplHtml, map[string]*bintree{}},
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
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: k}
	}
	panic("unreachable")
}
