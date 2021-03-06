# httpd

纯golang实现的简单、快速的http服务


### 使用源码安装

```shell
go get -u github.com/zituocn/httpd
```

```shell
cd github.com/zituocn/httpd
go build
```

### 下载二进制版本

* [Windows x64](https://github.com/zituocn/httpd/releases/download/v0.0.1/httpd_windows_amd64.exe)
* [Linux x64](https://github.com/zituocn/httpd/releases/download/v0.0.1/httpd_linux_amd64)
* [MacOS x64](https://github.com/zituocn/httpd/releases/download/v0.0.1/httpd_darwin_amd64)
* [MacOS ARM64](https://github.com/zituocn/httpd/releases/download/v0.0.1/httpd_darwin_arm64)
* [FreeBSD](https://github.com/zituocn/httpd/releases/download/v0.0.1/httpd_freebsd_amd64)
* [更多下载...](https://github.com/zituocn/httpd/releases)



### 使用

#### 查看参数

```go
./httpd -h
```


```shell
Usage of ./httpd:
  -path string
    	path (default ".")
  -port string
    	http listen port (default "8080")
```

#### 在当前目录运行

```shell
./httpd
```

运行状态

```shell
https://github.com/zituocn/httpd

当前路径为：.
当前端口为：8080
使用 ctrl+c 停止服务
```

浏览器访问

```shell
http://127.0.0.1:8080
```

#### 上加参数运行

```shell
./httpd -port=9001 -path /Users/samsong/Documents 
```

运行状态

```shell
https://github.com/zituocn/httpd

当前路径为：/Users/samsong/Documents
当前端口为：9001
使用 ctrl+c 停止服务
```

