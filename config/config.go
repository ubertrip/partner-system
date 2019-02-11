package config

import (
	"flag"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/spf13/viper"
)

type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Schema   string `json:"schema"`
}

type Config struct {
	Port     string        `json:"port"`
	Database Database      `json:"database"`
	Cookie   time.Duration `json:"cookieExperetionTime"`
}

var atomicConfig atomic.Value

func init() {
	env := flag.String("env", "dev", "environment (dev, prod, staging)")
	flag.Parse()

	viper.SetConfigType("json")
	viper.AddConfigPath("./env/")
	viper.SetConfigName(*env)
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("%v", err)
	} else {
		fmt.Println("Config loaded")
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		fmt.Printf("Unable to decode into config struct, %v", err)
	}

	atomicConfig.Store(config)
}

func Get() *Config {
	return atomicConfig.Load().(*Config)
}
