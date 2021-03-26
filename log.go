package utils

import (
	"fmt"
	"os"
	"strings"
	"time"
)

//Prefix LOG前缀
const Prefix = "CLOUD-PROXY"

//Println 打印行
func Println(param ...interface{}) {
	pLen := len(param)
	if pLen == 0 {
		return
	}
	newParam := append([]interface{}{}, 0x1B, time.Now().Local().Format("2006-01-02 03:04:05"))
	newParam = append(newParam, param...)
	newParam = append(newParam, 0x1B)
	fmt.Printf("%c[0;30;38m["+Prefix+"]%s "+strings.Repeat("%v ", pLen)+"%c[0m\n", newParam...)
}

//Fatalln 打印行
func Fatalln(param ...interface{}) {
	pLen := len(param)
	if pLen == 0 {
		return
	}
	newParam := append([]interface{}{}, 0x1B, time.Now().Local().Format("2006-01-02 03:04:05"))
	newParam = append(newParam, param...)
	newParam = append(newParam, 0x1B)
	fmt.Printf("%c[0;30;38m["+Prefix+"]%s "+strings.Repeat("%v ", pLen)+"%c[0m\n", newParam...)
	os.Exit(0)
}

//Danger 警告
func Danger(param ...interface{}) {
	pLen := len(param)
	if pLen == 0 {
		return
	}
	newParam := append([]interface{}{}, 0x1B, time.Now().Local().Format("2006-01-02 03:04:05"))
	newParam = append(newParam, param...)
	newParam = append(newParam, 0x1B)
	fmt.Printf("%c[0;31;38m["+Prefix+"]%s "+strings.Repeat("%v ", pLen)+"%c[0m\n", newParam...)
}

//Info 信息
func Info(param ...interface{}) {
	pLen := len(param)
	if pLen == 0 {
		return
	}
	newParam := append([]interface{}{}, 0x1B, time.Now().Local().Format("2006-01-02 03:04:05"))
	newParam = append(newParam, param...)
	newParam = append(newParam, 0x1B)
	fmt.Printf("%c[0;30;38m["+Prefix+"]%s "+strings.Repeat("%v ", pLen)+"%c[0m\n", newParam...)
}

//Success 成功
func Success(param ...interface{}) {
	pLen := len(param)
	if pLen == 0 {
		return
	}
	newParam := append([]interface{}{}, 0x1B, time.Now().Local().Format("2006-01-02 03:04:05"))
	newParam = append(newParam, param...)
	newParam = append(newParam, 0x1B)
	fmt.Printf("%c[0;32;38m["+Prefix+"]%s "+strings.Repeat("%v ", pLen)+"%c[0m\n", newParam...)
}

//Warning 警告
func Warning(param ...interface{}) {
	pLen := len(param)
	if pLen == 0 {
		return
	}
	newParam := append([]interface{}{}, 0x1B, time.Now().Local().Format("2006-01-02 03:04:05"))
	newParam = append(newParam, param...)
	newParam = append(newParam, 0x1B)
	fmt.Printf("%c[0;33;38m["+Prefix+"]%s "+strings.Repeat("%v ", pLen)+"%c[0m\n", newParam...)
}
