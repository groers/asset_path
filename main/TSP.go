//存放了TSP结构体的定义和生成函数
package main


import "math"


type TSP struct {
	// tsp问题结构体，包括座位表，优先级表，座位距离表
	SeatsList []*Seats
	priorityList [priorityNum][] int  // 行号是优先级，每一行中的数据是每一个此优先级的座位的id号
	seatDistanceList [][] float64
}


func newTSP(SeatsList []*Seats) *TSP{
	//通过座位列表返回一个TSP对象，二维优先级列表，和二维座位距离表已算出
	length := len(SeatsList)
	t:=new(TSP)
	t.SeatsList=SeatsList
	for _,v := range SeatsList{
		// 生成二维优先级列表
		t.priorityList[v.priority]=append(t.priorityList[v.priority], v.id)
	}
	for i:=0;i<length;i++ {
		//确定二维座位距离表行数即座位数
		t.seatDistanceList = append(t.seatDistanceList, []float64{})
	}
	for i:=0;i<length;i++{
		//生成二维座位距离表，距离为欧几里得距离，即两点之间直线的距离
		t.seatDistanceList[i]=append(t.seatDistanceList[i],0)
		for j:=i+1;j<length;j++{
			distance:=math.Sqrt(math.Pow(SeatsList[i].x - SeatsList[j].x, 2) + math.Pow(SeatsList[i].y - SeatsList[j].y, 2))
			t.seatDistanceList[i]=append(t.seatDistanceList[i],distance)
			t.seatDistanceList[j]=append(t.seatDistanceList[j],distance)
		}
	}
	return t
}
