package shared

type DetailedError struct {
	Message string `json:"error"`

	Details struct {
		Stack string `json:"stack"`
	} `json:"details"`

	Url string `json:"url"`
}

func (e *DetailedError) Error() string {
	return e.Message
}
