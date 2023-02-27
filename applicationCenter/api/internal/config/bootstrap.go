package config

type Bootstrap struct {
	Nacos `yaml:"nacos"`
	*Server
}

type Nacos struct {
	Discovery `yaml:"discovery"`
	Configure `yaml:"configure"`
}

type Discovery struct {
	Host        string `yaml:"host"`
	Port        uint64 `yaml:"port"`
	Timeout     uint64 `yaml:"timeout"`
	NamespaceId string `yaml:"namespaceId"`
	Group       string `yaml:"group"`
}

type Configure struct {
	Host        string `yaml:"host"`
	Port        uint64 `yaml:"port"`
	Timeout     uint64 `yaml:"timeout"`
	NamespaceId string `yaml:"namespaceId"`
	Group       string `yaml:"group"`
	DataId      string `yaml:"dataId"`
}
