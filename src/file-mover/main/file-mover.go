package main

import (
	"log"
	"encoding/xml"
	"io/ioutil"
	"os"
	"io"
	"time"
	"sync"
	"os/signal"
	"syscall"
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
	Sleep   int      `xml:"sleep"`
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

	wgMovers := sync.WaitGroup{}
	wgMovers.Add(len(infos.InfoList))

	dataCh := make(chan string, 100)
	stopCh := make(chan struct{})
	// 定义系统信号的channl
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs,
		os.Kill,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		sig := <-sigs
		log.Println("系统正在退出...：%d", sig)
		close(stopCh)
	}()

	for _, info := range infos.InfoList {
		log.Println(info)
		if info.Sleep == 0 {
			info.Sleep = sleeptime
		}
		// 每个info节点起一个协程
		go func() {
			for {
				select {
				case <-stopCh:
					return
				default:
				}
				// 扫描目录
				rd, err := ioutil.ReadDir(info.Source)
				if err != nil {
					log.Println(err)
				}
				for _, fi := range rd {
					if fi.IsDir() {
						log.Printf("[%s]\n", fi.Name())
					} else {

						select {
						case <-stopCh:
							return
						case dataCh <- info.Source + "/" + fi.Name():
						}
					}
				}
				log.Println(info.Sleep)
				time.Sleep(time.Millisecond * time.Duration(info.Sleep));
			}
		}()
		go func() {
			defer wgMovers.Done()

			for {
				log.Println("for.....")
				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case fileName := <-dataCh:
					log.Println("收到数据" + fileName)
					//CopyFile
				}
				//log.Println("收 end")
			}
		}()
	}

	wgMovers.Wait()
	log.Println("file_mover正常退出")
}

///
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
