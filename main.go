package main

import (
  "fmt"
  "github.com/urfave/cli/v2"
  "k8s.io/client-go/kubernetes"
  "log"
  "os"
  "sort"
)
var client *kubernetes.Clientset
func main() {
  app := &cli.App{
    Flags: []cli.Flag{
      &cli.StringFlag{
        Name:  "output",
        Aliases: []string{"o"},
        //Value: "table",
        //Usage: "Output to display [json, table]",
      },
    },
    Commands: []*cli.Command{
      {
        Name:    "user-roles",
        Usage:   "oc-roles User-roles [username] \n Gets the Roles for a User",
        ArgsUsage:   "[username]",
        //Before: func(c *cli.Context) error {
        //  client = k8sConfig()
        //  return nil
        //},
        Action:  func(c *cli.Context) error {
          fmt.Println(c.String("output"))
          client = k8sConfig()
          if c.Args().Len() < 1 {
            return cli.Exit("Username not provided", 1)
          }
          UserToRoles(client, c.Args().Get(0), c.String("output"))
          return nil
        },
      },
      {
        Name:    "roles-user",
        Usage:   "oc-roles roles-User [Role] \n Gets the users for a Role",
        ArgsUsage:   "[username]",
        Before: func(c *cli.Context) error {
          client = k8sConfig()
          return nil
        },
        Action:  func(c *cli.Context) error {
          if c.Args().Len() < 1 {
            return cli.Exit("Role not provided", 1)
          }
          RolesToUser(client, c.Args().Get(0), c.String("output"))
          return nil
        },
      },
    },
  }

  sort.Sort(cli.FlagsByName(app.Flags))
  sort.Sort(cli.CommandsByName(app.Commands))

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}

