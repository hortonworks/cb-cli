package cli

import (
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	s "github.com/hortonworks/hdc-cli/client_cloudbreak/smartsensesubscriptions"
	"github.com/hortonworks/hdc-cli/models_cloudbreak"
	"github.com/urfave/cli"
)

var SmartSenseSubscriptionHeader []string = []string{"ID", "SubscriptionID"}

type SmartSenseSubscription struct {
	Id             string `json:"ID" yaml:"ID"`
	SubscriptionId string `json:"SubscriptionID" yaml:"SubscriptionID"`
}

func (b *SmartSenseSubscription) DataAsStringArray() []string {
	return []string{b.Id, b.SubscriptionId}
}

func CreateSmartSenseSubscription(c *cli.Context) error {
	defer timeTrack(time.Now(), "creation of SmartSenseSubscription")
	checkRequiredFlags(c)

	subscriptionId := c.String(FlSmartSenseSubscription.Name)
	if len(subscriptionId) == 0 {
		logMissingParameterAndExit(c, []string{FlSmartSenseSubscription.Name})
	}

	cbClient := NewCloudbreakOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))
	createSmartSenseSubscriptionImpl(subscriptionId,
		cbClient.Cloudbreak.Smartsensesubscriptions.PostPrivateSmartSenseSubscription)
	return nil
}

func createSmartSenseSubscriptionImpl(subscriptionId string,
	postPrivateSmartSenseSubscription func(*s.PostPrivateSmartSenseSubscriptionParams) (*s.PostPrivateSmartSenseSubscriptionOK, error)) {

	log.Infof("[createSmartSenseSubscription] create SmartSense subscription for: %s ", subscriptionId)
	resp, err := postPrivateSmartSenseSubscription(
		&s.PostPrivateSmartSenseSubscriptionParams{Body: &models_cloudbreak.SmartSenseSubscriptionJSON{SubscriptionID: subscriptionId}})

	if err != nil {
		logErrorAndExit(err)
	}

	log.Infof("[createSmartSenseSubscription] SmartSense subscription created: %v", *resp.Payload)
}

func DeleteSmartSenseSubscription(c *cli.Context) error {
	defer timeTrack(time.Now(), "deletion of SmartSenseSubscription")
	checkRequiredFlags(c)

	id := c.String(FlSmartSenseSubscriptionID.Name)
	if len(id) == 0 {
		logMissingParameterAndExit(c, []string{FlSmartSenseSubscriptionID.Name})
	}

	cbClient := NewCloudbreakOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))
	deleteSmartSenseSubscriptionImpl(cbClient.Cloudbreak.Smartsensesubscriptions.DeleteSmartsensesubscriptionsUserSubscriptionID, id)
	log.Infof("[deleteSmartSenseSubscription] SmartSense subscription has been deleted with id: %s ", id)
	return nil
}

func deleteSmartSenseSubscriptionImpl(deleteSmartSenseSubscriptionByID func(params *s.DeleteSmartsensesubscriptionsUserSubscriptionIDParams) error, id string) {
	err := deleteSmartSenseSubscriptionByID(&s.DeleteSmartsensesubscriptionsUserSubscriptionIDParams{SubscriptionID: id})

	if err != nil {
		logErrorAndExit(err)
	}
}

func DescribeSmartSenseSubscription(c *cli.Context) error {
	defer timeTrack(time.Now(), "describe of SmartSenseSubscription")
	checkRequiredFlags(c)

	cbClient := NewCloudbreakOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))
	output := Output{Format: c.String(FlOutput.Name)}
	describeSmartSenseSubscriptionImpl(cbClient.Cloudbreak.Smartsensesubscriptions.GetPrivateSmartSenseSubscriptions, output)
	return nil
}

func describeSmartSenseSubscriptionImpl(
	getPrivateSmartSenseSubscriptions func(params *s.GetPrivateSmartSenseSubscriptionsParams) (*s.GetPrivateSmartSenseSubscriptionsOK, error),
	output Output) {

	resp, err := getPrivateSmartSenseSubscriptions(&s.GetPrivateSmartSenseSubscriptionsParams{})

	if err != nil {
		logErrorAndExit(err)
	}

	var tableRows []Row
	for _, subscription := range resp.Payload {
		row := &SmartSenseSubscription{Id: strconv.FormatInt(*subscription.ID, 10), SubscriptionId: subscription.SubscriptionID}
		tableRows = append(tableRows, row)
		log.Infof("[describeSmartSenseSubscription] describe SmartSense subscription for: %v ", subscription.SubscriptionID)
	}

	output.WriteList(SmartSenseSubscriptionHeader, tableRows)
}
