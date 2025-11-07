package exam

type Server struct {
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
	Addr         string
}
type Option func(s *Server)

func NewServer(options ...Option) *Server {
	svr := &Server{}
	for _, o := range options {
		o(svr)
	}

	return svr
}
