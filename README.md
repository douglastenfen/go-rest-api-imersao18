# Imersão Fullcycle 18 - Desafio API REST GO

![Imersão Full Stack && Full Cycle](https://events-fullcycle.s3.amazonaws.com/events-fullcycle/static/site/img/grupo_4417.png)

Participe gratuitamente: <https://imersao.fullcycle.com.br/>

## Sobre o desafio

- Converter o arquivo `data.js` para JSON, e salvar na raiz do projeto GO.
  - O arquivo está no [repositório oficial](https://github.com/devfullcycle/imersao18/blob/main/nextjs-frontend/node-api/data.js).
- A aplicação irá expor algumas rotas e deverá ler o arquivo JSON ao iniciar o programa com o go run main.go, e manipular os dados sempre na memória (não salvando nada no arquivo JSON).

### Criação dos endpoints

- GET /events - Listar todos os eventos
- GET /events/:eventId - Listar os dados de um evento
- GET /events/:eventId/spots - Listar os lugares de um evento
- POST /event/:eventId/reserve - Reservar um lugar
