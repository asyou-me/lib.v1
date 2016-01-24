package kafka

type Producer struct {
	ZookeeperNodes []string
	Group          string
	Topics         string
}

func (k *KafkaProducer) Push() {

}
