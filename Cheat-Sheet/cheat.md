# Creating topic
`docker compose exec kafka kafka-topics --create --topic test_topic --partitions 1 --replication-factor 1 --bootstrap-server kafka:9092`

# produce message
`docker compose exec kafka kafka-console-producer --topic test_topic --broker-list kafka:9092`

# consume message
`docker compose exec kafka kafka-console-consumer --topic test_topic --from-beginning --bootstrap-server kafka:9092`