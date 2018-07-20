package util

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func ReadJson(filename string, def interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("ReadJson failed: %T %v", def, err)
	}
	if err = json.Unmarshal(data, def); err != nil {
		return fmt.Errorf("ReadJson failed: %T %v %v", def, filename, err)
	}
	return nil
}
