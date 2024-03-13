# IotsimMeta
IotsimMeta é uma atualização de
[IotsimHive](https://github.com/IgorSFG/Modulo9/tree/main/src/pond4),
adicionando armazenamento em um banco de dados SQlite e visualização em uma dashboard pelo Metabase.

## Visualizando e armazenando com IotsimMeta!
Para armazenar os dados do simulador de sensor de radiação solar e visualizá-los em uma dashboard basta seguir o tutorial a partir do diretório `pond6`.

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

### API para gerenciamento de dados
Para que os dados sejam armazenados no banco de dados, é necessário ativar um api, então rode o seguinte comando no diretório `/api`:
```
go run .
```

### Publicando e Subscrevendo no HiveMQ
Para o envio e recebimento de dados ao cluster do HiveMQ, rode o seguinte comando nos diretórios `/publisher` e `/subscriber`:
```
go run .
```

### Visualizando a Dashboard
Para a visualização dos dados, rode o seguinte comando:
```
sudo docker run -d -p 3000:3000 \
-v ~<caminho-absoluto>/Modulo9/src/pond6/metabase.db:/metabase.db \
-v ~<caminho-absoluto>/Modulo9/src/pond6/database:/database \
--name metabase metabase/metabase
```

No metabase, insira suas informção e conecte-se ao banco de dados com o seguinte comando:
```
database/data.db
```

## IotsimMeta em Ação!
Você pode conferir o funcionamento de IotsimMeta no vídeo a seguir:

https://drive.google.com/file/d/1j1XDhm1z2UYCpQuphqfuQU5veI6NnfUh/view?usp=sharing