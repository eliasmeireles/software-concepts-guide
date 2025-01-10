# **Transações (Atomicidade, Isolamento)**

Em sistemas de gerenciamento de banco de dados (SGBD), **transações** são unidades de trabalho que consistem em uma ou
mais operações, como leitura e escrita de dados. Essas operações precisam ser executadas de forma que o sistema sempre
se encontre em um estado consistente, mesmo diante de falhas ou concorrência de acesso. A garantia da integridade dos
dados é fundamental, e para isso, conceitos como **atomicidade** e **isolamento** são utilizados dentro do paradigma das
transações.

## **1. O que é uma Transação?**

Uma transação é um conjunto de operações que deve ser executado de forma completa ou não executado de forma alguma. Ela
é usada para garantir a integridade dos dados em bancos de dados relacionais e pode envolver várias operações, como
inserir, atualizar ou deletar registros. O principal objetivo de uma transação é garantir que o banco de dados se
mantenha consistente, mesmo quando ocorrem falhas, e que as operações concorrentes não interfiram umas nas outras.

### **Propriedades das Transações - ACID**

O conceito de **ACID** (Atomicidade, Consistência, Isolamento e Durabilidade) define as propriedades essenciais que uma
transação deve ter para garantir a integridade e confiabilidade do banco de dados.

1. **Atomicidade (Atomicity)**
2. **Consistência (Consistency)**
3. **Isolamento (Isolation)**
4. **Durabilidade (Durability)**

Aqui, vamos focar em **Atomicidade** e **Isolamento**, que são duas das propriedades mais importantes.

## **2. Atomicidade (Atomicity)**

**Atomicidade** é a propriedade que garante que todas as operações dentro de uma transação sejam completadas com sucesso
ou, caso contrário, nenhuma delas seja aplicada. Se uma transação falhar em algum ponto, todas as alterações feitas até
aquele momento são revertidas, mantendo o banco de dados em seu estado original. Em outras palavras, uma transação é
atômica: ou ela é completamente concluída, ou não é executada de forma alguma.

### **Exemplo de Atomicidade:**

Imagine uma transação que envolve duas operações:

- Transferir dinheiro de uma conta para outra.
    - Debitar R$100 da Conta A.
    - Creditar R$100 na Conta B.

Se a operação de débito for bem-sucedida, mas a operação de crédito falhar (por qualquer motivo), o sistema deve
garantir que o valor debitado da Conta A seja revertido para manter a consistência dos dados. Isso é feito através do
conceito de rollback (desfazer a transação).

#### **Características da Atomicidade:**

- **Commit:** Quando todas as operações da transação são bem-sucedidas, a transação é confirmada e suas mudanças são
  aplicadas permanentemente ao banco de dados.
- **Rollback:** Se houver qualquer falha durante a transação, todas as mudanças feitas até aquele ponto são revertidas.

### **Vantagens da Atomicidade:**

- **Consistência garantida:** A atomicidade assegura que o banco de dados nunca ficará em um estado parcialmente
  modificado.
- **Recuperação de falhas:** Caso ocorra uma falha, a transação pode ser revertida, evitando que dados corrompidos sejam
  persistidos.

## **3. Isolamento (Isolation)**

**Isolamento** é a propriedade que garante que as transações sejam executadas de maneira isolada, ou seja, que os
efeitos de uma transação não sejam visíveis para outras até que ela seja completada. O isolamento assegura que as
transações concorrentes não afetem umas às outras de maneira indesejada, evitando problemas como leitura de dados
intermediários ou inconsistentes.

Existem diferentes níveis de isolamento que controlam o comportamento de transações concorrentes. Eles são definidos no
padrão SQL por meio de um conceito chamado **nível de isolamento de transação**. A escolha do nível de isolamento
determina o grau de visibilidade das operações de uma transação para outras transações que estão sendo executadas
simultaneamente.

### **Níveis de Isolamento (SQL Standard)**

1. **Read Uncommitted (Leitura Não Confirmada):**
    - Transações podem ler dados que ainda não foram confirmados por outra transação (dados "sujo").
    - **Problema:** Pode ocorrer a leitura de dados inconsistentes, chamados de "dirty reads".

2. **Read Committed (Leitura Confirmada):**
    - Transações só podem ler dados que foram confirmados por outras transações.
    - **Problema:** Ainda pode ocorrer a "non-repeatable read" (leitura não repetível), onde o valor lido em uma
      transação pode mudar quando lido novamente.

3. **Repeatable Read (Leitura Repetível):**
    - Garante que os dados lidos por uma transação não sejam modificados por outras transações até que a transação seja
      concluída.
    - **Problema:** Pode ocorrer a "phantom read" (leitura fantasma), onde uma nova linha pode ser inserida entre duas
      leituras na mesma transação.

4. **Serializable (Serializável):**
    - O nível mais alto de isolamento, onde as transações são completamente isoladas umas das outras. Garante que o
      sistema se comporte como se as transações fossem executadas uma após a outra, sem sobreposição.
    - **Vantagem:** Garante que não ocorram condições de corrida, como dirty reads, non-repeatable reads ou phantom
      reads.
    - **Desvantagem:** Pode afetar a performance, pois limita a concorrência.

### **Problemas com Transações Concorrentes:**

- **Dirty Read (Leitura Suja):** Uma transação lê dados que foram modificados por outra transação, mas ainda não foram
  confirmados. Se a outra transação for revertida, os dados lidos pela primeira transação são inválidos.
- **Non-repeatable Read (Leitura Não Repetível):** Uma transação lê um valor, mas outra transação modifica esse valor
  antes que a primeira transação o leia novamente. Isso causa resultados inconsistentes.
- **Phantom Read (Leitura Fantasma):** Uma transação lê um conjunto de dados, mas outra transação insere, modifica ou
  apaga registros durante a execução da primeira. A transação original pode então ver um conjunto de dados diferente.

## **4. Relacionamento entre Atomicidade e Isolamento**

Enquanto a **atomicidade** garante que as transações sejam tratadas como unidades indivisíveis (ou são completadas ou
revertidas), o **isolamento** garante que as transações concorrentes não interfiram umas nas outras, proporcionando um
ambiente no qual as operações de uma transação não afetam outras até que elas sejam finalizadas.

A atomicidade protege a transação de falhas, enquanto o isolamento garante que as transações concorrentes não alterem os
dados de maneira imprevisível.

## **5. Conclusão**

A **atomicidade** e o **isolamento** são propriedades fundamentais para garantir a integridade e a confiabilidade de
transações em bancos de dados. A atomicidade assegura que uma transação seja tratada como uma unidade indivisível,
enquanto o isolamento controla como as transações concorrentes interagem para evitar problemas como leituras
inconsistentes e condições de corrida.

Essas propriedades são essenciais para garantir que os sistemas de banco de dados possam lidar com múltiplas transações
simultâneas de forma segura e eficiente, preservando a consistência dos dados e proporcionando uma experiência de uso
confiável para as aplicações.