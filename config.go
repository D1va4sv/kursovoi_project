package config

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

type CompressionConfig struct {
    Enabled bool `yaml:"enabled"`
    Level   int  `yaml:"level"`
}

type Config struct {
    Compression CompressionConfig `yaml:"compression"`
}

var ConfigData Config

func LoadConfig(configPath string) error {
    data, err := ioutil.ReadFile(configPath)
    if err != nil {
        return err
    }
    err = yaml.Unmarshal(data, &ConfigData)
    if err != nil {
        return err
    }
    return nil
}

func init() {
    err := LoadConfig("config.yaml")
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }
}
