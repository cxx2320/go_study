# GO MOD 使用

使用go mod 管理项目，就不需要非得把项目放到GOPATH指定目录下，你可以在你磁盘的任何位置新建一个项目

## 使用
* 新建一个文件夹（test）
* 进入文件夹，打开终端，输入`go mod init <dir> ` dir为当前目录名称。会在当前目录生成一个`go.mod`的文件
    测试运行一下代码
    ```go
    package main

    import (
        "fmt"
        "github.com/jinzhu/configor"
    )

    func main() {
        fmt.Println("test",configor.Config{})
    }
    ```
    第一因为需要引入外部包。
    运行后会在当前目录生成`go.sum`文件,是记录所依赖的项目的版本的锁定是记录所依赖的项目的版本的锁定

## 自定义包
* 在gomod文件夹下创建test_math文件夹
* 创建一个简单的go函数
    ```go
    package test_math

    func Add(a int,b int) int {
        return a+b
    }
    ```
* 在`main.go`使用`mod_test/test_math`
* 调用`test_math.Add(200,30)`函数
* 输出 `230`