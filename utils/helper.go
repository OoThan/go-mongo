package utils

import "go-mongo/logger"

func GlobalErrorException(message error) bool {
	if message != nil {
		logger.Sugar.Errorf("Caused error exceptions : %s", message.Error())
	}

	return message == nil
}

func GlobalErrorDatabaseException(message error) bool {
	if message != nil {
		logger.Sugar.Error("Caused database error exception : ", message.Error())
	}

	return message == nil
}
