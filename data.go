package main

import "github.com/sammy-t/avdu/vault"

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type EntryCode struct {
	Entry  vault.Entry   `json:"entry"`
	Code   string        `json:"code"`
	Groups []vault.Group `json:"groups"`
}
