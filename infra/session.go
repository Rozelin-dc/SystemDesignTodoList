package infra

import (
	"github.com/Rozelin-dc/SystemDesignTodoList/domain/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SessionInfra struct {
	DB *sqlx.DB
}

func NewSessionInfra(db *sqlx.DB) *SessionInfra {
	return &SessionInfra{DB: db}
}

func (si *SessionInfra) CreateSession(userId string) (*model.Session, error) {
	uu, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	uuidStr := uu.String()

	_, err = si.DB.Exec(
		"INSERT INTO `sessions` (`session_id`, `user_id`) VALUES (?, ?)",
		uuidStr,
		userId,
	)
	if err != nil {
		return nil, err
	}

	return &model.Session{
		SessionId: uuidStr,
		UserId:    userId,
	}, nil
}

func (si *SessionInfra) GetSession(sessionId string) (*model.Session, error) {
	sess := model.Session{}
	err := si.DB.Get(
		&sess,
		"SELECT * FROM `sessions` WHERE `session_id` = ?",
		sessionId,
	)
	if err != nil {
		return nil, err
	}

	return &sess, nil
}

func (si *SessionInfra) DeleteSessionsBySessionId(sessionId string) error {
	_, err := si.DB.Exec(
		"DELETE FROM `sessions` WHERE `session_id` = ?",
		sessionId,
	)
	return err
}

func (si *SessionInfra) DeleteSessionsByUserId(userId string) error {
	_, err := si.DB.Exec(
		"DELETE FROM `sessions` WHERE `user_id` = ?",
		userId,
	)
	return err
}

func (si *SessionInfra) CheckSession(sessionId string) error {
	sess := model.Session{}
	err := si.DB.Get(
		&sess,
		"SELECT * FROM `sessions` WHERE `session_id` = ?",
		sessionId,
	)
	if err != nil {
		return err
	}

	return nil
}
