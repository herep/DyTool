package lvccx

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	lvccx2 "tbTool/api/service/lvccx"
	"tbTool/api/tools/common"
	"tbTool/api/tools/lvccx"
	"tbTool/pkg"
	"tbTool/pkg/request"
	"time"
)

const LvApiType = "TB_ORDER_GUIDE=1"

type OrderList struct {
	Ret  string `json:"ret"`
	Data struct {
		Data struct {
			LogReferer string `json:"log_referer"`
			SsoDomain  string `json:"sso_domain"`
			Mtime      string `json:"mtime"`
			LockScreen int    `json:"lock_screen"`
			SessionID  string `json:"session_id"`
			AfsCheck   int    `json:"afs_check"`
			Loginout   string `json:"loginout"`
			State      int    `json:"state"`
			OrderList  []struct {
				CreateDate               string  `json:"create_date"`
				ID                       int     `json:"id"`
				OrderNumber              string  `json:"order_number"`
				OrderType                string  `json:"order_type"`
				OrderState               int     `json:"order_state"`
				ErpOrderState            int     `json:"erp_order_state"`
				ReceivingName            string  `json:"receiving_name"`
				ReceivingAddress         string  `json:"receiving_address"`
				ReceivingPhone           string  `json:"receiving_phone"`
				LogisticsCompanyCode     string  `json:"logistics_company_code"`
				LogisticsCompanyName     string  `json:"logistics_company_name,omitempty"`
				LogisticsNumber          string  `json:"logistics_number,omitempty"`
				LogisticsMoney           int     `json:"logistics_money"`
				DfMoney                  float64 `json:"df_money"`
				ServiceMoney             int     `json:"service_money"`
				PaymentState             int     `json:"payment_state"`
				CheckState               int     `json:"check_state"`
				RefundState              int     `json:"refund_state"`
				PaymentMoney             float64 `json:"payment_money"`
				ProductMoney             int     `json:"product_money"`
				ProductCount             int     `json:"product_count"`
				DiscountMoney            int     `json:"discount_money"`
				UserPayState             string  `json:"user_pay_state"`
				WarehouseID              int     `json:"warehouse_id"`
				WarehouseName            string  `json:"warehouse_name"`
				SendState                int     `json:"send_state"`
				SendProductCount         int     `json:"send_product_count"`
				ReturnState              string  `json:"return_state"`
				IsCanRefund              int     `json:"is_can_refund"`
				IsPreOrderCreate         string  `json:"is_pre_order_create"`
				EarnestMoney             int     `json:"earnest_money"`
				Mark                     int     `json:"mark"`
				MarkNote                 string  `json:"mark_note,omitempty"`
				SampleType               int     `json:"sample_type"`
				ActivityType             int     `json:"activity_type"`
				SurplusMoney             int     `json:"surplus_money"`
				IsOutstockOrder          int     `json:"is_outstock_order"`
				RestMoney                int     `json:"rest_money"`
				OrderEarnestMoney        int     `json:"order_earnest_money"`
				IsEarnestPay             int     `json:"is_earnest_pay"`
				PayType                  int     `json:"pay_type"`
				IsCancel                 int     `json:"is_cancel"`
				RefundOrder              int     `json:"refund_order"`
				OrderStateName           string  `json:"order_state_name"`
				OrderStateColor          string  `json:"order_state_color"`
				PlatformSubsidy          int     `json:"platform_subsidy"`
				SysjNew                  int     `json:"sysj_new"`
				IsReturn                 int     `json:"is_return"`
				ConfirmCountdown         int     `json:"confirm_countdown"`
				RefundProductCount       int     `json:"refund_product_count"`
				TkCouponDiscountMoney    int     `json:"tk_coupon_discount_money"`
				TkCouponID               int     `json:"tk_coupon_id"`
				RefundingProductCount    int     `json:"refunding_product_count"`
				ShopName                 string  `json:"shop_name"`
				StationedUserID          int     `json:"stationed_user_id"`
				IsSellerMessage          int     `json:"is_seller_message"`
				ReturningProductCount    int     `json:"returning_product_count"`
				OrderProductServiceMoney int     `json:"order_product_service_money"`
				CloseType                int     `json:"close_type"`
				ExpressInfo              struct {
					ExpressNumber  string `json:"express_number"`
					SendDate       string `json:"send_date"`
					ReceivingName  string `json:"receiving_name"`
					ReceivingPhone string `json:"receiving_phone"`
					Deliverystatus int    `json:"deliverystatus"`
					State          int    `json:"state"`
					Error          string `json:"error"`
					ExpressMessage string `json:"express_message"`
					ExpressDate    string `json:"express_date"`
				} `json:"express_info,omitempty"`
				Rn         int    `json:"rn"`
				CancelDate string `json:"cancel_date,omitempty"`
			} `json:"orderList"`
		} `json:"data"`
		Page struct {
			OrderList []struct {
				CreateDate               string  `json:"create_date"`
				ID                       int     `json:"id"`
				OrderNumber              string  `json:"order_number"`
				OrderType                string  `json:"order_type"`
				OrderState               int     `json:"order_state"`
				ErpOrderState            int     `json:"erp_order_state"`
				ReceivingName            string  `json:"receiving_name"`
				ReceivingAddress         string  `json:"receiving_address"`
				ReceivingPhone           string  `json:"receiving_phone"`
				LogisticsCompanyCode     string  `json:"logistics_company_code"`
				LogisticsCompanyName     string  `json:"logistics_company_name,omitempty"`
				LogisticsNumber          string  `json:"logistics_number,omitempty"`
				LogisticsMoney           int     `json:"logistics_money"`
				DfMoney                  float64 `json:"df_money"`
				ServiceMoney             int     `json:"service_money"`
				PaymentState             int     `json:"payment_state"`
				CheckState               int     `json:"check_state"`
				RefundState              int     `json:"refund_state"`
				PaymentMoney             float64 `json:"payment_money"`
				ProductMoney             int     `json:"product_money"`
				ProductCount             int     `json:"product_count"`
				DiscountMoney            int     `json:"discount_money"`
				UserPayState             string  `json:"user_pay_state"`
				WarehouseID              int     `json:"warehouse_id"`
				WarehouseName            string  `json:"warehouse_name"`
				SendState                int     `json:"send_state"`
				SendProductCount         int     `json:"send_product_count"`
				ReturnState              string  `json:"return_state"`
				IsCanRefund              int     `json:"is_can_refund"`
				IsPreOrderCreate         string  `json:"is_pre_order_create"`
				EarnestMoney             int     `json:"earnest_money"`
				Mark                     int     `json:"mark"`
				MarkNote                 string  `json:"mark_note,omitempty"`
				SampleType               int     `json:"sample_type"`
				ActivityType             int     `json:"activity_type"`
				SurplusMoney             int     `json:"surplus_money"`
				IsOutstockOrder          int     `json:"is_outstock_order"`
				RestMoney                int     `json:"rest_money"`
				OrderEarnestMoney        int     `json:"order_earnest_money"`
				IsEarnestPay             int     `json:"is_earnest_pay"`
				PayType                  int     `json:"pay_type"`
				IsCancel                 int     `json:"is_cancel"`
				RefundOrder              int     `json:"refund_order"`
				OrderStateName           string  `json:"order_state_name"`
				OrderStateColor          string  `json:"order_state_color"`
				PlatformSubsidy          int     `json:"platform_subsidy"`
				SysjNew                  int     `json:"sysj_new"`
				IsReturn                 int     `json:"is_return"`
				ConfirmCountdown         int     `json:"confirm_countdown"`
				RefundProductCount       int     `json:"refund_product_count"`
				TkCouponDiscountMoney    int     `json:"tk_coupon_discount_money"`
				TkCouponID               int     `json:"tk_coupon_id"`
				RefundingProductCount    int     `json:"refunding_product_count"`
				ShopName                 string  `json:"shop_name"`
				StationedUserID          int     `json:"stationed_user_id"`
				IsSellerMessage          int     `json:"is_seller_message"`
				ReturningProductCount    int     `json:"returning_product_count"`
				OrderProductServiceMoney int     `json:"order_product_service_money"`
				CloseType                int     `json:"close_type"`
				ExpressInfo              struct {
					ExpressNumber  string `json:"express_number"`
					SendDate       string `json:"send_date"`
					ReceivingName  string `json:"receiving_name"`
					ReceivingPhone string `json:"receiving_phone"`
					Deliverystatus int    `json:"deliverystatus"`
					State          int    `json:"state"`
					Error          string `json:"error"`
					ExpressMessage string `json:"express_message"`
					ExpressDate    string `json:"express_date"`
				} `json:"express_info,omitempty"`
				Rn         int    `json:"rn"`
				CancelDate string `json:"cancel_date,omitempty"`
			} `json:"orderList"`
			Urlnow  string `json:"urlnow"`
			Allpage int    `json:"allpage"`
			Index   int    `json:"index"`
		} `json:"page"`
	} `json:"data"`
	Msg string `json:"msg"`
}

type OrderOnGetHandler struct {
	is lvccx2.LvOrderService
}

func NewOrderOnGetHandler(is lvccx2.LvOrderService) *OrderOnGetHandler {
	return &OrderOnGetHandler{
		is: is,
	}
}

//GetOrderList 获得用户下单列表
func (oh *OrderOnGetHandler) GetOrderList(c *gin.Context) pkg.Render {
	var info OrderList

	parameter := common.PostOrderListParameter{}

	// 通过ShouldBind函数，将请求参数绑定到struct对象， 处理json请求代码是一样的。
	// 如果是post请求则根据Content-Type判断，接收的是json数据，还是普通的http请求参数
	if c.ShouldBind(&parameter) == nil {
		// 绑定成功， 打印请求参数
		fmt.Println(parameter.Page)
	}

	//获取请求 Token
	logToken, Cookies := lvccx.GetLogTokenAndCookie()

	_url := oh.is.GetOrderUrl(parameter.Page, parameter.State, parameter.OrderNumber, parameter.ThirdNumber,
		parameter.SubOrderNumber, parameter.StationedUserName, parameter.ReceiveName, parameter.ReceivingPhone,
		parameter.ItemNumber, parameter.ProductName, parameter.LogisticsCompanyName, parameter.LogisticsNumber,
		parameter.BeginDate, parameter.EndDate, parameter.SendBeginDate, parameter.SendEndDate, logToken, parameter.Async)

	headers := make(map[string]string)
	headers["cache-control"] = "no-cache"
	headers["Content-Type"] = "application/json"
	headers["X-Requested-With"] = "XMLHttpRequest"
	headers["Cookie"] = LvApiType + "; " + Cookies + "; ua="

	_, data, err := request.Get(_url, 2*time.Second, 3, request.Headers(headers))

	if err != nil {

	}
	fmt.Println(string(data))
	//处理接口返回数据
	if errJson := json.Unmarshal(data, &info); errJson != nil {

	}

	return common.Succ(info)
}
