# Envio de menssagens no canal do Slack
## Utilizando:
- Golang (Como servidor para receber json via POST)
- LocalStak (Simulador do serviço SQS da AWS)
- Python (Como worker para ler as menssagens da fila SQS e enviar ao canal do Slack)
- Docker e Docker Compose para subir o ambiente em containers

### Observações:
- Credencias de AWS(LocalStak) são mockadas, caso seja configurado em ambiente AWS Cloud a forma de utilização para comunicação é feita via IAM role

### Servidor Golang:
#### Necessário colocar o nome da fila criada no arquivo main.go [nome pardrão - alert]
```sh
sendMessage(sqsClient, string(newData), "http://localhost.localstack.cloud:4566/000000000000/<SUA_FILA>")
```
### Localstack:
#### Necessário colocar o nome da fila criada no arquivo sqs.sh [nome pardrão - alert]
```sh
awslocal --endpoint-url=http://0.0.0.0:4566 sqs create-queue --region us-east-1 --queue-name SUA_FILA
```

### Python Worker:
#### Necessário adicionar as informações para comunicação com a API do Slack
```sh
slack_token = 'YourSlackToken'
slack_icon_emoji = ':see_no_evil:'
slack_user_name = 'YourUserName'
queue_name = 'alert'
```

## Iniciando ambiente com docker compose
```sh
$ git clone https://github.com/crypto-br/ambienteSlackSQS.git
$ cd ambienteSlackSQS
$ docker-compose build .
$ docker-compose up -d
```

### Testes
#### Realizando o start do ambiente com docker-compose
- Inicializando ambiente

![start](/img/start.jpg)

- Containers rodando

![containers](/img/containers.jpg)

- Enviando post para servidor GO

![server](/img/post.jpg)

- Menssagem enviada para o Slack

![slack](/img/slack.jpg)