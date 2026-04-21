### Install Strimzi

$ kubectl create namespace kafka

$ kubectl apply -f https://strimzi.io/install/latest?namespace=kafka -n kafka

### Deploy Kafka cluster


```bash
apiVersion: kafka.strimzi.io/v1beta2
kind: Kafka
metadata:
  name: my-cluster
  namespace: kafka
spec:
  kafka:
    replicas: 1
    listeners:
      - name: plain
        port: 9092
        type: internal
        tls: false
    storage:
      type: ephemeral
  zookeeper:
    replicas: 1
    storage:
      type: ephemeral
  entityOperator:
    topicOperator: {}
    userOperator: {}
```

### Create topics
```bash
$ kubectl exec -it -n kafka my-cluster-kafka-pool-0 -c kafka -- \
/opt/kafka/bin/kafka-topics.sh --create \
--topic upload-events \
--bootstrap-server localhost:9092 \
--partitions 1 \
--replication-factor 1
```
### Verify the topics created

```bash
$ kubectl exec -it -n kafka my-cluster-kafka-pool-0 -c kafka -- sh
sh-5.1$ /opt/kafka/bin/kafka-topics.sh --list --bootstrap-server localhost:9092
upload-events
sh-5.1$
exit
```