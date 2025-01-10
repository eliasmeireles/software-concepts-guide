# Interface Segregation Principle (ISP)

## O que é o ISP?

O **Princípio da Segregação de Interfaces** (ISP, do inglês *Interface Segregation Principle*) é o quarto dos princípios
SOLID, também proposto por Robert C. Martin. Este princípio estabelece que:

> **Clientes não devem ser forçados a depender de interfaces que não utilizam.**

Isso significa que é preferível ter várias interfaces específicas e coesas ao invés de uma única interface geral e
extensa.

---

## Por que o ISP é importante?

O ISP promove a **cohesão** e **flexibilidade** das interfaces, melhorando a **manutenibilidade** e **reutilização** do
código. Os benefícios incluem:

- **Redução de Acoplamento:** Interfaces menores e específicas diminuem a dependência entre componentes.
- **Facilidade de Implementação:** Classes implementam apenas as funcionalidades que realmente necessitam.
- **Melhoria na Legibilidade:** Interfaces coesas são mais fáceis de entender e utilizar corretamente.
- **Flexibilidade:** Facilita a modificação e extensão das funcionalidades sem impacto em outras partes do sistema.

---

## Como identificar uma violação do ISP?

Você pode estar violando o ISP se:

1. **Interfaces Grandes e Genéricas:** Interfaces que contêm métodos que não são utilizados por todas as classes que as
   implementam.
2. **Implementações Vazia:** Classes que implementam interfaces e deixam métodos não utilizados sem implementação ou
   lançam exceções.
3. **Forçar Dependências:** Classes que são obrigadas a depender de métodos que não fazem sentido para elas.

### Exemplo de uma violação do ISP:

Se uma interface `IMachine` possui métodos `Print()`, `Scan()` e `Fax()`, mas uma classe `Printer` implementa `IMachine`
e não utiliza os métodos `Scan()` e `Fax()`, isso indica uma violação do ISP.

---

## Como aplicar o ISP?

1. **Crie Interfaces Coesas:** Defina interfaces que agrupem métodos relacionados e específicos a uma funcionalidade
   única.

2. **Favor Interfaces Pequenas:** Prefira ter várias interfaces pequenas e específicas ao invés de uma única interface
   grande.

3. **Segregue Funcionalidades:** Separe funcionalidades distintas em interfaces diferentes, permitindo que as classes
   implementem apenas o que realmente necessitam.

4. **Use Interfaces de Clientes Específicos:** Crie interfaces baseadas nas necessidades dos clientes que as utilizam,
   garantindo que não haja métodos desnecessários.

5. **Aplicar o Princípio da Responsabilidade Única (SRP):** Ao garantir que cada interface tem uma única
   responsabilidade, facilita a aplicação do ISP.

---

## Quando o ISP pode ser negligenciado?

Em **sistemas muito simples** ou **projetos de curto prazo**, a segregação excessiva de interfaces pode introduzir
complexidade desnecessária. No entanto, em sistemas **complexos** ou **em crescimento**, ignorar o ISP pode resultar em
**interfaces inchadas** e **código difícil de manter**.

---

## Conclusão

O **Interface Segregation Principle** é fundamental para a criação de sistemas modulares e de fácil manutenção. Ao
evitar interfaces genéricas e promover a criação de interfaces específicas e coesas, o ISP contribui para a **redução de
acoplamento**, **melhoria da flexibilidade** e **aumento da reutilização** do código. Aplicar este princípio é essencial
para desenvolver software robusto e escalável.
