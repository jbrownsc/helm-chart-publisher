package config

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const APP = "helm-charts-publisher"

func init() {
	initViper()
}

func initViper() {
  viper.SetDefault("base_url", "http://localhost:8080")
	// viper.SetConfigName("config")        // name of config file (without extension)
	// viper.AddConfigPath("/etc/" + APP)   // path to look for the config file in
	// viper.AddConfigPath("$HOME/." + APP) // call multiple times to add many search paths
	// viper.AddConfigPath(".")             // optionally look for config in the working directory
	// err := viper.ReadInConfig()          // Find and read the config file
	// if err != nil {                      // Handle errors reading the config file
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }
}

type Config map[string]interface{}

func GetStorage() (string, Config) {
	storageConfig := viper.GetStringMap("storage")

	var storages []string
	// Return only key in this map
	for k := range storageConfig {
		storages = append(storages, k)
	}

	if len(storages) != 1 {
		fmt.Println("Error: multiple or none storage specified in configuration")
		os.Exit(1)
	}

	return storages[0], viper.GetStringMap("storage." + storages[0])
}

func GetRepos() interface{} {
	return viper.Get("repos")
}

func (c *Config) GetBaseUrl() string {
	return viper.GetString("base_url")
}

func ReadConfigFile(configFile string) error {
	if configFile == "" {
		return errors.New("no configuration file specified")
	}

	file, err := os.Open(configFile)
	if err != nil {
		return errors.Wrap(err, "open config file failed")
	}

	viper.SetConfigType("yaml")
	err = viper.ReadConfig(file)
	if err != nil {
		return errors.Wrap(err, "read config file failed")
	}

	return nil
}
