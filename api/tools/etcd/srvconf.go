package etcd

import (
	"encoding/json"
	"strconv"
)

/*
	第三方服务，例如 微信小程序，腾讯STS 等第三方配置
*/

var prefix = "business."

// GetTmaId 字节跳动分配的小程序id
func GetTmaId() string {
	return GetStringByEtcd(prefix + "toutiao_config.tuotiao_tmaid")
}

// GetSecretId 腾讯STS的secret_id
func GetSecretId() string {
	rm := getMapByEtcd(prefix + "wx_sts_config")
	if v, ok := rm["secret_id"]; ok {
		return v
	}
	return ""
}

// GetSecretKey 腾讯STS的secret_key
func GetSecretKey() string {
	rm := getMapByEtcd(prefix + "wx_sts_config")
	if v, ok := rm["secret_key"]; ok {
		return v
	}
	return ""
}

// GetBucket  bucket-appId
func GetBucket() string {
	rm := getMapByEtcd(prefix + "wx_sts_config")
	if v, ok := rm["bucket"]; ok {
		return v
	}
	return ""
}

// GetRegion 所在地区
func GetRegion() string {
	rm := getMapByEtcd(prefix + "wx_sts_config")
	if v, ok := rm["region"]; ok {
		return v
	}
	return ""
}

//GetQiNiuKoDoInfo 千牛
func GetQiNiuKoDoInfo() (accessKey, secretKey string) {
	accessKey = getMapStringEtcd(prefix+"qiniu.qiniu_kodo", "accessKey")
	secretKey = getMapStringEtcd(prefix+"qiniu.qiniu_kodo", "secretKey")

	return
}

/*
	GetVolumeByKey 放量配置
*/
func GetVolumeByKey(key string) []string {
	vci := GetBytesByEtcd(prefix + "volume_conf")
	if len(vci) == 0 {
		return nil
	}
	mr := make(map[string][]string)
	err := json.Unmarshal(vci, &mr)
	if err != nil {
		return nil
	}

	if v, ok := mr[key]; ok {
		return v
	}

	return nil
}

/*
	数据部门 sts
*/
// 腾讯STS的secret_id
func GetDaStoreSecretId() string {
	rm := getMapByEtcd(prefix + "wx_sts_config_dastore")
	if v, ok := rm["secret_id"]; ok {
		return v
	}
	return ""
}

// GetDaStoreSecretKey 腾讯STS的secret_key
func GetDaStoreSecretKey() string {
	rm := getMapByEtcd(prefix + "wx_sts_config_dastore")
	if v, ok := rm["secret_key"]; ok {
		return v
	}
	return ""
}

// GetDaStoreBucket  bucket-appId
func GetDaStoreBucket() string {
	rm := getMapByEtcd(prefix + "wx_sts_config_dastore")
	if v, ok := rm["bucket"]; ok {
		return v
	}
	return ""
}

// GetDaStoreRegion 所在地区
func GetDaStoreRegion() string {
	rm := getMapByEtcd(prefix + "wx_sts_config_dastore")
	if v, ok := rm["region"]; ok {
		return v
	}
	return ""
}

/*
	数据部门 sts
*/
// 腾讯STS的secret_id
func GetCfcSecretId() string {
	rm := getMapByEtcd(prefix + "certchain-cos")
	if v, ok := rm["secret_id"]; ok {
		return v
	}
	return ""
}

// GetCfcSecretKey 腾讯STS的secret_key
func GetCfcSecretKey() string {
	rm := getMapByEtcd(prefix + "certchain-cos")
	if v, ok := rm["secret_key"]; ok {
		return v
	}
	return ""
}

// GetCfcBucket bucket  bucket-appId
func GetCfcBucket() string {
	rm := getMapByEtcd(prefix + "certchain-cos")
	if v, ok := rm["bucket"]; ok {
		return v
	}
	return ""
}

// GetCfcRegion bucket 所在地区
func GetCfcRegion() string {
	rm := getMapByEtcd(prefix + "certchain-cos")
	if v, ok := rm["region"]; ok {
		return v
	}
	return ""
}

func GetQrLoginFlowControl() int {
	str := GetStringByEtcd("business.qrlogin_flow_control")
	if len(str) == 0 {
		return 100
	}
	flow, err := strconv.Atoi(str)
	if err != nil {
		return 100
	}
	return flow
}
