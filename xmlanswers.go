package xmlanswers

import "encoding/xml"

// SecretResponce struct for the  request containing the secerts, I'll collapse this in to one later
type SecretResponce struct {
	XMLName xml.Name
	AA      struct {
		XMLName xml.Name
		BB      struct {
			CC struct {
				Token   string `xml:"Token"`
				XMLName xml.Name
				DD      struct {
					Name  []string `xml:"Name"`
					Items struct {
						SID struct {
							Svalue []string `xml:"Value"`
						} `xml:",any"`
					} `xml:",any"`
				} `xml:",any"`
			} `xml:",any"`
		} `xml:",any"`
	} `xml:",any"`
}

// UnwindSecret to pull out relavent data from the return token
func UnwindSecret(data []byte) []string {
	m := &SecretResponce{}
	xml.Unmarshal(data, m)
	a := m.AA.BB.CC.DD.Items.SID.Svalue
	return a

}

// UnwindToken to pull out relavent data from the return token
func UnwindToken(data []byte) string {
	m := &SecretResponce{}
	xml.Unmarshal(data, m)
	a := m.AA.BB.CC
	// If you want to see the xml result with key:values uncomment this
	//fmt.Printf("%+v\n", a)
	return a.Token

}

// WindCreds is to build the SOAP user,pass, domain string
func WindCreds(user string, pass string, domain string) string {
	authPayLoad := `<?xml version="1.0" encoding="utf-8"?>
<soap12:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
  <soap12:Body>
    <Authenticate xmlns="urn:thesecretserver.com">
      <username>` + user + `</username>
      <password>` + pass + `</password>
      <organization></organization>
      <domain>` + domain + `</domain>
    </Authenticate>
 </soap12:Body>
</soap12:Envelope>`

	return authPayLoad
}

// WindToken is to build the SOAP token and secretid
func WindToken(token string, secretid string) string {
	tokenPayLoad := `<?xml version="1.0" encoding="utf-8"?>
<soap12:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
  <soap12:Body>
    <GetSecret xmlns="urn:thesecretserver.com">
      <token>` + token + `</token>
      <secretId>` + secretid + `</secretId>
      <loadSettingsAndPermissions>0</loadSettingsAndPermissions>
      <codeResponses>
        <CodeResponse>
          <ErrorCode></ErrorCode>
          <Comment></Comment>
          <AdditionalComment></AdditionalComment>
        </CodeResponse>
        <CodeResponse>
          <ErrorCode></ErrorCode>
          <Comment></Comment>
          <AdditionalComment></AdditionalComment>
        </CodeResponse>
      </codeResponses>
    </GetSecret>
  </soap12:Body>
</soap12:Envelope>`

	return tokenPayLoad
}
