package main

import (
	"fmt"
)

type Estado struct{
	nome string 
	area float64
}

func main() {
	fmt.Println(os10maioresEstadosDoBrasil())
}

func todosEstados() ([]Estado){
	return ([]Estado{
		
		{"Minas Gerais",	587.5},
		{"Bahia", 564.6},
		{"Mato Grosso do Sul",	357.1},
		{"Goiás",	340.0},
		{"Maranhão",	332.0},
		{"Rio Grande do Sul",	281.7},
		{"Tocantins",	277.6},
		{"Piauí",	251.5},
		{"São Paulo",	248.2},
		{"Rondônia",	237.5},
		{"Roraima",	224.2},
		{"Paraná",	199.3},
		{"Acre",	152.5},
		{"Ceará",	148.8},
		{"Amapá",	142.8},
		{"Pernambuco",	98.3},
		{"Santa Catarina",	95.3},
		{"Paraíba",	56.4},
		{"Rio Grande do Norte",	52.8},
		{"Espírito Santo",	46.0},
		{"Amazonas",1570.7},
		{"Pará",	1247.6},
		{"Mato Grosso",	903.3},
		{"Rio de Janeiro",	43.6},
		{"Alagoas",	27.7},
		{"Sergipe",	21.9},
		{"Distrito Federal", 5.8}	})
}

func HeapSort(array []Estado) (){
	BuildHeap(array)

	for length:= len(array); length > 1; length-- {
		RemoveTop(array, length)
	}

	
}
func BuildHeap(array[]Estado)(){
	for i := len(array) / 2; i >= 0; i-- {
		Heapify(array, i, len(array))
	}
} 	
func RemoveTop(array []Estado, length int)(){
	var lastIndex = length - 1
	array[0], array[lastIndex] = array[lastIndex], array[0]
	Heapify(array, 0, lastIndex)
}
func Heapify(array[]Estado, root, length int)(){
	var max = root
	var l, r = Left(root), Right( root)

	if l < length && array[l].area > array[max].area {
		max = l
	}

	if r < length && array[r].area > array[max].area {
		max = r
	}

	if max != root {
		array[root], array[max] = array[max], array[root]
		Heapify(array, max, length)
	}
}

func Left(root int) int {
	return (root * 2) + 1
}

func Right(root int) int {
	return (root * 2) + 2
}




func os10maioresEstadosDoBrasil() ([10]string, error) {
	var data [10]string
	estados:=todosEstados()
	HeapSort(estados)
	//os10estados := make(v:=todosEstados(), len(v)) 
		
	for j , i := 0 ,26; i >= 17; i--{
		data[j] = estados[i].nome
		j++
	}
	

	return data, nil
}
