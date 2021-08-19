package main

type Server struct {
	*Logger
	Name string
}

func NewServer(name string, logger *Logger) Server {
	return Server{
		logger,
		name,
	}
}

func (sr *Server) Start() {
	sr.Info("Starting server " + sr.Name)
}

func (sr *Server) GetInfo() string {
	return "this is all the info I have for now"
}
