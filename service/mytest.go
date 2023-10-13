package service

import "go-echo-stockport/repository"

type serviceTest struct {
	repo repository.RepositoryInterfaceTest
}

type UserDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewServiceTest(repo repository.RepositoryInterfaceTest) *serviceTest {

	serviceTest := serviceTest{
		repo: repo,
	}
	return &serviceTest
}

func (s *serviceTest) GetHello() string {
	return s.repo.GetHello()
}

func (s *serviceTest) UpdateHello(text string) {
	s.repo.UpdateHello(text)
}

func (s *serviceTest) GetAllUser() []UserDto {
	repousr := s.repo.GetAllUser()
	ret := make([]UserDto, len(repousr))
	for i, v := range repousr {
		userdto := UserDto{
			Name: v.Name,
		}
		ret[i] = userdto
	}
	return ret
}

func (s *serviceTest) AddUser(udto *UserDto) {
	repoUser := repository.User{
		Name:  udto.Name,
		Email: udto.Email,
	}
	s.repo.AddUser(repoUser)
}
