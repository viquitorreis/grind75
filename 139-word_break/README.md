# Word Break DP - Trace com "leetcode" e ["leet", "code"]

Inicialização:

```
dp = [T, F, F, F, F, F, F, F, F]
      0  1  2  3  4  5  6  7  8
```

dp[0] = true (base case: string vazia sempre válida)
Resto = false (ainda não sabemos)

i=1 (testa formar "l"):

j=0: dp[0]=T? Sim. "l" existe? Não → dp[1]=F

i=2 (testa formar "le"):

j=0: dp[0]=T? Sim. "le" existe? Não
j=1: dp[1]=F? Não, pula → dp[2]=F

i=3 (testa formar "lee"):

j=0: dp[0]=T? Sim. "lee" existe? Não
j=1: dp[1]=F? Não, pula
j=2: dp[2]=F? Não, pula → dp[3]=F

i=4 (testa formar "leet"):

j=0: dp[0]=T? Sim. "leet" existe? Sim! → dp[4]=T

i=5 (testa formar "leetc"):

j=0: dp[0]=T? Sim. "leetc" existe? Não
j=4: dp[4]=T? Sim. "c" existe? Não → dp[5]=F

i=6 (testa formar "leetco"):

Nenhuma quebra funciona → dp[6]=F

i=7 (testa formar "leetcod"):

Nenhuma quebra funciona → dp[7]=F

i=8 (testa formar "leetcode"):

j=0: dp[0]=T? Sim. "leetcode" existe? Não
j=4: dp[4]=T? Sim. "code" existe? Sim! → dp[8]=T

Resultado final: dp[8]=T → Consegue formar "leetcode" = "leet" + "code"