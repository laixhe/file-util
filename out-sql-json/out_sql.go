package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var replacerSql = strings.NewReplacer(
	"'", "\\'",
	)

func outSql(getData [][]string) error {

	oper := strings.ToLower(conf.Sql.Oper)
	switch oper {
	case SQL_OPER_INSERT:
	case SQL_OPER_UPDATE:
	default:
		return errors.New(fmt.Sprintf("不支持 %v sql 操作的类型", oper))
	}

	selectLen := len(conf.Sql.Select)
	selectData := ""

	for _, data := range getData {
		if len(data) != selectLen {
			log.Printf("数据长度不等于 select:%v\n", data)
			continue
		}

		if oper == SQL_OPER_INSERT {
			selectData += sqlInsert(data)
		}

	}

	if len(selectData) == 0 {
		return errors.New("没有数据输出到文件 sql")
	}

	// 写入文件
	err := ioutil.WriteFile(conf.OutPath, []byte(selectData), 0666)
	if err != nil {
		return err
	}

	return nil
}

func sqlInsert(data []string) string {

	selectData := fmt.Sprintf("INSERT INTO `%v`(", conf.Sql.TableName)

	for _, v := range conf.Sql.Select {
		selectData += fmt.Sprintf("`%v`,", v)
	}

	// 去掉最后一个 ,
	selectData = selectData[:len(selectData)-1]
	selectData += ") VALUES ("

	for _, v := range data {
		selectData += fmt.Sprintf("'%v',", replacerSql.Replace(v))
	}

	// 去掉最后一个 ,
	selectData = selectData[:len(selectData)-1]
	selectData += ");\r\n"

	return selectData
}