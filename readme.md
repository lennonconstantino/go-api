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
16. ... apos muitas etapas de implemetacao e testes
17. Vamos criar uma imagem docker da nossa aplicacao, vamos para a raiz do projeto
    $ touch Dockerfile
18. Editamos o arquivo Dockerfile e conn.go
19. Criamos a imagem da nossa api
    $ docker build -t go-api .
    $ docker image ls
20. vamos alterar o nosso Docker Compose
21. Derrubar e pausar imagens anteriores
22. Executar o Docker compose novamente
    $ docker compose up -d

------

## proximos passos
Criar uma rota de PUT para atualizar produtos
Criar uma rota de DELETE para deletar algum produto
Criar uma autenticacao jwt
