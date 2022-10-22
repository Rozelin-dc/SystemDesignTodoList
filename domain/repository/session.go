package repository

import "github.com/Rozelin-dc/SystemDesignTodoList/domain/model"

type SessionRepository interface {
	CreateSession(userId string) (*model.Session, error)
	DeleteSessionBySessionId(sessionId string) error
	DeleteSessionByUserId(userId string) error
	IsValidSession(session *model.Session) bool
}
