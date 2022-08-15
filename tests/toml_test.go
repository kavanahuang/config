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

type Cfg struct {
	Log tomlCfg
}

type tomlCfg struct {
	Relative bool   // true: Default use relative path, false: use absolute path.
	Dir      string // Log storage directory.
	Name     string // Log filename.
	Prefix   string // Prefix for cutting log filename.
	Level    string // The log level, if the value is OFF, then log off.
}

func TestTomlStart(t *testing.T) {
	test.Test.Start("toml")
}

func TestNewStructToml(t *testing.T) {
	var cfg Cfg
	msg := "NewStructToml: Relative is true: "
	cfgFile := "config.toml"
	config.Toml.NewStructToml("config", cfgFile, &cfg)
	relative := cfg.Log.Relative
	level := cfg.Log.Level

	if cfg.Log.Relative {
		test.Test.T(t).Logs(msg).Ok(relative)
	} else {
		test.Test.T(t).Logs(msg).No(relative)
	}

	msg = "NewStructToml: Level is INFO: "
	assert := "INFO"
	if level == assert {
		test.Test.T(t).Logs(msg).Ok(level)
	} else {
		test.Test.T(t).Logs(msg).No(level)
	}
}

func TestGet(t *testing.T) {
	msg := "Get: Relative is true: "
	cfgFile := "config.toml"
	newToml := config.Toml.NewToml("config", cfgFile)
	relative := newToml.Zone("log").Get("relative").AtBool()
	level := newToml.Zone("log").Get("level").AtStr()

	if relative {
		test.Test.T(t).Logs(msg).Ok(relative)
	} else {
		test.Test.T(t).Logs(msg).No(relative)
	}

	msg = "Get: Level is INFO: "
	assert := "INFO"
	if level == assert {
		test.Test.T(t).Logs(msg).Ok(level)
	} else {
		test.Test.T(t).Logs(msg).No(level)
	}
}

func TestGetTo(t *testing.T) {
	msg := "GetTo: Relative is true: "
	cfgFile := "config.toml"
	newToml := config.Toml.NewToml("config", cfgFile)
	relative := newToml.Zone("log").Get("relative").To()
	level := newToml.Zone("log").Get("level").To()

	if relative.(bool) {
		test.Test.T(t).Logs(msg).Ok(relative)
	} else {
		test.Test.T(t).Logs(msg).No(relative)
	}

	msg = "GetTo: Level is INFO: "
	assert := "INFO"
	if level.(string) == assert {
		test.Test.T(t).Logs(msg).Ok(level)
	} else {
		test.Test.T(t).Logs(msg).No(level)
	}
}

func TestFetch(t *testing.T) {
	msg := "Fetch: Relative is true: "
	cfgFile := "config.toml"
	newToml := config.Toml.NewToml("config", cfgFile)
	relative := newToml.Zone("log").Fetch("relative").ToBool()
	level := newToml.Zone("log").Fetch("level").ToStr()

	if relative {
		test.Test.T(t).Logs(msg).Ok(relative)
	} else {
		test.Test.T(t).Logs(msg).No(relative)
	}

	msg = "Fetch: Level is INFO: "
	assert := "INFO"
	if level == assert {
		test.Test.T(t).Logs(msg).Ok(level)
	} else {
		test.Test.T(t).Logs(msg).No(level)
	}
}

func TestTomlEnd(t *testing.T) {
	test.Test.End("toml")
}
