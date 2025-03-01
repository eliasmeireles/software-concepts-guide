# Introdução à Criptografia

## O que é Criptografia?

Criptografia é a prática de proteger informações através do uso de algoritmos matemáticos, transformando dados
legíveis (texto plano) em uma forma ilegível (texto cifrado). O objetivo principal é garantir a **confidencialidade**,
**integridade**, e **autenticidade** das informações.

### Definições Básicas

- **Texto Plano (Plaintext):** Dados em sua forma original, legível.
- **Texto Cifrado (Ciphertext):** Dados transformados em uma forma ilegível.
- **Chave (Key):** Um valor secreto usado para cifrar e decifrar dados.
- **Algoritmo de Criptografia:** Um conjunto de regras matemáticas usadas para cifrar e decifrar dados.

---

## Para que serve a Criptografia?

A criptografia é usada em diversas áreas para proteger informações sensíveis. Aqui estão alguns dos principais usos:

1. **Confidencialidade:** Impedir que pessoas não autorizadas acessem informações.
2. **Integridade:** Garantir que os dados não foram alterados durante a transmissão ou armazenamento.
3. **Autenticidade:** Verificar a identidade do remetente ou destinatário.
4. **Não Repúdio:** Impedir que uma parte negue a autoria de uma ação ou mensagem.

---

## Tipos de Criptografia

### 1. Criptografia Simétrica

- Usa a **mesma chave** para cifrar e decifrar dados.
- **Vantagens:** Rápida e eficiente para grandes volumes de dados.
- **Desvantagens:** Distribuição segura da chave é um desafio.
- **Exemplos:** AES (Advanced Encryption Standard), DES (Data Encryption Standard).

### 2. Criptografia Assimétrica

- Usa um **par de chaves**: uma pública (para cifrar) e uma privada (para decifrar).
- **Vantagens:** Resolve o problema de distribuição de chaves.
- **Desvantagens:** Mais lenta e computacionalmente intensiva.
- **Exemplos:** RSA, ECC (Elliptic Curve Cryptography).

### 3. Funções de Hash

- Transforma dados em um valor fixo (hash) que não pode ser revertido.
- Usado para verificar integridade de dados.
- **Exemplos:** SHA-256, MD5 (não recomendado para segurança).

---

## O que é Hash?

Hash é uma função matemática que transforma um conjunto de dados (de qualquer tamanho) em um valor fixo de comprimento,
conhecido como **valor de hash** ou **digest**. Diferente da criptografia, o hash é uma operação **unidirecional**, ou
seja, não é possível reverter o valor de hash para obter os dados originais.

### Características do Hash:

1. **Determinístico:** O mesmo conjunto de dados sempre produzirá o mesmo valor de hash.
2. **Rápido:** O cálculo do hash deve ser eficiente.
3. **Resistente a Colisões:** É difícil encontrar dois conjuntos de dados diferentes que produzam o mesmo valor de hash.
4. **Irreversível:** Não é possível reconstruir os dados originais a partir do valor de hash.

### Usos Comuns de Hash:

- Verificação de integridade de arquivos (ex.: checksum).
- Armazenamento seguro de senhas (ex.: em bancos de dados).
- Assinaturas digitais e certificados.

---

## Diferença entre Hash e Criptografia

| **Característica**      | **Hash**                                | **Criptografia**                   |
|-------------------------|-----------------------------------------|------------------------------------|
| **Reversibilidade**     | Unidirecional (não pode ser revertido). | Bidirecional (pode ser decifrado). |
| **Uso de Chaves**       | Não usa chaves.                         | Usa chaves para cifrar e decifrar. |
| **Propósito Principal** | Verificar integridade e autenticidade.  | Proteger confidencialidade.        |
| **Exemplos**            | SHA-256, MD5.                           | AES, RSA.                          |

---

## Conceitos Importantes

### 1. Certificados Digitais

- São usados para autenticar a identidade de entidades (pessoas, servidores, etc.).
- Baseados em criptografia assimétrica e emitidos por Autoridades Certificadoras (CAs).

### 2. SSL/TLS

- Protocolos que usam criptografia para proteger a comunicação na internet.
- Garantem que os dados trafegados entre cliente e servidor estejam seguros.

### 3. Criptografia de Ponta a Ponta (End-to-End Encryption)

- Apenas o remetente e o destinatário podem ler as mensagens.
- Usado em aplicativos como WhatsApp e Signal.

### 4. Blockchain

- Tecnologia que usa criptografia para garantir a segurança e integridade de transações.

---

## Boas Práticas em Criptografia

1. **Use Algoritmos Modernos:** Prefira AES, RSA, ECC e SHA-256.
2. **Proteja as Chaves:** Armazene chaves de forma segura (HSMs, cofres de senhas).
3. **Atualize Sistemas:** Mantenha bibliotecas e sistemas atualizados para evitar vulnerabilidades.
4. **Não Implemente sua Própria Criptografia:** Use bibliotecas e frameworks testados e aprovados pela comunidade.

---

## Referências e Links Úteis

- [SSL/TLS Explained](https://www.cloudflare.com/learning/ssl/)
- [AES Encryption](https://en.wikipedia.org/wiki/Advanced_Encryption_Standard)
- [RSA Algorithm](https://en.wikipedia.org/wiki/RSA_(cryptosystem))
- [SHA-256 Hash Function](https://en.wikipedia.org/wiki/SHA-2)
- [OWASP Cryptographic Practices](https://owasp.org/www-project-top-ten/2017/A6_2017-Security_Misconfiguration)

---

## Conclusão

A criptografia é uma ferramenta essencial para proteger informações em um mundo cada vez mais digital. Entender seus
conceitos básicos e boas práticas é fundamental para qualquer pessoa que trabalhe com tecnologia ou segurança da
informação. Para se aprofundar, explore os links de referência e pratique a implementação de algoritmos em projetos
reais.