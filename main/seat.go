//存放座位结构体的定义和产生座位数组的一个函数
package main

import "math/rand"


type Seats struct {
	// 资产座位结构体
	id int  // 座位id号
	x float64  // 座位x坐标
	y float64  // 资产y坐标
	priority int  // 资座位先级，0为最高
}

func getSeatsList(seatNum int)[]*Seats{
	// 通过随机数产生seatNum个座位，将它们指针做成座位表数组返回
	var x float64
	var y float64
	var priority int
	var t *Seats  // 单个资产指针
	var s []*Seats  // 座位表数组
	for  i:=0;i<seatNum;i++{
		x=float64(rand.Intn(100))+rand.Float64()  // 通过随机产生的int和float相加获得x，y值
		y=float64(rand.Intn(100))+rand.Float64()
		priority = rand.Intn(priorityNum)
		t=new(Seats)  // 给结构体指针分配地址
		t.x = x
		t.y = y
		t.id = i
		t.priority =priority
		s=append(s, t)
	}
	return s
}