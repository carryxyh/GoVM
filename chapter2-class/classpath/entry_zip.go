package classpath

import (
	"path/filepath"
	"archive/zip"
	"io/ioutil"
	"errors"
	"fmt"
)

type ZipEntry struct {
	absPath string
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	fmt.Println("------------------" + self.absPath)
	//打开压缩文件,错误直接返回
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	for _, f := range r.File {
		//打开一个ReadCloser，用来获取文件的内容
		rc, err := f.Open()
		if err != nil {
			return nil, nil, err
		}
		defer rc.Close()
		//这里用readAll，从Reader中
		data, err := ioutil.ReadAll(rc)
		if err != nil {
			return nil, nil, err
		}
		return data, self, err
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}