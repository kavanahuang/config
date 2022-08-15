/**
 * Created by IntelliJ IDEA.
 * User: kernel
 * Mail: kernelman79@gmail.com
 * Date: 2017/8/22
 * Time: 01:38
 */

package tests

import (
	"github.com/kavanahuang/config"
	"github.com/kavanahuang/test"
	"testing"
)

type yamlCfg struct {
	Relative bool   // true: Default use relative path, false: use absolute path.
	Dir      string // Log storage directory.
	Name     string // Log filename.
	Prefix   string // Prefix for cutting log filename.
	Level    string // The log level, if the value is OFF, then log off.
}

func TestYamlStart(t *testing.T) {
	test.Test.Start("yaml")
}

func TestNewYamlParse(t *testing.T) {
	var cfg yamlCfg
	msg := "NewYamlParse: Relative is true: "
	cfgFile := "config.yaml"
	config.Yaml.NewYaml("config", cfgFile).Parse(&cfg)
	relative := cfg.Relative
	level := cfg.Level

	if cfg.Relative {
		test.Test.New(t).Msg(msg).Ok(relative)
	} else {
		test.Test.New(t).Msg(msg).No(relative)
	}

	msg = "NewYamlParse: Level is INFO: "
	assert := "INFO"
	if level == assert {
		test.Test.New(t).Msg(msg).Ok(level)
	} else {
		test.Test.T(t).Msg(msg).No(level)
	}
}

func TestYamlEnd(New *testing.T) {
	test.Test.End("yaml")
}
