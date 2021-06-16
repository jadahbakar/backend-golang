package utils

import (
	"log"

	"github.com/jadahbakar/backend-golang/pkg/version"
)

func AppName() {
	log.Println(appName + version.Version)
}
