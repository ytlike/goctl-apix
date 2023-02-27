package config

type Server struct {
	Application     `yaml:"application"`
	Mysql           `yaml:"mysql"`
	Redis           `yaml:"redis"`
	Paas            `yaml:"paas"`
	Rocketmq        `yaml:"rocketmq"`
	AbilityUrl      `yaml:"abilityUrl"`
	OpenApplication `yaml:"openApplication"`
}

type Application struct {
	ListenOn    string `yaml:"listenOn"`
	Timeout     int64  `yaml:"timeout"`
	Health      bool   `yaml:"health"`
	ServiceConf `yaml:"serviceConf"`
}

type ServiceConf struct {
	Name      string `yaml:"name"`
	Mode      string `yaml:"mode"`
	Log       `yaml:"log"`
	Telemetry `yaml:"telemetry"`
	DevServer `yaml:"devServer"`
}

type Mysql struct {
	DataSource string `yaml:"dataSource"`
}

type Redis struct {
	Host string `yaml:"host"`
	Pass string `yaml:"pass"`
	Type string `yaml:"type"`
}

type Rocketmq struct {
	NameServer string `yaml:"name-server"`
}

type Paas struct {
	ClientId     string `yaml:"clientId"`
	ClientSecret string `yaml:"clientSecret"`
	PaasUrl      string `yaml:"paasUrl"`
	Retry        int32  `yaml:"retry"`
}

type Log struct {
	ServiceName         string `yaml:"serviceName"`
	Level               string `yaml:"level"`
	Mode                string `yaml:"mode"`
	TimeFormat          string `yaml:"timeFormat"`
	Path                string `yaml:"path"`
	Encoding            string `yaml:"encoding"`
	KeepDays            string `yaml:"keepDays"`
	Rotation            string `yaml:"rotation"`
	StackCooldownMillis int    `yaml:"stackCooldownMillis"`
}

type AbilityUrl struct {
	AliCloudUrl         string `yaml:"aliCloudUrl"`
	DigitalChongQingUrl string `yaml:"digitalChongQingUrl"`
}

type OpenApplication struct {
	AliCloudAppIdBegin         int64    `yaml:"aliCloudAppIdBegin"`
	DigitalChongQingAppIdBegin int64    `yaml:"digitalChongQingAppIdBegin"`
	PId                        int32    `yaml:"pid"`
	AccessTokenValidity        int32    `yaml:"accessTokenValidity"`
	RefreshTokenValidity       int32    `yaml:"refreshTokenValidity"`
	GrantTypes                 []string `yaml:"grantTypes"`
	ResourceIds                []string `yaml:"resourceIds"`
}

type Telemetry struct {
	Name     string  `yaml:"name"`
	Endpoint string  `yaml:"endpoint"`
	Sampler  float64 `yaml:"sampler"`
	Batcher  string  `yaml:"batcher"`
}

type DevServer struct {
	Enabled       bool   `yaml:"enabled"`
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	MetricsPath   string `yaml:"metricsPath"`
	HealthPath    string `yaml:"healthPath"`
	EnableMetrics bool   `yaml:"enableMetrics"`
	EnablePprof   bool   `yaml:"enablePprof"`
}
