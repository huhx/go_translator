package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string
var history bool
var cleanUp bool

var rootCmd = &cobra.Command{
	Use:   "trans",
	Short: "Translate text from cn to en and force",
	Long:  `A command line tool that translate chinese to english and vice versa, it also provides the feature of management the translation history.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 && !history && !cleanUp {
			_, _ = fmt.Fprintln(os.Stderr, "Usage: trans [-l -c] <text>")
			os.Exit(1)
		}

		if history {
			ListRecords()
		} else if cleanUp {
			ClearRecords()
		} else {
			Translate(args[0])
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go_translator.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&history, "history", "l", false, "list the history translation")
	rootCmd.PersistentFlags().BoolVarP(&cleanUp, "cleanUp", "c", false, "cleanUp the history translation")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".go_translator")
	}

	viper.AutomaticEnv()
	_ = viper.ReadInConfig()
}
