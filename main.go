package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"os/exec"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func main() {
	var localDirectory = os.Getenv("LOCAL_DIRECTORY")
	var remoteDirectory = os.Getenv("REMOTE_DIRECTORY")
	var postCheckout = os.Getenv("POST_CHECKOUT")
	var server = os.Getenv("SERVER_IP")
	var user = os.Getenv("USER")
	fmt.Println(localDirectory)

	cmd := exec.Command("git", "-C", localDirectory, "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Stderr = os.Stderr
	branch, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Current branch:", string(branch))

	// log into the remote server using ssh and update to the branch - run the post checkout script
	sshCommand := "cd " + remoteDirectory + " && git pull origin " + string(branch) + postCheckout

	fmt.Println("Logging into server:", server)
	fmt.Println("Running Command:", sshCommand)

	cmd = exec.Command("ssh", "-t", user+"@"+server, sshCommand)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
