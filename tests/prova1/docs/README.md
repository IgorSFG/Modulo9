## Como rodar

### Configuração das variaveis de ambiente
Primeiramente, para acessar o binário do Go, rode o seguinte comando:
```
source .bashrc
```

### Configuração de um broker MQTT
Em seguida, é necessário configurar um broker MQTT. Isso pode ser feito inicializando o arquivo de configuração `mosquitto.conf` com o comando a seguir:

```
mosquitto -c mosquitto.conf
```

### Ativação do Publisher & Subscriber
Por último, falta apenas ativar o publisher do sensor, e o subscriber para o recebimento de dados, que devem ser feitos rodando o seguinte comando nos diretórios `/publisher` e `/subscriber`.

```
go run .
```

### Utilizando testes
Para testar as funçoes do publisher e do subscriber, rode o seguinte comando nos diretórios `/publisher` e `/subscriber`.
```
go test -v
```

## Em Ação!
Você pode conferir imagens do funcionamento a seguir:

(https://github.com/IgorSFG/Modulo9/tree/main/tests/prova1/docs)