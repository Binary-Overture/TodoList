package api

func ErrorResponse(err error) map[string]string {

	if err != nil {
		return map[string]string{
			"error": "Have Error" + err.Error(),
		}
	}
	return map[string]string{
		"wrong": "Here may be something wrong",
	}
}
