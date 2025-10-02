package db

type Clip struct {
	userID  int
	content string
}

func (db DB) getClip(userID int) (Clip, error) {
	// dummy value
	return Clip{userID: 1, content: ""}, nil
}
