package db

type Clip struct {
	userID  int
	content string
}

func (db DB) getClip(userID int) (Clip, error) {
	return Clip{userID: 1, content: ""}, nil
}
