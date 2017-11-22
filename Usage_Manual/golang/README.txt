
[golang install](https://golang.org/doc/install)

1.
echo $PATH
/usr/lib64/qt-3.3/bin:/usr/local/bin:/usr/local/sbin:/usr/bin:/usr/sbin:/bin:/sbin

2.
下载
[下载](https://golang.org/dl/)

3.1
标准安装
标准位置
/usr/local/go 文件夹

3.2
自定义安装
自定义位置
~/local/usr/local/go 文件夹
解压：
把 go$VERSION.$OS-$ARCH.tar.gz 放到 ~/local/usr/local 里，然后执行命令 
> tar -C . -xzf go1.9.2.linux-amd64.tar.gz 即可
自定义path
nano ~/local/dotfiles/.bashrc
```
export GOROOT=$HOME/local/usr/local/go
export GOPATH=$HOME
export PATH=$PATH:$GOROOT/bin
```
> go version
# go version go1.9.2 linux/amd64

2.
> go run hello-world.go
https://gobyexample.com/hello-world

3.


