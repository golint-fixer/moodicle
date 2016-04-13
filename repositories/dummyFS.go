package repositories

import (
	"errors"
	"fmt"
	"github.com/blang/vfs"
	"os"
	"strings"
)

type dummyFS struct {
	vfs.Filesystem
}

func (fs dummyFS) ReadDir(path string) ([]os.FileInfo, error) {

	return []os.FileInfo{DumFileInfo{FileName: "1234.json"}, DumFileInfo{FileName: "5678.json"}}, nil

}

func (fs dummyFS) OpenFile(name string, flag int, perm os.FileMode) (vfs.File, error) {

	s := strings.Split(name, "/bob/")
	id := strings.Split(s[1], ".")[0]
	if id == "UNKNOWN" {
		return nil, errors.New("An error")
	}

	return DumFile{Id: id}, nil
}

func (fs dummyFS) Remove(name string) error {
	return nil
}

// Stat returns dummy error
func (fs dummyFS) Stat(name string) (os.FileInfo, error) {
	return DumFileInfo{}, nil
}

type DumFile struct {
	Id string
	vfs.DumFile
}

// Read returns dummy error
func (f DumFile) Read(p []byte) (n int, err error) {

	copy(p, []byte(fmt.Sprintf("{\"id\":\"%s\"}", f.Id)))

	return 13, nil
}

func (f DumFile) Close() (err error) {
	return nil
}

type DumFileInfo struct {
	FileName string
	vfs.DumFileInfo
}

// Read returns dummy error
func (f DumFileInfo) Size() int64 {
	return 13
}

// Read returns dummy error
func (f DumFileInfo) Name() string {
	return f.FileName
}
