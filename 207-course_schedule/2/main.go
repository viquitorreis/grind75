package main

func main() {

}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	dependCount := make([]int, numCourses) // iniciamos zerados, vamos aumentar as dependencias

	for _, p := range prerequisites {
		// prerequisitos do curso
		graph[p[1]] = append(graph[p[1]], p[0])
		dependCount[p[0]]++
	}

	// processamos os in degree = 0 primeiro
	depends := []int{}
	for course, count := range dependCount {
		if count == 0 {
			depends = append(depends, course)
		}
	}

	for i := 0; i < len(depends); i++ {
		for _, j := range graph[depends[i]] {
			dependCount[j]--

			if dependCount[j] == 0 {
				depends = append(depends, j)
			}
		}
	}

	return len(depends) == numCourses
}
