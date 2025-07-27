package docker

import (
	"fmt"
	"image-uploader/pkg/common"
	"log"
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

	err := common.RunCmd("docker", args...)
	if err != nil {
		log.Fatal(err)
	}

}

func (d *DockerCmd) SaveImage(dir string) (pathToFile string) {

	err := common.RunCmd("mkdir", "0777", "-p", dir)
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
	err = common.RunCmd("docker", args...)
	if err != nil {
		log.Fatal(err)
	}

	return fpath

}
