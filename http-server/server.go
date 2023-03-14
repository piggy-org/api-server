package http_server

type Server interface {
	Run() error
	Stop() error
}
