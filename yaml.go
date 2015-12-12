package main

import (
	"gopkg.in/yaml.v2"
)

func importFromYAML(data []byte) (*SplitsData, error) {
	splits := SplitsData{}
	err := yaml.Unmarshal(data, &splits)
	if err != nil {
		return nil, err
	}
	return &splits, nil
}

func exportToYAML(data *SplitsData) ([]byte, error) {
	marshalled, err := yaml.Marshal(data)
	if err != nil {
		return nil, err
	}
	return marshalled, nil
}
