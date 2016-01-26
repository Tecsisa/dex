package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type oAuth2Token struct {
	AccessToken  string `json:"access_token"`
	IDToken      string `json:"id_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func main1() {

	// Auth using client credentials
	urlclients := "http://localhost:5556/token"
	fmt.Println("URL:>", urlclients)

	// Create request
	req, err := http.NewRequest("POST", urlclients, strings.NewReader("grant_type=client_credentials"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Encode cliendID:password using base64
	var userCreds = []byte(`9jCU4aaDHjV-y59SSlGwfrmpdo4mIkGBW4E41QvI-X0=@127.0.0.1:uYDECfyJv3rSpf7oxN8KxmUB1o2Ea2N0y1pOIeoXHTZfBV7jNo2f85S1c4OOKEqnDilHseDGA_vNiSxxighuN89OVJhcj7eX`)

	// Add header according OAuth spec
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString(userCreds))
	client := &http.Client{}

	// Make call
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print response
	fmt.Println("Response Status Auth:", resp.Status)
	fmt.Println("Response Headers Auth:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	var m oAuth2Token
	err = json.Unmarshal(body, &m)
	fmt.Println("Response Body Auth access token:", m.AccessToken)
	fmt.Println("Response Body Auth IDToken:", m.IDToken)
	fmt.Println("Response Body Auth RefreshToken:", m.RefreshToken)

	// Call clients API
	urlclients = "http://localhost:5556/api/v1/clients"
	req, err = http.NewRequest("GET", urlclients, nil)

	// Use access token retrieved
	req.Header.Add("Authorization", "Bearer "+m.AccessToken)
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print response
	fmt.Println("Response Status Api clients:", resp.Status)
	fmt.Println("Response Headers Api clients:", resp.Header)
	body, _ = ioutil.ReadAll(resp.Body)
	fmt.Println("Response Body Api clients:", string(body))

}

func main2() {

	// Auth using client credentials
	urlclients := "http://localhost:5556/token"
	fmt.Println("URL:>", urlclients)

	// Create request
	code := "ofulQb7SbLQ%3D&state="
	req, err := http.NewRequest("POST", urlclients, strings.NewReader("grant_type=authorization_code&code="+code))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Encode cliendID:password using base64
	//var userCreds = []byte(`9jCU4aaDHjV-y59SSlGwfrmpdo4mIkGBW4E41QvI-X0=@127.0.0.1:uYDECfyJv3rSpf7oxN8KxmUB1o2Ea2N0y1pOIeoXHTZfBV7jNo2f85S1c4OOKEqnDilHseDGA_vNiSxxighuN89OVJhcj7eX`)

	// Add header according OAuth spec
	req.Header.Add("Authorization", "Basic "+"eyJhbGciOiJSUzI1NiIsImtpZCI6IjFBVVlUY1ZfWUMtUmdCMDNqVEtCcV9mQ2NQWEJDdHloMGF4MndJcVc0R05qOWFacWlEVDMydUl0Rk03Rk1pRVg5ZEhxX1RiTnloWjc1bU9jTk1GV0NqMUxYUFM5R3RMd01rc0NoRW4ydE80SjRyQlRiaDlGSUI5blRTSHpzd0tqdWdsbFV6MWZpUmdNakF0dExsLUJZSmlEX1dNc3BqNGhrdXA2UDBHd3VCUlRPcG1NTjZLUGI5eTNKWFBWaWtNbjNEcUhVMGN2N2JRbkpXR0JDQk51TmloUnZoU1hZNHQtV1Bya2dDUFFieEpHRExKU1RGdXZjYmJVYVdjZUY2Yy1pcnBpMDdlc1RfM3J2bzkwZ0ZOcldqTTRHS0t5VVlacWhvaXdTOVBNQWJEM0R6ZlRTYi16TGhVdWl1cDZqNW9WSDZPSVlqSUZFS2pjT3dMRVNqM0s5dz09IiwidHlwIjoiSldUIn0.eyJhdWQiOiI5akNVNGFhREhqVi15NTlTU2xHd2ZybXBkbzRtSWtHQlc0RTQxUXZJLVgwPUAxMjcuMC4wLjEiLCJlbWFpbCI6ImdlcnNvbi5wb3pvLmdvbnphbGV6QHRlY3Npc2EuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImV4cCI6MS40NTM0MTE5MjRlKzA5LCJpYXQiOjEuNDUzMzY4NzI0ZSswOSwiaXNzIjoiaHR0cDovLzEyNy4wLjAuMTo1NTU2IiwibmFtZSI6IiIsInN1YiI6IjU2MDNhMDhkLTlkMDktNDI2Mi1hOThhLTIzMWNjNzAwM2U2ZCJ9.OjRsJruGLzf9tMXL2mvzlVm4q_m5ZLm4_htGqtYJZ9u-Nu0hjMc8XfSTZXlx0GASAyepkgtaIz-MdjgbhMAjpedh79VRFg2bjrEnJ86LqVSNYZpgHTl68LxM5hEV8nepH4EMc3lbWh9UzSchCpGs6vtIOQAVkDwhEzDNvXjhPCfoJ6ZyQJuD3-suyYVWnUQXX5myyu8Y_YTcny9RhaGTDiT4PD8MA5_tPRKmeXqVTwMfYbdEd8S67wV-6eXOQ5-Owyr7Aig1Xad-4pxJ2xk1668KfKAobvbqDR0mRnB4V2SJ-gX9nmZVbLZASkKn3tZ5itREiMAw_46s0WBB1Uj7hQ")

	client := &http.Client{}

	// Make call
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print response
	fmt.Println("Response Status Auth:", resp.Status)
	fmt.Println("Response Headers Auth:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	var m oAuth2Token
	err = json.Unmarshal(body, &m)
	fmt.Println("Response Body Auth access token:", m.AccessToken)
	fmt.Println("Response Body Auth IDToken:", m.IDToken)
	fmt.Println("Response Body Auth RefreshToken:", m.RefreshToken)

}
