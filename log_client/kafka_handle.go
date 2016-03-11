package log_client

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strings"
	"sync"
	"time"
)

// 创建kafka处理对象
func NewKafkaHandle(lconf LogConf, log *Logger) (*KafkaHandle, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	addr_pre := strings.Replace(lconf.Addr, " ", "", -1)
	strings.Trim(",", addr_pre)
	addrs := strings.Split(addr_pre, ",")
	producer, err := sarama.NewSyncProducer(addrs, config)
	if err != nil {
		return nil, err
	}
	flog := KafkaHandle{
		Producer: producer,
		Topic:    lconf.Area,
		log:      log,
	}
	return &flog, nil
}

// kafka处理对象
type KafkaHandle struct {
	Producer sarama.SyncProducer
	Topic    string
	log      *Logger
	errNum   int64
	num      int64
	// 读写锁
	mu sync.RWMutex
}

// kafka服务健康检查
func (r *KafkaHandle) CheckHealth() bool {
	return true
}

// kafka处理句柄
func (r *KafkaHandle) WriteTo(msg LogBase) {
	NowTime := time.Now().Unix()
	msg.SetTime(NowTime)

	if r.log.PrintKey {
		fmt.Println(msg)
	}
	reader := jsonFormat(msg)
	message := &sarama.ProducerMessage{Topic: r.Topic, Key: sarama.StringEncoder("log"),
		Value: sarama.ByteEncoder(reader)}

	if _, _, err := r.Producer.SendMessage(message); err != nil {
		return
	}

	r.mu.Lock()
	r.num = r.num + 1
	r.mu.Unlock()
	if err != nil {
		r.mu.Lock()
		r.errNum = r.errNum + 1
		r.mu.Unlock()
		go func() {
			r.log.NewsChannel <- msg
			r.log.Err <- err
		}()
		return
	}
	reader = nil
	msg = nil
}

// kafka处理句柄
func (r *KafkaHandle) RecoveryTo(msg string) {

	msg = ""
}
