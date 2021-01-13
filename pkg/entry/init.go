package entry

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

var (
	readme      = "README.md"
	dirs        = []string{"build", "cmd", "configs", "deployments", "docs", "internal", "pkg", "scripts", "test"}
	mainContent = "package main\n\nfunc main() {\n\n}"
)

// Initiator interface
type Initiator interface {
	Init()
}

type initiator struct {
	appName string
}

// NewInitiator Factory
func NewInitiator(appName string) Initiator {
	return &initiator{appName: appName}
}

// Init inits a go project
func (i *initiator) Init() {

	i.createRootDir()
	i.createL1Dirs()
	i.createCMDDir()
	i.createMainFile()
	i.createREADME()

	fmt.Println("Go project init complete, happy coding :)")

}

func (i *initiator) createRootDir() {
	mkdir("", i.appName)
}

func (i *initiator) createL1Dirs() {
	mkdirs(i.appName+"/", dirs)
}

func (i *initiator) createCMDDir() {
	mkdir(i.appName+"/", "/cmd/"+i.appName)
}

func (i *initiator) createMainFile() {
	mainFile := i.appName + "/cmd/" + i.appName + "/" + i.appName + ".go"
	touch(mainFile)
	write(mainContent, mainFile)
}

func (i *initiator) createREADME() {
	readmeFile := i.appName + "/" + readme
	touch(readmeFile)
	write("# "+i.appName+"\n", readmeFile)
}

func mkdir(rootDir, dir string) error {
	cmd := exec.Command("mkdir", rootDir+dir)
	return cmd.Run()
}

func mkdirs(rootDir string, dirs []string) {
	for _, dir := range dirs {
		err := mkdir(rootDir, dir)
		if err != nil {
			fmt.Printf("./%s/ already exists, skipping.\n", dir)
		}
	}
}

func touch(fileName string) error {
	cmd := exec.Command("touch", fileName)
	return cmd.Run()
}

func write(content, fileName string) error {
	return ioutil.WriteFile(fileName, []byte(content), 644)
}
