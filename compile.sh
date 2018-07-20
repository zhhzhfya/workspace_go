#清理
go clean all
#需要维护的包，只需要列出含有执行文件的包名
MYSERVICE=("package1" "package2")
#生成本地运行文件
 for SERVER in ${MYSERVICE[*]}
        do
        echo "正在编译服务$SERVER mac版"
        go install $SERVER/...
 done
#生成linux运行文件
 for SERVER in ${MYSERVICE[*]}
        do
        echo "正在编译服务$SERVER linux版"
       GOOS=linux GOARCH=amd64 go install $SERVER/...
 done