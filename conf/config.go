package conf

import "gopkg.in/ini.v1"

var ()

func Init() {
	file, err = ini.Load("./config/config.ini")
}
