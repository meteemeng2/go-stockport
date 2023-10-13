package repository

type repositoryTest struct {
	hellotext string
	userlist  []User
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewRepositoryTest(options ...string) *repositoryTest {
	text := "hello default" // Default value

	if len(options) > 0 {
		text = options[0]
	}
	r := repositoryTest{
		hellotext: text,
	}
	return &r
}

func (r *repositoryTest) GetHello() string {
	return r.hellotext
}

func (r *repositoryTest) UpdateHello(text string) {
	r.hellotext = text
}

func (r *repositoryTest) GetAllUser() []User {
	return r.userlist
}

func (r *repositoryTest) AddUser(u User) {
	r.userlist = append(r.userlist, u)
}
