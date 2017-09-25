package cli

import (
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	f "github.com/hortonworks/hdc-cli/client_cloudbreak/flexsubscriptions"
	"github.com/hortonworks/hdc-cli/models_cloudbreak"
	"github.com/urfave/cli"
)

var FlexSubscriptionHeader []string = []string{"Name", "SubscriptionID", "IsDefault", "UsedForController"}

type FlexSubscription struct {
	Name              string `json:"Name" yaml:"Name"`
	SubscriptionId    string `json:"SubscriptionID" yaml:"SubscriptionID"`
	IsDefault         string `json:"IsDefault" yaml:"IsDefault"`
	UsedForController string `json:"UsedForController" yaml:"UsedForController"`
}

func (b *FlexSubscription) DataAsStringArray() []string {
	return []string{b.Name, b.SubscriptionId, b.IsDefault, b.UsedForController}
}

func CreateFlexSubscription(c *cli.Context) error {
	ss, err := getAvailableSmartSenseSubscription(c)
	if err != nil {
		return err
	}
	defer timeTrack(time.Now(), "creation of Flex subscription")
	checkRequiredFlags(c)

	name := c.String(FlFlexSubscriptionName.Name)
	if len(name) == 0 {
		logMissingParameterAndExit(c, []string{FlFlexSubscriptionName.Name})
	}

	subscriptionId := c.String(FlFlexSubscriptionID.Name)
	if len(subscriptionId) == 0 {
		logMissingParameterAndExit(c, []string{FlFlexSubscriptionID.Name})
	}

	cbClient := NewCloudbreakOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))
	createFlexSubscriptionImpl(cbClient.Cloudbreak.Flexsubscriptions.PostPrivateFlexSubscription, ss, name, subscriptionId)
	return nil
}

func createFlexSubscriptionImpl(
	postPrivateFlexSubscription func(params *f.PostPrivateFlexSubscriptionParams) (*f.PostPrivateFlexSubscriptionOK, error),
	ss *models_cloudbreak.SmartSenseSubscriptionJSON,
	name string,
	subscriptionId string) {

	smartSenseSubscriptionId := ss.ID
	fa := false
	resp, err := postPrivateFlexSubscription(f.NewPostPrivateFlexSubscriptionParams().
		WithBody(
		&models_cloudbreak.FlexSubscriptionRequest{
			Name:                     &name,
			SubscriptionID:           subscriptionId,
			SmartSenseSubscriptionID: smartSenseSubscriptionId,
			Default:                  &fa,
			UsedForController:        &fa}))

	if err != nil {
		logErrorAndExit(err)
	}

	log.Infof("[createFlexSubscription] Flex subscription created: %v", *resp.Payload)
}

func DeleteFlexSubscription(c *cli.Context) error {
	defer timeTrack(time.Now(), "deletion of Flex subscription")
	checkRequiredFlags(c)

	name := c.String(FlFlexSubscriptionName.Name)
	if len(name) == 0 {
		logMissingParameterAndExit(c, []string{FlFlexSubscriptionName.Name})
	}

	cbClient := NewCloudbreakOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))
	deleteFlexSubscriptionImpl(cbClient.Cloudbreak.Flexsubscriptions.DeletePrivateFlexSubscriptionByName, name)
	log.Infof("[deleteFlexSubscription] Flex subscription has been deleted with name: %s ", name)
	return nil
}

func deleteFlexSubscriptionImpl(
	deletePublicFlexSubscriptionByName func(params *f.DeletePrivateFlexSubscriptionByNameParams) error,
	name string) {

	err := deletePublicFlexSubscriptionByName(f.NewDeletePrivateFlexSubscriptionByNameParams().WithName(name))

	if err != nil {
		logErrorAndExit(err)
	}
}

func ListFlexSubscriptions(c *cli.Context) error {
	defer timeTrack(time.Now(), "describe of SmartSenseSubscription")
	checkRequiredFlags(c)

	cbClient := NewCloudbreakOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))
	output := Output{Format: c.String(FlOutput.Name)}
	listFlexSubscriptionImpl(cbClient.Cloudbreak.Flexsubscriptions.GetPrivateFlexSubscriptions, output)
	return nil
}

func listFlexSubscriptionImpl(
	getPrivateFlexSubscriptions func(params *f.GetPrivateFlexSubscriptionsParams) (*f.GetPrivateFlexSubscriptionsOK, error),
	output Output) {

	resp, err := getPrivateFlexSubscriptions(f.NewGetPrivateFlexSubscriptionsParams())

	if err != nil {
		logErrorAndExit(err)
	}

	var tableRows []Row
	for _, subscription := range resp.Payload {
		row := &FlexSubscription{
			Name:              *subscription.Name,
			SubscriptionId:    subscription.SubscriptionID,
			IsDefault:         strconv.FormatBool(*subscription.Default),
			UsedForController: strconv.FormatBool(*subscription.UsedForController)}
		tableRows = append(tableRows, row)
		log.Infof("[listFlexSubscription] subscription name: %v ", subscription.Name)
	}

	log.Infof("[listFlexSubscription] number of available Flex subscriptions: %d ", len(resp.Payload))

	if len(resp.Payload) > 0 {
		output.WriteList(FlexSubscriptionHeader, tableRows)
	}
}

func SetFlexSubscriptionAsDefault(c *cli.Context) error {
	defer timeTrack(time.Now(), "set the Flex subscription as default")
	checkRequiredFlags(c)

	name := c.String(FlFlexSubscriptionName.Name)
	if len(name) == 0 {
		logMissingParameterAndExit(c, []string{FlFlexSubscriptionName.Name})
	}

	cbClient := NewCloudbreakOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))
	setFlexSubscriptionAsDefaultImpl(cbClient.Cloudbreak.Flexsubscriptions.PutPublicDefaultFlexSubscriptionByName, name)
	log.Infof("[setFlexSubscriptionAsDefault] Flex subscription has been set as default with name: %s ", name)
	return nil
}

func setFlexSubscriptionAsDefaultImpl(
	putPublicDefaultFlexSubscriptionByName func(params *f.PutPublicDefaultFlexSubscriptionByNameParams) error,
	name string) {

	err := putPublicDefaultFlexSubscriptionByName(f.NewPutPublicDefaultFlexSubscriptionByNameParams().WithName(name))
	if err != nil {
		logErrorAndExit(err)
	}
}

func UseFlexSubscriptionForController(c *cli.Context) error {
	defer timeTrack(time.Now(), "use-flexsubscription-for-controller")
	checkRequiredFlags(c)

	name := c.String(FlFlexSubscriptionName.Name)
	if len(name) == 0 {
		logMissingParameterAndExit(c, []string{FlFlexSubscriptionName.Name})
	}

	cbClient := NewCloudbreakOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))
	useFlexSubscriptionForControllerImpl(cbClient.Cloudbreak.Flexsubscriptions.PutPublicUsedForControllerFlexSubscriptionByName, name)
	log.Infof("[useFlexSubscriptionForController] Flex subscription has been set for Controller with name: %s ", name)
	return nil
}

func useFlexSubscriptionForControllerImpl(
	useForControllerPut func(params *f.PutPublicUsedForControllerFlexSubscriptionByNameParams) error,
	name string) {

	err := useForControllerPut(f.NewPutPublicUsedForControllerFlexSubscriptionByNameParams().WithName(name))
	if err != nil {
		logErrorAndExit(err)
	}
}

func (cb *Cloudbreak) GetFlexSubscriptionByName(name string) *models_cloudbreak.FlexSubscriptionResponse {
	defer timeTrack(time.Now(), "GetFlexSubscriptionByName")

	resp, err := cb.Cloudbreak.Flexsubscriptions.GetPrivateFlexSubscriptionByName(f.NewGetPrivateFlexSubscriptionByNameParams().WithName(name))

	if err != nil {
		logErrorAndExit(err)
	}
	log.Infof("[GetFlexSubscriptionByName] Flex subscription has been found for name: %s ", name)
	return resp.Payload
}
