# Api Gateway

Um **API Gateway** é um padrão arquitetural que atua como um ponto centralizado para o gerenciamento e roteamento de
chamadas de APIs em um sistema distribuído. Ele é responsável por redirecionar as solicitações dos clientes para os
serviços backend apropriados, realizando funções adicionais, como autenticação, balanceamento de carga, monitoramento,
caching, transformação de protocolos e controle de tráfego.

### **O que é um API Gateway?**

Em um sistema distribuído de microserviços, como aqueles baseados na arquitetura de microservices, o **API Gateway** age
como um ponto de entrada único para todas as requisições feitas aos serviços. Ele simplifica a interação do cliente com
múltiplos serviços internos, ao mesmo tempo em que centraliza preocupações transversais, como segurança e monitoramento.

O API Gateway oferece diversas vantagens, incluindo a redução do número de chamadas de rede necessárias, a consolidação
da lógica de roteamento, e o desacoplamento entre clientes e serviços internos.

---

## **Principais Funções do API Gateway**

1. **Roteamento de Requisições**:
    - O API Gateway direciona as requisições dos clientes para os serviços backend adequados. Ele pode mapear endpoints
      externos para APIs internas.

2. **Autenticação e Autorização**:
    - O gateway pode ser configurado para autenticar e autorizar solicitações antes de encaminhá-las para os serviços.
      Isso pode incluir a verificação de tokens de autenticação (como JWTs), API keys ou integração com serviços de
      identidade como OAuth.

3. **Balanceamento de Carga**:
    - O API Gateway pode distribuir as requisições entre múltiplas instâncias de um serviço backend, implementando o
      balanceamento de carga, garantindo alta disponibilidade e escalabilidade.

4. **Agregação de Resultados**:
    - O API Gateway pode combinar respostas de múltiplos serviços em uma única resposta, facilitando a integração entre
      sistemas e evitando que os clientes façam múltiplas chamadas.

5. **Monitoramento e Logging**:
    - O API Gateway pode coletar métricas de desempenho, como latência e número de requisições, além de gerar logs que
      são valiosos para debugging e análise de tráfego.

6. **Caching**:
    - O API Gateway pode armazenar em cache respostas de APIs frequentemente solicitadas, melhorando a performance e
      reduzindo a carga sobre os serviços backend.

7. **Transformação de Protocolos**:
    - Pode realizar transformações de protocolo, por exemplo, de XML para JSON, ou vice-versa, para garantir que os
      serviços internos se comuniquem corretamente com os clientes, sem a necessidade de modificar os serviços
      existentes.

8. **Rate Limiting e Controle de Tráfego**:
    - Para evitar sobrecarga de um serviço ou sistema, o API Gateway pode aplicar políticas de controle de tráfego, como
      limitar o número de requisições por usuário ou IP, ou até mesmo definir períodos de "cooldown" entre requisições.

---

## **Arquitetura do API Gateway**

A arquitetura do API Gateway geralmente envolve os seguintes componentes principais:

1. **Gateway de API**:
    - É o ponto de entrada principal para todas as requisições. Ele pode ser uma solução implementada internamente ou um
      serviço fornecido por provedores de nuvem (como **Amazon API Gateway**, **Azure API Management**, etc.).

2. **Microserviços ou Backends**:
    - Os microserviços são as entidades internas que realizam o processamento da lógica de negócios. O API Gateway se
      comunica com eles para rotear as requisições e retornar as respostas.

3. **Sistema de Autenticação e Autorização**:
    - O API Gateway pode delegar autenticação a um sistema externo (como OAuth, JWT ou integração com LDAP).

4. **Protocolo de Comunicação**:
    - O gateway pode ser configurado para suportar diversos tipos de protocolos (HTTP, HTTPS, WebSocket, etc.),
      garantindo que ele possa interagir com diferentes sistemas e arquiteturas.

---

## **Benefícios do API Gateway**

1. **Centralização da Lógica de Tráfego**:
    - Ao centralizar o controle do tráfego entre clientes e serviços, o API Gateway facilita a manutenção e a gestão do
      sistema.

2. **Desacoplamento entre Cliente e Backend**:
    - O cliente se comunica diretamente com o API Gateway, não precisando conhecer a estrutura dos serviços internos ou
      como eles se comunicam.

3. **Segurança**:
    - Com o controle de autenticação e autorização centralizado, o API Gateway pode proteger o acesso a diferentes
      serviços de maneira mais eficiente, implementando políticas de segurança de forma unificada.

4. **Facilidade de Escalabilidade**:
    - O API Gateway pode ser escalado de forma independente dos serviços backend, permitindo uma maior flexibilidade na
      arquitetura e atendendo melhor ao aumento de demanda.

5. **Simplificação da Comunicação**:
    - O cliente não precisa realizar múltiplas chamadas para diferentes serviços; o API Gateway pode agregar dados de
      vários serviços e retornar uma resposta consolidada.

6. **Redução de Sobrecarga nos Serviços Backend**:
    - O uso de caching, balanceamento de carga e rate limiting ajuda a reduzir a carga sobre os serviços internos,
      garantindo uma maior disponibilidade e performance.

---

## **Desvantagens do API Gateway**

1. **Ponto Único de Falha**:
    - Se o API Gateway falhar, toda a comunicação do sistema pode ser afetada. Para mitigar isso, pode-se implementar
      estratégias como **high availability** (alta disponibilidade) e **failover**.

2. **Sobrecarga de Processamento**:
    - Em algumas arquiteturas, o API Gateway pode se tornar um gargalo de desempenho se não for configurado corretamente
      ou se ele tiver que processar muitas requisições complexas.

3. **Complexidade**:
    - A introdução de um API Gateway pode aumentar a complexidade da arquitetura do sistema, exigindo mais manutenção e
      monitoramento.

---

## **Exemplos Populares de API Gateway**

1. **Kong**:
    - Um dos mais populares gateways de API open source. O Kong pode ser usado para roteamento de tráfego, autenticação,
      segurança e muito mais.

2. **Nginx**:
    - Embora seja conhecido como um servidor web, o **Nginx** também pode ser configurado como um API Gateway,
      fornecendo roteamento, balanceamento de carga, caching e monitoramento.

3. **Amazon API Gateway**:
    - Serviço da Amazon Web Services (AWS) que permite criar e gerenciar APIs, com integração a outros serviços da AWS
      como Lambda, DynamoDB e S3.

4. **Zuul (Netflix)**:
    - O Zuul é uma solução popular de API Gateway, desenvolvida pela Netflix, para roteamento de tráfego e integração
      com microserviços. Ele pode ser usado em conjunto com o **Eureka**, para descoberta de serviços.

5. **Apigee**:
    - Plataforma de gerenciamento de API do Google, que oferece funcionalidades para criar, gerenciar, monitorar e
      proteger APIs em grande escala.

6. **Traefik**:
    - Traefik é um proxy reverso que também funciona como um API Gateway, com foco em ambientes de containers e
      orquestração de microserviços, como o Docker e o Kubernetes.

---

## **Exemplo Simples de Configuração de API Gateway**

Vamos usar o **Kong** como exemplo para configurar um simples API Gateway:

1. **Instalação do Kong com Docker**:

```bash
# Baixar e rodar o Kong no Docker
docker run -d --name kong-database \
  -e "KONG_DATABASE=postgres" \
  -e "KONG_PG_HOST=kong-database" \
  -p 5432:5432 \
  postgres:9.6

docker run -d --name kong \
  -e "KONG_DATABASE=postgres" \
  -e "KONG_PG_HOST=kong-database" \
  -p 8000:8000 \
  -p 8443:8443 \
  kong:latest
```

2. **Adicionar uma API no Kong**:

```bash
# Adicionando uma API para roteamento
curl -i -X POST http://localhost:8001/services/ \
  --data "name=example-service" \
  --data "url=http://example.com"
  
# Criando um route para a API
curl -i -X POST http://localhost:8001/services/example-service/routes \
  --data "paths[]=/example"
```

Agora, qualquer requisição para `http://localhost:8000/example` será roteada para o serviço `http://example.com`.

---

## **Conclusão**

O **API Gateway** é uma peça crucial em sistemas modernos de microserviços, proporcionando uma maneira centralizada de
gerenciar tráfego, segurança, escalabilidade e integridade de dados entre os clientes e os serviços internos. Sua
implementação oferece benefícios como redução da complexidade do cliente, maior segurança e controle, além de otimizar a
comunicação entre os serviços.

Entretanto, como qualquer padrão arquitetural, o API Gateway vem com seus desafios, como o risco de ser um ponto único
de falha e a sobrecarga de processamento. Por isso, é importante projetá-lo e configurá-lo adequadamente para atender
aos requisitos de desempenho e resiliência do sistema.

