## Cenário

Vamos criar um sistema simples de gerenciamento de produtos, onde os produtos têm informações sensíveis como custo de produção. Vamos expor desnecessariamente esses dados ao usuário final.

## Objetivo
Criar uma API em Go utilizando Gin Gonic e GORM com PostgreSQL que simule a exposição excessiva de dados, permitindo que um atacante obtenha mais informações do que o necessário.

## Estrutura do Projeto

### Configuração do Projeto

1. Crie uma pasta para o projeto.
2. Inicialize um novo módulo Go.
3. Instale as dependências necessárias (`gin`, `gorm`, `postgres`).

### Configuração do Banco de Dados

1. Conecte-se ao PostgreSQL usando GORM.
2. Configure o banco de dados no arquivo `main.go`.

### Modelos

1. Defina o modelo `Product` em `models/models.go`.

### Rotas e Controladores

1. Implemente as rotas para criar e listar produtos em `controllers/controllers.go`.
2. Adicione a lógica que expõe dados desnecessários.

## Teste do Cenário

1. Criar Produto.
2. Listar Produtos.