package models

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
}

var Users = []User{
    {ID: 1, Name: "Juan Garcia", Email: "juan@sena.edu.co", Age: 28},
    {ID: 2, Name: "Juan Gonzalez", Email: "JuanFe@gmail.com", Age: 18},
    {ID: 3, Name: "Susana Gomez", Email: "susana@sena.edu.co", Age: 45},
}