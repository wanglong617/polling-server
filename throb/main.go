package main

import (
	"fmt"
	"github.com/fatih/color"
	_ "happy.work/throb/routers"
)

var UserId string

func main() {
	slogan()
	config()
	errors()
}

// 初始化配置文件
func config() {

}

// 设置错误控制器
func errors() {

}

// 欢迎字符图形

func slogan() {
	/*
	 * http://www.network-science.de/ascii
	 * Font:standard
	 */
	ascii := `
 _____ _               _
|_   _| |__  _ __ ___ | |__
  | | | '_ \| '__/ _ \| '_ \
  | | | | | | | | (_) | |_) |
  |_| |_| |_|_|  \___/|_.__/  v1.0
`

	color.Set(color.FgHiMagenta, color.Bold)
	defer color.Unset()
	fmt.Println(ascii)
}
