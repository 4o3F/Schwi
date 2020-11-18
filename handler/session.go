package handler

import (
	"fmt"
	"github.com/CardinalDevLab/Schwi-Backend/database"
	"github.com/alexedwards/scs/sqlite3store"
	"github.com/alexedwards/scs/v2"
	"net/http"
	"time"
)

var SessionManager *scs.SessionManager

func InitSession() {
	SessionManager = scs.New()
	SessionManager.Store = sqlite3store.New(database.MainDatabase)
	SessionManager.Lifetime = 24 * time.Hour
	SessionManager.Cookie.SameSite = http.SameSiteNoneMode
	SessionManager.Cookie.Secure = true
	SessionManager.Cookie.Domain = database.CookieDomain

	fmt.Printf("Init Session")
}
