# **Banco de Dados Relacional vs Não Relacional: Quando Usar e Por Quê**

Ao decidir entre usar um **banco de dados relacional (RDBMS)** ou **não relacional (NoSQL)**, é crucial considerar as
necessidades específicas do projeto, os requisitos de escalabilidade, flexibilidade e a natureza dos dados que você
precisa gerenciar. A seguir, apresentamos um detalhamento completo sobre os **dois tipos de banco de dados**, seus usos
ideais e as razões para escolher um ou outro.

---

## **1. [Bancos de Dados Relacionais (RDBMS)](./relational-db/README.md)**

Os **bancos de dados relacionais** são baseados no modelo de dados relacional, onde os dados são organizados em tabelas
com linhas e colunas. Esses sistemas utilizam a linguagem **SQL (Structured Query Language)** para manipulação e
consulta de dados. Exemplos populares de RDBMS incluem **MySQL**, **PostgreSQL**, **Oracle** e **SQL Server**.

### **Características Principais de RDBMS:**

- **Estrutura Tabular:** Dados são organizados em tabelas (ou relações) com linhas (registros) e colunas (atributos).
  Cada tabela possui uma chave primária (PK) única.
- **Integridade Referencial:** Relacionamentos entre tabelas são mantidos por meio de **chaves estrangeiras (FK)**, o
  que assegura a integridade e consistência dos dados.
- **ACID:** Suporte a transações com propriedades **ACID** (Atomicidade, Consistência, Isolamento, Durabilidade),
  garantindo que as transações sejam seguras e consistentes.
- **Schema Definido:** O esquema (estrutura de dados) do banco é predefinido, o que significa que você precisa definir
  as tabelas e tipos de dados antes de inserir os dados.

### **Quando Usar Bancos de Dados Relacionais?**

- **Estrutura de Dados Bem Definida:** Quando os dados possuem uma estrutura rígida e predefinida, como em sistemas de *
  *gestão financeira**, **ERP**, **CRM**, **sistemas bancários**, e **e-commerce**.
- **Relacionamentos Complexos entre Dados:** Quando você precisa armazenar dados que estão fortemente relacionados entre
  si, como **clientes**, **pedidos**, **produtos** e **pagamentos**.
- **Consistência e Integridade dos Dados:** Quando a consistência dos dados e a integridade referencial são de extrema
  importância, por exemplo, em sistemas bancários onde transações precisam ser garantidas e auditáveis.
- **Consultas Complexas e Join Operations:** Quando há necessidade de consultas complexas que envolvem múltiplas
  tabelas, junções (**JOINs**) e agregações, como **relatórios financeiros**, **análises de dados** e **consultas
  transacionais**.
- **Transações ACID:** Quando a aplicação exige **garantia de transações atômicas** e quer garantir que todas as
  operações sejam executadas de forma confiável e segura (ex.: transferências bancárias, compras online).
- **Exigências de Conformidade:** Em setores regulamentados, como o **financeiro**, onde há requisitos rigorosos sobre
  integridade e rastreabilidade dos dados.

### **Vantagens de Bancos de Dados Relacionais:**

- **Consultas Complexas:** Capacidade de realizar consultas complexas e transações com muitas interações entre os dados.
- **ACID e Consistência:** Garantia de transações seguras, com forte consistência.
- **Ferramentas Maturadas:** RDBMSs possuem ferramentas robustas de administração e análise, como **backup**, *
  *restauração**, **replicação** e **monitoramento**.

---

## **2. [Bancos de Dados Não Relacionais (NoSQL)](./non-relational-db/README.md)**

Os **bancos de dados não relacionais (NoSQL)** são uma categoria que abrange diferentes tipos de sistemas de banco de
dados que não utilizam o modelo relacional de tabelas e linhas. Esses bancos de dados são projetados para trabalhar com
dados mais flexíveis e escaláveis. Exemplos incluem **MongoDB**, **Cassandra**, **Redis**, **Couchbase** e **DynamoDB**.

### **Características Principais de NoSQL:**

- **Flexibilidade de Dados:** Não requer um esquema fixo, permitindo que os dados sejam armazenados de forma flexível (
  como documentos JSON, chave-valor, grafos ou colunas).
- **Escalabilidade Horizontal:** A maioria dos bancos NoSQL é projetada para escalar horizontalmente (adicionando mais
  servidores), o que os torna adequados para grandes volumes de dados e aplicações distribuídas.
- **Desempenho em Grandes Volumes de Dados:** Ideal para lidar com grandes volumes de dados não estruturados e dados de
  alta variabilidade, como **logs de eventos**, **feeds de redes sociais**, **sensores IoT** e **dados de dispositivos
  móveis**.
- **Falta de Suporte a ACID:** Em muitos bancos NoSQL, a ênfase está na escalabilidade e desempenho, e não em garantias
  de transações ACID. Em vez disso, seguem um modelo de consistência eventual (eventual consistency).

### **Quando Usar Bancos de Dados Não Relacionais?**

- **Dados Não Estruturados ou Semiestruturados:** Quando os dados não se ajustam bem a um esquema fixo, como em **redes
  sociais**, **logs de aplicação**, **feeds de notícias** ou **arquivos de mídia**.
- **Escalabilidade Horizontal:** Quando a necessidade de escalabilidade é grande, como em **aplicações web e móveis**
  que devem lidar com grandes volumes de dados e acessos simultâneos (ex.: plataformas de streaming, e-commerce de
  grande escala).
- **Desempenho e Velocidade:** Quando a aplicação exige altíssimo desempenho e resposta rápida, como em sistemas de *
  *cache** e **tempo real** (ex.: **Redis** ou **Cassandra**).
- **Desempenho para Consultas Simples:** Quando a aplicação não exige consultas complexas, mas sim operações rápidas de
  leitura e escrita, como **armazenamento de sessões de usuário** ou **listas de favoritos**.
- **Alta Disponibilidade e Tolerância a Falhas:** Quando você precisa de alta disponibilidade e tolerância a falhas em
  uma aplicação distribuída, como sistemas de **recomendação**, **IoT** ou **dispositivos móveis**.
- **Mudanças Frequentes no Modelo de Dados:** Quando o esquema de dados pode mudar frequentemente, como em ambientes
  ágeis que necessitam de **iterações rápidas** de desenvolvimento.

### **Vantagens de Bancos de Dados Não Relacionais:**

- **Escalabilidade Horizontal:** Suporte para escalar para grandes volumes de dados e usuários, distribuindo a carga por
  vários servidores ou clusters.
- **Desempenho em Grandes Volumes de Dados:** Capacidade de lidar com grandes quantidades de dados em tempo real, com
  baixa latência.
- **Flexibilidade no Modelo de Dados:** Facilidade para armazenar dados não estruturados, como documentos JSON, imagens,
  vídeos e dados em formato binário.
- **Baixo Custo de Escalabilidade:** Em muitos casos, sistemas NoSQL podem ser mais baratos para escalar, pois podem
  rodar em infraestrutura distribuída mais simples.

---

## **3. Quando Usar um Banco Relacional vs Não Relacional?**

| **Critério**                    | **Banco de Dados Relacional**                                       | **Banco de Dados Não Relacional**                                     |
|---------------------------------|---------------------------------------------------------------------|-----------------------------------------------------------------------|
| **Estrutura dos Dados**         | Dados bem estruturados (tabelas e relações).                        | Dados flexíveis e não estruturados (documentos, chave-valor, grafos). |
| **Modelo de Dados**             | Esquema fixo e normalização de dados.                               | Sem esquema fixo, dados semiestruturados.                             |
| **Consultas Complexas**         | Alta capacidade para consultas SQL complexas (joins, agrupamentos). | Consultas simples e rápidas; limitações em joins complexos.           |
| **Escalabilidade**              | Escalabilidade vertical (potência de hardware).                     | Escalabilidade horizontal (distribuição entre múltiplos servidores).  |
| **Transações**                  | Suporte completo a transações ACID.                                 | Suporte a consistência eventual, menos garantias ACID.                |
| **Desempenho em Grande Escala** | Não ideal para grandes volumes de dados e alta carga.               | Ideal para grandes volumes e cargas distribuídas.                     |
| **Exemplo de Uso**              | Sistemas financeiros, ERP, CRM.                                     | Redes sociais, logs, aplicações em tempo real, IoT.                   |
| **Consistência de Dados**       | Alta consistência e integridade referencial.                        | Consistência eventual (dependendo da implementação).                  |
| **Custo de Infraestrutura**     | Pode ser mais caro para escalar.                                    | Menor custo de escalabilidade, especialmente para grandes volumes.    |

---

## **Conclusão**

A escolha entre **banco de dados relacional** e **não relacional** depende diretamente dos requisitos específicos de *
*escalabilidade**, **complexidade de dados**, **transações**, **consultas** e **flexibilidade** do projeto.

- **Use bancos de dados relacionais** quando a integridade dos dados, transações complexas e consultas SQL são
  essenciais, e quando a estrutura de dados é bem definida e estável.
- **Use bancos de dados não relacionais** quando você precisa de flexibilidade de dados, grande escalabilidade e
  desempenho em sistemas distribuídos, especialmente para dados não estruturados ou semi-estruturados.

Ambos os tipos de banco de dados têm seus méritos, e muitas vezes, as empresas optam por uma **abordagem híbrida**,
usando ambos em diferentes partes do sistema, conforme as necessidades de cada componente ou microserviço.