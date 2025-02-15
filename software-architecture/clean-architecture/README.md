# Padrão de Projeto Adicional - Clean Architecture

![clean-architecture.png](clean-architecture.png)

### Introdução

O padrão **Clean Architecture**, proposto por Robert C. Martin (Uncle Bob), é uma abordagem arquitetural que visa criar
sistemas de software de alta coesão e baixo acoplamento. O principal objetivo desse padrão é tornar o código mais
simples de manter e escalar, separando responsabilidades em diferentes camadas e mantendo as regras de negócio (core)
isoladas de outras dependências, como frameworks, serviços externos ou bancos de dados.

### Arquiteturas Relacionadas

Nos últimos anos, diversas ideias sobre arquitetura de sistemas foram propostas, convergindo no objetivo comum de
separação de responsabilidades. Algumas delas incluem:

- **Arquitetura Hexagonal (Ports and Adapters)**, por Alistair Cockburn, amplamente adotada por Steve Freeman e Nat
  Pryce em seu livro *Growing Object Oriented Software*.
- **Arquitetura Onion**, por Jeffrey Palermo.
- **Screaming Architecture**, proposta em um dos blogs de Uncle Bob.
- **DCI (Data, Context, and Interaction)**, por James Coplien e Trygve Reenskaug.
- **BCE (Boundary-Control-Entity)**, por Ivar Jacobson, definido em seu livro *Object Oriented Software Engineering: A
  Use-Case Driven Approach*.

Apesar das diferenças nos detalhes, essas arquiteturas são bastante similares e compartilham os seguintes princípios:

- **Independência de Frameworks**: A arquitetura não depende de bibliotecas específicas, permitindo o uso de frameworks
  como ferramentas secundárias, em vez de integrar o sistema diretamente a elas.
- **Testabilidade**: As regras de negócio podem ser testadas isoladamente, sem a UI, banco de dados, servidor web ou
  qualquer outro elemento externo.
- **Independência da UI**: A interface do usuário pode ser alterada sem impacto no restante do sistema. Por exemplo, uma
  interface web pode ser substituída por uma interface console sem afetar as regras de negócio.
- **Independência do Banco de Dados**: O banco de dados é tratado como um detalhe que pode ser substituído (por exemplo,
  Oracle ou SQL Server pode ser trocado por MongoDB, CouchDB, etc.).
- **Independência de Agentes Externos**: As regras de negócio não têm dependência de elementos externos.

### Regras Fundamentais

No diagrama representativo da Clean Architecture, os círculos concêntricos indicam diferentes áreas do software. Em
geral:

- Quanto mais interno, maior o nível de abstração.
- Quanto mais externo, mais concreto e detalhado.

A regra primordial da Clean Architecture é a **Regra de Dependência**, que determina que as dependências de código-fonte
devem sempre apontar para dentro (ou seja, em direção às camadas mais abstratas e próximas da lógica de negócio). Nada
em uma camada interna pode depender de algo de uma camada externa, incluindo funções, classes, variáveis ou qualquer
entidade nomeada.

Além disso, formatos de dados usados em camadas externas não devem ser utilizados diretamente nas camadas internas.
Assim, mudanças no exterior não afetam o núcleo do sistema.

### Camadas da Clean Architecture

#### Entities (Entidades)

As entidades encapsulam as regras de negócio fundamentais do sistema, ou seja, aquelas que seriam aplicáveis em
diferentes contextos ou em várias aplicações dentro de uma organização. Por exemplo:

- Podem ser objetos com métodos ou conjuntos de estruturas de dados e funções.
- Não devem ser afetadas por mudanças operacionais do sistema, como navegabilidade de páginas ou mudanças na segurança.

#### Use Cases (Casos de Uso)

Essa camada contém as regras de negócio específicas da aplicação, encapsulando toda a lógica dos casos de uso do
sistema. Seu propósito é orquestrar o fluxo de dados:

- Dirigem as entidades para utilizarem suas regras internas e atingirem os objetivos.
- São isoladas de dependências externas, como banco de dados ou frameworks.
- Mudanças nas operações da aplicação impactam apenas essa camada.

#### Interface Adapters (Adaptadores de Interface)

Essa camada contém adaptadores responsáveis por:

- Converter os dados no formato mais conveniente para uso em outras camadas, como a de casos de uso e a de entidades.
- Conter a arquitetura MVC (Model-View-Controller) de uma GUI, assim como conversões de dados para persistência (banco
  de dados) ou serviços externos.

Restrições importantes:

- SQL ou outros detalhes do banco de dados são restritos a esta camada.
- Nenhuma camada interna deve possuir conhecimento específico sobre bancos de dados ou frameworks externos.

#### Frameworks and Drivers (Frameworks e Drivers)

A camada mais externa contém os detalhes concretos do sistema, como:

- Comunicação com banco de dados.
- Frameworks web.
- Outras ferramentas externas.

Essa camada serve como "cola" para conectar os elementos internos às tecnologias externas e, em geral, possui o menor
volume de código.

### Travessia de Fronteiras das Camadas

A comunicação entre as camadas é feita por **Inversão de Dependência**:

1. Camadas internas expõem interfaces que devem ser implementadas por camadas externas.
2. Por exemplo, um *Use Case* pode chamar o *Presenter* por meio de uma interface, sem conhece-la diretamente.

Dessa forma:

- As dependências seguem sempre a regra de apontar para dentro.
- O fluxo de controle é direcionado para camadas externas quando necessário, mantendo a abstração e modularidade.

### Dados que cruzam fronteiras

Os dados trocados entre camadas devem ser simples e isolados, como:

- Estruturas básicas de dados (*structs* ou DTOs - Data Transfer Objects).
- Mapas (*hashmaps*) ou objetos construídos para transferência de dados.
- Devem evitar dependências externas, como estruturas geradas por frameworks.

Por exemplo:

- Estruturas de linhas de banco de dados (*Row Structures*) não devem ser passadas diretamente para camadas internas.
- Dados devem ser convertidos em um formato independente antes de cruzar fronteiras.

### Conclusão

Seguir as regras simples da Clean Architecture ajuda a criar sistemas altamente coesos, testáveis e independentes das
tecnologias externas. Ao manter as responsabilidades bem definidas em camadas e aderir à Regra de Dependência, reduzimos
significativamente as dores de cabeça futuras ao escalar, modificar ou substituir partes do sistema em resposta a novas
demandas.
