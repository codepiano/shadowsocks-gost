package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
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
		writeTo("parse gostL tpl err: %v", err)
	}
	gostF, err := template.New("gostF").Parse(gostFTpl)
	if err != nil {
		writeTo("parse gostF tpl err: %v", err)
	}
	// 渲染执行命令
	var LArg bytes.Buffer
	err = gostL.Execute(&LArg, config)
	if err != nil {
		writeTo("render LArg error: %v", err)
	}
	var FArg bytes.Buffer
	err = gostF.Execute(&FArg, config)
	if err != nil {
		writeTo("render FArg error: %v", err)
	}
	// 执行
	cmd := exec.Command(config.GostPath, "-L", LArg.String(), "-F", FArg.String())
	writeTo(cmd.String())
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		writeTo("get stdout error: %v", err)
	}
	cmd.Stderr = cmd.Stdout
	writeTo(strings.Join(os.Environ(), "|"))
	err = cmd.Start()
	if err != nil {
		writeTo("start gost failed, err: %v", err)
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		writeTo(string(tmp))
		if err != nil {
			break
		}
	}
}

func writeTo(format string, v ...any) {
	test := fmt.Sprintf(format, v...)
	f, err := os.OpenFile("/tmp/log.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(fmt.Sprintf("%s\n", test)); err != nil {
		log.Println(err)
	}
}
