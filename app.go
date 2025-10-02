package main

import (
	"context"
	"fmt"
	"time"

	"github.com/noksa/gokeenapi/pkg/config"
	"github.com/noksa/gokeenapi/pkg/gokeenrestapi"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"go.uber.org/multierr"
)

// App struct
type App struct {
	ctx context.Context
}

// RouterConfig holds router connection data
type RouterConfig struct {
	URL      string `json:"url"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

// AWGConfig holds AWG connection data
type AWGConfig struct {
	Name     string `json:"name"`
	FilePath string `json:"filePath"`
}

// ProgressUpdate represents progress information
type ProgressUpdate struct {
	Message string `json:"message"`
	Step    int    `json:"step"`
	Total   int    `json:"total"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// OpenFileDialog opens a file selection dialog
func (a *App) OpenFileDialog() (string, error) {
	selection, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Выбор AWG конфиг файла",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Config Files (*.conf)",
				Pattern:     "*.conf",
			},
			{
				DisplayName: "All Files (*.*)",
				Pattern:     "*.*",
			},
		},
	})
	return selection, err
}

// ValidateRouterConfig validates router configuration data
func (a *App) ValidateRouterConfig(routerConfig RouterConfig) error {
	var mErr error
	if routerConfig.Login == "" {
		mErr = multierr.Append(mErr, fmt.Errorf("логин от роутера не введен"))
	}
	if routerConfig.Password == "" {
		mErr = multierr.Append(mErr, fmt.Errorf("пароль от роутера не введен"))
	}
	if routerConfig.URL == "" {
		mErr = multierr.Append(mErr, fmt.Errorf("URL роутера не введен"))
	}
	return mErr
}

// CreateAWGConnection creates a new AWG connection
func (a *App) CreateAWGConnection(routerConfig RouterConfig, awgConfig AWGConfig) error {
	// Set configuration
	config.Cfg.Keenetic.URL = routerConfig.URL
	config.Cfg.Keenetic.Login = routerConfig.Login
	config.Cfg.Keenetic.Password = routerConfig.Password

	// Validate router data
	if err := a.ValidateRouterConfig(routerConfig); err != nil {
		return err
	}

	// Authenticate
	if err := gokeenrestapi.Common.Auth(); err != nil {
		return fmt.Errorf("ошибка авторизации: %w", err)
	}

	// Check AWG interface from config file
	if err := gokeenrestapi.Checks.CheckAWGInterfaceExistsFromConfFile(awgConfig.FilePath); err != nil {
		return fmt.Errorf("ошибка чтения conf файла: %w", err)
	}

	// Create interface
	createdInterface, err := gokeenrestapi.AwgConf.AddInterface(awgConfig.FilePath, awgConfig.Name)
	if err != nil {
		return fmt.Errorf("ошибка создания соединения: %w", err)
	}

	// Configure interface
	time.Sleep(time.Second * 1)
	if err := gokeenrestapi.AwgConf.ConfigureOrUpdateInterface(awgConfig.FilePath, createdInterface.Created); err != nil {
		return fmt.Errorf("ошибка настройки соединения: %w", err)
	}

	// Set global IP
	if err := gokeenrestapi.Interface.SetGlobalIpInInterface(createdInterface.Created, true); err != nil {
		return fmt.Errorf("ошибка настройки IP соединения: %w", err)
	}

	// Bring interface up
	if err := gokeenrestapi.Interface.UpInterface(createdInterface.Created); err != nil {
		return fmt.Errorf("ошибка включения соединения: %w", err)
	}

	// Wait for interface to be up
	if err := gokeenrestapi.Interface.WaitUntilInterfaceIsUp(createdInterface.Created); err != nil {
		return fmt.Errorf("соединение с сервером не установлено: %w", err)
	}

	return nil
}

// GetRouterWebURL returns the router web interface URL
func (a *App) GetRouterWebURL(routerURL string) string {
	return fmt.Sprintf("%v/otherConnections", routerURL)
}
