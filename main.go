package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

var(
	tip = `
https://github.com/zituocn/httpd

当前路径为：%s
当前端口为：%s
使用 ctrl+c 停止服务
`

)

func main(){
	port:=flag.String("port","8080","http listen port")
	path:=flag.String("path",".","path")
	flag.Parse()

	if *path =="/"{
		*path = getPath()
	}
	fmt.Printf(tip +"\n",*path,*port)
	err := http.ListenAndServe(fmt.Sprintf(":%s",*port),http.FileServer(http.Dir(*path)))
	if err!=nil{
		fmt.Printf("httpd 启动失败 -> %s \n",err.Error())
	}
}

func getPath() string{
	fmt.Println("路径为根目录 或 未输入目录，使用当前目录")
	path,err:=getCurrentPath()
	if err!=nil{
		fmt.Println("获取当前路径失败: ",err)
	}
	return path
}


func getCurrentPath() (string,error){
	return filepath.Abs(filepath.Dir(os.Args[0]))
}