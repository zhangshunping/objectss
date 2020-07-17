/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (

	cfgFile string
	ObjectStorgeLink string
	Sqllimmit int  // sql limit
	Sqldays  int   // 查询多少天
	ChannelCap int // channel的容量大小
	ConsumerNum int //comusers Numebr
	SqlConnect string // mysql 链接方式

)


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "objectss",
	Short: "upload files to cloud oss",
	Long: ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
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

	rootCmd.PersistentFlags().StringVarP(&ObjectStorgeLink,"ObjectStorgeLink","l","oss://educoder.tmp","oss/obs link (-o oss://educoder.tmp )")
	rootCmd.PersistentFlags().IntVarP(&Sqllimmit,"sqlLimits","n",1000,"sql limit nums (-n 1000) ")
	rootCmd.PersistentFlags().IntVarP(&Sqldays,"sqldays","d",-15,"select data from mysql 15 days ago (-d -15) ")
	rootCmd.PersistentFlags().StringVar(&SqlConnect, "sqlcon", "root:123456789@tcp(127.0.0.1:3306)/gitlab", "connect sql (default is $HOME/.objectss.yaml) ")
	rootCmd.PersistentFlags().IntVarP(&ConsumerNum,"consusmerNum","s",100,"Run the number of comsumer goroutines (-s 100)")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.objectss.yaml) ")

	rootCmd.PersistentFlags().IntVarP(&ChannelCap,"ChannelCap","c",10,"channle cap (-c 10)")
	// Cobra also supports local flags, which will only ru
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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

		// Search config in home directory with name ".objectss" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".objectss")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}



