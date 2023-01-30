package middleware

import (
    "net/http"
    "io"
    "mime"
    "gorm.io/gorm"
	"github.com/go-redis/redis/v8"
)

type {{.name}} struct {
	RedisClient *redis.Client
	BoDB        *gorm.DB
}

func New{{.name}}(redisClient *redis.Client, boDB *gorm.DB) *{{.name}} {
	return &{{.name}}{
        RedisClient: redisClient,
        BoDB:        boDB,
    }
}

func (m *{{.name}})Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        // 驗證是否為 application/json
        ctHdr := r.Header.Get("Content-Type")
        contentType, _, err := mime.ParseMediaType(ctHdr)
        if contentType != "application/json" {
            respx.Response(w, nil, errors.New("request form format error."))
            return
        }
        // 讀取Body後
        bodyBytes, _ := io.ReadAll(r.Body)
        r.Body.Close()
        r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		// TODO generate middleware implement function, delete after code implementation
		log.Println(bodyBytes)
		// var auth form.MerchantAuth
		// err = json.Unmarshal(bodyBytes, &auth)
		// Passthrough to next handler if need
		next(w, r)
	}
}
