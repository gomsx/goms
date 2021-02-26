package model

// MakePongMsg make pong msg.
func MakePongMsg(s string) string {
	if s == "" {
		s = "NONE!"
	}
	return "pong" + " " + s
}
