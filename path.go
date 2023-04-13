package godash

import (
	"os"
	"path"
	"strings"
)

// GetRootPath 设置项目根路径名称，在项目下任意目录运行，返回项目根路径地址
func GetRootPath(rootDirName string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	arr := strings.Split(strings.ReplaceAll(dir, "\\", "/"), rootDirName)
	return path.Join(
		arr[0],
		rootDirName,
	), nil
}
