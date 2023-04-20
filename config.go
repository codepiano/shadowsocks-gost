package main

import "text/template"


type Config struct {
	/**
	shadowsocks 官方不支持禁用加密，综合现有信息来看，chacha20-ietf-poly1305 性能比较快
	refer: https://github.com/shadowsocks/shadowsocks-libev/issues/762
	*/
	Method     string `json:"method"`
	SSPassword string `json:"ss_password"`
	SSPort     int    `json:"ss_port"`

	// Gost 服务器配置，gost 客户端通过下面的配置连接 gost 服务器
	GostHost     string `json:"gost_host"`
	GostUser     string `json:"gost_user"`
	GostPassword string `json:"gost_password"`
}

func InitConfig() {
	c := &Config{
		Method: "chacha20-ietf-poly1305",
	}
}


sweaters := Inventory{"wool", 17}
if err != nil { panic(err) }
err = tmpl.Execute(os.Stdout, sweaters)
if err != nil { panic(err) }
