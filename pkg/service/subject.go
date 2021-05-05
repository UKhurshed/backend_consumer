package service

import (
	"backend_consumer/pkg/domain"
	"backend_consumer/pkg/repository"
)

type SubjectService struct {
	repo repository.Subject
}

func NewSubjectService(repo repository.Subject) *SubjectService{
	return &SubjectService{repo: repo}
}

func (s *SubjectService) GetAllSubjects() ([]domain.Subject, error){
	return s.repo.GetAllSubjects()
}

func (s *SubjectService) CreateSubject(subject domain.Subject) (int, error){
	return s.repo.CreateSubject(subject)
}

func (s *SubjectService) UpdateSubject(subjectId int, subject domain.SubjectInput) error{
	return s.repo.UpdateSubject(subjectId, subject)
}

func (s* SubjectService) DeleteSubject(subjectId int) error{
	return s.repo.DeleteSubject(subjectId)
}



