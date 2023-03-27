package middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const userkey = "session_id"

// Use cookie to store session id
func SetSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(userkey))
	return sessions.Sessions("mysession", store)
}

func AuthSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionId := session.Get(userkey)
		if sessionId == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message:": "需要登入",
			})
			return
		}
		c.Next()
	}
}

// Save Session for User
func SaveSession(c *gin.Context, userId int) {
	session := sessions.Default(c)
	session.Set(userkey, userId)
	session.Options(sessions.Options{MaxAge: 3600})
	session.Save()
}

// Clear Session for User
func ClearSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Options(sessions.Options{MaxAge: -1})
	session.Save()

}

// Get Session for User
func GetSession(c *gin.Context) int {
	session := sessions.Default(c)
	sessionId := session.Get(userkey)
	if sessionId == nil {
		return 0
	}
	return sessionId.(int)
}

// Check Session for User
func CheckSession(c *gin.Context) bool {
	session := sessions.Default(c)
	sessionId := session.Get(userkey)
	return sessionId != nil
}
