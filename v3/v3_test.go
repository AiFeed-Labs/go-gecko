package coingecko

import (
	"io/ioutil"
	"os"
	"testing"

	gock "gopkg.in/h2non/gock.v1"
)

func init() {
	defer gock.Off()
}

var c = NewClient(nil)
var mockURL = "https://api.coingecko.com/api/v3"

func TestPing(t *testing.T) {
	if err := setupGock("json/ping.json", "/ping"); err != nil {
		t.Errorf("setupGock failed: %v", err)
		t.FailNow()
	}
	ping, err := c.Ping()
	if err != nil {
		t.Errorf("Ping failed: %v", err)
		t.FailNow()
	}
	if ping.GeckoSays != "(V3) To the Moon!" {
		t.Errorf("Ping rsp not right, want: %v, got: %v", "(V3) To the Moon!", ping.GeckoSays)
		t.FailNow()
	}
}

func TestSimpleSinglePrice(t *testing.T) {
	if err := setupGock("json/simple_single_price.json", "/simple/price"); err != nil {
		t.Errorf("setupGock failed: %v", err)
		t.FailNow()
	}
	simplePrice, err := c.SimpleSinglePrice("bitcoin", "usd")
	if err != nil {
		t.Errorf("SimpleSinglePrice failed: %v", err)
		t.FailNow()
	}
	if simplePrice.ID != "bitcoin" || simplePrice.Currency != "usd" || simplePrice.MarketPrice != float64(5013.61) {
		t.Error("SimpleSinglePrice rsp not right")
		t.FailNow()
	}
}

func TestSimplePrice(t *testing.T) {
	if err := setupGock("json/simple_price.json", "/simple/price"); err != nil {
		t.Errorf("setupGock failed: %v", err)
		t.FailNow()
	}
	ids := []string{"bitcoin", "ethereum"}
	vc := []string{"usd", "myr"}
	sp, err := c.SimplePrice(ids, vc)
	if err != nil {
		t.Errorf("SimplePrice failed: %v", err)
		t.FailNow()
	}
	bitcoin := (*sp)["bitcoin"]
	eth := (*sp)["ethereum"]

	if bitcoin["usd"] != 5005.73 || bitcoin["myr"] != 20474 {
		t.Error("SimplePrice bitcoin rsp not right")
		t.FailNow()
	}
	if eth["usd"] != 163.58 || eth["myr"] != 669.07 {
		t.Error("SimplePrice eth rsp not right")
		t.FailNow()
	}
}

func TestSimpleSupportedVSCurrencies(t *testing.T) {
	if err := setupGock("json/simple_supported_vs_currencies.json", "/simple/supported_vs_currencies"); err != nil {
		t.Errorf("setupGock failed: %v", err)
		t.FailNow()
	}
	s, err := c.SimpleSupportedVSCurrencies()
	if err != nil {
		t.Errorf("SimpleSupportedVSCurrencies failed: %v", err)
		t.FailNow()
	}
	if len(*s) != 54 {
		t.Error("SimpleSupportedVSCurrencies rsp not right")
		t.FailNow()
	}
}

func TestCoinsList(t *testing.T) {
	if err := setupGock("json/coins_list.json", "/coins/list"); err != nil {
		t.Errorf("setupGock failed: %v", err)
		t.FailNow()
	}
	list, err := c.CoinsList()
	if err != nil {
		t.Errorf("CoinsList failed: %v", err)
		t.FailNow()
	}
	item := (*list)[0]
	if item.ID != "01coin" {
		t.Error("CoinsList rsp not right")
		t.FailNow()
	}
}

// Util: Setup Gock
func setupGock(filename string, url string) error {
	testJSON, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer testJSON.Close()

	testByte, err := ioutil.ReadAll(testJSON)
	if err != nil {
		return err
	}

	gock.New(mockURL).
		Get(url).
		Reply(200).
		JSON(testByte)

	return nil
}
