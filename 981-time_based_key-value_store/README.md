## Time Based Key Value Store

Esse problema, comecei pensando que poderia simplesmente armazenar o timestamp nas chaves para obter mais rapidamente. Porém isso se torna problemático ao buscar, pois precisamos poder buscar quando tem timestamps maiores que o que foi inserido, e se for muito maior, o tempo fica muito grande para fazer.

Primeira intuição *ERRADA*:

```go
type TimeMap struct {
	keyVal map[string]string
}

func (this *TimeMap) Get(key string, timestamp int) string {
	for ts := timestamp; ts > 0; ts-- {
		k := fmt.Sprintf("%s_%d", key, ts)
		if val, ok := this.keyVal[k]; ok {
			return val
		}
	}

	return ""
}
```

Isso não resolve o problema, assim como mencionado.

Precisa armazenar **múltiplos valores** por chave, e organizar esses valores por timestamp. Logo, sempre que inserirmos, podemos ordenar. Mas nesse problema não precisamos, pois está sempre inserindo de forma crescente de timestamp.

Quando buscarmos, se não existir o valor naquele timestamp, buscamos o valor menor mais próximo dele. É a forma de se fazer isso funcionar.

**Estrutura**

```go
type TimeMap struct {
    store map[string][]Pair  // chave -> lista de (timestamp, value)
}

type Pair struct {
    timestamp int
    value     string
}
```