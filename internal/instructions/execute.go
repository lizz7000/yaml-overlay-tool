// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: MIT

package instructions

import (
	"context"
	"fmt"

	"github.com/op/go-logging"
	"github.com/vmware-tanzu-labs/yaml-overlay-tool/internal/overlays"
	"golang.org/x/sync/errgroup"
)

var log = logging.MustGetLogger("instructions") //nolint:gochecknoglobals

func Execute(cfg *Config) error {
	eg, ctx := errgroup.WithContext(context.Background())

	instructions, err := ReadInstructionFile(&cfg.InstructionsFile)
	if err != nil {
		return err
	}

	instructions.addCommonOverlays()

	pChan := make(chan *YamlFile, len(instructions.YamlFiles))

	for _, yamlFile := range instructions.YamlFiles {
		yf := yamlFile

		eg.Go(
			func() error {
				stream := overlays.NewWorkStream()

				stream.StartStream()

				go yf.queueOverlays(stream)

				if err := stream.StartHandler(); err != nil {
					return fmt.Errorf("%w", err)
				}

				select {
				case pChan <- yf:
					return nil
				case <-ctx.Done():
					if err := ctx.Err(); err != nil {
						return fmt.Errorf("%w", err)
					}

					return nil
				}
			},
		)
	}

	go func() {
		eg.Wait() //nolint:errcheck // we use this to gracefully close the channel and receive the err later
		close(pChan)
	}()

	if err := PostProcessHandler(cfg, pChan); err != nil {
		return err
	}

	if err := eg.Wait(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func PostProcessHandler(cfg *Config, pChan <-chan *YamlFile) error {
	for yf := range pChan {
		if err := yf.doPostProcessing(cfg); err != nil {
			return err
		}
	}

	return nil
}