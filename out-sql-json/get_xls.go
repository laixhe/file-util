package main

import (
	"errors"

	"github.com/xuri/excelize/v2"
)

// 解析 xlsx xls 文件
func getXls() ([][]string, error) {

	data := make([][]string, 0)

	file, err := excelize.OpenFile(conf.GetPath)
	if err != nil {
		return nil, errors.New("打开解析 xlsx 或 xls 文件失败：" + err.Error())
	}

	// 获取所有工作表
	for _, sheet := range file.GetSheetMap() {

		// 获取 工作表 上所有单元格
		rows, err := file.GetRows(sheet)
		if err != nil {
			return nil, errors.New("打开" + sheet + "单元格失败：" + err.Error())
		}

		// 获取单元格数据
		for rowK, row := range rows {
			if rowK == 0 {
				continue
			}

			if len(row) > 0 {
				data = append(data, row)
			}
		}

	}

	if len(data) == 0 {
		return nil, errors.New("解析 xlsx 或 xls 文件没有数据！")
	}

	return data, nil
}
