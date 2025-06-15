package main

import (
	"log"
	"vijju/graph"
	"vijju/post/model"
	postRepository "vijju/post/repository"
	"vijju/post/service"
	userModel "vijju/user/model"
	userRepository "vijju/user/repository"
	userSer "vijju/user/service"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	dsn := "host=localhost port=5432 user=postgres password=1575250377 dbname=mydb sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate models
	db.AutoMigrate(&userModel.User{}, &model.Post{})

	// Initialize repositories
	userRepo := userRepository.NewUserRepository(db)
	postRepo := postRepository.NewPostRepository(db)

	// Initialize services
	userService := userSer.NewUserService(userRepo)
	postService := service.NewPostService(postRepo)

	// Set up Gin router
	r := gin.Default()

	// GraphQL handler
	graphQLHandler := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{
			Resolvers: &graph.Resolver{
				UserService: userService,
				PostService: postService,
			},
		}),
	)

	// Playground route
	r.GET("/", gin.WrapH(playground.Handler("GraphQL Playground", "/query")))
	r.POST("/query", gin.WrapH(graphQLHandler))

	log.Println("Server running at http://localhost:8080/")
	r.Run(":8080")
}
