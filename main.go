package main

import "time"

type Coordinate struct {
	X, Y float64
}

type Car interface {
	Start() (bool, error)
	Stop() (bool, error)
	IsRunning() bool
	Drive(Coordinate) error
	CurrentLocation() (Coordinate, error)
}

var _ Car = (*BlueFerrari)(nil)
 
type BlueFerrari struct {
	Manufactured time.Time
}

func (b *BlueFerrari) Start() (bool, error) {
	panic("implement me")
}

func (b *BlueFerrari) Stop() (bool, error) {
	panic("implement me")
}

func (b *BlueFerrari) IsRunning() bool {
	panic("implement me")
}

func (b *BlueFerrari) Drive(coordinate Coordinate) error {
	panic("implement me")
}

func (b *BlueFerrari) CurrentLocation() (Coordinate, error) {
	panic("implement me")
}

func main() {
	myBlueFerrari := BlueFerrari{
		Manufactured: time.Now(),
	}

	log.Printf("%+v", myBlueFerrari)
}
