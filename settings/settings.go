package settings

const (
	// Modes
	PvP  string = "PvP"
	PvAI string = "PvAI"

	// Difficulties
	Easy   string = "easy"
	Normal string = "normal"
)

type Settings struct {
	Modes              [2]string
	Difficulties       [2]string
	SelectedMode       string
	SelectedDifficulty string
}

func NewSettings() *Settings {
	return &Settings{
		Modes:              [2]string{PvP, PvAI},
		Difficulties:       [2]string{Easy, Normal},
		SelectedMode:       PvAI,
		SelectedDifficulty: Easy,
	}
}

func (s *Settings) GetSelectedMode() string {
	return s.SelectedMode
}

func (s *Settings) SetSelectedMode(m string) {
	s.SelectedMode = m
}

func (s *Settings) GetSelectedDifficulty() string {
	return s.SelectedDifficulty
}

func (s *Settings) SetSelectedDifficulty(d string) {
	s.SelectedDifficulty = d
}

func (s *Settings) GetModes() [2]string {
	return s.Modes
}

func (s *Settings) GetDifficulties() [2]string {
	return s.Difficulties
}
