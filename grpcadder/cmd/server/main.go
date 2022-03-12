package main

func main(){
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	api.RegisterAdderServer(s, srv)

	l, err := net.listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Server(l); err != nil{
		log.Fatal(err)
	}
}