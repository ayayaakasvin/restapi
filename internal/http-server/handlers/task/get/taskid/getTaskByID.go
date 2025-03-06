package get

import (
	"errors"
	"log/slog"
	"net/http"

	"restapi/internal/errorset"
	"restapi/internal/lib/sl"
	"restapi/internal/models"
	"restapi/internal/models/status"
	"restapi/internal/http-server/handlers/task/get"

	"github.com/gin-gonic/gin"
)

type Request struct {
	TaskID int64 `json:"taskId" binding:"required,gt=0"`
}

type Response struct {
	Status 	status.Status `json:"status"`
	Task 	*models.Task `json:"task,omitempty"`
}

func GetTaskHandler (log *slog.Logger, tg get.TasksGetter) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "handlers.task.get.GetTaskHandler"
		requestID, exists := c.Get("RequestID")
        if !exists {
            requestID = "unknown"
        }

        log = log.With(
            slog.String("op", op),
            slog.String("request_id", requestID.(string)),
        )

		var req Request
		// Reads the body of the request and binds it to the Request struct
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Error("failed to bind request", sl.Err(err))
			responseError(c, http.StatusBadRequest, "failed to bind request")
			return
		}

		log.Info("decoded request", slog.Any("req", req))

		task, err := tg.GetTaskByTaskID(req.TaskID)
		if err != nil {
			if errors.Is(err, errorset.ErrTaskNotFound) {
				log.Error(err.Error(), sl.Err(err))
				responseError(c, http.StatusNotFound, err.Error())
				return
			}
			
			log.Error("failed to get task", sl.Err(err))
			responseError(c, http.StatusInternalServerError, "failed to get task")
			return
		}

		responseOk(c, task)
	}
}

func responseOk(c *gin.Context, task *models.Task) {
	c.JSON(http.StatusOK, Response{
		Status: status.OK(),
		Task: task,
	})
}

func responseError(c *gin.Context, code int, errormsg string) {
	c.JSON(code, Response{
		Status: status.Error(errormsg),
	})
}