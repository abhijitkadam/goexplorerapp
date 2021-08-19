package main

import "fmt"

func main() {
	logger := NewLogger()
	server := NewServer("Info server", &logger)
	server.Start()
	ServerOperations(&server)
	server.Logger.Error("Recording a error")
	server.DrainToConsole()

}

func ServerOperations(server *Server) string {

	server.Info("getting data")
	info := server.GetInfo()
	server.Debug("Got the data" + info)
	return info
}

func LoggerUsage() {
	logger := NewLogger()

	logger.Info("Inside main")
	logger.Info("Invoking task 1")
	logger.Error("Error in executing task")

	logger.DrainToConsole()
}

func RepoDemoForInterfaces() {
	fRepo := NewHotelRepoFile("data/golddata.json")

	HotelsRepoUsage(&fRepo)

	apiRepo := NewApiHotelRepo("http://localhost:3000/inventory")

	HotelsRepoUsage(&apiRepo)
}

func HotelsRepoUsage(repo HotelRepo) {
	hData := repo.GetAllHotels()
	fmt.Println(hData)
}

func CounterDemoForInterfaces() {
	sc := NewSimpleCounter()
	CounterOperations(&sc)

	stCounter := NewStepCounter(5, 0)
	CounterOperations(&stCounter)
}

//CounterOperations works on abstraction i.e. the Counter interface
func CounterOperations(counter Counter) {
	counter.Increment()
	counter.Increment()
	fmt.Println(counter.GetValue())
}
