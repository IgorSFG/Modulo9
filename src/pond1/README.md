# Iotsim
Iotsim é um simulador de dispositivos IoT utilizando o protocolo MQTT através do uso da biblioteca Eclipse Paho e a linguagem de programação Go.

## Simulando um Sensor de Radiação Solar
Para a simulação do dispositivo IoT, foi escolhido o
[Sensor de Radiação Solar](https://sigmasensors.com.br/produtos/sensor-de-radiacao-solar-sem-fio-hobonet-rxw-lib-900),
que tem sua faixa de medição variando de 0 a 1280 W/m² e com uma taxa de dados de até 250 kbps, o que possibilita uma transferência de até 22000 mensagens por segundo.

## Simulando com Iotsim!
Para simular um sensor de radiação solar, basta seguir o tutorial a partir do diretório `pond1`.

### Configuração das variaveis de ambiente
Primeiramente, para acessar o binário do Go, rode o seguinte comando:
```
source .bashrc
```

### Configuração de um broker MQTT
Primeiramente, é necessário configurar um broker MQTT. Isso pode ser feito inicializando o arquivo de configuração `mosquitto.conf` com o comando a seguir:

```
mosquitto -c mosquitto.conf
```

### Ativação do Publisher & Subscriber
A segunda etapa consiste na ativação do publisher do sensor, e o subscriber para o recebimento de dados, que devem ser feitos rodando o seguinte comando nos diretórios `/publisher` e `/subscriber`.

```
go run .
```

## Iotsim em Ação!
Você pode conferir o funcionamento de Iotsim no vídeo a seguir:

https://drive.google.com/file/d/1X4G-q1IAJhRuZhomQylVpPES22cRnxkC/view?usp=sharing