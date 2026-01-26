# Coin change

Exemplo:

coins = [1, 2], amount = 4

## Passo 1 - Inicialização

Iniciamos um array do tamanho amount (sim, pode ser que fique gigante, e é o contra de dynamic programming).

Cada posição do array precisa ter max infinito.

A **posição 0** precisa ser zero.

```
minim = [0, ∞, ∞, ∞, ∞]
índice   0  1  2  3  4

Tradução: minim[i] = quantas moedas mínimas preciso para formar valor i?
minim[0] = 0 (não preciso de moedas para formar zero)
minim[1] = ∞ (ainda não sei)
minim[2] = ∞ (ainda não sei)
minim[3] = ∞ (ainda não sei)
minim[4] = ∞ (ainda não sei)
```

**Obs:** O primeiro elemento precisa ser 0, pois precisamos de 0 coins para formar esse amount.

## Passo 2 - Loop em coins (loop externo)

```go
for i := coin; i <= amount; i++
```

Nesse loop, poderiamos começar do índice 0.

Porém, não faz sentido processar índices zero até o valor da coin. É uma otimização, não é obrigatório, mas economiza iterações inúteis.

**Comparação:**

```go
if minim[i-coin] != math.MaxInt
```

Nessa comparação estamos **checando se o anterior é diferente de maxInt**.

Então estamos vendo se é possível formar o valor ```i-coin```.

Se for maxInt, significa que ainda não conseguiu formar aquele valor. Então não adianta tentar usar essa moeda.

**Atribuição:**

Vamos atribuir valor ao elemento atual.

```go
minim[i] = min(minim[i], minim[i-coin]+1)
```

- minim[i]: Nesse elemento, pode ser tanto max inf, quanto o valor da ultima iteração com a coin passada.

O que significa **minim[i-coin]+1**

**=>** é o valor que está guardado na posição anterior + 1.

Significa **quantas moedas eu preciso para formar o resto, mais a moeda que eu acabei de usar.**

Exemplo: se formar valor dez e tem moeda de valor tres:

- resto = 10 - 3 = 7
- minim[7] te diz quantas moedas precisa para formar 7 (digamos que seja duas)
- minim[7] + 1 = duas moedas para formar 7, mais a moeda de 3 que usamos agora = 3 moedas total

---

**Iteração 1: i = 1**

- Pergunta: "Posso formar valor 1 usando uma moeda de valor 1?"

Verifico: minim[i-coin] = minim[1-1] = minim[0] = 0 (não é ∞, pois começa em 0)

```
Cálculo: minim[1] = min(minim[1], minim[0]+1)
                  = min(∞, 0+1)
                  = min(∞, 1)
                  = 1
```

*Resultado*: minim[1] = 1
*Significado*: "Preciso de 1 moeda para formar valor 1"

Estado atual: minim = [0, 1, ∞, ∞, ∞]

**Iteração 2: i = 2**

Pergunta: "Posso formar valor 2 usando uma moeda de valor 1?"

Verifico: minim[i-coin] = minim[2-1] = minim[1] = 1 (não é ∞)

```
Cálculo: minim[2] = min(minim[2], minim[1]+1)
                  = min(∞, 1+1)
                  = min(∞, 2)
                  = 2
```

*Resultado*: minim[2] = 2
*Significado*: "Preciso de 2 moedas de valor 1 para formar valor 2"

Estado atual: minim = [0, 1, 2, ∞, ∞]

**Iteração 3: i = 3**

Pergunta: "Posso formar valor 3 usando uma moeda de valor 1?"

Verifico: minim[i-coin] = minim[3-1] = minim[2] = 2 (não é ∞)

```
Cálculo: minim[3] = min(minim[3], minim[2]+1)
                  = min(∞, 2+1)
                  = min(∞, 3)
                  = 3
```

*Resultado*: minim[3] = 3
*Significado*: "Preciso de 3 moedas de valor 1 para formar valor 3"

Estado atual: minim = [0, 1, 2, 3, ∞]

**Iteração 4: i = 4**

Pergunta: "Posso formar valor 4 usando uma moeda de valor 1?"

Verifico: minim[i-coin] = minim[4-1] = minim[3] = 3 (não é ∞)

```
Cálculo: minim[4] = min(minim[4], minim[3]+1)
                  = min(∞, 3+1)
                  = min(∞, 4)
                  = 4
```

Resultado: minim[4] = 4
Significado: "Preciso de 4 moedas de valor 1 para formar valor 4"

Estado final após processar coin=1: minim = [0, 1, 2, 3, 4]
