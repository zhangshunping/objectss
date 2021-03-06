package handler

import (
	"bytes"
	"fmt"
	"objectss/database"
	"objectss/utils"
	"os/exec"
	"sync"
)

var mu sync.Mutex
var ComsuNum int

// 生产者产生 shell的命令，放入管道中等待生产者执行
func Productor(channel chan string, dbw *database.DbWorker) {
	defer close(channel)
	for i := 0; i < len(dbw.Repositories); i++ {
		channel <- fmt.Sprintf("%s", dbw.Repositories[i].Path.String)
	}
}

// 消费者负责执行shell
func Consumer(channel chan string, dbw *database.DbWorker, wg *sync.WaitGroup, osslink string, obsCommand string) {
	for {
		path, ok := <-channel // 此处会阻塞, 如果信道中没有数据的话
		if ok {
			//取出命令执行
			if len(path) != 0 {
				//pathDir := strings.Replace(path, ".git", "", -1)
				//utils.Log.Info("path:", pathDir, "ComsumerNume:", ComsuNum)
				// 拷贝版本库到oss
				ossPath := fmt.Sprintf("%s%s", osslink, path)
				copyExec := fmt.Sprintf("%s %s %s", obsCommand, path, ossPath)

				fmt.Println("shell:", copyExec)
				s, err := Exec_linux_shell(copyExec)
				// 拷贝完成后，更新数据 oss=1
				if err == nil {
					dbw.UpdateRepositoryOssbyPath(path)
				} else {
					
					utils.Log.Errorf("对象存储命令行执行失败,错误为: %s ; 执行命令为： %s  ; 命令返回值: %s ",err,copyExec,s )

				}
			}

			mu.Lock()
			ComsuNum++
			mu.Unlock()
			wg.Done()
		}
	}
}

//执行shell 命令
func Exec_linux_shell(s string) (string, error) {
	//阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", s)

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out
	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()
	return out.String(), err

}
