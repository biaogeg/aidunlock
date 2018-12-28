// Copyright © 2018 mritd <mritd1234@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"

	"github.com/mritd/aidunlock/unlock"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const version = "1.0.2"

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "aidunlock",
	Short: "A simple Apple ID unlock tool.",
	Long: `
A simple Apple ID unlock tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		unlock.Boot()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is aidunlock.yaml)")
}

func initConfig() {

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		cfgFile = "aidunlock.yaml"
		viper.SetConfigFile(cfgFile)

		if _, err := os.Stat(cfgFile); err != nil {
			os.Create(cfgFile)
			viper.Set("AppleIDs", unlock.ExampleConfig())
			viper.Set("Email", unlock.SMTPExampleConfig())
			viper.WriteConfig()
		}

	}

	viper.AutomaticEnv()
	viper.ReadInConfig()

}
