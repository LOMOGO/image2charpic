# image2charpic
这个golang命令行工具实现了将图片文字化的功能，目前只支持png格式图片的输入输出

如何使用：(在终端中运行以下语句)

```
go build -o Image2CharPic main.go
./Image2CharPic conv -h
图片转文字

Usage:
   conv [flags]

Flags:
  -h, --help            help for conv
      --size int8       文字的大小 (default 7)
      --source string   原图片的地址
      --target string   字符图片的输出地址
      --text string     图片文字化的文字(目前只支持英文、数字) (default "hello")

```
后续版本留坑：

1. 引入并发以优化图片生成的速度
2. 控制生成图片的大小在一个合适的范围

~~更加完善的命令行工具~~

3. 支持更多的图片格式
