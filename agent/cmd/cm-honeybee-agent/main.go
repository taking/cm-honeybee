package main

import (
	"embed"
	_ "embed"
	"github.com/cloud-barista/cm-honeybee/agent/common"
	"github.com/cloud-barista/cm-honeybee/agent/lib/config"
	"github.com/cloud-barista/cm-honeybee/agent/lib/privileged"
	"github.com/cloud-barista/cm-honeybee/agent/pkg/api/rest/controller"
	"github.com/cloud-barista/cm-honeybee/agent/pkg/api/rest/server"
	"github.com/jollaman999/utils/logger"
	"github.com/jollaman999/utils/syscheck"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

//go:embed .latest_version
var versionFile embed.FS

func copyVersionFile() error {
	fileContent, err := versionFile.ReadFile(".latest_version")
	if err != nil {
		return err
	}

	f, err := os.Create("/tmp/honeybee_agent_version")
	if err != nil {
		return err
	}

	_, err = f.Write(fileContent)
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}

func init() {
	err := syscheck.CheckRoot()
	if err != nil {
		log.Fatalln(err)
	}

	err = copyVersionFile()
	if err != nil {
		log.Fatal(err)
	}

	err = privileged.CheckPrivileged()
	if err != nil {
		log.Fatalln(err)
	}

	err = config.PrepareConfigs()
	if err != nil {
		log.Fatalln(err)
	}

	err = common.InitAgentUUID()
	if err != nil {
		log.Panicln(err)
	}

	err = logger.InitLogFile(common.RootPath+"/log", strings.ToLower(common.ModuleName))
	if err != nil {
		log.Panicln(err)
	}

	logger.Println(logger.INFO, false, "Agent UUID: "+common.AgentUUID)

	controller.OkMessage.Message = "API server is not ready"

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		server.Init()
	}()

	controller.OkMessage.Message = "CM-Honeybee Agent is ready"
	controller.IsReady = true

	wg.Wait()
}

func end() {
	logger.CloseLogFile()
}

func main() {
	// Catch the exit signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		logger.Println(logger.INFO, false, "Exiting "+common.ModuleName+" module...")
		end()
		os.Exit(0)
	}()
}
