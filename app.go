package main

import (
	"context"
	"fmt"
	"path/filepath"
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

type WgInterface struct {
	Id          string
	Description string
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

// RouteConfig holds route configuration data
type RouteConfig struct {
	InterfaceId string   `json:"interfaceId"`
	BatFiles    []string `json:"batFiles"`
	BatUrls     []string `json:"batUrls"`
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

// OpenBatFileDialog opens a file selection dialog for BAT files
func (a *App) OpenBatFileDialog() (string, error) {
	selection, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Выбор BAT файла с маршрутами",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "BAT Files (*.bat)",
				Pattern:     "*.bat",
			},
			{
				DisplayName: "All Files (*.*)",
				Pattern:     "*.*",
			},
		},
	})
	return selection, err
}

func (a *App) DeleteRoutes(interfaceId string) error {
	routes, err := gokeenrestapi.Ip.GetAllUserRoutesRciIpRoute(interfaceId)
	if err != nil {
		return err
	}
	err = gokeenrestapi.Ip.DeleteRoutes(routes, interfaceId)
	return err
}

func (a *App) AddRoutes(interfaceId string, batFiles []string, batUrls []string) error {
	for _, file := range batFiles {
		absFilePath, err := filepath.Abs(file)
		if err != nil {
			return err
		}
		err = gokeenrestapi.Ip.AddRoutesFromBatFile(absFilePath, interfaceId)
		if err != nil {
			return err
		}
	}
	for _, url := range batUrls {
		err := gokeenrestapi.Ip.AddRoutesFromBatUrl(url, interfaceId)
		if err != nil {
			return err
		}
	}
	return nil
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

// TestRouterConnection tests actual connection to router
func (a *App) TestRouterConnection(routerConfig RouterConfig) error {
	// Set configuration
	config.Cfg.Keenetic.URL = routerConfig.URL
	config.Cfg.Keenetic.Login = routerConfig.Login
	config.Cfg.Keenetic.Password = routerConfig.Password

	// Validate router data first
	if err := a.ValidateRouterConfig(routerConfig); err != nil {
		return err
	}

	// Authenticate
	if err := gokeenrestapi.Common.Auth(); err != nil {
		return fmt.Errorf("ошибка авторизации: %w", err)
	}

	return nil
}

// ValidateAWGConfig validates AWG configuration and authenticates
func (a *App) ValidateAWGConfig(routerConfig RouterConfig, awgConfig AWGConfig) error {
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

	return nil
}

// CreateAWGInterface creates AWG interface
func (a *App) CreateAWGInterface(awgConfig AWGConfig) (string, error) {
	// Create interface
	createdInterface, err := gokeenrestapi.AwgConf.AddInterface(awgConfig.FilePath, awgConfig.Name)
	if err != nil {
		return "", fmt.Errorf("ошибка создания соединения: %w", err)
	}

	return createdInterface.Created, nil
}

// ConfigureAWGInterface configures the created AWG interface
func (a *App) ConfigureAWGInterface(awgConfig AWGConfig, interfaceName string) error {
	// Configure interface
	time.Sleep(time.Second * 1)
	if err := gokeenrestapi.AwgConf.ConfigureOrUpdateInterface(awgConfig.FilePath, interfaceName); err != nil {
		return fmt.Errorf("ошибка настройки соединения: %w", err)
	}

	// Set global IP
	if err := gokeenrestapi.Interface.SetGlobalIpInInterface(interfaceName, true); err != nil {
		return fmt.Errorf("ошибка настройки IP соединения: %w", err)
	}

	return nil
}

func (a *App) ShowWgInterfaces() ([]WgInterface, error) {
	var ifaces []WgInterface
	interfaces, err := gokeenrestapi.Interface.GetInterfacesViaRciShowInterfaces(false, "Wireguard")
	if err != nil {
		return nil, err
	}
	for _, iface := range interfaces {
		ifaces = append(ifaces, WgInterface{
			Id:          iface.Id,
			Description: iface.Description,
		})
	}
	return ifaces, err
}

// ActivateAWGInterface brings the interface up and waits for it to be ready
func (a *App) ActivateAWGInterface(interfaceName string) error {
	// Bring interface up
	if err := gokeenrestapi.Interface.UpInterface(interfaceName); err != nil {
		return fmt.Errorf("ошибка включения соединения: %w", err)
	}

	// Wait for interface to be up
	if err := gokeenrestapi.Interface.WaitUntilInterfaceIsUp(interfaceName); err != nil {
		return fmt.Errorf("соединение с сервером не установлено: %w", err)
	}

	return nil
}

// GetRouterWebURL returns the router web interface URL with specific path
func (a *App) GetRouterWebURL(routerURL string, path string) string {
	if path == "" {
		return routerURL
	}
	return fmt.Sprintf("%v/%v", routerURL, path)
}

// ValidateRouteConfig validates route configuration data
func (a *App) ValidateRouteConfig(routeConfig RouteConfig) error {
	var mErr error

	if routeConfig.InterfaceId == "" {
		mErr = multierr.Append(mErr, fmt.Errorf("ID интерфейса не указан"))
	}

	if len(routeConfig.BatFiles) == 0 && len(routeConfig.BatUrls) == 0 {
		mErr = multierr.Append(mErr, fmt.Errorf("укажите хотя бы один BAT файл или URL ссылку"))
	}

	return mErr
}
