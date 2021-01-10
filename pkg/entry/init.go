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

// Init inits a go project
func Init(appName string) {

	mkdir("", appName)

	mkdirs(appName+"/", dirs)
	mkdir(appName+"/", "/cmd/"+appName)

	mainFile := appName + "/cmd/" + appName + "/" + appName + ".go"
	touch(mainFile)
	write(mainContent, mainFile)

	readmeFile := appName + "/" + readme
	touch(readmeFile)
	write("# "+appName+"\n", readmeFile)

	fmt.Println("Go project init complete, happy coding :)")

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
