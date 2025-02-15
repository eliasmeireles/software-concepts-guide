# Design Pattern: Observer (Padrão de Design Observador)

## **Documentação detalhada**

### **Definição**

O padrão de design Observer (ou Observador) é um dos padrões comportamentais do livro *Design Patterns: Elements of
Reusable Object-Oriented Software*. Ele é utilizado quando existe a necessidade de notificar múltiplos objetos de alguma
mudança no estado de outro objeto. A ideia principal é desacoplar o objeto principal (observado) dos objetos
dependentes (observadores), permitindo uma comunicação mais flexível e passível de extensão.

---

### **Estrutura**

1. **Sujeito (Subject):**
    - O objeto que contém o estado compartilhado.
    - Mantém uma lista de observadores.
    - Oferece métodos para registrar/remover notificações e para notificar os observadores sobre mudanças no estado.

2. **Observador (Observer):**
    - Define uma interface para ser implementada por todos os observadores.
    - Recebe atualizações do estado do sujeito.

3. **Implementação concreta de ambos:**
    - Sujeito concreto: Implementa as operações de registro e notificação.
    - Observador concreto: Define ações específicas quando recebe uma notificação de mudança.

---

### **Diagrama UML**

```plaintext
+-------------------+          +-------------------+
|     Subject       |          |     Observer      |
|-------------------|          |-------------------|
| - observers       |<>------->| + Update()        |
| + Attach(o)       |          +-------------------+
| + Detach(o)       |
| + Notify()        |
+-------------------+ 
        ▲
        |
 +--------------------+
 | ConcreteSubject    |
 +--------------------+
 | - state            |
 | + GetState()       |
 | + SetState(state)  |
 +--------------------+ 
        |
 +--------------------+
 | ConcreteObserver   |
 +--------------------+
 | - subject          |
 | - observerState    |
 | + Update()         |
 +--------------------+
```

---

### **Exemplo de código em Go Lang**

Neste exemplo, implementaremos o padrão Observer com um `WeatherStation` (estação meteorológica) como o _Sujeito_, e
dispositivos como telefones ou painéis eletrônicos como _Observadores_.

```go
package main

import "fmt"

// --- Interface Observer ---
type Observer interface {
	Update(temperature float64) // Método chamado quando o estado muda
}

// --- Interface Subject ---
type Subject interface {
	Attach(observer Observer) // Adiciona um observador
	Detach(observer Observer) // Remove um observador
	Notify()                  // Notifica todos os observadores
}

// --- Sujeito Concreto ---
type WeatherStation struct {
	temperature float64
	observers   []Observer
}

func NewWeatherStation() *WeatherStation {
	return &WeatherStation{}
}

// Adicionar novo observador
func (w *WeatherStation) Attach(observer Observer) {
	w.observers = append(w.observers, observer)
}

// Remover um observador
func (w *WeatherStation) Detach(observer Observer) {
	var index int
	for i, o := range w.observers {
		if o == observer {
			index = i
			break
		}
	}
	w.observers = append(w.observers[:index], w.observers[index+1:]...)
}

// Notificar todos os observadores
func (w *WeatherStation) Notify() {
	for _, observer := range w.observers {
		observer.Update(w.temperature)
	}
}

// Atualizar o estado do sujeito
func (w *WeatherStation) SetTemperature(temp float64) {
	w.temperature = temp
	w.Notify() // Sempre que o estado muda, notifica os observadores
}

// --- Observador Concreto ---
type PhoneDisplay struct {
	subject *WeatherStation
}

func (p *PhoneDisplay) Update(temp float64) {
	fmt.Printf("PhoneDisplay atualizado com a temperatura: %.2f°C\n", temp)
}

type WindowDisplay struct {
	subject *WeatherStation
}

func (w *WindowDisplay) Update(temp float64) {
	fmt.Printf("WindowDisplay atualizado com a temperatura: %.2f°C\n", temp)
}

// --- Função principal ---
func main() {
	// Instancia o Subject (Sujeito)
	weatherStation := NewWeatherStation()

	// Cria Observadores
	phoneDisplay := &PhoneDisplay{}
	windowDisplay := &WindowDisplay{}

	// Registra os observadores
	weatherStation.Attach(phoneDisplay)
	weatherStation.Attach(windowDisplay)

	// Atualiza a temperatura e notifica os observadores
	weatherStation.SetTemperature(25.0)
	weatherStation.SetTemperature(30.5)
}
```

**Saída esperada do programa:**

```plaintext
PhoneDisplay atualizado com a temperatura: 25.00°C
WindowDisplay atualizado com a temperatura: 25.00°C
PhoneDisplay atualizado com a temperatura: 30.50°C
WindowDisplay atualizado com a temperatura: 30.50°C
```

---

### **Vantagens**

1. **Desacoplamento:** O padrão reduz o acoplamento entre o objeto observado e os observadores.
2. **Escalabilidade:** É fácil adicionar novos observadores sem alterar o código do sujeito.
3. **Extensibilidade:** Permite diferentes tipos de observadores com comportamentos personalizados.

---

### **Desvantagens**

1. **Complexidade:** A implementação básica do padrão pode adicionar complexidade, especialmente ao gerenciar a lista de
   observadores.
2. **Dependência indireta:** Problemas podem surgir quando muitos observadores estão conectados ao sujeito, criando
   dependências ocultas.
3. **Potencial para loops:** Se implementado de forma errada, um loop infinito de notificações pode ocorrer.

---

### **Conclusão**

O padrão Observer é amplamente utilizado quando se deseja manter diferentes partes do sistema sincronizadas. Ele é
essencial para cenários em que mudanças em um objeto precisam ser refletidas em muitos outros, mas sem criar
dependências rígidas. No entanto, como qualquer padrão, ele deve ser usado com cuidado para evitar aumento desnecessário
de complexidade.

---

### **Referências**

1. Livro *Design Patterns: Elements of Reusable Object-Oriented Software* (GoF).
2. Documentação oficial de Go: https://golang.org/doc/
3. Artigos e tutoriais sobre padrões de design: https://refactoring.guru/design-patterns/observer
