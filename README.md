# Guia de Conceitos de Software

> Conceitos essenciais que **todo desenvolvedor de software deve conhecer** para criar sistemas mais robustos,
> manuteníveis e escaláveis.

---

## **Fundamentos de Programação**

- **[POO (Programação Orientada a Objetos)](./fundamentals/poo/README.md)**: Conceitos básicos como classes, objetos,
  herança,
  encapsulamento, abstração e polimorfismo.
- **[Abstração](./fundamentals/abstraction/README.md)**: Como simplificar sistemas complexos focando nos aspectos
  essenciais.
- **[Polimorfismo](./fundamentals/polymorphism/README.md)**: Capacidade de objetos de diferentes classes responderem ao
  mesmo método
  de formas diferentes.
- **[Injeção de Dependência (DI)](./fundamentals/dependency-injection/README.md)**: Padrão que promove o desacoplamento
  entre
  classes, facilitando testes e manutenção.
- **[Beans e Estereótipos](./fundamentals/beans-stereotypes/README.md)**: Conceitos de componentes gerenciados em
  frameworks como
  Spring.
- **[ORM (Object-Relational Mapping)](./fundamentals/orm/README.md)**: O que é, para que serve e como mapeia objetos
  para bancos de
  dados relacionais.
- **[JPA vs Hibernate](./fundamentals/jpa-vs-hibernate/README.md)**: Diferenças entre a especificação JPA e sua
  implementação mais
  popular, o Hibernate.

---

## **Padrões de Projeto (Design Patterns)**

- **[Strategy](./design-patterns/strategy/README.md)**: Permite que algoritmos sejam selecionados em tempo de execução.
- **[Chain of Responsibility](./design-patterns/chain-of-responsibility/README.md)**: Encadeia objetos para processar
  uma requisição.
- **[Fluent API](design-patterns/fluent-interface/README.md)**: Técnica para criar APIs mais legíveis e expressivas.
- **[Builder](./design-patterns/builder/README.md)**: Separa a construção de objetos complexos de sua representação.
- **[Factory](./design-patterns/factory/README.md)**: Centraliza a criação de objetos, promovendo flexibilidade.
- **[Singleton](./design-patterns/singleton/README.md)**: Garante que uma classe tenha apenas uma instância.
- **[Observer](./design-patterns/observer/README.md)**: Notifica objetos sobre mudanças de estado.
- **[Decorator](./design-patterns/decorator/README.md)**: Adiciona comportamentos a objetos dinamicamente.
- **[Adapter](./design-patterns/adapter/README.md)**: Permite que classes com interfaces incompatíveis trabalhem juntas.
- **[Facade](./design-patterns/facade/README.md)**: Simplifica a interação com subsistemas complexos.

---

## **Arquitetura de Software**

- **[Arquitetura de Camadas](./software-architecture/layered-architecture/README.md)**: Separação de responsabilidades
  em camadas (
  apresentação, lógica de negócios e dados).
- **[Microserviços](./software-architecture/microservices/README.md)**: Divisão da aplicação em serviços independentes.
- **[API Gateway](./software-architecture/api-gateway/README.md)**: Ponto único de entrada para microserviços.
- **[Domain-Driven Design (DDD)](./software-architecture/ddd/README.md)**: Foco no domínio do negócio para modelagem de
  software.
- **[Event Sourcing](./software-architecture/event-sourcing/README.md)**: Armazenamento de mudanças de estado como
  eventos imutáveis.
- **[CQRS (Command Query Responsibility Segregation)](./software-architecture/cqrs/README.md)**: Separação de comandos (
  escrita) e consultas (
  leitura).
- **[Service Mesh](./software-architecture/service-mesh/README.md)**: Gerenciamento de comunicação entre microserviços.
- **[Orquestração vs Coreografia](./software-architecture/orchestration-vs-choreography/README.md)**: Centralização vs
  descentralização de
  controle em sistemas distribuídos.
- **[Stateless vs Stateful](./software-architecture/stateless-vs-stateful/README.md)**: Diferenças entre sistemas sem
  estado e com estado.

---

## **Banco de Dados**

- **[RDBMS vs NoSQL](./databases/rdbms-vs-nosql/README.md)**: Comparação entre bancos relacionais e não relacionais.
- **[Transações: Atomicidade e Isolamento](./databases/transactions/README.md)**: Garantia de consistência em operações.
- **[Optimistic vs Pessimistic Locking](./databases/locking/README.md)**: Técnicas de controle de concorrência.
- **[Idempotência](./databases/idempotence/README.md)**: Garantia de que operações repetidas tenham o mesmo efeito.

---

## **Desenvolvimento de Software**

- **[TDD (Test-Driven Development)](./software-development/tdd/README.md)**: Desenvolvimento guiado por testes.
- **[SOLID](./software-development/solid/README.md)**: Princípios de design orientado a objetos.
- **[Fluent Interface](./software-development/fluent-interface/README.md)**: Técnica para criar APIs mais legíveis.
- **[Cache](./software-development/cache/README.md)**: Armazenamento temporário para otimização de desempenho.
- **[DevOps](./software-development/devops/README.md)**: Integração entre desenvolvimento e operações.
- **[Containers e Orquestração (Docker, Kubernetes)](./software-development/containers-orchestration/README.md)**:
  Gerenciamento de sistemas
  distribuídos.

---

## **Resiliência e Escalabilidade**

- **[Tolerância a Falhas](./resilience-scalability/fault-tolerance/README.md)**: Capacidade de operar mesmo com falhas.
- **[Circuit Breaker](./resilience-scalability/circuit-breaker/README.md)**: Prevenção de falhas catastróficas em
  sistemas distribuídos.
- **[Pub/Sub (Publish/Subscribe)](./resilience-scalability/pub-sub/README.md)**: Comunicação assíncrona entre produtores
  e consumidores.

---

## **Observabilidade e Monitoramento**

- **[Observabilidade (Logs, Distributed Tracing, NewRelic)](./observability/README.md)**: Monitoramento e rastreamento
  de sistemas em produção.

---

## **Outros Conceitos Importantes**

- **[REST vs RESTful](./other-concepts/rest-vs-restful/README.md)**: Diferenças entre o estilo arquitetural REST e sua
  implementação
  prática.
- **[Clean Code](./other-concepts/clean-code/README.md)**: Boas práticas para escrever código legível e manutenível.
- **[YAGNI (You Aren't Gonna Need It)](./other-concepts/yagni/README.md)**: Princípio de evitar funcionalidades
  desnecessárias.
- **[GRASP (General Responsibility Assignment Software Patterns)](./other-concepts/grasp/README.md)**: Padrões para
  atribuição de responsabilidades em software orientado a objetos.
- **[KISS (Keep It Simple, Stupid)](./other-concepts/kiss/README.md)**: Simplicidade como princípio de design.
- **[DRY (Don't Repeat Yourself)](./other-concepts/dry/README.md)**: Evitar duplicação de código.

---

## **Padrões de Projeto Adicionais**

- **[Proxy](./additional-design-patterns/proxy/README.md)**: Controla o acesso a um objeto.
- **[Composite](./additional-design-patterns/composite/README.md)**: Trata objetos individuais e compostos de forma
  uniforme.
- **[State](./additional-design-patterns/state/README.md)**: Permite que um objeto altere seu comportamento quando seu
  estado muda.
- **[Template Method](./additional-design-patterns/template-method/README.md)**: Define o esqueleto de um algoritmo,
  permitindo que subclasses redefinam etapas específicas.

- **[Clean Architecture](./software-architecture/clean-architecture/README.md)**: Organização de código que promove
  separação de responsabilidades e independência de frameworks, interfaces e detalhes de implementação.
---
