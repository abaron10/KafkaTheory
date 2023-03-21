# Single Broker - Single Partition
![image](https://user-images.githubusercontent.com/64280930/226711541-46043984-63f1-408b-ad25-d375a26dba1d.png)

Since there is only one partition even if you set kafka.PartitionAny on message config, your event will be routed to Partition 0. If you set manually a greater number, an error will be thrown to the program.


### Benefits:

- Simplicity: This configuration is the simplest possible Kafka setup, requiring only a single broker and a single partition.
- Low latency: Messages can be produced and consumed with low latency because there is no need to coordinate between multiple brokers or partitions.
- High throughput: A single partition can handle a high volume of messages, and a single broker can handle a large number of partitions.

### Limitations:

- Scalability: A single broker can handle only a limited number of partitions and may become a bottleneck as the volume of messages and partitions increases.
- Fault tolerance: A single broker is a single point of failure, so if the broker fails, the entire Kafka cluster becomes unavailable. In addition, if the single partition becomes corrupt or unavailable, messages cannot be produced or consumed until the partition is restored.
- Limited replication: With a single partition, there is no replication of data to other brokers, which can result in data loss if the broker fails or if the partition becomes corrupted.

# Notes:
### - Using deliveryChan:
  In Apache Kafka, a delivery channel (also known as a delivery handler or delivery callback) is a mechanism used to notify the producer of the status of a produced message.

  When a producer sends a message to a Kafka broker, it does not immediately know whether the message was successfully written to the broker or not. The delivery channel allows the producer to register a callback function that will be called by the broker once the message has been successfully written to a Kafka topic or if there was an error writing the message.

  The delivery channel callback function can be used by the producer to perform actions based on the status of the message, such as logging success or failure, updating internal state, or retrying failed messages.

  The delivery channel is particularly useful when producing messages with a high level of reliability or when dealing with large volumes of messages. By using the delivery channel, the producer can ensure that messages are being delivered correctly and can take corrective action if necessary.

  It's worth noting that the delivery channel is not a required feature of Apache Kafka. Producers can choose not to use it and simply assume that messages are successfully written to the broker unless an error is thrown. However, for many use cases, using the delivery channel can provide greater reliability and flexibility.
