package midlewares

import (
	"context"
	"diplomPlugService/internal/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type wrappedServerStream struct {
	grpc.ServerStream
	ctx context.Context
}

var restRoutsWithoutAuthorization [2]string = [2]string{"/login", "/refresh"}
var grpcRoutsWithoutAuthorization [1]string = [1]string{"Login"}

func CheckAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, rout := range restRoutsWithoutAuthorization {
			if r.URL.String() == rout {
				next.ServeHTTP(w, r)
				return
			}
		}
		user, err := checkToken(r.Header.Get("Authorization"))
		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte("1invalid token"))
			return
		}
		ctx := context.WithValue(r.Context(), models.UserKeyForContext, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAndFirstTokenCheck(tokensArr []string) (string, error) {
	if len(tokensArr) != 1 {
		return "", fmt.Errorf("invalid authorization")
	}
	return tokensArr[0], nil
}

func checkToken(token string) (models.UserInfo, error) {
	var infoFromToken models.JwtClaims
	authHeader := strings.Split(token, " ")
	if len(authHeader) != 2 {
		return infoFromToken.UserInfo, fmt.Errorf("invalid token")
	}

	tokenPrefix := authHeader[0]
	tokenString := authHeader[1]

	if tokenPrefix != "Bearer" {
		return infoFromToken.UserInfo, fmt.Errorf("invalid token")
	}

	_, err := jwt.ParseWithClaims(tokenString, &infoFromToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret word"), nil
	})
	if err != nil {
		return infoFromToken.UserInfo, fmt.Errorf("invalid token")
	}
	if infoFromToken.Type != "access" {
		return infoFromToken.UserInfo, fmt.Errorf("invalid token")
	}

	return infoFromToken.UserInfo, nil
}

func CheckAuthorizationUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	for _, rout := range grpcRoutsWithoutAuthorization {
		checkRouts := strings.Split(info.FullMethod, "/")
		if checkRouts[len(checkRouts)-1] == rout {
			return handler(ctx, req)
		}
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("invalid token")
	}
	text, err := getAndFirstTokenCheck(md.Get("authorization"))
	if err != nil {
		return "", fmt.Errorf("invalid token")
	}
	user, err := checkToken(text)
	if err != nil {
		return "", fmt.Errorf("invalid token")
	}
	ctxWithUser := context.WithValue(ctx, models.UserKeyForContext, user)
	return handler(ctxWithUser, req)
}

func CheckAuthorizationStreamInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	for _, rout := range grpcRoutsWithoutAuthorization {
		checkRouts := strings.Split(info.FullMethod, "/")
		if checkRouts[len(checkRouts)-1] == rout {
			return handler(srv, ss)
		}
	}
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return fmt.Errorf("invalid token")
	}
	text, err := getAndFirstTokenCheck(md.Get("authorization"))
	if err != nil {
		return fmt.Errorf("invalid token")
	}
	user, err := checkToken(text)
	if err != nil {
		return fmt.Errorf("invalid token")
	}
	ctxWithUser := context.WithValue(ss.Context(), models.UserKeyForContext, user)
	return handler(srv, &wrappedServerStream{ServerStream: ss, ctx: ctxWithUser})
}
