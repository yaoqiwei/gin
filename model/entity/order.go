package entity

import "time"

type Order struct {
	ID        int64     `json:"id"`
	OrderId   string    `json:"order_id"`
	OrderType byte      `json:"order_type"` // 1钻石2VIP
	Status    byte      `json:"status"`     // 状态 0 未付款 1已付款 2未发货 3已发货 4交易成功 5交易关闭
	Uid       int64     `json:"uid"`
	HandleId  int64     `json:"handle_id"` // 操作人id
	TotalFee  float64   `json:"total_fee"`
	PayType   byte      `json:"pay_type"`   // 支付类型 2微信 1支付宝 3银行卡 4人工充值
	PayTime   time.Time `json:"pay_time"`   // 支付时间
	EndTime   time.Time `json:"end_time"`   // 交易完成时间
	CloseTime time.Time `json:"close_time"` // 订单关闭时间
	PaymentId int64     `json:"payment_id"`
	Channel   string    `json:"channel"` // 通道
	GoodsId   int64     `json:"goods_id"`
	CreateAt  time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
	VipType   byte      `json:"vip_type"` // VIP类型，0无1首充2续费
}
