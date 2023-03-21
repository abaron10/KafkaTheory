## Single broker - Double partition - Manual commit

This scenario handles two consumers asociated with the same group, one that fails the polling of the message and another that consumes the message after the 
sessionTimeoutMsKey finishes.

```
const (
	maxPollInterval      = 10800000
	sessionTimeout       = 60000
	bootstrapServersKey  = "bootstrap.servers"
	groupIdKey           = "group.id"
	maxPollIntervalMsKey = "max.poll.interval.ms"
	sessionTimeoutMsKey  = "session.timeout.ms"
	autoCommit           = "enable.auto.commit" // 
)

	config = &kafka.ConfigMap{
		bootstrapServersKey:  host,
		groupIdKey:           groupID,
		maxPollIntervalMsKey: maxPollInterval,
		// amount of time in ms for a message to be retained by a consumer, if the consumer does not commit the message
		// kafka release it so other consumers can use it.
		sessionTimeoutMsKey: sessionTimeout,
    // If false, manual commit must be done in order to release messages from the partition.
		autoCommit:          false,
	}

```


From the consumer see that if the message is processed correctly, then we commit the message to delete it from the topic, so other consumers don't poll it.
```
func (ki *KafkaInput) Poll() chan []byte {

	go func() {
		for {
			ev := ki.Consumer.Consumer.Poll(1000)
			switch e := ev.(type) {
			case *kafka.Message:
				// Process the message
				messageProcessedSuccessfully := ki.processMessage(e.Value)
				if messageProcessedSuccessfully {
					ki.Consumer.Consumer.CommitMessage(e)
					ki.InputChannel <- e.Value
				}

				data, _ := ki.Consumer.Consumer.GetConsumerGroupMetadata()
				fmt.Printf("Consumer instance ID: %s\n", data)

			case kafka.Error:
				log.Errorf("%v consume_message_error in topics[%s], error: %v\n", e.Code(), strings.Join(ki.Consumer.Topics, ","), e)
			default:
			}
		}
	}()

	return ki.InputChannel

}
```

## When to use autocommit?
- Control over message processing: With auto-commit enabled, a consumer will automatically commit its current offset position to the broker after consuming a message. This can make it difficult to control how messages are processed, especially if the processing of messages is complex and requires more time. By disabling auto-commit, the consumer can have more control over when to commit the offset position, allowing it to process messages more carefully.

- Exactly-once message processing: If your application requires exactly-once message processing semantics, disabling auto-commit is necessary. With auto-commit enabled, it's possible for a consumer to process a message but not commit the offset, leading to duplicate processing of messages in certain scenarios. To achieve exactly-once semantics, the consumer must explicitly commit the offset position after processing the message.

- Limited availability of messages: In some cases, messages may not be readily available in the Kafka broker, and consumers may need to wait for messages to be produced before they can be consumed. In such cases, auto-commit may lead to overconsumption of messages, as the consumer may commit its offset position even if no messages were actually consumed. Disabling auto-commit can help avoid overconsumption of messages and improve overall efficiency.

- Large batch processing: If your application processes messages in large batches, it may be more efficient to disable auto-commit and commit offsets after processing a batch of messages. This can help reduce the number of network requests made to the Kafka broker, leading to improved performance.
