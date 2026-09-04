# Simple Compiler

Um compilador simples desenvolvido em Go.

O projeto possui:

- Analisador léxico (`scanner`)
- Tokens (`token`)
- Analisador sintático (`parser`)
- Geração de código intermediário baseado em pilha
- Interpretador (`interpreter`)
- Operações `+`, `-`, `*` e `/`
- Atribuição com `let`
- Impressão com `print`

## Estrutura

```text
simple_compiler/
├── cmd/
│   └── compiler/
│       └── main.go
├── scanner/
│   └── scanner.go
├── parser/
│   └── parser.go
├── interpreter/
|   ├── command.go
│   └── interpreter.go
├── token/
│   └── token.go
├── go.mod
├── Dockerfile
├── compose.yaml
└── README.md
```

## Pré-requisitos

### Executar diretamente no computador

Instale o Go e verifique:

```bash
go version
```

### Executar com Docker

Instale:

- Docker
- Docker Compose

Verifique:

```bash
docker --version
docker compose version
```

No Windows, o Docker Desktop precisa estar em execução para os comandos Docker funcionarem.

---

# 1. Executando sem Docker

## Executar diretamente com `go run`

A forma mais simples durante o desenvolvimento é:

```bash
go run ./cmd/compiler
```

Isso compila temporariamente o programa e já executa.

---

# 2. Gerando um executável

Para gerar um executável do projeto:

```bash
go build -o simple_compiler ./cmd/compiler
```

No Linux/macOS será gerado:

```text
simple_compiler
```

No Windows, é possível gerar diretamente:

```bash
go build -o simple_compiler.exe ./cmd/compiler
```

Depois disso, basta executar.

### Windows

```powershell
./simple_compiler.exe
```

### Linux/macOS

```bash
./simple_compiler
```

Assim, depois do primeiro `go build`, você não precisa mais usar `go run`.

Sempre que alterar o código, basta recompilar:

```bash
go build -o simple_compiler.exe ./cmd/compiler
```

---

# 3. Executar com Docker

O projeto também pode ser executado dentro de um container.

Primeiro, certifique-se de que o Docker Desktop está em execução.

Depois:

```bash
docker compose up --build
```

O `--build` garante que a imagem seja reconstruída quando necessário.

Para parar:

```bash
docker compose down
```

---

# 4. Exemplo do programa

O `main.go` atualmente pode conter uma entrada semelhante a:

```go
package main

import (
    "simple_compiler/interpreter"
    "simple_compiler/parser"
)

func main() {
    input := `
let a = 42 + 2;
let b = 15 + 3;
print a + b;
`

    p := parser.NewParser([]byte(input))
    p.Parse()

    i := interpreter.NewInterpreter(p.Output())
    i.Run()
}
```

Executando:

```bash
go run ./cmd/compiler
```

o resultado esperado é:

```text
62
```

---

# 5. O que acontece durante a execução?

O fluxo do projeto é:

```text
Código-fonte
     │
     ▼
  Scanner
     │
     ▼
   Tokens
     │
     ▼
   Parser
     │
     ▼
Código intermediário
     │
     ▼
Interpretador
     │
     ▼
  Resultado
```

Por exemplo:

```text
let a = 42 + 2;
let b = 15 + 3;
print a + b;
```

O parser produz instruções semelhantes a:

```text
push 42
push 2
add
pop a
push 15
push 3
add
pop b
push a
push b
add
print
```

O interpretador executa essas instruções utilizando uma pilha e um mapa de variáveis.

Resultado:

```text
62
```

---

# 6. Gramática da linguagem

A linguagem implementada neste projeto utiliza uma gramática simples para representar programas formados por comandos de atribuição e impressão.

A gramática atual é:

```text
program         → statements

statements      → statement*

statement       → letStatement
                | printStatement

letStatement    → 'let' IDENT '=' expression ';'

printStatement  → 'print' expression ';'

expression      → term oper

oper            → '+' term oper
                | '-' term oper
                | ε

term            → factor termOper

termOper        → '*' factor termOper
                | '/' factor termOper
                | ε

factor          → NUMBER
                | IDENT
```

## 6.1 Programa

Um programa é uma sequência de zero ou mais `statements`:

```text
program → statements

statements → statement*
```

O `*` significa "zero ou mais ocorrências".

Por exemplo:

```text
let a = 10;
let b = 20;
print a + b;
```

possui três `statements`.

---

## 16.2 Statements

Existem dois tipos de comandos:

```text
statement → letStatement
          | printStatement
```

### Atribuição

```text
letStatement → 'let' IDENT '=' expression ';'
```

Exemplo:

```text
let a = 42 + 5;
```

### Impressão

```text
printStatement → 'print' expression ';'
```

Exemplo:

```text
print a + 10;
```

---

## 6.3 Expressões

Uma expressão é formada por um `term` seguido de zero ou mais operações de soma ou subtração:

```text
expression → term oper
```

As operações são:

```text
oper → '+' term oper
     | '-' term oper
     | ε
```

O símbolo `ε` representa a **produção vazia**: significa que não há mais operações.

Exemplos:

```text
42
42 + 5
42 - 5
42 + 5 - 8
```

---

## 6.4 Termos

Os termos são utilizados para implementar a precedência de multiplicação e divisão:

```text
term → factor termOper
```

```text
termOper → '*' factor termOper
         | '/' factor termOper
         | ε
```

Assim, `*` e `/` possuem maior precedência que `+` e `-`.

Por exemplo:

```text
2 + 3 * 4
```

é interpretado como:

```text
2 + (3 * 4)
```

e não:

```text
(2 + 3) * 4
```

---

## 6.5 Fatores

Um `factor` representa o elemento básico de uma expressão:

```text
factor → NUMBER
       | IDENT
```

Portanto, atualmente uma expressão pode utilizar:

- números;
- identificadores (variáveis).

Exemplos:

```text
42
123
a
preco
valor_total
```

---

## 6.6 Números

Os números são reconhecidos pelo analisador léxico como uma sequência de dígitos.

Exemplos:

```text
0
5
42
876
12345
```

O token produzido é:

```text
NUMBER
```

---

## 6.7 Identificadores

Identificadores são formados por letras, `_` e, depois do primeiro caractere, também podem conter números.

A regra utilizada pelo Scanner é equivalente a:

```text
IDENT → letra | '_' seguido de zero ou mais caracteres alfanuméricos ou '_'
```

Exemplos válidos:

```text
a
abc
preco
valor1
_meuValor
```

O Scanner primeiro identifica a sequência como um identificador e depois consulta a tabela de palavras-chave.

---

## 6.8 Palavras-chave

Atualmente a linguagem possui duas palavras-chave:

```text
let
print
```

Elas são classificadas como:

```text
let   → LET
print → PRINT
```

Enquanto:

```text
a
preco
valor
```

continuam sendo classificados como:

```text
IDENT
```

---

## 6.9 Operadores e símbolos

Os símbolos reconhecidos atualmente são:

| Símbolo | Token | Função |
|---|---|---|
| `+` | `PLUS` | Soma |
| `-` | `MINUS` | Subtração |
| `*` | `MULT` | Multiplicação |
| `/` | `DIV` | Divisão |
| `=` | `EQ` | Atribuição |
| `;` | `SEMICOLON` | Finaliza um statement |

Além deles, o Scanner produz:

```text
EOF
```

para indicar o fim da entrada.

---

## 6.10 Precedência dos operadores

A estrutura da gramática estabelece a seguinte ordem de precedência:

```text
Maior precedência
       │
       ▼
      * /
       │
       ▼
      + -
       │
       ▼
Menor precedência
```

Por exemplo:

```text
10 + 20 * 3
```

gera uma sequência intermediária equivalente a:

```text
push 10
push 20
push 3
mul
add
```

O resultado é:

```text
70
```

A precedência é definida pela estrutura da gramática, e não pelo Interpretador.

---

## 6.11 Exemplo completo

Entrada:

```text
let a = 42 + 5 * 2;
let b = 10;
print a + b;
```

Estrutura simplificada:

```text
program
└── statements
    ├── letStatement
    │   ├── let
    │   ├── a
    │   ├── =
    │   ├── expression
    │   │   ├── 42
    │   │   ├── +
    │   │   └── 5 * 2
    │   └── ;
    │
    ├── letStatement
    │   └── ...
    │
    └── printStatement
        └── expression
```

O Parser produz código intermediário semelhante a:

```text
push 42
push 5
push 2
mul
add
pop a
push 10
pop b
push a
push b
add
print
```

O Interpretador então executa essas instruções utilizando:

- uma **pilha** para os valores;
- um **mapa de variáveis** para armazenar os identificadores.
