# Guia de Conceitos de Software

> Conceitos essenciais que **todo desenvolvedor de software deve conhecer** para criar sistemas mais robustos,
> manuteníveis e escaláveis.

---

## **Fundamentos de Programação**

- **[POO (Programação Orientada a Objetos)](scg/fundamentals/poo/README.md)**: Conceitos básicos como classes, objetos,
  herança,
  encapsulamento, abstração e polimorfismo.
- **[Abstração](scg/fundamentals/abstraction/README.md)**: Como simplificar sistemas complexos focando nos aspectos
  essenciais.
- **[Polimorfismo](scg/fundamentals/polymorphism/README.md)**: Capacidade de objetos de diferentes classes responderem ao
  mesmo método
  de formas diferentes.
- **[Injeção de Dependência (DI)](scg/fundamentals/dependency-injection/README.md)**: Padrão que promove o desacoplamento
  entre
  classes, facilitando testes e manutenção.
- **[Beans e Estereótipos](scg/fundamentals/beans-stereotypes/README.md)**: Conceitos de componentes gerenciados em
  frameworks como
  Spring.
- **[ORM (Object-Relational Mapping)](scg/fundamentals/orm/README.md)**: O que é, para que serve e como mapeia objetos
  para bancos de
  dados relacionais.
- **[JPA vs Hibernate](scg/fundamentals/jpa-vs-hibernate/README.md)**: Diferenças entre a especificação JPA e sua
  implementação mais
  popular, o Hibernate.

---

## **Padrões de Projeto (Design Patterns)**

- **[Strategy](scg/design-patterns/strategy/README.md)**: Permite que algoritmos sejam selecionados em tempo de execução.
- **[Chain of Responsibility](scg/design-patterns/chain-of-responsibility/README.md)**: Encadeia objetos para processar
  uma requisição.
- **[Fluent API](scg/design-patterns/fluent-interface/README.md)**: Técnica para criar APIs mais legíveis e expressivas.
- **[Builder](scg/design-patterns/builder/README.md)**: Separa a construção de objetos complexos de sua representação.
- **[Factory](scg/design-patterns/factory/README.md)**: Centraliza a criação de objetos, promovendo flexibilidade.
- **[Singleton](scg/design-patterns/singleton/README.md)**: Garante que uma classe tenha apenas uma instância.
- **[Observer](scg/design-patterns/observer/README.md)**: Notifica objetos sobre mudanças de estado.
- **[Decorator](scg/design-patterns/decorator/README.md)**: Adiciona comportamentos a objetos dinamicamente.
- **[Adapter](scg/design-patterns/adapter/README.md)**: Permite que classes com interfaces incompatíveis trabalhem juntas.
- **[Facade](scg/design-patterns/facade/README.md)**: Simplifica a interação com subsistemas complexos.

---

## **Arquitetura de Software**

- **[Arquitetura de Camadas](scg/software-architecture/layered-architecture/README.md)**: Separação de responsabilidades
  em camadas (
  apresentação, lógica de negócios e dados).
- **[Microserviços](scg/software-architecture/microservices/README.md)**: Divisão da aplicação em serviços independentes.
- **[API Gateway](scg/software-architecture/api-gateway/README.md)**: Ponto único de entrada para microserviços.
- **[Domain-Driven Design (DDD)](scg/software-architecture/ddd/README.md)**: Foco no domínio do negócio para modelagem de
  software.
- **[Event Sourcing](scg/software-architecture/event-sourcing/README.md)**: Armazenamento de mudanças de estado como
  eventos imutáveis.
- **[CQRS (Command Query Responsibility Segregation)](scg/software-architecture/cqrs/README.md)**: Separação de comandos (
  escrita) e consultas (
  leitura).
- **[Service Mesh](scg/software-architecture/service-mesh/README.md)**: Gerenciamento de comunicação entre microserviços.
- **[Orquestração vs Coreografia](scg/software-architecture/orchestration-vs-choreography/README.md)**: Centralização vs
  descentralização de
  controle em sistemas distribuídos.
- **[Stateless vs Stateful](scg/software-architecture/stateless-vs-stateful/README.md)**: Diferenças entre sistemas sem
  estado e com estado.

---

## **Banco de Dados**

- **[RDBMS vs NoSQL](scg/databases/rdbms-vs-nosql/README.md)**: Comparação entre bancos relacionais e não relacionais.
- **[Transações: Atomicidade e Isolamento](scg/databases/transactions/README.md)**: Garantia de consistência em operações.
- **[Optimistic vs Pessimistic Locking](scg/databases/locking/README.md)**: Técnicas de controle de concorrência.
- **[Idempotência](scg/databases/idempotence/README.md)**: Garantia de que operações repetidas tenham o mesmo efeito.

---

## **Desenvolvimento de Software**

- **[TDD (Test-Driven Development)](scg/software-development/tdd/README.md)**: Desenvolvimento guiado por testes.
- **[SOLID](scg/software-development/solid/README.md)**: Princípios de design orientado a objetos.
- **[Fluent Interface](scg/software-development/fluent-interface/README.md)**: Técnica para criar APIs mais legíveis.
- **[Cache](scg/software-development/cache/README.md)**: Armazenamento temporário para otimização de desempenho.
- **[DevOps](scg/software-development/devops/README.md)**: Integração entre desenvolvimento e operações.
- **[Containers e Orquestração (Docker, Kubernetes)](scg/software-development/containers-orchestration/README.md)**:
  Gerenciamento de sistemas
  distribuídos.

---

## **Resiliência e Escalabilidade**

- **[Tolerância a Falhas](scg/resilience-scalability/fault-tolerance/README.md)**: Capacidade de operar mesmo com falhas.
- **[Circuit Breaker](scg/resilience-scalability/circuit-breaker/README.md)**: Prevenção de falhas catastróficas em
  sistemas distribuídos.
- **[Pub/Sub (Publish/Subscribe)](scg/resilience-scalability/pub-sub/README.md)**: Comunicação assíncrona entre produtores
  e consumidores.

---

## **Observabilidade e Monitoramento**

- **[Observabilidade (Logs, Distributed Tracing, NewRelic)](scg/observability/README.md)**: Monitoramento e rastreamento
  de sistemas em produção.

---

## **Outros Conceitos Importantes**

- **[REST vs RESTful](scg/other-concepts/rest-vs-restful/README.md)**: Diferenças entre o estilo arquitetural REST e sua
  implementação
  prática.
- **[Clean Code](scg/other-concepts/clean-code/README.md)**: Boas práticas para escrever código legível e manutenível.
- **[YAGNI (You Aren't Gonna Need It)](scg/other-concepts/yagni/README.md)**: Princípio de evitar funcionalidades
  desnecessárias.
- **[GRASP (General Responsibility Assignment Software Patterns)](scg/other-concepts/grasp/README.md)**: Padrões para
  atribuição de responsabilidades em software orientado a objetos.
- **[KISS (Keep It Simple, Stupid)](scg/other-concepts/kiss/README.md)**: Simplicidade como princípio de design.
- **[DRY (Don't Repeat Yourself)](scg/other-concepts/dry/README.md)**: Evitar duplicação de código.

---

## **Padrões de Projeto Adicionais**

- **[Proxy](scg/additional-design-patterns/proxy/README.md)**: Controla o acesso a um objeto.
- **[Composite](scg/additional-design-patterns/composite/README.md)**: Trata objetos individuais e compostos de forma
  uniforme.
- **[State](scg/additional-design-patterns/state/README.md)**: Permite que um objeto altere seu comportamento quando seu
  estado muda.
- **[Template Method](scg/additional-design-patterns/template-method/README.md)**: Define o esqueleto de um algoritmo,
  permitindo que subclasses redefinam etapas específicas.

- **[Clean Architecture](scg/software-architecture/clean-architecture/README.md)**: Organização de código que promove
  separação de responsabilidades e independência de frameworks, interfaces e detalhes de implementação.
---
