package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// crear la estructura json de como vas a guardar los datos
type Country struct {
	ID        string  `json:"ID"`
	Name      string  `json:"name"`
	Continent string  `json:"continent"`
	MinSalary float64 `json:"minsalary"`
	Language  string  `json:"language"`
}

// Creemos nuestra base de datos local!
// Declarar la variable countries como una variable global
var countries = []Country{
	{
		ID:        "1",
		Name:      "Spain",
		Continent: "Europe",
		MinSalary: 1500.00,
		Language:  "Spanish",
	},
	{
		ID:        "2",
		Name:      "Japan",
		Continent: "Asia",
		MinSalary: 2000.50,
		Language:  "Japanese",
	},
	{
		ID:        "3",
		Name:      "Australia",
		Continent: "Oceania",
		MinSalary: 2500.75,
		Language:  "English",
	},
	{
		ID:        "4",
		Name:      "Brazil",
		Continent: "South America",
		MinSalary: 1000.25,
		Language:  "Portuguese",
	},
	{
		ID:        "5",
		Name:      "Egypt",
		Continent: "Anfrica",
		MinSalary: 800.50,
		Language:  "Arabic",
	},
}

func main() {
	//Marshal convierte structuras de GO a Json
	countriesJson, _ := json.Marshal(countries)
	fmt.Println(string(countriesJson))
	//Instaciamos una variable gin.Default()
	router := gin.Default()
	//HandlerFunc de tipo GET
	router.GET("/countries", getCountries)
	//HandlerFunc de tipo POST
	router.POST("/countries", AddCountries)
	//HandlerFunc que permite mostrarte un pais y su informaciòn en especifico!
	router.GET("/countries/:id", ShowCountry)
	//HandlerFunc que permite elimnar algun pais 3:)
	router.DELETE("/countries/:id", DeleteCountry)
	//Establecemos el IP adress y el puerto!
	router.Run("localhost:8080")
}

// Esta funcion de acà es considerada una HandlerFunc
func getCountries(c *gin.Context) {
	c.JSON(http.StatusOK, countries)
}

// Creamos una nueva funciòn para agregar Countries
func AddCountries(c *gin.Context) {
	var NewCountry Country
	if err := c.BindJSON(&NewCountry); err != nil {
		return
	}
	countries = append(countries, NewCountry)
	c.IndentedJSON(http.StatusCreated, NewCountry)
}

// Crearemos una funciòn para que nos muestre un item Especìfico
func ShowCountry(c *gin.Context) {
	//Con esta Funcion Param
	id := c.Param("id")
	for _, a := range countries {
		if a.ID == id {
			c.IndentedJSON(http.StatusFound, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Pais inexistente"})
}

// HandlerFunc para eliminar Paises!
func DeleteCountry(c *gin.Context) {
	id := c.Param("id")
	for idx, a := range countries {
		if a.ID == id {
			countries = append(countries[:idx], countries[idx+1:]...)
			c.IndentedJSON(http.StatusAccepted, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No puedes eliminar un pais que no existe"})
}
