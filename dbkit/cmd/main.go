package main

// import (
// 	"log"
// 	"math/rand"
// 	"time"

// 	"github.com/DucTran999/shared-pkg/dbkit"
// )

// func main() {
// 	conn, err := dbkit.NewDBConnector(
// 		dbkit.PostgresDriver,
// 		"localhost",
// 		dbkit.WithUsername("test"),
// 		dbkit.WithPassword("test"),
// 		dbkit.WithDatabase("atlana_shop"),
// 	)
// 	if err != nil {
// 		log.Fatal("e", err)
// 	}

// 	conn2, err := dbkit.NewDBConnector(
// 		dbkit.PostgresDriver,
// 		"localhost",
// 		dbkit.WithUsername("test"),
// 		dbkit.WithPassword("test"),
// 		dbkit.WithDatabase("athena"),
// 	)
// 	if err != nil {
// 		log.Fatal("e", err)
// 	}

// 	db, err := conn.Open()
// 	if err != nil {
// 		log.Panicln(err)
// 	}

// 	db2, err := conn2.Open()
// 	if err != nil {
// 		log.Panicln(err)
// 	}

// 	for i := range 4 {
// 		go func() {
// 			for {
// 				if i%2 == 0 {
// 					log.Println("routine", i, "do query")
// 					err := db.DB().Exec("Select pg_sleep(?);", i+2).Error
// 					if err != nil {
// 						log.Println("routine:", i, "err", err)
// 					}
// 				} else {
// 					log.Println("routine", i, "do query")
// 					err := db2.DB().Exec("Select pg_sleep(?);", i+2).Error
// 					if err != nil {
// 						log.Println("routine:", i, "err", err)
// 					}
// 				}

// 				time.Sleep(time.Second*10 + time.Duration(rand.Intn(1000)))
// 			}
// 		}()
// 	}
// 	time.Sleep(time.Second * 20)
// 	// err = db.Close()
// 	// log.Print(err)
// 	// err = db2.DB().Exec("Select pg_sleep(5);").Error
// 	// log.Fatal("routine:", "err", err)

// 	// time.Sleep(time.Second * 10)
// 	// go func() {
// 	// 	err = db2.DB().Exec("Select pg_sleep(5);").Error
// 	// 	log.Fatal("routine:", "err", err)

// 	// }()
// 	// go func() {
// 	// 	err = db.DB().Exec("Select pg_sleep(5);").Error
// 	// 	log.Fatal("routine:", "err", err)
// 	// }()

// 	time.Sleep(time.Minute * 10)

// 	defer func() {
// 		err := db2.Close()
// 		log.Println(err)
// 		err = db.Close()
// 		log.Print(err)
// 	}()
// }
