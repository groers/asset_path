//辅助文件，存放了常数值的定义和一些辅助函数

package main


import "math/rand"


const priorityNum = 20  // 优先级数目，从0到19


func pop(s []int, index int) []int{
	// 弹出int型切片s中下标为index的元素，并返回新int切片
	copyS := s[:index]
	length:=len(s)
	for i:=index+1;i<length;i++ {
		copyS=append(copyS,s[i])
	}
	return copyS
}


func rouletteChoose(rouletteWeight []float64)int{
	// 轮盘赌法，给定一个权重列表，然后随机给出被选择的元素的下标
	sum:=0.0
	for _,v := range rouletteWeight{
		sum+=v
	}
	randNum:=rand.Float64()
	point:=0.0
	for i:=0;i<len(rouletteWeight);i++{
		point+=rouletteWeight[i]/sum
		if point>randNum{
			return i
		}
	}
	return -1
}