package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
	"text/template"
)

func main() {
	config := InitConfig()
	startGostClient(config)
}

var gostLTpl = "ss://{{.Method}}:{{.SSPassword}}@{{.SSLocalAddress}}:{{.SSPort}}"
var gostFTpl = "https://{{.GostAuth}}@{{.GostAddress}}:{{.GostPort}}"

func startGostClient(config *Config) {
	if config == nil {
		log.Fatalf("config is empty")
	}
	// 初始化参数模板
	if config.Debug {
		debug(fmt.Sprintf("config: %s", config.toString()))
	}
	gostL, err := template.New("gostL").Parse(gostLTpl)
	if err != nil {
		log.Fatalf("parse gostL tpl err: %v", err)
	}
	gostF, err := template.New("gostF").Parse(gostFTpl)
	if err != nil {
		log.Fatalf("parse gostF tpl err: %v", err)
	}
	// 渲染执行命令
	var LArg bytes.Buffer
	err = gostL.Execute(&LArg, config)
	if err != nil {
		log.Fatalf("render LArg error: %v", err)
	}
	var FArg bytes.Buffer
	err = gostF.Execute(&FArg, config)
	if err != nil {
		log.Fatalf("render FArg error: %v", err)
	}
	// 执行
	if config.Debug {
		debug(fmt.Sprintf("%s -L %s -F %s", config.GostPath, LArg.String(), FArg.String()))
	}
	cmd := exec.Command(config.GostPath, "-L", LArg.String(), "-F", FArg.String())
	if !config.Debug {
		cmd.Stdout = io.Discard
		cmd.Stderr = io.Discard
	}
}

func debug(text string) {
	fmt.Println(fmt.Sprintf("shadowsocks-gost debug: %s", text))
}
