package auth

import "context"

func GetUserID(ctx context.Context) string {
	v := ctx.Value("user_id")
	if v == nil {
		return ""
	}
	return v.(string)
}

func GetRole(ctx context.Context) string {
	v := ctx.Value("role")
	if v == nil {
		return ""
	}
	return v.(string)
}
