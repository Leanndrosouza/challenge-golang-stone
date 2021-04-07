# Aplicação para o Desafio Técnico - Go(lang)

Código feito para solucionar desafio proposto pela Stone que se trata de criar uma API de transferência entre contas Internas de um banco digital. 

### Aplicação

A aplicação trata-se de uma api que tem como intuito criação de contas, listagem de contas e saldos, autenticação e realização de transferências entre contas. Algumas das rotas são permitidas apenas para usuários autenticados e envolvem algumas regrinhas para manter as consistência dos dados. Essa aplicação foi feita na linguagem Go(lang) e é possível executá-la em ambiente Docker.

Para mais informações sobre o desafio, acesse o [Desafio](https://gist.github.com/guilhermebr/fb0d5896d76634703d385a4c68b730d8 "Desafio").

### Requerimento

Para executar essa aplicação é necessário que você tenha configurado o seu ambiente da linguagem Go e Banco de dados MySQL. Caso não queira se preocupar com essa configuração, o ideal é utilizar o Docker e assim não precisar se preocupar com a configuração da máquina local. Vou continuar esse tutorial imaginando que foi decidido que ela irá rodar no Docker, e nesse caso vai ser necessário instalar e configurar o **Docker** e **docker compose** em sua máquina.

### Como Começar

Vamos começar criando uma pasta com um nome qualquer, vou chamá-la de **api**. Dentro da pasta criada, clone este repositório e crie uma outra pasta com o nome **mysql_data**, esta vai servir para manter os dados do banco de dados guardados mesmo após derrubar o container da aplicação. Agora acesse a pasta do repositório e copie os arquivo **docker-compose.yml** e **docker-compose.test.yml** para a pasta que você criou inicialmente, no meu caso para a pasta **api**.

Dessa forma, teremos a seguinte estrutura:

	/api
		/challenge-golang-stone
		/mysql_data
		docker-compose.yml
		docker-compose.test.yml

Dentro da pasta **api** execute o comando `docker-compose build` para construir e baixar imagens do Docker. Após o comando finalizar com sucesso, execute `docker-compose up` levantar todos os serviços/containers, inclusive a aplicação em Go.

Feito isso, o banco de dados, o phpmyadmin e a aplicação estarão rodando em sua máquina. Agora vamos criar as tabelas do banco. Acesse o phpmyadmin **localhost:8888** com as seguintes credenciais:

 - **username**: root
 - **password**: password

Já logado no phpmyadmin, acesse a aba SQL para executar o conteúdo do arquivo sql.sql que se encontra em *challenge-golang-stone/sql/sql.sql*. Após a executação do códgo sql as tabelas do banco serão criadas.

### Como utilizar

Será necessário utilizar um programa que permita realização de requisições HTTP com métodos POST e GET, que seja possível passar informações no corpo da requisição por formato JSON. Recomendo utilizar o [Insomnia](https://insomnia.rest/ "Insomnia") ou [Postman](https://www.postman.com/ "Postman").

Para detalhes sobre as rotas disponíveis e como utilizá-las, recomendo acessar o [README do desafio](https://gist.github.com/guilhermebr/fb0d5896d76634703d385a4c68b730d8 "README do desafio") e verificar as regras e restrições de cada rota.

### Testes

Para executar os testes, vá para a pasta criada inicialmente, nesse caso a pasta **api** e execute o seguinte comando para rodar os testes:

`docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit`

**OBS**: Lembre-se de estar com o banco de dados configurado antes de executar os testes e que os testes utilizam a mesma base de dados da aplicação, dessa forma seus dados anteriormente cadastrados serão apagados ao finalizar os testes.

