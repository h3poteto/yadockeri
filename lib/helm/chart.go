package helm

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"strings"
	"time"

	"gopkg.in/src-d/go-git.v4"
)

func GitClone(url string, userName, token string) (string, error) {
	authURL := strings.Replace(url, "https://", "https://"+userName+":"+token+"@", 1)

	now := time.Now()
	hasher := md5.New()
	hasher.Write([]byte(now.String()))
	downloaded := "/tmp/" + hex.EncodeToString(hasher.Sum(nil))
	_, err := git.PlainClone(downloaded, false, &git.CloneOptions{
		URL:      authURL,
		Progress: os.Stdout,
	})
	if err != nil {
		return "", err
	}
	return downloaded, nil
}
