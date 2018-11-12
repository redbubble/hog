package cmd

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"os"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// Variables to hold flag values
var target string
var port int
var limit int
var timeoutMs int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hog",
	Short: "Takes all the connections. Doesn't give them back.",
	Long:  `Hog is a testing tool for finding how many simultaneous TCP connections a service will accept.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		if viper.GetBool("version") {
			versionCmd()
			os.Exit(0)
		}

		if target == "" {
			return errors.New("--target is missing")
		}

		fmt.Printf("Testing port %d on %s with %d simultaneous connections.\n", port, target, limit)

		conns := make([]net.Conn, limit)
		results := make([]bool, limit)

		for i := 0; i < limit; i++ {
			conns[i], err = net.DialTimeout("tcp", fmt.Sprintf("%s:%d", target, port), time.Duration(timeoutMs)*time.Millisecond)

			if err != nil {
				results[i] = false
			} else {
				results[i] = true
				defer conns[i].Close()
			}
		}

		success, fail := 0, 0

		for i := 0; i < limit; i++ {
			if results[i] {
				success++
			} else {
				fail++
			}
		}

		fmt.Printf("Made %d connections to %s:%d. %d successful, %d failed.\n", limit, target, port)
		return nil
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is $HOME/.hog.yaml)")
	rootCmd.PersistentFlags().Bool("version", false, "Print the current version and exit")
	viper.BindPFlag("version", rootCmd.PersistentFlags().Lookup("version"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&target, "target", "t", "", "Hostname or IP address of the target service")
	rootCmd.Flags().IntVarP(&port, "port", "p", 80, "TCP port of the target service")
	rootCmd.Flags().IntVarP(&limit, "limit", "l", 100, "Maximum number of simultaneous connections to attempt")
	rootCmd.Flags().IntVar(&timeoutMs, "timeout-ms", 500, "Connection timeout in milliseconds")
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

		// Search config in home directory with name ".hog" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".hog")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func versionCmd() {
	fmt.Printf("hog v%s\n", viper.GetString("hog.version"))

	yabytes, _ := base64.StdEncoding.DecodeString(`
ICAgICAgICAgICAgICBfLC0iIiIiLS4uX18KICAgICAgICAgfGAsLSdfLiBg
ICBgIGBgICBgLS0nIiIiLgogICAgICAgICA7ICAsJyAgfCBgYCAgYCBgICBg
IGBgYCAgYC4KICAgICAgICwtJyAgIC4uLScgYCBgIGBgIGAgIGBgIGAgIGAg
fD09LgogICAgICwnICAgIF4gICAgYCAgYCAgICBgYCBgICBgIGAuICA7ICAg
XAogICAgYH1fLC1eLSAgIF8gLiAgYCBcIGAgIGAgX18gYCAgIDsgICAgIwog
ICAgICAgYCItLS0iJyBgLWAuIGAgXC0tLSIiYC5gLiAgYDsKICAgICAgICAg
ICAgICAgICAgXFxgIDsgICAgICAgOyBgLiBgLAogICAgICAgICAgICAgICAg
ICAgfHxgOyAgICAgIC8gLyB8IHwKICAgICAgICAgICAgICAgICAgLy9fO2Ag
ICAgLF87JyAsXzsiCg==`)
	var yascii = string(yabytes)
	fmt.Printf("\n%s\n", yascii)
}
