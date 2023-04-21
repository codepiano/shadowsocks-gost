package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"os"
	"strconv"
)

type Config struct {
	SSLocalAddress string `json:"ss_local_address"`
	SSPort         int    `json:"ss_port"`
	SSPassword     string `json:"ss_password"`
	/**
	shadowsocks 官方不支持禁用加密，综合现有信息来看，chacha20-ietf-poly1305 性能比较快
	refer: https://github.com/shadowsocks/shadowsocks-libev/issues/762
	*/
	Method string `json:"method"`

	// Gost 服务器配置，gost 客户端通过下面的配置连接 gost 服务器
	GostAddress string `json:"gost_Address"`
	GostPort    int    `json:"gost_port"`
	// 认证串，格式为 user:password 的 base64 形式
	GostAuth string `json:"gost_auth"`
}

func InitConfig() {
	c := &Config{}
	c.getEnvConfigs()
	if c.SSPort == 0 || c.SSPassword == "" || c.GostAddress == "" || c.GostPort == 0 || c.GostAuth == "" {
		log.Fatalf("config invalid")
	}
	if c.Method == "" {
		c.Method = "chacha20-ietf-poly1305"
	}
	if c.SSLocalAddress == "" {
		c.SSLocalAddress = "127.0.0.1"
	}
	if c.SSPassword == "" {

	}
}

func (c *Config) getEnvConfigs() {
	var port string
	var remotePort string
	var pluginOpts string
	var err error

	c.SSLocalAddress = os.Getenv("SS_LOCAL_HOST")

	port = os.Getenv("SS_LOCAL_PORT")
	c.SSPort, err = strconv.Atoi(port)
	if err != nil {
		log.Fatalf("ss local port invalid, port: %s, err: %v", port, err)
	}

	c.GostAddress = os.Getenv("SS_REMOTE_HOST")

	remotePort = os.Getenv("SS_REMOTE_PORT")
	c.GostPort, err = strconv.Atoi(remotePort)
	if err != nil {
		log.Fatalf("ss local port invalid, port: %s, err: %v", remotePort, err)
	}

	pluginOpts = os.Getenv("SS_PLUGIN_OPTIONS")
	if pluginOpts == "" {
		log.Fatal("no gost auth config")
	}
	c.GostAuth = base64.StdEncoding.EncodeToString([]byte(pluginOpts))
	writeTo(c.toString())
}

func (c *Config) toString() string {
	data, err := json.Marshal(c)
	if err != nil {
		log.Fatalf("marshal config err: %v", err)
	}
	return string(data)
}
