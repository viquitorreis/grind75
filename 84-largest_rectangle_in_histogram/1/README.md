# Histogram

## Sacada do problema

Primeiro pensei em usar two pointers, nos examples do LC mostra apenas quando tem 2 histogramas juntos, mas isso não funciona, podemos ter diversos histogramas formando um retangulo.

Portanto: a **barra do meio vai ser fator limitante**.

- Para uma barra de altura `h[i]`:

Para uma barra `h[i]` a maior área usando ela como **altura mínima** é:

`h[i] * (índice da primeira barra maior à direita - índice da primeira barra maior à esquerda - 1)`

Precisamos saber: `até onde essa barra se expande para esquerda e direita`, antes de encontrar uma barra menor que ela.

## Stack

É aqui que a stack entra. É uma estrutura que mantém ordem e permite remover do topo de forma eficiente, dessa forma conseguimos pegar a menor barra anterior facilmente.

-> Não tem por que usar two pointers, seria O(n²) no pior caso.
-> Não tem como usar heap, não faz sentido. Heap seria se precisassemos da menor global, que não é o caso.

- Como fazer:

- Empilha enquanto `h[i] >= h[topo]`
- Quando `h[i] < h[topo]`: a barra no topo encontrou "sua primeira menor à direita", e a primeira menor à esquerda é o novo topo após o pop

Exemplo com `[2, 1, 5, 6, 2, 3]`:

```
i=0: stack=[0]
i=1: h[1]=1 < h[0]=2 -> pop 0, width=1 (top vazio), area=2*1=2
     stack=[1]
i=2: stack=[1,2]
i=3: stack=[1,2,3]
i=4: h[4]=2 < h[3]=6 -> pop 3, width=4-2-1=1, area=6*1=6
          h[4]=2 == h[2]=5? não, < -> pop 2, width=4-1-1=2, area=5*2=10 ✓
     stack=[1,4]
i=5: stack=[1,4,5]
fim: adiciona 0 no final pra forçar flush do stack
```

R = 10.

## Código

1. Adicionamos uma barra com altura 0 no final para forçar o flush. Quando chegar nela, toda transação que ainda estiver aberta na stack vai ser fechada.

```go
heights = append(heights, 0)
```

2. Loop externo

Vai varrer cada barra da esquerda para a direita.

Para cada posição `i`, fazemos duas coisas:

- Fecha transações que não podem mais se expandir
- Depois abrimos uma nova.

```go
for i, stack, top := 0, make([]int, len(heights)), -1; i < len(heights); i++ {
```

Dentro desse loop, o `for` interno fecha todas as transações mais caras que a barra atual

```go
for top != -1 && heights[i] <= heights[stack[top]] {
```

-> enquando a barra atual `heights[i]` for menor ou igual ao topo da stack, vamos fazer pop, pois a barra no topo encontrou sua fronteira a direita e nunca mais vai conseguir expandir além de i.

Esse loop vai continuar enquanto o topo da stack for **maior ou igual a heights[i]**. Por que? POrque a barra atual i é uma fronteira, nenhuma transação mais alta que ela pode se expandir além daqui

3. Depois disso, precisamos fechar a transação do **topo**. Essa é a **altura do retângulo** que vamos calcular **agora**.

```go
height := heights[stack[top]]
top--
```

4. Precisamos calcular a janela de tempo (**largura**).

- `top == -1`: não tem transação mais barata ainda aberta -> a janela começa do índice 0. Então `width = i`
- Caso contrário: a janela vai do novo topo até `i-1` -> `width = i - 1 - stack[top]`

```go
if top == -1 {
    width = i
} else {
    width = i - 1 - stack[top]
}
```

Explicando:

Primeiro: a stack é **crescente**.

Quando damos o pop na barra `i` e olha para o novo topo `stack[top]`, sabemos que **tudo entre** `stack[top]` e `i` é **maior ou igual à barra que acabamos de popar**. Se não já teria sido feito o pop antes.

Então a barra que fazemos o pop consegue se expandir por toda essa janela sem encontrar nada menor que ela.

Exemplo:

```go
índices: 0  1  2  3  4
alturas: 1  3  5  3  2
stack antes do pop de 5: [0, 1, 2]  (índices)
```

Quando chega `h[3]=3`. Pop no índice 2 (altura 5). Novo topo = índice 1 (altura 3). Width = `3 - 1 - 1 = 1`.  Barra 5 só ocupa 1 slot pois `esquerda já tem altura 3 < 5.

Chega `h[4] = 2`. Pop índice 3 (altura 3). Novo topo = índice 0 (altura 1). Width = `4 - 1 - 0 = 3`. Altura 3 se expande pelos índice 1, 2, 3

5. Área dessa transação fechada. Compara com o máximo global.

```go
square = max(square, height*width)
```

6. Abrimos uma nova transação (empilha o índice atual)

```go
top++
stack[top] = i
```