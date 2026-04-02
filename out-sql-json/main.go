package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	file_util "github.com/laixhe/file-util"
)

// 配置结构
type Conf struct {
	GetPath string
	OutPath string

	Json struct {
		Select []string
	}

	Sql struct {
		Oper      string
		TableName string
		Select    []string
	}
}

type GetExt int

// 解析文件类型
const (
	TXT_EXT_TYPE GetExt = 1
	XLS_EXT_TYPE GetExt = 2
)

// 解析文件扩展类型，目前支持 txt xlsx xls
var GetExtType = map[string]GetExt{
	"txt":  TXT_EXT_TYPE,
	"xlsx": XLS_EXT_TYPE,
	"xls":  XLS_EXT_TYPE,
}

type OutExt int

// 输出文件类型
const (
	SQL_OUT_TYPE  OutExt = 1
	JSON_OUT_TYPE OutExt = 2
)

// 输出文件扩展类型，目前支持 sql json
var OutExtType = map[string]OutExt{
	"sql":  SQL_OUT_TYPE,
	"json": JSON_OUT_TYPE,
}

// sql 操作的类型
const (
	SQL_OPER_INSERT = "insert"
	SQL_OPER_UPDATE = "update"
)

var conf = &Conf{}

func main() {

	fileName := "./conf.toml"
	if len(os.Args) >= 2 {
		fileName = os.Args[1]
	}

	// 解析配置文件
	_, err := toml.DecodeFile(fileName, conf)
	if err != nil {
		log.Println("解析配置文件错误：", err)
		return
	}

	// 获取要解析文件扩展名
	getExt := file_util.GetFileExt(conf.GetPath)
	if getExt == "" {
		log.Println("解析文件扩展名为空：", conf.GetPath)
		return
	}

	getExtType, isExt := GetExtType[getExt]
	if !isExt {
		log.Println("不支持要解析文件扩展名：", getExt)
		return
	}

	// 获取输出文件扩展类型
	outExt := file_util.GetFileExt(conf.OutPath)
	if outExt == "" {
		log.Println("输出文件扩展类型为空：", conf.GetPath)
		return
	}

	outExtType, isExt := OutExtType[outExt]
	if !isExt {
		log.Println("不支持要输出文件扩展名：", outExt)
		return
	}

	// 解析文件
	var getData [][]string
	switch getExtType {
	case TXT_EXT_TYPE:
		// txt 文件
		getData, err = getTxt()
	case XLS_EXT_TYPE:
		getData, err = getXls()
	}

	if err != nil {
		log.Println("解析文件错误：", err)
		return
	}

	switch outExtType {
	case SQL_OUT_TYPE:
		err = outSql(getData)
	case JSON_OUT_TYPE:
		err = outJson(getData)
	}

	if err != nil {
		log.Println("输出文件错误：", err)
		return
	}

	log.Println("输出文件完成")
}
