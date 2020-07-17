package database

import (
	"bytes"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"objectss/utils"
	"os/exec"
	"time"
)

// 定义query log 位置
var (
	err error
)

type DbWorker struct {
	Dsn          string
	Db           *sql.DB
	Repositories []Repositorytb //一万条记录
}

// 仓库数据
type Repositorytb struct {
	Id int
	// NullString代表一个为NULL的字符串
	Path sql.NullString
}

func Initmysql(RepositoriesLenth int, sqlconnect string) *DbWorker {

	Repositories:=make([]Repositorytb ,RepositoriesLenth)
	//初始化结构体，保存数据库参数
	Dbw := DbWorker{
		//Dsn: "root:123456789@tcp(127.0.0.1:3306)/gitlab",
		Dsn: sqlconnect,
		Repositories: Repositories,
	}
	//打开数据库,并保存给结构体内db
	Dbw.Db, err = sql.Open("mysql", Dbw.Dsn)
	if err != nil {
		utils.Log.Error(err.Error())
		panic(err.Error())
	}
	//如果打开失败，panic退出
	return &Dbw
}

// 查询数据
//d  负数表示多少天之前
func (dbw *DbWorker) QuerData(days int,sqlLimits int) {


	// 取出15天没更新的数据
	currentTime := time.Now()
	oldTime := currentTime.AddDate(0, 0, days).Format("2006-01-02 15:04:05")
	//取15天没更新的前10000条数据
	sql := fmt.Sprintf("SELECT id, path From repositories where oss = 0 and updated_at > '%s' limit %d", oldTime,sqlLimits)
	rows, err := dbw.Db.Query(sql)

	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	i := 0
	for rows.Next() {
		//取出数据库数据
		err = rows.Scan(&dbw.Repositories[i].Id, &dbw.Repositories[i].Path)
		if err != nil {
			fmt.Println(err.Error())
			utils.Log.Error(err.Error())
			continue
		}
		i++
	}
	err = rows.Err()
	if err != nil {
		utils.Log.Error("DatabaseQuery:",err.Error())
		fmt.Printf(err.Error())
	}

}

// 更新数据库
func (dbw *DbWorker) UpdateRepositoryOssbyPath(path string) {
	sql := fmt.Sprintf("UPDATE repositories SET oss = 1 where path = '%s'", path)
	_, err := dbw.Db.Exec(sql)
	if err != nil {
		utils.Log.Debugf("updateRepositoryOssbyPath error:", err, ", sql:", sql)
	}
}

//阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
func exec_shell(s string) (string, error) {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", s)

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()

	return out.String(), err
}
