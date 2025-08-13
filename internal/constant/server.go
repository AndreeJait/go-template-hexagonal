package constant

type AppMode string

const (
	DevelopmentMode AppMode = "development"
	StagingMode     AppMode = "staging"
	ProductionMode  AppMode = "production"
)
