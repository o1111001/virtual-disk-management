// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/o1111001/virtual-disk-management/server/disks"
)

// Injectors from modules.go:

func ComposeApiServer(port HTTPPortNumber) (*APIServer, error) {
	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}
	store := disks.NewStore(db)
	httpHandlerFunc := disks.HTTPHandler(store)
	apiServer := &APIServer{
		Port:          port,
		Handler: 			 httpHandlerFunc,
	}
	return apiServer, nil
}