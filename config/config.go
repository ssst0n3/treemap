package config

import (
	"errors"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"log"
	"os"
	"strings"
)

const ProjectName = "/treemap"

var ProjectDir string

func init() {
	GetProjectDir()
}

func GetProjectDir() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	dir, err := os.Getwd()
	awesome_error.CheckFatal(err)
	index := strings.Index(dir, ProjectName)
	if index < 0 {
		log.Println(dir, ProjectName)
		awesome_error.CheckFatal(errors.New("index must be bigger than -1"))
	}
	ProjectDir = dir[:index+len(ProjectName)] + "/"
}
