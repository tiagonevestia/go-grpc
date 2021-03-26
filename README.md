# GO - gRPC

Projeto de estudo do módulo de introdução à comunicação, no qual colocamos em prática usando o protocolo gRPC com GO.

### O que é o gRPC?

É um framework projetado para facilitar o processo de comunicação entre sistemas de forma rápida, leve e independente de linguagem.

> **RPC** é a definição de um protocolo para executar procedimentos em outros computadores em rede.

### Principais vantagens

- Ideal para microsserviços.
- Reduza a transmissão binária de dados, reduza o uso da rede e a latência.
- Streaming de Bidirecional e HTTP/2
- Geração de lib de forma automática

### Como usar

Para clonar e executar este projeto, você precisará do Git e Docker instalado no seu computador.
Na sua linha de comando:

```bash
# Clonando o projeto
$ git clone https://github.com/tiagonevestia/go-grpc go-grpc

# Vamos entrar dentro do repositório
$ cd go-grpc

# Executando o docker-compose
$ docker-compose up -d --build

# Entrando no container e executando o servidor
$ docker-compose exec app bash
$ go run cmd/server/server.go

# Entrando no container e executando o cliente
$ docker-compose exec app bash
$ go run cmd/client/client.go

```

Feito com ♥ por Tiago Neves 👋
