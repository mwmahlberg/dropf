// Code generated by go-bindata.
// sources:
// templates/index.html
// templates/userspace.html
// DO NOT EDIT!

package main

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

var _templatesIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x91\x31\x4f\xc3\x30\x10\x85\xf7\xfe\x8a\xe3\x76\x64\xc1\xc4\xe0\x78\xa1\x6c\x48\x74\x80\x81\xd1\xa9\x2f\xd8\x52\x92\xb3\xec\x8b\xa0\xff\x1e\xc7\x71\xda\x05\x31\xe5\xe2\xf7\xbd\x77\xf2\xb3\xbe\x3b\xbe\x3d\xbf\x7f\x9e\x5e\xc0\xcb\x34\x9a\x83\xde\x3e\x00\xda\x93\x75\xeb\x50\xc6\x89\xc4\xc2\xd9\xdb\x94\x49\x3a\x5c\x64\xb8\x7f\xc2\xca\xa8\x1d\xd2\x3d\xbb\x8b\x39\x6c\xb8\x7f\x30\x2e\x71\x1c\xb4\x7f\x6c\x01\xd1\x1c\xcb\x01\x5c\x78\x49\x30\x84\x91\x32\x78\x4a\xa4\x55\x6c\xfa\xc0\x69\x82\xb2\xc5\xb3\xeb\x30\x72\x16\x04\x7b\x96\xc0\x73\x87\x6a\xe4\xaf\x30\xe3\x06\x16\x74\xb4\x3d\x8d\x50\x0c\x1d\xce\x76\x22\x34\x1f\x99\x12\xac\xa3\x56\x55\xbb\x92\x61\x8e\x8b\x80\x5c\x22\x75\x28\xf4\x53\x32\x83\x6b\xa6\xca\xef\x01\x7f\x24\x47\x9b\xf3\x37\x27\x87\xe6\xd4\xa6\xff\xc2\xaf\x74\x5d\x70\xfb\xdb\x96\xdc\xb2\x76\x6f\xbf\x88\xf0\xdc\xcc\x79\xe9\xa7\x20\x68\x5e\xd7\x6b\x6a\xb5\x69\xad\x16\xb5\xf6\x52\x5b\x2d\x42\x2d\xb8\x34\x5e\xdf\xe7\x37\x00\x00\xff\xff\xd4\x41\x1f\x1c\xb7\x01\x00\x00")

func templatesIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesIndexHtml,
		"templates/index.html",
	)
}

func templatesIndexHtml() (*asset, error) {
	bytes, err := templatesIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/index.html", size: 439, mode: os.FileMode(420), modTime: time.Unix(1464604172, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesUserspaceHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x51\xbd\x6e\xf3\x30\x0c\xdc\xf3\x14\xfc\xb4\x27\x42\x90\xe5\x1b\x6c\x2f\x6d\x33\xa6\x1d\xba\x74\x64\x22\xba\x16\x20\x4b\x86\x4d\x15\x08\x04\xbf\x7b\xf5\x63\x3b\x19\x3a\xe9\x24\xde\x1d\x8f\x54\xf5\xef\xf5\xfd\xe5\xf3\xeb\xe3\x0d\x3a\xee\x4d\xb3\xab\xca\x01\x50\x75\x84\x2a\x81\x08\x7b\x62\x84\x5b\x87\xe3\x44\x5c\x0b\xcf\xed\xfe\xbf\xc8\x1c\xb9\x92\xaa\xab\x53\xf7\x66\x97\x75\xc7\xc6\x0f\xc6\xa1\x02\x84\x56\x1b\x8a\xa4\x63\xa6\xb4\x6e\xec\x81\xec\x8d\xef\x03\xd5\xa2\xf7\x86\xf5\x80\x23\xcb\xf4\xbe\x57\xc8\x28\x00\x6f\xac\x9d\xad\x85\x2c\x0e\x02\x62\xe7\xce\xa9\x5a\x0c\x6e\x62\xb1\xa4\xd1\x76\xf0\x0c\xc5\x25\x35\x10\x60\xb1\xdf\xb0\x56\x0b\x92\x7f\xd0\x27\x7f\xed\x35\x0b\xf8\x41\xe3\xe3\x75\xed\x22\xcb\x30\x29\x48\x42\x21\x80\x6e\xe1\x70\x8e\x2e\xd3\x3c\xa7\x12\xe3\xd5\xd0\xe2\xc7\x8f\xc5\xe4\x5b\x73\xce\x43\x46\x50\xea\xf2\x89\x10\xc2\x88\xf6\x9b\x9e\xbd\x92\x66\x7c\xc8\x55\x13\xc2\xe1\x12\xf3\xcf\x73\x54\xaa\xcd\x63\x5c\x0d\xc8\xaa\x92\x41\x6e\x21\x62\x3e\x32\x13\x41\x79\xef\x4e\xcd\xc5\xe5\x4d\x4f\x50\x06\x22\x05\x77\xe2\xb8\xf7\xd3\xca\xb6\x6a\x21\xcb\xf2\x4f\xb1\x96\xbf\xf9\x37\x00\x00\xff\xff\x95\x60\x76\x98\xfe\x01\x00\x00")

func templatesUserspaceHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesUserspaceHtml,
		"templates/userspace.html",
	)
}

func templatesUserspaceHtml() (*asset, error) {
	bytes, err := templatesUserspaceHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/userspace.html", size: 510, mode: os.FileMode(420), modTime: time.Unix(1465028590, 0)}
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
	"templates/index.html": templatesIndexHtml,
	"templates/userspace.html": templatesUserspaceHtml,
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
		"index.html": &bintree{templatesIndexHtml, map[string]*bintree{}},
		"userspace.html": &bintree{templatesUserspaceHtml, map[string]*bintree{}},
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
