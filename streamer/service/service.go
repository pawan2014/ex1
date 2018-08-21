package service

import (
	"github.com/ex1/streamer/dao"
	"github.com/ex1/streamer/model"
)

type Service struct {
	dao dao.Dao
}

func (s *Service) GetAllMachines() []model.Machine {

	return s.dao.GetMachinesAndTags()
}
func NewService(dao dao.Dao) *Service {
	Myservice = Service{dao}
	return &Myservice
}

var Myservice Service
