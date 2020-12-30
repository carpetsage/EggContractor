package cmd

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/fanaticscripter/EggContractor/db"
	"github.com/fanaticscripter/EggContractor/util"
)

var _statusCommand = &cobra.Command{
	Use:   "status",
	Short: "Print statuses of active solo contracts & coops from last refresh",
	Long: `Print statuses of active solo contracts & coops from last refresh

This command does not touch the network, only the database.`,
	Args:    cobra.NoArgs,
	PreRunE: subcommandPreRunE,
	RunE: func(cmd *cobra.Command, args []string) error {
		timestamp, solos, coops, err := db.GetSoloAndCoopStatusesFromRefresh(time.Now())
		if err != nil {
			return err
		}
		if timestamp.IsZero() {
			log.Warn("no refresh found in the database, try using the refresh subcommand")
			return nil
		}
		fmt.Printf("Last refreshed at %s (%s)\n\n",
			util.FormatDatetime(timestamp), humanize.Time(timestamp))
		if len(solos) == 0 && len(coops) == 0 {
			fmt.Println(util.MsgNoActiveContracts)
		}
		for _, solo := range solos {
			solo.Display()
		}
		for _, coop := range coops {
			coop.Display(_sortBy.by)
		}
		return nil
	},
}

func init() {
	_rootCmd.AddCommand(_statusCommand)
}
