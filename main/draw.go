// 绘制结果图相关函数
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

//func drawLine(x0, y0, x1, y1 float64, brush func(x int,y int)) {
//	if x0>x1{
//		x0,x1=x1,x0
//		y0,y1=y1,y0
//	}
//	var a float64 = (y1-y0)/(x1-x0) // 斜率
//	var b float64 = (x1*y0-x0*y1)/(x1-x0) // 在y轴截距
//	i:=0.0
//	for i=x0;i<=x1;i++{
//		brush(int(i),int(a*i+b))
//	}
//	fmt.Println(i)
//}

func abs(x int) int {
	// 求绝对值
	if x >= 0 {
		return x
	}
	return -x
}

func drawLine(x0, y0, x1, y1 int, brush func(x int, y int)) {
	// 划线函数
	dx := abs(x1 - x0)
	dy := abs(y1 - y0)
	sx, sy := 1, 1
	if x0 >= x1 {
		sx = -1
	}
	if y0 >= y1 {
		sy = -1
	}
	err := dx - dy

	for {
		brush(x0, y0)
		if x0 == x1 && y0 == y1 {
			return
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

func (ag *AntGroup) draw(outPath string) {
	const (
		dx = 120
		dy = 120
	)
	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	imgFile, _ := os.Create(outPath)
	currentPoint := ag.bestPathSeats[0]
	img.Set(int(ag.SeatsList[currentPoint].x), int(ag.SeatsList[currentPoint].y), color.RGBA64{255, 0, 0, 255})
	for i := 1; i < len(ag.bestPathSeats); i++ {
		nextPoint := ag.bestPathSeats[i]
		drawLine(int(ag.SeatsList[currentPoint].x), int(ag.SeatsList[currentPoint].y), int(ag.SeatsList[nextPoint].x), int(ag.SeatsList[nextPoint].y), func(x, y int) {
			img.Set(x, y, color.Black)
		})
		currentPoint = nextPoint
	}
	err := png.Encode(imgFile, img)
	if err != nil {
		fmt.Println(err)
	}
}
