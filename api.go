package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	router := gin.Default()

	router.StaticFile("/favicon.ico", "./favicon.ico")
	router.NoRoute(func(c *gin.Context) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":                 http.StatusBadRequest,
				"message":              "Bad Request",
				"suggestions":          "Check the official API documentation to see how to properly format endpoints/routes",
				"v1_documentation_url": "https://www.pharmacodb.com/docs/api",
			},
		})
	})

	v1 := router.Group("v1")
	{
		v1.GET("/cell_lines", GetCells)
		v1.GET("/cell_lines/stats", GetCellStats)
		v1.GET("/cell_lines/ids", GetCellIDs)
		v1.GET("/cell_lines/ids/:id", GetCellByID)
		v1.GET("/cell_lines/names", GetCellNames)

		v1.GET("/tissues", GetTissues)
		v1.GET("/tissues/stats", GetTissueStats)
		v1.GET("/tissues/ids", GetTissueIDs)
		v1.GET("/tissues/ids/:id", GetTissueByID)
		v1.GET("/tissues/names", GetTissueNames)

		v1.GET("/drugs", GetDrugs)
		v1.GET("/drugs/stats", GetDrugStats)
		v1.GET("/drugs/ids", GetDrugIDs)
		v1.GET("/drugs/ids/:id", GetDrugByID)
		v1.GET("/drugs/names", GetDrugNames)

		v1.GET("/datasets", GetDatasets)
		v1.GET("/datasets/stats", GetDatasetStats)
		v1.GET("/datasets/ids", GetDatasetIDs)
		v1.GET("/datasets/ids/:id", GetDatasetByID)
		v1.GET("/datasets/names", GetDatasetNames)
	}

	router.Run(":3000")
}
