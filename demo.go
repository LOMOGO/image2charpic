//这个小程序的功能是将文字图片化
package main

import (
	"bytes"
	"fmt"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/vgimg"
	"image"
	"image/color"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func main() {
	//测试程序用时
	t1 := time.Now()

	//获取命令行输入的参数
	source := os.Args[1]//输入图片地址
	target := os.Args[2]//输出图片地址
	text := os.Args[3]//将文字图片化所需要的文字
	size := os.Args[4]//文字上大小

	f1, _ := ioutil.ReadFile(source) //读取文件
	buffer := bytes.NewBuffer(f1)
	img, _, _ := image.Decode(buffer)
	//获取原图片尺寸
	bounds := img.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()

	fontsize, _ := strconv.Atoi(size)
	vgfont, _ := vg.MakeFont("Courier", vg.Length(fontsize))
	//文字的宽度和高度
	fontWidth := vgfont.Width(text)
	fontHeight := vgfont.Extents().Height
	//创建vg.Canvas
	c := vgimg.NewWith(
		vgimg.UseWH(vg.Length(dx)-fontWidth, vg.Length(dy)),
		vgimg.UseBackgroundColor(color.RGBA{}),
	)
	for x := vg.Length(dx); x >= vg.Length(0) ; x -= fontWidth {
		for y := vg.Length(dy); y >= vg.Length(0); y -= fontHeight*0.585 {
			//获取(x, (dy + 1) - y)点的RGBA信息，之所以y坐标这样表示是因为img.At默认的坐标起始点是左上角，而vgimg的SetColor方法的默认坐标起始点是左下角，如果不转换，输出的图片将会倒立
			colorRGBA := img.At(int(x), (dy + 1) - int(y))
			r, g, b, a := colorRGBA.RGBA()
			rUint8 := uint8(r >> 8)
			gUint8 := uint8(g >> 8)
			bUint8 := uint8(b >> 8)
			aUint8 := uint8(a >> 8)
			//将文字打印在图片上
			c.FillString(vgfont, vg.Point{X: x - fontWidth, Y: y}, text)
			//设置像素点颜色
			c.SetColor(color.RGBA{R: rUint8, G: gUint8, B: bUint8, A: aUint8})
		}
	}
	nc := vgimg.PngCanvas{
		Canvas: c,
	}

	f,_ := os.Create(target)
	_, _ = nc.WriteTo(f)
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)
}
