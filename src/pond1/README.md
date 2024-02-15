# Iotsim

## Iotsim em Ação!
Para simular um sensor de radiação solar utilizando o protocolo MQTT, basta seguir o tutorial a seguir a partir do diretório `pond1`.

### Configuração de um broker MQTT
Primeiramente, é necessário configurar um broker MQTT. Isso pode ser feito inicializando o arquivo de configuração `mosquitto.conf` com o comando a seguir:

```
mosquitto -c mosquitto.conf
```

### 
A segunda etapa consiste na ativação do publisher do sensor, e o subscriber para o recebimento de dados, que devem ser feitos rodando o seguinte comando nos diretórios `/publisher` e `/subscriber`.

```
go run .
```