package web

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/seetohjinwei/repostats/models"
	"github.com/seetohjinwei/repostats/postgres"
)

func Start(pool *pgxpool.Pool) {
	// TODO: quick testing code

	// postgres.AddUser(pool, "seetohjinwei")
	postgres.AddRepository(pool, "seetohjinwei", "adventofcode", "main")
	postgres.AddRepository(pool, "seetohjinwei", "repostats", "main")

	d1 := []models.TypeData{
		{"java", 1, 420},
		{"python", 10, 69},
	}
	d2 := []models.TypeData{
		{"go", 5, 1234},
	}
	postgres.AddTypeData(pool, "seetohjinwei", "adventofcode", d1)
	postgres.AddTypeData(pool, "seetohjinwei", "repostats", d2)
}
