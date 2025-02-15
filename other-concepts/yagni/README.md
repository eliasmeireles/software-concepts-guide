# YAGNI (You Aren’t Gonna Need It) - Princípio no Desenvolvimento de Software

## O que é o YAGNI?

YAGNI, ou "You Aren't Gonna Need It" (Você Não Vai Precisar Disso), é um princípio fundamental no desenvolvimento de
software que defende que os desenvolvedores não devem implementar funcionalidades que não são necessárias no momento —
ou seja, funcionalidades que não atendem aos requisitos presentes.

Este conceito está fortemente relacionado ao desenvolvimento ágil, à simplicidade e à redução de código desnecessário,
promovendo um design mais eficiente, menos propenso a bugs e mais fácil de manter.

Segundo YAGNI, ao implementar apenas o que é necessário no momento, evitamos desperdícios de tempo e esforço na criação
de soluções que podem nunca ser usadas.

---

## Benefícios do YAGNI

**Prós**:

1. **Redução de Complexidade**: Ajuda a evitar código desnecessário que pode confundir outros desenvolvedores no futuro.
2. **Custo de Manutenção Reduzido**: Como o software é mais simples, é mais fácil de manter e refatorar.
3. **Foco nos Requisitos Atuais**: Garante que o time se concentre apenas no que o usuário ou cliente realmente precisa
   no momento.
4. **Velocidade de Entrega Inicial Mais Alta**: Evita desperdício de tempo na resolução de problemas em partes do código
   que nem sequer eram necessárias.
5. **Facilita Refatoração**: É mais fácil ajustar o software para novos requisitos, pois não há excesso de código
   desnecessário.

---

## Considerações e Limitações

**Contras**:

1. **Riscos de Mudanças Futuras**: Se novos requisitos relacionados ao que foi ignorado surgirem, pode ser necessário
   reimplementar ou até refatorar partes do sistema.
2. **Perda de Contexto**: Em alguns casos, pode ser necessário relembrar e reanalisar decisões que poderiam já ter sido
   pensadas anteriormente.
3. **Possível Falta de Visão Estratégica**: O foco exclusivo no presente pode ignorar a possibilidade de crescimento
   futuro do sistema.

---

## Quando Usar o YAGNI?

1. Utilize sempre que for possível evitar implementar funcionalidades que "talvez sejam usadas no futuro".
2. Certifique-se de que você está focando nos requisitos claramente definidos pelos clientes ou produtos.
3. Combine YAGNI com práticas ágeis, como **refatoração contínua** e **testes automatizados**, para garantir
   flexibilidade no futuro.

---

## Exemplo em Go

Aqui está um exemplo de como o YAGNI pode ser aplicado em Go. Imagine que estamos desenvolvendo uma aplicação para
processar números. No entanto, **não há requisitos para armazenar os números em um banco de dados neste momento**.

**Código sem aplicar YAGNI (com código desnecessário para gravação em banco de dados):**

```go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func processNumber(num int) error {
	// Processamento simples
	fmt.Printf("Número processado: %d\n", num)

	// Código desnecessário para armazenar no banco de dados,
	// pois atualmente não há essa exigência
	db, err := sql.Open("postgres", "user=placeholder dbname=placeholder sslmode=disable")
	if err != nil {
		return fmt.Errorf("erro ao conectar no banco: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO numbers (value) VALUES ($1)", num)
	if err != nil {
		return fmt.Errorf("erro ao salvar o número no banco: %v", err)
	}

	return nil
}

func main() {
	err := processNumber(42)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
	}
}
```

**Código simplificado aplicando YAGNI:**

```go
package main

import (
	"fmt"
)

func processNumber(num int) {
	// Apenas processa o número sem adicionar complexidade desnecessária
	fmt.Printf("Número processado: %d\n", num)
}

func main() {
	processNumber(42)
}
```

No segundo exemplo, não há tentativa de conectar ao banco, pois **não há requisitos atuais para armazenamento dos
números processados**. Essa decisão utiliza o princípio YAGNI para economizar esforço e evitar complexidade
desnecessária até que um requisito real exija essa funcionalidade.

---

## Conclusão

O YAGNI é uma prática simples, mas poderosa, que ajuda os desenvolvedores a economizarem tempo e criar software de
qualidade com menos esforço. Ao evitar a implementação de funcionalidades desnecessárias, as equipes permanecem focadas
no que realmente importa e entregam valor rapidamente. No entanto, é fundamental equilibrar o uso do YAGNI com a visão
de longo prazo, considerando o impacto de possíveis mudanças nos requisitos futuros.

---

## Recursos Adicionais

- [Princípios Ágeis: YAGNI - Martin Fowler](https://martinfowler.com/bliki/Yagni.html)
- [Extreme Programming - YAGNI Definition](http://c2.com/cgi/wiki?YouArentGonnaNeedIt)
- [12 Princípios Ágeis](https://agilemanifesto.org/principles.html)
