package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
	"time"
)

type Card struct {
	ID   string `json:"id"`
	Name string `json:"cardName"`
}

type set struct {
	abbreviation string
	cardNum      int
}

var sets = []set{
	{
		abbreviation: "SOR",
		cardNum:      252,
	},
	{
		abbreviation: "SHD",
		cardNum:      262,
	},
	{
		abbreviation: "TWI",
		cardNum:      257,
	},
	{
		abbreviation: "JTL",
		cardNum:      262,
	},
	{
		abbreviation: "LOF",
		cardNum:      264,
	},
}

func main() {
	cards := make([]Card, 0)

	for _, currentSet := range sets {
		for i := 1; i < currentSet.cardNum+1; i++ {
			istr := numToString(i, 3)
			cardID := currentSet.abbreviation + istr

			card, err := getCardDetails(currentSet.abbreviation, istr, cardID)
			if err != nil {
				panic(fmt.Sprintf("failed to get card details: %s: %s", cardID, err.Error()))
			}

			cards = append(cards, card)
			fmt.Println(card)

			if err = downloadCardImage(currentSet.abbreviation, istr, cardID); err != nil {
				panic(fmt.Sprintf("failed to download card image: %s: %s", cardID, err.Error()))
			}
			fmt.Println("Saved image")

			time.Sleep(time.Second)
		}
	}

	fmt.Println(cards)
	// Generate and save the "goswucards.go" file

	// Read the template file
	tmplContent, err := os.ReadFile("template")
	if err != nil {
		panic(fmt.Sprintf("failed to read template file: %s", err.Error()))
	}

	// Parse the template
	tmpl, err := template.New("goswucards").Parse(string(tmplContent))
	if err != nil {
		panic(fmt.Sprintf("failed to parse template: %s", err.Error()))
	}

	// Create the output file
	outputFile, err := os.Create("../goswucards.go")
	if err != nil {
		panic(fmt.Sprintf("failed to create output file: %s", err.Error()))
	}
	defer outputFile.Close()

	// Execute the template with the cards data
	err = tmpl.Execute(outputFile, cards)
	if err != nil {
		panic(fmt.Sprintf("failed to execute template: %s", err.Error()))
	}

	fmt.Println("Successfully generated goswucards.go")
}

func getCardDetails(set, cardNumber, cardID string) (c Card, err error) {
	reqBody := map[string]any{
		"cardNumber":            cardNumber,
		"expansionAbbreviation": set,
		"isFoil":                nil,
		"language":              "",
	}

	respBody, err := makeRequest("https://swudb.com/api/card/getPrintingInfo", reqBody)
	if err != nil {
		err = fmt.Errorf("failed to make request: %s: %w", cardID, err)
		return
	}

	cardName, ok := respBody["cardName"]
	if !ok {
		err = fmt.Errorf("failed to get card name: %s", cardID)
		return
	}

	cardNameString, ok := cardName.(string)
	if !ok {
		err = fmt.Errorf("failed to coerce card name to string: %v", cardName)
		return
	}

	return Card{
		ID:   cardID,
		Name: cardNameString,
	}, nil
}

func downloadCardImage(set string, cardNumber, cardID string) (err error) {
	imageURL := fmt.Sprintf("https://swudb.com/images/cards/%s/%s.png", set, cardNumber)
	imageResp, err := http.Get(imageURL)
	if err != nil {
		err = fmt.Errorf("failed to GET image URL: %s: %w", cardID, imageResp)
		return
	}

	defer imageResp.Body.Close()
	imageData, err := io.ReadAll(imageResp.Body)
	if err != nil {
		err = fmt.Errorf("failed to read image data: %w", err)
		return
	}

	filename := fmt.Sprintf("../images/%s.png", cardID)
	if err = os.WriteFile(filename, imageData, 0644); err != nil {
		err = fmt.Errorf("failed to write image data to file: %w", err)
		return
	}

	return nil
}

func makeRequest(url string, body map[string]any) (response map[string]any, err error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		err = fmt.Errorf("failed to marshal body: %w", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		err = fmt.Errorf("failed to make request: %w", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Go-HTTP-Client/1.0")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("failed to do request: %w", err)
		return
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("failed to read response body: %w", err)
		return
	}

	if err = json.Unmarshal(respBody, &response); err != nil {
		err = fmt.Errorf("failed to unmarshal response: %w", err)
		return
	}

	return response, nil
}

func numToString(i int, l int) (s string) {
	s = fmt.Sprint(i)
	if len(s) < l {
		d := l - len(s)
		for w := 0; w < d; w++ {
			s = "0" + s
		}
	}
	return s
}
