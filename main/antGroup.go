//存放蚁群结构体的定义和蚁群专属函数
package main

import (
	"fmt"
	"math"
	"math/rand"
)

type AntGroup struct{
	// 蚁群结构体
	a int  // α参数
	b int  // β参数
	p float64  // 信息素消散率
	count int  // 每轮结果相同次数计数
	pointNum int  // 点数目，包括起始点
	antNum int //蚁群数目
	maxRound int  // 最大循环轮数
	maxSameTime int  // 每轮结果最大相同次数


	bestPathSeats []int  //最佳遍历座位序列
	roundPathSeats []int  //  某一轮遍历的座位序列
	bestDistance float64  // 最佳距离
	roundDistance float64  // 某一轮产生的距离

	pheromoneList [][]float64  // 路径信息素二维数组
	SeatsList []*Seats  // 座位数组=资产数+1（加上起始点）
	seatDistanceList [][] float64  // 座位距离数组
	priorityList [][] int  // 二维优先级数组,可供后面按比例选择座位的时候改变,该表不包含起始点
	basicPriorityList [priorityNum][] int // 基础优先级数组，不可改变，该表不包含起始点
}


func newAntGroup(t *TSP, startX float64, startY float64) *AntGroup{
	// 利用一个TSP对象，即固定的座位状态，和一个起始点生成一个新的蚁群。好处是可以利用同一个TSP对象与不同的起始点生成许多
	// 不同的蚁群，即用不同的出发点生成路线。减少了重新构建距离表的时间。
	ag:=new(AntGroup)
	ag.a=1
	ag.b=2
	ag.p=0.5
	ag.count=0
	ag.antNum = 50
	ag.maxRound=1000
	ag.maxSameTime=20
	ag.basicPriorityList=t.priorityList
	// 复制一个新的priorityList，作为ag.priorityList的内容，以免修改此表的时候ag.basicPriorityList也被改变
	ag.priorityList =make([][]int,0)
	for _,v :=range ag.basicPriorityList{
		temList:=make([]int, len(v))
		copy(temList, v)
		ag.priorityList=append(ag.priorityList,temList)
	}

	//以下内容皆是将临时起始点加入到TSP问题座位列表和距离二维数组中
	SeatsList := make([]*Seats, len(t.SeatsList))
	copy(SeatsList, t.SeatsList)  // copy一个座位表，以免加入起始点的操作改变TSP中的座位表
	SeatsList = append(SeatsList, &Seats{len(SeatsList), startX, startY, 0})  // 起始点id为资产数目，优先级为0，起始点优先级无作用

	// 将起始点与其他点的距离加入，生成antGroup自己的距离表
	seatDistanceList := make([][] float64, len(t.SeatsList))
	copy(seatDistanceList, t.seatDistanceList)
	temList := make([]float64, 0)  // temList是新表的最后一行
	for _, v := range t.SeatsList {
		distance := math.Sqrt(math.Pow(v.x-startX, 2) + math.Pow(v.y-startY, 2))
		temList = append(temList, distance)
	}
	temList = append(temList, 0)  // 往最后一行末尾加上个0，作为起始点到起始点的距离
	seatDistanceList = append(seatDistanceList, temList)  // 将新表最后一行加入旧距离表生成新表
	for i := 0; i < len(t.SeatsList); i++ {
		// 将最后一行元素的内容写入新表对称的元素里去，生成对称二维数组
		seatDistanceList[i] = append(seatDistanceList[i], seatDistanceList[len(t.SeatsList)][i])
	}

	ag.seatDistanceList = seatDistanceList
	ag.SeatsList = SeatsList
	ag.pointNum=len(ag.SeatsList)
	return ag
}

func (ag *AntGroup)SeatExtraction(rateList []float64){
	// 座位提取函数，根据记录了每个优先级盘查比例的比率列表，将优先级表进行裁剪

	//初始化，复制基础优先级表内容给优先级表
	ag.priorityList =make([][]int,0)
	for _,v :=range ag.basicPriorityList{
		temList:=make([]int, len(v))
		copy(temList, v)
		ag.priorityList=append(ag.priorityList,temList)
	}

	for i:=0;i<len(rateList);i++{
		for j:=0;j<int((1.0-rateList[i])*float64(len(ag.basicPriorityList[i])));j++{
			ag.priorityList[i]=pop(ag.priorityList[i],rand.Intn(len(ag.priorityList[i])))
		}
	}
}

func (ag *AntGroup)run()(float64, []int){
	// 运行函数，返回最短路径长度，和路径经过的点，其中第一个点是起始点编号，即座位数目
	ag.roundDistance,ag.roundPathSeats = ag.firstFind()  // 先利用贪心算法进行首次查找
	ag.bestDistance,ag.bestPathSeats = ag.roundDistance,ag.roundPathSeats

	for i:=0;i<ag.maxRound;i++{
		var roundDistance float64
		roundDistance,ag.roundPathSeats = ag.findWay()  // 后续查找是以蚁群行动一次为单位
		// 利用蚁群找出的最短值，与上一轮结果进行比对，如果轮结果20次为改变则说明已收敛，结束蚁群算法
		if ag.roundDistance==roundDistance{
			ag.count++
			if ag.count==ag.maxSameTime{
				break
			}
		}else{
			ag.roundDistance=roundDistance
			ag.count=0
		}
		if ag.roundDistance < ag.bestDistance{
			ag.bestDistance,ag.bestPathSeats = ag.roundDistance,ag.roundPathSeats
		}
	}
	fmt.Printf("最短距离为：%f 最短路径为:", ag.bestDistance)
	fmt.Println(ag.bestPathSeats)
	return ag.bestDistance,ag.bestPathSeats
}


func (ag *AntGroup)firstFind()(float64,[]int){
	// 蚁群首次查找，使用贪心算法，下一个选择点是离现在点距离最短的点，查找完毕后初始化所有路径信息素浓度
	currentPoint := len(ag.SeatsList)-1  // 将初始点id设置为起始点
	distance:=0.0
	path:=[]int{currentPoint}  // 将初始点加入到路径中

	//复制一个优先级表，以供蚁群行走途中进行修改
	priorityList :=make([][]int,0)
	for _,v :=range ag.priorityList{
		temList:=make([]int, len(v))
		copy(temList, v)
		priorityList=append(priorityList,temList)
	}

	for i:=0;i<len(priorityList);i++ {
		for len(priorityList[i])>=1{  // 当此优先级中还有元素时，即此行长度不为0时持续查找下一个点
			length:=ag.seatDistanceList[currentPoint][priorityList[i][0]]
			point:=priorityList[i][0]
			column:=0
			for j:=1;j<len(priorityList[i]);j++{
				temLength:=ag.seatDistanceList[currentPoint][priorityList[i][j]]
				temPoint:=priorityList[i][j]
				if temLength<length{
					length=temLength
					point=temPoint
					column=j
				}
			}
			distance+=length
			path=append(path, point)
			priorityList[i]=pop(priorityList[i], column)  // 从优先级表去掉出已经路过的点
		}
	}
	initPheromone:=float64(ag.antNum)/distance  //每条路段信息素量为蚁群数目乘以首次寻找后得出路径的倒数
	for i:=0;i<ag.pointNum;i++{  // 初始化信息素表中的信息素
		temList:=make([]float64,0)
		for j:=0;j<ag.pointNum;j++{
			if i==j{
				temList=append(temList, 0)
			} else{
				temList=append(temList, initPheromone)
			}
		}
		ag.pheromoneList=append(ag.pheromoneList, temList)
	}
	return distance, path
}


func (ag *AntGroup)findWay()(float64,[]int){
	// 根据蚁群返回的距离列表和路径列表更新信息素和最佳路径
	distanceList:=make([]float64,0)
	pathList:=make([][]int,0)
	distanceList,pathList=ag.antFindWay(distanceList,pathList)
	//所有路径信息素首先挥发掉ag.p比例的部分
	for i:=0;i<len(ag.pheromoneList);i++{
		for j:=0;j<len(ag.pheromoneList);j++{
			ag.pheromoneList[i][j]*=1-ag.p
		}
	}

	bestDistance:=distanceList[0]
	bestPath:=pathList[0]
	for i:=0;i<len(pathList);i++{
		if distanceList[i]<bestDistance{  // 更新最短距离
			bestDistance=distanceList[i]
			bestPath=pathList[i]
		}
		currentPoint:=pathList[i][0]
		for j:=1;j<len(pathList[i]);j++{  // 每一条路径第一个点都是必定是起始点，所以j下标从1开始
			ag.pheromoneList[currentPoint][pathList[i][j]]+=1.0/distanceList[i]  // 往蚂蚁走过的路段上加上此蚂蚁走过的距离的倒数的信息素
			ag.pheromoneList[pathList[i][j]][currentPoint]+=1.0/distanceList[i]
			currentPoint=pathList[i][j]
		}
	}
	return bestDistance,bestPath
}


func (ag *AntGroup)antFindWay(distanceList []float64, pathList [][]int)([]float64, [][]int){
	// 蚁群找路的函数，返回蚁群的经过的所有路径的长度和经过的所有路径
	float64Ch:=make(chan float64 ,ag.antNum)
	pathCh:=make(chan []int,ag.antNum)
	for i:=0;i<ag.antNum;i++{
		go ag.antFunc(float64Ch,pathCh)
	}
	//生成距离列表和路径列表
	for i:=0;i<ag.antNum;i++{
		distance:=<-float64Ch
		temPath:=<-pathCh
		distanceList=append(distanceList, distance)
		pathList=append(pathList, temPath)
	}
	return  distanceList,pathList
}


func (ag *AntGroup)antFunc(float64Ch chan float64,pathCh chan []int){
	// 单个蚂蚁执行的函数，方便多线程调用
	priorityList :=make([][]int,0)                     //复制一个priorityList的副本
	for _,v :=range ag.priorityList{                   //
		temList:=make([]int, len(v))                   //
		copy(temList, v)                               //
		priorityList=append(priorityList,temList)      //
	}                                                  //
	currentPoint:=ag.pointNum-1
	temPath:=[]int{currentPoint}
	distance:=0.0
	for i:=0;i<len(priorityList);i++ {
		for len(priorityList[i])>=1{
			rouletteWeight:=make([]float64,len(priorityList[i]))
			for j:=0;j<len(priorityList[i]);j++{
				multiplier1:=math.Pow(ag.pheromoneList[currentPoint][priorityList[i][j]],float64(ag.a))
				multiplier2:=math.Pow(1.0/ag.seatDistanceList[currentPoint][priorityList[i][j]],float64(ag.b))
				weight:=multiplier1*multiplier2
				rouletteWeight[j]=weight  // 依据现在与可达点之间的距离和信息素更新轮盘选择权重列表，以得到下一个点
			}
			chosenPoint:=rouletteChoose(rouletteWeight)  // 返回i优先级的点中被选中点的下标，此点就作为下一个被选择的点
			temPath=append(temPath,priorityList[i][chosenPoint])
			distance+=ag.seatDistanceList[currentPoint][priorityList[i][chosenPoint]]
			currentPoint=priorityList[i][chosenPoint]
			priorityList[i]=pop(priorityList[i], chosenPoint)
		}
	}
	float64Ch<-distance  // 函数结束后将距离和路径送入通道内
	pathCh<-temPath
}
