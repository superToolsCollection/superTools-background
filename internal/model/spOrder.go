package model

/**
* @Author: super
* @Date: 2021-01-22 20:30
* @Description: 订单表
**/

type SpOrder struct {
	TradeNo            string  `gorm:"column:trade_no" json:"trade_no"`
	OrderFapiaoTitle   string  `gorm:"column:order_fapiao_title" json:"order_fapiao_title"`
	IsSend             string  `gorm:"column:is_send" json:"is_send"`
	OrderFapiaoCompany string  `gorm:"column:order_fapiao_company" json:"order_fapiao_company"`
	OrderFapiaoContent string  `gorm:"column:order_fapiao_content" json:"order_fapiao_content"`
	// 订单状态： 0未付款、1已付款
	PayStatus          int  `gorm:"column:pay_status" json:"pay_status"`
	OrderNumber        string  `gorm:"column:order_number" json:"order_number"`
	OrderPrice         float64 `gorm:"column:order_price" json:"order_price"`
	// 支付方式  0未支付 1支付宝  2微信  3银行卡
	OrderPay           int  `gorm:"column:order_pay" json:"order_pay"`
	OrderID            int     `gorm:"column:order_id;primary_key" json:"order_id"`
	UserID             int     `gorm:"column:user_id" json:"user_id"`
	CreateTime         int64   `gorm:"column:create_time" json:"create_time"`
	ConsigneeAddr      string  `gorm:"column:consignee_addr" json:"consignee_addr"`
	UpdateTime         int64   `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (s *SpOrder) TableName() string {
	return "sp_order"
}
