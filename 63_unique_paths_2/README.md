# Unique Paths 2

Esse exercício é quase o 1, porém:

**Ponto 1:** Se uma célula tem obstáculo, não existem caminhos passando por ela, então você coloca zero nessa posição do DP.

```
Grid original (0 = livre, 1 = obstáculo):
[0] [0] [0]
[0] [1] [0]  <- obstáculo aqui
[0] [0] [0]

DP que você vai construir:
[1] [1] [1]
[1] [0] [1]  <- obstáculo vira 0 no DP
[1] [1] [2]
```

**Ponto 2:** Quando preenche a primeira linha e primeira coluna, se encontrar um obstáculo, tudo depois dele também fica zero. Pois não tem como contonar.

Exemplo:

```
Grid com obstáculo na primeira linha:
[0] [1] [0] [0]
     ↑ obstáculo

DP resultante na primeira linha:
[1] [0] [0] [0]  ← depois do obstáculo, tudo zero (não tem como chegar)
```

