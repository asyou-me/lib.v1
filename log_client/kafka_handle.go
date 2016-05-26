package log_client

import (
	"strings"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/asyoume/lib/pulic_type"
)

// 创建kafka处理对象
func NewKafkaHandle(lconf LogConf, log *Logger) (*KafkaHandle, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.Flush.Frequency = 500 * time.Millisecond

	// 删除可能出现的空格
	addr_pre := strings.Replace(lconf.Addr, " ", "", -1)
	// 删除可能出现两边的，
	strings.Trim(",", addr_pre)
	// 通过[,]分割多个地址
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
func (r *KafkaHandle) WriteTo(msg pulic_type.LogBase) {
	NowTime := time.Now().Unix()
	msg.SetTime(NowTime)

	// 格式化数据到json
	reader := jsonFormat(msg)
	// 生成消息主体
	message := &sarama.ProducerMessage{Topic: r.Topic, Key: sarama.StringEncoder("log"),
		Value: sarama.ByteEncoder(reader)}
	// 发送消息到kafka
	_, _, err := r.Producer.SendMessage(message)

	r.mu.Lock()
	r.num = r.num + 1
	r.mu.Unlock()
	if err != nil {
		r.mu.Lock()
		r.errNum = r.errNum + 1
		r.mu.Unlock()
		go func() {
			r.log.MsgChannel <- msg
			r.log.Err <- err
		}()
		return
	}
	reader = nil
	msg = nil
}

// kafka处理句柄
func (r *KafkaHandle) RecoveryTo(msg string) {
	message := &sarama.ProducerMessage{Topic: r.Topic, Key: sarama.StringEncoder("log"),
		Value: sarama.StringEncoder(msg)}
	_, _, err := r.Producer.SendMessage(message)

	r.mu.Lock()
	r.num = r.num + 1
	r.mu.Unlock()
	if err != nil {
		r.mu.Lock()
		r.errNum = r.errNum + 1
		r.mu.Unlock()
		go func() {
			r.log.RecoveryChannel <- msg
			r.log.Err <- err
		}()
		return
	}
	msg = ""
}
