Nesse problema, precisamos ordenar por **START CRESCENTE**, porém, se os starts forem iguais, precisamos ordenar por **END DECRESCENTE**.

- Exemplo SEM ORDENAR POR END (quando starts começam iguais):
```
Sorted apenas por start: [[1,2],[1,4]]

maxEnd = 0

[1,2]: end=2 <= 0? NÃO → maxEnd = 2
[1,4]: end=4 <= 2? NÃO → maxEnd = 4

Resultado: 2 intervalos sobrevivem
```

Isso está ERRADO, pois [1,2] está coberto por [1,4], mas processou [1,2] primeiro e foi configurado maxEnd=2. Quando chegou em [1,4] ao comparar 4 <= 2 (falso) achou que [1,4] não estava coberto.


- Cenário ordenando POR end (quando starts começam iguais):

```
Sorted: [[1,4],[1,2]]  ← [1,4] vem primeiro agora

maxEnd = 0

[1,4]: end=4 <= 0? NÃO → maxEnd = 4
[1,2]: end=2 <= 4? SIM → coberto!

Resultado: 1 intervalo sobrevive ([1,4])
```

### Importante da ordenação decrescente no segundo

Exemplo: [[1,3], [1,5]]

Com ordenação crescente: processa [1,3] depois [1,5]

Com ordenação correta (decrescente): processa [1,5] depois [1,3]

Na segunda forma, quando chega em [1,3], você já processou [1,5] que cobre o elemento. Fica mais fácil detectar.

## Lógica de atualização

Depois de ordenar corretamente, se intervals[i][1] <= maxEnd, significa que existe um intervalo anterior que:

- Começa antes ou igual (pela ordenação)
- Termina depois ou igual (porque maxEnd é o maior fim visto)

```go
    count := 0 // intervalo não cobertos
	maxEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][1] > maxEnd {
			count++
			maxEnd = intervals[i][1]
		}
		// se interval[1] <= maxEnd, está coberto, não precisa incrementar
	}

	return count
}
```