# Liskov Substitution Principle (LSP)

## O que é o LSP?

O **Princípio da Substituição de Liskov** (LSP, do inglês *Liskov Substitution Principle*) é o terceiro dos princípios
SOLID, proposto por Barbara Liskov. Este princípio afirma que:

> **Objetos de uma classe derivada devem ser substituíveis por objetos da classe base sem alterar as propriedades
desejáveis do programa (correção, tarefa realizada, etc.).**

Em outras palavras, subclasses devem poder substituir suas superclasses sem que o comportamento do sistema seja
comprometido.

---

## Por que o LSP é importante?

O LSP assegura que as **heranças** sejam utilizadas de forma correta, promovendo **compatibilidade** e **reutilização**
de código. Os benefícios incluem:

- **Consistência:** Garante que as subclasses mantenham o comportamento esperado das classes base.
- **Flexibilidade:** Facilita a substituição de implementações sem impacto no restante do sistema.
- **Manutenibilidade:** Simplifica a manutenção e a extensão do código, evitando surpresas e comportamentos inesperados.
- **Reutilização Segura:** Permite que componentes sejam reutilizados em diferentes contextos sem riscos de falhas.

---

## Como identificar uma violação do LSP?

Você pode estar violando o LSP se:

1. **Comportamento Alterado:** Subclasses que alteram o comportamento esperado da classe base.
2. **Métodos que Lançam Exceções:** Subclasses que lançam exceções adicionais que não são previstas pela classe base.
3. **Pré-condições Mais Restritivas ou Pós-condições Mais Flexíveis:** Subclasses que impõem restrições adicionais ou
   relaxam garantias da classe base.

### Exemplo de uma violação do LSP:

Se uma classe base `Avião` possui um método `voar()`, e uma subclasse `Pato` que herda de `Avião` não consegue voar,
isso viola o LSP, pois `Pato` não pode substituir `Avião` sem alterar o comportamento esperado.

---

## Como aplicar o LSP?

1. **Assure Compatibilidade de Tipos:** As subclasses devem ser substituíveis pelas superclasses sem necessidade de
   alterações no código cliente.

2. **Mantenha Contratos Consistentes:** As pré-condições e pós-condições dos métodos na subclasse devem ser compatíveis
   com as da classe base.

3. **Evite Restrições Adicionais:** Não imponha restrições adicionais nas subclasses que não existem na classe base.

4. **Use Interfaces Apropriadas:** Defina interfaces que capturam corretamente o comportamento esperado, permitindo
   implementações que respeitem o LSP.

5. **Favor Composição sobre Herança:** Quando a herança pode levar a violações do LSP, prefira utilizar composição para
   reutilizar comportamentos.

---

## Quando o LSP pode ser negligenciado?

Em **situações de prototipagem rápida** ou **sistemas extremamente simples**, a aplicação rigorosa do LSP pode ser
dispensável. No entanto, em sistemas **complexos** ou **em evolução**, ignorar o LSP pode levar a **problemas de
integridade** e **manutenção difícil**.

---

## Conclusão

O **Liskov Substitution Principle** é essencial para garantir que a herança seja utilizada de maneira correta e eficaz.
Ao assegurar que subclasses possam substituir suas superclasses sem alterar o comportamento esperado do sistema, o LSP
contribui para a **robustez**, **flexibilidade** e **manutenibilidade** do código. Adotar este princípio é crucial para
o desenvolvimento de software orientado a objetos de alta qualidade.
