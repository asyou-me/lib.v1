package log_client

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// 创建一个文档日志处理对象
func NewFileHandle(conf LogConf, log *Logger) (*FileHandle, error) {
	flog := FileHandle{
		Path: conf.Addr,
		Area: conf.Area,
		log:  log,
	}
	err := flog.Init()
	if err != nil {
		return nil, err
	}
	return &flog, nil
}

// 文档日志处理对象
type FileHandle struct {
	Path       string
	Out        *os.File
	log        *Logger
	Area       string
	errNum     int64
	num        int64
	NowDayTime int64
	// 读写锁
	mu sync.RWMutex
}

// 初始化日志写入对象
func (r *FileHandle) Init() error {

	NowTime := time.Now()
	r.NowDayTime = NowTime.Unix() - int64(NowTime.Hour()*3600) - int64(NowTime.Minute()*60) -
		int64(NowTime.Second())

	r.Cut(false)

	_, err := file_path_check(r.Path + "/" + r.Area)
	if err != nil {
		os.Mkdir(r.Path+"/"+r.Area, 0755)
	}

	out_file, logErr := os.OpenFile(r.Path+"/"+r.Area+"/log.json", os.O_CREATE|
		os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		return logErr
	}
	r.Out = out_file
	return nil
}

func (r *FileHandle) Cut(must bool) {

	fileinfo, err := file_path_check(r.Path + "/" + r.Area + "/log.json")
	if err != nil {
		return
	}
	mtime := fileinfo.ModTime()

	// 检查是否需要迁移数据
	if mtime.Unix() < r.NowDayTime || must {
		os.Rename(r.Path+"/"+r.Area+"/log.json", r.Path+"/"+r.Area+"/"+
			mtime.Format(FileTimeFormat)+".json")
		if must {
			out_file, logErr := os.OpenFile(r.Path+"/"+r.Area+"/log.json", os.O_CREATE|
				os.O_RDWR|os.O_APPEND, 0666)
			if logErr == nil {
				r.Out = out_file
			}
		}
		NowTime := time.Now()
		r.NowDayTime = NowTime.Unix() - int64(NowTime.Hour()*3600) - int64(NowTime.Minute()*60) -
			int64(NowTime.Second())
	}
}

// 文档日志对象健康检查
func (r *FileHandle) CheckHealth() bool {
	_, err := file_path_check(r.Path + "/" + r.Area + "/log.json")
	if err != nil {
		_base_log.WriteTo(&logErr{
			Level: "ERROR",
			Err:   err.Error(),
			Msg: "检查日志文档" + r.Path + "/" + r.Area + ".log无法使用(" +
				fmt.Sprintf("%d/%d", r.errNum, r.num) + ")",
			Time: time.Now().Unix(),
		})
		return false
	}
	_base_log.WriteTo(&logErr{
		Level: "INFO",
		Err:   "",
		Msg: "检查日志文档" + r.Path + "/" + r.Area + ".log可以使用(" +
			fmt.Sprintf("%d/%d", r.errNum, r.num) + ")",
		Time: time.Now().Unix(),
	})
	return true
}

// 文档日志处理句柄
func (l *FileHandle) WriteTo(msg LogBase) {
	var NowTime = time.Now().Unix()
	// 自动分片日志
	if NowTime-l.NowDayTime > 86400 {
		// 多线程写锁定,需要更新文件句柄
		l.mu.Lock()
		l.Cut(true)
		l.mu.Unlock()
	}

	msg.SetTime(NowTime)
	msgbyte := append(jsonFormat(msg), '\n')
	if l.log.PrintKey {
		fmt.Print(string(msgbyte))
	}
	reader := bytes.NewBuffer(msgbyte)

	// 多线程写锁定
	l.mu.Lock()
	_, err := io.Copy(l.Out, reader)
	l.num = l.num + 1
	l.mu.Unlock()

	reader = nil
	msgbyte = nil
	if err != nil {

		// 多线程写锁定
		l.mu.Lock()
		l.errNum = l.errNum + 1
		l.mu.Unlock()

		go func() {
			l.log.MsgChannel <- msg
			l.log.Err <- err
		}()
		return
	}
	msg = nil
}

// 文档日志处理句柄
func (l *FileHandle) RecoveryTo(msg string) {
	reader := bytes.NewBuffer(append([]byte(msg), '\n'))

	// 多线程写锁定
	l.mu.Lock()
	_, err := io.Copy(l.Out, reader)
	l.num = l.num + 1
	l.mu.Unlock()

	if err != nil {
		// 多线程写锁定
		l.mu.Lock()
		l.errNum = l.errNum + 1
		l.mu.Unlock()
		go func() {
			l.log.RecoveryChannel <- msg
			l.log.Err <- err
		}()
		return
	}
	msg = ""
}
