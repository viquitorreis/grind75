## DP com Memoização

Partition Equal Subset Sum - Passo a Passo:

1. Soma todos números do array
2. Se soma é ímpar -> return false (impossível dividir igualmente)
3. Soma target é sum / 2 (queremos a metade, como pede no problema)
4. Criamos um array dp, de tamanho target + 1, que irá sinalizar se conseguimos formar a soma com os números vistos até então

Antes do loop: dp = [T, F, F, F, F, T=dp[5], F, F, F, F, F]
    (marcamos dp[5]=true na iteração anterior ou inicial)

```go
func canPartition(nums []int) bool {
    // Calcula soma total
    sum := 0
    for _, v := range nums {
        sum += v
    }
    
    // Se ímpar, impossível dividir igualmente
    if sum%2 != 0 {
        return false
    }
    
    target := sum / 2
    
    // dp[j] = consigo formar soma j com números vistos até agora?
    dp := make([]bool, target+1)
    dp[0] = true  // soma 0 sempre possível (não pegar nada)
    
    // Para cada número
    for _, num := range nums {
        // Percorre de trás pra frente para não usar mesmo número 2x
        for j := target; j >= num; j-- {
            // Se conseguia (j-num) antes, agora consigo j pegando num
            dp[j] = dp[j] || dp[j-num]
        }
    }
    
    return dp[target]
}
```

## Analogia

Esse algoritmo é chato de aprender. Vamos usar uma analogia.

Vamos supor que estamos construindo um motor. Temos várias peças, pistão, virabrequim, biela e válvula. Cada peça tem um peso específico, e queremos saber "consigo montar um conjunto de peças que pese exatamente 11 quilos?"

Agora, pensa que temos uma prateleira com gavetas numeradas de 0 até 11. Cada gaveta representa um peso possível. Se uma gaveta está marcada com um "sim", significa que descobrimos uma combinação de peças que chega naquele peso exato.

Começamos com todas gavetas vazias, mascadas com "NÃO", com excessão da 0, pois é só não pegar nenhuma peça nela... Por isso ```dp[0] = true```

1. Dai pegamos a primeira peça, o pistão, que pesa 1 KG, quais gavetas podemos marcar SIM AGORA com ela?
    a. Começamos na gaveta 11 e vamos descendo. Quando chega na 11, perguntamos "eu já tinha uma combinação de peso 10? Pois, se tinhamos de peso
        10, é só somar com 1 e teremos a soma que precisamos... Então olhamos na gaveta 10, no caso, não vamos ter.
        Continuamos descendo, até chegar na gaveta do peso atual, que é a 1. Como tinhamos combinação 0, para formar a 1, conseguimos. Então marcamos a 1 como SIM.

2. Pegamos a próxima, virabrequim com 5KG. Fazemos o mesmo processo.
    a. Quando chegamos na gaveta 6, perguntamos "eu tinha peso 1?" (6 - 5), e sim, tinhamos. Então pensamos, se eu tinha as peças que pesavam 1 quilo (o pistão), e agora pego esse virabrequim de 5KG, consigo fazer 6KG.
        Então, quando chegamos aqui, conseguimos provas que conseguimos fazer a soma para essa gaveta, usando uma gaveta que a complemente, que já tenha o peso necessário para complementar ela.