package router

import "github.com/labstack/echo/v4"

func rootRouter(e *echo.Echo) {
	v2 := e.Group("/api/v2")
	{
		v2.POST("/login", dummyHandler)
		v2.POST("/logout", dummyHandler)
		v2.POST("/refresh", dummyHandler)

		user := v2.Group("/users")
		{
			user.GET("/me", dummyHandler)
			user.PUT("/me/password", dummyHandler)

			user.GET("/:id", dummyHandler)
			user.POST("/", userHandler.CreateUser)
			user.POST("/verify/:token", dummyHandler)
		}

		problem := v2.Group("/problems")
		{
			problem.POST("/", dummyHandler)

			problem.GET("/:id", dummyHandler)
			problem.PUT("/:id", dummyHandler)

			problem.POST("/:id/sets", dummyHandler)
			problem.PUT("/:id/sets/:setId", dummyHandler)
			problem.DELETE("/:id/sets/:setId", dummyHandler)

			problem.POST("/:id/sets/:setId/cases", dummyHandler)
			problem.PUT("/:id/sets/:setId/cases/:caseId", dummyHandler)
			problem.DELETE("/:id/sets/:setId/cases/:caseId", dummyHandler)
		}

		contest := v2.Group("/contests")
		{
			contest.POST("/", contestHandler.CreateContest)
			contest.GET("/:id", contestHandler.FindContestByID)
			contest.PUT("/:id", dummyHandler)
			contest.POST("/:id/join", dummyHandler)
			contest.GET("/:id/problems", dummyHandler)
			contest.GET("/:id/ranking", dummyHandler)

			contest.POST("/:id/submissions", dummyHandler)
			contest.GET("/:id/submissions", dummyHandler)
			contest.GET("/:id/submissions/:submissionId", dummyHandler)
		}
	}
}

func dummyHandler(c echo.Context) error {
	return c.String(200, "ok")
}
