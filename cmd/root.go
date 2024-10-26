package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Esta aplicación te ayuda a gestionar tus pendientes",
	Long:  "Aplicación contruida en Go para gestionar las pendientes del usuario",
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
