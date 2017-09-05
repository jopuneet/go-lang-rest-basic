package util

import mgo "gopkg.in/mgo.v2"

//GetSession creates a session with mongodb
func GetSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}
