# IotsimHive
IotsimHive representa uma integração com o simulador
[Iotsim](https://github.com/IgorSFG/Modulo9/tree/main/src/pond1)
com um cluster configurado na HiveMQ com a utilização da autenticação em camada de transporte (TLS).

## Integrando com IotsimHive!
Para enviar os dados do simulador de sensor de radiação solar para um cluster, basta seguir o tutorial a partir do diretório `pond4`.

### Configuração das variaveis de ambiente
Primeiramente, para acessar o binário do Go, rode o seguinte comando:
```
source .bashrc
```

Em seguida, crie um arquivo `.env` com as suas credenciais:
```
BROKER_ADDR = "379d67d20bd940f2921461046040735b.s1.eu.hivemq.cloud"
HIVE_USER = "<seu-nome-de-usuario-aqui>"
HIVE_PSWD = "<sua-senha-cadastrada-aqui>"
```

### Publicando e Subscrevendo no HiveMQ
Finalmente, chega a hora de integrar o publisher do sensor, e o subscriber para o recebimento de dados, ao cluster do HiveMQ. Rode o seguinte comando nos diretórios `/publisher` e `/subscriber`.
```
go run .
```

## IotsimHive em Ação!
Você pode conferir o funcionamento de IotsimHive no vídeo a seguir:

https://drive.google.com/file/d/1_arLLDNqxZm74vxMnuy2q5TbwlMzR4Wa/view?usp=sharing