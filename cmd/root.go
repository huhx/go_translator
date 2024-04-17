package cmd

import (
	"fmt"
	"os"
	"trans/api"
	"trans/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "trans",
	Short: "Translate text from cn to en and force",
	Long:  `A command line tool that translate chinese to english and vice versa, it also provides the feature of querying the translation history.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			_, _ = fmt.Fprintln(os.Stderr, "Usage: trans <text>")
			os.Exit(1)
		}
		text := args[0]

		if util.IsChinese(text) {
			result := api.ToEnglish(text)
			println(result)
		} else {
			result := api.ToChinese(text)
			println(result)
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
