package tests

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/coreos/etcd/clientv3"
	conf2 "gitlab.xfq.com/tech-lab/dionysus/conf"
	"gitlab.xfq.com/tech-lab/dionysus/pkg/conf"
	"gitlab.xfq.com/tech-lab/dionysus/pkg/logger"
	dredis "gitlab.xfq.com/tech-lab/dionysus/pkg/redis"
	"google.golang.org/grpc"
)

var (
	etcdClient      *clientv3.Client
	salePrefix      = "/gapi/business/"
	saleRedisPrefix = "/gapi/watch/redis/"
	saleWatchPrefix = "/gapi/watch/"
)

type regrouter func()

//  加载测试准备工作
func DioInit(regrouters ...regrouter) error {
	var err error

	_ = os.Setenv("GAPI_PROJECT_NAME", "gapi")
	_ = os.Setenv("GAPI_CONFIG_ETCD", "127.0.0.1:2379")

	for _, fn := range regrouters {
		fn()
	}

	// 初始化etc
	DioEtc()
	if err := conf2.Setup(); err != nil {
		panic(fmt.Sprintf("err:%v", err))
	}
	// 加载redis
	err = conf.RegisterEtcdWatch(&dredis.RedisEvent{Prefix: "watch.redis"})
	if err != nil {
		panic(fmt.Sprintf("err:%v", err))
	}
	//
	err = conf.StartWatchConfig("business")
	if err != nil {
		panic(fmt.Sprintf("err:%v", err))
	}
	err = conf.StartWatchConfig("watch.mysql")
	if err != nil {
		panic(fmt.Sprintf("err:%v", err))
	}

	return nil
}

// 初始化etc,请开启本地的Ectd服务
func DioEtc() {
	etcConf := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
	}

	var err error
	if etcdClient, err = clientv3.New(etcConf); err != nil {
		logger.Error("connect etcd failed")
	}

	/*
	* 加载 配置信息
	 */

	// 1 加载redis 配置信息
	etcRedisKey := saleRedisPrefix + "REDIS_CODIS_BUSINESS_HOST"

	// test服或本地 使用
	//etcRedisVal := `{"addr":"10.3.7.39:19000", "db":0, "password":"wpt-codis-auth"}`

	// gray pr使用
	etcRedisVal := `{"addr":"127.0.0.1:6379", "db":0, "password":"foobared"}`

	DioEtcdConf(etcRedisKey, etcRedisVal)

	codisSocket := saleRedisPrefix + "CODIS_SOCKET"

	// test服或本地 使用
	//codisSocketVal := `{"addr":"10.3.7.39:19000", "db":0, "password":"wpt-codis-auth"}`

	// gray pr使用
	codisSocketVal := `{"addr":"127.0.0.1:6379", "db":0, "password":"foobared"}`

	DioEtcdConf(codisSocket, codisSocketVal)

	// 2 配置 toutiao_config.tuotiao_tmaid
	etctTmaid := salePrefix + "toutiao_config/tuotiao_tmaid"
	DioEtcdConf(etctTmaid, "tt6e391884f113e094")

	// 3 配置 ms_gw_host
	msGwHost := salePrefix + "ms_gw_host"
	DioEtcdConf(msGwHost, "http://10.3.7.10:60777")

	// 4 字节跳动.配置
	imgRecognitionAddress := salePrefix + "img_recognition_address"
	DioEtcdConf(imgRecognitionAddress, "http://10.3.7.67:5050/classifyc/api/inference")

	// 7 腾讯sts
	wxStsConfig := salePrefix + "wx_sts_config"
	wxStsConfigValue := `{"secret_id":"AKIDPNvTGLJKCZRmFp66CTIX7LAIiNiJqLOt",
"secret_key":"8elLN8S5jmSJIMzVhHTK9HZSvK93r1c6","bucket":"appwpt-10002380","region":"ap-shanghai"}`
	DioEtcdConf(wxStsConfig, wxStsConfigValue)

	// 8 腾讯cos
	certchainCosConfig := salePrefix + "certchain-cos"
	certchainCosConfigValue := `{"secret_id":"AKIDPNvTGLJKCZRmFp66CTIX7LAIiNiJqLOt",
"secret_key":"8elLN8S5jmSJIMzVhHTK9HZSvK93r1c6","bucket_url":"http://wpt-certchain-1251022884.cos.ap-shanghai.myqcloud.com"}`
	DioEtcdConf(certchainCosConfig, certchainCosConfigValue)

	// 8 图片识别白名单
	coinRecognitionWhite := salePrefix + "coin-recognition-white-list"
	coinRecognitionWhiteValue := `{"list":[47833988,51832975,50001432,59479740,28355039,36661964,1,50851532,10560309,10520049]}`
	DioEtcdConf(coinRecognitionWhite, coinRecognitionWhiteValue)

	// 9 数据库bundle_db
	dundleDb := saleWatchPrefix + "mysql/bundle_db"
	dundleDbvalue := `{"ormName":"mysql","host":"10.3.7.39","port":"3306","password":"hWhG4tjxFV*49ROY","dbname":"bundle_db","account":"wpt_test"}`
	DioEtcdConf(dundleDb, dundleDbvalue)

	// 10 七牛云存储
	qiniuConf := salePrefix + "qiniu/qiniu_kodo"
	qiniuConfValue := `{"accessKey":"FxAlYUt9lUQxNsIRjlIgPqTB8TyEGHSYo_-xRQnV","secretKey":"YdgcEwDEuX8n9VZdVCVneh_oDBjN4PWzr6s-j9lP"}`
	DioEtcdConf(qiniuConf, qiniuConfValue)

	qiniuVcdnurl := salePrefix + "qiniu/VCDNURL"
	qiniuVcdnurValue := `//media.weipaitang.com/`
	DioEtcdConf(qiniuVcdnurl, qiniuVcdnurValue)

	appCdnurl := salePrefix + "CDNURL"
	appCdnurlValue := "//cdn01t.weipaitang.com/"
	DioEtcdConf(appCdnurl, appCdnurlValue)

	appApiHost := salePrefix + "API_HOST"
	appApiHostValue := "//apit.weipaitang.com/"
	DioEtcdConf(appApiHost, appApiHostValue)

	appImApiInsideHost := salePrefix + "IM_API_INSIDE_HOST"
	appImApiInsideHostValue := "//shopimapit.weipaitang.com/"
	DioEtcdConf(appImApiInsideHost, appImApiInsideHostValue)

	app_DELICERY_CODE_SCAN_HOST := salePrefix + "DELICERY_CODE_SCAN_HOST"
	app_DELICERY_CODE_SCAN_HOST_value := "http://imgt.wpt.la"
	DioEtcdConf(app_DELICERY_CODE_SCAN_HOST, app_DELICERY_CODE_SCAN_HOST_value)

	app_SALE_GATEWAY := salePrefix + "SALE_GATEWAY"
	app_SALE_GATEWAY_value := "http://10.3.7.20:8080"
	DioEtcdConf(app_SALE_GATEWAY, app_SALE_GATEWAY_value)

	appGApiHost := salePrefix + "GAPI_HOST"
	appGApiHostValue := "//t-gapi.weipaitang.com/"
	DioEtcdConf(appGApiHost, appGApiHostValue)

	// 9 有证链肉网地址
	cfcInsideHost := salePrefix + "cfc-inside-host"
	cfcInsideHostValue := `http://cfct.weipaitang.com`
	DioEtcdConf(cfcInsideHost, cfcInsideHostValue)

	// 10 证书详情页地址
	certDetailUrl := salePrefix + "cfc-cert-detail-url"
	certDetailUrlValue := `https://wwwt.youzheng.link/detailnew`
	DioEtcdConf(certDetailUrl, certDetailUrlValue)

	// 11 放量
	VolumeConf := salePrefix + "volume_conf"
	VolumeConfValue := `{"wx_image":["2"],"scan_qrcode":["2"]}`
	DioEtcdConf(VolumeConf, VolumeConfValue)

	wxStsConfigDastore := salePrefix + "wx_sts_config_dastore"
	wxStsConfigValueDastore := `{"secret_id":"AKIDPNvTGLJKCZRmFp66CTIX7LAIiNiJqLOt",
"secret_key":"8elLN8S5jmSJIMzVhHTK9HZSvK93r1c6","bucket":"dastore-1251022884","region":"ap-shanghai"}`
	DioEtcdConf(wxStsConfigDastore, wxStsConfigValueDastore)
}

// 加载 配置信息
func DioEtcdConf(key, val string) {
	_, err := etcdClient.Put(context.Background(), key, val)
	if err != nil {
		panic(fmt.Sprintf("ETCD 测试添加key:%s value:%s 配置信息失败err:%v \n", key, val, err))
	}

	// wait to handle etc watch event
	time.Sleep(time.Millisecond * 100)
}

