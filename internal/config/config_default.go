// Package config ...
package config

import "strings"

// DefaultConfig struct for default config
type DefaultConfig struct {
	configData map[string]interface{}
}

func (c *DefaultConfig) get(name string) (result interface{}, found bool) {
	data := c.configData

	for _, key := range strings.Split(name, ":") {
		result, found = data[key]
		if newSection, ok := result.(map[string]interface{}); ok && found {
			data = newSection
		} else {
			return
		}
	}

	return
}

// GetSection found and return section in configurations file
func (c *DefaultConfig) GetSection(name string) (section Configuration, found bool) {
	value, found := c.get(name)

	if found {
		if sectionData, ok := value.(map[string]interface{}); ok {
			section = &DefaultConfig{configData: sectionData}
		}
	}

	return section, found
}

// GetString found and return value in string
func (c *DefaultConfig) GetString(name string) (result string, found bool) {
	value, found := c.get(name)

	if found {
		result = value.(string)
	}

	return result, found
}

// GetInt found and return value in integer
func (c *DefaultConfig) GetInt(name string) (result int, found bool) {
	value, found := c.get(name)

	if found {
		result = int(value.(float64))
	}
	return result, found
}

// GetBool found and return value in booling
func (c *DefaultConfig) GetBool(name string) (result bool, found bool) {
	value, found := c.get(name)

	if found {
		result = value.(bool)
	}

	return result, found
}

// GetFloat found and return value in float64
func (c *DefaultConfig) GetFloat(name string) (result float64, found bool) {
	value, found := c.get(name)

	if found {
		result = value.(float64)
	}

	return result, found
}
