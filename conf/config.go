package conf

import "time"

type Config struct {
	App      App      `mapstructure:"app" yaml:"app"`
	Api      Api      `mapstructure:"api" yaml:"api"`
	Database Database `mapstructure:"database" yaml:"database"`
	Redis    Redis    `mapstructure:"redis" yaml:"redis"`
	Kafka    Kafka    `mapstructure:"kafka" yaml:"kafka"`
	Server   Server   `mapstructure:"server" yaml:"server"`
	Zap      Zap      `mapstructure:"zap" yaml:"zap"`
	Wechat   Wechat   `mapstructure:"wechat" yaml:"wechat"`
	Express  Express  `mapstructure:"express" yaml:"express"`
}

type App struct {
	Domain    string `mapstructure:"domain" yaml:"domain"`
	JwtSecret string `mapstructure:"jwt-secret" yaml:"jwt-secret"`
	PageSize  int    `mapstructure:"page-size" yaml:"page-size"`
	PrefixUrl string `mapstructure:"prefix-url" yaml:"prefix-url"`

	RuntimeRootPath string `mapstructure:"runtime-root-path" yaml:"runtime-root-path"`

	ImageSavePath  string   `mapstructure:"image-save-path" yaml:"image-save-path"`
	ImageMaxSize   int      `mapstructure:"image-max-size" yaml:"image-max-size"`
	ImageAllowExts []string `mapstructure:"image-allow-exts" yaml:"image-allow-exts"`

	ExportSavePath string `mapstructure:"export-save-path" yaml:"export-save-path"`
	QrCodeSavePath string `mapstructure:"qrcode-save-path" yaml:"qrcode-save-path"`
	FontSavePath   string `mapstructure:"font-save-path" yaml:"font-save-path"`

	LogSavePath string `mapstructure:"log-save-path" yaml:"log-save-path"`
	LogSaveName string `mapstructure:"log-save-name" yaml:"log-save-name"`
	LogFileExt  string `mapstructure:"log-file-ext" yaml:"log-file-ext"`
	TimeFormat  string `mapstructure:"time-format" yaml:"time-format"`
}

type Api struct {
	SearchProductAK string `mapstructure:"search-product-ak" yaml:"search-product-ak"`
	SearchProductSK string `mapstructure:"search-product-sk" yaml:"search-product-sk"`
}

type Database struct {
	Type        string `mapstructure:"type" yaml:"type"`
	User        string `mapstructure:"user" yaml:"user"`
	Password    string `mapstructure:"password" yaml:"password"`
	Host        string `mapstructure:"host" yaml:"host"`
	Name        string `mapstructure:"name" yaml:"name"`
	TablePrefix string `mapstructure:"table-prefix" yaml:"table-prefix"`
}

type Redis struct {
	Host        string        `mapstructure:"host" yaml:"host"`
	Password    string        `mapstructure:"password" yaml:"password"`
	IdleTimeout time.Duration `mapstructure:"idle-timeout" yaml:"idle-timeout"`
}

type Server struct {
	RunMode      string        `mapstructure:"run-mode" yaml:"run-mode"`
	HttpPort     int           `mapstructure:"http-port" yaml:"http-port"`
	ReadTimeout  time.Duration `mapstructure:"read-timeout" yaml:"read-timeout"`
	WriteTimeout time.Duration `mapstructure:"write-timeout" yaml:"write-timeout"`
}

type Zap struct {
	LogFilePath     string `mapstructure:"log-filepath" yaml:"log-filepath"`
	LogInfoFileName string `mapstructure:"log-info-filename" yaml:"log-info-filename"`
	LogWarnFileName string `mapstructure:"log-warn-filename" yaml:"log-warn-filename"`
	LogFileExt      string `mapstructure:"log-fiile-ext" yaml:"log-fiile-ext"`
	LogConsole      bool   `mapstructure:"log-console" yaml:"log-console"`
}

type Kafka struct {
	Hosts []string `mapstructure:"hosts" yaml:"hosts"`
}

type Wechat struct {
	AppID          string `mapstructure:"app_id" yaml:"app_id"`                     //appid
	AppSecret      string `mapstructure:"app_secret" yaml:"app_secret"`             //appsecret
	Token          string `mapstructure:"token" yaml:"token"`                       //token
	EncodingAESKey string `mapstructure:"encoding_aes_key" yaml:"encoding_aes_key"` //EncodingAESKey
}

type Express struct {
	EBusinessId string `eBusinessId:"host" yaml:"eBusinessId"`
	AppKey      string `mapstructure:"appKey" yaml:"appKey"`
}
