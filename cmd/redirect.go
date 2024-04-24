/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// redirectCmd represents the redirect command
var redirectCmd = &cobra.Command{
	Use:   "redirect",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("redirect called")
	},
}

func init() {
	var s string
	fmt.Println("どのサイトにリダイレクトしますか？")
	fmt.Println("1: Google")
	fmt.Println("2: 琉球大学教務システム")
	fmt.Println("3: WebClass")
	fmt.Println("4: GitHub")
	fmt.Scan(&s)
	var err error
	if s == "1" {
		err = exec.Command("open", "https://google.com").Start()
	} else if s == "2" {
		err = exec.Command("open", "https://tiglon.jim.u-ryukyu.ac.jp/portal/Login.aspx").Start()
	} else if s == "3" {
		err = exec.Command("open", "https://webclass.cc.u-ryukyu.ac.jp/webclass/login.php").Start()
	} else {
		err = exec.Command("open", "https://github.com").Start()
	}
	fmt.Println(err)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// redirectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// redirectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
