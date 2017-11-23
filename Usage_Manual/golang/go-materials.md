
https://golang.org/doc/install

https://gobyexample.com/

G learn golang the hard way <br>
https://github.com/gophergala/learn-Go-the-hard-way <br>
https://andy-zhangtao.gitbooks.io/golang/content/ <br>

[https://tour.golang.org/](https://tour.golang.org/list) (golang array slice)

https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/preface.md
https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/01.2.md#获取远程包



-

learnxinyminutes
https://learnxinyminutes.com/docs/go/

golang cheatsheet

golang underscore

golang xpath


= = = = = =

GOBIN: need not to set, if GOPATH(global GOPATH or current GOPATH) is set. If you have set GOPATH, don't set GOBIN: its default value will be GOPATH/bin. 但是，仍然推荐显式设置，因为可以在当前终端里运行 go env 里看到 [-](https://stackoverflow.com/questions/40067997/how-set-gobin-automatically)

（错误的GOBIN: $GOROOT/bin 是错误的GOBIN）

= = = = = =

require, include

(packages, install, go get)

https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies

= = = = = =

list packages:

(无用) `go list` 的作用：cd AppF/src/rainyApp > `go list` # rainyApp ； cd AppF/src/c21 > `go list` # c21

`go list ./...` 的作用：列出所有packages 在 workspace(src+bin+pkg)Dir 里的：cd AppF > `go list ./...` # rainyApp, c21

[-](https://stackoverflow.com/questions/28166249/how-to-list-installed-go-packages) [-](http://www.techietown.info/2017/02/list-installed-packages-gogolang/)

