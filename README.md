# SpamMail
## Go utility to send bulk emails


## Setup:

1. Ensure SMTP server is running on localhost port 25.  Go Script tested with smtp4dev.
	* Further documentation found at [SMTPDEV](https://github.com/rnwood/smtp4dev)
2. Email Entry should have the following format. 
	* {"From": "what@who.com", "To": "who@who.com","Subject": "Some Subject","Text":"Some Text"}

## Execute:

go run email.go some_mail_list.txt

## Notes:

* Processing Time returned to console after completion.