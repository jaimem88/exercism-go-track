package robot

import "fmt"

// step 1
const (
	N Dir = iota
	S
	E
	W
)

func Right() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Dir = E
	case S:
		Step1Robot.Dir = W
	case E:
		Step1Robot.Dir = S
	case W:
		Step1Robot.Dir = N
	}
}
func Left() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Dir = W
	case S:
		Step1Robot.Dir = E
	case E:
		Step1Robot.Dir = N
	case W:
		Step1Robot.Dir = S
	}

}
func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y += 1
	case S:
		if Step1Robot.Y > 0 {
			Step1Robot.Y -= 1
		}
	case E:
		Step1Robot.X += 1
	case W:
		if Step1Robot.X > 0 {
			Step1Robot.X -= 1
		}
	}
}

func (d Dir) String() string {
	switch d {
	case N:
		return "N"
	case S:
		return "S"
	case E:
		return "E"
	case W:
		return "W"
	}
	return ""
}

type Action Command

// step 2
func StartRobot(cmd chan Command, act chan Action) {
	for {

		select {
		case c, ok := <-cmd:
			if !ok {
				close(act)
				return

			}
			act <- Action(c)

		}
	}

}
func Room(i int, extent Rect, robot Step2Robot, act chan Action, rep chan Step2Robot) {
	fmt.Printf("RRRRRR start: %d - %s %+v\n", i, robot.Dir, robot.Pos)

	for {
		select {
		case a, ok := <-act:
			if !ok {
				// sending robot
				fmt.Printf("RRRRRR final: %d - %s %+v\n", i, robot.Dir, robot.Pos)
				rep <- robot
				// close rep channel
				close(rep)
				return
			}
			fmt.Printf("the action! %s\n", string(a))
			switch Command(a) {
			case 'A':
				switch robot.Dir {
				case N:
					if robot.Pos.Northing < extent.Max.Northing {
						robot.Pos.Northing++
					}
				case S:
					if robot.Pos.Northing > extent.Max.Northing {
						robot.Pos.Northing--
					}
				case E:
					if robot.Pos.Easting < extent.Max.Easting {
						robot.Pos.Easting++
					}
				case W:
					if robot.Pos.Easting > extent.Max.Easting {
						robot.Pos.Easting--
					}
				}
			case 'R':
				switch robot.Dir {
				case N:
					robot.Dir = E
				case E:
					robot.Dir = S
				case S:
					robot.Dir = W
				case W:
					robot.Dir = N
				}
			case 'L':
				switch robot.Dir {
				case N:
					robot.Dir = W
				case E:
					robot.Dir = N
				case S:
					robot.Dir = E
				case W:
					robot.Dir = S
				}
			}
			fmt.Printf("robot after action: %d - %s %+v\n", i, robot.Dir, robot.Pos)
		}
	}
}
