package timestamp

import "time"

func GetTimestamp() (string, error) {

	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return "", err
	}

	currentTime := time.Now().In(location)
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	return formattedTime, nil
}
