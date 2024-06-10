package lvccx

import (
	"bytes"
	"fmt"
	"gitlab.xfq.com/tech-lab/dionysus/pkg/logger"
	"regexp"
	"tbTool/pkg/request"
	"time"
)

const (
	//LCUrl 吕橙子官网-初始
	LCUrl = "https://passport.lvccx.com/"

	//LCLoginUrl LCUrl吕橙子官网-登錄
	LCLoginUrl = "https://passport.lvccx.com/?cburl=https://www.lvccx.com&exit="

	//UserName 吕橙子官网-用户名
	UserName = "herec"

	//PasWorld 吕橙子官网-密码
	PasWorld = "117339d9fe4b4f89f259aff03688fb7f"

	//MustCompileStr 正则规则
	MustCompileStr = `<input type="hidden" name="logtoken" value="(.*?)">`

	//UabCollina Cookie前缀 TODO：后期可能出问题
	UABCollina = "_uab_collina=171752826184269165788333;"
)

var LogToken string
var Cookies string
var initLogToken string

func getLVLogin() {
	resp, data, err := request.Get(LCUrl, 2*time.Second, 3)

	Cookies = ""
	for _, cookie := range resp.Cookies() {
		Cookies = cookie.Name + "=" + cookie.Value
	}

	if err != nil {
		logger.Warnf("LC LogToken Get err, request err: %v", err)
	}

	//提取 initLogToken
	initLogToken = getUrlDataMustCompile(string(data))

	//urlParameter 吕橙登录接口参数
	urlParameter := getLoginParameters(initLogToken)

	//Headers 添加请求头
	headers := make(map[string]string)
	headers["X-Requested-With"] = "XMLHttpRequest"
	headers["Content-Type"] = "application/x-www-form-urlencoded; charset=UTF-8"
	headers["Cookie"] = UABCollina + Cookies

	//根據初始 LogToken 替換 賬號的 LogToken
	_, postData, postErr := request.Post(LCLoginUrl, urlParameter, 2*time.Second, 3, request.Headers(headers))
	if postErr != nil {
		logger.Warnf("LC LogLoginToken Post err, request err: %v", postErr)
	}

	fmt.Println(string(postData))
	LogToken = getUrlDataMustCompile(string(postData))
}

//GetLogTokenAndCookie 获取请求token
func GetLogTokenAndCookie() (string, string) {
	//保存缓存

	getLVLogin()
	//添加etcd配置

	return LogToken, Cookies
}

func getLoginParameters(initLogToken string) []byte {
	var buff bytes.Buffer

	buff.WriteString("action=login&logtoken=")
	buff.WriteString(initLogToken)
	buff.WriteString("&async=1&username=")
	buff.WriteString(UserName)
	buff.WriteString("&pwd=")
	buff.WriteString(PasWorld)

	return buff.Bytes()
}

func getUrlDataMustCompile(urlData string) string {
	//提取 logToken
	compileRegex := regexp.MustCompile(MustCompileStr)
	resLogToken := compileRegex.FindAllStringSubmatch(urlData, -1)

	token := ""
	for _, v := range resLogToken {
		token = v[1]
	}

	return token
}
