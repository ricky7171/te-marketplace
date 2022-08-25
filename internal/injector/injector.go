//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/ricky7171/te-marketplace/internal/helper"
	"github.com/ricky7171/te-marketplace/internal/library_wrapper"
	accountappservice "github.com/ricky7171/te-marketplace/internal/modules/account/application/service"
	accountdomrepository "github.com/ricky7171/te-marketplace/internal/modules/account/domain/repository"
	accountinfrarepository "github.com/ricky7171/te-marketplace/internal/modules/account/infrastructure/repository"
	accountpresent "github.com/ricky7171/te-marketplace/internal/modules/account/presentation"
	"github.com/ricky7171/te-marketplace/internal/router"
)

// singkatan domain
// acc : account

// singkatan layer
// app : application
// dom : domain
// infra : infrastructure
// present : presentation

// singkatan sub layer
// serv : service

// singkatan service name / repo name
// authn : authentication
// autho : authorization

var accDomRepoSet = wire.NewSet(accountinfrarepository.NewAccountRepositoryPg, wire.Bind(new(accountdomrepository.AccountRepository), new(*accountinfrarepository.AccountRepositoryPg)))

var myJwtSet = wire.NewSet(library_wrapper.NewMyJwtImpl, wire.Bind(new(library_wrapper.MyJwt), new(*library_wrapper.MyJwtImpl)))

var helperJwtSet = wire.NewSet(myJwtSet, helper.NewHelperJwtImpl, wire.Bind(new(helper.HelperJwt), new(*helper.HelperJwtImpl)))

var accAppServAuthnSet = wire.NewSet(accDomRepoSet, helperJwtSet, accountappservice.NewAuthenticationServiceImpl, wire.Bind(new(accountappservice.AuthenticationService), new(*accountappservice.AuthenticationServiceImpl)))

var accPresentHandlerSet = wire.NewSet(accAppServAuthnSet, accountpresent.NewHandler)

var routerSet = wire.NewSet(gin.Default, accPresentHandlerSet, router.NewRouter)

func InitializedRouter() *router.Router {
	wire.Build(routerSet)
	return nil
}
