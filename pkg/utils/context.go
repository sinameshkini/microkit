package utils

import (
	"context"
	"github.com/sinameshkini/microkit/models"
)

func SetUserIDToCtx(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, models.USERID, id)
}

func GetUserIDFromCtx(ctx context.Context) string {
	uid, _ := ctx.Value(models.USERID).(string)
	return uid
}

func SetRequestIDToCtx(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, models.REQUEST, id)
}

func GetRequestIDFromCtx(ctx context.Context) string {
	uid, _ := ctx.Value(models.REQUEST).(string)
	return uid
}

func SetLangToCtx(ctx context.Context, lang string) context.Context {
	return context.WithValue(ctx, models.LANG, lang)
}

func GetLangFromCtx(ctx context.Context) string {
	uid, _ := ctx.Value(models.LANG).(string)
	return uid
}

func SetRoleToCtx(ctx context.Context, role string) context.Context {
	return context.WithValue(ctx, models.ROLE, role)
}

func GetRoleFromCtx(ctx context.Context) string {
	uid, _ := ctx.Value(models.ROLE).(string)
	return uid
}
