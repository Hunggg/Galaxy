package config

/**
 * Config
 */
type Config struct {
	ID               uint64 `gorm:"primaryKey"`
	LastFetchedBlock uint64
}

type ConfigField string

const (
	// ConfigFieldMetronionBaseURI ConfigField = "metronion_base_uri"
	ConfigFieldLastFetchedBlock ConfigField = "last_fetched_block"
)
