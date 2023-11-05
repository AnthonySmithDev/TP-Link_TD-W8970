package main

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"tp-link/td-w8970/wireless"
)

// App struct
type App struct {
	ctx context.Context
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

func (a *App) WirelessRead() *wireless.Class {
	s := wireless.NewSlice().SetEnable().SetSSID().SetPassword()
	class, err := wireless.Get(s)
	if err != nil {
		return nil
	}
	return class
}

func (a *App) WirelessWrite(enable bool, name, password string) error {
	m := wireless.NewMap().SetEnable(enable).SetSSID(name).SetPassword(password)
	return wireless.Put(m)
}

const Password = "admin"

func (a *App) SystemRead() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Password
	}
	newFile := filepath.Join(homeDir, "pwd.txt")
	data, err := os.ReadFile(newFile)
	if err != nil {
		return Password
	}
	if string(data) == "" {
		return Password
	}
	return string(data)
}

func (a *App) SystemWrite(password string) error {
	time.Sleep(time.Millisecond * 500)
	homeDir, _ := os.UserHomeDir()
	newFile := filepath.Join(homeDir, "pwd.txt")
	err := os.WriteFile(newFile, []byte(password), 0644)
	if err != nil {
		return err
	}
	return nil
}
