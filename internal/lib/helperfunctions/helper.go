package helper

import (
	"fmt"
	"log/slog"
	"restapi/internal/errorset"
	"restapi/internal/lib/jwtutil"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	TaskIDKey 			= "taskId"
	UserIDKey 			= "userId"
	TaskKey 			= "task"
	TasksKey 			= "tasks"
	UserKey 			= "user"
	ReqKey 				= "request"
	UsernameKey 		= "username"
	AuthorizationHeader = "Authorization"
)

func GetIDFromParams(c *gin.Context, idkey string) int64 {
	if idkey == "" {
		return -1
	}

	taskIDString := c.Param(idkey)
	if taskIDString == "" {
		return -1
	}
	taskID, err := strconv.ParseInt(taskIDString, 10, 0)
	if err != nil {
		return -1
	}

	return taskID
}

func LoadLogger(log *slog.Logger, c *gin.Context, operation string) *slog.Logger  {
	requestID, exists := c.Get("X-Request-ID")
	if !exists {
		requestID = "unknown"
	}

	newLogger := log.With(
		slog.String("op", operation),
		slog.String("X-Request-ID", requestID.(string)),
	)

	return newLogger
}

func FetchIDFromToken(c *gin.Context, idkey string) (int64) {
	token, err := FetchTokenFromContext(c)
	if err != nil {
		return -1
	}
	
	claims, err := jwtutil.ValidateJWT(token)
	if err != nil {
		return -1
	}

	id, ok := claims[idkey].(int64)
	if !ok || id == 0 {
		return -1
	}

	return id
}

func FetchTokenFromContext(c *gin.Context) (string, error) {
	authHeader := c.GetHeader(AuthorizationHeader)
	if authHeader == "" {
		return "", fmt.Errorf(errorset.ErrAuthorizationMissing)
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return "", fmt.Errorf(errorset.ErrAuthorizationMissing)
	}

	return tokenString, nil
}