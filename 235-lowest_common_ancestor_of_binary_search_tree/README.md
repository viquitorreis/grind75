LCA é o nó mais baixo na árvore que tem tanto p quanto q como descendentes.

Ponto chave disso: **Um nó pode ser descendente de si mesmo**...

Por exemplo:

```
      6
     / \
    2   8
   / \
  0   4
     / \
    3   5
```

Tendo p=2, e q=4

O nó 2 é **ancestral de si mesmo**, e também é ancestral de 4...

Não é 6 pois é ancestral comum de 2 e 4, mas **não é o mais baixo**.

Lógica:

- Se ambos p e q são menores que curr -> LCA está à esquerda, desce
- Se ambos p e q são maiores que curr -> LCA está à direita, desce
- Se estão em lados opostos OU curr é p ou q -> curr é o LCA