package service



func ShortenURL(url string) string {

	//I'm going to shortnen the URL using the base 62 encoding
	// and then store it in a map (let's keep like this for now)
	// the flow is simple take the url  generate the new one with the encoder and then
	// check the length of the current shortned URLs exist already pass it  to the encoder func

	return ""
}

func Encoder(num int) string {
	// Base62 encoding
	const base62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var encoded string
	num += 1

	for num > 0{
	    remainder := num % 62
	    encoded = string(base62[remainder]) + encoded
	    num /= 62
	}

	return encoded
}
