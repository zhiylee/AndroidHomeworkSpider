package setting

import (
	"github.com/go-ini/ini"
	"log"
)

var (
	Cfg *ini.File

	RunMode string

	TablePrefix string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadDatabase()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadDatabase() {
	sec, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}

	TablePrefix = sec.Key("TABLE_PREFIX").MustString("")
}