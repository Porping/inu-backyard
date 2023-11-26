package repository

import (
	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type SemesterRepository struct {
	gorm *gorm.DB
}

func NewSemesterRepositoryGorm(gorm *gorm.DB) entity.SemesterRepository {
	return &SemesterRepository{gorm: gorm}
}

func (r *SemesterRepository) GetAll() ([]entity.Semester, error) {
	var semesters []entity.Semester
	if err := r.gorm.Find(&semesters).Error; err != nil {
		return nil, err
	}
	return semesters, nil
}

func (r *SemesterRepository) GetByID(id string) (*entity.Semester, error) {
	var semester entity.Semester
	if err := r.gorm.First(&semester, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &semester, nil
}

func (r *SemesterRepository) Create(semester *entity.Semester) error {
	if err := r.gorm.Create(semester).Error; err != nil {
		return err
	}
	return nil
}

func (r *SemesterRepository) Update(semester *entity.Semester) error {
	if err := r.gorm.Save(semester).Error; err != nil {
		return err
	}
	return nil
}

func (r *SemesterRepository) Delete(id string) error {
	if err := r.gorm.Delete(&entity.Semester{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
