/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// redCmd represents the red command
var redCmd = &cobra.Command{
	Use:   "red",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("red called")
	},
}

func init() {
	var s string
	fmt.Println("どのサイトにリダイレクトしますか？")
	fmt.Println("1: Google")
	fmt.Println("2: 琉球大学教務システム")
	fmt.Println("3: WebClass")
	fmt.Println("4: GitHub")
	fmt.Println("5: AtCoder")
	fmt.Scan(&s)
	var err error
	if s == "1" {
		err = exec.Command("open", "https://google.com").Start()
	} else if s == "2" {
		err = exec.Command("open", "https://tiglon.jim.u-ryukyu.ac.jp/portal/Login.aspx").Start()
	} else if s == "3" {
		err = exec.Command("open", "https://webclass.cc.u-ryukyu.ac.jp/webclass/login.php").Start()
	} else if s == "4" {
		var user string
		fmt.Printf("\n検索したいユーザー名を入力してください。\n無ければEnterをクリックしてください。\n")
		fmt.Scanln(&user)
		err = exec.Command("open", "https://github.com/"+user).Start()
	} else {
		var user string
		fmt.Printf("\n検索したいユーザー名を入力してください。\n無ければEnterをクリックしてください。\n")
		fmt.Scanln(&user)
		if user != "" {
			err = exec.Command("open", "https://atcoder.jp/users/"+user).Start()
		} else {
			err = exec.Command("open", "https://atcoder.jp/").Start()
		}
	}
	fmt.Println(err)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// redCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// redCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
