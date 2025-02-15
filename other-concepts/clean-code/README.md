## **Clean Code (Código Limpo)**

**Clean Code** é uma filosofia de desenvolvimento de software que enfatiza a escrita de código legível, simples e de
fácil manutenção. O conceito foi popularizado por Robert C. Martin (também conhecido como Uncle Bob) em seu livro *"
Clean Code: A Handbook of Agile Software Craftsmanship"*. A ideia central é que o código deve ser escrito para ser
entendido por humanos, não apenas para ser executado por máquinas.

### **Princípios do Clean Code**

1. **Legibilidade**: O código deve ser fácil de ler e entender.
2. **Simplicidade**: Evite complexidade desnecessária.
3. **Manutenibilidade**: O código deve ser fácil de modificar e estender.
4. **Testabilidade**: O código deve ser fácil de testar.
5. **Consistência**: Siga padrões e convenções consistentes.
6. **Responsabilidade Única**: Cada função, classe ou módulo deve ter uma única responsabilidade.

### **Quando aplicar Clean Code?**

- Em qualquer projeto de software, independentemente do tamanho ou complexidade.
- Quando você deseja melhorar a qualidade do código e facilitar a colaboração em equipe.
- Quando você quer reduzir o custo de manutenção a longo prazo.

---

### **Exemplo de Clean Code em Go**

Vamos comparar um código "sujo" com um código "limpo" em Go. O exemplo é uma função que calcula a média de uma lista de
números.

#### **Código "Sujo"**

```go
package main

import "fmt"

func calcAvg(nums []float64) float64 {
	var sum float64
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum / float64(len(nums))
}

func main() {
	numbers := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	avg := calcAvg(numbers)
	fmt.Println("Média:", avg)
}
```

#### **Problemas do Código "Sujo"**

1. Nome da função pouco descritivo (`calcAvg`).
2. Uso de `i` como nome de variável no loop, sem significado claro.
3. Falta de tratamento para casos em que a lista está vazia (divisão por zero).

---

#### **Código "Limpo"**

```go
package main

import (
	"errors"
	"fmt"
)

// Função com nome descritivo e tratamento de erros
func CalculateAverage(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("a lista de números está vazia")
	}

	sum := 0.0
	for _, number := range numbers {
		sum += number
	}

	average := sum / float64(len(numbers))
	return average, nil
}

func main() {
	numbers := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

	average, err := CalculateAverage(numbers)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Printf("A média dos números é: %.2f\n", average)
}
```

#### **Melhorias no Código "Limpo"**

1. Nome da função descritivo (`CalculateAverage`).
2. Uso de `number` no loop, que é mais significativo.
3. Tratamento de erro para lista vazia.
4. Formatação de saída para melhor legibilidade.

---

### **Práticas Recomendadas para Clean Code em Go**

1. **Nomes Significativos**: Use nomes descritivos para variáveis, funções e pacotes.
    - Exemplo: `CalculateAverage` em vez de `calcAvg`.
2. **Funções Pequenas**: Mantenha as funções curtas e com uma única responsabilidade.
3. **Evite Código Duplicado**: Reutilize código sempre que possível.
4. **Comentários Úteis**: Use comentários para explicar o "porquê", não o "como".
5. **Tratamento de Erros**: Sempre trate erros de forma explícita.
6. **Testes Automatizados**: Escreva testes unitários para garantir que o código funcione conforme o esperado.

---

### **Prós do Clean Code**

1. **Legibilidade**: Código fácil de ler e entender.
2. **Manutenibilidade**: Facilita a modificação e extensão do código.
3. **Colaboração**: Melhora a comunicação entre membros da equipe.
4. **Redução de Bugs**: Código mais claro e simples tende a ter menos erros.
5. **Testabilidade**: Código bem estruturado é mais fácil de testar.

### **Contras do Clean Code**

1. **Tempo Inicial**: Pode levar mais tempo para escrever código limpo.
2. **Curva de Aprendizado**: Requer prática e disciplina para aplicar os princípios corretamente.
3. **Overhead**: Em projetos pequenos, pode parecer excessivo.

---

### **Conclusão**

O Clean Code é uma prática essencial para qualquer desenvolvedor que deseja escrever software de alta qualidade. Ele
promove a legibilidade, manutenibilidade e escalabilidade do código, reduzindo custos a longo prazo. Embora possa exigir
um esforço inicial maior, os benefícios superam os custos, especialmente em projetos grandes e complexos.

### **Links de Referência**

- [Clean Code: A Handbook of Agile Software Craftsmanship](https://www.amazon.com/Clean-Code-Handbook-Software-Craftsmanship/dp/0132350882) (
  Livro do Uncle Bob)

---

