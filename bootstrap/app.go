package bootstrap

import (
	"github.com/hibiken/asynq"
	"gitlab.com/mandoor/internal/validator"
)

type Application struct {
	Env *Env
	//Mongo mongo.Client
	Mysql      MysqlClient
	RedisCache asynq.RedisClientOpt
	Validator  validator.IValidator
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	////app.Mongo = NewMongoDatabase(app.Env)
	app.Mysql = NewMysqlDatabase(app.Env)

	//app.RedisCache = NewRedisClient(app.Env)

	redisOpt := asynq.RedisClientOpt{
		Addr: app.Env.DBRedisSource,
	}
	app.RedisCache = redisOpt
	app.Validator = validator.NewCustomValidator()
	////if err := logger.InitLogger(); err != nil {
	////	log.Fatal(err)
	////}
	//
	//app.HttpClient = httpclient.NewClient()
	//
	//log := logger.NewZapLogger("", context.Background())
	//app.ZapLogger = log

	return *app
}
