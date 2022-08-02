package api

import (
	"net/http"

	"github.com/Hammond95/bartender/bartender/dbo"
	"github.com/gin-gonic/gin"
)

/*func RecipeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "recipe.html", nil)
}*/

func (env *HandlersEnv) GetCocktailsHandler(c *gin.Context) {
	c.DefaultQuery("lastId", "")
	c.DefaultQuery("perPage", "30")

	cocktails, err := dbo.GetCocktails(env.MongoDB, c)
	if err != nil {
		c.AbortWithError(
			http.StatusInternalServerError,
			err,
		)
	}

	c.JSON(
		http.StatusOK,
		gin.H{"cocktails": cocktails},
	)
}

func (env *HandlersEnv) GetCocktailDetailsHandler(c *gin.Context) {
	cocktailName := c.Params.ByName("name")

	cocktail, err := dbo.GetCocktail(env.MongoDB, cocktailName)
	if err != nil {
		panic(err)
	}

	c.JSON(
		http.StatusOK,
		cocktail,
	)
}

func (env *HandlersEnv) GetCocktailIngredientsHandler(c *gin.Context) {
	cocktailName := c.Params.ByName("name")
	ingredients, err := dbo.GetCocktailIngredients(env.MongoDB, cocktailName)
	if err != nil {
		panic(err)
	}
	c.JSON(
		http.StatusOK,
		ingredients,
	)
}

func V1DefineCocktailRoutes(group *gin.RouterGroup, env HandlersEnv) {
	ckt := group.Group("/cocktail")

	ckt.GET("/all", env.GetCocktailsHandler)
	ckt.GET("/:name", env.GetCocktailDetailsHandler)
	ckt.GET("/:name/ingredients", env.GetCocktailIngredientsHandler)

	//ckt.GET("/recipe", RecipeHandler)
}
