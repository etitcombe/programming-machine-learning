package main

import (
	"fmt"
	"log"

	"github.com/etitcombe/programming-machine-learning/data"
	"gonum.org/v1/gonum/mat"
)

func main() {
	run()
}

func run() error {
	const cols = 4

	records, err := data.ReadSeparatedValuesFloat64("./input/more_pizza.txt", ' ', cols, true)
	if err != nil {
		log.Fatal(err)
	}

	mx := mat.NewDense(len(records), cols-1, nil)
	my := mat.NewDense(len(records), 1, nil)

	for j, record := range records {
		for i := range record {
			if i < len(record)-1 {
				mx.Set(j, i, record[i])
			} else {
				my.Set(j, 0, record[i])
			}
		}
	}

	fmt.Println(mx.Dims())
	fmt.Printf("%+v\n", mx)
	fmt.Println(my.Dims())
	fmt.Printf("%+v\n", my)

	w := mat.NewDense(len(records), 1, nil)
	fmt.Println(w)

	return nil
}
