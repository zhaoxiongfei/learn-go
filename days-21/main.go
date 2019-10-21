package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

/*
练习：HTTP 处理
实现下面的类型，并在其上定义 ServeHTTP 方法。在 web 服务器中注册它们来处理指定的路径。

type String string

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}
例如，可以使用如下方式注册处理方法：

http.Handle("/string", String("I'm a frayed knot."))
http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
注意： 这个例子无法在基于 web 的用户界面下运行。 为了尝试编写 web 服务，你可能需要 安装 Go。
*/
