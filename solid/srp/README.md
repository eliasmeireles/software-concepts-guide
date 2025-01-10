# Single Responsibility Principle (SRP)

## O que é o SRP?

O **Princípio da Responsabilidade Única** (SRP, do inglês *Single Responsibility Principle*) é o primeiro dos cinco
princípios SOLID, introduzidos por Robert C. Martin (também conhecido como Uncle Bob). Este princípio estabelece que:

> **Uma classe deve ter apenas uma razão para mudar.**

Em outras palavras, uma classe deve ser responsável por **apenas uma tarefa ou funcionalidade específica** dentro do
sistema.

---

## Por que o SRP é importante?

O SRP ajuda a evitar o chamado *"God Object"*, que é uma classe que tenta fazer "tudo". Esse tipo de design resulta em
código difícil de manter, testar e modificar.

Os benefícios do SRP incluem:

- **Facilidade de Manutenção:** Cada classe se concentra em uma única responsabilidade, tornando o código mais legível e
  fácil de modificar.
- **Reutilização de Código:** Classes bem definidas podem ser reutilizadas em outros contextos ou projetos.
- **Testabilidade:** Classes pequenas e com responsabilidade única são mais fáceis de testar.
- **Menor Acoplamento:** Promove uma separação clara de responsabilidades, reduzindo o impacto de mudanças em outras
  partes do sistema.

---

## Como identificar uma violação do SRP?

Você pode estar violando o SRP se:

1. Uma classe possui métodos que lidam com **diferentes aspectos de um sistema** (por exemplo, lógica de negócios,
   manipulação de dados e interface de usuário).
2. Alterar uma funcionalidade exige modificações em **várias partes da classe**, aumentando o risco de erros.
3. O nome da classe é genérico demais (*Manager*, *Processor*, *Handler*, etc.), indicando que ela está tentando
   realizar muitas coisas.

### Exemplo de uma violação do SRP:

Se uma classe é responsável por:

- Validar dados de entrada.
- Salvar os dados no banco de dados.
- Enviar notificações por e-mail.

Nesse caso, a classe possui três razões para mudar:

1. Mudança nas regras de validação.
2. Alterações no banco de dados.
3. Modificações no envio de notificações.

---

## Como aplicar o SRP?

1. **Divida responsabilidades:** Separe funcionalidades distintas em classes ou módulos diferentes.
    - Por exemplo, uma classe deve lidar apenas com **validação de dados**, enquanto outra deve cuidar da **persistência
      no banco de dados**.

2. **Use nomes claros e descritivos:** O nome de uma classe deve refletir sua responsabilidade única.
    - Exemplo: `ValidadorDeUsuario`, `RepositorioDeUsuarios`, `ServicoDeEmail`.

3. **Delegue tarefas:** Em vez de sobrecarregar uma classe, delegue tarefas específicas para outras classes que são
   especialistas naquela função.

---

## Quando o SRP pode ser negligenciado?

Em sistemas muito simples ou protótipos, seguir o SRP de forma rígida pode parecer um esforço desnecessário. No entanto,
em sistemas maiores ou em crescimento, ignorar o SRP pode levar a um **aumento significativo na complexidade e no custo
de manutenção**.

---

## Conclusão

O **Single Responsibility Principle** é um dos pilares do design orientado a objetos bem estruturado. Ele incentiva a
criação de classes com foco em uma única responsabilidade, facilitando a manutenção, teste e evolução do software.
Embora sua aplicação possa parecer exagerada em projetos pequenos, ele se torna essencial para a escalabilidade de
sistemas complexos.

Lembre-se: **"Uma classe, uma responsabilidade."** Essa simples regra pode transformar a qualidade do seu código e do
seu sistema como um todo.
