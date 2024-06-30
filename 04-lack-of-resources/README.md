## Cenário

Vamos criar um sistema simples onde os usuários podem fazer requisições para uma rota específica. Vamos demonstrar uma implementação sem limitação de taxa, permitindo que um atacante faça muitas requisições em um curto período de tempo.

## Objetivo
Criar uma API em Go utilizando Gin Gonic e GORM com PostgreSQL que simule a falta de controle de recursos e limitação de taxa, permitindo que um atacante possa abusar da API com um número excessivo de requisições, potencialmente causando negação de serviço (DoS).

## Estrutura do Projeto

### Configuração do Projeto

1. Crie uma pasta para o projeto.
2. Inicialize um novo módulo Go.
3. Instale as dependências necessárias (`gin`, `gorm`, `postgres`, `tollbooth`).

### Configuração do Banco de Dados

1. Conecte-se ao PostgreSQL usando GORM.
2. Configure o banco de dados no arquivo `main.go`.

### Modelos

1. Defina o modelo `RequestLog` em `models/models.go`.

### Rotas e Controladores

1. Implemente a rota para fazer requisições em `controllers/controllers.go`.
2. Adicione a lógica sem limitação de taxa.

## Teste do Cenário

1. Fazer várias requisições para a rota `/ping`.

## Correção da Vulnerabilidade

1. Use a biblioteca `tollbooth` para implementar a limitação de taxa.

Com isso, mitigamos a vulnerabilidade de falta de controle de recursos e limitação de taxa, garantindo que a API não seja abusada com um número excessivo de requisições.