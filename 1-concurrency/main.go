// Нужно создать приложение с 2-мя горутинами, где:

//     Первая создаёт slice из 10 случайных элементов от 0 до 100 и передаёт их по одному во вторую горутину.
//     Вторая получает числа от 1-й и возводит в квадрат передавая результат в main.
//     В main дожидаемся всех 10 чисел, которые были возведены в квадрат и выводим их в консоль.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch := make(chan int)
	result := make(chan int)

	go generateSlice(ch)
	go powWorker(ch, result)

	for v := range result {
		fmt.Println("Результат возведения в квадрат:", v)
	}
}

func generateSlice(ch chan<- int) {
	defer close(ch)
	size := 10
	slice := make([]int, size)

	for i := range slice {
		random := rand.Intn(101)
		slice[i] = random
	}

	for _, v := range slice {
		ch <- v
	}
}

func powWorker(ch <-chan int, result chan<- int) {
	defer close(result)
	for v := range ch {
		result <- v * v
	}
}
