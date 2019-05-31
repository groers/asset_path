package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func init() {
	// 初始化函数，在main执行前自动执行
	rand.Seed(time.Now().Unix()) // 初始化随机数seed
}

func main() {
	startTime := time.Now().UnixNano()

	//list:=[]float64{0.6,0.5,0.3,0.3,0.1}

	//seatList:=getSeatsList(10)
	//parseSeats(seatList,"asset_table_10.csv")
	seatList := importSeats("asset_table_20.csv")

	g := newAntGroup(newTSP(seatList), 64, 3.83)
	//g.SeatExtraction(list)

	g.printBasicPriorityList()
	g.printPriorityList()
	g.printSeats()

	g.run()
	g.draw("path.png")
	fmt.Printf("耗时：%fs\n", float64(time.Now().UnixNano()-startTime)/math.Pow(10, 9))
}
