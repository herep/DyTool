package lvccx

import (
	"bytes"
)

//LvOrderUrl 吕橙下单订单接口
const (
	LvOrderUrl = "https://www.lvccx.com/user/order/getorderlist.html"
)

//LvOrderService 下单列表接口
type LvOrderService interface {
	GetOrderUrl(page, state, orderNumber, thirdNumber, subOrderNumber, stationedUserName,
		receiveName, receivingPhone, itemNumber, productName, logisticsCompanyName, logisticsNumber, beginDate, endDate,
		sendBeginDate, sendEndDate, logToken, async string) string
}

type LvOrderServiceImpl struct{}

func NewLvOrderServiceImpl() LvOrderService {
	return &LvOrderServiceImpl{}
}

//GetOrderUrl 拼接订单接口
func (is *LvOrderServiceImpl) GetOrderUrl(page, state, orderNumber, thirdNumber, subOrderNumber, stationedUserName,
	receiveName, receivingPhone, itemNumber, productName, logisticsCompanyName, logisticsNumber, beginDate, endDate,
	sendBeginDate, sendEndDate, logToken, async string) string {

	var buff bytes.Buffer
	buff.WriteString(LvOrderUrl)
	buff.WriteString("?page=")
	buff.WriteString(page)
	buff.WriteString("&state=")
	buff.WriteString(state)
	buff.WriteString("&order_number=")
	buff.WriteString(orderNumber)
	buff.WriteString("&third_number=")
	buff.WriteString(thirdNumber)
	buff.WriteString("&sub_order_number=")
	buff.WriteString(subOrderNumber)
	buff.WriteString("&stationed_user_name=")
	buff.WriteString(stationedUserName)
	buff.WriteString("&receive_name=")
	buff.WriteString(receiveName)
	buff.WriteString("&receiving_phone=")
	buff.WriteString(receivingPhone)
	buff.WriteString("&itemnumber=")
	buff.WriteString(itemNumber)
	buff.WriteString("&product_name=")
	buff.WriteString(productName)
	buff.WriteString("&logistics_company_name=")
	buff.WriteString(logisticsCompanyName)
	buff.WriteString("&logistics_number=")
	buff.WriteString(logisticsNumber)
	buff.WriteString("&begin_date=")
	buff.WriteString(beginDate)
	buff.WriteString("&end_date=")
	buff.WriteString(endDate)
	buff.WriteString("&send_begin_date=")
	buff.WriteString(sendBeginDate)
	buff.WriteString("&send_end_date=")
	buff.WriteString(sendEndDate)
	buff.WriteString("&logtoken=")
	buff.WriteString(logToken)
	buff.WriteString("&async=")
	buff.WriteString(async)

	return buff.String()
}
