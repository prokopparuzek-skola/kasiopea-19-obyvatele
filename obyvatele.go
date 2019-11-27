//nefunguje spocti permutace
package main

import "fmt"

func addOne(count []int, graph [][]int, K int) bool {
	for i := 1; i < len(count); i++ {
		c := count[i]
		if c != K {
			c++
		} else {
			continue
		}
		ver := make([]bool, K-c+1)
		for j := 0; j < len(graph[i]); j++ {
			if count[graph[i][j]] >= c {
				ver[K-count[graph[i][j]]] = true
			}
		}
		for _, v := range ver {
			if v == false {
				return true
			}
		}
	}
	return false
}

func main() {
	var T int
	fmt.Scanf("%d", &T)
	for ; T > 0; T-- {
		var N, M, K int
		fmt.Scanf("%d%d%d", &N, &M, &K)
		count := make([]int, N+1)
		graph := make([][]int, N+1)
		for i := 1; i < len(graph); i++ {
			graph[i] = make([]int, 0)
		}
		for i := 1; i < len(count); i++ {
			fmt.Scanf("%d", &count[i])
		}
		for i := 0; i < M; i++ {
			var from, to int
			fmt.Scanf("%d%d", &from, &to)
			graph[from] = append(graph[from], to)
			graph[to] = append(graph[to], from)
		}
		can := addOne(count, graph, K)
		if can {
			fmt.Println("ANO")
		} else {
			fmt.Println("NE")
		}
	}
}
