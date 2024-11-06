package application

/**
* Here business logic is implemented
**/

type HelloService struct {
}

func (s *HelloService) GenerateHello(name string) string {
	return "Hello " + name
}
