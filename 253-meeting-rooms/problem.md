# Meeting Rooms II

Given an array of meeting time interval objects consisting of start and end times [[start_1,end_1],[start_2,end_2],...] (start_i < end_i), find the minimum number of days required to schedule all meetings without any conflicts.

Note: (0,8),(8,10) is not considered a conflict at 8.

Example 1:

Input: intervals = [(0,40),(5,10),(15,20)]

Output: 2

Explanation:
day1: (0,40)
day2: (5,10),(15,20)

Example 2:

Input: intervals = [(4,9)]

Output: 1

```go
type Interval struct {
	start int
	end int
}

func minMeetingRooms(intervals []Interval) int {

}
```

## Solução

É um problema de merge intervals, e como sempre, esse tipo de problema precisa de ordenação dos intervalos. Porém, aqui é diferente:

1. Ordenamos os INTERVALOS INICIAIS por START

    Isso é padrão, precisamos colocar ordem neles. Dessa forma vamos poder organizar corretamente na fila.

2. A FILA de prioridades precisa ordenar por END (término da reunião)

    Isso por quê precisamos ter sempre no topo a reunião que termina mais cedo, para saber se podemos reusar a sala...
