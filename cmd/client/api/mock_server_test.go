package api

// type mockServer struct {
// 	portpb.UnimplementedPortServiceServer
// }

// func (*mockServer) PortsUpdate(stream portpb.PortService_PortsUpdateServer) error {
// 	count := 0
// 	for {
// 		_, err := stream.Recv()
// 		if err == io.EOF {
// 			return stream.SendAndClose(&portpb.PortResponse{
// 				Result: fmt.Sprintf("Received %d records.", count),
// 			})
// 		}
// 		count++
// 	}
// }

// func (s *mockServer) PortsList(req *portpb.ListPortsRequest, stream portpb.PortService_PortsListServer) error {

// 	lPorts := []*portpb.Port{}
// 	lPorts = append(lPorts, &portpb.Port{
// 		Id:          1,
// 		Name:        "Ajman",
// 		City:        "Ajman",
// 		Country:     "United Arab Emirates",
// 		Alias:       []string{},
// 		Regions:     []string{},
// 		Coordinates: &portpb.Coordinates{Lat: 55.5136433, Long: 25.4052165},
// 		Province:    "Ajman",
// 		Timezone:    "Asia/Dubai",
// 		Unlocs:      &portpb.Unlocs{Unloc: []string{"AEAJM"}},
// 		Code:        "52000",
// 	})
// 	for _, p := range lPorts {
// 		err := stream.Send(&portpb.ListPortsResponse{Port: p})
// 		if err != nil {
// 			log.Printf("Error while reading client stream loop: %v", err)
// 			return nil
// 		}
// 	}

// 	return nil
// }

// func dialer() func(context.Context, string) (net.Conn, error) {
// 	list := bufconn.Listen(1024 * 1024)
// 	server := grpc.NewServer()
// 	portpb.RegisterPortServiceServer(server, &mockServer{})

// 	go func() {
// 		if err := server.Serve(list); err != nil {
// 			log.Fatal(err)
// 		}
// 	}()

// 	return func(context.Context, string) (net.Conn, error) {
// 		return list.Dial()
// 	}
// }
