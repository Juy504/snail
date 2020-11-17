package conf

import (
	"os"
)

type Wxconf struct{
	Appid  string
	Appkey string
	Rurl   string
}

var ServerConf = make(map[string]string)

func init() {
	curDir, err := os.Getwd()
	if err != nil{
		panic("directory error")
	}

	ServerConf["logDir"] = curDir + "/../log"
}
