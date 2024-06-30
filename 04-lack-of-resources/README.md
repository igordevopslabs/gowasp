## Cenário

Vamos criar um sistema simples onde os usuários podem fazer requisições para uma rota específica. Vamos demonstrar uma implementação sem limitação de taxa, permitindo que um atacante faça muitas requisições em um curto período de tempo.

## Objetivo
Criar uma API em Go utilizando Gin Gonic e GORM com PostgreSQL que simule a falta de controle de recursos e limitação de taxa, permitindo que um atacante possa abusar da API com um número excessivo de requisições, potencialmente causando negação de serviço (DoS).

