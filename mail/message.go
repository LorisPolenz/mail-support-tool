package mail

func GenerateMessageBody() string {
	return `
		Hallo Anita,
		<br>
		Im Anhang findest du meinen Kassen und Stundenraport.
		<br>
		Viele Grüße
	`

}

func GenerateMessageSubject() string {
	return "Kassen und Stundenraport"
}
