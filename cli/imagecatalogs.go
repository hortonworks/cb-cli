package cli

import (
	"fmt"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/cb-cli/cli/cloud"
	"github.com/hortonworks/cb-cli/cli/utils"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v3_organization_id_imagecatalogs"
	"github.com/hortonworks/cb-cli/models_cloudbreak"
	"github.com/urfave/cli"
)

var imagecatalogHeader = []string{"Name", "Default", "URL"}

type imagecatalogOut struct {
	Name    string `json:"Name" yaml:"Name"`
	Default bool   `json:"Default" yaml:"Default"`
	URL     string `json:"URL" yaml:"URL"`
}

func (r *imagecatalogOut) DataAsStringArray() []string {
	return []string{r.Name, strconv.FormatBool(r.Default), r.URL}
}

var imageHeader = []string{"Date", "Description", "Version", "ImageID"}

type imageOut struct {
	Date        string `json:"Date" yaml:"Date"`
	Description string `json:"Description" yaml:"Description"`
	Version     string `json:"Version" yaml:"Version"`
	ImageID     string `json:"ImageID" yaml:"ImageID"`
}

func (r *imageOut) DataAsStringArray() []string {
	return []string{r.Date, r.Description, r.Version, r.ImageID}
}

var imageDetailsHeader = []string{"Date", "Description", "Ambari Version", "ImageID", "OS", "OS Type", "Images", "Package Versions"}

type imageDetailsOut struct {
	Date            string                       `json:"Date" yaml:"Date"`
	Description     string                       `json:"Description" yaml:"Description"`
	Version         string                       `json:"Version" yaml:"Version"`
	ImageID         string                       `json:"ImageID" yaml:"ImageID"`
	Os              string                       `json:"OS" yaml:"OS"`
	OsType          string                       `json:"OSType" yaml:"OSType"`
	Images          map[string]map[string]string `json:"Images" yaml:"Images"`
	PackageVersions map[string]string            `json:"PackageVersions" yaml:"PackageVersions"`
}

func (r *imageDetailsOut) DataAsStringArray() []string {
	var images string
	for prov, imgs := range r.Images {
		images += fmt.Sprintf("%s:\n", prov)
		for region, image := range imgs {
			images += fmt.Sprintf("  %s: %s\n", region, image)
		}
	}

	var packageVersions string
	for pkg, ver := range r.PackageVersions {
		packageVersions += fmt.Sprintf("%s: %s\n", pkg, ver)
	}

	return []string{r.Date, r.Description, r.Version, r.ImageID, r.Os, r.OsType, images, packageVersions}
}

func CreateImagecatalogFromUrl(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)

	log.Infof("[CreateImagecatalogFromUrl] creating imagecatalog from a URL")
	cbClient := NewCloudbreakHTTPClientFromContext(c)
	createImagecatalogImpl(
		cbClient.Cloudbreak.V3OrganizationIDImagecatalogs,
		c.Int64(FlOrganizationOptional.Name),
		c.String(FlName.Name),
		c.String(FlURL.Name))
}

type imagecatalogClient interface {
	CreateImageCatalogInOrganization(params *v3_organization_id_imagecatalogs.CreateImageCatalogInOrganizationParams) (*v3_organization_id_imagecatalogs.CreateImageCatalogInOrganizationOK, error)
}

func createImagecatalogImpl(client imagecatalogClient, orgID int64, name string, imagecatalogURL string) {
	defer utils.TimeTrack(time.Now(), "create imagecatalog")
	imagecatalogRequest := &models_cloudbreak.ImageCatalogRequest{
		Name: &name,
		URL:  &imagecatalogURL,
	}
	var ic *models_cloudbreak.ImageCatalogResponse
	log.Infof("[createImagecatalogImpl] sending create imagecatalog request")
	resp, err := client.CreateImageCatalogInOrganization(v3_organization_id_imagecatalogs.NewCreateImageCatalogInOrganizationParams().WithOrganizationID(orgID).WithBody(imagecatalogRequest))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	ic = resp.Payload
	log.Infof("[createImagecatalogImpl] imagecatalog created: %s (id: %d)", *ic.Name, ic.ID)
}

func ListImagecatalogs(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "list imagecatalogs")

	cbClient := NewCloudbreakHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(FlOutputOptional.Name)}
	orgID := c.Int64(FlOrganizationOptional.Name)
	listImagecatalogsImpl(cbClient.Cloudbreak.V3OrganizationIDImagecatalogs, orgID, output.WriteList)
}

type listImageCatalogsByOrganizationClient interface {
	ListImageCatalogsByOrganization(*v3_organization_id_imagecatalogs.ListImageCatalogsByOrganizationParams) (*v3_organization_id_imagecatalogs.ListImageCatalogsByOrganizationOK, error)
}

func listImagecatalogsImpl(client listImageCatalogsByOrganizationClient, orgID int64, writer func([]string, []utils.Row)) {
	log.Infof("[listImagecatalogsImpl] sending imagecatalog list request")
	imagecatalogResp, err := client.ListImageCatalogsByOrganization(v3_organization_id_imagecatalogs.NewListImageCatalogsByOrganizationParams().WithOrganizationID(orgID))
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	tableRows := []utils.Row{}
	for _, ic := range imagecatalogResp.Payload {
		tableRows = append(tableRows, &imagecatalogOut{*ic.Name, ic.UsedAsDefault, *ic.URL})
	}

	writer(imagecatalogHeader, tableRows)
}

func DeleteImagecatalog(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "delete imagecatalog")

	cbClient := NewCloudbreakHTTPClientFromContext(c)
	orgID := c.Int64(FlOrganizationOptional.Name)
	name := c.String(FlName.Name)
	log.Infof("[DeleteImagecatalog] sending delete imagecatalog request with name: %s", name)
	if _, err := cbClient.Cloudbreak.V3OrganizationIDImagecatalogs.DeleteImageCatalogInOrganization(v3_organization_id_imagecatalogs.NewDeleteImageCatalogInOrganizationParams().WithOrganizationID(orgID).WithName(name)); err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[DeleteImagecatalog] imagecatalog deleted, name: %s", name)
}

func SetDefaultImagecatalog(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "set default imagecatalog")

	cbClient := NewCloudbreakHTTPClientFromContext(c)
	orgID := c.Int64(FlOrganizationOptional.Name)
	name := c.String(FlName.Name)
	log.Infof("[SetDefautlImagecatalog] sending set default imagecatalog request with name: %s", name)

	if _, err := cbClient.Cloudbreak.V3OrganizationIDImagecatalogs.PutSetDefaultImageCatalogByNameInOrganization(v3_organization_id_imagecatalogs.NewPutSetDefaultImageCatalogByNameInOrganizationParams().WithOrganizationID(orgID).WithName(name)); err != nil {
		utils.LogErrorAndExit(err)
	}

	log.Infof("[SetDefaultImagecatalog] imagecatalog is set as default, name: %s", name)
}

func ListAwsImages(c *cli.Context) {
	cloud.SetProviderType(cloud.AWS)
	listImages(c)
}

func ListAzureImages(c *cli.Context) {
	cloud.SetProviderType(cloud.AZURE)
	listImages(c)
}

func ListGcpImages(c *cli.Context) {
	cloud.SetProviderType(cloud.GCP)
	listImages(c)
}

func ListOpenstackImages(c *cli.Context) {
	cloud.SetProviderType(cloud.OPENSTACK)
	listImages(c)
}

func listImages(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "list available images")

	cbClient := NewCloudbreakHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(FlOutputOptional.Name)}
	listImagesImpl(cbClient.Cloudbreak.V3OrganizationIDImagecatalogs, output.WriteList, c.Int64(FlOrganizationOptional.Name), c.String(FlImageCatalog.Name))
}

func ListImagesValidForUpgrade(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "list images valid for stack upgrade")

	cbClient := NewCloudbreakHTTPClientFromContext(c)

	clusterName := c.String(FlClusterToUpgrade.Name)
	imageCatalogName := c.String(FlImageCatalogOptional.Name)
	output := utils.Output{Format: c.String(FlOutputOptional.Name)}
	orgID := c.Int64(FlOrganizationOptional.Name)
	if imageCatalogName != "" {
		log.Infof("[ListImagesValidForUpgrade] sending list images request, stack: %s, catalog: %s", clusterName, imageCatalogName)
		imageResp, err := cbClient.Cloudbreak.V3OrganizationIDImagecatalogs.GetImagesByStackNameAndCustomImageCatalogInOrganization(v3_organization_id_imagecatalogs.NewGetImagesByStackNameAndCustomImageCatalogInOrganizationParams().WithOrganizationID(orgID).WithName(imageCatalogName).WithStackName(clusterName))
		if err != nil {
			utils.LogErrorAndExit(err)
		}
		writeImageListInformation(output.WriteList, toImageResponseList(imageResp.Payload))
	} else {
		log.Infof("[ListImagesValidForUpgrade] sending list images request using the default catalog, stack: %s", clusterName)
		imageResp, err := cbClient.Cloudbreak.V3OrganizationIDImagecatalogs.GetImagesByStackNameAndDefaultImageCatalogInOrganization(v3_organization_id_imagecatalogs.NewGetImagesByStackNameAndDefaultImageCatalogInOrganizationParams().WithOrganizationID(orgID).WithStackName(clusterName))
		if err != nil {
			utils.LogErrorAndExit(err)
		}
		writeImageListInformation(output.WriteList, toImageResponseList(imageResp.Payload))
	}
}

func toImageResponseList(images *models_cloudbreak.ImagesResponse) []*models_cloudbreak.ImageResponse {
	var imagesList = make([]*models_cloudbreak.ImageResponse, 0, len(images.BaseImages)+len(images.HdfImages)+len(images.HdpImages))
	for _, i := range images.BaseImages {
		imagesList = append(imagesList, toImageResponse(i))
	}
	for _, i := range images.HdfImages {
		imagesList = append(imagesList, i)
	}
	for _, i := range images.HdpImages {
		imagesList = append(imagesList, i)
	}

	return imagesList
}

func DescribeAwsImage(c *cli.Context) {
	cloud.SetProviderType(cloud.AWS)
	describeImage(c)
}

func DescribeAzureImage(c *cli.Context) {
	cloud.SetProviderType(cloud.AZURE)
	describeImage(c)
}

func DescribeGcpImage(c *cli.Context) {
	cloud.SetProviderType(cloud.GCP)
	describeImage(c)
}

func DescribeOpenstackImage(c *cli.Context) {
	cloud.SetProviderType(cloud.OPENSTACK)
	describeImage(c)
}

func describeImage(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "describe image")

	cbClient := NewCloudbreakHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(FlOutputOptional.Name)}
	describeImageImpl(cbClient.Cloudbreak.V3OrganizationIDImagecatalogs, output.WriteList, c.Int64(FlOrganizationOptional.Name), c.String(FlImageCatalog.Name), c.String(FlImageId.Name))
}

func describeImageImpl(client getPublicImagesClient, writer func([]string, []utils.Row), orgID int64, imagecatalog string, imageid string) {
	log.Infof("[listImagesImpl] sending list images request")
	provider := cloud.GetProvider().GetName()
	imageResp, err := client.GetImagesByProviderAndCustomImageCatalogInOrganization(v3_organization_id_imagecatalogs.NewGetImagesByProviderAndCustomImageCatalogInOrganizationParams().WithOrganizationID(orgID).WithName(imagecatalog).WithPlatform(*provider))
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	image := findImageByUUID(imageResp.Payload, imageid)
	if image == nil {
		utils.LogErrorMessage(fmt.Sprintf("Image not found by id: %s for cloud: %s", imageid, *provider))
	}

	writeImageInformation(writer, image)
}

func findImageByUUID(imageResponse *models_cloudbreak.ImagesResponse, imageID string) *models_cloudbreak.ImageResponse {
	image := findBaseImage(imageResponse.BaseImages, imageID)
	if image == nil {
		warmupImage := findWarmupImage(imageResponse.HdpImages, imageID)
		if warmupImage == nil {
			warmupImage = findWarmupImage(imageResponse.HdfImages, imageID)
		}
		return warmupImage
	}
	return toImageResponse(image)
}

func toImageResponse(image *models_cloudbreak.BaseImageResponse) *models_cloudbreak.ImageResponse {
	return &models_cloudbreak.ImageResponse{
		Date:            image.Date,
		Description:     image.Description,
		Version:         "",
		UUID:            image.UUID,
		Os:              image.Os,
		OsType:          image.OsType,
		Images:          image.Images,
		PackageVersions: image.PackageVersions,
	}
}

func findBaseImage(images []*models_cloudbreak.BaseImageResponse, imageID string) *models_cloudbreak.BaseImageResponse {
	for _, i := range images {
		if i.UUID == imageID {
			return i
		}
	}
	return nil
}

func findWarmupImage(images []*models_cloudbreak.ImageResponse, imageID string) *models_cloudbreak.ImageResponse {
	for _, i := range images {
		if i.UUID == imageID {
			return i
		}
	}
	return nil
}

func writeImageInformation(writer func([]string, []utils.Row), image *models_cloudbreak.ImageResponse) {
	tableRows := []utils.Row{}
	tableRows = append(tableRows, &imageDetailsOut{image.Date, image.Description, image.Version, image.UUID, image.Os, image.OsType, image.Images, image.PackageVersions})
	writer(imageDetailsHeader, tableRows)
}

func writeImageListInformation(writer func([]string, []utils.Row), payload []*models_cloudbreak.ImageResponse) {
	tableRows := []utils.Row{}

	for _, i := range payload {
		tableRows = append(tableRows, &imageDetailsOut{i.Date, i.Description, i.Version, i.UUID, i.Os, i.OsType, i.Images, i.PackageVersions})
	}
	writer(imageDetailsHeader, tableRows)
}

type getPublicImagesClient interface {
	GetImagesByProviderAndCustomImageCatalogInOrganization(*v3_organization_id_imagecatalogs.GetImagesByProviderAndCustomImageCatalogInOrganizationParams) (*v3_organization_id_imagecatalogs.GetImagesByProviderAndCustomImageCatalogInOrganizationOK, error)
}

func listImagesImpl(client getPublicImagesClient, writer func([]string, []utils.Row), orgID int64, imagecatalog string) {
	log.Infof("[listImagesImpl] sending list images request")
	provider := cloud.GetProvider().GetName()
	imageResp, err := client.GetImagesByProviderAndCustomImageCatalogInOrganization(v3_organization_id_imagecatalogs.NewGetImagesByProviderAndCustomImageCatalogInOrganizationParams().WithOrganizationID(orgID).WithName(imagecatalog).WithPlatform(*provider))
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	tableRows := []utils.Row{}
	for _, i := range imageResp.Payload.BaseImages {
		tableRows = append(tableRows, &imageOut{i.Date, i.Description, i.Version, i.UUID})
	}

	writer(imageHeader, tableRows)

}
