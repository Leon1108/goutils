package goutils

import (
	"os"
	"path/filepath"
)

/**
 * 路径是否存在，可能是文件或目录
 */
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

/**
 * 路径是否为文件
 */
func IsFile(filePath string) bool {
	fi, err := os.Stat(filePath)
	return err == nil && !fi.IsDir()
}

/*
 * 路径是否为目录
 */
func IsDir(path string) bool {
	fi, err := os.Stat(path)
	return err == nil && fi.IsDir()
}

func MkDir(dirPath string) (err error) {
	if !IsDir(dirPath) {
		err = os.Mkdir(dirPath, 0744)
	}
	return
}

/*
 * 在工作目录中创建目录
 */
func MkDirInWd(dirName string) (dirPath string, err error) {
	var wd string
	if wd, err = os.Getwd(); err != nil {
		return
	}

	dirPath = filepath.Join(wd, dirName)
	err = MkDir(dirPath)
	return
}
