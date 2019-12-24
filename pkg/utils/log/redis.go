package log

import (
	"github.com/rogierlommers/logrus-redis-hook"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

func NewLogRedis(log *logrus.Logger) {
	hookConfig := logredis.HookConfig{
		Host:     "localhost",
		Password: "1234",
		Key:      "redis:log",
		Format:   "v0",
		App:      "my_app_name",
		Port:     6379,
		Hostname: "my_app_hostname", // will be sent to field @source_host
		DB:       0,                 // optional
		TTL:      3600,
	}

	hook, err := logredis.NewHook(hookConfig)
	if err == nil {
		log.AddHook(hook)
	} else {
		log.Errorf("logredis error: %q", err)
	}
	log.SetOutput(ioutil.Discard)
}
