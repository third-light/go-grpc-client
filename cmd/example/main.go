package main

import (
	"context"
	"flag"
	"log"

	"github.com/third-light/go-grpc-client/chorus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	host   = flag.String("hostname", "", "Hostname of Chorus Site")
	user   = flag.String("username", "", "Username")
	pass   = flag.String("password", "", "Password")
	apiKey = flag.String("api-key", "", "API-Key")
)

type Session struct {
	Session string
}

func (s *Session) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"session": s.Session,
	}, nil
}

func (s *Session) RequireTransportSecurity() bool {
	return false
}

func main() {
	// Remove log timestamp
	log.SetFlags(0)

	// Parse flags
	flag.Parse()

	// Check for non-empty Chorus address
	if *host == "" {
		log.Fatal("You must specify a Chorus hostname")
	}
	grpcAddr := *host + ":443"

	// Check for either user+password auth, or api-key auth
	if *apiKey == "" && (*user == "" || *pass == "") {
		log.Fatalf("You must provide either an `api-key` or a `username` + `password` for authentication")
	}
	userPassAuth := (*apiKey == "")

	// Dial new GRPC conn
	creds := credentials.NewTLS(nil)
	session := &Session{}
	conn, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(session))
	if err != nil {
		log.Fatalf("Error connecting to chorus instance: %s", err.Error())
	}

	// Create new Chorus client from connection
	client := chorus.NewClient(conn)

	// Authenticate this session
	var loginDetails *chorus.LoginDetails
	if userPassAuth {
		// Create new login request object
		loginRequest := &chorus.AuthRequestsLogin{
			Username: *user,
			Password: *pass,
		}

		// Attempt to login
		loginDetails, err = client.Auth().Login(loginRequest)
		if err != nil {
			log.Fatalf("Failed to authenticate with Chorus as user %s: %s", *user, err.Error())
		}
		session.Session = loginDetails.SessionId
	} else {
		// Create new login request object
		loginRequest := &chorus.AuthRequestsLoginWithKey{
			ApiKey: *apiKey,
		}

		// Attempt to login
		loginDetails, err = client.Auth().LoginWithKey(loginRequest)
		if err != nil {
			log.Fatalf("Failed to authenticate with Chorus using supplied API key: %s", err.Error())
		}
		session.Session = loginDetails.SessionId
	}
	log.Printf("Authenticated with chorus as: %s, %s", loginDetails.UserDetails.Name, loginDetails.UserDetails.Username)
	log.Println()

	// Fetch contexts available to current user
	contextChildren, err := client.Contexts().GetChildren(&chorus.ContextsRequestsGetChildren{ContextId: "dom0"})
	if err != nil {
		log.Fatalf("Failed to retrieve context children: %s", err.Error())
	}

	for _, contextId := range contextChildren.Response {
		// Fetch further context information
		contextDetails, err := client.Contexts().Get(&chorus.ContextsRequestsGet{ContextId: contextId})
		if err != nil {
			log.Fatalf("Failed to retrieve context details: %s", err.Error())
		}
		log.Printf("Context: %s", contextDetails.Name)

		// If no backing folder id, nothing to do here
		if contextDetails.BackingFolderId == nil {
			continue
		}

		// Get backing folder children
		files, err := client.Folders().GetChildFiles(&chorus.FoldersRequestsGetChildFiles{FolderId: *contextDetails.BackingFolderId})
		if err != nil {
			log.Fatalf("Failed to retrieve child files for folder: %s", err.Error())
		}
		folders, err := client.Folders().GetChildFolders(&chorus.FoldersRequestsGetChildFolders{FolderId: *contextDetails.BackingFolderId})
		if err != nil {
			log.Fatalf("Failed to retrieve child folders for folder: %s", err.Error())
		}

		// Print these files and folders
		for _, folder := range folders.Response {
			log.Println("[DIR] ", folder.Name)
		}
		for _, file := range files.Response {
			log.Println("[FILE]", file.Filename)
		}
		log.Println()
	}
}
