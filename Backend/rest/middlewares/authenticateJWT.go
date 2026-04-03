package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"
)

func base64URLEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}

func (m *Middlewares) AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// parse jwt
		// parse header and payload or claims
		// hmac-sha-256 algorithm -> hash hmac(header, payload, secret key)
		// parse signature part from the jwt
		// if the signature and hash is same -> forward to create products
		// other 401 status code with Unauthorized

		header := r.Header.Get("Authorization") // taking the Authorization header from the request

		if header == "" {
			http.Error(w, "Unathorized", http.StatusUnauthorized)
			return
		}
		headerArr := strings.Split(header, " ") // split into Bearer and the token

		if len(headerArr) != 2 {
			http.Error(w, "Unathorized", http.StatusUnauthorized)
			return
		}
		accessToken := headerArr[1] // taking the token part from the header

		tokenParts := strings.Split(accessToken, ".") // split the token into header, payload and signature parts

		if len(tokenParts) != 3 {
			http.Error(w, "Unathorized", http.StatusUnauthorized)
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		signature := tokenParts[2]

		message := jwtHeader + "." + jwtPayload

		byteArrMessage := []byte(message)
		byteArrSecret := []byte(m.cnf.JwtSecretKey)

		h := hmac.New(sha256.New, byteArrSecret) // create a new HMAC hash using the SHA-256 algorithm and the secret key
		h.Write(byteArrMessage)                  // write the message (header and payload) to the hash

		hash := h.Sum(nil)
		newsignature := base64URLEncode(hash)

		if newsignature != signature {
			http.Error(w, "hacker", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
