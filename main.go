package main

import (
	"image-uploader/internal/docker"
	"image-uploader/internal/scp"
	"os"
)

func main() {
	scp.ScpFile(dockerBuildAndSave())
}

func dockerBuildAndSave() (filepath string) {
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

	docker := docker.NewDockerCmd(dockerfilePath, imageName, tagName)
	docker.BuildImage()
	return docker.SaveImage("/tmp/dimages/")
}
