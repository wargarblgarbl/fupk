package main
import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"strconv"
)
type Thing struct {
	Name           string  `json:"name"`
}

type eve []struct {
	KillID   int    `json:"killID"`
	KillTime string `json:"killTime"`
	MoonID   int    `json:"moonID"`
	Attackers []struct {
		AllianceID      int    `json:"allianceID"`
		AllianceName    string `json:"allianceName"`
		CharacterID     int    `json:"characterID"`
		CharacterName   string `json:"characterName"`
		CorporationID   int    `json:"corporationID"`
		CorporationName string `json:"corporationName"`
		DamageDone      int    `json:"damageDone"`
		FactionID       int    `json:"factionID"`
		FactionName     string `json:"factionName"`
		FinalBlow       int    `json:"finalBlow"`
		SecurityStatus  float64    `json:"securityStatus"`
		ShipTypeID      int    `json:"shipTypeID"`
		WeaponTypeID    int    `json:"weaponTypeID"`
	} `json:"attackers"`

	SolarSystemID int `json:"solarSystemID"`
	Victim        struct {
		AllianceID      int    `json:"allianceID"`
		AllianceName    string `json:"allianceName"`
		CharacterID     int    `json:"characterID"`
		CharacterName   string `json:"characterName"`
		CorporationID   int    `json:"corporationID"`
		CorporationName string `json:"corporationName"`
		DamageTaken     int    `json:"damageTaken"`
		FactionID       int    `json:"factionID"`
		FactionName     string `json:"factionName"`
		ShipTypeID      int    `json:"shipTypeID"`
	} `json:"victim"`
	Zkb struct {
		Hash       string  `json:"hash"`
		LocationID int     `json:"locationID"`
		Points     int     `json:"points"`
		TotalValue float64 `json:"totalValue"`
	} `json:"zkb"`
}


func curlJson(url string) (stuff []byte){
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	} else {
		defer r.Body.Close()
		contents, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		stuff :=  contents
		return stuff
	}
}
func main()  {
	activeapi := "https://zkillboard.com/api/no-items/allianceID/1081578607/limit/100"
	jsonstring := curlJson(activeapi)
	var f eve
	err := json.Unmarshal(jsonstring, &f)
	if err != nil {
		fmt.Println(err)
	}
	for _,element := range f {
		var killer string
		var inship string
		place := checkCrestLocation(strconv.Itoa(element.SolarSystemID))
		if element.Victim.AllianceID != 1081578607 {
			for _, thing := range element.Attackers {
				if thing.FinalBlow != 0 {
					killer = thing.CharacterName
					inship = strconv.Itoa(thing.ShipTypeID)
				}
			}
			ourship := checkCrest(strconv.Itoa(element.Victim.ShipTypeID))
			cheese := checkCrest(inship)
			var welp string
			if killer == "deadlypie" {
				welp = "this must be a mistake"
			} else {
				welp = "YAY"
			}
			fmt.Println("<tr bgcolor=\"#6CBB3C\"><td><a href=\"https://zkillboard.com/kill/"+strconv.Itoa(element.KillID)+"/\">",element.KillTime,"</a></td><td>",killer,"</td><td>", cheese, "</td><td>", element.Victim.CharacterName,"</td><td>", ourship, "</td><td>", place ,"</td><td>", welp,"</td></tr>")
		} else {
			for _, thing := range element.Attackers {
				if thing.FinalBlow !=0 {
					killer = thing.CharacterName
					inship = strconv.Itoa(thing.ShipTypeID)
				}
			}
			ourship := checkCrest(strconv.Itoa(element.Victim.ShipTypeID))
			cheese := checkCrest(inship)
			var welp string
			/*
			if element.Victim.CharacterName == "deadlypie" {
				welp ="MORE LIKE DEADPIE"
			} else {
*/
			welp = "WELP"
//			}
			
                        fmt.Println("<tr bgcolor=\"#FF2400\"><td><a href=\"https://zkillboard.com/kill/"+strconv.Itoa(element.KillID)+"/\">",element.KillTime,"</a></td><td>",killer,"</td><td>", cheese, "</td><td>", element.Victim.CharacterName,"</td><td>", ourship, "</td><td>", place ,"</td><td>", welp,"</td></tr>")
		}
	}
}
func checkCrest (what string) (string) {
	thething := "https://crest-tq.eveonline.com/types/"+what+"/"
	jston := curlJson(thething)
	var x Thing
	brr := json.Unmarshal(jston, &x)
	if brr != nil {
		fmt.Println(brr)
	}
	return x.Name
}

func checkCrestLocation (what string) (string) {
	thething := "https://crest-tq.eveonline.com/solarsystems/"+what+"/"
	jston := curlJson(thething)
	var x Thing
	brr := json.Unmarshal(jston, &x)
	if brr != nil {
		fmt.Println(brr)
	}
	return x.Name
}
