
https://golang.org/doc/install

https://gobyexample.com/

G learn golang the hard way <br>
https://github.com/gophergala/learn-Go-the-hard-way <br>
https://andy-zhangtao.gitbooks.io/golang/content/#静态链接库 <br>

[https://tour.golang.org/](https://tour.golang.org/list) (golang array slice)

https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/preface.md
https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/01.2.md#获取远程包

[Go 语言简介 - CoolShell](https://coolshell.cn/articles/8460.html)
[Go语言的修饰器](https://coolshell.cn/articles/17929.html)

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

(packages, install, go get; 如果想 include 一个文件(a package)，那么文件存放位置: vendor 的优先级 [高](#vendorDirCanShadowSrcPackages)[于](https://golang.org/cmd/go/#hdr-Vendor_Directories) src_packages) [-](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies)

include by vendor: <br>
/src/rainyApp/test1.go <br>
/src/rainyApp/vendor/dasMath/1.go <br><br>

include by packages: <br>
/src/rainyApp/test1.go <br>
/src/myMath/test1.go <br><br>

= = = = = =

list packages:

(无用) `go list` 的作用：cd AppF/src/rainyApp > `go list` # rainyApp ； cd AppF/src/c21 > `go list` # c21

`go list ./...` 的作用：列出所有packages 在 workspace(src+bin+pkg)Dir 里的：cd AppF > `go list ./...` # rainyApp, c21

[-](https://stackoverflow.com/questions/28166249/how-to-list-installed-go-packages) [-](http://www.techietown.info/2017/02/list-installed-packages-gogolang/)

= ==  = =

print
https://golang.org/pkg/fmt/#hdr-Printing
https://golang.org/pkg/fmt/#Println
https://www.reddit.com/r/golang/comments/6xizkf/why_would_you_want_to_use_printf_over_print_or/

