/*
 Author: Kernel.Huang
 Mail: kernelman79@gmail.com
 File: config.go
 Date: 8/9/22 1:57 PM
*/
package config

import (
	"github.com/BurntSushi/toml"
	"github.com/kavanahuang/common"
	"github.com/kavanahuang/log"
	goToml "github.com/pelletier/go-toml"
)

type tomlConfig struct {
	keyName    string
	value      interface{}
	Structured interface{}
	cfg        *goToml.Tree
}

var Toml = new(tomlConfig)

func (tf *tomlConfig) NewToml(dirname string, filename string) *tomlConfig {
	name := common.GetCustomConfigPath(dirname, filename)
	conf, err := goToml.LoadFile(name)

	if err != nil {
		log.Logs.Error("Load toml file error: ", err)
	}

	tf.cfg = conf
	return tf
}

// Example: result := Tome.NewToml(dirname, filename).Zone("zoneName").Get("key").To()
func (tf *tomlConfig) Zone(key string) *tomlConfig {
	tf.keyName = key
	return tf
}

// Example: result := Tome.NewToml(dirname, filename).Zone("zoneName").Get("key").To()
func (tf *tomlConfig) Get(key string) *tomlConfig {
	tf.keyName = tf.keyName + "." + key
	return tf
}

/**
 * Example: result := Tome.NewToml(dirname, filename).Zone("zoneName").Get("key").To()
 */
func (tf *tomlConfig) To() interface{} {
	return tf.cfg.Get(tf.keyName)
}

// Example: result := Tome.NewToml(dirname, filename).Zone("zoneName").Get("key").AtStr()
func (tf *tomlConfig) AtStr() string {
	tf.value = tf.cfg.Get(tf.keyName)
	return tf.value.(string)
}

// Example: result := Tome.NewToml(dirname, filename).Zone("zoneName").Get("key").AtInt()
func (tf *tomlConfig) AtInt() int {
	tf.value = tf.cfg.Get(tf.keyName)
	return tf.value.(int)
}

// Example: result := Tome.NewToml(dirname, filename).Zone("zoneName").Get("key").AtBool()
func (tf *tomlConfig) AtBool() bool {
	tf.value = tf.cfg.Get(tf.keyName)
	return tf.value.(bool)
}

// Example: result := Tome.NewToml(dirname, filename).Zone("zoneName").Fetch("key").ToStr()
func (tf *tomlConfig) Fetch(key string) *tomlConfig {
	tf.keyName = tf.keyName + "." + key
	tf.value = tf.cfg.Get(tf.keyName)
	return tf
}

// Example: result := Tome.NewToml(dirname, filename).Read("zoneName.key").ToStr() or ToInt()
func (tf *tomlConfig) Read(key string) *tomlConfig {
	tf.keyName = key
	tf.value = tf.To()
	return tf
}

// Example: result := Tome.NewToml(dirname, filename).Read("zoneName.key").ToStr()
func (tf *tomlConfig) ToStr() string {
	return tf.value.(string)
}

// Example: result := Tome.NewToml(dirname, filename).Read("zoneName.key").ToInt()
func (tf *tomlConfig) ToInt() int {
	return tf.value.(int)
}

// Example: result := Tome.NewToml(dirname, filename).Read("zoneName.key").ToBool()
func (tf *tomlConfig) ToBool() bool {
	return tf.value.(bool)
}

/*
Example:
	var structured structuredConfig
	config.Toml.NewStructToml("config", "config.toml", &structured)
*/
func (tf *tomlConfig) NewStructToml(dirname string, filename string, structured any) any {

	path := common.GetCustomConfigPath(dirname, filename)
	_, err := toml.DecodeFile(path, structured)
	if err != nil {
		log.Logs.Error("Load or decode toml file error: ", err)
	}

	return structured
}
