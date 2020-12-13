package main

import (
	"fmt"
	"github.com/maxheckel/advent2020/common"
	"math"
	"strconv"
)

type Ship struct {
	longitude int
	latitude int
	heading string
}

func (s *Ship) applyAction(action string, amount int, asWaypoint bool) {

	switch action {
	case left:
		if asWaypoint {
			switch amount {
			case 90:
				currentLat := s.latitude
				currentLon := s.longitude
				s.longitude = currentLat*-1
				s.latitude = currentLon
			case 180:
				s.latitude*=-1
				s.longitude*=-1
			case 270:
				currentLat := s.latitude
				currentLon := s.longitude
				s.longitude = currentLat
				s.latitude = currentLon*-1
			}
		} else {
			s.heading = degreesToHeading((amount+headingToDegrees(s.heading))%360)
		}

	case right:
		if asWaypoint {
			switch amount {
			case 90:
				currentLat := s.latitude
				currentLon := s.longitude
				s.longitude = currentLat
				s.latitude = currentLon*-1
			case 180:
				s.latitude*=-1
				s.longitude*=-1
			case 270:
				currentLat := s.latitude
				currentLon := s.longitude
				s.longitude = currentLat*-1
				s.latitude = currentLon
			}
		} else {
			newHeading := (headingToDegrees(s.heading)-amount)%360
			if newHeading < 0 {
				newHeading = 360 + newHeading
			}
			s.heading = degreesToHeading(newHeading)
		}

	case forward:
		s.moveShip(s.heading, amount)
	default:
		s.moveShip(action, amount)
	}
}

func (s *Ship) moveToOtherShip(waypoint Ship, amount int){
	s.latitude += waypoint.latitude*amount
	s.longitude += waypoint.longitude*amount
}

func (s *Ship) moveShip(direction string, amount int) {
	switch direction {
	case north:
		s.latitude += amount
	case south:
		s.latitude -= amount
	case east:
		s.longitude += amount
	case west:
		s.longitude -= amount
	}
}

const (
	east = "E"
	north = "N"
	south = "S"
	west = "W"
	left = "L"
	right = "R"
	forward = "F"
)
func main() {
	directions := common.StringListFromFile("./12/input.txt")

	ship := Ship{
		longitude: 0,
		latitude:  0,
		heading:   east,
	}
	waypoint := Ship{
		longitude: 10,
		latitude:  1,
		heading:   "",
	}
	for _, direction := range directions {
		action := direction[:1]
		amount, _ := strconv.Atoi(direction[1:])
		ship.applyAction(action, amount, false)
	}
	sum := math.Abs(float64(ship.longitude)) + math.Abs(float64(ship.latitude))
	ship.longitude = 0
	ship.latitude = 0

	fmt.Printf("Part 1: %f\n", sum)
	for _, direction := range directions {
		action := direction[:1]
		amount, _ := strconv.Atoi(direction[1:])
		if action != forward {

			waypoint.applyAction(action, amount, true)
		} else {
			ship.moveToOtherShip(waypoint, amount)
		}

		fmt.Println("waypoint", waypoint)
		fmt.Println("ship", ship)
	}
	sum = math.Abs(float64(ship.longitude)) + math.Abs(float64(ship.latitude))
	fmt.Printf("Part 2: %f", sum)
}

func degreesToHeading(degrees int) string {
	if degrees < 0 {

	}
	switch degrees {
	case 0:
		return east
	case 90:
		return north
	case 180:
		return west
	case 270:
		return south
	}
	return ""
}

func headingToDegrees(heading string) int {
	switch heading {
	case east:
		return 0
	case north:
		return 90
	case west:
		return 180
	case south:
		return 270
	}
	return 0
}