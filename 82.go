package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)


func readMatrix(fileName string) []int {

	file, error := os.Open(fileName, os.O_RDONLY, 0)
	if error != nil {
		panic(error)
	}
	defer file.Close()


	fileContentBytes, error := ioutil.ReadAll(file)
	if error != nil {
		panic(error)
	}


	numbers := bytes.Split(fileContentBytes, []byte{','}, -1)

	matrix := make([]int, len(numbers))

	for i := 0; i < len(numbers); i += 1 {
		number, _ := strconv.Atoi(string(numbers[i]))
		matrix[i] = number
	}

	return matrix[0:len(matrix)-1]

}


func right(pos, dimx, dimy int) (neighborPos int, works bool) {
	if (pos+1) % dimy == 0 {
		return 0, false
	} else {
		return pos+1, true
	}
	panic("w00t")
}

func bottom(pos, dimx, dimy int) (neighborPos int, error bool) {
	if pos >= dimx * dimy - dimx {
		return 0, false
	} else {
		return pos+dimx, true
	}
	panic("w00t")
}

func top(pos, dimx, dimy int) (neighborPos int, error bool) {
	if pos < dimx {
		return 0, false
	} else {
		return pos-dimx, true
	}
	panic("w00t")
}


func smallestDist(dist []int64, allNodes map[int]bool) int {

	minDist := int64(0x7fffffffffffffff)
	pos := -1

	for k, _ := range allNodes {
		if dist[k] < minDist {
			minDist = dist[k]
			pos = k
		}
	}

	return pos
}


func dijkstra(matrix []int, dimx, dimy int, start int) []int {

	//fmt.Println(matrix)

	allNodes := make(map[int]bool)
	for i, _ := range matrix {
		allNodes[i] = true
	}

	dist := make ([]int64, len(matrix))
	previous := make ([]int, len(matrix))

	for i, _ := range matrix {
		dist[i] = 0x7fffffffffffffff //max
		previous[i] = -1 // undefined
	}


	dist[start] = int64(matrix[start]) //value of the first matrix entry

	for ;len(allNodes) > 0; {

		u := smallestDist(dist, allNodes)
		if u == -1 {
			panic("no smallest dist")
		}


		//fmt.Println(u, dist[u])


		if dist[u] == 0x7fffffffffffffff {
			break
		}

		allNodes[u] = false, false


		rightNode, rightWorks := right(u,dimx,dimy)
		bottomNode, bottomWorks := bottom(u,dimx,dimy)
		topNode, topWorks := top(u,dimx,dimy)

		// fmt.Println("  ", rightNode, rightWorks, matrix[rightNode], dist[u] + int64(matrix[rightNode]))
		// fmt.Println("  ", bottomNode, bottomWorks, matrix[bottomNode], dist[u] + int64(matrix[bottomNode]))
		// fmt.Println("  ", topNode, topWorks, matrix[topNode], dist[u] + int64(matrix[topNode]))

		/* look at the neighbors */
		if rightWorks == true {
			alt := dist[u] + int64(matrix[rightNode])
			if alt < dist[rightNode] {
				dist[rightNode] = alt
				previous[rightNode] = u
			}
		}
		if bottomWorks == true {
			alt := dist[u] + int64(matrix[bottomNode])
			if alt < dist[bottomNode] {
				dist[bottomNode] = alt
				previous[bottomNode] = u
			}
		}
		if topWorks == true {
			alt := dist[u] + int64(matrix[topNode])
			if alt < dist[topNode] {
				dist[topNode] = alt
				previous[topNode] = u
			}
		}
	}

	return previous
}

func main() {

	matrix := readMatrix("82.txt")

	dimx := 80
	dimy := 80

	smallestSumSum := int64(0x7fffffffffffffff)
	smallestStart := -1
	smallestGoal2 := -1

	for start := 0; start < len(matrix); start += dimx {

		fmt.Println(start)

		previous := dijkstra(matrix, dimx, dimy, start)

		smallestSum := int64(0x7fffffffffffffff)
		smallestGoal := -1

		for goal := dimx - 1; goal < len(matrix); goal += dimx {

			sum := int64(0)
			for element := goal; element != -1; element = previous[element] {
				sum += int64(matrix[element])
			}

			if sum < smallestSum {
				smallestSum = sum
				smallestGoal = goal
			}
		}

		if smallestSum < smallestSumSum {
			smallestStart = start
			smallestGoal2 = smallestGoal
			smallestSumSum = smallestSum
		}

	}

	fmt.Print(smallestStart, smallestGoal2, smallestSumSum)
}

