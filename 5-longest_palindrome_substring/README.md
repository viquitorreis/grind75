Ideia: ao invés de identificar palíndromos de fora para dentro, expandimos de dentro para fora.

Para cada posição da string, assumimos que aquela posição é o CENTRO de um palindromo e tenta expandir para os lados ENQUANTO os caracteres são iguais.

Sempre que acharmos palavras palindromas maiores, vamos atualizar a resposta.

## Checagem de palindromo

A seguinte função, só funciona pois estamos buscando o maior palíndromo formado por *string contígua*:

```go
palindrome := func(s string, l, r int) string {
    for l >= 0 && r < len(s) && s[l] == s[r] {
        l--
        r++
    }

    return s[l+1 : r]
}
```

`l+1 pois o loop para em s[l] != s[r], mas precisamos de strings iguais, e sabemos que a janela anterior era...`

**Ponto importante:** esse algoritmo é muito inteligente, abaixo mostro o exemplo. O ponto crucial dele, é que ele **começa do centro**, ele assume que a posição que está é o centro, compara o caractere atual consigo próprio (se é centro, vai dar certo, pois o caractere atual é sempre igual ele mesmo), e vai expandindo para as bordas.
    - Left -> uma casa para esquerda a cada iteração
    - Right -> uma casa para a direita a cada iteração

A cada iteração compara o elemento atual na left com o da right, se são iguais então é palindromo.

## Loop principal

Vamos fazer a checagem de todas as substrings da string a partir de um loop principal:

```go
for i := 0; i < len(s); i++ {
```

Na iteração do nosso loop principal, vamos checar:

1. Palíndromo ímpar

```go
first := palindrome(s, i, i)    // checando se valor ímpar é palindromo
```

- Aqui começamos com `l=i e r=i` -> mesmo caractere
- Expandimos: comparamos s[i-1] com s[i+1], e depois s[i-2] com s[i+2]

2. Palíndromo par

```go
second := palindrome(s, i, i+1) // checando se valor par é palindromo
```

- Começa com `l=i e r=i+1` (dois caracteres adjacentes)
- Expande: compara s[i-1] com s[i+2], depois s[i-2] com s[i+3]

Trace com "babad":

```
i=0 (b):
  palindrome(0,0): 
    inicio: l=0, r=0
    loop: s[0]==s[0]? Sim ('b'=='b') → l=-1, r=1
    loop: l<0, para
    retorna s[0:1] = "b"
  
  palindrome(0,1):
    inicio: l=0, r=1
    loop: s[0]==s[1]? Não ('b'!='a'), não entra no loop
    retorna s[1:1] = "" (string vazia)
  
  res = "b"

i=1 (a):
  palindrome(1,1):
    inicio: l=1, r=1
    loop: s[1]==s[1]? Sim ('a'=='a') → l=0, r=2
    loop: s[0]==s[2]? Sim ('b'=='b') → l=-1, r=3
    loop: l<0, para
    retorna s[0:3] = "bab"
  
  palindrome(1,2):
    inicio: l=1, r=2
    loop: s[1]==s[2]? Não ('a'!='b'), não entra
    retorna s[2:2] = ""
  
  res = "bab"

i=2 (b):
  palindrome(2,2):
    inicio: l=2, r=2
    loop: s[2]==s[2]? Sim ('b'=='b') → l=1, r=3
    loop: s[1]==s[3]? Sim ('a'=='a') → l=0, r=4
    loop: s[0]==s[4]? Não ('b'!='d'), para
    retorna s[1:4] = "aba"
  
  palindrome(2,3):
    inicio: l=2, r=3
    loop: s[2]==s[3]? Não ('b'!='a'), não entra
    retorna s[3:3] = ""
  
  res = "bab" (mesmo tamanho que "aba", mantém)

i=3 (a):
  palindrome(3,3):
    inicio: l=3, r=3
    loop: s[3]==s[3]? Sim ('a'=='a') → l=2, r=4
    loop: s[2]==s[4]? Não ('b'!='d'), para
    retorna s[3:4] = "a"
  
  palindrome(3,4):
    inicio: l=3, r=4
    loop: s[3]==s[4]? Não ('a'!='d'), não entra
    retorna s[4:4] = ""
  
  res = "bab"

i=4 (d):
  palindrome(4,4):
    inicio: l=4, r=4
    loop: s[4]==s[4]? Sim ('d'=='d') → l=3, r=5
    loop: r>=len(s), para
    retorna s[4:5] = "d"
  
  palindrome(4,5):
    inicio: l=4, r=5
    loop: r>=len(s), não entra
    retorna s[5:5] = ""
  
  res = "bab"
```