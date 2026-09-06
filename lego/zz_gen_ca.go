package lego

const (
	CodeActalis = "actalis"

	CodeDigicert = "digicert"

	CodeFreeSSL = "freessl"

	CodeGlobalSign = "globalsign"

	CodeGoogleTrust = "googletrust"

	CodeGoogleTrustStaging = "googletrust-staging"

	CodeLetsEncrypt = "letsencrypt"

	CodeLetsEncryptStaging = "letsencrypt-staging"

	CodeLiteSSL = "litessl"

	CodePeeringHub = "peeringhub"

	CodeSSLComECC = "sslcomecc"

	CodeSSLComRSA = "sslcomrsa"

	CodeSectigo = "sectigo"

	CodeSectigoEV = "sectigoev"

	CodeSectigoOV = "sectigoov"

	CodeZeroSSL = "zerossl"
)

const (
	DirectoryURLActalis = "https://acme-api.actalis.com/acme/directory"

	DirectoryURLDigicert = "https://one.digicert.com/mpki/api/v1/acme/v2/directory"

	DirectoryURLFreeSSL = "https://acmepro.freessl.cn/v2/DV"

	DirectoryURLGlobalSign = "https://emea.acme.atlas.globalsign.com/directory"

	DirectoryURLGoogleTrust = "https://dv.acme-v02.api.pki.goog/directory"

	DirectoryURLGoogleTrustStaging = "https://dv.acme-v02.test-api.pki.goog/directory"

	DirectoryURLLetsEncrypt = "https://acme-v02.api.letsencrypt.org/directory"

	DirectoryURLLetsEncryptStaging = "https://acme-staging-v02.api.letsencrypt.org/directory"

	DirectoryURLLiteSSL = "https://acme.litessl.com/acme/v2/directory"

	DirectoryURLPeeringHub = "https://stica.peeringhub.io/acme"

	DirectoryURLSSLComECC = "https://acme.ssl.com/sslcom-dv-ecc"

	DirectoryURLSSLComRSA = "https://acme.ssl.com/sslcom-dv-rsa"

	DirectoryURLSectigo = "https://acme.sectigo.com/v2/DV"

	DirectoryURLSectigoEV = "https://acme.sectigo.com/v2/EV"

	DirectoryURLSectigoOV = "https://acme.sectigo.com/v2/OV"

	DirectoryURLZeroSSL = "https://acme.zerossl.com/v2/DV90"
)

func GetDirectoryURL(code string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func GetAllCodes() []string { _ = "STUB: not implemented"; return nil }
