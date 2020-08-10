package model

//
type Ping struct {
	Type  string
	Count int64
}

//
func MakePongMsg(s string) string {
	if s == "" {
		s = "NONE!"
	}
	return "pong" + " " + s
}
