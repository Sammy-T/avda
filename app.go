package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sammy-t/avdu"
	"github.com/sammy-t/avdu/vault"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx       context.Context
	vaultData *vault.Vault
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
		Title:   "Select vault file",
		Filters: filters,
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
	}

	return response
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

		time.Sleep(time.Duration(ttn) * time.Millisecond)
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

func (a *App) GetGroups() []vault.Group {
	return a.vaultData.Db.Groups
}
