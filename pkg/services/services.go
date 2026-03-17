package services

import (
	"database/sql"
	"fmt"

	"github.com/EnterpriseGradeSystem/pkg/models"
	"github.com/EnterpriseGradeSystem/pkg/adapters"
)

func NewService() Service {
	return &service{adapter: &adapters.SqlAdapter{}}
}

type Service interface {
	GetUsers() ([]models.User, error)
}

type service struct {
	adapter adapters.Adapter
}

func (s *service) GetUsers() ([]models.User, error) {
	return s.adapter.GetUsers()
}