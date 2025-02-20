# **Observability (Logs, Distributed Tracing, New Relic, Datadog, Grafana, Prometheus, etc.)**

**Observabilidade** é a capacidade de medir o estado interno de um sistema com base na sua saída externa. Em vez de
apenas monitorar eventos conhecidos (como erros ou falhas), a observabilidade fornece a capacidade de entender o
comportamento de um sistema distribuído e identificar a causa raiz de problemas, mesmo que esses problemas não sejam
previsíveis ou sejam complexos. As principais técnicas de observabilidade incluem **logs**, **tracing distribuído** e *
*métricas**.

---

## **Componentes Principais de Observabilidade**

1. [**Logs**](#1-logs)
2. [**Distributed Tracing**](#2-distributed-tracing)
3. [**Métricas**](#3-métricas)
4. [**Ferramentas de Observabilidade (Datadog, New Relic, Grafana, Prometheus, etc.)
   **](#4-ferramentas-de-observabilidade)

---

## **1. Logs**

**Logs** são registros de eventos que acontecem durante a execução de um sistema. Eles são uma das formas mais antigas e
utilizadas para observabilidade. Logs podem fornecer uma visão detalhada sobre o que está acontecendo em uma aplicação,
como mensagens de erro, status de tarefas, interações de usuários, e muito mais.

### **Características dos Logs**:

- **Simples e Direto**: Permite registrar eventos em tempo real de forma simples.
- **Diversidade de Formatos**: Logs podem ser formatados de diversas maneiras (texto simples, JSON, XML).
- **Fácil de Implementar**: Pode ser integrado rapidamente a qualquer aplicação.
- **Exemplo**: Mensagens de erro, eventos de processamento, etc.

### **Tipos de Logs**:

1. **Logs de Aplicação**:

- Registram a atividade da própria aplicação, como requisições HTTP, falhas de autenticação, mensagens de erro, etc.

2. **Logs de Servidor**:

- Registram eventos relacionados ao servidor, como falhas de hardware, uso de CPU, memória, etc.

3. **Logs de Rede**:

- Registram o tráfego de rede, como pacotes de dados, status de conexão, latência, etc.

### **Ferramentas para Logs**:

- [**Elasticsearch, Logstash e Kibana (ELK Stack)**](https://www.elastic.co/what-is/elk-stack): Popular stack para
  agregação, busca e visualização de logs.
- [**Fluentd**](https://www.fluentd.org/): Ferramenta de coleta e envio de logs para diversos destinos.
- [**Grafana Loki**](https://grafana.com/oss/loki/): Ferramenta de agregação de logs focada na integração com o Grafana.
- [**Splunk**](https://www.splunk.com/): Plataforma para coletar, indexar e visualizar logs.

---

## **2. Distributed Tracing**

**Tracing Distribuído** (ou rastreamento distribuído) é uma técnica que permite rastrear a jornada de uma requisição ou
evento por um sistema distribuído. Em sistemas modernos, a aplicação pode ser dividida em vários serviços, e cada
serviço pode processar diferentes partes de uma requisição. O tracing distribuído permite que você "siga" a requisição
enquanto ela atravessa múltiplos serviços, ajudando a identificar gargalos e pontos de falha.

### **Características do Distributed Tracing**:

- **Rastreamento de Requisições**: Ajuda a visualizar como uma requisição se move através de microserviços ou
  componentes.
- **Contexto de Performance**: Permite identificar onde o tempo está sendo gasto em uma requisição e onde ocorrem os
  gargalos.
- **Rastreio de Latência**: Ajuda a identificar latências entre serviços e áreas onde o desempenho pode ser melhorado.
- **Visibilidade no Fluxo de Dados**: Facilita a análise de transações complexas que atravessam múltiplos sistemas.

### **Exemplo**:

Em uma aplicação de e-commerce, uma requisição de compra pode passar por serviços de autenticação, cálculo de
preços, estoque, pagamento e confirmação de pedido. O tracing ajuda a identificar onde ocorrem atrasos ou falhas.

### **Ferramentas para Distributed Tracing**:

- [**Jaeger**](https://www.jaegertracing.io/): Sistema de rastreamento distribuído de código aberto.
- [**Zipkin**](https://zipkin.io/): Uma plataforma de rastreamento distribuído que coleta e exibe dados de performance e
  latência.
- [**OpenTelemetry**](https://opentelemetry.io/): Um conjunto de ferramentas, APIs e SDKs que ajudam a coletar dados de
  observabilidade, incluindo tracing e métricas.
- [**Datadog**](https://www.datadoghq.com/): Oferece tracing distribuído como parte de sua plataforma de
  observabilidade.

---

## **3. Métricas**

**Métricas** fornecem dados quantitativos sobre o comportamento e o desempenho de um sistema. Elas medem aspectos como
latência, taxa de erro, uso de recursos (CPU, memória), contagem de requisições, e outros indicadores chave de
desempenho (KPIs). As métricas são essenciais para ter uma visão geral da saúde do sistema.

### **Características das Métricas**:

- **Quantitativo**: Fornece valores numéricos mensuráveis sobre o sistema.
- **Ideal para Monitoramento Contínuo**: Métricas podem ser coletadas em intervalos regulares (como a cada segundo ou
  minuto).
- **Análise de Tendências**: Permite detectar padrões e anomalias em períodos de tempo, como picos de latência ou
  aumento da taxa de erros.
- **Exemplo**: Tempo médio de resposta de uma API, quantidade de requisições por segundo, uso de CPU.

### **Métricas Comuns**:

- **Latência**: O tempo que leva para uma requisição ser processada.
- **Taxa de Erros**: A porcentagem de requisições que resultam em erro.
- **Taxa de Requisições**: A quantidade de requisições processadas por um sistema em um período de tempo.
- **Uso de Recursos**: CPU, memória, espaço de disco.

### **Ferramentas para Métricas**:

- [**Prometheus**](https://prometheus.io/): Sistema de monitoramento de código aberto que coleta e armazena métricas em
  tempo real.
- [**Grafana**](https://grafana.com/): Mais conhecido como uma plataforma de visualização de dados, o Grafana também
  pode ser usado para exibir métricas coletadas pelo Prometheus.
- [**Datadog**](https://www.datadoghq.com/): Fornece métricas em tempo real e análises de dados.
- [**StatsD**](https://github.com/statsd/statsd): Ferramenta para enviar métricas para backend, como Graphite,
  Prometheus ou Datadog.

---

## **4. Ferramentas de Observabilidade**

Ferramentas de observabilidade integradas, como **New Relic**, **Datadog**, **Prometheus**, entre outras, permitem
coletar, analisar e visualizar logs, métricas e traces. Elas ajudam as equipes a ter uma visão completa do sistema,
identificando pontos de falha, monitorando a saúde e analisando o desempenho em tempo real.

### **Datadog**

- [**Site oficial**](https://www.datadoghq.com/)
- **Funcionalidades**:
    - **APM (Application Performance Monitoring)**: Monitora o desempenho de aplicativos em tempo real.
    - **Distributed Tracing**: Rastreia transações e requisições em um sistema distribuído.
    - **Métricas e Dashboards**: Oferece visualizações de métricas e análise de dados em tempo real.
    - **Alertas e Notificações**: Configura alertas para quando o sistema se desvia dos padrões estabelecidos.

### **New Relic**

- [**Site oficial**](https://newrelic.com/)
- **Funcionalidades**:
    - **APM (Application Performance Monitoring)**: Monitora o desempenho de aplicativos em tempo real.
    - **Distributed Tracing**: Rastreia transações e requisições em um sistema distribuído.
    - **Métricas e Dashboards**: Oferece visualizações de métricas e análise de dados em tempo real.
    - **Alertas e Notificações**: Configura alertas para quando o sistema se desvia dos padrões estabelecidos.

### **Grafana**

- [**Site oficial**](https://grafana.com/)
- **Funcionalidades**:
    - **Integração com Prometheus**: Exibe métricas coletadas pelo Prometheus.
    - **Plugins**: Suporta uma variedade de plugins para integração com outras ferramentas.
    - **Alertas**: Configura alertas baseados em métricas.

### **Prometheus**

- [**Site oficial**](https://prometheus.io/)
- **Funcionalidades**:
    - **Coleta de Métricas**: Coleta métricas de sistemas e aplicações.
    - **Alertas**: Configura alertas baseados em métricas.
    - **Integração com Grafana**: Exibe métricas coletadas pelo Prometheus.

---

## **Conclusão**

A **observabilidade** é crucial para entender o comportamento de sistemas distribuídos, identificar problemas e melhorar
o desempenho. Ferramentas como [**Datadog**](https://www.datadoghq.com/), [**New Relic**](https://newrelic.com/), [*
*Grafana**](https://grafana.com/) e [**Prometheus**](https://prometheus.io/) facilitam a análise e visualização de logs,
métricas e tracing, promovendo uma manutenção mais eficiente em sistemas modernos.
