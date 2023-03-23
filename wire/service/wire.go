//go:build wireinject
//+build wireinject

package service

import "github.com/google/wire"

func BuildClient() *Client {
	wire.Build(NewClient, NewFoo, NewBar)
	return nil
}