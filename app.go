package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kassisol/flip/datasource"
	"github.com/kassisol/flip/datasource/driver"
	"github.com/kassisol/flip/pkg/docker"
	"github.com/kassisol/flip/pkg/ip"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func runApp(cmd *cobra.Command, args []string) {
	var i *driver.IP

	if cmdDebug {
		log.SetLevel(log.DebugLevel)
	}

	d, err := datasource.NewDriver(cmdDatasource, cmdDatasourceOptions)
	if err != nil {
		log.Fatal(err)
	}

	for {
		if err := d.IsAvailable(); err != nil {
			log.Error(err)
		}

		i, err = d.GetIP()
		if err != nil {
			log.Error(err)
			continue
		}

		break
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)

	ticker := time.NewTicker(time.Second * time.Duration(cmdKeepalive))
	go func(nic string, ipaddr *driver.IP) {
		for _ = range ticker.C {
			d, err := docker.NewDockerClient()
			if err != nil {
				log.Error(err)
				continue
			}

			if err := d.GetNodeID(); err != nil {
				log.Error(err)
				continue
			}

			service, err := d.IsServiceRunning()
			if err != nil {
				log.Error(err)
				continue
			}

			if service {
				// Check if IP is already set
				i := ip.NewIP(nic, ipaddr.Address, ipaddr.Netmask)
				if i.IsSet() {
					log.Debugf("IP address '%s' is already configured", ipaddr.Address)
					continue
				}

				// Ping IP
				if i.Ping() {
					log.Debugf("IP address '%s' is reachable and so seems to be already set on another node", ipaddr.Address)
					continue
				}

				// If IP is not set then add it
				if err := i.Set(); err != nil {
					log.Warning(err)
					continue
				}

				log.Infof("IP address '%s' has been set", ipaddr.Address)
			} else {
				// Check if IP is already set
				i := ip.NewIP(nic, ipaddr.Address, ipaddr.Netmask)
				if i.IsSet() {
					log.Debugf("IP address '%s' is already configured", ipaddr.Address)

					// If IP is set then delete it
					if err := i.Unset(); err != nil {
						log.Warning(err)
						continue
					}

					log.Infof("IP address '%s' has been unset", ipaddr.Address)
				}
			}
		}
	}(cmdInterface, i)

	s := <-ch
	ticker.Stop()
	log.Infof("Processing signal '%s'", s)
}
