package docker

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type DockerCmd struct {
	dockerfilePath, imageName, tagName string
}

func NewDockerCmd(dockerfilePath, imageName, tagName string) *DockerCmd {
	return &DockerCmd{
		dockerfilePath: dockerfilePath,
		imageName:      imageName,
		tagName:        tagName,
	}
}

func (d *DockerCmd) BuildImage() {
	args := []string{
		"build",
		"-t",
		fmt.Sprintf("%s:%s", d.imageName, d.tagName),
		d.dockerfilePath,
	}

	fmt.Printf("RUN cmd: docker %s %s %s %s \n", args[0], args[1], args[2], args[3])

	err := runCmd("docker", args...)
	if err != nil {
		log.Fatal(err)
	}

}

func (d *DockerCmd) SaveImage(dir string) (pathToFile string) {

	err := runCmd("mkdir", "0777", "-p", dir)
	if err != nil {
		log.Fatal(err)
	}

	imageWithTag := fmt.Sprintf("%s:%s", d.imageName, d.tagName)

	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}

	fpath := fmt.Sprintf("%s%s_%s.tar", dir, d.imageName, d.tagName)

	args := []string{
		"save",
		imageWithTag,
		"-o",
		fpath,
	}

	fmt.Printf("RUN cmd: docker %s %s %s %s \n", args[0], args[1], args[2], args[3])
	err = runCmd("docker", args...)
	if err != nil {
		log.Fatal(err)
	}

	return fpath

}

func runCmd(cmdStr string, args ...string) error {
	cmd := exec.Command(cmdStr, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()

}
