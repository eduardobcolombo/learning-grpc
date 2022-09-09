package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/eduardobcolombo/learning-grpc/internal/pkg/foundation"
	"github.com/eduardobcolombo/learning-grpc/internal/pkg/port"
	"go.uber.org/zap"
)

var testFileName = "../../../test/data/ports_test.json"
var testFileName2Records = "../../../test/data/ports_test_2_records.json"
var cfg = GetEnvTest()

// Construct the application logger.
var testLog = func() *zap.SugaredLogger {
	log, err := foundation.NewLogger("GW-API")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()
	return log
}()
var core = CoreConfig{
	Port: port.NewCore(testLog, cfg.psc),
}
var rt = ServerWithMiddlewares(core, cfg)
var assertCorrectMessage = func(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %s[%T] want %s[%T]", got, got, want, want)
	}
}
var assertNil = func(t *testing.T, got interface{}) {
	t.Helper()
	if got != nil {
		t.Errorf("got %q want nil", got)
	}
}

type ReqTest struct {
	Verb string
	URL  string
	Body string
}

func GetEnvTest() *Config {
	testing.Init()
	cfg := &Config{}
	return cfg
}

func ServerWithMiddlewares(core CoreConfig, cfg *Config) http.Handler {
	api := newAPI()
	api.routes(core, cfg)
	api.mid(cfg)

	return api.router
}

func MakeRequest(r ReqTest) (res *httptest.ResponseRecorder) {
	b := strings.NewReader(r.Body)
	req, _ := http.NewRequest(r.Verb, r.URL, b)
	if r.Verb == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	res = httptest.NewRecorder()
	rt.ServeHTTP(res, req)

	return res
}
