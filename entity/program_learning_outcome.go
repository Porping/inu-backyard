package entity

type ProgramLearningOutcome struct {
	ID              string `json:"id" gorm:"primaryKey;type:char(255)"`
	Code            string `json:"code"`
	DescriptionThai string `json:"descriptionThai"`
	DescriptionEng  string `json:"descriptionEng"`
	ProgramYear     int    `json:"programYear"`
}

type ProgramLearningOutcomeRepository interface {
	GetAll() ([]ProgramLearningOutcome, error)
	GetByID(id string) (*ProgramLearningOutcome, error)
	Create(programLearningOutcome *ProgramLearningOutcome) error
	Update(programLearningOutcome *ProgramLearningOutcome) error
	Delete(id string) error
}

type ProgramLearningOutcomeUsecase interface {
	GetAll() ([]ProgramLearningOutcome, error)
	GetByID(id string) (*ProgramLearningOutcome, error)
	Create(code string, descriptionThai string, descriptionEng string, programYear int) (*ProgramLearningOutcome, error)
	Delete(id string) error
}
