/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-06-06 04:58:45
# File Name: file.go
# Description:
####################################################################### */

package util

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

// 返回绝对路径
func AbsPath(targetPath string, basePath string) string {
	if filepath.IsAbs(targetPath) {
		return targetPath
	}
	return path.Join(basePath, targetPath)
}

func PathExists(path string) (exists bool, isdir bool, err error) {
	f, err := os.Stat(path)
	if err == nil {
		return true, f.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, false, nil
	}
	return false, false, err
}

// 遍历目录下所有指定后缀的文件
func WalkDir(path string, suffixes []string) (files []string, err error) {
	for k, suffix := range suffixes {
		suffixes[k] = strings.ToUpper(suffix)
	}

	err = filepath.Walk(path, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}

		if len(suffixes) == 0 {
			files = append(files, filename)
		}
		for _, suffix := range suffixes {
			if strings.HasSuffix(strings.ToUpper(fi.Name()), fmt.Sprintf(".%s", suffix)) {
				files = append(files, filename)
			}
		}
		return nil
	})
	return files, err
}

//MD5
func Md5File(file string) (string, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return "", err
	}

	h := md5.New()
	_, err = io.Copy(h, bufio.NewReader(f))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

//SHA1
func SHA1File(file string) (string, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return "", err
	}

	h := sha1.New()
	_, err = io.Copy(h, bufio.NewReader(f))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

//SHA256
func SHA256File(file string) (string, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return "", err
	}

	h := sha256.New()
	_, err = io.Copy(h, bufio.NewReader(f))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func ReadFile(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func WriteFile(data string, file string) error {
	return ioutil.WriteFile(file, []byte(data), 0666)
}

func ReadLine(file string, fn func(line []byte)) (err error) {
	var f *os.File
	if f, err = os.Open(file); err != nil {
		return
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	for {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		fn(line)
	}
	return
}

func GetCurPath() (r string, err error) {
	if r, err = exec.LookPath(os.Args[0]); err != nil {
		return
	}
	if r, err = filepath.Abs(r); err != nil {
		return
	}
	r = filepath.Dir(r)
	return
}
