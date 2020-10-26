package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

type Data struct {
	FirstClass  string `json:"first_class"`  // 一级
	SecondClass string `json:"second_class"` // 二级
	ThirdClass  string `json:"third_class"`  //三级
}

func ReadExcel(filepath string) ([]Data, error) {
	// 首先读excel
	xlsx, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rows := xlsx.GetRows("Daily Record")
	var datas []Data
	for i, row := range rows {
		// 去掉第一行，第一行是表头
		if i == 0 {
			continue
		}
		var data Data
		for j, colCell := range row {
			// 排除第一列为Null
			if j == 0 && colCell == "Null" {
				continue
			}
			// 第一列即是一级
			if j == 0 && colCell != "Null" {
				data.FirstClass = colCell
			}
			// 第二列即是二级
			if j == 1 {
				data.SecondClass = colCell
			}
			// 三级
			if j == 2 {
				data.ThirdClass = colCell
			}

		}
		datas = append(datas, data)
	}
	return datas, nil
}

func main() {
	path := "/Users/yanghaha/WebstormProjects/badge/equ/csv_data/01北京作业中心/11.xlsx"
	arr, err := ReadExcel(path)
	if err != nil {
		fmt.Print(err)
	}
	for _, val := range arr {
		fmt.Print(val)
	}
}
