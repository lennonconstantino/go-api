# GO API

1. Iniciando o projeto
    $ go mod init go-api

2. Criar pasta cmd e main.go dentro dela
3. Fazer o download do pacote que vamos usar para criar a API
    $ go get github.com/gin-gonic/gin
4. Criamos o server com o endpoint para teste /ping no main.go
5. Criando o docker-compose na raiz do projeto
6. Para executar o docker
    $ docker compose up -d go_db
