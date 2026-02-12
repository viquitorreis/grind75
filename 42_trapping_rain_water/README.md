Nesse problema, para cada posição i, a água acumulada é `min(maxLeft, maxRight) - height[i]`, ou seja:

```min(altura do maior muro à esquerda, altura do maior muro à direita) - altura atual```

A largura sai de graça, **não precisamos calcular ela explicitamente**

## Two Pointers

Versão naive:

- Faz dois loops, pré-computando `maxLeft[i]` e `maxRight[i]` para cada i.
- O two pointers elimina os arrays auxiliares.

**Ponto importante**: não precisamos saber o `maxRight` exato se ja sabemos que maxLeft < maxRight. Nesse caso o maxLeft é fator limitante e vamos precisar calcular a água da esquerda sem olhar para a direita.

---

### 1. Por que `if height[l] >= maxL` não soma?

a. Quando chegamos a uma posição e ela é mais alta que tudo que vimos á esquerda, ela se torna o **novo muro esquerdo**. NÃO TEM ÁGUA EM CIMA DE UM MURO, pois ela iria escorrer. Então aqui atualizamos, `maxL = heigh[l]` e o loop segue.

b. Quando a posição é MENOR que maxL, estamos em um vale. Então aqui temos água, que seria `maxL - height[l]`. que é a coluna de água naquele ponto, pois o fator limitante está a esquerda.

### 2. Por que processar apenas um lado é suficiente?

Quando entramos na condição `height[l] <= height[r]`, isso significa que `maxR` atual é no minimo tão alto quanto `height[r]`, e tão alto quanto `height[l]`. 

Então, sabemos que o **lado direito não vai ser o fator limitante para a posição l**, o lado esquerdo é, pois se o lado esquerdo é menor, a água entre os dois precisa estar no máximo até a altura do menor. Então quando calculamos a agua em l, usamos só `maxL` sem nem precisar saber o `maxR` exato, pois sabemos que não vai transbordar (estamos pegando a altura do menor...).

Então quando calculamos e estamos na esquerda, olhamos para `maxL`, a parte esquerda precisa saber a parte direita para saber se fica ou escorre? Só se o muro direito for menor que o esquerdo. MAS, ja sabemos que `height[r] >= height[l]` então existe pelo menos esse muro à direita. Então a gota fica, e a quantidade é `maxL - height[l]`

### 3. Calculo trapped += maxL - height[l] OU trapped += maxR - height[r]

Como sabemos que esse calculo funciona?

**Ponto**: estamos calculando quanto de água cabe NAQUELA COLUNA ESPECÍFICA.

Pensa que estamos em uma posição i, com `height[i] = 1`, e o maior muro a esquerda é `maxL = 3`, o maior muro à direita é `maxR = 5`. Quanto de água cabe NESSA COLUNA?

A água sobe até o nível do MENOR dos dois muros. Logo, o nível da água vai ser `min(3,5) = 3`. A coluna de água é `nível - chão = 3 - 1 = 2`, que é exatamente min(maxL, maxR) - height[i].

O two pointers só faz ser desnecessário o uso do min explicito, pois dentro do bloco específico já temos a altura minima.

---


**Algoritmo**

```go
l, r := 0, len(height)-1
maxL, maxR := 0, 0
trapped := 0

for l < r {
    if height[l] <= height[r] {
        // 1. quando max left é fator limitante, aumentamos o muro à esquerda
        // 		pois não tem como represar a agua nesse intervalo
        // 2. caso contrário aumentamos a quantidade de água represada
        if height[l] >= maxL {
            maxL = height[l]
        } else {
            trapped += maxL - height[l]
        }

        l++
    } else {
        // 1. max right é fator limitante, aumentamos o muro a direita.
        // 2. caso contrario somamos na agua represada
        if height[r] >= maxR {
            maxR = height[r]
        } else {
            trapped += maxR - height[r]
        }

        r--
    }
}
```


