package v1

import "log"

func (s SDKImpl) OpenConnection() (err error) {
	return nil
}

func (s SDKImpl) CloseConnection() (err error) {
	mtd := "CloseConnection: "

	// Stop processing messages
	running = false

	// close WS channel if its not yet fully closed!
	if stopC != nil {
		close(stopC)
	}

	if con != nil {
		// close connection
		err = con.Close()
		if err != nil {
			log.Println(mtd + "Can't close connection")
			log.Println(err)
			return err
		}
	}
	return nil
}

func (s SDKImpl) ResetConnection() (err error) {

	err = s.CloseConnection()
	logError(err)

	err = s.OpenConnection()
	logError(err)

	return err
}
