package auth

import (
	"net/http"
	"fmt"
	"strings"
	//"os"
)

type Auth struct {
	ServerUrl string
}

func Register(login string, serverUrl string) error {
	res, err := http.Post(serverUrl + "/users", "text/plain", strings.NewReader(login))
	if (err != nil) {
		fmt.Printf("error : %s\n", err)
		return err
	}
	fmt.Printf("result : %s\n", res.Status)
	return nil
}
