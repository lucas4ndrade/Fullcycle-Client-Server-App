## Client Server APP

Repositório para guardar o desafio 1 do curso GO Expert da Fullcycle

## Execução

Essa aplicação contém duas partes, o cliente e o servidor.

### Executando o servidor

Para executar o servidor, basta rodar:

```shell
go run server/main.go \
    --port= \
    --db-path=
```

### Parâmetros:

| parâmetro      | descrição                                                             | Valor default (se existir)|
|----------------|-----------------------------------------------------------------------|---------------------------|
| port           | Porta em que o servidor deverá subir                                  | 8080                      |
| db-path        | Caminho desejado do arquivo de banco de dados                         | ./fullcycle.db            |

O servidor irá subir uma API HTTP na porta informada, ou na porta 8080.

### Executando o cliente

Para executar o cliente, basta rodar:

```shell
go run client/main.go \
    --server-url= \
    --file-path=
```

### Parâmetros:

| parâmetro      | descrição                                                 | Valor default (se existir)|
|----------------|-----------------------------------------------------------|---------------------------|
| server-url     | URL aonde o servidor está rodando                         | http://localhost:8080     |
| file-path      | Caminho desejado do arquivo onde será salva a cotação     | cotacao.txt               |

O cliente irá fazer uma requisição HTTP para o servidor na URL desejada, e então irá gravar a cotação atual do Dólar no arquivo desejado
