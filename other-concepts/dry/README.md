# Princípio DRY (Don't Repeat Yourself) no Desenvolvimento de Software

O princípio **DRY (Don't Repeat Yourself)** é um dos conceitos fundamentais no desenvolvimento de software. Ele tem como
objetivo minimizar a repetição de código e lógica em um sistema, tornando-o mais eficiente, sustentável e fácil de
manter. Este princípio foi formalizado no livro *The Pragmatic Programmer*, de Andy Hunt e Dave Thomas.

## O Que É o Princípio DRY?

O DRY prega que "cada pedaço de conhecimento deve ter uma única representação não ambígua e definitiva dentro do
sistema". Em outras palavras, devemos evitar duplicar código. Em vez disso, criar abstrações e reutilizar módulos
específicos sempre que eles desempenharem funções similares ou compartilharem lógica semelhante.

## Benefícios do DRY

Ao implementar o DRY corretamente, você pode obter diversos benefícios, como:

- **Manutenção mais fácil**: Alterar a lógica em um único local reflete em todo o sistema.
- **Menos erros**: Reduz duplicações e inconsistências, minimizando riscos de bugs.
- **Legibilidade**: Código mais limpo e fácil de entender.
- **Eficiência de tempo**: Desenvolvedores podem reutilizar o mesmo código ou lógica de maneira consistente,
  economizando esforço.

---

## Exemplo em Go

Aqui está um exemplo ilustrando o DRY:

### Código sem DRY (repetição desnecessária)

```go
package main

import "fmt"

func calcularDescontoPrecoA(preco float64) float64 {
	return preco * 0.9
}

func calcularDescontoPrecoB(preco float64) float64 {
	return preco * 0.9
}

func main() {
	precoA := 100.0
	precoB := 200.0

	fmt.Printf("Preço A com desconto: %.2f\n", calcularDescontoPrecoA(precoA))
	fmt.Printf("Preço B com desconto: %.2f\n", calcularDescontoPrecoB(precoB))
}
```

### Código refatorado usando o DRY

No exemplo abaixo, removemos a duplicação ao criar uma função genérica reutilizável:

```go
package main

import "fmt"

// Função genérica para cálculo de desconto
func calcularDesconto(preco float64) float64 {
	return preco * 0.9
}

func main() {
	precoA := 100.0
	precoB := 200.0

	fmt.Printf("Preço A com desconto: %.2f\n", calcularDesconto(precoA))
	fmt.Printf("Preço B com desconto: %.2f\n", calcularDesconto(precoB))
}
```

Neste exemplo, consolidamos a lógica do desconto em uma única função (`calcularDesconto`), reduzindo duplicações. Com
isso, se a lógica do desconto precisar mudar, basta alterar apenas um único local.

---

## Prós e Contras do DRY

### Vantagens (Prós)

1. **Reutilização de Código**: Promove consistência ao reutilizar os mesmos módulos ou funções.
2. **Facilidade de Manutenção**: Modificações impactam menos partes do sistema, reduzindo o custo de manutenção.
3. **Redução de Bugs**: Menos duplicação significa menos chances de erros em diferentes partes do código.
4. **Desenvolvimento Mais Rápido**: Reaproveitar código existente economiza tempo de desenvolvimento.

### Desvantagens (Contras)

1. **Overengineering (Excesso de Abstração)**: Tentar remover duplicações em casos excessivamente complexos pode
   resultar em soluções difíceis de entender.
2. **Impacto no Desempenho**: Abstrações excessivas podem adicionar sobrecarga de performance em alguns casos.
3. **Curva de Aprendizado**: Equipes podem ter dificuldades de entender partes do código muito abstraídas ou genéricas.

---

## Conclusão

O princípio DRY é uma diretriz altamente valiosa para o desenvolvimento sustentável de software, e sua aplicação correta
pode resultar em projetos de alta qualidade, mais fáceis de manter e menos propensos a erros. No entanto, é importante
usá-lo com bom senso, evitando cenários de abstração excessiva ou desnecessária. Em última análise, como em qualquer
princípio de engenharia de software, o DRY deve sempre servir ao objetivo principal: criar código simples, legível e
eficiente.

---

## Referências

- [The Pragmatic Programmer - Livro](https://pragprog.com/titles/tpp20/the-pragmatic-programmer-20th-anniversary-edition/)
- [Documentação da linguagem Go](https://go.dev/doc/)
- [Wiki sobre DRY](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself)
