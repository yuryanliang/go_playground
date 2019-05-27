package controller

func init() {
	s.grpcService = services.GRPCServices
}
func (c *controllers) SayHelloWorld() (response string) {
	res, err := s.grpcService.HelloService.GreeterClient.SayHello(nil, nil)
	if res.word == "Hi" {
		response = "Hi there! how are you doing?"
	} else if res.word == "Bye" {
		response = "Bye, see you again soon."
	}
	return response
}
