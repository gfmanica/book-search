# Microsserviço de Enriquecimento de Dados de Livros

Este microsserviço recebe um ISBN via requisição HTTP POST, consulta a API do Google Books e retorna os dados estruturados do livro.

## Como rodar

```bash
go run main.go
```

O serviço ficará disponível em `http://localhost:8080/enrich`.

## Exemplo de requisição

POST `/enrich`

```json
{
  "isbn": "9788532530783"
}
```

## Exemplo de resposta

```json
{
  "title": "Harry Potter e a Pedra Filosofal",
  "authors": ["J.K. Rowling"],
  "publisher": "Rocco",
  "publishedDate": "2017-08-15",
  "description": "Harry Potter é um garoto cujos pais, feiticeiros, foram assassinados...",
  "pageCount": 264,
  "coverUrl": "http://books.google.com/books/content?id=...&printsec=frontcover"
}
```
