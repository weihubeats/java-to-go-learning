package main

import (
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"os"
	"time"
)

const (
	SecretAppID     = "test"
	SecretNameSpace = "application"
	SecretKey       = "bd5ebb9fc733476384333b8f6a56a05f"
)

func main() {

	err := os.Setenv("AGOLLO_CONF", "/Users/xiaozoujishu/GolandProjects/java-to-go-learning/student/apollo/app.properties")
	if err != nil {
		return
	}
	agollo.SetLogger(&DefaultLogger{})

	client, err := agollo.Start()

	if err != nil {
		fmt.Println("err:", err)
		panic(err)
	}

	checkKey("application", client)

	c := &config.AppConfig{
		AppID:          SecretAppID,
		Cluster:        "dev",
		IP:             "http://127.0.0.1:8080",
		NamespaceName:  SecretNameSpace,
		IsBackupConfig: false,
		//Secret:         SecretKey,
	}

	client, err = agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})

	if err != nil {
		fmt.Println("err:", err)
		panic(err)
	}

	checkKey(c.NamespaceName, client)

	time.Sleep(5 * time.Second)

}

func checkKey(namespace string, client agollo.Client) {
	cache := client.GetConfigCache(namespace)
	count := 0
	cache.Range(func(key, value interface{}) bool {
		fmt.Println("key : ", key, ", value :", value)
		count++
		return true
	})
	if count < 1 {
		panic("config key can not be null")
	}
}

type DefaultLogger struct {
}

func (this *DefaultLogger) Debugf(format string, params ...interface{}) {
	this.Debug(format, params)
}

func (this *DefaultLogger) Infof(format string, params ...interface{}) {
	this.Debug(format, params)
}

func (this *DefaultLogger) Warnf(format string, params ...interface{}) {
	this.Debug(format, params)
}

func (this *DefaultLogger) Errorf(format string, params ...interface{}) {
	this.Debug(format, params)
}

func (this *DefaultLogger) Debug(v ...interface{}) {
	fmt.Println(v)
}
func (this *DefaultLogger) Info(v ...interface{}) {
	this.Debug(v)
}

func (this *DefaultLogger) Warn(v ...interface{}) {
	this.Debug(v)
}

func (this *DefaultLogger) Error(v ...interface{}) {
	this.Debug(v)
}