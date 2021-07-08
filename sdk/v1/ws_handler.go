package v1

import "log"

func (s SDKImpl) OpenConnection() (err error) {
	err = s.ws.Connect(*s.url, nil)
	return err
}

func (s SDKImpl) CloseConnection() (err error) {
	// Stop processing messages
	running = false

	// close connection
	err = s.ws.Close()
	if err != nil {
		log.Println("Can't close connection")
		log.Println(err)
	}
	return err
}

func (s SDKImpl) ResetConnection() (err error) {

	err = s.CloseConnection()
	logError(err)

	err = s.OpenConnection()
	logError(err)

	return err
}
