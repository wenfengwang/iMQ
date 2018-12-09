# infinite Message Queue
iMQ is a message service Cloud-Native oriented and Topic-based, it is unconcerned what's storage engine is used for iMQ
through abstraction of storage module. so, you can use any storage you want such as TiKV, S3, HDFS etc by implementing
Storage API. So far iMQ only support TiKV and TiKV is default engine, more engines may be implemented in the future.
iMQ was born for scale-out free in any dimensions, which means you can increase or decrease your iMQ cluster size to reach
performance you want. In iMQ, Topic is a logic concept, it consist of multiple queues that distributed in multiple broker.
Queue is similar to Partition of Kafka or Queue of RocketMQ, but the different is the queue of iMQ can be consumed by
multiple clients and Kafka or RocketMQ can't. Currently, iMQ at an early stage, if you interested in iMQ, any
contribution are very welcome.

## Roadmap
- Normal Message: DOING
- Delay Message : TODO
- Ordered Message: TODO
- Transaction Message: TODO
- OPS Easier: TODO
- Monitor: TODO
- Integration with Kubernetes: TODO
- Authorization: TODO
- Multiple Language Clients: TODO
- MQTT Support: TODO
- Integration with BigData System: TODO
- Function & Streaming: TODO
