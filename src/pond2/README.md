# IotsimTest
IotsimTest é um conjunto de testes feitos para garantir a qualidade do simulador
[Iotsim](https://github.com/IgorSFG/Modulo9/tree/main/src/pond1).
Estes testes incluem: garantia da taxa de registro, validação do dado lido pelo sensor e o dado enviado para o broker e autenticação do recebibemento do dado de forma autêntica.

## Testando com IotsimTest!
Para testar o simulador de sensor de radiação solar, basta seguir o tutorial a partir do diretório `pond2`.

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

### Teste do Publisher & Subscriber
Enfim, chaga a hora de testar o publisher do sensor, e o subscriber para o recebimento de dados, que devem ser feitos rodando o seguinte comando nos diretórios `/sub_test` e `/pub_test` respectivamente.

```
go test -v
```

## IotsimTest em Ação!
Você pode conferir o funcionamento de IotsimTest no vídeo a seguir:

https://drive.google.com/file/d/1YW559eGErvN_-tbbt5KJLMb85I7fXqQR/view?usp=sharing