# **Bases de Dados Não Relacionais (NoSQL)**

As **bases de dados não relacionais** (ou **NoSQL**, que significa "Not Only SQL") são sistemas de gerenciamento de
banco de dados que não seguem o modelo relacional tradicional. Elas foram desenvolvidas para atender a casos de uso que
os bancos de dados relacionais não conseguem escalar ou lidar de forma eficiente, como grandes volumes de dados, dados
não estruturados ou semi-estruturados, e alta escalabilidade horizontal.

As bases de dados não relacionais são flexíveis e podem ser adaptadas a uma ampla gama de modelos de dados. Elas não
utilizam SQL como linguagem principal e muitas vezes não possuem o esquema rígido encontrado nos bancos de dados
relacionais.

## **1. Definição**

As **bases de dados não relacionais** são caracterizadas por não seguirem o modelo de tabelas, linhas e colunas, como
ocorre nas bases de dados relacionais. Elas podem armazenar dados de diferentes formas, incluindo documentos, pares
chave-valor, grafos ou colunas. O principal objetivo dessas bases de dados é oferecer maior flexibilidade, desempenho e
escalabilidade, especialmente quando se trabalha com grandes volumes de dados.

## **2. Tipos de Bancos de Dados Não Relacionais**

Existem diferentes tipos de bases de dados não relacionais, cada uma adequada a cenários específicos de uso. A seguir,
veremos os tipos mais comuns:

### **1. Banco de Dados de Documentos**

- **Exemplo:** MongoDB, CouchDB
- **Estrutura:** Armazenam dados no formato de documentos, geralmente em JSON ou BSON (uma versão binária do JSON). Cada
  documento pode ter uma estrutura diferente, o que proporciona flexibilidade para armazenar dados não estruturados ou
  semi-estruturados.

  **Exemplo de documento JSON em MongoDB:**
  ```json
  {
    "_id": "12345",
    "nome": "João",
    "idade": 25,
    "enderecos": [
      { "cidade": "São Paulo", "estado": "SP" },
      { "cidade": "Rio de Janeiro", "estado": "RJ" }
    ]
  }
  ```

- **Vantagens:**
  - Flexibilidade na modelagem de dados.
  - Suporta dados aninhados e estruturas complexas.
  - Adequado para aplicativos com dados dinâmicos e que mudam frequentemente.

### **2. Banco de Dados de Chave-Valor**

- **Exemplo:** Redis, DynamoDB
- **Estrutura:** Armazenam dados como pares **chave-valor**. A chave é única e pode ser usada para buscar o valor
  associado. Esses bancos são rápidos e altamente escaláveis.

  **Exemplo de chave-valor em Redis:**
  ```text
  "user:12345" => {"nome": "João", "idade": 25}
  ```

- **Vantagens:**
  - Simples e rápidos para acesso a dados.
  - Escalabilidade horizontal fácil.
  - Ótimos para armazenar sessões de usuários, caches e dados temporários.

### **3. Banco de Dados de Colunas**

- **Exemplo:** Cassandra, HBase
- **Estrutura:** Armazenam dados em **colunas** em vez de linhas, o que pode ser mais eficiente para consultas em
  grandes volumes de dados. Cada linha pode ter diferentes colunas.

  **Exemplo de estrutura de dados em Cassandra:**
  ```text
  id      | nome   | idade  | cidade
  -------------------------------------
  1       | João   | 25     | São Paulo
  2       | Maria  | 30     | Rio de Janeiro
  ```

- **Vantagens:**
  - Excelente desempenho para consultas de leitura em larga escala.
  - Ideal para dados analíticos e em tempo real.
  - Oferece boa escalabilidade horizontal.

### **4. Banco de Dados de Grafos**

- **Exemplo:** Neo4j, ArangoDB
- **Estrutura:** Armazenam dados em formato de **grafos**, onde os dados são representados como nós (entidades) e
  arestas (relações). Ideal para modelar dados inter-relacionados, como redes sociais, recomendações e redes de
  computadores.

  **Exemplo de grafo em Neo4j:**
  ```text
  (João)-[:AMIGO_DE]->(Maria)
  ```

- **Vantagens:**
  - Eficiência em consultas que envolvem relacionamentos complexos.
  - Adequado para redes sociais, sistemas de recomendação e aplicativos de análise de grafos.
  - Facilita a navegação entre os dados inter-relacionados.

### **5. Banco de Dados de Armazenamento em Eventos (Event Store)**

- **Exemplo:** EventStoreDB
- **Estrutura:** São especializados em armazenar eventos ou logs de ações de sistemas, com a ideia de que o estado de um
  sistema pode ser reconstruído a partir da sequência de eventos que ocorreram.

  **Exemplo de evento armazenado:**
  ```json
  {
    "evento": "compra_realizada",
    "data": "2023-01-01T12:00:00",
    "dados": {
      "usuario": "João",
      "produto": "Smartphone",
      "quantidade": 1
    }
  }
  ```

- **Vantagens:**
  - Ideal para sistemas de processamento de eventos.
  - Permite a auditoria e rastreabilidade de ações no sistema.
  - Pode ser usado em sistemas baseados em CQRS (Command Query Responsibility Segregation).

## **3. Vantagens das Bases de Dados Não Relacionais**

- **Escalabilidade Horizontal:** Bancos de dados NoSQL são projetados para escalar horizontalmente, distribuindo dados
  por várias máquinas. Isso é crucial para lidar com grandes volumes de dados e tráfego em tempo real.
- **Desempenho em Grande Escala:** Bancos NoSQL, como o MongoDB ou Cassandra, são otimizados para consultas de leitura e
  escrita em grande escala.
- **Flexibilidade de Modelagem de Dados:** Ao contrário dos bancos relacionais, que exigem esquemas rígidos, bancos de
  dados NoSQL permitem modelos de dados dinâmicos e altamente flexíveis, ideais para dados semi-estruturados ou não
  estruturados.
- **Alta Disponibilidade:** Muitos bancos NoSQL são projetados com replicação automática, oferecendo alta
  disponibilidade e resistência a falhas.

## **4. Desvantagens das Bases de Dados Não Relacionais**

- **Falta de Conformidade com ACID:** Muitos bancos NoSQL não oferecem garantias completas de **ACID** (Atomicidade,
  Consistência, Isolamento e Durabilidade) como os bancos relacionais. Isso pode ser um problema para aplicações que
  exigem transações rigorosas.
- **Consistência Eventual:** Algumas bases NoSQL utilizam o modelo de **consistência eventual**, o que significa que, em
  vez de garantir a consistência imediata, elas garantem que, após algum tempo, os dados se tornem consistentes.
- **Consultas Complexas:** Bancos NoSQL podem ser menos eficientes em consultas complexas e junções entre dados, que são
  fáceis de realizar em bancos relacionais com SQL.
- **Menos Ferramentas de Suporte:** Embora esteja mudando, os bancos de dados NoSQL ainda têm menos ferramentas de
  suporte em comparação com bancos relacionais, que têm uma comunidade e ecossistema muito mais maduros.

## **5. Exemplos de Bancos de Dados Não Relacionais Populares**

- **MongoDB:** Um banco de dados de documentos muito popular que usa BSON (uma forma binária de JSON) para armazenar
  dados. Muito utilizado em aplicativos web e mobile.
- **Cassandra:** Um banco de dados de colunas altamente escalável, adequado para sistemas que exigem alta
  disponibilidade e grandes volumes de dados.
- **Redis:** Um banco de dados de chave-valor extremamente rápido e eficiente, frequentemente usado para caching e
  armazenamento de sessões.
- **Neo4j:** Um banco de dados de grafos, ideal para modelar e consultar relacionamentos complexos.
- **DynamoDB:** Um banco de dados de chave-valor altamente escalável oferecido pela AWS, ideal para aplicativos com alta
  demanda de leitura e gravação.

## **6. Casos de Uso Típicos para Bancos de Dados Não Relacionais**

- **Big Data:** Bancos NoSQL são ideais para trabalhar com grandes volumes de dados, como logs de servidores, dados de
  sensores ou dados de redes sociais.
- **Aplicações em Tempo Real:** Sistemas de recomendação, como os usados por e-commerce e streaming, podem se beneficiar
  de bancos de dados NoSQL para oferecer consultas rápidas em grandes volumes de dados.
- **Armazenamento de Dados Semi-Estruturados:** Quando os dados não possuem um esquema fixo ou mudam frequentemente (
  como dados de usuário em aplicativos), bancos de dados de documentos são uma escolha popular.
- **Redes Sociais e Conectividade:** Bancos de dados de grafos são amplamente usados para representar redes sociais,
  onde os relacionamentos entre os usuários são essenciais.

---

## **Conclusão**

As **bases de dados não relacionais (NoSQL)** oferecem flexibilidade e escalabilidade para lidar com grandes volumes de
dados, dados não estruturados e sistemas distribuídos. Embora possuam limitações em relação à consistência transacional
e à complexidade das consultas, elas são ideais para cenários que exigem alta disponibilidade, performance e modelagem
de dados flexível. A escolha entre bancos de dados relacionais e não relacionais depende das necessidades específicas da
aplicação, como volume de dados, modelo de dados e requisitos de consistência.