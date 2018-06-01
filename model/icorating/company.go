package model

type ICORatingCompany struct {
	Title            string
	Type             string
	Date	    	 string
	Industry         string
	Description      string
	Features         string
	Founded			 string
	Technical		 string
	Website          string
	Whitepaper       string
	InvestmentRating string
	HypeScore        float32
	RiskScore        float32
	SocialMedia      map[string]string
}
