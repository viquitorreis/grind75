package main

func main() {

}

func canFinish(numCourses int, prerequisites [][]int) bool {
	startGraph := make([][]int, numCourses) // lista de adjacentcia onde cada indice representa as dependencias de um curso
	inDegree := make([]int, numCourses)     // quantidade de dependencias atuais de cada curso (por indice)

	for _, p := range prerequisites {
		// ex [1,0]: Para fazer curso 1, precisa fazer curso 0 antes
		// 1. mapeando todas todas dependencias usando p[1]. Primeiro elemento (p[0]) é o curso, segundo (p[1]) é a dependencia
		// DEPENDENCIA APONTA PARA QUEM LIBERA.
		startGraph[p[1]] = append(startGraph[p[1]], p[0])
		inDegree[p[0]]++ // o curso p[0] tem mais uma dependencia
	}

	// acha os nodes com in-degree = 0 -> precisamos fazer eles primeiro pois não tem dependências
	inDegreeList := make([]int, 0)
	for node, d := range inDegree {
		if d == 0 {
			inDegreeList = append(inDegreeList, node)
		}
	}

	// de forma iterativa removemos os in-degree = 0 e atualizamos os in-degrees
	for i := 0; i < len(inDegreeList); i++ {
		for _, j := range startGraph[inDegreeList[i]] {
			inDegree[j]--
			// adiciona o node para a lista de in-degree se ela se tornar 0
			if inDegree[j] == 0 {
				inDegreeList = append(inDegreeList, j)
			}
		}
	}

	return len(inDegreeList) == numCourses
}
