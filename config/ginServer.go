package config

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

// GinServerMode modo
type GinServerMode int

const (
	DebugMode GinServerMode = iota
	ReleaseMode
	TestMode
)

// GinServer : the struct gathering all the server details
type GinServer struct {
	port   int
	Router *gin.Engine
}

// NewServer server
func NewServer(port int, mode GinServerMode) GinServer {
	s := GinServer{}
	s.port = port

	s.Router = gin.New()

	switch mode {
	case DebugMode:
		gin.SetMode(gin.DebugMode)
	case TestMode:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	s.Router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	s.Router.Use(gin.Recovery())

	SetCors(s.Router, "*")

	return s
}

// SetCors is a helper to set current engine cors
func SetCors(engine *gin.Engine, allowedOrigins string) {
	engine.Use(cors.Middleware(cors.Config{
		Origins:         allowedOrigins,
		Methods:         strings.Join([]string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodOptions, http.MethodPatch}, ","),
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
}

// Start the server
func (s GinServer) Start() {
	s.Router.Run(":" + strconv.Itoa(int(s.port)))
}
