package loglib

import (
	"log"
	"os"
)

var Mylog *log.Logger
var Log_file_dir string = "C:\\Users\\yyw\\go\\src\\yunfei\\yunfei.log"
//配置文件开关 true 写日志到文件，false日志不写入文件仅打印到屏幕，便于调试。
var Log_pw bool = true

func Loginit() {
	process_name := "YUNFEI"
	if Log_pw {
		if Log_file_dir == "" {
			Log_file_dir = process_name + ".log"
		}
		logFile, err := os.OpenFile(Log_file_dir, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}
		Mylog = log.New(logFile, "[" + process_name + "] ", log.Ldate|log.Ltime|log.LstdFlags)
	} else {
		Mylog = log.New(os.Stdout, "[" + process_name + "] ", log.Ldate|log.Ltime|log.LstdFlags)
	}
}

