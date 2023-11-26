package entity

type Enrollment struct {
	ID        string `json:"id" gorm:"primaryKey;type:char(255)"`
	CourseID  string `json:"course_id"`
	StudentID string `json:"student_id"`

	Course  Course
	Student Student
}

type EnrollmentRepository interface {
	GetAll() ([]Enrollment, error)
	GetByID(id string) (*Enrollment, error)
	Create(enrollment *Enrollment) error
	Update(id string, enrollment *Enrollment) error
	Delete(id string) error
}

type EnrollmentUseCase interface {
	GetAll() ([]Enrollment, error)
	GetByID(id string) (*Enrollment, error)
	Create(courseID string, studentID string) (*Enrollment, error)
	Update(id string, enrollment *Enrollment) error
	Delete(id string) error
	Enroll(studentID string, courseID string) error
	Withdraw(studentID string, courseID string) error
}
