package web

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/seetohjinwei/repostats/data"
)

func Start(pool *pgxpool.Pool) {
	// TODO: quick testing code

	// postgres.AddUser(pool, "seetohjinwei")
	// data.AddRepository(pool, "seetohjinwei", "adventofcode", "main")
	// data.AddRepository(pool, "seetohjinwei", "repostats", "main")

	// d1 := []models.TypeData{
	// 	{"java", 4, 420},
	// 	{"python", 100, 69},
	// }
	// d2 := []models.TypeData{
	// 	{"go", 5, 1234556},
	// }
	// data.UpsertTypeData(pool, "seetohjinwei", "adventofcode", "main", d1)
	// data.UpsertTypeData(pool, "seetohjinwei", "repostats", "main", d2)

	// x1, e1 := data.QueryRepositoryLastUpdated(pool, "seetohjinwei", "adventofcode")
	// fmt.Println(x1, e1)

	// x2, e2 := data.QueryRepository(pool, "seetohjinwei", "adventofcode")
	// if e2 != nil {
	// 	fmt.Println(e2)
	// } else {
	// 	for _, td := range x2 {
	// 		fmt.Println(td.Type, td.FileCount, td.Bytes)
	// 	}
	// }

	// x, _ := data.QueryRepositoryLastUpdated(pool, "x", "y")
	// fmt.Println(x)
	// x, _ = data.QueryRepositoryLastUpdated(pool, "seetohjinwei", "adventofcode")
	// fmt.Println(*x)
	// x, err := data.QueryRepositoryLastUpdated(pool, "seetohjinwei", "ad")
	// fmt.Println(x, err)

	// // data.GetReposWithData("seetohjinwei")

	x2, e2 := data.QueryRepository(pool, "seetohjinwei", "adventofcode")
	fmt.Println(x2, e2)
}
