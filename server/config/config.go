package config

import(
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)	
// Capitalize the first letter of all variables! 
type Config_guard struct {
	Dbconf struct {
		Debug bool `yaml:"debug"`
		Db_type string `yaml:"db_type"`
		Max_lifetime uint64 `yaml:"max_lifetime"`
		Max_open_conns uint64 `yaml:"max_open_conns"`
		Max_idle_conns uint64 `yaml:"max_idle_conns"`
		Table_prefix string `yaml:"table_prefix"`
	}
	Sqlite3 struct {
		Path string `yaml:"path"`
	}
}
var Data Config_guard


func Parse_config(path string) error {
	conf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err2 := yaml.Unmarshal(conf, &Data)
	if err2!=nil{
		return err2
	}
	return nil
}