# Combination Sum - Backtracking

Esse algoritmo precisa ser resolvido com **backtracking**. Se não encontrarmos a solução, voltamos atrás e escolhemos outras.

Precisamos **explorar cada ramo** até:

- Chegar em target = 0 = solução encontrada
- Passar do target = caminho inválido -> **backtrack**

## Estrutura do Bactracking

1. Base cases
    a. remaining == 0. Adiciona current ao resultado e retorna
    b. remaining < 0. Caminho inválido, retorna (backtrack)
2. Tenta cada candidato a partir de 'start'
    a. escolhe um current (atual candidato)
    b. explora o caminho, chama a função recursiva com - current, valor restante
    c. desfaz -> remove current do array de currents

```go
func backtrack(current []int, remaining int, start int) {
    // Base cases
    if remaining == 0 {
        // Encontrou uma combinação válida!
        adiciona current no resultado
        return
    }
    if remaining < 0 {
        // Passou do target, caminho inválido
        return
    }
    
    // Tenta cada candidato a partir de 'start'
    for i := start; i < len(candidates); i++ {
        current = append(current, candidates[i])  // Escolhe
        backtrack(current, remaining - candidates[i], i)  // Explora tudo antes de ir para próxima linha
        current = current[:len(current)-1]  // Desfaz (backtrack)
    }
}
```

## Por Que `start` é Necessário?

Para evitar duplicatas. Se você sempre começa do índice zero, vai gerar `[2,3]` e `[3,2]` que são a mesma combinação.

Ao passar `i` (posição atual) como `start` na próxima chamada, você garante que só olha para frente, não para trás.

Start não garante que você não viu o número. Garante que você não olha para trás no array.

Exemplo: candidates = [2,3,5]

Se start = 1, você pode pegar 3 múltiplas vezes (3, 3, 3...). Mas você não pode voltar e pegar o 2. Isso evita gerar [2,3] e [3,2] que são duplicatas.

## Importante

Por que `i` e não `i + 1` na recursão?

-> Pois o problema permite **reutilizar** o MESMO NÚMERO. Se passar i+1, nesse caso cada número seria usado apenas uma vez.

