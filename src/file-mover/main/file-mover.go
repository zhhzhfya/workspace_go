package main

import (
	"log"
	"encoding/xml"
	"io/ioutil"
	"fmt"
	"os"
	"io"
	"time"
)

type SInfos struct {
	Infos     xml.Name `xml:"infos"`      //
	SleepTime int      `xml:"sleep_time"` //
	InfoList  []Info   `xml:"info"`       //
}
type Info struct {
	Infos   xml.Name `xml:"info"`
	Source  string   `xml:"source"`
	Desc    string   `xml:"desc"`
	SubDir  string   `xml:"sub_dir"`
	IsTmp   bool     `xml:"istmp"`
	AddType string   `xml:"addType"`
}

var run = false

func main() {
	run = true
	log.SetFlags(log.Lshortfile | log.LstdFlags);
	log.Println("app begin");
	content, err := ioutil.ReadFile("src/file-mover/conf/info_dev.xml")

	if err != nil {
		log.Fatal(err)
	}
	infos := SInfos{}
	err = xml.Unmarshal(content, &infos)
	if err != nil {
		log.Fatal(err)
	}
	sleeptime := infos.SleepTime;
	log.Println("默认休眠时间:%d", sleeptime)
	scaners := make(chan string, 10)

	for _, info := range infos.InfoList {
		log.Println(info)
		// 每个info节点起一个协程
		go func() {
			// 扫描目录
			rd, err := ioutil.ReadDir(info.Source)
			if err != nil {
				fmt.Println(err)
			}
			for _, fi := range rd {
				if fi.IsDir() {
					log.Printf("[%s]\n", fi.Name())
				} else {
					scaners <- info.Source + "/" + fi.Name()
					log.Println("sended " + info.Source + "/" + fi.Name())
				}
			}
			time.Sleep(time.Millisecond * 500);
			log.Println("scan over")

		}()
		go func() {
			select {
			case fileName := <-scaners:
				log.Println("收到数据" + fileName)
				//CopyFile
			}
			log.Println("收 end")
		}()
	}

	for {
		if !run {
			break
		}
		time.Sleep(time.Millisecond * 500);
	}
	log.Println("app over")
}

func CopyFile(dst, src string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close()
	dstFile, err := os.Open(dst)
	if err != nil {
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}
