# GoBus

Esta aplicação Go utiliza o Azure Service Bus para enviar e receber mensagens de uma fila ou tópico. A aplicação permite configurar a conexão com o Service Bus e controlar a quantidade de mensagens processadas em paralelo.

## Pré-requisitos
- Go 1.16 ou superior

## Instalação

- Clone o repositório:
```
git clone https://github.com/seu-usuario/seu-repositorio.git
cd seu-repositorio
```

## Instale as dependências:
```
go mod tidy
 ```

## Configuração
Certifique-se de definir a variável de ambiente necessária:
```
export SERVICE_BUS_CONNECTION_STRING="sua-connection-string"
```

## Execução
Para executar a aplicação, use o comando:
```
go run main.go
```

## Funcionalidades
### Configuração do Service Bus

A aplicação solicita ao usuário a configuração da conexão com o Service Bus:

- SERVICE_BUS_CONNECTION_STRING: String de conexão do Azure Service Bus como variavel de ambiente ou no input da execução.

## Parâmetros de Execução
A aplicação solicita ao usuário os seguintes parâmetros:

- MAX_MESSAGES_COUNT: Número máximo de mensagens a serem processadas por vez (padrão: 100).
- SENDER_PARALLEL_COUNT: Número de mensagens a serem enviadas em paralelo (padrão: 50).

## Envio e Recebimento de Mensagens
A aplicação realiza as seguintes operações:

Receber Mensagens: Recebe mensagens da fila ou tópico de origem.
Enviar Mensagens: Envia mensagens para a fila ou tópico de destino.
Processamento Paralelo: Processa mensagens em paralelo utilizando goroutines e um canal de semáforo para controlar a quantidade de mensagens processadas simultaneamente.

## Contribuição
Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou pull requests.

## Licença
Este projeto está licenciado sob a licença MIT. Consulte o arquivo LICENSE para obter mais informações.