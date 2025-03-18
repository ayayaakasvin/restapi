package handlers

import (
	"log/slog"

	"github.com/ayayaakasvin/restapigolang/internal/http-server/handlers/task"
	"github.com/ayayaakasvin/restapigolang/internal/http-server/handlers/user"
	"github.com/ayayaakasvin/restapigolang/internal/storage"
)

type Handlers struct {
	Task task.TaskHandlers
	User user.UserHandlers
}

func NewHandlers(db storage.Storage, log *slog.Logger) *Handlers {
	return &Handlers{
		Task: task.NewTaskHandler(log, db),
		User: user.NewUserHandler(log, db),
	}
}
