package inject

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func (c *Injector) ProvideRedisClient() *redis.Client {
	if c.redisClient == nil {
		addr := fmt.Sprintf("%v:%v", viper.GetString("METROGALAXY_API_REDIS_HOST"), viper.GetString("METROGALAXY_API_REDIS_PORT"))
		redisClient := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: viper.GetString("METROGALAXY_API_REDIS_PASSWORD"),
		})
		return redisClient
	}
	return c.redisClient
}