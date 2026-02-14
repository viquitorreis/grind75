Esse problema, precisa ser resolvido usando inorder traversal.

É necessário pois precisamos organizar os elementsopara buscar de forma crescente.

Dessa forma:

```
      5
     / \
    3   6
   / \
  2   4
 /
1
```

Inorder visita: 1, 2, 3, 4, 5, 6 (ordenado crescente)

Se precisarmos de:

- k = 1 -> retornamos 1
- k = 3 -> retornamos 3

Vamos usar:

- Pilha
- Count (até k)
- Ponteiro para o current (nó atual)

Fazemos um loop enquanto current != nil e o tamanho da pilha for menor que k

1. Primeiro vai tudo para a esquerda
2. Processa os nodes e soma count, se chegar em k retorna
3. Vai para a direita

**Entendendo o código**

1. Stack

Precisamos de uma stack para armazenar os nodes visitamos.

A sequencia é:

- Visita tudo a esquerda
- Visita a raiz
- Visita tudo a direita

Essa stack precisa iniciar vazia, pois não processamos a root inicialmente. Pois a root não é o menor elemento, e precisamos de uma stack ordenada pelos melhores elementos.

2. Loop externo

Precisamos de uma lógica onde o curr != nil OU o tamanho da stack for maior que 0.

- Quando curr == nil, não tem mais nada para processar, não tem mais direita para descer
- Quando stack está vazia, não tem mais nós pendentes para processar, então tem que sair

```go
for curr != nil || len(stack) > 0 {
```

3. Loop interno

Nesse loop interno, vamos descer o máximo possível para a esquerda, empurrando TODOS os nós incluindo o atual
