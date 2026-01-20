# Permutations - Backtracking

Esse problema usa backtracking (ou backtracking) assim como o Combination Sum. Mas tem diferença.

**Combination Sum:** Pode reutilizar números, usa ```start``` para evitar duplicatas.
**Permutations:** Cada número usado ```uma vez```, precisa rastrear quais já foram usados.

## Árvore de decisões

**Input:** [1, 2, 3]

```
                        []
           /            |            \
        [1]           [2]           [3]
       /   \         /   \         /   \
    [1,2] [1,3]   [2,1] [2,3]   [3,1] [3,2]
      |     |       |     |       |     |
   [1,2,3][1,3,2][2,1,3][2,3,1][3,1,2][3,2,1]
```

Em cada nível, você escolhe um número que **ainda não usou.** Quando usou todos, encontrou uma permutação.

### A Lógica em desfazer a escolha

Ao fazer o backtracking, precisamos **remover aquela escolha**, nesse caso usamos um simples array boolean com os índices

Quando você volta do backtrack, está desfazendo aquela escolha. Precisa desfazer TUDO relacionado a ela:

Remove de current (tira da permutação)

Marca used[i] = false (libera para outras permutações)

Se você só remove de current mas não libera no used, o número fica "preso" como usado mesmo não estando mais na permutação.
