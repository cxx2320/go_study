package main

import (
	"net/http"

	"log"
)

type HandleFnc func(http.ResponseWriter, *http.Request)

func main() {

	http.HandleFunc("/test1", logPanics(SimpleServer))

	if err := http.ListenAndServe(":8088", nil); err != nil {

		panic(err)

	}

}

func SimpleServer(writer http.ResponseWriter, req *http.Request) {

	// 手动抛异常，交给logPanics函数的defer处理
	// panic("error SimpleServer")

	writer.Write([]byte("Hello Word"))
}

func logPanics(function HandleFnc) HandleFnc {

	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {

			if x := recover(); x != nil {

				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)

				// 下面一行代码是译者添加，默认出现 panic 只会记录日志，页面就是一个无任何输出的白页面，
				// 可以给页面一个错误信息，如下面的示例返回了一个 500
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			}

		}()

		function(writer, request)

	}

}
