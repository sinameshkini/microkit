package logger

import (
	"context"
	"github.com/sinameshkini/microkit/pkg/utils"
)

func (l *Logs) logContext(ctx context.Context) {
	l.With(FieldUserID, utils.GetUserIDFromCtx(ctx))
	l.With(FieldRequestID, utils.GetRequestIDFromCtx(ctx))
}
