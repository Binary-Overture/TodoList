package util

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"time"
)

var LogrusObj *logrus.Logger

func InitLog() {
	if LogrusObj != nil {
		src, _ := setOutputFile()
		// 设置输出
		LogrusObj.Out = src
		return
	}
	//实例化
	logger := logrus.New()
	file, err := setOutputFile()
}

// setOutputFile
// 设置输出文件路径
func setOutputFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err != nil {
		logFilePath += dir + "/logs/"
	}
	//查询文件状态
	_, err := os.Stat(logFilePath)
	//判断文件是否存在
	if os.IsNotExist(err) {
		//如果新建文件并赋权限错误
		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	//命名log文件名格式
	logFileName := now.Format("2006-01-02") + ".log"
	//创建文件夹
	fileName := path.Join(logFileName, logFilePath)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return src, nil
}
