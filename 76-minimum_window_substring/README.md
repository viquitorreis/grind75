Esse problema precisamos de:

**Três estruturas de dados**

- 1 array ou hashmap (need): guarda quantas vezes cada caractere aparece em t
- 1 array ou hashmap (window): guarda quantas vezes cada caractere aparece na janela atual de s
- 1 variável int (formed): rastreia quantos caracteres únicos já satisfazemos completamente

Dois ponteiros:

- ponteiro left: começa em 0, avança quando E ENQUANTO tivermos uma janela válida e queremos encolher ela. Quando a janela deixa de ser válida porque removeu um caractere, paramos de mover left e voltamos a expandir right.
- ponteiro right: começa em 0 e avança junto com a string s. Sempre que right avançar, adicionamos o caractere na janela e atualizamos o contador.

## Lógica:

1. Constroi o map / array need iterando em t. Precisamos saber quantas vezes cada caractere aparece.
2. Loop principal: itera right de 0 até s.
    a. Pega o caractere em s na posição right e adiciona ele no map window -> aumentando sua contagem na janela.
    b. Se o caractere existe em need E a contagem no window agora é igual a contagem no need, incrementamos formed porque satisfez mais um caractere único.
    c. Checamos se formed é igual ao número de caracteres únicos em t.
        Se sim: significa que a janela atual tem todos os caracteres necessários de t.
            
            **Loop de encolhimento**
            
            Aqui entramos em um **loop interno** onde tentamos encolher a janela movendo left para a direita.
                Nesse loop interno, primeiro checamos se a janela atual é menor que a janela que tinhamos guardado antes. Se for, atualizamos a menor janela com as posições atuais de left e right.
            Depois pegamos o caractere em s na posição left (primeiro caractere), remove ele da janela decrementando sua contagem no map / array window.
                Se esse caractere existe em need e a contagem agora no window ficou menor que a contagem no need (necessária), decrementamos formed, pois deixou de satisfazer esse caractere específico. 
            No fim, incrementamos left, encolhendo a janela pela esquerda.

            Continua o loop de encolhimento enquanto formed ainda for igual ao número de caracteres únicos necessários, ou seja, enquanto a janela for válida.

            No fim do loop de encolhimento, a janela deixou de ser válida, então voltamos para o loop externo e right avançá para o próximo caractere, expandindo a janela novamente.

### Códgio

```go
// map need conta quantas vezes cada caractere aparece em t
need := make(map[byte]int)
for i := 0; i < len(t); i++ {
    need[t[i]]++
}
```

```go
// número de caracteres únicos em t que precisamos satisfazer
required := len(need)
```

```go
// map que rastreia quantas vezes cada caractere aparece na janela atual
window := make(map[byte]int)
```

```go
// variável que conta quantos caracteres únicos já satisfizemos completamente
formed := 0
    
// ponteiros da janela deslizante
left := 0
```

```go
// variáveis para guardar a menor janela válida encontrada
// inicializamos minLen com um valor muito grande
minLen := len(s) + 1
minLeft := 0
minRight := 0
```

- Loop Principal

```go
// loop principal: right expande a janela pela direita
for right := 0; right < len(s); right++ {
    // pegamos o caractere na posição right
    char := s[right]
    // adicionamos o caractere na janela incrementando sua contagem
    window[char]++
```

```go
// caso esse caractere existe em need e agora temos a quantidade exata dele,
// incrementamos formed porque satisfizemos mais um caractere único
if count, exists := need[char]; exists && window[char] == count {
    formed++
}
```

- Loop Interno

```go
// agora tentamos encolher a janela enquanto ela ainda for válida
// uma janela é válida quando formed == required
for formed == required && left <= right {
    // calcula o tamanho da janela atual
    currentLen := right - left + 1
```

```go
// se essa janela é menor que a menor janela que encontramos até agora,
// atualizamos os valores da menor janela
if currentLen < minLen {
    minLen = currentLen
    minLeft = left
    minRight = right
}
```

```go
// agora tentamos encolher a janela removendo o caractere da esquerda
leftChar := s[left]

// remove o caractere da janela decrementando sua contagem
window[leftChar]--

// se esse caractere existe em need e agora temos menos dele do que precisamos,
// decrementamos formed porque deixamos de satisfazer esse caractere
if count, exists := need[leftChar]; exists && window[leftChar] < count {
    formed--
}

// move o ponteiro left para a direita, encolhendo a janela
left++
```

Final:

```go
// se nunca encontramos uma janela válida, minLen ainda será maior que len(s)
if minLen > len(s) {
    return ""
}

// retorna a substring correspondente à menor janela válida
return s[minLeft : minRight+1]
```