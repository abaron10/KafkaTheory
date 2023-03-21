## Single broker - Double Partition
![image](https://user-images.githubusercontent.com/64280930/226721244-65e78726-d853-49cd-8de6-8e2ba230adf9.png)


Create a topic with the following configuration:
```
```
kafka-topics --create --bootstrap-server localhost:9092  --replication-factor 1 --partitions 2 --topic message-log
```
```

After this we can decide if a message goes to partition 0 or partition 1.
```
	// Change partition attribute from 0 to desired partition, otherwise set it to kafka.PartitionAny
	// if round Robin is desired.
	err = p.p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.topic, Partition: 1},
		Key:            []byte(uuid.New().String()),
		Value:          messageBytes,
	}, deliveryChan)

```

From the point of view of the consumer we can have two options for consuming the partitions:

1. if your consumer group can read from any partition
```
	  if err := kc.SubscribeTopics(topics, nil); err != nil {
		  log.Errorf("Error subscribing to topics: %v (%v)", err, topics)
		  kc.Close()
		  panic(err)
	  }
```
2. If consumer group must read from specific partition (To guarentee order for example)

```
		topicPartition := kafka.TopicPartition{Topic: &topics[0], Partition: 1}
		err = kc.Assign([]kafka.TopicPartition{topicPartition})
		if err != nil {
      kc.Close()
			panic(err)
		}
```

Partitioning is a key feature of Apache Kafka, and it provides several benefits for building distributed, scalable, and fault-tolerant messaging systems. Here are some of the benefits of partitioning in Kafka:

Scalability: Partitioning allows Kafka to scale horizontally by distributing message data across multiple brokers, which can process messages in parallel. As the number of brokers and partitions increases, Kafka can handle larger volumes of messages and support more consumers.

Availability: By replicating data across multiple brokers, Kafka can provide high availability in the event of broker failure or network partitions. If one broker fails, another broker can take over serving requests for the partitions hosted on the failed broker.

Performance: Partitioning can improve message throughput and reduce message latency by allowing Kafka to process messages in parallel. Consumers can also read from multiple partitions concurrently, further improving performance.

Retention and data management: Kafka can store large volumes of data across multiple partitions, and each partition can be configured with its own retention policy. This makes it easier to manage data and maintain message history for longer periods of time.

Flexibility: Partitioning allows Kafka to support different use cases and workloads, such as event streaming, log aggregation, and messaging. Different partitions can be configured with different replication factors, retention policies, and message processing requirements.
