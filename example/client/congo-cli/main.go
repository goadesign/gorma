package main

import (
	"fmt"
	"github.com/goadesign/gorma/example/client"
	"github.com/spf13/cobra"
	"os"
	"time"
)

// PrettyPrint is true if the tool output should be formatted for human consumption.
var PrettyPrint bool

func main() {
	// Create command line parser
	app := &cobra.Command{
		Use:   "congo-cli",
		Short: `CLI client for the congo service (https://gopheracademy.com/congo/getting-started.html)`,
	}
	c := client.New(nil)
	c.UserAgent = "congo-cli/1.0"
	app.PersistentFlags().StringVarP(&c.Scheme, "scheme", "s", "", "Set the requests scheme")
	app.PersistentFlags().StringVarP(&c.Host, "host", "H", "api.gopheracademy.com", "API hostname")
	app.PersistentFlags().DurationVarP(&c.Timeout, "timeout", "t", time.Duration(20)*time.Second, "Set the request timeout")
	app.PersistentFlags().BoolVar(&c.Dump, "dump", false, "Dump HTTP request and response.")
	app.PersistentFlags().BoolVar(&PrettyPrint, "pp", false, "Pretty print response body")
	RegisterCommands(app, c)
	if err := app.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "request failed: %s", err)
		os.Exit(-1)
	}
}

// RegisterCommands all the resource action subcommands to the application command line.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "bootstrap",
		Short: `Render single page app HTML`,
	}
	tmp1 := new(BootstrapUICommand)
	sub = &cobra.Command{
		Use:   `ui ["/"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "callback",
		Short: `OAUTH2 callback endpoint`,
	}
	tmp2 := new(CallbackAuthCommand)
	sub = &cobra.Command{
		Use:   `auth ["/api/auth/:provider/callback"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "create",
		Short: `create action`,
	}
	tmp3 := new(CreateProposalCommand)
	sub = &cobra.Command{
		Use:   `proposal ["/api/users/:userID/proposals"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp4 := new(CreateReviewCommand)
	sub = &cobra.Command{
		Use:   `review ["/api/users/:userID/proposals/:proposalID/review"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp5 := new(CreateUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/api/users"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "delete",
		Short: `delete action`,
	}
	tmp6 := new(DeleteProposalCommand)
	sub = &cobra.Command{
		Use:   `proposal ["/api/users/:userID/proposals/:proposalID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp7 := new(DeleteReviewCommand)
	sub = &cobra.Command{
		Use:   `review ["/api/users/:userID/proposals/:proposalID/review/:reviewID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp7.Run(c, args) },
	}
	tmp7.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp8 := new(DeleteUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/api/users/:userID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp8.Run(c, args) },
	}
	tmp8.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "list",
		Short: `list action`,
	}
	tmp9 := new(ListProposalCommand)
	sub = &cobra.Command{
		Use:   `proposal ["/api/users/:userID/proposals"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp9.Run(c, args) },
	}
	tmp9.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp10 := new(ListReviewCommand)
	sub = &cobra.Command{
		Use:   `review ["/api/users/:userID/proposals/:proposalID/review"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp10.Run(c, args) },
	}
	tmp10.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp11 := new(ListUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/api/users"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp11.Run(c, args) },
	}
	tmp11.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "oauth",
		Short: `OAUTH2 login endpoint`,
	}
	tmp12 := new(OauthAuthCommand)
	sub = &cobra.Command{
		Use:   `auth ["/api/auth/:provider"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp12.Run(c, args) },
	}
	tmp12.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "refresh",
		Short: `Obtain a refreshed access token`,
	}
	tmp13 := new(RefreshAuthCommand)
	sub = &cobra.Command{
		Use:   `auth ["/api/auth/refresh"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp13.Run(c, args) },
	}
	tmp13.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "show",
		Short: `show action`,
	}
	tmp14 := new(ShowProposalCommand)
	sub = &cobra.Command{
		Use:   `proposal ["/api/users/:userID/proposals/:proposalID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp14.Run(c, args) },
	}
	tmp14.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp15 := new(ShowReviewCommand)
	sub = &cobra.Command{
		Use:   `review ["/api/users/:userID/proposals/:proposalID/review/:reviewID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp15.Run(c, args) },
	}
	tmp15.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp16 := new(ShowUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/api/users/:userID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp16.Run(c, args) },
	}
	tmp16.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "token",
		Short: `Obtain an access token`,
	}
	tmp17 := new(TokenAuthCommand)
	sub = &cobra.Command{
		Use:   `auth ["/api/auth/token"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp17.Run(c, args) },
	}
	tmp17.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "update",
		Short: `update action`,
	}
	tmp18 := new(UpdateProposalCommand)
	sub = &cobra.Command{
		Use:   `proposal ["/api/users/:userID/proposals/:proposalID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp18.Run(c, args) },
	}
	tmp18.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp19 := new(UpdateReviewCommand)
	sub = &cobra.Command{
		Use:   `review ["/api/users/:userID/proposals/:proposalID/review/:reviewID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp19.Run(c, args) },
	}
	tmp19.RegisterFlags(sub, c)
	command.AddCommand(sub)
	tmp20 := new(UpdateUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/api/users/:userID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp20.Run(c, args) },
	}
	tmp20.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)

}
