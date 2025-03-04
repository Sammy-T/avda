package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/sammy-t/avdu"
	"github.com/sammy-t/avdu/vault"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx       context.Context
	vaultData *vault.Vault
	config    *Config
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// Load config
	config, err := loadConfig()
	if err != nil {
		log.Printf("Error loading config: %v\n", err)
		config = &Config{}
	}
	a.config = config
}

// GetAppInfo returns the info from the app's config file.
func (a *App) GetAppInfo() Response {
	return Response{
		Status: "success",
		Data:   appInfo,
	}
}

// SelectVault opens a file selection dialog
// and returns the selected file's path.
func (a *App) SelectVault() Response {
	filters := []runtime.FileFilter{
		{DisplayName: "Vault Files (*.json)", Pattern: "*.json"},
	}

	options := runtime.OpenDialogOptions{
		Title:            "Select vault file",
		Filters:          filters,
		DefaultFilename:  a.config.LastVaultPath,
		DefaultDirectory: filepath.Dir(a.config.LastVaultPath),
	}

	filePath, err := runtime.OpenFileDialog(a.ctx, options)

	response := Response{}

	if err != nil {
		log.Printf("unable to open file: %v\n", err)

		response.Status = "error"
		response.Message = err.Error()
	} else {
		response.Status = "success"
		response.Data = filePath

		// Save the selected path to config
		if a.config.LastVaultPath != filePath && filePath != "" {
			a.config.LastVaultPath = filePath
			if err := saveConfig(a.config); err != nil {
				log.Printf("Error saving config: %v\n", err)
			}
		}
	}

	return response
}

// GetLastVaultPath returns the last used vault path from config
func (a *App) GetLastVaultPath() Response {
	return Response{
		Status: "success",
		Data:   a.config.LastVaultPath,
	}
}

// OpenVault reads the vault file at the provided path,
// sets the vault data on the app struct,
// starts OTP watching in a goroutine,
// and returns whether the read was successful.
func (a *App) OpenVault(filePath string, password string) Response {
	var vaultData *vault.Vault
	var err error

	if password == "" {
		vaultData, err = avdu.ReadVaultFile(filePath)
	} else {
		vaultData, err = avdu.ReadAndDecryptVaultFile(filePath, password)
	}

	var response Response

	if err != nil {
		errMsg := fmt.Sprintf("unable to open vault at %q: %v", filePath, err)
		log.Println(errMsg)

		response.Status = "error"
		response.Message = errMsg
	} else {
		response.Status = "success"

		a.vaultData = vaultData

		go a.watchOTPs()
	}

	return response
}

// CloseVault removes the vault data.
func (a *App) CloseVault() {
	a.vaultData = nil
}

// watchOTPs loops through updating the OTPs
// until the current vault data is removed.
//
// NOTE: This is intended to be called as a goroutine.
func (a *App) watchOTPs() {
	a.updateOTPs()

	for a.vaultData != nil {
		ttn := avdu.GetTTN()

		if ttn > 29000 {
			a.updateOTPs()
		}

		runtime.EventsEmit(a.ctx, "onTimeUpdated", ttn)

		time.Sleep(1000 * time.Millisecond)
	}
}

// updateOTPs loads the OTPs for the current vault data
// and emits an "onCodesUpdated" event containing
// an array of each entry with its current OTP.
func (a *App) updateOTPs() {
	otps, err := avdu.GetOTPs(a.vaultData)
	if err != nil {
		log.Println(err)
	}

	var entryCodes []EntryCode

	for _, entry := range a.vaultData.Db.Entries {
		entryCode := EntryCode{
			Entry: entry,
			Code:  fmt.Sprintf("%v", otps[entry.Uuid]),
		}

		entryCodes = append(entryCodes, entryCode)
	}

	runtime.EventsEmit(a.ctx, "onCodesUpdated", entryCodes)
}
