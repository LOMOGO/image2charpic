# image2charpic
这个golang小程序实现了将文字图片化的功能，目前只支持png格式图片的输入输出

如何使用：(在终端中运行以下语句)

```
go build demo.go
./demo 带转换图片.png 输出的图片.png 要转换的文字 字体大小(./demo 梁家辉.png test.png LJH 10)
```
后续版本留坑：

1. 引入并发以优化图片生成的速度
2. 控制图片的大小在一个合适的范围
3. 更加完善的命令行工具
