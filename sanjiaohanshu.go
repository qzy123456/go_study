package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"log"
)
func main()  {
	const size  =200

	pic := image.NewGray(image.Rect(0,0,size,size))

	for x:=0; x<size;x++  {
		for y:=0;y<size;y++ {
			pic.SetGray(x,y,color.Gray{255})
		}
	}
	// 从0到最大像素生成x坐标
	for x := 0; x < size; x++ {
		// 让sin的值的范围在0~2Pi之间
		s := float64(x) * 2 * math.Pi / size
		// sin的幅度为一半的像素。向下偏移一半像素并翻转
		y := size/2 - math.Sin(s)*size/2
		// 用黑色绘制sin轨迹
		pic.SetGray(x, int(y), color.Gray{0})
	}
	// 创建文件
	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}
	// 使用png格式将数据写入文件
	png.Encode(file, pic) //将image信息写入文件中
	// 关闭文件
	file.Close()
}