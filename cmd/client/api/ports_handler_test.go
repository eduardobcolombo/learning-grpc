package api

import (
	"context"
	"crypto/tls"
	"net/http"
	"testing"

	"github.com/eduardobcolombo/learning-grpc/internal/pkg/port"
	"github.com/eduardobcolombo/learning-grpc/internal/pkg/portpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func TestPortsIndexOPTIONS(t *testing.T) {
	t.Run("Test if the verb OPTIONS endpoint is working", func(t *testing.T) {
		rq := ReqTest{Verb: "OPTIONS", URL: "/v1/ports"}
		res := MakeRequest(rq)

		got := res.Code
		want := http.StatusOK
		assertCorrectMessage(t, got, want)
	})
}

// func TestFillPortpbWithJSON(t *testing.T) {
// 	t.Run("Test if fill Portpb with JSON is filling on the right way", func(t *testing.T) {
// 		fileName := testFileName
// 		f, err := os.Open(fileName)
// 		assertNil(t, err)
// 		defer f.Close()
// 		r := bufio.NewReader(f)
// 		dec := json.NewDecoder(r)

// 		for dec.More() {
// 			var m map[string]interface{}
// 			err := dec.Decode(&m)
// 			assertNil(t, err)

// 			for _, v := range m {
// 				jsonbody, err := json.Marshal(v)
// 				assertNil(t, err)

// 				filledPort, err := port.fillPortpbWithJSON(jsonbody)
// 				assertNil(t, err)

// 				p := filledPort.GetPort()
// 				assertCorrectMessage(t, p.GetName(), "Ajman")
// 				assertCorrectMessage(t, p.GetCity(), "Ajman")
// 				assertCorrectMessage(t, p.GetCountry(), "United Arab Emirates")
// 				assertCorrectMessage(t, p.GetCoordinates().GetLat(), float32(55.5136433))
// 				assertCorrectMessage(t, p.GetCoordinates().GetLong(), float32(25.4052165))
// 				assertCorrectMessage(t, p.GetProvince(), "Ajman")
// 				assertCorrectMessage(t, p.GetTimezone(), "Asia/Dubai")
// 				assertCorrectMessage(t, p.GetUnlocs().GetUnloc()[0], "AEAJM")
// 				assertCorrectMessage(t, p.GetCode(), "52000")

// 				return
// 			}
// 		}
// 	})
// }

func TestClientPortUpdate(t *testing.T) {

	tests := []struct {
		name     string
		file     string
		expected interface{}
	}{
		{
			"Request update with an valid JSON with 1 record",
			testFileName,
			"Received 1 records.",
		},
		{
			"Request update with an invalid JSON ",
			testFileName2Records,
			"Received 2 records.",
		},
	}

	ctx := context.Background()

	opts := []grpc.DialOption{
		// grpc.WithBlock(),
	}
	opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})))
	opts = append(opts, grpc.WithContextDialer(dialer()))

	cc, err := grpc.DialContext(ctx, "", opts...)
	assertNil(t, err)
	defer cc.Close()
	cfg.psc = portpb.NewPortServiceClient(cc)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := core.Port.UpdatePortsOnServer(tt.file)
			assertNil(t, err)
			assertCorrectMessage(t, response, tt.expected)
		})
	}
}
func TestClientRetrievePorts(t *testing.T) {
	tests := []struct {
		name     string
		file     string
		expected string
	}{
		{
			"Retrieve records from server",
			testFileName,
			"Ajman",
		},
	}

	ctx := context.Background()

	opts := []grpc.DialOption{
		grpc.WithBlock(),
	}
	opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})))

	cc, err := grpc.DialContext(ctx, "", opts...)
	assertNil(t, err)
	defer cc.Close()
	cfg.psc = portpb.NewPortServiceClient(cc)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := port.Core.RetrievePortsFromServer(core.Port)
			assertNil(t, err)
			for _, res := range response {
				assertCorrectMessage(t, res.GetName(), tt.expected)
				assertCorrectMessage(t, res.GetCity(), tt.expected)
			}
		})
	}
}
