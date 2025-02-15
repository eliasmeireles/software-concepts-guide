# **Bancos de Dados Relacionais (RDBMS)**

As **bases de dados relacionais** (RDBMS - Relational Database Management Systems) são sistemas de gerenciamento de
banco de dados que armazenam dados de forma estruturada em tabelas, que são compostas por linhas e colunas. O modelo
relacional foi introduzido por **Edgar F. Codd** em 1970 e é baseado em teoria matemática de conjuntos e álgebra
relacional. As bases de dados relacionais são amplamente utilizadas em aplicações empresariais e sistemas que exigem
consistência e integridade de dados.

## **1. Definição**

Em um **banco de dados relacional**, os dados são organizados em **tabelas**, onde cada tabela é composta por **linhas (
tuplas)** e **colunas (atributos)**. Cada linha representa uma entidade única, e cada coluna armazena um tipo de dado
específico sobre essa entidade.

As bases de dados relacionais seguem a **linguagem SQL (Structured Query Language)** para consulta, inserção,
atualização e exclusão de dados.

### **Exemplo de Estrutura Relacional:**

| **ID** | **Nome** | **Idade** | **Cidade**     |
|--------|----------|-----------|----------------|
| 1      | João     | 25        | São Paulo      |
| 2      | Maria    | 30        | Rio de Janeiro |
| 3      | Carlos   | 28        | Belo Horizonte |

- **Tabelas:** Cada entidade (ex.: pessoas, produtos, vendas) é representada por uma tabela.
- **Colunas:** Atributos ou características da entidade (ex.: Nome, Idade, Cidade).
- **Linhas:** Instâncias específicas de dados.

## **2. Princípios Fundamentais das Bases de Dados Relacionais**

### **Modelo Relacional**

O modelo relacional organiza os dados em tabelas, e cada tabela pode se relacionar com outras tabelas, usando **chaves
primárias** e **chaves estrangeiras** para garantir a integridade e as relações entre os dados.

- **Chave Primária (Primary Key):** É um campo ou conjunto de campos que identificam de forma única cada registro em uma
  tabela.
- **Chave Estrangeira (Foreign Key):** É um campo em uma tabela que se refere à chave primária de outra tabela,
  estabelecendo uma relação entre elas.

### **Integridade dos Dados**

As bases de dados relacionais garantem a **integridade referencial** e **integridade de dados** através de restrições,
como:

- **Integridade de Entidade:** Garante que cada linha da tabela seja única, usando chaves primárias.
- **Integridade Referencial:** Garante que, quando uma chave estrangeira é usada, ela corresponda a uma chave primária
  existente em outra tabela.

### **Relacionalidade**

As tabelas em um banco de dados relacional são **relacionadas** entre si. Isso significa que é possível fazer *
*junções (JOINs)** entre tabelas para combinar dados de diferentes fontes de forma eficiente.

### **Normalização**

A **normalização** é o processo de organizar as tabelas de um banco de dados para evitar redundâncias e dependências
desnecessárias, garantindo uma estrutura eficiente e sem anomalias.

## **3. Operações Básicas em Bancos de Dados Relacionais**

As operações em um banco de dados relacional são realizadas usando SQL (Structured Query Language). As operações mais
comuns incluem:

### **Operações de Manipulação de Dados (DML - Data Manipulation Language):**

- **SELECT:** Usado para consultar dados.
  ```sql
  SELECT * FROM clientes WHERE idade > 30;
  ```

- **INSERT:** Usado para inserir novos dados na tabela.
  ```sql
  INSERT INTO clientes (nome, idade, cidade) VALUES ('João', 25, 'São Paulo');
  ```

- **UPDATE:** Usado para atualizar dados existentes.
  ```sql
  UPDATE clientes SET idade = 26 WHERE nome = 'João';
  ```

- **DELETE:** Usado para excluir dados.
  ```sql
  DELETE FROM clientes WHERE nome = 'João';
  ```

### **Operações de Definição de Dados (DDL - Data Definition Language):**

- **CREATE TABLE:** Usado para criar uma nova tabela.
  ```sql
  CREATE TABLE clientes (
    id INT PRIMARY KEY,
    nome VARCHAR(100),
    idade INT,
    cidade VARCHAR(100)
  );
  ```

- **ALTER TABLE:** Usado para modificar a estrutura de uma tabela existente.
  ```sql
  ALTER TABLE clientes ADD email VARCHAR(100);
  ```

- **DROP TABLE:** Usado para excluir uma tabela.
  ```sql
  DROP TABLE clientes;
  ```

### **Operações de Controle de Dados (DCL - Data Control Language):**

- **GRANT:** Usado para conceder permissões de acesso.
  ```sql
  GRANT SELECT, INSERT ON clientes TO usuario;
  ```

- **REVOKE:** Usado para revogar permissões de acesso.
  ```sql
  REVOKE SELECT, INSERT ON clientes FROM usuario;
  ```

## **4. Tipos de Relacionamentos entre Tabelas**

Em um banco de dados relacional, as tabelas podem se relacionar de várias maneiras:

### **1. Relacionamento Um-para-Um (1:1)**

Um relacionamento onde cada linha de uma tabela está associada a no máximo uma linha de outra tabela.

Exemplo: Cada **pessoa** tem um único **passaporte**, e cada passaporte pertence a uma única pessoa.

### **2. Relacionamento Um-para-Muitos (1:N)**

Uma linha em uma tabela pode ser associada a várias linhas em outra tabela, mas cada linha da segunda tabela está
associada a no máximo uma linha da primeira tabela.

Exemplo: Um **cliente** pode fazer várias **compras**, mas cada compra pertence a um único cliente.

### **3. Relacionamento Muitos-para-Muitos (N:M)**

Quando várias linhas de uma tabela podem ser associadas a várias linhas de outra tabela. Esse tipo de relacionamento é
geralmente implementado através de uma tabela intermediária, chamada de **tabela de junção**.

Exemplo: **Estudantes** podem estar matriculados em várias **disciplinas**, e cada disciplina pode ter vários
estudantes.

## **5. Vantagens das Bases de Dados Relacionais**

- **Estrutura Bem Definida:** A estrutura baseada em tabelas proporciona um modelo claro e fácil de entender.
- **Integridade de Dados:** Regras de integridade garantem a consistência e precisão dos dados.
- **Suporte a SQL:** O uso da linguagem SQL para consultas e manipulação de dados é um padrão amplamente aceito.
- **Normalização:** Ajuda a reduzir a redundância e a aumentar a eficiência no armazenamento de dados.

## **6. Desvantagens das Bases de Dados Relacionais**

- **Escalabilidade Horizontal Limitada:** Bancos de dados relacionais podem ter dificuldades em escalar
  horizontalmente (distribuindo a carga de trabalho entre várias máquinas).
- **Desempenho:** Consultas complexas em grandes volumes de dados podem resultar em queda de desempenho, especialmente
  quando não são bem indexadas.
- **Flexibilidade Limitada:** O modelo de dados rígido de tabelas pode ser menos flexível para representar dados
  complexos ou não estruturados.

## **7. Exemplos de Bancos de Dados Relacionais Populares**

- **MySQL:** Um dos sistemas de gerenciamento de banco de dados relacionais mais populares, usado amplamente em
  aplicações web.
- **PostgreSQL:** Um RDBMS robusto e de código aberto, conhecido por suas funcionalidades avançadas e conformidade com
  os padrões SQL.
- **SQLite:** Um banco de dados relacional leve e autônomo, popular para dispositivos móveis e aplicativos de desktop.
- **Microsoft SQL Server:** Uma solução empresarial da Microsoft, amplamente usada em ambientes corporativos.
- **Oracle Database:** Uma das soluções mais poderosas e amplamente utilizadas em grandes empresas e aplicações de
  missão crítica.

---

## **Conclusão**

As **bases de dados relacionais** são a espinha dorsal de muitas aplicações empresariais, web e sistemas que necessitam
de integridade de dados, confiabilidade e um modelo bem estruturado. Embora existam limitações em termos de
escalabilidade horizontal e flexibilidade, seu modelo baseado em tabelas e SQL fornece uma solução sólida e eficiente
para muitos tipos de aplicação. A normalização, integridade referencial e a utilização de SQL tornam as RDBMS uma
escolha popular para gerenciamento de dados em ambientes organizacionais.