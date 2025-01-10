## **TDD (Test-Driven Development)**.

---

### O que é TDD?

**TDD (Test-Driven Development)** é uma abordagem de desenvolvimento onde o desenvolvedor escreve primeiro os testes que
descrevem o comportamento desejado, e depois escreve o código para fazer esses testes passarem. Esse processo ajuda a
melhorar a qualidade do código, promovendo testes automáticos que garantem que o sistema funcione corretamente à medida
que evolui.

### Ciclo Red-Green-Refactor do TDD

O ciclo de TDD segue três etapas principais:

1. **Red (Escrever o Teste)**

- Escrever um teste que falha, pois o código ainda não foi implementado.

2. **Green (Escrever o Código para Passar no Teste)**

- Escrever o código necessário para passar no teste.

3. **Refactor (Refatorar o Código)**

- Refatorar o código, mantendo todos os testes passados, para melhorar sua qualidade e clareza.

---

### Benefícios do TDD

1. **Qualidade do Código**

- O código é testado desde o início, o que aumenta a qualidade e reduz bugs.

2. **Facilidade de Refatoração**

- Como você tem uma suíte de testes, pode refatorar o código com segurança, sabendo que qualquer quebra será detectada
  rapidamente.

3. **Documentação Automática**

- Os testes servem como documentação de como o código deve se comportar.

4. **Design Melhorado**

- O TDD força você a pensar sobre como o código deve ser estruturado antes de escrevê-lo.

---

### Exemplo Prático de TDD em Go

Vamos criar um simples exemplo de uma função que soma dois números e escrevê-la com TDD.

#### 1. **Escrever o Teste (Red)**

Primeiro, criamos o teste para a função de soma. O teste ainda falhará, pois a função ainda não foi implementada.

```go
// sum_test.go
package main

import "testing"

func TestSoma(t *testing.T) {
	resultado := soma(2, 3)
	if resultado != 5 {
		t.Errorf("Soma(2, 3) = %d; esperado 5", resultado)
	}
}
```

Aqui, escrevemos o teste da função `soma` que, neste caso, deve somar 2 e 3. Como a função `soma` ainda não existe, o
teste falhará.

#### 2. **Escrever o Código para Passar no Teste (Green)**

Agora, implementamos a função `soma` de forma simples para fazer o teste passar.

```go
// sum.go
package main

func soma(a, b int) int {
	return a + b
}
```

Agora, se executarmos os testes, o teste que escrevemos passará, porque a função `soma` retorna a soma correta de dois
números.

#### 3. **Refatorar o Código (Refactor)**

No nosso exemplo, a função já está bem simples, então não há necessidade de refatoração. No entanto, se a implementação
fosse mais complexa, você refatoraria o código para melhorar sua legibilidade ou desempenho, sempre garantindo que os
testes continuem passando.

---

### Execução dos Testes

Para rodar os testes em Go, você usa o comando:

```sh
go test
```

Se tudo estiver correto, o Go mostrará uma saída informando que todos os testes passaram.

---

### Benefícios do TDD em Go

1. **Cobertura de Testes**

- O Go possui um excelente suporte nativo para testes através do pacote `testing`, que facilita a criação de testes
  automatizados.

2. **Facilidade de Manutenção**

- O TDD ajuda a manter o código fácil de manter e refatorar, pois sempre que uma mudança for feita, você poderá executar
  os testes para garantir que o comportamento esperado foi mantido.

3. **Aumento da Confiabilidade**

- O TDD aumenta a confiança de que o sistema está funcionando como esperado, pois a cada nova funcionalidade
  implementada, ela é coberta por testes.

---

### Conclusão

O **TDD** em Go segue a mesma filosofia de outras linguagens, com a vantagem de contar com um excelente suporte de
testes do próprio Go. Ao escrever testes primeiro e depois o código para fazer esses testes passarem, você pode garantir
um código mais robusto, fácil de manter e refatorar. O ciclo **Red-Green-Refactor** é simples, mas poderoso para
garantir a qualidade do software durante o desenvolvimento.

Ao aplicar TDD, você não só melhora a qualidade do seu código, mas também cria uma base sólida de testes que servem como
documentação viva e segurança contra regressões.