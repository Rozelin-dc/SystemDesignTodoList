package repository

import "github.com/Rozelin-dc/SystemDesignTodoList/domain/model"

type SessionRepository interface {
	CreateSession(userId string) (*model.Session, error)
	GetSession(sessionId string) (*model.Session, error)
	DeleteSessionsBySessionId(sessionId string) error
	DeleteSessionsByUserId(userId string) error
	CheckSession(sessionId string) error
}
