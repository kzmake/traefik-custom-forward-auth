package main

import (
	"golang.org/x/xerrors"

	"github.com/kzmake/traefik-custom-forward-auth/pkg/logger"

	"github.com/kzmake/traefik-custom-forward-auth/service/randauth/infrastructure/api"
)

func main() {
	s := api.New()

	if err := s.Run(); err != nil {
		logger.Errorf("%+v", xerrors.Errorf("server の起動に失敗しました: %w", err))
	}
}
