package kafka

import (
	"context"
	"fmt"
	"sync"
	"tank-api/app/entity/logs"

	"github.com/IBM/sarama"
)

// const version = sarama.ParseKafkaVersion("0.0.1")

type KafkaLogger struct {
	producer *producerProvider
	topic    string
}

func NewKafkaLogger(version sarama.KafkaVersion, topic string, brokers ...string) *KafkaLogger {
	producer := newProducerProvider(brokers, func() *sarama.Config {
		config := sarama.NewConfig()
		config.Version = version
		config.Producer.Idempotent = true
		config.Producer.Return.Errors = false
		config.Producer.RequiredAcks = sarama.WaitForAll
		config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
		config.Producer.Transaction.Retry.Backoff = 10
		config.Producer.Transaction.ID = "txn_producer"
		config.Net.MaxOpenRequests = 1
		return config
	})
	return &KafkaLogger{
		producer: producer,
		topic:    topic,
	}
}

func (logger *KafkaLogger) Context(ctx context.Context) logs.ILogger {
	return &KafkaLogger{}
}

func (logger *KafkaLogger) Error(message string) {
	producer := logger.producer.borrow()
	defer logger.producer.release(producer)
	producer.Input() <- &sarama.ProducerMessage{Topic: logger.topic, Value: sarama.StringEncoder(fmt.Sprintf("[ERROR] %s\n", message))}
}

func (logger *KafkaLogger) Fatal(message string) {
	producer := logger.producer.borrow()
	defer logger.producer.release(producer)
	producer.Input() <- &sarama.ProducerMessage{Topic: logger.topic, Value: sarama.StringEncoder(fmt.Sprintf("[FATAL] %s\n", message))}
}

func (logger *KafkaLogger) Info(message string) {
	producer := logger.producer.borrow()
	defer logger.producer.release(producer)
	producer.Input() <- &sarama.ProducerMessage{Topic: logger.topic, Value: sarama.StringEncoder(fmt.Sprintf("[INFO] %s\n", message))}
}

// ************************************
// TAKEN FROM SARAMA EXAMPLE
// https://github.com/IBM/sarama/blob/main/examples/txn_producer/main.go
// ***********************************
type producerProvider struct {
	transactionIdGenerator int32

	producersLock sync.Mutex
	producers     []sarama.AsyncProducer

	producerProvider func() sarama.AsyncProducer
}

func newProducerProvider(brokers []string, producerConfigurationProvider func() *sarama.Config) *producerProvider {
	provider := &producerProvider{}
	provider.producerProvider = func() sarama.AsyncProducer {
		config := producerConfigurationProvider()
		suffix := provider.transactionIdGenerator
		if config.Producer.Transaction.ID != "" {
			provider.transactionIdGenerator++
			config.Producer.Transaction.ID = config.Producer.Transaction.ID + "-" + fmt.Sprint(suffix)
		}
		producer, err := sarama.NewAsyncProducer(brokers, config)
		if err != nil {
			return nil
		}
		return producer
	}
	return provider
}

func (p *producerProvider) borrow() (producer sarama.AsyncProducer) {
	p.producersLock.Lock()
	defer p.producersLock.Unlock()

	if len(p.producers) == 0 {
		for {
			producer = p.producerProvider()
			if producer != nil {
				return
			}
		}
	}

	index := len(p.producers) - 1
	producer = p.producers[index]
	p.producers = p.producers[:index]
	return
}

func (p *producerProvider) release(producer sarama.AsyncProducer) {
	p.producersLock.Lock()
	defer p.producersLock.Unlock()

	// If released producer is erroneous close it and don't return it to the producer pool.
	if producer.TxnStatus()&sarama.ProducerTxnFlagInError != 0 {
		// Try to close it
		_ = producer.Close()
		return
	}
	p.producers = append(p.producers, producer)
}
