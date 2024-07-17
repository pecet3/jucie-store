package auth

import (
	"github.com/pecet3/my-api/utils"
)

// to do: change name email activator to email session, replace it to session services

type email struct{}

type emailServices interface {
	SendActivateEmail(to, url, username string) error
}

func (e email) SendActivateEmail(to, url, userName string) error {
	subject := "🔒Auth - Confirm Email🔒 pecet.it (no reply)"
	body := `
    <html>
    	<body>
    		<h1>Hello ` + userName + `,</h1>
    			<p>This is link to <b>activate</b> your profile:</p>
				<h2>
					<i>` + url + `</i>
				</h2>
    	</body>
    </html>
    `
	if err := utils.SendEmail(to, subject, body); err != nil {
		return err
	}
	return nil
}
