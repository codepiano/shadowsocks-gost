package main

import (
	"bytes"
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
	// 初始化参数模板
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
	cmd := exec.Command(config.GostPath, "-L", LArg.String(), "-F", FArg.String())
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("get stdout error: %v", err)
	}
	cmd.Stderr = cmd.Stdout
	err = cmd.Start()
	if err != nil {
		log.Fatalf("start gost failed, err: %v", err)
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		if err != nil {
			break
		}
	}
}
