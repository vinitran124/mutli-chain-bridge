package config_test

import (
	"bridge/config"
	"flag"
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
	"os"
	"reflect"
	"strings"
	"testing"
)

func Test_Defaults(t *testing.T) {
	tcs := []struct {
		path          string
		expectedValue interface{}
	}{
		{
			path:          "Environment",
			expectedValue: config.LogEnvironment("development"),
		},
		{
			path:          "Database.User",
			expectedValue: "postgres",
		},
		{
			path:          "Database.Password",
			expectedValue: "123456",
		},
		{
			path:          "Database.Name",
			expectedValue: "state_db",
		},
		{
			path:          "Database.Host",
			expectedValue: "postgres",
		},
		{
			path:          "Database.Port",
			expectedValue: "5432",
		},
	}

	file, err := os.CreateTemp("", "genesisConfig")
	require.NoError(t, err)
	defer func() {
		require.NoError(t, os.Remove(file.Name()))
	}()
	require.NoError(t, os.WriteFile(file.Name(), []byte("{}"), 0600))

	flagSet := flag.NewFlagSet("", flag.PanicOnError)
	flagSet.String(config.FlagEnvironment, "development", "")
	ctx := cli.NewContext(cli.NewApp(), flagSet, nil)
	cfg, err := config.Load(ctx)
	if err != nil {
		t.Fatalf("Unexpected error loading default config: %v", err)
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.path, func(t *testing.T) {
			fmt.Println(cfg)
			actual := getValueFromStruct(tc.path, cfg)
			require.Equal(t, tc.expectedValue, actual)
		})
	}
}

func getValueFromStruct(path string, object interface{}) interface{} {
	keySlice := strings.Split(path, ".")
	v := reflect.ValueOf(object)

	for _, key := range keySlice {
		for v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		v = v.FieldByName(key)
	}
	return v.Interface()
}
