package config

type PersistentDb interface {
	/**
	 * Config
	 */
	SaveConfig(config Config) error
	UpdateConfig(params map[ConfigField]interface{}) error
	GetConfig() (Config, error)
}
