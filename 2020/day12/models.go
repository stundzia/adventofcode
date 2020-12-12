package day12


type Ship struct {
	Heading int
	Coords [2]int
}

type ShipNav struct {
	Ship *Ship
	Waypoint *Ship // Let's pretend the waypoint is an imaginary ship
}


func NewShip(startingHeading int, startingCoords [2]int) *Ship {
	ship := &Ship{
		Heading:     startingHeading,
		Coords: startingCoords,
	}
	return ship
}


func deltaFromHeading(heading int) (delta [2]int) {
	mod := (heading / 90) % 4
	switch mod {
	case 0:
		delta[1] = 1
		break
	case 1:
		delta[0] = 1
		break
	case 2:
		delta[1] = -1
		break
	case 3:
		delta[0] = -1
		break
	default:
		panic("weird heading mod")
	}
	return delta
}

func deltaFromDirection(direction string) (delta [2]int) {
	switch direction {
	case "N":
		delta[1] = 1
		break
	case "E":
		delta[0] = 1
		break
	case "S":
		delta[1] = -1
		break
	case "W":
		delta[0] = -1
		break
	default:
		panic("weird direction")
	}
	return delta
}

func (s *Ship) GetDeltaFromHeading() (delta [2]int) {
	return deltaFromHeading(s.Heading)
}

func (s *Ship) RotateAroundCenter(degrees int) {
	if degrees > 0 {
		switch degrees / 90 % 4 {
		case 0:
			return
		case 1:

		}
	}
}

func (s *Ship) HandleNavCommand(direction string, value int)  {
	switch direction {
	case "L":
		s.Heading -= value
		return
	case "R":
		s.Heading += value
	case "F":
		d := s.GetDeltaFromHeading()
		for i, delta := range d {
			s.Coords[i] += delta * value
		}
	default:
		d := deltaFromDirection(direction)
		for i, delta := range d {
			s.Coords[i] += delta * value
		}
	}
	if s.Heading < 0 {
		s.Heading += 360
	}
}

func (sn *ShipNav) HandleNavCommand(direction string, value int)  {
	switch direction {
	case "L":
		rot := value / 90 % 2
		for i := 0; i < rot; i++ {
			//sn.Waypoint
		}

	}
}


func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func (s *Ship) getManhattenDistance() int {
	return abs(s.Coords[0]) + abs(s.Coords[1])
}