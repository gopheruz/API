package main

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/nurmuhammaddeveloper/API/api"
	"github.com/nurmuhammaddeveloper/API/config"
	"github.com/nurmuhammaddeveloper/API/models"
	"github.com/nurmuhammaddeveloper/API/storage"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load(".")
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("filed to connect to database: %v", err)
	}
	defer db.Close()
	fmt.Println("database connected successfully")
	s := storage.NewDBManager(db)
	// Create
	// cretedDataINfo, err := database.Create(models.CreateStudentRequest{
	// 	FirstName:   "fake",
	// 	LastName:    "fake",
	// 	Email:       "fakeemail@gmail.com",
	// 	PhoneNumber: "+1234567890",
	// 	UserName:    "@fakeuser",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(cretedDataINfo)

	// Update
	updatedData, err := s.Update(&models.UpdateStudentRequest{
		FirstName:   "Sug'diyona",
		LastName:    "Ahmadjonova",
		UserName:    "@sugdiyona",
		Email:       "Ahmadjonovasugdiyona@gmail.com",
		PhoneNumber: "+998901970882",
		ID:          2,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(updatedData)

	// Get

	// result, err := s.Get(2)
	// if err != nil {
	// 	log.Fatalf("Filed to get data from database: %v", err)
	// }
	// fmt.Println(result)

	//Delete
	// err = database.Delete(4)
	// if err!= nil {
	//     log.Fatal(err)
	// }

	// data, err := database.GetAll(models.GetStudentsQueryParam{
	// 	Limit: 3,
	// 	Page:  1,
	// })
	// if err!= nil {
	//     log.Fatal(err)
	// }
	// for _, val := range data.Students{
	// 	fmt.Println( val)
	// }
	// fmt.Println(data.Count)

	storage := storage.NewDBManager(db)
	server := api.NewServer(storage)
	err = server.Run(cfg.HttpPort)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("server is running: ", server)
}
