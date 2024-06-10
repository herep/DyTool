package etcd

import (
	"encoding/json"

	jsoniter "github.com/json-iterator/go"

	"gitlab.xfq.com/tech-lab/dionysus/pkg/conf"
	"gitlab.xfq.com/tech-lab/dionysus/pkg/logger"
)

func GetStringByEtcd(key string) string {
	rb, err := conf.LoadBytes(key)
	if err != nil {
		logger.Errorf("从配置信息获取k:%s的值,获取失败 err:%v-_- ...", key, err)
		return ""
	}
	return string(rb)
}

func GetBytesByEtcd(key string) []byte {
	rb, err := conf.LoadBytes(key)
	if err != nil {
		logger.Errorf("从配置信息获取k:%s,获取失败 err:%v-_- ...", key, err)
		return nil
	}
	return rb
}

func GetDbMapForEtcd(prefix, key string) string {
	return getMapStringEtcd(prefix, key)
}

func getMapByEtcd(key string) map[string]string {
	rb, err := conf.LoadBytes(key)
	if err != nil {
		logger.Errorf("从配置信息获取k:%v的值,获取失败 err:%v-_- ...", key, err)
		return nil
	}
	m := make(map[string]string)
	if err := json.Unmarshal(rb, &m); err != nil {
		logger.Errorf("从配置信息获取k:%v的值,json unmarshal 失败 err:%v-_- ...", key, err)
		return nil
	}
	return m
}

func getMapStringEtcd(prefix, key string) string {
	rm := getMapByEtcd(prefix)
	if v, ok := rm[key]; ok {
		return v
	}
	return ""
}

func GetSliceString(key string) []string {
	empty := make([]string, 0)
	data, err := conf.LoadBytes(key)
	if err != nil {
		logger.Errorf("从配置信息获取k:%v的值,获取失败 err:%v-_- ...", key, err)
		return empty
	}
	c := make([]string, 0)
	err = jsoniter.Unmarshal(data, &c)
	if err != nil {
		logger.Errorf("从配置信息获取k:%v的值,获取失败 err:%v-_- ...", key, err)
		return empty
	}

	return c
}
