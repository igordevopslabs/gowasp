
## Cenário
Vamos criar um sistema simples onde os usuários podem buscar informações de um banco de dados de produtos. Vamos demonstrar uma implementação vulnerável a injeções SQL, permitindo que um atacante execute comandos SQL maliciosos.

Na rota searchProduct, estamos construindo uma consulta SQL dinâmica concatenando o valor da entrada diretamente na consulta, o que torna a aplicação vulnerável a injeções SQL.

```bash
curl -X GET "http://localhost:3000/products/search?name='1'='1"
```

## Objetivo
Criar uma API em Go utilizando Gin Gonic e GORM com PostgreSQL que simule vulnerabilidades de injeção, como SQL, NoSQL, comando OS, injeções LDAP, etc.

