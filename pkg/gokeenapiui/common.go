package gokeenapiui

import (
	"fmt"

	"github.com/noksa/gokeenapi/pkg/config"
	"github.com/spf13/viper"
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
	viper.Set(config.ViperKeeneticLogin, login)
	viper.Set(config.ViperKeeneticPassword, password)
	viper.Set(config.ViperKeeneticUrl, url)
}

func (*common) RouterDataFilled() error {
	var mErr error
	if viper.GetString(config.ViperKeeneticLogin) == "" {
		mErr = multierr.Append(mErr, fmt.Errorf("логин от роутера не введен"))
	}
	if viper.GetString(config.ViperKeeneticPassword) == "" {
		mErr = multierr.Append(mErr, fmt.Errorf("пароль от роутера не введен"))
	}
	if viper.GetString(config.ViperKeeneticUrl) == "" {
		mErr = multierr.Append(mErr, fmt.Errorf("URL роутера не введен"))
	}
	return mErr
}
