package models

type About struct {
	SlackUsername string `json:"slackUsername"`
	Backend       bool   `json:"backend"`
	Age           int    `json:"age"`
	Bio           string `json:"bio"`
}

func NewBioProfile() *About {
	return &About{}
}

func (a *About) CreateMyBio() {
	a.SlackUsername = "Pelumi"
	a.Backend = true
	a.Age = 23
	a.Bio = "Hello everyone, my name is Pelumi and I'm trying to get better using golang"
}
