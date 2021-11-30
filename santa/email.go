package santa

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

type Email struct {
	from Santa
	to Santa
}

func NewEmail(from Santa, to Santa) Email {
	e := Email{from, to}
	return e
}

func (e Email) To() string {
	return e.to.Email()
}

func (e Email)MessageBody()  string{
	message := "Salut, "+ e.from.FirstName() +", Elful șef aici. \n"
	message += "Misiunea ta este să-i trimiți un cadou frumos crețofolinei " + e.to.FullName() +"! \n"
	message += "Adresa ei este: " + e.to.Address() + "\nTelefon: " + e.to.PhoneNumber()
	message += "\nEa te roagă să ai grijă la următoarele lucruri: \n" + e.to.Mentions()
	message += "\nDacă nu ai idei, mi-a șoptit că-i plac astea: \n" + e.to.Wishes()
	message += "\nȘi că-i vin lucrurile măsura asta: \n" + e.to.Sizing()
	message += "\nHai la mulți ani, dar din dar se face rai, dacă ai nelămuriri poți să mă intrebi la email"
	return message
}

func (e Email)Send() {
	m := gomail.NewMessage()

	m.SetHeader("From", "from email")
	m.SetHeader("To", e.from.Email())
	m.SetHeader("Subject", "Secret Curly Santa 2021")
	m.SetBody("text/plain", e.MessageBody())
	d := gomail.NewDialer("smtp.gmail.com", 587, "username", "password")
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
}