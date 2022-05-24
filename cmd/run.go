package cmd

import (
	"fmt"
	"log"
	"migrate/util"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run migration.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "No database specified.")
			return
		}
		options.database = args[0]

		db, err := util.InitDB(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", options.user, viper.GetString("password"), options.host, options.port, options.database))
		if err != nil {
			log.Fatalln(err)
		}

		filesExecute, err := util.GetFilesExecute(options.dir, db)
		if err != nil {
			log.Fatalln(err)
		}

		if len(filesExecute) == 0 {
			fmt.Println("There's no need to migrate.")
			return
		} else {
			fmt.Printf("Following files will be executed: %s\n", filesExecute)
			err = util.ExecMigration(db, options.dir, filesExecute)
			if err != nil {
				log.Fatalln("Migration failed:", err)
			}
		}

		fmt.Println("Migration has been completed.")
	},
}
