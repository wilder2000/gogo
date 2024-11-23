package config

import (
	"fmt"
	"os"
	"time"
)

type AppConfig struct {
	Port            string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	DataSource      DBConfig
	ExpireTime      time.Duration
	MinioInfo       Minio
	UserNames       []string
	DefaultAvatar   string
	DocProcPoolSize int
	MaxCurrThread   int
	CoverWidth      int
	CoverHeight     int
	AppSecret       AppSecret
	StaticDir       StaticDir
	UserService     bool
	ReportError     bool
	Security        Security
	Web             map[string]string

	AccessControlAllowOrigin  bool
	AccessControlAllowHost    string
	AccessControlAllowMethods string
	AccessControlAllowHeaders string
}
type Security struct {
	Registration     bool
	MaxTryTimes      int
	ForbidAccessTime float64
}
type StaticDir struct {
	RelativePath     string
	AbsoluteFileDir  string
	RelativePath2    string
	AbsoluteFileDir2 string
}
type AppSecret struct {
	AccessKey string
	SecretKey string
}
type Minio struct {
	Url           string
	Account       string
	Password      string
	BucketName    string
	PrivateBucket string
	Endpoint      string
}

var (
	AConfig *AppConfig
)

func initAvatorHome() {
	//create static image home
	fmt.Printf("AbsoluteFileDir:%s\n", AConfig.StaticDir.AbsoluteFileDir)
	if len(AConfig.StaticDir.AbsoluteFileDir) == 0 {
		fmt.Printf("Http server www directory not config.\n")
		return
	}
	imgHome := AConfig.StaticDir.AbsoluteFileDir + IconHome
	_, err := os.Stat(imgHome)
	if err != nil && os.IsNotExist(err) {
		err2 := os.MkdirAll(imgHome, os.ModePerm)
		if err2 != nil {
			panic(err2)
		}
	}

}
func AvatorHome() string {
	return AConfig.StaticDir.AbsoluteFileDir + IconHome
}

const (
	IconHome = "/images"
)

func LoadConfig() {
	AConfig = &AppConfig{}
	yaml, err := ReadYAML(AppConfigFile(), ConfDir())
	if err != nil {
		panic(err)
	}
	yaml.Sub("application").Unmarshal(AConfig)
	initAvatorHome()
}

//读取配置文件

func ReadYAMLConfig[T interface{}](conf string) *T {
	model := new(T)
	yaml, err := ReadYAML(conf, ConfDir())
	if err != nil {
		panic(err)
	}
	yaml.Sub("application").Unmarshal(model)
	return model
}
func init() {
	LoadConfig()
	//log.Logger.InfoF("%v", *AConfig)
	//log.Logger.InfoF("%d", 1<<20)
	//log.Logger.InfoF("%s", 1<<20)
	//log.Logger.InfoF("names:%v", AConfig.UserNames)
}

type DBConfig struct {
	Dir                string
	Type               string
	Name               string
	DSN                string
	MaxOpenConnections int
	MaxIdleConnections int
}
