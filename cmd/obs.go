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
	"objectss/database"
	"objectss/handler"
	"objectss/utils"
	"sync"
	"time"
)

// obsCmd represents the obs command
var obsCmd = &cobra.Command{
	Use:   "obs",
	Short: "huawei cloud obs",
	Long: ``,
	Example: "## 迁移git仓库到obs\n " +
		"objectss  obs uploadgit",
	Run: func(cmd *cobra.Command, args []string) {

		cmd.Help()
		fmt.Println(cmd.Example)
		return
	},
}

func init() {
	rootCmd.AddCommand(obsCmd)
	obsCmd.AddCommand(uploadgitCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// obsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// obsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}



var uploadgitCmd = &cobra.Command{
	Use:   "uploadgit cmd ",
	Short: "huawei cloud obs upaloadgit ",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		startT:=time.Now()
		utils.Log.Infof("Staring,开始迁移任务.")
		wg := new(sync.WaitGroup)   //需要值引用
		var ch = make(chan string, ChannelCap)
		// 初始化databse
		dbw := database.Initmysql(Sqllimmit, SqlConnect)
		// 默认查询查询15天前的tpis
		dbw.QuerData(Sqldays, Sqllimmit)
		wgNums := len(dbw.Repositories)
		go handler.Productor(ch, dbw)
		wg.Add(wgNums)
		for i := 0; i < ConsumerNum; i++ {
			go handler.Consumer(ch, dbw, wg, ObjectStorgeLink)
		}
		wg.Wait()
		eT := time.Since(startT)
		utils.Log.Infof("Ending,消费者消费的gitpath个数为:%d,耗时：%v", handler.ComsuNum,eT)
	},
}
