# **Apache Kafka Theory**

![image](https://user-images.githubusercontent.com/64280930/226430739-e72a3529-0781-4683-aa8d-b8cb8f4a1412.png)

**Zookeeper:**

Zookeeper is a distributed orchestration service that uses Apache Kafka to manage and orchestrate your cluster.
In a Kafka cluster, Zookeeper is responsible for the following functions:

1. Cluster management: Zookeeper maintains a list of all brokers in the Kafka cluster and their current status (for example, alive, dead, or unavailable).
2. Leader Election: Zookeeper manages the election of a leader for each partition in the Kafka cluster. The leader is responsible for handling all read and write requests for that partition (REMEMBER: Kafka manages a replication factor for partitions, if one dies the other is ready).
3. Topic configuration: Zookeeper stores the configuration information for each topic in the Kafka cluster, including the number of partitions and the replication factor.
4. Consumer Group Management: Zookeeper tracks the offsets for messages that each consumer group has consumed from each partition. This allows consumers to continue reading messages from where they left off in case of failure.
5. Broker registration – When a new broker is added to the Kafka cluster, it is registered with Zookeeper. This allows other brokers and clients to discover and communicate with you.

**Topic:**

Logical concept for grouping multiple partitions within a logical business unit.

**Partition:**

Kafka's topics are divided into several **partitions**
. While the topic is a logical concept in Kafka, a partition is the smallest storage unit that ***holds a subset of records owned by a topic***.
Each partition is a single log file where records are written to it in an append-only fashion.

**Partitions are the way that Kafka provides scalability**

A Kafka cluster is made of one or more servers. In the Kafka universe, they are called Brokers. Each broker holds a subset of records that belongs to the entire cluster.

Kafka distributes the partitions of a particular topic across multiple brokers. By doing so, we’ll get the following benefits.

- If we are to put all partitions of a topic in a single broker, the scalability of that topic will be constrained by the broker’s IO throughput. A topic will never get bigger than the biggest machine in the cluster. By spreading partitions across multiple brokers, a single topic can be scaled horizontally to provide performance far beyond a single broker’s ability.
- A single topic can be consumed by multiple consumers in parallel. Serving all partitions from a single broker limits the number of consumers it can support. Partitions on multiple brokers enable more consumers.
- Multiple instances of the same consumer can connect to partitions on different brokers, allowing very high message processing throughput. Each consumer instance will be served by one partition, ensuring that each record has a clear processing owner.

**Message:**
- TopicPartition:
A topic name and partition number, involves Offset of current message.
```
Type TopicPartition struct {
    Topic     *string
    Partition int32
    Offset    Offset
    Metadata  *string
    Error     error
}
```
- Value:
Type []byte, passed value of a message.

- Key:
Type []byte, Keys are mostly useful/necessary if you require strong order for a key and are developing something like a state machine. If you require that messages with the same key (for instance, a unique id) are always seen in the correct order, attaching a key to messages will ensure messages with the same key always go to the same partition in a topic. Kafka guarantees order within a partition, but not across partitions in a topic, so alternatively not providing a key - which will result in round-robin distribution across partitions - will not maintain such order.

- Timestamp:
Type time.Time.
```
2023-03-20 14:25:26.984 -0500 -05

```
- TimestampType:
The timestamp type of the records.

- Headers:
Additional data for the message.

![image](https://user-images.githubusercontent.com/64280930/226462847-973d0693-efd9-474e-a285-7addb9050bd1.png)

**Example of consumed message**

![image](https://user-images.githubusercontent.com/64280930/226445901-b1ff7571-bcb5-4ad2-bb64-be43daa98e26.png)


# Commands
- Create a topic
```
kafka-topics --create --bootstrap-server localhost:9092  --replication-factor 1 --partitions 2 --topic message-log
```
- List topics
```
kafka-topics --bootstrap-server localhost:9092 --list
```
- Delete a topic
```
kafka-topics --bootstrap-server localhost:9092 --delete --topic message-log
```

