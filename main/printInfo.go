// 打印antGroup结构体中数据结构的函数集合
package main

import "fmt"

func (g *AntGroup) printSeats() {
	//输出座位表
	fmt.Println("座位表如下：")
	for _, v := range g.SeatsList {
		fmt.Println(*v)
	}
	fmt.Println()
}

func (g *AntGroup) printDistanceList() {
	//输出二维座位距离表
	fmt.Println("距离表如下：")
	for _, v := range g.seatDistanceList {
		fmt.Println(v)
	}
	fmt.Println()
}

func (g *AntGroup) printPriorityList() {
	//输出优先级表
	fmt.Println("优先级表如下：")
	for k, v := range g.priorityList {
		fmt.Printf("%d: ", k)
		fmt.Println(v)
	}
	fmt.Println()
}

func (g *AntGroup) printBasicPriorityList() {
	//输出基本优先级表
	fmt.Println("基本优先级表如下：")
	for k, v := range g.basicPriorityList {
		fmt.Printf("%d: ", k)
		fmt.Println(v)
	}
	fmt.Println()
}
