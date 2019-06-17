package configutils

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"outstagram/server/pkg/network"
	"strings"

	"github.com/spf13/viper"
)

var (
	ConfigUrlEntry         = "configurl"
	RemoteConfigLogPrefix  = "[REMOTE CONFIG]"
	LocalConfigLogPrefix   = "[LOCAL CONFIG]"
	WarningLevel           = "[WARNING]"
	ErrorLevel             = "[ERROR]"
	InfoLevel              = "[INFO]"
	ConfigUrlRequiredError = errors.New("remote configs url is required")
)

func logConfigMessage(prefix, level, msg string) {
	log.Println(fmt.Sprintf("%s %s %s", prefix, level, msg))
}

func LoadRemoteConfiguration(url string) ([]byte, error) {
	if len(url) == 0 {
		return nil, ConfigUrlRequiredError
	}

	httpClient := network.NewClient()
	ctx := context.Background()
	resp := ConfigServiceResponse{}
	_, err := httpClient.Get(ctx, url, &resp)
	if err != nil {
		return nil, err
	}

	if resp.ReturnCode != 1 {
		return nil, fmt.Errorf("loading configuration from remote failed with StatusCode :%d", resp.ReturnCode)
	}

	logConfigMessage(
		RemoteConfigLogPrefix,
		InfoLevel,
		"Loaded configuration from remote successfully",
	)

	log.Printf(
		"[Configuration Info]: PrjectName: `%s`, ServiceName: `%s`, EnvName: `%s`",
		resp.Data.ProjectName,
		resp.Data.ServiceName,
		resp.Data.EnvName,
	)

	return []byte(resp.Data.ConfigJson), nil
}

func LoadConfiguration(serviceName, configFile, configPath string) {
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	configURL := viper.GetString(ConfigUrlEntry)
	configByte, err := LoadRemoteConfiguration(configURL)
	if err == nil {
		viper.SetConfigType("json")
		err = viper.ReadConfig(bytes.NewBuffer(configByte))
		if err == nil {
			logConfigMessage(
				RemoteConfigLogPrefix,
				InfoLevel,
				"Read configuration from remote successfully",
			)
			return
		}
	}

	logConfigMessage(
		RemoteConfigLogPrefix,
		ErrorLevel,
		err.Error(),
	)
	logConfigMessage(
		LocalConfigLogPrefix,
		WarningLevel,
		"Loading configuration from local file system as a fallback",
	)
	viper.SetConfigName(configFile)
	viper.AddConfigPath(configPath)
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", serviceName))
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		logConfigMessage(
			LocalConfigLogPrefix,
			ErrorLevel,
			"Can't load configuration from local file system",
		)
	}

	logConfigMessage(
		LocalConfigLogPrefix,
		InfoLevel,
		"Loaded configuration from local file system successfully",
	)
}
