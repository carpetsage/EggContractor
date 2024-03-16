package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/carpetsage/EggContractor/config"
)

const _configFileEnvVar = "EGGCONTRACTOR_CONFIG_FILE"

var _config *config.Config

var _configCommand = &cobra.Command{
	Use:   "config",
	Short: "Print current configurations",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := initConfig(); err != nil {
			return err
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintf(w, "config file:\t%s\n", viper.ConfigFileUsed())
		fmt.Fprintf(w, "player.id:\t%s\n", _config.Player.Id)
		fmt.Fprintf(w, "player.device_id:\t%s\n", _config.Player.DeviceId)
		fmt.Fprintf(w, "database.path:\t%s\n", _config.Database.Path)
		fmt.Fprintf(w, "notification.pushover.on:\t%t\n", _config.Notification.Pushover.On)
		fmt.Fprintf(w, "notification.pushover.api_key:\t%s\n", _config.Notification.Pushover.APIKey)
		fmt.Fprintf(w, "notification.pushover.user_key:\t%s\n", _config.Notification.Pushover.UserKey)
		w.Flush()
		return nil
	},
}

var _configTemplateCommand = &cobra.Command{
	Use:   "config-template",
	Short: "Print a config file template",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(config.ConfigTemplate)
	},
}

func init() {
	_rootCmd.AddCommand(_configCommand)
	_rootCmd.AddCommand(_configTemplateCommand)
}

func initConfig() error {
	if _cfgFile != "" {
		viper.SetConfigFile(_cfgFile)
	} else if os.Getenv(_configFileEnvVar) != "" {
		viper.SetConfigFile(os.Getenv(_configFileEnvVar))
	} else {
		viper.AddConfigPath("$HOME/.config/" + config.ProgName)
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
	}
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	_config = &config.Config{}
	if err := viper.Unmarshal(_config); err != nil {
		return err
	}
	if err := _config.ResolveAndValidate(); err != nil {
		return errors.Wrap(err, "invalid config")
	}
	return nil
}
