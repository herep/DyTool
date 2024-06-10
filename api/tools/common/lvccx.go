package common

type PostOrderListParameter struct {
	Page                 string `json:"page" form:"page"`
	State                string `json:"state" form:"state"`
	OrderNumber          string `json:"order_number" form:"order_number"`
	ThirdNumber          string `json:"third_Number" form:"third_Number"`
	SubOrderNumber       string `json:"sub_order_number" form:"sub_order_number"`
	StationedUserName    string `json:"stationed_user_name" form:"stationed_user_name"`
	ReceiveName          string `json:"receive_name" form:"receive_name"`
	ReceivingPhone       string `json:"receiving_phone" form:"receiving_phone"`
	ItemNumber           string `json:"item_number" form:"item_number"`
	ProductName          string `json:"product_name" form:"product_name"`
	LogisticsCompanyName string `json:"logistics_company_name" form:"logistics_company_name"`
	LogisticsNumber      string `json:"logistics_number" form:"logistics_number"`
	BeginDate            string `json:"begin_date" form:"begin_date"`
	EndDate              string `json:"end_date" form:"end_date"`
	SendBeginDate        string `json:"send_begin_date" form:"send_begin_date"`
	SendEndDate          string `json:"send_end_date" form:"send_end_date"`
	Async                string `json:"async" form:"async"`
}
