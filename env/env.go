package env

type ContractConfig struct {
	Address        string `mapstructure:"address"`
	CreatedAtBlock uint64 `mapstructure:"created_at_block"`
}

type DatabaseConfig struct {
	Database string `mapstructure:"database"`
}

type HttpConfig struct {
	EnableCors bool `mapstructure:"enable_cors"`
}

type MetronionNFTConfig struct {
	BaseURI  string `mapstructure:"base_uri"`
	ImageURI string `mapstructure:"image_uri"`
}

type Config struct {
	ChainId                        uint64             `mapstructure:"chain_id"`
	ChainName                      string             `mapstructure:"chain_name"`
	NodeUrl                        string             `mapstructure:"node_url"`
	WsNodeUrl                      string             `mapstructure:"ws_node_url"`
	MetronionAccessoriesContract   ContractConfig     `mapstructure:"metronion_accessories_contract"`
	MetronionNFTContract           ContractConfig     `mapstructure:"metronion_nft_contract"`
	MetronionSaleContract          ContractConfig     `mapstructure:"metronion_sale_contract"`
	MetroGalaxyMarketplaceContract ContractConfig     `mapstructure:"metrogalaxy_marketplace_contract"`
	MetronionNFTConfig             MetronionNFTConfig `mapstructure:"metronion_nft"`
	Postgres                       DatabaseConfig     `mapstructure:"postgres"`
	Http                           HttpConfig         `mapstructure:"http"`
}

// maintain one single instance of config
var (
	config *Config
)

func GetConfig() Config {
	if config == nil {
		return Config{}
	}
	return *config
}

func InitConfig(cfg *Config) {
	config = cfg
}
