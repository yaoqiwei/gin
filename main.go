package main

func main() {
	// config.HttpConf()
	// config.Database()
	// config.Redis()
	// config.RabbitMQ()
	// config.ModeConf()

	// lib.Init()
	// rabbitmq.InitRabbitMQ()
	// routes.HttpServerRun()

	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, os.Interrupt)
	// <-quit

	// aaa := TreeNode{
	// 	Val: 3,
	// 	Left: &TreeNode{
	// 		Val: 9,
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 20,
	// 		Left: &TreeNode{
	// 			Val: 15,
	// 		},
	// 		Right: &TreeNode{
	// 			Val: 7,
	// 		},
	// 	},
	// }

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertTemperature(celsius float64) []float64 {

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
