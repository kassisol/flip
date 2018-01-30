// Flip is an application to manage a Floating IP address
// on a Docker Swarm cluster.
// Copyright (C) 2018 Kassisol inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	cmdDatasource        string
	cmdDatasourceOptions string
	cmdDebug             bool
	cmdInterface         string
	cmdKeepalive         int
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "flip",
		Short: "Manage Floating IP",
		Long:  "Flip runs as a Docker container (with special capabilities) on each Docker Swarm manager host and set or unset the Floating IP.",
		Run:   runApp,
	}

	rootCmd.Flags().StringVarP(&cmdDatasource, "datasource", "d", "", "Select datasource")
	rootCmd.Flags().StringVarP(&cmdDatasourceOptions, "datasource-opts", "o", "", "Datasource options")
	rootCmd.Flags().BoolVarP(&cmdDebug, "debug", "D", false, "Enable debug mode")
	rootCmd.Flags().StringVarP(&cmdInterface, "interface", "i", "eth0", "Select interface to set Floating IP")
	rootCmd.Flags().IntVarP(&cmdKeepalive, "keepalive", "k", 5, "Set how many seconds between heartbeats")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
