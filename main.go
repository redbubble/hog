package main

import (
	"github.com/spf13/viper"

	"github.com/redbubble/hog/cmd"
)

var version string

func main() {
	viper.Set("hog.version", version)
	cmd.Execute()
}
