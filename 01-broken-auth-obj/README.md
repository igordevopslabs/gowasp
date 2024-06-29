# Resumo do Cenário: Broken Object Level Authorization

## Objetivo
Criar uma API em Go utilizando Gin Gonic e GORM com PostgreSQL que simule uma falha de autorização a nível de objeto, permitindo que um usuário acesse dados de outro usuário.

## Estrutura do Projeto

### Configuração do Projeto

1. Crie uma pasta para o projeto.
2. Inicialize um novo módulo Go.
3. Instale as dependências necessárias (`gin`, `gorm`, `postgres`).

### Configuração do Banco de Dados

1. Conecte-se ao PostgreSQL usando GORM.
2. Configure o banco de dados no arquivo `main.go`.

### Modelos

1. Defina os modelos `User` e `Task` em `models/models.go`.

### Rotas e Controladores

1. Implemente as rotas para criar e visualizar tarefas em `controllers/controllers.go`.
2. Adicione a lógica para autenticação básica e autorização.

## Teste do Cenário

1. Crie usuários.
2. Crie tarefas para cada usuário.
3. Tente acessar as tarefas de outros usuários para verificar a vulnerabilidade.

## Correção da Vulnerabilidade

1. Adicione a verificação de propriedade da tarefa para corrigir a falha de autorização.
