# Golang Technology Learning

## To start developing Golang Project

If you want to build this project right away here is an option:

- ### You have a working [Go environment](https://golang.org/doc/install).

```shell
mkdir -p $GOPATH/src/github.com
cd $GOPATH/src/github.com
git clone git@github.com:landy8530/Golang-Course.git
cd Golang-Course
```

## Golang现状

- Go语言的作者很多是Google或者Unix的作者，之前是写C的
- Go是编译型语言，更快
- 为现代而生的Go语言
  - 支持交叉编译，编译快速
  - 语法简洁（只有25个关键字）
  - 代码风格统一（自带格式化工具）
  - 开发效率高
  - 执行性能好
  - 天生支持并发

> Go于2009年发布，当时多核处理器已经上市。Go语言在多核并发上拥有原生的设计优势，Go语言从底层原生支持并发，无须第三方库、开发者的编程技巧和开发经验。
>
> Go语言的并发是基于 `goroutine` 的，`goroutine` 类似于线程，但并非线程。可以将 `goroutine` 理解为一种虚拟线程。Go 语言运行时会参与调度 `goroutine`，并将 `goroutine` 合理地分配到每个 CPU 中，最大限度地使用CPU性能。开启一个`goroutine`的消耗非常小（大约2KB的内存），你可以轻松创建数百万个`goroutine`。
>
> `goroutine`的特点：
>
> 1. `goroutine`具有可增长的分段堆栈。这意味着它们只在需要时才会使用更多内存。
> 2. `goroutine`的启动时间比线程快。
> 3. `goroutine`原生支持利用channel安全地进行通信。
> 4. `goroutine`共享数据结构时无需使用互斥锁。

## Go 性能强悍

> 与其他现代高级语言（如Java/Python）相比，使用C，C++的最大好处是它们的性能。因为C/ C++是编译型语言而不是解释的语言。 处理器只能理解二进制文件，Java和Python这种高级语言在运行的时候需要先将人类可读的代码翻译成字节码，然后由专门的解释器再转变成处理器可以理解的二进制文件。

![image-20211012202454791](/Users/landyl/Learning/Golang/src/github.com/images/image-20211012202454791.png)

