package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
	"strings"
)

type Data struct {
	No         string `json:"no"`
	PrintDate  string `json:"printDate"`
	AddAt      string `json:"addAt"`
	ClickCount int    `json:"clickCount"`
	Color      string `json:"color"`
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
	var titleMap = make(map[string][3]string)
	var titleLine = rows[0]
	for j, title := range titleLine {
		if title == "时间戳记" || title == "生产日期" {
			continue
		}
		tempArr := GetMacColorAndSize(title)
		if i := titleMap[tempArr[0]]; i[0] == "" {
			titleMap[strconv.Itoa(j)] = tempArr
		}
	}
	fmt.Print(titleMap)

	for i, row := range rows {
		if i == 0 {
			continue
		}
		var data Data
		for j, colCell := range row {
			// 排除第一列为Null
			if j == 0 || colCell == "Null" {
				continue
			}
			tempMacMetaInfo := titleMap[strconv.Itoa(j)]
			data.AddAt = "2020-10-27"
			data.ClickCount, _ = strconv.Atoi(colCell)
			data.Color = tempMacMetaInfo[1]
			data.No = tempMacMetaInfo[0]
			datas = append(datas, data)
		}
	}
	return datas, nil
}

func GetMacColorAndSize(str string) [3]string {

	color := "black"
	size := "A4"
	if find := strings.Contains(str, "彩色"); find {
		color = "color"
	}
	if find := strings.Contains(str, "彩印"); find {
		color = "color"
	}
	if find := strings.Contains(str, "A3"); find {
		size = "A3"
	}

	return [3]string{str, color, size}
}

func main() {
	path := "/Users/yanghaha/goproj/src/go-mega/test/11.xlsx"
	arr, err := ReadExcel(path)
	if err != nil {
		fmt.Print(err)
	}
	for _, val := range arr {
		fmt.Print(val)
	}
}
