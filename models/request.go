package models

import ()

type Request struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}
