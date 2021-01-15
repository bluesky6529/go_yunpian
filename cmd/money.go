package cmd

import (
	"fmt"
	_ "yunpian/configs"
	"yunpian/internal/Money"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var money = &cobra.Command{
	Use:   "Money",
	Short: "查詢帳戶餘額",
	Long:  "用來查詢帳戶餘額使用方法如下 : \nQueryAccountBalance -a <帳戶名稱> ex: QueryAccountBalance -a account",
	Run: func(cmd *cobra.Command, args []string) {
		var Apikey = viper.GetString(account + ".apikey")

		request := Money.Money(Apikey)
		fmt.Printf("%s 帳戶餘額: %f \n", account, request)
	},
}

func init() {
	money.Flags().StringVarP(&account, "account", "a", "", "帳號 (require)")
	money.MarkFlagRequired("account")
}
