package models

type (
	APIInformation struct {
		API       APIInfo   `json:"api"`
		Status    APIStatus `json:"status"`
		Timestamp string    `json:"timestamp"`
	}

	APIInfo struct {
		Commit  string `json:"commit"`
		Release string `json:"release"`
		Version int    `json:"version"`
	}

	APIStatus struct {
		Error    int    `json:"error"`
		HTTPCode int    `json:"http_code"`
		Message  string `json:"message"`
	}
)
