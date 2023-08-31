package config

import (
	"time"

	"go.uber.org/fx"

	"github.com/abdivasiyev/microservice_template/pkg/config/viper"
)

var FxOption = fx.Provide(
	fx.Annotate(viper.New, fx.As(new(Config))),
)

type Config interface {
	Get(key string) any
	GetBool(key string) bool
	GetFloat64(key string) float64
	GetInt(key string) int
	GetUInt32(key string) uint32
	GetUInt8(key string) uint8
	GetIntSlice(key string) []int
	GetString(key string) string
	GetStringMap(key string) map[string]any
	GetStringMapString(key string) map[string]string
	UnmarshalKey(key string, val any) error
	GetStringSlice(key string) []string
	GetDuration(key string) time.Duration
}
