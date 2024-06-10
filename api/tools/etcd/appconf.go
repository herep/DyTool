package etcd

var prefixApp = "business.qiniu."

// 七牛读取视频路径（后面带视频的Key）qiniu-cdn
func VcdnURL() string {
	return GetStringByEtcd(prefixApp + "VCDNURL")
}

// CND
func CdnURL() string {
	return GetStringByEtcd(prefix + "CDNURL")
}

// php接口域名
func ApiHost() string {
	return GetStringByEtcd(prefix + "API_HOST")
}

// IM接口地址
func ImApiInsideHost() string {
	return GetStringByEtcd(prefix + "IM_API_INSIDE_HOST")
}

// 获取订单号扫描的host
func DeliveryCodeHost() string {
	return GetStringByEtcd(prefix + "DELICERY_CODE_SCAN_HOST")
}

// g-api项目接口域名
func GApiHost() string {
	return GetStringByEtcd(prefix + "GAPI_HOST")
}

// 获取网关SALE_GATEWAY
func SaleGateway() string {
	return GetStringByEtcd("business.SALE_GATEWAY")
}

// 百川接口
func GetBaiChuanHost() string {
	return GetStringByEtcd("business.BAICHUAN_HOST")
}

// 二维码登陆白名单列表
func GetQrloginWhiteHost() string {
	return GetStringByEtcd("business.qrlogin_white_host")
}
