package configs

type DNSProvider struct {
	Name string
	IPs  []string
}

var Providers = []DNSProvider{
	{"Shecan", []string{"178.22.122.100", "185.51.200.2"}},
	{"Electro", []string{"78.157.42.100", "78.157.42.101"}},
	{"Begzar", []string{"185.55.226.26", "185.55.226.25"}},
	{"DNS Pro", []string{"87.107.110.109", "87.107.110.110"}},
	{"Google", []string{"8.8.8.8", "8.8.4.4"}},
	{"Cloudflare", []string{"1.1.1.1", "1.0.0.1"}},
	{"403", []string{"10.202.10.202", "10.202.10.102"}},
	{"Radar", []string{"10.202.10.10", "10.202.10.11"}},
	{"Shelter", []string{"94.103.125.157", "94.103.125.158"}},
	{"Pishgaman", []string{"5.202.100.100", "5.202.100.101"}},
	{"Shatel", []string{"85.15.1.14", "85.15.1.15"}},
}
