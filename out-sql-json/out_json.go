package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

var replacerJson = strings.NewReplacer(
	`"`, "\\\"",
)

func outJson(getData [][]string) error {

	selectLen := len(conf.Json.Select)
	selectData := "[\r\n"

	for _, data := range getData {
		if len(data) != selectLen {
			log.Printf("数据长度不等于 select:%v %v\n", data, conf.Json.Select)
			//continue
		}

		selectData += "{"
		for k, v := range data {
			selectData += fmt.Sprintf(`"%v":"%v",`, conf.Json.Select[k], replacerJson.Replace(v))
		}

		// 去掉最后一个 ,
		selectData = selectData[:len(selectData)-1]

		selectData += "},\r\n"

	}

	if len(selectData) == 1 {
		return errors.New("没有数据输出到文件 json")
	}

	// 去掉最后一个 ,
	selectData = selectData[:len(selectData)-3]
	selectData += "\r\n]"

	// 写入文件
	err := os.WriteFile(conf.OutPath, []byte(selectData), 0666)
	if err != nil {
		return err
	}

	return nil
}
