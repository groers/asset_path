//存放了csv相关操作的函数*
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func parseSeats(seatsList []*Seats, outPath string){
	//解析座位表，将座位表写入csv文件中保存
	file, _ := os.OpenFile(outPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	w := csv.NewWriter(file)
	for _,v := range seatsList{
		id:=strconv.Itoa(v.id)
		x:= strconv.FormatFloat(v.x, 'f', -1, 64)
		y:= strconv.FormatFloat(v.y, 'f', -1, 64)
		priority:=strconv.Itoa(v.priority)
		stringList:=[]string{id,x,y,priority}
		err:=w.Write(stringList)
		if err != nil{
			fmt.Println("读文件错误",outPath)
			os.Exit(-1)
		}
	}
	w.Flush()
	err:=file.Close()
	if err != nil{
		fmt.Println("关闭文件错误",outPath)
		os.Exit(-1)
	}
}


func importSeats(inPath string)[]*Seats{
	// 从csv文件中引入座位表
	seatsList:=make([]*Seats,0)
	file, err := os.Open(inPath)
	if err!=nil{
		fmt.Println("打开文件错误",inPath)
		os.Exit(-1)
	}
	r := csv.NewReader(file)
	out,_ := r.ReadAll()
	for _, v:= range out {
		id,_:=strconv.Atoi(v[0])
		x,_:=strconv.ParseFloat(v[1], 64)
		y,_:=strconv.ParseFloat(v[2], 64)
		priority,_:=strconv.Atoi(v[3])
		s:=&Seats{id,x,y,priority}
		seatsList=append(seatsList,s)
	}
	err=file.Close()
	if err != nil{
		fmt.Println("关闭文件错误",inPath)
		os.Exit(-1)
	}
	return seatsList
}