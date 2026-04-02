package file_util

import "strings"

// 获取文件扩展名
func GetFileExt(fileName string) string {

	exts := strings.Split(fileName, ".")
	if len(exts) <= 1 {
		return ""
	}

	return strings.ToLower(exts[len(exts)-1])
}