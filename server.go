package main

import (
	"flag"
	"fmt"
	"gggo/controller"
	"gggo/model"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gorilla/context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	logFileName = flag.String("log", "run/server.log", "Log file name") //日志
)

func init() {
	//生成,更新日志
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "cServer start Failed")
		os.Exit(1)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	//连接数据库
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)
	//开启路由
	controller.Startup()
	//监听
	http.ListenAndServe(":8888", context.ClearHandler(http.DefaultServeMux))

}

// func watchApp() {
// 	if *autoWatch == "true" {
// 		autoAppName := appName + ".auto"
// 		if err := copyApp(); err != nil {
// 			log.Fatal(err)
// 		}

// 		if err := exec.Command("./"+autoAppName, "-auto", "watch", "-dir", *confDir).Run(); err != nil {
// 			log.Fatal(err)
// 		}
// 	} else {
// 		//Watcher Begin
// 		goFileWatcher, err := fsnotify.NewWatcher()
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		defer goFileWatcher.Close()

// 		// teplFileWatcher, err := fsnotify.NewWatcher()
// 		// if err != nil {
// 		// 	log.Fatal(err)
// 		// }

// 		// defer teplFileWatcher.Close()
// 		if err := goFileWatcher.Add("./static"); err != nil {
// 			log.Warn("watch dir:", "./static", " error:", err)
// 		}
// 		//Watcher End

// 		cmd := startApp()
// 		planToDoTime := time.Now()

// 		for {
// 			select {
// 			case ev, ok := <-goFileWatcher.Events:
// 				if ok {
// 					log.Debug("goFileWatcher event:", ev)
// 					if ev.Op&fsnotify.Chmod != fsnotify.Chmod {
// 						if time.Since(planToDoTime) > *deferToAuto {
// 							time.AfterFunc((*deferToAuto + 2500*time.Millisecond), func() {
// 								if time.Since(planToDoTime) > (*deferToAuto + 2500*time.Millisecond) {
// 									if installApp() == nil {
// 										if err := cmd.Process.Kill(); err == nil {
// 											cmd = startApp()
// 										}
// 									} else {
// 										log.Debug("cmd.Process.Kill error:", err)
// 									}
// 								}
// 							})
// 							planToDoTime = time.Now()
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// }

// func copyApp() error {
// 	src := appName
// 	dst := appName + ".auto"
// 	in, err := os.Open(src)
// 	if err != nil {
// 		return err
// 	}
// 	defer in.Close()
// 	tmp, err := ioutil.TempFile(filepath.Dir(dst), "autofile")
// 	if err != nil {
// 		return err
// 	}
// 	_, err = io.Copy(tmp, in)
// 	if err != nil {
// 		tmp.Close()
// 		os.Remove(tmp.Name())
// 		return err
// 	}
// 	if err := tmp.Close(); err != nil {
// 		os.Remove(tmp.Name())
// 		return err
// 	}
// 	const perm = 0744
// 	if err := os.Chmod(tmp.Name(), perm); err != nil {
// 		os.Remove(tmp.Name())
// 		return err
// 	}
// 	if err := os.Rename(tmp.Name(), dst); err != nil {
// 		os.Remove(tmp.Name())
// 		return err
// 	}

// 	return nil
// }

// func installApp() error {
// 	cmd := exec.Command("go", "build")
// 	outp, err := cmd.CombinedOutput()
// 	log.Debug("Rebuild " + appName)
// 	log.Debug("Error: ", err)
// 	log.Debug("OutPut: " + string(outp))
// 	return err
// }

// func startApp() *exec.Cmd {
// 	cmd := exec.Command("./"+appName, "-dir", *confDir)
// 	var b bytes.Buffer
// 	cmd.Stdout = &b
// 	cmd.Stderr = &b
// 	err := cmd.Start()
// 	log.Debug("Error: ", err)
// 	log.Debug("OutPut: " + b.String())
// 	// if err == nil {
// 	// 	pid := cmd.Process.Pid
// 	// 	ioutil.WriteFile(filepath.Clean(conf.GetCDPath()+conf.TempDirPath+appName+".pid"), []byte(strconv.Itoa(pid)), 0700)
// 	// 	log.Debug("Start "+appName+" - pid:", pid)
// 	// }

// 	return cmd
// }
