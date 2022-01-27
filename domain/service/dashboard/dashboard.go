package dashboardservice

import "github.com/AndreySpies/doccer/domain/contract"

type dashboardService struct {
	userRepo contract.UserRepo
}

func NewDashboardService(repo contract.UserRepo) contract.DashboardService {
	return &dashboardService{
		userRepo: repo,
	}
}
