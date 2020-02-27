package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/JayceChant/drivel/pkg/confuse"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	marker      string
	filePath    string
	wordSegment bool
	useMartian  bool
	overwrite   bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "drivel [text]",
	Short: "A CLI tool that make your article like drivel",
	Long: `A tool that make your article confusing.
try 'drivel -f your_article.txt'`,
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
	rootCmd.Flags().BoolVarP(&wordSegment, "word", "w", false, "enable word segmentation mode for Chinese, which won't swap characters btween words")
	rootCmd.Flags().BoolVarP(&useMartian, "martian", "m", false, "enable martian(huoxingwen) mode")
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
			cmd.Println(err)
			os.Exit(1)
		}
		text = string(b)
	} else if len(args) == 1 {
		text = args[0]
	} else {
		cmd.Println("Error: neither text nor filepath is given")
		cmd.Printf("Run '%v --help' for usage.\n", cmd.CommandPath())
		os.Exit(1)
	}

	text, err := confuse.Trans(text, marker, wordSegment, useMartian)
	if err != nil {
		cmd.Println("Error:", err)
		os.Exit(1)
	}

	if overwrite && filePath != "" {
		ioutil.WriteFile(filePath, []byte(text), 0666)
		return
	}

	cmd.Println(text)
}
