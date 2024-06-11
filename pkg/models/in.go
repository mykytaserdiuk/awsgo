package models

type TodoIn struct {
	Topic       string `json:"topic"`
	Description string `json:"description"`
}

func (t *TodoIn) Validate() error {

	if t.Topic == "" {
		return ErrorUnvalidTopic
	}
	if t.Description == "" {
		return ErrorUnvalidDescription
	}

	return nil
}
