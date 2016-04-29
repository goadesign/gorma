package main

import (
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	"github.com/goadesign/gorma/example/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"os"
)

type (
	// CallbackAuthCommand is the command line data structure for the callback action of auth
	CallbackAuthCommand struct {
		Provider string
	}
	// OauthAuthCommand is the command line data structure for the oauth action of auth
	OauthAuthCommand struct {
		Provider string
	}
	// RefreshAuthCommand is the command line data structure for the refresh action of auth
	RefreshAuthCommand struct {
		Payload string
	}
	// TokenAuthCommand is the command line data structure for the token action of auth
	TokenAuthCommand struct {
		Payload string
	}
	// CreateProposalCommand is the command line data structure for the create action of proposal
	CreateProposalCommand struct {
		Payload string
		UserID  int
	}
	// DeleteProposalCommand is the command line data structure for the delete action of proposal
	DeleteProposalCommand struct {
		// Proposal ID
		ProposalID int
		UserID     int
	}
	// ListProposalCommand is the command line data structure for the list action of proposal
	ListProposalCommand struct {
		UserID int
	}
	// ShowProposalCommand is the command line data structure for the show action of proposal
	ShowProposalCommand struct {
		ProposalID int
		UserID     int
	}
	// UpdateProposalCommand is the command line data structure for the update action of proposal
	UpdateProposalCommand struct {
		Payload    string
		ProposalID int
		UserID     int
	}
	// CreateReviewCommand is the command line data structure for the create action of review
	CreateReviewCommand struct {
		Payload    string
		ProposalID int
		UserID     int
	}
	// DeleteReviewCommand is the command line data structure for the delete action of review
	DeleteReviewCommand struct {
		ProposalID int
		// Review ID
		ReviewID int
		UserID   int
	}
	// ListReviewCommand is the command line data structure for the list action of review
	ListReviewCommand struct {
		ProposalID int
		UserID     int
	}
	// ShowReviewCommand is the command line data structure for the show action of review
	ShowReviewCommand struct {
		ProposalID int
		ReviewID   int
		UserID     int
	}
	// UpdateReviewCommand is the command line data structure for the update action of review
	UpdateReviewCommand struct {
		Payload    string
		ProposalID int
		ReviewID   int
		UserID     int
	}
	// BootstrapUICommand is the command line data structure for the bootstrap action of ui
	BootstrapUICommand struct {
	}
	// CreateUserCommand is the command line data structure for the create action of user
	CreateUserCommand struct {
		Payload string
	}
	// DeleteUserCommand is the command line data structure for the delete action of user
	DeleteUserCommand struct {
		// User ID
		UserID int
	}
	// ListUserCommand is the command line data structure for the list action of user
	ListUserCommand struct {
	}
	// ShowUserCommand is the command line data structure for the show action of user
	ShowUserCommand struct {
		UserID int
	}
	// UpdateUserCommand is the command line data structure for the update action of user
	UpdateUserCommand struct {
		Payload string
		UserID  int
	}
)

// Run makes the HTTP request corresponding to the CallbackAuthCommand command.
func (cmd *CallbackAuthCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/auth/%v/callback", cmd.Provider)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CallbackAuth(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CallbackAuthCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var provider string
	cc.Flags().StringVar(&cmd.Provider, "provider", provider, ``)
}

// Run makes the HTTP request corresponding to the OauthAuthCommand command.
func (cmd *OauthAuthCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/auth/%v", cmd.Provider)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.OauthAuth(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *OauthAuthCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var provider string
	cc.Flags().StringVar(&cmd.Provider, "provider", provider, ``)
}

// Run makes the HTTP request corresponding to the RefreshAuthCommand command.
func (cmd *RefreshAuthCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/auth/refresh"
	}
	var payload client.RefreshAuthPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.RefreshAuth(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *RefreshAuthCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the TokenAuthCommand command.
func (cmd *TokenAuthCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/auth/token"
	}
	var payload client.TokenAuthPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.TokenAuth(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *TokenAuthCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the CreateProposalCommand command.
func (cmd *CreateProposalCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/users/%v/proposals", cmd.UserID)
	}
	var payload client.CreateProposalPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateProposal(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateProposalCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, ``)
}

// Run makes the HTTP request corresponding to the DeleteProposalCommand command.
func (cmd *DeleteProposalCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/users/%v/proposals/%v", cmd.ProposalID, cmd.UserID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteProposal(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteProposalCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var proposalID int
	cc.Flags().IntVar(&cmd.ProposalID, "proposalID", proposalID, `Proposal ID`)
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, ``)
}

// Run makes the HTTP request corresponding to the ListProposalCommand command.
func (cmd *ListProposalCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/users/%v/proposals", cmd.UserID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListProposal(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListProposalCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, ``)
}

// Run makes the HTTP request corresponding to the ShowProposalCommand command.
func (cmd *ShowProposalCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/users/%v/proposals/%v", cmd.ProposalID, cmd.UserID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowProposal(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowProposalCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var proposalID int
	cc.Flags().IntVar(&cmd.ProposalID, "proposalID", proposalID, ``)
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, ``)
}

// Run makes the HTTP request corresponding to the UpdateProposalCommand command.
func (cmd *UpdateProposalCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/users/%v/proposals/%v", cmd.ProposalID, cmd.UserID)
	}
	var payload client.UpdateProposalPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateProposal(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateProposalCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
	var proposalID int
	cc.Flags().IntVar(&cmd.ProposalID, "proposalID", proposalID, ``)
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, ``)
}

// Run makes the HTTP request corresponding to the CreateReviewCommand command.
func (cmd *CreateReviewCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/users/%v/proposals/%v/review", cmd.ProposalID, cmd.UserID)
	}
	var payload client.CreateReviewPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateReview(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateReviewCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
	var proposalID int
	cc.Flags().IntVar(&cmd.ProposalID, "proposalID", proposalID, ``)
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, ``)
}

// Run makes the HTTP request corresponding to the DeleteReviewCommand command.
func (cmd *DeleteReviewCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/users/%v/proposals/%v/review/%v", cmd.ProposalID, cmd.ReviewID, cmd.UserID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteReview(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteReviewCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var proposalID int
	cc.Flags().IntVar(&cmd.ProposalID, "proposalID", proposalID, ``)
	var reviewID int
	cc.Flags().IntVar(&cmd.ReviewID, "reviewID", reviewID, `Review ID`)
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, ``)
}

// Run makes the HTTP request corresponding to the ListReviewCommand command.
func (cmd *ListReviewCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/users/%v/proposals/%v/review", cmd.ProposalID, cmd.UserID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListReview(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListReviewCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var proposalID int
	cc.Flags().IntVar(&cmd.ProposalID, "proposalID", proposalID, ``)
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, ``)
}

// Run makes the HTTP request corresponding to the ShowReviewCommand command.
func (cmd *ShowReviewCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/users/%v/proposals/%v/review/%v", cmd.ProposalID, cmd.ReviewID, cmd.UserID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowReview(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowReviewCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var proposalID int
	cc.Flags().IntVar(&cmd.ProposalID, "proposalID", proposalID, ``)
	var reviewID int
	cc.Flags().IntVar(&cmd.ReviewID, "reviewID", reviewID, ``)
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, ``)
}

// Run makes the HTTP request corresponding to the UpdateReviewCommand command.
func (cmd *UpdateReviewCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/users/%v/proposals/%v/review/%v", cmd.ProposalID, cmd.ReviewID, cmd.UserID)
	}
	var payload client.UpdateReviewPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateReview(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateReviewCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
	var proposalID int
	cc.Flags().IntVar(&cmd.ProposalID, "proposalID", proposalID, ``)
	var reviewID int
	cc.Flags().IntVar(&cmd.ReviewID, "reviewID", reviewID, ``)
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, ``)
}

// Run makes the HTTP request corresponding to the BootstrapUICommand command.
func (cmd *BootstrapUICommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.BootstrapUI(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *BootstrapUICommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the CreateUserCommand command.
func (cmd *CreateUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/users"
	}
	var payload client.CreateUserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateUser(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the DeleteUserCommand command.
func (cmd *DeleteUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/users/%v", cmd.UserID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteUser(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, `User ID`)
}

// Run makes the HTTP request corresponding to the ListUserCommand command.
func (cmd *ListUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/users"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListUser(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the ShowUserCommand command.
func (cmd *ShowUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/users/%v", cmd.UserID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowUser(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, ``)
}

// Run makes the HTTP request corresponding to the UpdateUserCommand command.
func (cmd *UpdateUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/users/%v", cmd.UserID)
	}
	var payload client.UpdateUserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateUser(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, ``)
}
