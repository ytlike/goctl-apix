package config

type Server struct {
	Application `yaml:"application"`
	Auth        `yaml:"auth"`
	Redis       `yaml:"redis"`
	Rpc         `yaml:"rpc"`
}

type Application struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Timeout     int64  `yaml:"timeout"`
	MaxConns    int    `yaml:"maxConns"`
	MaxBytes    int64  `yaml:"maxBytes"`
	ServiceConf `yaml:"serviceConf"`
}

type ServiceConf struct {
	Name      string `yaml:"name"`
	Mode      string `yaml:"mode"`
	Log       `yaml:"log"`
	Telemetry `yaml:"telemetry"`
	DevServer `yaml:"devServer"`
}

type Auth struct {
	CheckUrl   string   `yaml:"checkUrl"`
	IgnoreUrls []string `yaml:"ignoreUrls"`
}

type Redis struct {
	Host string `yaml:"host"`
	Pass string `yaml:"pass"`
	Type string `yaml:"type"`
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

type Rpc struct {
	App []struct {
		Name    string `yaml:"name"`
		Timeout int64  `yaml:"timeout"`
	}
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
