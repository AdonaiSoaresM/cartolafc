package usecase

import (
	"context"
	"strconv"
	"strings"

	"consolidador/internal/domain/entity"
	"consolidador/internal/domain/repository"
	"consolidador/pkg/uow"
)

type MatchUpdateResultInput struct {
	ID     string
	Result string
}

type MatchUpdateResultUseCase struct {
	Uow uow.UowInterface
}

func NewMatchUpdateResultUseCase(uow uow.UowInterface) *MatchUpdateResultUseCase {
	return &MatchUpdateResultUseCase{
		Uow: uow,
	}
}

func (u *MatchUpdateResultUseCase) Execute(ctx context.Context, input MatchUpdateResultInput) error {
	err := u.Uow.Do(ctx, func(_ *uow.Uow) error {
		matchRepo := u.getMatchRepository(ctx)
		match, err := matchRepo.FindByID(ctx, input.ID)
		if err != nil {
			return err
		}
		matchResult := strings.Split(input.Result, "-")
		// convert results to int
		teamAResult, _ := strconv.Atoi(matchResult[0])
		teamBResult, _ := strconv.Atoi(matchResult[1])
		match.Result = *entity.NewMatchResult(teamAResult, teamBResult)

		err = matchRepo.Update(ctx, match)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (u *MatchUpdateResultUseCase) getMatchRepository(ctx context.Context) repository.MatchRepositoryInterface {
	matchRepository, err := u.Uow.GetRepository(ctx, "MatchRepository")
	if err != nil {
		panic(err)
	}
	return matchRepository.(repository.MatchRepositoryInterface)
}
