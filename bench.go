package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var (
	client    *redis.Client
	workerCnt = 1
)

func initLog() {
	logname := time.Now().Format("2006-01-02-03_04_05.log")
	file, _ := os.OpenFile(filepath.Join("logs", logname), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	mw := io.MultiWriter(os.Stdout, file)
	logrus.SetOutput(mw)

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
}

func init() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	initLog()
}

func setNow(id int) {
	key := fmt.Sprintf("key_data/%v", id)
	now := time.Now().UnixNano()
	_, err := client.Set(context.Background(), key, now, 0).Result()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"key": key,
			"err": err,
			"id":  id,
		}).Info()
	}
}

func worker(id int, wg *sync.WaitGroup, stop chan struct{}) {
	tick := time.NewTicker(time.Millisecond * 1000)
OUT:
	for {
		select {
		case <-tick.C:
			setNow(id)
		case <-stop:
			wg.Done()
			break OUT
		}
	}
	tick.Stop()
	logrus.Infof("Stop worker %v\n", id)
}

func main() {
	var wg sync.WaitGroup
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	stop := make(chan struct{})
	for i := 0; i < workerCnt; i++ {
		wg.Add(1)
		go worker(i, &wg, stop)
	}

	select {
	case <-sigc:
		close(stop)
	}

	wg.Wait()
	logrus.Info("Stop all worker")
}
