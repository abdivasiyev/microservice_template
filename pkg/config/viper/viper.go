package viper

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	cfg *viper.Viper
}

func New() (*Config, error) {
	cfg := viper.New()
	cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	cfg.AutomaticEnv()
	cfg.WatchConfig()

	return &Config{cfg: cfg}, nil
}

func (c *Config) Get(key string) any {
	return c.cfg.Get(key)
}

func (c *Config) GetBool(key string) bool {
	return c.cfg.GetBool(key)
}

func (c *Config) GetFloat64(key string) float64 {
	return c.cfg.GetFloat64(key)
}

func (c *Config) GetInt(key string) int {
	return c.cfg.GetInt(key)
}

func (c *Config) GetUInt32(key string) uint32 {
	return c.cfg.GetUint32(key)
}

func (c *Config) GetUInt8(key string) uint8 {
	return uint8(c.cfg.GetUint(key))
}

func (c *Config) GetIntSlice(key string) []int {
	return c.cfg.GetIntSlice(key)
}

func (c *Config) GetString(key string) string {
	return c.cfg.GetString(key)
}

func (c *Config) GetStringSlice(key string) []string {
	return c.cfg.GetStringSlice(key)
}

func (c *Config) GetStringMap(key string) map[string]any {
	return c.cfg.GetStringMap(key)
}
func (c *Config) GetStringMapString(key string) map[string]string {
	return c.cfg.GetStringMapString(key)
}

func (c *Config) UnmarshalKey(key string, val any) error {
	return c.cfg.UnmarshalKey(key, &val)
}

func (c *Config) GetDuration(key string) time.Duration {
	return c.cfg.GetDuration(key)
}
