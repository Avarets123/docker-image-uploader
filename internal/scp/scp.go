package scp

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh"
)

func ScpFile(filepath string) {

	username := os.Getenv("INPUT_USERNAME")
	if username == "" {
		log.Fatal("Username not passed!")
	}

	host := os.Getenv("INPUT_HOST")
	if host == "" {
		log.Fatal("Host not passed!")
	}

	port := os.Getenv("INPUT_PORT")
	if port == "" {
		port = "22"
	}

	privateKey := os.Getenv("INPUT_PRIVATE_KEY")
	if privateKey == "" {
		log.Fatal("PrivateKey not setted!")
	}

	remoteDir := os.Getenv("INPUT_REMOTE_COPY_PATH")
	if remoteDir == "" {
		log.Fatal("PrivateKey not setted!")
	}

	b := []byte(privateKey)

	signer, err := ssh.ParsePrivateKey(b)
	if err != nil {
		log.Fatal(err)
	}

	config := &ssh.ClientConfig{
		User:            username,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	con, err := ssh.Dial("tcp", host+":"+port, config)
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	scp(con, filepath, remoteDir)
}

func scp(con *ssh.Client, localFile, remoteDir string) {
	f, err := os.Open(localFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fInfo, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	sess, err := con.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	stdinPipe, err := sess.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stdinPipe.Close()

	remoteFile := filepath.Join(remoteDir, filepath.Base(localFile))

	copyCmd := fmt.Sprintf("scp -t %s", remoteDir)
	err = sess.Start(copyCmd)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Start copy file.\n Size: %d\n Remote dir: %s\n", fInfo.Size(), remoteFile)

	fmt.Fprintf(stdinPipe, "C%04o %d %s\n", 0777, fInfo.Size(), filepath.Base(remoteFile))

	cb, err := io.Copy(stdinPipe, f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d bytes was copied!", cb)
	fmt.Fprint(stdinPipe, "\x00")

}
