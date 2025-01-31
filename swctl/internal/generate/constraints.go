package generate

const (
	RepoAddr        = "https://github.com/ugabiga/swan.git"
	BoostrapDirName = "bootstrap"
	BootstrapPath   = "swan/bootstrap"
)

const (
	AppRootPath   = "./internal/app"
	AppPath       = AppRootPath + "/app.go"
	ContainerPath = AppRootPath + "/container.go"
	EventPath     = AppRootPath + "/event.go"
	CommandPath   = AppRootPath + "/commands.go"
	RouterPath    = AppRootPath + "/server/router.go"
)
