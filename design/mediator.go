package design

import "fmt"

type Train interface {
	arrive()
	depart()
	permitArrival()
}

type PassengerTrain struct {
	mediator Mediator
}

func (p *PassengerTrain) arrive() {
	if !p.mediator.canArrive(p) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (p *PassengerTrain) depart() {
	fmt.Println("PassengerTrain: Leaving")
	p.mediator.notifyAboutDeparture()
}

func (p *PassengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	p.arrive()
}

type FreightTrain struct {
	mediator Mediator
}

func (g *FreightTrain) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")
}

func (g *FreightTrain) depart() {
	fmt.Println("FreightTrain: Leaving")
	g.mediator.notifyAboutDeparture()
}

func (g *FreightTrain) permitArrival() {
	fmt.Println("FreightTrain: Arrival permitted")
	g.arrive()
}

type Mediator interface {
	canArrive(Train) bool
	notifyAboutDeparture()
}

type StationManager struct {
	isPlatformFree bool
	trainQueue     []Train
}

func newStationManager() *StationManager {
	return &StationManager{
		isPlatformFree: true,
	}
}

func (s *StationManager) canArrive(t Train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return true
}

func (s *StationManager) notifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[:1]
		firstTrainInQueue.permitArrival()
	}
}

func DoMediator() {
	stationManager := newStationManager()
	passengerTrain := &PassengerTrain{mediator: stationManager}
	freightTrain := &FreightTrain{mediator: stationManager}
	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
