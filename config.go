package goflexer

import (
	"io/ioutil"
	"log"

	"github.com/go-yaml/yaml"
)

// Config holds the configuration options for the flexer context
type Config struct {
	URL         string `envconfig:"cmp_url" yaml:"cmp_url"`
	Username    string `envconfig:"cmp_username" yaml:"cpi_api_user"`
	Password    string `envconfig:"cmp_password" yaml:"cmp_api_key"`
	AccessToken string `envconfig:"cmp_access_token"`
	CodeDir     string `envconfig:"nflex_codedir"`
	ModuleID    string `envconfig:"nflex_module_id"`
	VerifySSL   bool   `envconfig:"verify_ssl" default:"true"`
}

type flexerYAMLConfig struct {
	Regions   map[string]map[string]string `yaml:"regions"`
	VerifySSL bool                         `yaml:"verify_ssl"`
}

// NewConfigFromYAML loads config from a .flexer yaml file
func NewConfigFromYAML() *Config {

	conf := &Config{}
	fConf := &flexerYAMLConfig{}
	yamlFile, err := ioutil.ReadFile("/home/mike/.flexer.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &fConf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	// TODO Review format as it's really ugly
	df, ok := fConf.Regions["default"]
	if ok {
		conf.URL = df["cmp_url"]
		conf.Username = df["cmp_api_key"]
		conf.Password = df["cmp_api_secret"]
	}
	conf.VerifySSL = fConf.VerifySSL
	return conf
}
