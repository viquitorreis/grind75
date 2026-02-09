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

