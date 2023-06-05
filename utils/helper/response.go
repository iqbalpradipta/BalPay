package helper

func FailedResponseHelper(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status": "Failed",
		"message": msg,
	}
}

func SuccessResponseHelper(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
		"message": msg,
	}
}

func SuccessResponseDataHelper(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
		"message": msg,
		"data": data,
	}
}