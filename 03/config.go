package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// pod defines deployment configuration
type deployment struct {
	Metadata struct {
		Name    string  `yaml:"name" json:"name"`
		Version float32 `yaml:"version" json:"version"`
	} `yaml:"metadata" json:"metadata"`
	Type  string `yaml:"type" json:"type"`
	Ports []struct {
		Name          string `yaml:"name" json:"name"`
		HostPort      string `yaml:"hostPort" json:"hostPort"`
		ContainerPort string `yaml:"containerPort" json:"containerPort"`
	} `yaml:"ports" json:"ports"`
	Env []string `yaml:"env" json:"env"`
}

// pod defines Pod configuration
type pod struct {
	Replicas float32 `yaml:"replicas" json:"replicas"`
	Health   struct {
		Wait     float32 `yaml:"wait" json:"wait"`
		Interval float32 `yaml:"interval" json:"interval"`
		Method   string  `yaml:"method" json:"method"`
		Path     string  `yaml:"path" json:"path"`
	} `yaml:"health" json:"health"`
}

// Config defines the user configuration
type Config struct {
	Deployment deployment `yaml:"deployment" json:"deployment"`
	Pod        pod        `yaml:"pod" json:"pod"`
}

// LoadConfig reads file bytes and parses it into configuration
func LoadConfig(fileName string, config *Config) error {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	return yaml.Unmarshal(data, &config)
}
