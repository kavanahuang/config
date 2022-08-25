/*
 Author: Kernel.Huang
 Mail: kernelman79@gmail.com
 File: config.go
 Date: 8/9/22 1:57 PM
*/
package config

import (
	"github.com/kavanahuang/common"
	"github.com/kavanahuang/system"
	"gopkg.in/yaml.v3"
	"log"
)

type YamlConfig struct {
	cfg []byte
}

var Yaml = new(YamlConfig)

func (yc *YamlConfig) NewYaml(dirname string, filename string) *YamlConfig {
	name := system.GetFilepath(dirname, filename)
	if common.Files.PathExists(name) {
		yc.cfg = common.Files.GetFile(name)
	}

	return yc
}

/*
Example:
	var structured structuredConfig
	config.Yaml.NewYaml("config", "config.toml").Parse(&structured)
*/
func (yc *YamlConfig) Parse(structured any) any {
	err := yaml.Unmarshal(yc.cfg, structured)
	if err != nil {
		log.Println("Parse yaml file error: ", err)
	}

	return structured
}
