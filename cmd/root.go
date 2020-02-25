package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/JayceChant/drivel/pkg/confuse"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	// confuse bool
	marker     string
	filePath   string
	useMartian bool
	overwrite  bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "drivel [text]",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. `,
	Args: cobra.MaximumNArgs(1),
	Run:  run,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVarP(&marker, "marker", "k", "@", "set target sentence marker")
	rootCmd.Flags().StringVarP(&filePath, "file", "f", "", "read input from file given")
	rootCmd.Flags().BoolVarP(&useMartian, "martian", "m", false, "enable martian mode")
	rootCmd.Flags().BoolVarP(&overwrite, "overwrite", "o", false, "overwrite result to file instead of printing; will be ignored if --file not specified")
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

func run(cmd *cobra.Command, args []string) {
	var text string
	if filePath != "" {
		b, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		text = string(b)
	} else {
		text = strings.Join(args, "")
	}

	text = confuse.Trans(text, marker, 2, useMartian)

	if overwrite && filePath != "" {
		ioutil.WriteFile(filePath, []byte(text), 0666)
		return
	}

	fmt.Println(text)
}
