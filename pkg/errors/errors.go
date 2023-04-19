package errors

import "log"

// PrintInStdout 输出命令行打印
func PrintInStdout(err error) {
	if err != nil {
		log.Fatalln("出现错误: ", err.Error())
	}
}
