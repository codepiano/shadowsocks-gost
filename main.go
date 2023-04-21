package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	for _, arg := range os.Args {
		writeTo(arg)
	}
}

var configTpl = "gost -L 'ss://{{.Method}}:{{.SSPassword}}@{{.SSLocalAddress}}:{{.SSPort}}' -F 'https://{{.GostAddress}}:{{.GostPort}}?auth={{.GostAuth}}'"

func startGostClient(config *Config) {
	cmdTpl, err := template.New("gost").Parse(configTpl)
	if err != nil {
		log.Fatalf("parse gost cmd err: %v", err)
	}
	// 渲染执行命令
	var b *bytes.Buffer
	err = cmdTpl.Execute(b, config)
	if err != nil {
		log.Fatalf("render config error: %v", err)
	}
}

func writeTo(test string) {
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
