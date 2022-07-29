package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type TestData struct {
	Name string `json:"name"`
}

func TestGet(c *gin.Context) {
	fmt.Println("aaa")
	// p := TestData{}
	// if err := c.ShouldBind(&p); err != nil {
	// 	context.ValidateError(c)
	// 	return
	// }
	// comma := fmt.Sprintf("%c", p.Name[0])
	// arr := []string{"A", "C", "D", "E", "F", "G", "H", "I", "K", "L", "M", "N", "P", "Q", "R", "S", "T", "V", "W", "Y"}
	// res := ""
	// for _, v := range arr {
	// 	if v == comma {
	// 		continue
	// 	}
	// 	res = p.Name + v
	// 	fmt.Printf("%s\n\n", res)
	// }

	// context.Success(c, nil, nil)
}
