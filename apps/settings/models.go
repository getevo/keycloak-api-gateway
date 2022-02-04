package settings
/*
* Copyright © 2022 Allan Nava <EVO TEAM>
* Created 04/02/2022
* Updated 04/02/2022
*
 */
var Settings settings

type Params struct {
	ID    int    `gorm:"column:id_parameter"`
	Class string `gorm:"column:class_parameter"`
	Name  string `gorm:"column:name_parameter"`
	Value string `gorm:"column:value_parameter"`
}

func (Params) TableName() string {
	return "parameter"
}

type settings struct {
	Keycloak struct {
		Server   string `key:"SERVER" yaml:"server"`
		Realm    string `key:"REALM" yaml:"realm"`
		Client   string `key:"CLIENT" yaml:"client"`
		Username string `key:"USERNAME" yaml:"username"`
		Password string `key:"PASSWORD" yaml:"password"`
	} `key:"KEYCLOAK" yaml:"keycloak"`

	API struct {
		Server        string `key:"SERVER" yaml:"server"`
		Key           string `key:"KEY" yaml:"key"`
		Secure        bool   `key:"SECURE" yaml:"secure"`
		TrustedSubnet string `key:"TRUSTED_SUBNET" yaml:"trusted_subnet"`
	} `key:"API" yaml:"api"`
}
