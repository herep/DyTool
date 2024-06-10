package helloWorld

type HelloWorldService interface {
	GetInfo()string
}

type HelloWorldServiceImpl struct {
	
}

func (h HelloWorldServiceImpl) GetInfo() string {
	panic("implement me")
}

func NewHelloWorldServiceImpl() HelloWorldService {
	return &HelloWorldServiceImpl{}
}
