package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "Poder",
	Short:            "Poder is a like docker",
	Long:             `Poder is a self to learn docker by write it`,
	TraverseChildren: true,
}

func Execute() {
	rand.Seed(time.Now().UnixNano())

	// 需要使用 chroot和写特权文件，需要使用管理员权限
	if os.Geteuid() != 0 {
		log.Fatal("You need root privileges to run this program.")
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
