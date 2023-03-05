package mail

func GenerateMessageBody() string {
	return `
		Hallo Anita, 

		Im Anhang findest du meinen Kassen und Stundenraport. 

		Viele Grüße
	`

}

func GenerateMessageSubject() string {
	return "Kassen und Stundenraport"
}
