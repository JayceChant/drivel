package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/JayceChant/drivel/pkg/martian"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile    string
	overwrite  bool
	useMartian bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "drivel",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. `,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("file path is not given")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		b, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		text := string(b)
		if useMartian {
			text = martian.Trans(text)
		}

		if overwrite {
			ioutil.WriteFile(args[0], []byte(text), 0666)
			return
		}

		fmt.Print(text)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolVarP(&overwrite, "overwrite", "w", false, "overwrite result to file instead of printing")
	rootCmd.Flags().BoolVarP(&useMartian, "martian", "m", false, "enable martian mode")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".drivel" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".drivel")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
