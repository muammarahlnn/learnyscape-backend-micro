package provider

import (
	"github.com/jmoiron/sqlx"
	"github.com/muammarahlnn/learnyscape-backend/auth-service/internal/config"
	"github.com/muammarahlnn/learnyscape-backend/auth-service/internal/repository"
	"github.com/muammarahlnn/learnyscape-backend/pkg/database"
	encryptutil "github.com/muammarahlnn/learnyscape-backend/pkg/util/encrypt"
	jwtutil "github.com/muammarahlnn/learnyscape-backend/pkg/util/jwt"
	"github.com/redis/go-redis/v9"
)

var (
	db           *sqlx.DB
	rdb          *redis.ClusterClient
	bcryptHasher encryptutil.Hasher
	jwtUtil      jwtutil.JWTUtil
	dataStore    repository.DataStore
)

func BootstrapGlobal(cfg *config.Config) {
	db = database.NewPostgres((*database.PostgresOptions)(cfg.Postgres))
	rdb = database.NewRedisCluster((*database.RedisClusterOptions)(cfg.Redis))
	bcryptHasher = encryptutil.NewBcryptHasher(cfg.App.BCryptCost)
	jwtUtil = jwtutil.NewJWTUtil()
	dataStore = repository.NewDataStore(db)
}
