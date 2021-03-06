package plugin_config

import "encoding/json"

type PluginData struct {
	Plugins map[string]string
}

func NewData() *PluginData {
	return &PluginData{
		Plugins: make(map[string]string),
	}
}

func (pd *PluginData) JsonMarshalV3() (output []byte, err error) {
	return json.MarshalIndent(pd, "", "  ")
}

func (pd *PluginData) JsonUnmarshalV3(input []byte) (err error) {
	return json.Unmarshal(input, pd)
}
