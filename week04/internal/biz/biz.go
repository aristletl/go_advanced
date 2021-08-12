package biz

import (
	"context"
	"errors"
	"fmt"
	"week04/internal/data"
)

func SayHello(ctx context.Context, m *data.DBModel, account string) string {
	query := `select name from user where account =?`
	name, err := m.GetUserName(account, query)
	switch errors.Unwrap(err) {
	case data.ErrNoUser:
		return "Hello, my friend."
	default:
		return fmt.Sprintf("hello,%s", name)
	}
}
