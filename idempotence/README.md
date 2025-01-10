# **Idempotência**

A **idempotência** é um conceito fundamental na ciência da computação, especialmente em sistemas distribuídos e APIs.
Ela descreve a propriedade de uma operação que pode ser executada várias vezes sem que o resultado final seja alterado
após a primeira execução.

Em outras palavras, uma operação **idempotente** é aquela que, quando realizada múltiplas vezes, tem o mesmo efeito que
se fosse realizada apenas uma vez, ou seja, ela não causa alterações indesejadas após a primeira execução bem-sucedida.

## **1. Definição de Idempotência**

### **O que é?**

Idempotência refere-se à propriedade de uma operação de produzir o mesmo resultado, independentemente de quantas vezes a
operação seja executada.

Por exemplo, se uma operação de **atualização de recurso** for idempotente, isso significa que, não importa quantas
vezes a solicitação de atualização seja feita, o estado final do recurso será o mesmo após a primeira execução.

### **Exemplo Prático:**

- Se você fizer uma **requisição HTTP PUT** para atualizar um recurso (como um perfil de usuário), a primeira vez que a
  requisição for feita, o recurso será atualizado. Se a mesma requisição for repetida, o recurso permanecerá no mesmo
  estado, pois ele já está atualizado com os dados enviados. Não haverá mudança adicional.

## **2. Exemplos de Operações Idempotentes**

- **GET:** Uma requisição HTTP `GET` para buscar um recurso é idempotente porque ela não altera o estado do recurso. Ela
  pode ser feita quantas vezes for necessário, mas o resultado será sempre o mesmo.

  Exemplo: Fazer uma requisição para buscar informações sobre um usuário, como `GET /user/123`. O resultado será sempre
  as mesmas informações do usuário, sem modificar nada.

- **PUT:** Uma requisição HTTP `PUT` para atualizar um recurso é idempotente, pois, independentemente de quantas vezes
  ela seja chamada com os mesmos dados, o estado do recurso não será alterado após a primeira requisição bem-sucedida.

  Exemplo: Se você fizer uma requisição `PUT /user/123` com os dados `{ "name": "João" }`, a primeira vez ela atualizará
  o nome para "João", e todas as chamadas subsequentes terão o mesmo efeito.

- **DELETE:** Uma requisição HTTP `DELETE` também pode ser idempotente, pois, se o recurso já foi excluído, uma segunda
  chamada para deletá-lo novamente não mudará o estado.

  Exemplo: Se você fizer uma requisição `DELETE /user/123` e o usuário for excluído, a segunda requisição para excluir o
  mesmo usuário não terá efeito, pois ele já foi removido.

## **3. Não Idempotentes vs Idempotentes**

| **Operação** | **Idempotente?** | **Justificativa**                                                                              |
|--------------|------------------|------------------------------------------------------------------------------------------------|
| **GET**      | Sim              | Não altera o estado do recurso.                                                                |
| **POST**     | Não              | Pode criar um novo recurso, causando mudanças a cada execução.                                 |
| **PUT**      | Sim              | Atualiza o recurso para um estado específico, sem mudar o resultado com múltiplas requisições. |
| **DELETE**   | Sim              | Deleta um recurso, e múltiplas requisições não alteram o estado (já excluído).                 |
| **PATCH**    | Não              | Pode modificar parcialmente um recurso, alterando-o com cada requisição.                       |

## **4. Importância da Idempotência**

### **Evitar Efeitos Colaterais**

A idempotência é crucial para evitar efeitos colaterais inesperados, como a criação de registros duplicados ou a
realização de uma ação de forma repetida, quando o objetivo é garantir que ela ocorra apenas uma vez.

### **Robustez de Sistemas Distribuídos**

Em sistemas distribuídos, onde as falhas de rede e a duplicação de requisições podem ocorrer, a idempotência permite que
as requisições sejam repetidas sem causar problemas, garantindo que o sistema se recupere de falhas de maneira suave.

### **Confiabilidade nas APIs**

Quando uma API é idempotente, os desenvolvedores podem confiar nela, pois sabem que podem fazer várias tentativas sem
que isso gere efeitos adversos. Isso é especialmente importante em sistemas que fazem chamadas para APIs externas, onde
a latência e a possibilidade de falhas podem ocorrer.

## **5. Implementação de Idempotência**

Em APIs e sistemas, a implementação de idempotência pode envolver várias abordagens, incluindo:

- **Uso de Identificadores Únicos para Requisições:** Em sistemas que implementam **POST** de maneira idempotente,
  pode-se usar um identificador único (como um ID de transação) para garantir que a mesma operação não seja repetida.

- **Controle de Versão:** Manter o controle de versões de recursos e garantir que uma operação de **PUT** não cause
  alterações desnecessárias.

- **Verificação de Estado:** Verificar o estado de um recurso antes de realizar uma operação para garantir que a
  operação não será repetida de forma indesejada.

## **6. Conclusão**

A **idempotência** é uma propriedade crucial para garantir a confiabilidade e a previsibilidade em sistemas de software,
especialmente quando se trata de APIs, sistemas distribuídos e manipulação de dados em ambientes com alta concorrência
ou redes instáveis.

Ao garantir que uma operação possa ser repetida sem causar efeitos adversos, ela permite que o sistema se recupere de
falhas de forma robusta e que os usuários e sistemas não sejam impactados por repetições de requisições.

Implementar idempotência corretamente é essencial para a integridade e a escalabilidade de sistemas modernos.