Importante: tem que atualizar o lastEnd para a **intersecção** (menor end). E quando não tem overlap.

Ideia do problema: Balões sobrepostos podem ser estourados com uma única flecha atirada na interseção deles. Você precisa de nova flecha só quando não há mais overlap.

Exemplo:

```
Input: [[10,16],[2,8],[1,6],[7,12]]
Sorted by END: [[1,6],[2,8],[7,12],[10,16]]
```

- Pense assim: onde você atira a primeira flecha para pegar o máximo de balões possível?

```
arrows = 1
lastEnd = 6  (atirar em x=6 pega [1,6])

i=1: [2,8]
  2 <= 6? SIM → overlap!
  Interseção: [2,6] (posso atirar em qualquer x entre 2 e 6)
  Para pegar ambos, devo atirar no máximo em x=6
  lastEnd = min(6, 8) = 6  (atualiza para interseção)

i=2: [7,12]
  7 <= 6? NÃO → sem overlap, nova flecha!
  arrows = 2
  lastEnd = 12

i=3: [10,16]
  10 <= 12? SIM → overlap!
  lastEnd = min(12, 16) = 12

Resultado: 2 flechas
```

Visualização:

```
Flecha 1 em x=6:  [1----6]
                      [2------8]

Flecha 2 em x=12:        [7--------12]
                            [10-----------16]
```