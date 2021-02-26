package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// C 全局配置(需要先执行MustLoad，否则拿不到配置)
var C *Config

// Config 配置参数
type Config struct {
	RunMode string
	PORT    string
	ENV     string
	Log     Log
	Redis   Redis
	DB      DB
	NSQ     NSQ
	Swagger bool
	HTTP    HTTP
	JWTAuth JWTAuth
	CORS    CORS
}

//Init config
func Init() {
	viper.SetConfigName("config") //  设置配置文件名 (不带后缀)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("config") // 比如添加当前目录
	//viper.AddConfigPath("/workspace/appName1") // 可以多次调用添加路径
	err := viper.ReadInConfig() // 搜索路径，并读取配置数据
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}
	C = &Config{
		RunMode: viper.GetString("runmode"),
		PORT:    viper.GetString("port"),
		ENV:     viper.GetString("env"),
	}
	C.Log = Log{
		Level:         viper.GetInt("log.level"),
		TestLevel:     viper.GetInt("log.test_level"),
		Format:        "",
		Output:        viper.GetString("log.file"),
		OutputFile:    "",
		EnableHook:    false,
		HookLevels:    nil,
		HookMaxThread: 0,
		HookMaxBuffer: 0,
	}
	C.DB = DB{
		Debug:             viper.GetBool("db.debug"),
		Type:              viper.GetString("db.type"),
		Addr:              viper.GetString("db.addr"),
		User:              viper.GetString("db.username"),
		Password:          viper.GetString("db.password"),
		Name:              viper.GetString("db.dbname"),
		Parameters:        viper.GetString("db.parameters"),
		MaxLifetime:       viper.GetInt("db.max_lifetime"),
		MaxOpenConns:      viper.GetInt("db.max_open_conns"),
		MaxIdleConns:      viper.GetInt("db.max_idle_conns"),
		TablePrefix:       viper.GetString("db.table_prefix"),
		EnableAutoMigrate: viper.GetBool("db.enable_auto_migrate"),
	}
	C.Redis = Redis{
		Addr:        viper.GetString("redis.addr"),
		Password:    viper.GetString("redis.password"),
		DB:          viper.GetInt("redis.db"),
		MaxIdle:     viper.GetInt("redis.max_idle"),
		MaxActive:   viper.GetInt("redis.max_active"),
		IdleTimeout: viper.GetInt("redis.idle_timeout"),
	}
	C.NSQ = NSQ{
		Addr:    viper.GetString("nsq.addr"),
		Topic:   viper.GetString("nsq.topic"),
		Channel: viper.GetString("nsq.channel"),
	}
}

// Log 日志配置参数
type Log struct {
	Level         int
	TestLevel     int
	Format        string
	Output        string
	OutputFile    string
	EnableHook    bool
	HookLevels    []string
	HookMaxThread int
	HookMaxBuffer int
}

// NSQ 配置参数
type NSQ struct {
	Addr    string
	Topic   string
	Channel string
}

// Redis redis配置参数
type Redis struct {
	Addr        string
	Password    string
	DB          int
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}

// DB Gorm gorm配置参数 mysql配置参数
type DB struct {
	Debug             bool
	Type              string
	Addr              string
	User              string
	Password          string
	Name              string
	Parameters        string
	MaxLifetime       int
	MaxOpenConns      int
	MaxIdleConns      int
	TablePrefix       string
	EnableAutoMigrate bool
}

// DSN 数据库连接串
func (d DB) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		d.User, d.Password, d.Addr, d.Name, d.Parameters)
}

// JWTAuth 用户认证
type JWTAuth struct {
	Enable        bool
	SigningMethod string
	SigningKey    string
	Expired       int
	Store         string
	FilePath      string
	RedisDB       int
	RedisPrefix   string
}

// HTTP http配置参数
type HTTP struct {
	Host             string
	Port             int
	CertFile         string
	KeyFile          string
	ShutdownTimeout  int
	MaxContentLength int64
	MaxLoggerLength  int `default:"4096"`
}

// CORS 跨域请求配置参数
type CORS struct {
	Enable           bool
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
	MaxAge           int
}
