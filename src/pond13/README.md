# IotsimMongoKafka
IotsimMongoKafka é uma atualização de
[IotsimMeta](https://github.com/IgorSFG/Modulo9/tree/main/src/pond6),
alterando o método de armazenamento para um banco de dados MongoDB e garantindo um índice sequencial com Kafka via Confluent.

## Consumindo e armazenando com IotsimMongoKafka!
Para consumir e armazenar dados de um simulador de sensor de radiação solar basta seguir o tutorial a partir do diretório `pond13`.

### Configuração das variaveis de ambiente
Primeiramente, para acessar o binário do Go, rode o seguinte comando:
```
source .bashrc
```

Em seguida, crie um arquivo `.env` na pasta `config` com as suas credenciais:
```
BROKER_ADDR = "379d67d20bd940f2921461046040735b.s1.eu.hivemq.cloud"
HIVE_USER = "<seu-nome-de-usuario-aqui>"
HIVE_PSWD = "<sua-senha-cadastrada-aqui>"

BOOTSTRAP_SERVERS = "pkc-ldjyd.southamerica-east1.gcp.confluent.cloud:9092"
CLUSTER_ID = "lkc-055wzp"
CLUSTER_NAME = "Cluster0"
API_KEY = "<sua-chave-api-aqui>"
API_SECRET = "<seu-segredo-api-aqui>"

MONGODB = "<sua-connection-string-aqui>"
```

### Publicando no HiveMQ e Consumindo no Kafka
Para o envio e recebimento de dados, rode o seguinte comando nos diretórios `/publisher` e `/subscriber`:
```
go run .
```

## IotsimMongoKafka em Ação!
Você pode conferir o funcionamento de IotsimMongoKafka no vídeo a seguir:

https://drive.google.com/file/d/1zsprxmdH0f5QDbrDzzIsZnYOp2D4mJzM/view?usp=sharing