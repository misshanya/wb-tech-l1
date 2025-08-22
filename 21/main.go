package main

import "fmt"

// В примере GeometricSquare использует интерфейс NumberOperations
// Структура lowLevelNumberOperations не предоставляет метода Square
// Но она предоставляет метод Multiple, который мы обернули в методе Square структуры numberOperations

// Паттерн применяется, когда надо заставить работать объект с несовместимым интерфейсом
// Также он широко применяется для декаплинга разных частей проекта
// Он позволяет отделить бизнес-логику от транспортного слоя, что дает больше гибкости

// Плюсы:
// 1. SRP. У каждого своя зона ответственности
// 2. Переиспользование
// 3. Клиент не знает, что на самом деле может происходить под капотом.
// 	  Это позволяет менять логику, не трогая вызов

// Из минусов можно выделить разве что усложнение кода

func main() {
	lowLeveNumOps := lowLevelNumberOperations{}
	numOps := numberOperations{lowLevel: lowLeveNumOps}

	sq := NewGeometricSquare(5, &numOps)

	fmt.Println("Area of square is", sq.Area())
}

type GeometricSquare struct {
	sideLen int
	math    NumberOperations
}

func NewGeometricSquare(sideLen int, math NumberOperations) *GeometricSquare {
	return &GeometricSquare{
		sideLen: sideLen,
		math:    math,
	}
}

func (s *GeometricSquare) Area() int {
	return s.math.Square(s.sideLen)
}

type NumberOperations interface {
	Square(number int) int
}

type numberOperations struct {
	lowLevel lowLevelNumberOperations
}

func (o *numberOperations) Square(number int) int {
	return o.lowLevel.Multiple(number, number)
}

type LowLevelNumberOperations interface {
	Multiple(num1, num2 int) int
}

type lowLevelNumberOperations struct{}

func (o *lowLevelNumberOperations) Multiple(num1, num2 int) int {
	return num1 * num2
}
