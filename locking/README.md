# **Optimistic Locking vs Pessimistic Locking**

Em sistemas de banco de dados e concorrência, o **locking** (bloqueio) é um mecanismo que controla o acesso simultâneo a
dados compartilhados. Ele é utilizado para evitar condições de corrida e garantir a integridade dos dados durante
transações concorrentes. Existem duas abordagens principais para implementar o bloqueio: **Optimistic Locking** e *
*Pessimistic Locking**.

## **1. Pessimistic Locking**

### **Definição:**

O **Pessimistic Locking** assume que conflitos entre transações serão comuns. Portanto, ele bloqueia os recursos (dados)
durante o processo de leitura e escrita para garantir que nenhum outro processo ou transação interfira durante o seu
uso.

### **Como Funciona:**

Quando uma transação começa, ela bloqueia explicitamente o recurso (como uma linha de uma tabela) para garantir que
nenhuma outra transação possa modificar esses dados até que a transação em questão seja concluída. Isso garante que a
transação seja realizada de forma isolada e sem interferência.

#### **Exemplo:**

1. A Transação A começa e bloqueia uma linha no banco de dados.
2. A Transação B tenta acessar a mesma linha, mas o banco de dados impede que ela faça isso até que a Transação A
   termine e libere o bloqueio.
3. A Transação A realiza a atualização e finaliza a transação, liberando o bloqueio.
4. Agora, a Transação B pode acessar a linha e realizar a atualização.

### **Vantagens:**

- **Segurança:** Garante que não haja conflitos ou condições de corrida, pois as transações não podem acessar dados
  bloqueados simultaneamente.
- **Simplicidade:** A lógica de implementação é direta e fácil de entender.

### **Desvantagens:**

- **Desempenho:** Pode gerar gargalos, pois as transações ficam bloqueadas e não podem ser acessadas por outros
  processos, aumentando a latência.
- **Escalabilidade:** Em sistemas de alta concorrência, pode resultar em muitos bloqueios, prejudicando a performance do
  sistema.

---

## **2. Optimistic Locking**

### **Definição:**

O **Optimistic Locking** assume que conflitos entre transações serão raros. Em vez de bloquear explicitamente os dados,
ele permite que múltiplas transações acessem os dados simultaneamente, mas verifica se houve alteração antes de
confirmar a transação.

### **Como Funciona:**

Ao iniciar uma transação, não é feito um bloqueio explícito sobre os dados. Em vez disso, a transação mantém um
registro (geralmente uma **versão** ou um **timestamp**) dos dados no momento em que os acessou. Quando a transação
tenta salvar os dados, o sistema verifica se a versão ou timestamp dos dados foi alterado por outra transação. Se houver
uma alteração (conflito), a transação falha e deve ser repetida.

#### **Exemplo:**

1. A Transação A lê os dados de uma linha e recebe a versão ou timestamp.
2. A Transação B também lê os mesmos dados e recebe a versão ou timestamp.
3. A Transação A faz alterações e tenta gravar, verificando se a versão dos dados ainda é a mesma.
4. A Transação A grava os dados com sucesso, se não houver conflito.
5. A Transação B tenta gravar os dados, mas o sistema detecta que a versão foi modificada por A, e a transação falha.
6. A Transação B deve lidar com o conflito, geralmente tentando refazer a operação.

### **Vantagens:**

- **Desempenho:** Como não há bloqueios explícitos, o sistema é mais escalável e pode manipular mais transações
  simultâneas sem um grande impacto de desempenho.
- **Escalabilidade:** É mais eficiente em ambientes com alta concorrência, pois não impede que múltiplas transações
  acessem os dados ao mesmo tempo.

### **Desvantagens:**

- **Complexidade:** A implementação é mais complexa, pois o sistema precisa monitorar versões ou timestamps e verificar
  conflitos antes de permitir a confirmação da transação.
- **Risco de Conflito:** Se houver muitos conflitos entre transações, pode ser necessário repetir várias vezes a mesma
  transação, o que pode aumentar a carga do sistema e diminuir a eficiência.

---

## **Comparação entre Pessimistic Locking e Optimistic Locking**

| Característica                    | **Pessimistic Locking**                                     | **Optimistic Locking**                              |
|-----------------------------------|-------------------------------------------------------------|-----------------------------------------------------|
| **Mecanismo**                     | Bloqueio explícito de recursos durante a transação.         | Não há bloqueio explícito; usa versões/timestamps.  |
| **Conflitos de Concorrência**     | Prevê conflitos e impede o acesso simultâneo.               | Assume que conflitos são raros e verifica no final. |
| **Desempenho**                    | Pode ser afetado por bloqueios, reduzindo a escalabilidade. | Melhor desempenho e escalabilidade.                 |
| **Segurança**                     | Garante a consistência absoluta.                            | Pode falhar se houver muitos conflitos.             |
| **Simplicidade de Implementação** | Simples de implementar, direta.                             | Mais complexo, exige controle de versões.           |

---

## **Conclusão**

A escolha entre **Pessimistic Locking** e **Optimistic Locking** depende dos requisitos específicos do sistema:

- **Pessimistic Locking** é mais adequado quando a probabilidade de conflitos é alta ou quando a integridade dos dados
  deve ser garantida sem falhas.
- **Optimistic Locking** é preferível em ambientes com alta concorrência onde os conflitos são raros e a performance e
  escalabilidade são mais críticas.

Cada abordagem tem suas vantagens e desvantagens, e a escolha ideal depende das características da aplicação e da carga
de trabalho que ela precisa suportar.