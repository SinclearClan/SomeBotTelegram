package dhl

import (
	"github.com/SinclearClan/SomeBotTelegram/config"
	"gitlab.com/equalsGibson/godhl"
)

var (
	cfg = config.GetConfig()
)

func GetDHLPackageInfo(trackingNumber string) (*godhl.Payload, error) {

	dhl_client := godhl.DHLClient{
		ApiKey: cfg.DHL.Key,
		ApiEndpoint: cfg.DHL.Endpoint,
	}
	trackingPackage, err := dhl_client.SendAPIRequest(trackingNumber)
	if err != nil {
		return nil, err
	}

	return trackingPackage.Payload, nil

}
