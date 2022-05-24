package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Options struct {
	user     string
	password string
	host     string
	port     int
	database string
	dir      string
}

var (
	options = &Options{}
)

var rootCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A database migration tool.",
}

func Execute() {
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(runCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&options.user, "user", "u", "root", "database username")
	rootCmd.PersistentFlags().StringVarP(&options.password, "password", "p", "", "database password")
	viper.BindEnv("password", "DB_PASSWORD")
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	rootCmd.PersistentFlags().StringVarP(&options.host, "host", "h", "127.0.0.1", "database hostname")
	rootCmd.PersistentFlags().IntVarP(&options.port, "port", "P", 3306, "database port")
	rootCmd.PersistentFlags().StringVarP(&options.dir, "dir", "D", "./migrations", "directory that migration files exist")
	rootCmd.PersistentFlags().BoolP("help", "", false, "help for this command")
}
