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