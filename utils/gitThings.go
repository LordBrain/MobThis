package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

func GitCloneSSH(url, path, key string) error {
	if url == "" || path == "" || key == "" {
		return errors.New("Missing config value")
	}
	fmt.Println("Mobbing code dir:", path)
	//Create mobbing directory if it doesn't exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(filepath.Dir(path), 0770)
	}
	//Create mobbing subdirectory if it doesn't exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0770)
	}

	var publicKey *ssh.PublicKeys

	sshKey, _ := ioutil.ReadFile(key)
	publicKey, keyError := ssh.NewPublicKeys("git", []byte(sshKey), "")
	if keyError != nil {
		fmt.Println("key error:", keyError)
		return keyError
	}
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
		Auth:     publicKey,
	})
	if err != nil {
		return err
	}
	return nil
}

func GitCloneAuth(url, path, username, token string) error {

	if url == "" || path == "" || username == "" || token == "" {
		return errors.New("Missing config value")
	}
	fmt.Println("Mobbing code dir:", path)
	//Create mobbing directory if it doesn't exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(filepath.Dir(path), 0770)
	}
	//Create mobbing subdirectory if it doesn't exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0770)
	}

	_, err := git.PlainClone(path, false, &git.CloneOptions{
		// The intended use of a GitHub personal access token is in replace of your password
		// because access tokens can easily be revoked.
		// https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/
		Auth: &http.BasicAuth{
			Username: username, // yes, this can be anything except an empty string
			Password: token,
		},
		URL:      url,
		Progress: os.Stdout,
	})
	if err != nil {
		return err
	}

	return nil
}

func GitBranch() {
	// r, err := git.PlainClone(directory, false, &git.CloneOptions{
	// 	URL: url,
	// })
	// CheckIfError(err)

	// // Create a new branch to the current HEAD
	// Info("git branch my-branch")

	// headRef, err := r.Head()
	// CheckIfError(err)

	// // Create a new plumbing.HashReference object with the name of the branch
	// // and the hash from the HEAD. The reference name should be a full reference
	// // name and not an abbreviated one, as is used on the git cli.
	// //
	// // For tags we should use `refs/tags/%s` instead of `refs/heads/%s` used
	// // for branches.
	// ref := plumbing.NewHashReference("refs/heads/my-branch", headRef.Hash())

	// // The created reference is saved in the storage.
	// err = r.Storer.SetReference(ref)
	// CheckIfError(err)

}
