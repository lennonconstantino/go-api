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
7. Criamos o script sql
8. fizemos os pacotes model e controller para fazer os primeiros testes mocados
9. Criamos a rota /products no main.go
10. Criamos a camada de use case
11. ...
12. go get package github/lib/pq
13. Rodar para concertar dependencias
    $ go mod tidy
14. Executar no main.go para verificar se esta tudo ok
15. Vamos criar a rota para inserir produtos

