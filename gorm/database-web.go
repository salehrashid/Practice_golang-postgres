package main

import (
	"fmt"
	"go-postgres/constants"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"html/template"
	"net/http"
)

type Cars struct {
	gorm.Model
	Id     int
	Name   string
	Engine string
}

func main() {
	httpHandler()
	dbInsertRecords()
}

func httpHandler() {
	fmt.Println("server nya lagi jalan nih bang, http://localhost:7000")

	http.HandleFunc("/", root)
	http.HandleFunc("/index", index)

	/**
	digunakan untuk menghidupkan server sekaligus menjalankan aplikasi
	menggunakan server tersebut. Di Go, 1 web aplikasi adalah 1 buah
	server berbeda.
	*/
	if err := http.ListenAndServe(":7000", nil); err != nil {
		panic(err)
	}
}

func root(writer http.ResponseWriter, _ *http.Request) {
	tmplt := template.Must(template.ParseFiles("gorm/template/root.html"))
	if err := tmplt.Execute(writer, nil); err != nil {
		panic(err)
	}
}

func index(writer http.ResponseWriter, _ *http.Request) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		constants.Host, constants.Port, constants.User, constants.Password, constants.Dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var cars []Cars
	db.Find(&cars)

	tmplt := template.Must(template.ParseFiles("gorm/template/index.html"))
	if err := tmplt.Execute(writer, cars); err != nil {
		panic(err)
	}
}

func dbInsertRecords() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		constants.Host, constants.Port, constants.User, constants.Password, constants.Dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&Cars{}); err != nil {
		panic(err)
	}
	//db.Create(&Cars{Name: "skyline r34", Engine: "RB26DETT twin-turbocharged 2.6-liter inline-six"})
	//db.Create(&Cars{Name: "dodge charger srt", Engine: "Supercharged 6.2L HEMI V8"})
	//db.Create(&Cars{Name: "supra mk4", Engine: "2JZ-GTE 3.0-litre twin-turbocharged straight 6"})
	//db.Create(&Cars{Name: "skyline r35", Engine: "VR38DETT V6 3.8 L twin-turbocharged"})
	//db.Create(&Cars{Name: "subaru wrx sti", Engine: "flat-four turbocharged 2.5-liter"})
	//db.Create(&Cars{Name: "shelby gt500", Engine: "5.2 L Predator V8"})
	//db.Create(&Cars{Name: "corvette z06 2006", Engine: "7.0 L LS7 V8"})
	//db.Create(&Cars{Name: "slivia s15", Engine: "2.0 L SR20DET I4 turbo"})
	//db.Create(&Cars{Name: "mazda rx7", Engine: "13B-REW twin-turbo twin-rotor"})
	//db.Create(&Cars{Name: "lexus lfa", Engine: "4.8 L 1LR-GUE V10"})
	//db.Create(&Cars{Name: "mitsubishi evo 9", Engine: "2.0 L 4G63 Inline 4 Turbocharged"})
	//db.Create(&Cars{Name: "370z fairlady", Engine: "3.7 L VQ37VHR V6"})
}
