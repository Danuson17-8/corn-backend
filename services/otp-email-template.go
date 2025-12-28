package services

import (
	"fmt"
)

func GenerateOTPEmailTemplate(code int) string {
	return fmt.Sprintf(`
<html>
  <body style="font-family: Arial, sans-serif; display: flex; justify-content: center;">
    <div style="height: 250px; width: 100%%; max-width: 500px; background-color: #ffe817; padding: 30px; text-align: center;">
      <h1 style="color: #030303;">Corn Cornn!!</h1>
      <p>Your OTP code is:</p>
      <div style="background-color: #000000; padding: 10px; margin-right: 100px; margin-left: 100px;">
        <h2 style="color: #ffffff; font-size: 2em; margin: 20px 0;">%d</h2>
      </div>
      <hr style="width: 310px;">
      <p style="font-size: 0.9em; color: #777;">This code is valid for 5 minutes.</p>
    </div>
  </body>
</html>
`, code)
}
