package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	var localHost string
	var localPort string
	var remoteHost string
	var remotePort string
	var pluginOpts string

	localHost = os.Getenv("SS_LOCAL_HOST")
	localPort = os.Getenv("SS_LOCAL_PORT")
	remoteHost = os.Getenv("SS_REMOTE_HOST")
	remotePort = os.Getenv("SS_REMOTE_PORT")
	pluginOpts = os.Getenv("SS_PLUGIN_OPTIONS")
	writeTo(fmt.Sprintf("%s|%s|%s|%s|%s", localHost, localPort, remoteHost, remotePort, pluginOpts))
	for _, arg := range os.Args {
		writeTo(arg)
	}
	gostCmd := fmt.Sprintf("gost -L ss://%s:%s@:%d -F 'https://%s:%s@%s:%d'")
}

func startGostClient() {

	tmpl, err := template.New("gost").Parse("gost -L 'ss://{{.Method}}:{{.SSPassword}}@:{{.SSPort}}' -F ''")
	if err != nil {
		log.Fatalf("parse gost cmd err: %v", err)
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
