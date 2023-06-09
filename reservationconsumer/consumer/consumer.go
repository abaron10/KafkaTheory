package consumer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"
	"strings"
)

type (
	Offset int64

	KafkaInput struct {
		Consumer     *KafkaConsumer
		InputChannel chan []byte
	}

	KafkaConsumer struct {
		Consumer        *kafka.Consumer
		BootstrapServer string
		GroupID         string
		Topics          []string
	}

	ConsumerConfig struct {
		Host    string
		Topics  []string
		GroupId string
	}
)

const (
	maxPollInterval      = 10800000
	sessionTimeout       = 60000
	bootstrapServersKey  = "bootstrap.servers"
	groupIdKey           = "group.id"
	maxPollIntervalMsKey = "max.poll.interval.ms"
	sessionTimeoutMsKey  = "session.timeout.ms"
)

func NewConsumer(config ConsumerConfig) *KafkaInput {
	var (
		consumer     = NewKafkaConsumer(config.Host, config.Topics, config.GroupId)
		inputChannel = make(chan []byte, 1)
	)
	return &KafkaInput{Consumer: consumer, InputChannel: inputChannel}
}

func NewKafkaConsumer(host string, topics []string, groupID string) *KafkaConsumer {
	var config *kafka.ConfigMap
	config = &kafka.ConfigMap{
		bootstrapServersKey:  host,
		groupIdKey:           groupID,
		maxPollIntervalMsKey: maxPollInterval,
		sessionTimeoutMsKey:  sessionTimeout,
	}

	log.Infof("Setting up ConfigMap %v for topic cpgs-elastic-executor-eventmanager", config)

	kc, err := kafka.NewConsumer(config)
	if err != nil {
		log.Errorf("Error creating a reservationconsumer: error=%v, bootstrap.server=%v, group.id=%v", err, host, groupID)
		kc.Close()
		panic(err)
	}

	if err := kc.SubscribeTopics(topics, nil); err != nil {
		log.Errorf("Error subscribing to topics: %v (%v)", err, topics)
		kc.Close()
		panic(err)
	}

	return &KafkaConsumer{
		Consumer:        kc,
		BootstrapServer: host,
		GroupID:         groupID,
		Topics:          topics,
	}
}

func (ki *KafkaInput) Poll() chan []byte {

	go func() {
		for {
			ev := ki.Consumer.Consumer.Poll(1000)
			switch e := ev.(type) {
			case *kafka.Message:
				ki.InputChannel <- e.Value
			case kafka.Error:
				log.Errorf("%v consume_message_error in topics[%s], error: %v\n", e.Code(), strings.Join(ki.Consumer.Topics, ","), e)
			default:
			}
		}
	}()

	return ki.InputChannel

}

func (ki *KafkaInput) Close() {
	ki.Consumer.Consumer.Close()
}
