package robot

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
	switch Step1Robot.Dir {
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
