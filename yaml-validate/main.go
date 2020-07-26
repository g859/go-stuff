package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
)

type YamlFile struct {
	Update         bool              `yaml:"package_update"`
	Upgrade        bool              `yaml:"package_upgrade"`
	Reboot         bool              `yaml:"package_reboot_if_required"`
	PackageSources []Sources         `yaml:"apt_sources"`
	Packages       []string          `yaml:"packages,flow"`
	Files          []File            `yaml:"write_files"`
	Commands       []Commands        `yaml:"runcmd,flow"`
	PowerSettings  map[string]string `yaml:"power_state"`
}

type Sources struct {
	Source string `yaml:"source"`
}

type Commands []string

type File struct {
	Owner   string `yaml:"owner"`
	Path    string `yaml:"path"`
	Content string `yaml:"content`
}

func main() {
	fmt.Println("Parsing YAML file")

	var fileName string
	flag.StringVar(&fileName, "f", "", "YAML file to parse.")
	flag.Parse()

	if fileName == "" {
		fmt.Println("Please provide yaml file by using -f option")
		return
	}

	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return
	}

	var yamlConfig YamlFile
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}

	v := reflect.ValueOf(yamlConfig)

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%v\n", v.Field(i).Interface())
	}
}
