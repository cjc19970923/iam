package crmapiserver

import (
	"github.com/marmotedu/iam/internal/crmapiserver/options"
	"github.com/marmotedu/iam/internal/crmapiserver/store"
	"github.com/marmotedu/iam/internal/crmapiserver/store/mysql"
	genericapiserver "github.com/marmotedu/iam/internal/pkg/server"
	"github.com/marmotedu/iam/pkg/app"
	"github.com/marmotedu/iam/pkg/log"
	"github.com/marmotedu/iam/pkg/shutdown"
	"github.com/marmotedu/iam/pkg/shutdown/shutdownmanagers/posixsignal"
)

func NewApp() *app.App {
	opts := options.NewOptions()
	application := app.NewApp(
		"crm-apiserver",
		"crm-apiserver",
		app.WithOptions(opts),
		app.WithDescription("crm-apiserver"),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run(opts)),
	)

	return application
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		log.Init(opts.Log)
		gs := shutdown.New()
		gs.AddShutdownManager(posixsignal.NewPosixSignalManager())
		genericConfig := genericapiserver.NewConfig()
		if err := opts.GenericServerRunOptions.ApplyTo(genericConfig); err != nil {
			return err
		}

		if err := opts.FeatureOptions.ApplyTo(genericConfig); err != nil {
			return err
		}

		if err := opts.InsecureServing.ApplyTo(genericConfig); err != nil {
			return err
		}

		genericApiServer, err := genericConfig.Complete().New()
		if err != nil {
			return err
		}

		//在这里设置仓库层，可以替换成mysql redis etcd等
		storeIns, _ := mysql.GetMySQLFactory(opts.MySQLOptions)
		store.SetClient(storeIns)
		initRouter(genericApiServer.Engine)
		gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
			mysqlStore, _ := mysql.GetMySQLFactory(nil)
			mysqlStore.Close()
			genericApiServer.Close()
			return nil
		}))

		// start shutdown managers
		if err := gs.Start(); err != nil {
			log.Fatalf("start shutdown manager failed: %s", err.Error())
		}

		return genericApiServer.Run()

	}
}
