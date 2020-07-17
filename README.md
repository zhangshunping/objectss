
#objectss

---

### 实现
#### 基于生产者多个消费者模型
- 生产者
    - 通过查询sql中记录的gitpath路径，放入到带缓冲区的管道中
    
- 消费者
    - 根据参数-s 指定的int值(默认10) 启用的消费者go程执行 obs/oss 迁移任务
    - 如果迁移成功，则修改数据中的flag oss=1证明已经迁移到cloud 对象存储上了，如果迁移命令失败
      则oss 仍然为0


### Help 
```go

$ go run main.go -h
upload files to cloud oss

Usage:
  objectss [command]

Available Commands:
  help        Help about any command
  obs         huawei cloud obs
  oss         aliyun object oss

Flags:
  -c, --ChannelCap int            channle cap (-c 10) (default 10)
  -l, --ObjectStorgeLink string   oss/obs link (-o oss://educoder.tmp ) (default "oss://educoder.tmp")
      --config string             config file (default is $HOME/.objectss.yaml)
  -s, --consusmerNum int          Run the number of comsumer goroutines (-s 100) (default 100)
  -h, --help                      help for objectss
  -n, --sqlLimits int             sql limit nums (-n 1000)  (default 1000)
      --sqlcon string             connect sql (default is $HOME/.objectss.yaml)  (default "root:123456789@tcp(127.0
.0.1:3306)/gitlab")
  -d, --sqldays int               select data from mysql 15 days ago (-d -15)  (default -15)
  -t, --toggle                    Help message for toggle

Use "objectss [command] --help" for more information about a command.

```


### 需补充
oss和obs存储迁移的 命令嵌入

