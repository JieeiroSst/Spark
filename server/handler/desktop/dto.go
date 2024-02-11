package desktop

import "Spark/utils/melody"

type desktop struct {
	uuid       string
	device     string
	srcConn    *melody.Session
	deviceConn *melody.Session
}

var desktopSessions = melody.New()
