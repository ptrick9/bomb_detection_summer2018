package main
import "fmt"
import (
	"math/rand"
	"os"
	"log"
	"bytes"
)

//var grid[][]Node
//var nodeList[]NodeParent
var maxX int
var maxY int

var xDiv int
var yDiv int

const squareRow = 20
const squareCol = 20

const numNodes = 100

//var grid = [][]*Square{}

func rangeInt(min, max int) int {
	return rand.Intn(max-min) + min
}




func printGrid(g [][]*Square) bytes.Buffer{
	var buffer bytes.Buffer
	for i, _ := range g {
		for _, x := range g[i] {
			buffer.WriteString(fmt.Sprintf("%.2f\t", x.avg))
			//fmt.Printf("%.2f\t", x.avg)
		}
		buffer.WriteString(fmt.Sprintf("\n"))
		//fmt.Println("")
	}
	return buffer
}


func main() {

	file, err := os.Create("result.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()



	maxX = 100
	maxY = 100

	xDiv = maxX/squareCol
	yDiv = maxY/squareRow

	nodeList := make([]NodeMovement, numNodes)
	b := &bomb{x:5, y:5}
	for i := 0; i < numNodes; i++ {
		if i%3 == 0 {
			nodeList[i] = &bn{NodeImpl:&NodeImpl{x: rand.Intn(maxX), y: rand.Intn(maxY), id: i}, x_speed: rangeInt(-3, 3), y_speed: rangeInt(-3, 3)}
			//fmt.Print("added bn")

		} else if i%3 == 1 {
			nodeList[i] = &wn{NodeImpl:&NodeImpl{x: rand.Intn(maxX), y: rand.Intn(maxY), id: i}, speed: rangeInt(-3, 3), dir: rand.Intn(2)}
		} else {
			nodeList[i] = &rn{NodeImpl:&NodeImpl{x: rand.Intn(maxX), y: rand.Intn(maxY), id: i}}
		}
	}

	grid := make([][]*Square, squareRow)
	for i := range grid {
		grid[i] = make([]*Square, squareCol)
	}
	for i := 0; i < squareRow; i++ {
		for j := 0; j < squareCol; j++ {
			grid[i][j] = &Square{0.0, 0, []float32{0.0, 0.0, 0.0, 0.0}, 4, 0.0}
		}
	}


	for i := 0; i < 1000; i++ {
		for _, v := range nodeList {
			v.move()
			fmt.Fprintln(file, v)

			//fmt.Println(v.distance(*b))

			//fmt.Println(v)
			//fmt.Println(v.distance(*b))
			grid[v.row(xDiv)][v.col(yDiv)].takeMeasurement(v.distance(*b))

		}
		/*fmt.Println("")
		printGrid(grid)
		fmt.Println("-------------")*/
		x := printGrid(grid)
		//fmt.Fprintf(file, "-----------\n")
		fmt.Fprintln(file, x.String())
		fmt.Fprintf(file,"----------------\n")
	}

}
