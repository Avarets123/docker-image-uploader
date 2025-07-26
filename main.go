package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

	dockerfilePath := os.Getenv("INPUT_DOCKERFILE_PATH")
	if dockerfilePath == "" {
		dockerfilePath = "./"
	}

	imageName := os.Getenv("INPUT_IMAGE_NAME")
	if imageName == "" {
		imageName = "upload_image"
	}

	tagName := os.Getenv("INPUT_IMAGE_TAG")
	if tagName == "" {
		tagName = "latest"
	}

	err := runCmd("env")
	if err != nil {
		log.Fatal(err)
	}

	docker := NewDockerCmd(dockerfilePath, imageName, tagName)
	docker.buildImage()
	docker.saveImage("/tmp/dimages/")

}

type dockerCmd struct {
	dockerfilePath, imageName, tagName string
}

func NewDockerCmd(dockerfilePath, imageName, tagName string) *dockerCmd {
	return &dockerCmd{
		dockerfilePath: dockerfilePath,
		imageName:      imageName,
		tagName:        tagName,
	}
}

func (d *dockerCmd) buildImage() {
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

func (d *dockerCmd) saveImage(dir string) {

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

}

func runCmd(cmdStr string, args ...string) error {
	cmd := exec.Command(cmdStr, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()

}
