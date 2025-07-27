# Docker Image Uploader GitHub Action

![GitHub Action](https://img.shields.io/badge/GitHub-Action-blue)
![Docker Support](https://img.shields.io/badge/Docker-✓-blue)
![SCP Transfer](https://img.shields.io/badge/SCP-Transfer-green)

A GitHub Action that builds Docker images and securely transfers them to remote servers via SSH/SCP.

## Features

- 🐳 Build Docker images from specified Dockerfiles
- 🔒 Secure transfer using SSH with private key authentication
- ⚙️ Customizable image naming and tagging
- 🖥️ Configurable remote server settings

## Usage


## Input Parameters

| Parameter           | Required     | Default          | Description |
|---------------------|--------------|------------------|-------------|
| `dockerfile_path`   | No           | `./`             | Path to directory containing Dockerfile |
| `image_name`        | No           | Repository name  | Name for the Docker image |
| `image_tag`         | No           | `latest`         | Tag for the Docker image |
| `private_key`       | Yes          | -                | SSH private key for authentication |
| `host`              | Yes          | -                | Remote server hostname/IP address |
| `port`              | No           | `22`             | SSH port number |
| `remote_copy_path`  | Yes          | -                | Target directory on remote server |


### Basic Example

```yaml
steps:
- uses: actions/checkout@v4
- uses: Avarets123/docker-image-uploader@master
  with:
    dockerfile_path: ./docker/
    image_name: my-app
    private_key: ${{ secrets.SSH_PRIVATE_KEY }}
    host: production.example.com
    remote_copy_path: /opt/docker-images/
```



## 🖥️ Server Requirements

| Requirement | Details |
|------------|---------|
| 🔑 **SSH access enabled** | Port 22 (or custom port specified in `port` parameter) |
| 💾 **Sufficient storage space** | Minimum 2x the size of your Docker image |


## 🔄 Post-Deployment

After successful image transfer, these next steps are recommended:

### Load the Docker Image
```bash
docker load < /path/to/transferred/image.tar