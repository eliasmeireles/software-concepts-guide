# REST e RESTful: Uma Visão Detalhada

## O que é REST?

REST (“Representational State Transfer”) é um estilo arquitetônico para criação de sistemas distribuídos. Ele foi
introduzido por Roy Fielding em sua tese de doutorado em 2000. REST não é um protocolo ou padronização formal, mas sim
um conjunto de princípios que guiam o design de APIs para torná-las simples, escaláveis e eficientes.

### Princípios Fundamentais do REST

1. **Interface Uniforme:** É a base do REST, garantindo que as interações entre cliente e servidor sejam padronizadas.
   Isso envolve:
   - Identificação de recursos por URLs.
   - Manipulação de recursos através de representações (JSON, XML, etc.).
   - Uso de verbos HTTP (“GET”, “POST”, “PUT”, “DELETE”).

2. **Arquitetura Cliente-Servidor:**
   - O cliente e o servidor são separados e interagem através de uma interface bem definida.
   - Essa separação permite a evolução independente de ambos os lados.

3. **Sem Estado (Stateless):**
   - Cada requisição do cliente para o servidor deve conter todas as informações necessárias para o servidor
     processá-la.
   - Não é mantido nenhum estado de sessão no servidor entre requisições.

4. **Cacheabilidade:**
   - Respostas de recursos devem indicar se podem ser armazenadas em cache.
   - Isso melhora a eficiência ao evitar requisições desnecessárias ao servidor.

5. **Sistema em Camadas:**
   - A arquitetura pode ser composta por camadas (e.g., servidores intermediários) que não interferem umas nas outras.

6. **Código Sob Demanda (Opcional):**
   - O servidor pode enviar código executável (e.g., JavaScript) ao cliente para ampliar sua funcionalidade.

### Benefícios do REST

- Escalabilidade: devido à separação entre cliente e servidor.
- Flexibilidade: suporte a diferentes formatos de dados.
- Simplicidade: baseia-se nos padronizados verbos HTTP e estrutura de URLs.
- Interoperabilidade: APIs REST podem ser consumidas por diversos tipos de clientes.

---

## O que é RESTful?

RESTful refere-se à implementação prática dos princípios REST em sistemas e APIs. Uma API é considerada RESTful se segue
fielmente os princípios REST descritos acima.

### Características de uma API RESTful

1. **URLs como Identificadores de Recursos:**
   - Cada recurso é representado por uma URL clara e simples.
   - Exemplo: `https://api.exemplo.com/usuarios/123` representa o usuário com ID 123.

2. **Uso Correto dos Verbos HTTP:**
   - `GET`: Para recuperar informações.
   - `POST`: Para criar novos recursos.
   - `PUT`: Para atualizar recursos existentes.
   - `DELETE`: Para remover recursos.

3. **Representações Padronizadas:**
   - Os dados são transmitidos em formatos como JSON, XML ou outros.
   - Exemplo de resposta JSON:
     ```json
     {
         "id": 123,
         "nome": "João Silva",
         "email": "joao.silva@exemplo.com"
     }
     ```

4. **Cabeçalhos HTTP Bem Definidos:**
   - Uso de cabeçalhos para metadados (e.g., `Content-Type`, `Authorization`, `Cache-Control`).

5. **Links HATEOAS (Hypermedia as the Engine of Application State):**
   - Recursos incluem links para outras ações ou informações relacionadas.
   - Exemplo:
     ```json
     {
         "id": 123,
         "nome": "João Silva",
         "links": [
             {"rel": "self", "href": "https://api.exemplo.com/usuarios/123"},
             {"rel": "pedidos", "href": "https://api.exemplo.com/usuarios/123/pedidos"}
         ]
     }
     ```

6. **Mensagens Sem Estado:**
   - Cada requisição é independente e transporta todas as informações necessárias para processá-la.

7. **Documentação Clara:**
   - APIs RESTful geralmente incluem documentação detalhada para facilitar o uso, como Swagger ou Postman Collections.

### Exemplo Prático de uma API RESTful

#### Recursos:

- `https://api.exemplo.com/usuarios`
- `https://api.exemplo.com/usuarios/{id}`
- `https://api.exemplo.com/usuarios/{id}/pedidos`

#### Operações:

1. **GET /usuarios**
   - Retorna a lista de usuários.
2. **POST /usuarios**
   - Cria um novo usuário.
   - Corpo da requisição:
     ```json
     {
         "nome": "Maria Oliveira",
         "email": "maria.oliveira@exemplo.com"
     }
     ```
3. **GET /usuarios/{id}**
   - Retorna os dados de um usuário específico.
4. **PUT /usuarios/{id}**
   - Atualiza os dados de um usuário.
5. **DELETE /usuarios/{id}**
   - Remove um usuário.

---

## Diferenças entre REST e RESTful

- **REST:** Conjunto de princípios arquiteturais.
- **RESTful:** Implementação prática desses princípios em um sistema ou API.

---

## Conclusão

O REST é um estilo arquitetônico amplamente adotado por sua simplicidade e eficácia, enquanto o termo RESTful se refere
à aplicação correta dos princípios REST. Implementar uma API RESTful ajuda a criar sistemas padronizados, fáceis de
entender e interoperáveis, promovendo escalabilidade e manutenção eficiente.

