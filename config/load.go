package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	env "github.com/knadh/koanf/providers/env/v2"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"

	"strings"
)

func Load(configPath string) Config {
	k := koanf.New(".")

	// Load default values
	k.Load(confmap.Provider(defaultConfig, "."), nil)

	// Load YAML
	k.Load(file.Provider(configPath), yaml.Parser())

	// Load environment variables
	k.Load(
		env.Provider(
			".", // delimiter for nested keys
			env.Opt{
				Prefix: "GAMEAPP_",
				TransformFunc: func(key, val string) (string, any) {
					// remove prefix, lower-case, replace "_" with "."
					k := strings.ToLower(strings.TrimPrefix(key, "GAMEAPP_"))
					k = strings.ReplaceAll(k, "_", ".")
					return k, val
				},
			},
		),
		nil,
	)

	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		panic(err)
	}

	return cfg
}
