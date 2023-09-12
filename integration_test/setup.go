package integrationtest

import (
	"fmt"

	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gaogao-asia/golang-template/config"
	"github.com/gaogao-asia/golang-template/internal/server"
	"github.com/gaogao-asia/golang-template/internal/server/http/response"
	"github.com/gaogao-asia/golang-template/pkg/connection"
	"github.com/gaogao-asia/golang-template/pkg/testutil"
	"github.com/gaogao-asia/golang-template/pkg/utils"
	"github.com/gin-gonic/gin"
)

var (
	Conn    connection.Conn
	BaseURL string
	Engine  *gin.Engine
)

func Setup() {
	testutil.InitConfigForIntegrationTest()

	BaseURL = fmt.Sprintf("http://localhost:%s", config.AppConfig.Server.Port)

	var err error
	Conn, err = connection.GetConnection()
	if err != nil {
		panic(err)
	}

	Engine = SetupTestServer(Conn)
}

func ParseResponse(resp *http.Response) (statusCode int, respData *response.ResponseBody, err error) {
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return -1, nil, err
	}

	res := new(response.ResponseBody)
	err = utils.JsonUnmarshal(b, &res)
	if err != nil {
		return -1, nil, err
	}

	return resp.StatusCode, res, nil
}

func CallAPI(method, apiPath string, body *strings.Reader) (*httptest.ResponseRecorder, error) {
	var req *http.Request
	var err error

	w := httptest.NewRecorder()
	if body != nil {
		req, err = http.NewRequest(method, apiPath, body)
	} else {
		req, err = http.NewRequest(method, apiPath, http.NoBody)
	}
	if err != nil {
		return nil, err
	}
	req.RequestURI = apiPath
	Engine.ServeHTTP(w, req)

	return w, nil
}

func SetupTestServer(db connection.Conn) *gin.Engine {
	gin.SetMode(gin.TestMode)
	engine := gin.New()

	api := engine.Group("/api")
	server.NewRouter(api, db)
	return engine
}
