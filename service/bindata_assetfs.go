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

var _templatesIndexTplHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x57\x41\x4f\xe3\x3c\x10\xbd\xf7\x57\x8c\x0c\xc7\xaf\x8d\xbe\x2b\x4a\x23\x55\xc0\x69\x57\x7b\x58\x76\xf7\xee\xc6\x6e\x6b\xc9\xd8\x96\xed\x54\x20\xc4\x7f\xdf\x89\x13\xb7\x4e\x49\x52\xe8\x2e\xa8\x0b\x12\x52\xea\x79\x33\xf3\xec\x79\x33\x71\x9e\x9e\x40\xa8\x52\x56\x8c\x03\x59\x59\xad\x3c\x57\x8c\x00\x91\xf4\x51\x57\x3e\xdb\x70\xca\xb8\x9d\x79\x23\x67\x1b\x7f\x2f\x09\x3c\x3f\x4f\x26\x39\x13\x5b\xb0\x5a\xf2\x39\xf1\x74\x69\xa8\xe2\x92\x14\x13\xc0\xbf\xbc\x92\x50\x4a\xea\xdc\x9c\x28\xba\x05\xfc\x9f\x22\xc2\x91\x3d\x5a\x0a\xe7\x5b\x70\x70\x90\xa2\xb5\x19\xcb\x1d\x57\x9e\x7a\xa1\x15\x89\x41\x68\xe9\xc5\x96\x93\x22\xa7\xb0\xb1\x7c\x35\x27\x17\x8e\xdb\xad\x28\x39\x86\xa4\x56\xd0\x69\x89\x84\xd1\x1f\xa1\x7b\xc3\x2e\x17\x01\x46\x3d\x9d\x7a\xbd\x5e\xc7\x95\xe2\xae\x85\xe5\x19\x2d\xf2\x4c\x8a\xa3\x54\x92\xdc\xd4\x18\x29\xca\xb0\xfc\x32\x7f\xd7\x38\xca\x61\x91\x40\x4f\xe1\xe1\xb5\x11\xe5\x4b\x06\x71\x79\x34\xf7\x8f\x00\x3a\x25\xab\xab\x96\xae\xb4\xc2\xf4\x6f\xff\xc0\x3a\x5e\x83\x14\xdb\xa5\x92\x67\x95\x2c\x26\xcd\x63\xad\xb2\x56\x07\xe8\x16\x72\x21\xaf\x54\x3c\x3d\x3a\x4c\x3d\xea\x15\x68\x25\x04\x82\x25\x1a\xd9\xc7\xe8\xc9\x24\xf9\x14\x8f\xc0\x20\xb7\xa0\xbd\x0e\x34\xc0\x03\xa6\xe3\x00\x8d\xdb\x46\x6f\xb9\x6d\x9f\x9d\xc7\x1d\x72\xd6\xe3\xdf\xc4\xa8\x3b\x6b\xd0\x56\x2c\x02\x6b\x70\x3b\xb5\xe2\x5a\x3f\x38\x1b\x8d\xb4\xd4\xec\xb1\xdf\x86\x7d\x6f\xa9\x5a\x73\xb8\x54\xf4\x9e\xff\x07\x97\x6d\x2e\xb8\x9a\xc3\x2c\x76\x49\xdd\xee\xfd\x81\x6d\x7f\xd4\xc6\xc8\x0a\x8c\x1e\xe2\x62\x00\x64\x38\x44\x2f\x1b\x0a\x83\xee\x38\x86\x06\xb3\x67\x03\xfb\x42\x43\x7d\xf6\x07\xd5\xcd\xb0\xbc\x89\x68\x0e\x7e\xbe\x42\x43\x8d\x78\x3a\x0d\x7e\xd6\x02\x6a\xed\xc5\xc5\xb0\x6c\x22\x24\x19\x45\xb0\xf8\xfe\xed\xb8\xc3\x35\xb7\x5e\xac\x6a\x17\x0e\xfc\xc1\x08\x1b\x7c\x8f\xfb\xdd\x2a\x66\xb4\x50\x7e\x44\xca\x11\xfa\x95\x3a\x0f\x95\xc1\xb9\xc1\xdf\x55\xf7\x02\x45\x9f\x14\x35\x08\x3f\x1d\xcd\x7f\x22\x7e\xca\x18\x26\x80\xff\x47\xf5\x9f\xe0\x53\x22\xb3\x85\x55\x27\xf9\x25\xb5\xb9\xdd\x95\xe6\xa6\xae\xd4\x29\xd1\x76\x15\xbb\xd6\x95\xf2\x18\x02\x32\x18\xc4\xdc\xaa\x5a\xb3\x2c\x42\xdf\x9e\xad\x2e\xfa\xcf\x50\xf3\xb3\x1d\x19\x27\xcc\x8c\xf6\x95\xfc\x59\xa6\x45\xb8\x3c\x8c\xcf\x89\xbf\xd4\x97\xe1\xe0\x42\x47\x36\x17\x96\x0f\xed\xc5\x90\xbc\x49\x7c\xac\x15\xff\x31\x3d\x76\x6f\x69\x9f\x45\x96\xe9\x7d\xf2\x75\x6f\xb1\x38\xb7\xde\x5d\xc7\xe9\x81\x37\x37\xab\xb4\x02\x1f\xaa\xea\x94\x4a\x87\xc6\x5b\x5e\x37\x9d\x20\xf1\x14\xcf\xb6\x43\x92\xc7\xd8\x33\xb9\xd9\x49\x96\x3f\xf8\xa9\x15\xeb\x4d\xfc\xa4\xb8\xbb\xf9\x02\xa8\x5c\x87\x5b\xbb\xaa\xc9\xcd\x70\xe1\x57\xf3\xbb\x26\x99\x67\x06\xfd\xc7\x3f\x97\x57\x5a\xfb\x83\xcf\xe5\xdf\x01\x00\x00\xff\xff\xad\x67\x14\x9a\x60\x0f\x00\x00")

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

	info := bindataFileInfo{name: "templates/index.tpl.html", size: 3936, mode: os.FileMode(420), modTime: time.Unix(1472768784, 0)}
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
