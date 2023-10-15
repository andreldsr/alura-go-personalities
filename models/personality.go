package models

import "alura-go-personalities/database"

type Personality struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Story string `json:"story"`
}

func FindAllPersonalities() []Personality {
	var personalities []Personality
	database.DB.Find(&personalities)
	return personalities
}

func FindPersonalityById(id int) (p Personality) {
	database.DB.First(&p, id)
	return
}

func CreatePersonality(personality *Personality) {
	database.DB.Create(&personality)
}

func DeletePersonality(id int) (p Personality) {
	database.DB.Delete(&p, id)
	return
}

func UpdatePersonality(id int, personality Personality) (p Personality) {
	database.DB.First(&p, id)
	database.DB.Model(&p).Updates(personality)
	return
}
