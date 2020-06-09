package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

// 解析 txt 文件
func getTxt() ([][]string, error) {

	file, err := os.OpenFile(conf.GetPath, os.O_RDONLY, 0600)
	if err != nil {
		return nil, errors.New("打开解析txt文件失败："+err.Error())
	}

	data := make([][]string, 0)

	// 读入缓存
	buff := bufio.NewReader(file)
	for {
		line, _, err := buff.ReadLine()
		if err != nil || io.EOF == err {
			break
		}

		if len(line) > 0 {
			data = append(data, strings.Fields(string(line)))
		}
	}

	if len(data) == 0 {
		return nil, errors.New("解析txt文件没有数据！")
	}

	return data, nil
}
