package config

import (
	"fmt"
	"reflect"
	"testing"

	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

const envTestKey = "ENV_TEST_KEY"

func TestConfig_Get(t *testing.T) {
	testCases := []struct {
		name     string
		key      string
		value    any
		setToEnv bool
	}{
		{
			name:     "test with existing value",
			key:      envTestKey,
			value:    "test",
			setToEnv: true,
		},
		{
			name:     "test with non-existing value",
			key:      envTestKey,
			value:    nil,
			setToEnv: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setToEnv {
				t.Setenv(tc.key, fmt.Sprint(tc.value))
			}

			var cfg Config
			testApp := fxtest.New(t, FxOption, fx.Populate(&cfg))

			gotValue := cfg.Get(tc.key)
			if !reflect.DeepEqual(gotValue, tc.value) {
				t.Fatalf("expected: %v [%T], got: %v [%T]", tc.value, tc.value, gotValue, gotValue)
			}

			defer testApp.RequireStart().RequireStop()
		})
	}
}

func TestConfig_GetBool(t *testing.T) {
	testCases := []struct {
		name     string
		key      string
		value    bool
		setToEnv bool
	}{
		{
			name:     "test with existing value",
			key:      envTestKey,
			value:    true,
			setToEnv: true,
		},
		{
			name:     "test with non-existing value",
			key:      envTestKey,
			value:    false,
			setToEnv: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setToEnv {
				t.Setenv(tc.key, fmt.Sprint(tc.value))
			}

			var cfg Config
			testApp := fxtest.New(t, FxOption, fx.Populate(&cfg))

			gotValue := cfg.GetBool(tc.key)
			if !reflect.DeepEqual(gotValue, tc.value) {
				t.Fatalf("expected: %v [%T], got: %v [%T]", tc.value, tc.value, gotValue, gotValue)
			}

			defer testApp.RequireStart().RequireStop()
		})
	}
}

func TestConfig_GetFloat64(t *testing.T) {
	testCases := []struct {
		name     string
		key      string
		value    float64
		setToEnv bool
	}{
		{
			name:     "test with existing value",
			key:      envTestKey,
			value:    1.0,
			setToEnv: true,
		},
		{
			name:     "test with non-existing value",
			key:      envTestKey,
			value:    0,
			setToEnv: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setToEnv {
				t.Setenv(tc.key, fmt.Sprint(tc.value))
			}

			var cfg Config
			testApp := fxtest.New(t, FxOption, fx.Populate(&cfg))

			gotValue := cfg.GetFloat64(tc.key)
			if !reflect.DeepEqual(gotValue, tc.value) {
				t.Fatalf("expected: %v [%T], got: %v [%T]", tc.value, tc.value, gotValue, gotValue)
			}

			defer testApp.RequireStart().RequireStop()
		})
	}
}

func TestConfig_GetString(t *testing.T) {
	testCases := []struct {
		name     string
		key      string
		value    string
		setToEnv bool
	}{
		{
			name:     "test with existing value",
			key:      envTestKey,
			value:    "test",
			setToEnv: true,
		},
		{
			name:     "test with non-existing value",
			key:      envTestKey,
			value:    "",
			setToEnv: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setToEnv {
				t.Setenv(tc.key, fmt.Sprint(tc.value))
			}

			var cfg Config
			testApp := fxtest.New(t, FxOption, fx.Populate(&cfg))

			gotValue := cfg.GetString(tc.key)
			if !reflect.DeepEqual(gotValue, tc.value) {
				t.Fatalf("expected: %v [%T], got: %v [%T]", tc.value, tc.value, gotValue, gotValue)
			}

			defer testApp.RequireStart().RequireStop()
		})
	}
}
