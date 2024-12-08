# load-tester

Ferramenta CLI para realizar testes de carga em serviços web.

## Como usar

### Build da Imagem Docker

#### docker build -t load-tester .

#### docker run load-tester --url=http://google.com --requests=100 --concurrency=10
