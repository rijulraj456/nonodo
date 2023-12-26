// Copyright (c) Gabriel de Quadros Ligneul
// SPDX-License-Identifier: Apache-2.0 (see LICENSE)

package foundry

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path"

	"github.com/gligneul/nonodo/internal/supervisor"
)

// Default port for the Ethereum node.
const AnvilDefaultPort = 8545

// Start the anvil process in the host machine.
type AnvilWorker struct {
	Port    int
	Verbose bool
}

func (w AnvilWorker) String() string {
	return "anvil"
}

func (w AnvilWorker) Start(ctx context.Context, ready chan<- struct{}) error {
	// create temp dir
	tempDir, err := os.MkdirTemp("", "")
	if err != nil {
		return fmt.Errorf("anvil: failed to create temp dir: %w", err)
	}
	defer func() {
		err := os.RemoveAll(tempDir)
		if err != nil {
			slog.Warn("anvil: failed to remove temp file", "error", err)
		}
	}()

	// create state file in temp dir
	stateFile := path.Join(tempDir, "anvil_state.json")
	const permissions = 0644
	err = os.WriteFile(stateFile, devnetState, permissions)
	if err != nil {
		return fmt.Errorf("anvil: failed to write state file: %w", err)
	}

	// start server worker
	args := []string{
		"--port", fmt.Sprint(w.Port),
		"--load-state", stateFile,
	}
	if !w.Verbose {
		args = append(args, "--silent")
	}
	var server supervisor.ServerWorker
	server.Name = "anvil"
	server.Command = "anvil"
	server.Args = args
	server.Port = w.Port
	return server.Start(ctx, ready)
}
