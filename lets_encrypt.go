package nginxconf

type LetsEncrypt struct{}

func (e LetsEncrypt) SSLCertificatePath(domain string) string {
	return "/etc/letsencrypt/live/" + domain + "/fullchain.pem"
}

func (e LetsEncrypt) SSLCertificateKeyPath(domain string) string {
	return "/etc/letsencrypt/live/" + domain + "/privkey.pem"
}
