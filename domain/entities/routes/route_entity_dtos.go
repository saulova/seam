package routes

type RouteEntityInput struct {
	Id             string
	GatewayPath    string
	Methods        []string
	MiddlewaresIds []string
	Action         string
}
