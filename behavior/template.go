package behavior

import "fmt"

type IOtp interface {
	getRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
}

type Otp struct {
	iOtp IOtp
}

func (o *Otp) genAndSendOTP(otpLength int) error {
	otp := o.iOtp.getRandomOTP(otpLength)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	err := o.iOtp.sendNotification(message)
	if err != nil {
		return err
	}
	return nil
}

type Sms struct {
	Otp
}

func (s *Sms) getRandomOTP(len int) string {
	randomOtp := "1234"
	fmt.Printf("SMS: generating random otp %s\n", randomOtp)
	return randomOtp
}

func (s *Sms) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache \n", otp)
}

func (s *Sms) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *Sms) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}

type Email struct {
	Otp
}

func (e *Email) getRandomOTP(len int) string {
	randomOtp := "1234"
	fmt.Printf("Email: generating random otp %s\n", randomOtp)
	return randomOtp
}

func (e *Email) saveOTPCache(otp string) {
	fmt.Printf("Email: saving otp: %s to cache \n", otp)
}

func (e *Email) getMessage(otp string) string {
	return "Email OTP for login is " + otp
}

func (e *Email) sendNotification(message string) error {
	fmt.Printf("Email: sending sms: %s\n", message)
	return nil
}

func DoTemplate() {
	smsOTP := &Sms{}
	o := Otp{
		iOtp: smsOTP,
	}
	o.genAndSendOTP(4)

	fmt.Println("")

	emailOtp := &Email{}
	o = Otp{
		iOtp: emailOtp,
	}
	o.genAndSendOTP(4)
}
