package gokeenapiui

import "fyne.io/fyne/v2/data/binding"

var (
	Bindings = newBindings()
)

type bindings struct {
	RouterUrl      binding.String
	RouterLogin    binding.String
	RouterPassword binding.String
	AwgConfFile    binding.String
	AwgName        binding.String
}

func newBindings() *bindings {
	b := &bindings{}
	b.RouterUrl = binding.NewString()
	b.RouterLogin = binding.NewString()
	b.RouterPassword = binding.NewString()
	b.AwgConfFile = binding.NewString()
	b.AwgName = binding.NewString()
	return b
}
