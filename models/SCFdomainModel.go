//  "UUID": "857e0d76-e6b2-4893-b901-e8bf6518a8c8",
//  "#\u00a0": 1,
//  "SCF Domain": "Security & Privacy Governance",
//  "SCF Identifier": "GOV",
//  "Security & Privacy by Design (S|P) Principles":
//  Principle Intent

package models

type SCFDomain struct {
	//gorm.Model
	//Id_domain       int      `json:"id_domain"`
	//UUID            string   `json:"UUID"`
	//ID            int    `json:"#\u00a0"`
	SCFDomain     string
	SCFIdentifier string
	Controls      interface{}
	//SecurityPrivacy []string `json:"Security & Privacy by Design (S|P) Principles"`
	//PrincipleIntent string   `json:"Principle Intent"`
	//Date_created    string   `gorm:"-:all"`
	//Date_updated    string   `json:"Date_updated"`
}

// let arr = Array('ISO\n27001\nv2013','PCIDSS\nv3.2','COBIT\n2019');
// let stringed = JSON.stringify(arr);
// console.log(JSON.stringify(arr));

// let finalo = stringed.replace("[", "(");
//  finalo = finalo.replace("]", ")");
//  finalo = finalo.replaceAll('"', '\'');

// console.log(finalo);
