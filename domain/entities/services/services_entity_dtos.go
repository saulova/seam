package services

type ServiceEntityInput struct {
	Id              string
	MiddlewaresIds  []string
	GatewayBasePath string
	RoutesIds       []string
}
