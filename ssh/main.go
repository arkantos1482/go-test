package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"os"
	"time"
)

// Struct to store node information
type Node struct {
	Host     string
	Username string
}

func (node Node) getPass() string {
	//number := strings.TrimPrefix(node.Username, "node")
	return fmt.Sprintf("%s-blockstars@@", node.Username)
}

const (
	keyPath = "/home/abbas/.ssh/id_ed25519"
	port    = "22"
)

// SSH and execute a command on a given node
func runCommandOnNode(node Node, command string) (string, error) {
	// Read the private key from the specified file
	key, err := os.ReadFile(keyPath)
	if err != nil {
		return "", fmt.Errorf("unable to read private key: %w", err)
	}

	// Parse the private key
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return "", fmt.Errorf("unable to parse private key: %w", err)
	}

	// SSH Client configuration
	config := &ssh.ClientConfig{
		User: node.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
			ssh.Password(node.getPass()),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // WARNING: insecure, use ssh.FixedHostKey for production.
		Timeout:         10 * time.Second,
	}

	// Create an SSH connection to the node
	addr := net.JoinHostPort(node.Host, port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return "", fmt.Errorf("failed to dial: %w", err)
	}
	defer client.Close()

	// Create a new SSH session
	session, err := client.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create session: %w", err)
	}
	defer session.Close()

	// Capture the output from the command
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf

	// Execute the command
	if err := session.Run(command); err != nil {
		return "", fmt.Errorf("failed to run command: %w", err)
	}

	// Return the output
	return stdoutBuf.String(), nil
}

func main() {
	// List of nodes to SSH into
	nodes := []Node{
		{Host: "20.56.9.207", Username: "node1"},
		{Host: "74.249.178.206", Username: "node2"},
		{Host: "40.67.226.87", Username: "node3"},
		{Host: "20.196.28.88", Username: "node4"},
		{Host: "20.5.225.3", Username: "node5"},
		{Host: "20.0.96.145", Username: "node6"},
		{Host: "20.197.38.175", Username: "node7"},
		{Host: "20.164.22.2", Username: "node8"},
		{Host: "74.241.133.233", Username: "node9"},
		{Host: "20.164.18.102", Username: "node10"},
		{Host: "20.13.146.122", Username: "node11"},
		{Host: "52.231.108.190", Username: "node12"},
		{Host: "20.5.9.110", Username: "node13"},
		{Host: "20.243.8.13", Username: "node14"},
		{Host: "57.155.112.32", Username: "node15"},
	}

	// Command to run on each node
	command := "sudo rm -rf ~/btc_indexer.db"
	//command := "uptime"

	// Iterate through the list of nodes
	for _, node := range nodes {
		fmt.Printf("Connecting to %s...\n", node.Host)
		output, err := runCommandOnNode(node, command)
		if err != nil {
			log.Printf("Error on node %s: %s\n", node.Host, err)
		} else {
			fmt.Printf("Output from %s:\n%s\n", node.Host, output)
		}
	}
}
