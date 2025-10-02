package gokeenapiui

import (
	"fmt"

	"github.com/noksa/gokeenapi/pkg/config"
	"go.uber.org/multierr"
)

var (
	Common = common{}
)

type common struct {
}

func (*common) SetData() {
	login, _ := Bindings.RouterLogin.Get()
	password, _ := Bindings.RouterPassword.Get()
	url, _ := Bindings.RouterUrl.Get()
	config.Cfg.Keenetic.URL = url
	config.Cfg.Keenetic.Login = login
	config.Cfg.Keenetic.Password = password
}

func (*common) RouterDataFilled() error {
	var mErr error
	if config.Cfg.Keenetic.Login == "" {
		mErr = multierr.Append(mErr, fmt.Errorf("логин от роутера не введен"))
	}
	if config.Cfg.Keenetic.Password == "" {
		mErr = multierr.Append(mErr, fmt.Errorf("пароль от роутера не введен"))
	}
	if config.Cfg.Keenetic.URL == "" {
		mErr = multierr.Append(mErr, fmt.Errorf("URL роутера не введен"))
	}
	return mErr
}
