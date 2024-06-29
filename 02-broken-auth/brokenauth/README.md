# Resumo do Cenário: Broken Authentication

## Objetivo
Criar uma API em Go utilizando Gin Gonic e GORM com PostgreSQL que simule problemas na implementação de autenticação, permitindo que um atacante comprometa endpoints da API.

## Estrutura do Projeto

### Configuração do Projeto

1. Crie uma pasta para o projeto.
2. Inicialize um novo módulo Go.
3. Instale as dependências necessárias (`gin`, `gorm`, `postgres`).

### Configuração do Banco de Dados

1. Conecte-se ao PostgreSQL usando GORM.
2. Configure o banco de dados no arquivo `main.go`.

### Modelos

1. Defina os modelos `User` e `Token` em `models/models.go`.

### Rotas e Controladores

1. Implemente as rotas para registro, login e acesso protegido em `controllers/controllers.go`.
2. Adicione a lógica para autenticação básica e autorização.

## Teste do Cenário

1. Criar Usuários.
2. Fazer Login.
3. Tentar Acessar o Endpoint Protegido.

## Correção da Vulnerabilidade

### Armazenamento Seguro de Senhas

1. Use bcrypt para hash das senhas.

### Geração Segura de Tokens

1. Use uma biblioteca confiável para geração e gerenciamento de tokens, como JWT.
