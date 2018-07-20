#环境配置

GO_HOME=C:\go
GOPATH=E:\workspace_go;C:\Users\lenovo\go

#dep 安装

 $ go get -u github.com/golang/dep/cmd/dep

# 确保所有本地状态-代码树、清单、锁和供应商彼此同步
 $ dep ensure

#
 $ dep ensure -add github.com/bitly/go-simplejson

#还可以指定依赖的版本：

 $ dep ensure -add github.com/bitly/go-simplejson@=0.4.3

#执行dep ensure 为了更好地看到过程，加上参数-v。

#执行dep init -gopath -v查看初始化过程。

#获取mysql 驱动下载到了GO_PATH下
#C:\Users\lenovo\go\src\github.com\go-sql-driver
#go get github.com/go-sql-driver/mysql