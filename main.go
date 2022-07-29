package main

import (
	"gin/config"
	"gin/log"
	"gin/model"
	"gin/redis"
	"gin/routes"
)

func main() {

	routes.InitRouter()

	logConfig := config.Log()
	DatabaseConfig := config.Database()
	RedisConfig := config.Redis()

	// backup.BackupRedis()
	// excelConfig := config.Excel()
	log.Init(logConfig)
	model.Init(DatabaseConfig)
	redis.Init(RedisConfig)
	// excel.Init(excelConfig)

	// f, err := os.Open("output_path.txt")
	// if err != nil {
	// 	fmt.Println("err:", err)
	// 	return
	// }
	// defer f.Close()

	// content, err := readTxt(f)
	// if err != nil {
	// 	fmt.Println("err:", err)
	// 	return
	// }
	// // fmt.Println("content:", content)
	// outText(content)
}

// func readTxt(r io.Reader) ([]string, error) {
// 	reader := bufio.NewReader(r)
// 	l := make([]string, 0, 64)
// 	// 按行读取
// 	for {
// 		line, _, err := reader.ReadLine()
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			} else {
// 				return nil, err
// 			}
// 		}
// 		// l = append(l, strings.Trim(string(line), " "))
// 		l = append(l, string(line))
// 	}

// 	return l, nil
// }

// func outText(context []string) {
// 	fileName := "out.txt"
// 	dstFile, err := os.Create(fileName)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	defer dstFile.Close()
// 	for i := 0; i < len(context); i++ {
// 		// if len(context[i]) > 0 {
// 		// 	Z := context[i][:1] + "A" + context[i][1:] + ";"
// 		dstFile.WriteString(context[i] + "\n1\n")
// 		// }
// 	}
// }
