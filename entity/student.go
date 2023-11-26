package entity

type Student struct {
	ID             string `gorm:"primaryKey;type:char(255)"`
	Name           string
	FirstName      string
	LastName       string
	ProgrammeID    string
	DepartmentName string
	GPAX           float64
	MathGPA        float64
	EngGPA         float64
	SciGPA         float64
	School         string
	Year           string
	Admission      string
	Remark         string

	Programme  Programme
	Department Department
}

type StudentRepository interface {
	GetByID(id string) (*Student, error)
	GetAll() ([]Student, error)
	GetByParams(params *Student, limit int, offset int) ([]Student, error)
	Create(student *Student) error
	CreateMany(student []Student) error
	Update(student *Student) error
	Delete(id string) error
}

type StudentUseCase interface {
	GetByID(id string) (*Student, error)
	GetAll() ([]Student, error)
	GetByParams(params *Student, limit int, offset int) ([]Student, error)
	Create(student *Student) error
	CreateMany(student []Student) error
	Update(student *Student) error
}
