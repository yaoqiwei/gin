package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

type NameData struct {
	Name string
	Aval string
}

func main() {
	// logConfig := config.Log()
	// DatabaseConfig := config.Database()
	// RedisConfig := config.Redis()

	// log.Init(logConfig)
	// model.Init(DatabaseConfig)
	// redis.Init(RedisConfig)

	// routes.InitRouter()

	sumMap := make(map[string]float64, 0)
	countMap := make(map[string]float64, 0)

	file, _ := os.Open("/Volumes/yqw/xm/gin/ddg_predictions.out")

	defer file.Close()
	b, _ := ioutil.ReadAll(file)
	content := string(b)
	list := make([]*NameData, 0)
	for _, lineStr := range strings.Split(content, "\n") {
		if lineStr != "" {
			arr := strings.Fields(strings.TrimSpace(lineStr))
			if len(arr) == 0 {
				continue
			}
			if data, ok := sumMap[arr[1]]; ok {
				lineFloat, _ := strconv.ParseFloat(arr[2], 64)
				sumMap[arr[1]] = data + lineFloat
			} else {
				sumMap[arr[1]] = data
			}

			if data, ok := countMap[arr[1]]; ok {
				countMap[arr[1]] = data + float64(1)
			} else {
				countMap[arr[1]] = float64(1)
			}
		}
	}
	for v, i := range sumMap {
		count, _ := countMap[v]
		list = append(list, &NameData{
			Name: v,
			Aval: fmt.Sprintf("%.3f", i/count),
		})
	}

	Export(list)
	// fmt.Println(list)
}

// HeaderColumn 表头字段定义
type HeaderColumn struct {
	Field string // 字段，数据映射到的数据字段名
	Title string // 标题，表格中的列名称
}

func Export(data []*NameData) {
	file := xlsx.NewFile()                      // NewWriter 创建一个Excel写操作实例
	sheet, err := file.AddSheet("student_list") //表实例
	if err != nil {
		fmt.Printf(err.Error())
	}

	headers := []*HeaderColumn{
		{Field: "Name", Title: "名称"},
		{Field: "Aval", Title: "平均值"},
	}
	style := map[string]float64{
		"Name": 2.0,
		"Aval": 2.0,
	}
	sheet, _ = SetHeader(sheet, headers, style)

	for _, stu := range data {
		data := make(map[string]string)
		data["Name"] = stu.Name
		data["Aval"] = stu.Aval

		row := sheet.AddRow()
		row.SetHeightCM(0.8)
		for _, field := range headers {
			row.AddCell().Value = data[field.Field]
		}
	}
	outFile := "/Volumes/yqw/xm/gin/test.xlsx"
	err = file.Save(outFile)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println("\n\nexport success")
}

func SetHeader(sheet *xlsx.Sheet, header []*HeaderColumn, width map[string]float64) (*xlsx.Sheet, error) {
	if len(header) == 0 {
		return nil, errors.New("Excel.SetHeader 错误: 表头不能为空")
	}

	// 表头样式
	style := xlsx.NewStyle()

	font := xlsx.DefaultFont()
	font.Bold = true

	alignment := xlsx.DefaultAlignment()
	alignment.Vertical = "center"

	style.Font = *font
	style.Alignment = *alignment

	style.ApplyFont = true
	style.ApplyAlignment = true

	// 设置表头字段
	row := sheet.AddRow()
	row.SetHeightCM(1.0)
	row_w := make([]string, 0)
	for _, column := range header {
		row_w = append(row_w, column.Field)
		cell := row.AddCell()
		cell.Value = column.Title
		cell.SetStyle(style) //设置单元样式
	}

	// 表格列，宽度
	if len(row_w) > 0 {
		for k, v := range row_w {
			if width[v] > 0.0 {
				sheet.SetColWidth(k, k, width[v]*10)
			}
		}
	}

	return sheet, nil
}
