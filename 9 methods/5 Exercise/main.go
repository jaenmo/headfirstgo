package main

import ("fmt"
		"errors"
		"log")

type Date struct{
	Year int
	Month int
	Day int
}
func (d *Date) SetYear(year int) error{
	if year < 1900{
		return errors.New("invalid year")
	}
	d.Year = year
	return nil
}
func (d *Date) SetMonth(month int) error{
	if month < 1 || month > 12{
		return errors.New("invalid month")
	}
	d.Month = month
	return nil
}
func (d *Date) SetDay(day int) error{
		if day < 1 || day > 31{
		return errors.New("invalid day")
	}
	d.Day = day
	return nil
}

func main(){
	date := Date{}
	err := date.SetYear(1920)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(31)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(2)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
}