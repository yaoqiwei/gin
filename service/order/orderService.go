package order

import (
	"gin/model"
	"gin/model/body"
)

func OrderSuccess(p body.OrderNumParam) {
	model.OrderSuccess(p)
}
