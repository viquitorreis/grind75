Por que precisamos de uma stack?

Quando encontramos um `(`, está dizendo "vou começar uma sub-expressão nova guarda o que eu tinha até agora". Quando encontramos um `)`, estamos dizendo "a sub-expressão terminou, pega o que tinha guardado e combina com o resultado desta sub-expressão".

Isso é exatamente o comportamento de uma stack. Você empurra o **estado atual quando abre parênteses**, e restaura o estado anterior quando fecha.

---

**O que precisamos rastrear**

A cada momento que estamos lendo a string, precisamos saber três coisas:

- Resultado acumulado até agora: `result`
- Sinal do próximo número que vamos ler: `sign` -> começa com +1 (soma)
- Número que estamos construindo no momento: `num` -> número pode ter múltiplos dígitos como "10" por exemplo

Quando encontramos um operador `+` ou `-`:

1. Aplicamos o número atual ao resultado, (`resultado += sign + num`)
2. Resetamos o `num` para zero
3. Atualizamos o `sign` para o próximo número

---

Trace com `10+4-6`

Lendo '1':
    `num = 1`

Lendo '0':
    `num = 10`

Lendo '+':
    aplica `result += 1 * 10 = 10`
    reseta
        `num = 0`
        `sign = +1`

Lendo '4':
    `num = 4`

Lendo '-':
    aplica `result += 1 * 4 = 14`
    reseta
        `num = 0`
        `sign = -1`

Lendo '6':
    `num = 6`
    Fim da string: `result += -1 * 6 = 8`

---

**O que a stack adiciona para o parenteses?**

Quando encontramos `(`:

- Empurramos `result` atual e o `sign` atual na stack.
- E resetamos ambos, como se estivsse começando a expressão do zero.

Operação: `result = stack.pop() + stack.pop() * result`

    O primeiro pop, devolve `result` anterior
    O segundo pop devolve o `sign` que estava ativo quando o ( foi encontrado

Quando encontramos `)`:

- Pegamos / Consumimos o que estava na stack: o resultado interno que acabamos de calcular é combinado com o `result`.
- O resultado interno que acabamos de calcular é combinado com o `result` anterior e o `sign` que estava guardado

Na stack, vamos SEMPRE ter elementos agregados em PARES.
Vamos sempre pegar os dois últimos elementos.

```go
res += num * sign
res *= stack[len(stack)-1]
res += stack[len(stack)-2]
sign = 0
```