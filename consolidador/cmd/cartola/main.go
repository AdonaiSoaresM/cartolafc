package main

import (
	"consolidador/internal/infra/db"
	httphandler "consolidador/internal/infra/http"
	"consolidador/internal/infra/repository"
	"consolidador/pkg/uow"
	"context"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dtb, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/cartola?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer dtb.Close()
	uow, err := uow.NewUow(ctx, dtb)
	if err != nil {
		panic(err)
	}
	registerRepositories(uow)

	r := chi.NewRouter()
	r.Get("/players", httphandler.ListPlayersHandler(ctx, *db.New(dtb)))
	r.Get("/my-teams/{teamID}/players", httphandler.ListMyTeamPlayersHandler(ctx, *db.New(dtb)))
	r.Get("/my-teams/{teamID}/balance", httphandler.GetMyTeamBalanceHandler(ctx, *db.New(dtb)))
	r.Get("/matches", httphandler.ListMatchesHandler(ctx, repository.NewMatchRepository(dtb)))
	r.Get("/matches/{matchID}", httphandler.ListMatchByIDHandler(ctx, repository.NewMatchRepository(dtb)))

	//go http.ListenAndServe(":8081", r)
	if err = http.ListenAndServe(":8081", r); err != nil {
		panic(err)
	}
}

func registerRepositories(uow *uow.Uow) {
	uow.Register("PlayerRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewPlayerRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("MatchRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewMatchRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("TeamRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewTeamRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("MyTeamRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewMyTeamRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})
}
