package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/liangpengcheng/qcontinuum/base"
)

type config struct {
	Address string
	ID      int32
	Center  string
}

var cfg config

func init() {
	fp, err := os.Open("../runtime/gateway.json")
	if err == nil {
		buf, err := ioutil.ReadAll(fp)
		base.CheckError(err, "read buf error:")
		err = json.Unmarshal(buf, &cfg)
		base.CheckError(err, "json unmarshal error:")
	} else {
		base.LogPanic("open gateway.json failed :%s", err.Error())
	}
}
