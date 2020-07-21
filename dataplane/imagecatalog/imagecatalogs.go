package imagecatalog

import (
	"fmt"
	"strconv"
	"time"

	"github.com/hortonworks/cb-cli/dataplane/oauth"

	v4img "github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_imagecatalogs"
	"github.com/hortonworks/cb-cli/dataplane/api/model"
	"github.com/hortonworks/cb-cli/dataplane/cloud"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/dp-cli-common/utils"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var imagecatalogHeader = []string{"Name", "Description", "Default", "URL"}

type imagecatalogOut struct {
	Name        string `json:"Name" yaml:"Name"`
	Description string `json:"Description" yaml:"Description"`
	Default     bool   `json:"Default" yaml:"Default"`
	URL         string `json:"URL" yaml:"URL"`
}

type imagecatalogOutDescribe struct {
	*imagecatalogOut
	CRN string `json:"CRN" yaml:"CRN"`
}

func (r *imagecatalogOut) DataAsStringArray() []string {
	return []string{r.Name, r.Description, strconv.FormatBool(r.Default), r.URL}
}

func (b *imagecatalogOutDescribe) DataAsStringArray() []string {
	return append(b.imagecatalogOut.DataAsStringArray(), b.CRN)
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

	log.Infof("[CreateImagecatalogFromUrl] creating imagecatalog from a URL")
	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	createImagecatalogImpl(
		cbClient.Cloudbreak.V4WorkspaceIDImagecatalogs,
		c.Int64(fl.FlWorkspaceOptional.Name),
		c.String(fl.FlName.Name),
		c.String(fl.FlDescriptionOptional.Name),
		c.String(fl.FlURL.Name))
}

type imageCatalogClient interface {
	CreateImageCatalogInWorkspace(params *v4img.CreateImageCatalogInWorkspaceParams) (*v4img.CreateImageCatalogInWorkspaceOK, error)
}

func createImagecatalogImpl(client imageCatalogClient, workspaceID int64, name string, description string, imagecatalogURL string) {
	defer utils.TimeTrack(time.Now(), "create imagecatalog")
	imageCatalogRequest := &model.ImageCatalogV4Request{
		Name:        &name,
		Description: &description,
		URL:         &imagecatalogURL,
	}
	log.Infof("[createImagecatalogImpl] sending create imagecatalog request")
	resp, err := client.CreateImageCatalogInWorkspace(v4img.NewCreateImageCatalogInWorkspaceParams().WithWorkspaceID(workspaceID).WithBody(imageCatalogRequest))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	ic := resp.Payload
	log.Infof("[createImagecatalogImpl] imagecatalog created: %s (crn: %s)", *ic.Name, ic.Crn)
}

func ListImagecatalogs(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "list imagecatalogs")

	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	listImagecatalogsImpl(cbClient.Cloudbreak.V4WorkspaceIDImagecatalogs, workspaceID, output.WriteList)
}

type listImageCatalogsByWorkspaceClient interface {
	ListImageCatalogsByWorkspace(*v4img.ListImageCatalogsByWorkspaceParams) (*v4img.ListImageCatalogsByWorkspaceOK, error)
}

func listImagecatalogsImpl(client listImageCatalogsByWorkspaceClient, workspaceID int64, writer func([]string, []utils.Row)) {
	log.Infof("[listImagecatalogsImpl] sending imagecatalog list request")
	imagecatalogResp, err := client.ListImageCatalogsByWorkspace(v4img.NewListImageCatalogsByWorkspaceParams().WithWorkspaceID(workspaceID))
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	var tableRows []utils.Row
	for _, ic := range imagecatalogResp.Payload.Responses {
		tableRows = append(tableRows, &imagecatalogOut{*ic.Name, utils.SafeStringConvert(ic.Description), *ic.UsedAsDefault, *ic.URL})
	}

	writer(imagecatalogHeader, tableRows)
}

func DeleteImagecatalog(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "delete imagecatalog")

	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	name := c.String(fl.FlName.Name)
	log.Infof("[DeleteImagecatalog] sending delete imagecatalog request with name: %s", name)
	if _, err := cbClient.Cloudbreak.V4WorkspaceIDImagecatalogs.DeleteImageCatalogInWorkspace(v4img.NewDeleteImageCatalogInWorkspaceParams().WithWorkspaceID(workspaceID).WithName(name)); err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[DeleteImagecatalog] imagecatalog deleted, name: %s", name)
}

func DescribeImagecatalog(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "describe imagecatalog")

	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	resp, err := cbClient.Cloudbreak.V4WorkspaceIDImagecatalogs.GetImageCatalogInWorkspace(v4img.NewGetImageCatalogInWorkspaceParams().WithWorkspaceID(workspaceID).WithName(c.String(fl.FlName.Name)))
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	imgc := resp.Payload
	if len(imgc.Crn) == 0 {
		output.Write(imagecatalogHeader, &imagecatalogOut{*imgc.Name, utils.SafeStringConvert(imgc.Description), *imgc.UsedAsDefault, *imgc.URL})
	} else {
		output.Write(append(imagecatalogHeader, "CRN"), &imagecatalogOutDescribe{&imagecatalogOut{*imgc.Name, utils.SafeStringConvert(imgc.Description), *imgc.UsedAsDefault, *imgc.URL}, imgc.Crn})
	}
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
	defer utils.TimeTrack(time.Now(), "list available images")

	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	listImagesImpl(cbClient.Cloudbreak.V4WorkspaceIDImagecatalogs, output.WriteList, c.Int64(fl.FlWorkspaceOptional.Name), c.String(fl.FlImageCatalog.Name))
}

func ListImagesValidForUpgrade(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "list images valid for stack upgrade")

	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)

	clusterName := c.String(fl.FlClusterToUpgrade.Name)
	imageCatalogName := c.String(fl.FlImageCatalogOptional.Name)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	if imageCatalogName != "" {
		log.Infof("[ListImagesValidForUpgrade] sending list images request, stack: %s, catalog: %s", clusterName, imageCatalogName)
		imageResp, err := cbClient.Cloudbreak.V4WorkspaceIDImagecatalogs.GetImagesByNameInWorkspace(v4img.NewGetImagesByNameInWorkspaceParams().WithWorkspaceID(workspaceID).WithName(imageCatalogName).WithStackName(&clusterName))
		if err != nil {
			utils.LogErrorAndExit(err)
		}
		writeImageListInformation(output.WriteList, toImageResponseList(imageResp.Payload))
	} else {
		log.Infof("[ListImagesValidForUpgrade] sending list images request using the default catalog, stack: %s", clusterName)
		imageResp, err := cbClient.Cloudbreak.V4WorkspaceIDImagecatalogs.GetImagesInWorkspace(v4img.NewGetImagesInWorkspaceParams().WithWorkspaceID(workspaceID).WithStackName(&clusterName))
		if err != nil {
			utils.LogErrorAndExit(err)
		}
		writeImageListInformation(output.WriteList, toImageResponseList(imageResp.Payload))
	}
}

func toImageResponseList(images *model.ImagesV4Response) []*model.ImageV4Response {
	var imagesList = make([]*model.ImageV4Response, 0, len(images.BaseImages)+len(images.CdhImages)+len(images.CdhImages)+len(images.CdhImages))
	for _, i := range images.BaseImages {
		imagesList = append(imagesList, toImageResponse(i))
	}
	for _, i := range images.CdhImages {
		imagesList = append(imagesList, i)
	}
	for _, i := range images.CdhImages {
		imagesList = append(imagesList, i)
	}
	for _, i := range images.CdhImages {
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
	defer utils.TimeTrack(time.Now(), "describe image")

	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	describeImageImpl(cbClient.Cloudbreak.V4WorkspaceIDImagecatalogs, output.WriteList, c.Int64(fl.FlWorkspaceOptional.Name), c.String(fl.FlImageCatalog.Name), c.String(fl.FlImageId.Name))
}

func describeImageImpl(client getPublicImagesClient, writer func([]string, []utils.Row), workspaceID int64, imagecatalog string, imageid string) {
	log.Infof("[listImagesImpl] sending list images request")
	provider := cloud.GetProvider().GetName()
	imageResp, err := client.GetImagesByNameInWorkspace(v4img.NewGetImagesByNameInWorkspaceParams().WithWorkspaceID(workspaceID).WithName(imagecatalog).WithPlatform(provider))
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	image := findImageByUUID(imageResp.Payload, imageid)
	if image == nil {
		utils.LogErrorMessage(fmt.Sprintf("Image not found by id: %s for cloud: %s", imageid, *provider))
	}

	writeImageInformation(writer, image)
}

func findImageByUUID(imageResponse *model.ImagesV4Response, imageID string) *model.ImageV4Response {
	image := findBaseImage(imageResponse.BaseImages, imageID)
	if image == nil {
		warmupImage := findWarmupImage(imageResponse.CdhImages, imageID)
		if warmupImage == nil {
			warmupImage = findWarmupImage(imageResponse.CdhImages, imageID)
			if warmupImage == nil {
				warmupImage = findWarmupImage(imageResponse.CdhImages, imageID)
			}
		}
		return warmupImage
	}
	return toImageResponse(image)
}

func toImageResponse(image *model.BaseImageV4Response) *model.ImageV4Response {
	return &model.ImageV4Response{
		Date:            image.Date,
		Description:     image.Description,
		Version:         image.Version,
		UUID:            image.UUID,
		Os:              image.Os,
		OsType:          image.OsType,
		Images:          image.Images,
		PackageVersions: image.PackageVersions,
	}
}

func findBaseImage(images []*model.BaseImageV4Response, imageID string) *model.BaseImageV4Response {
	for _, i := range images {
		if i.UUID == imageID {
			return i
		}
	}
	return nil
}

func findWarmupImage(images []*model.ImageV4Response, imageID string) *model.ImageV4Response {
	for _, i := range images {
		if i.UUID == imageID {
			return i
		}
	}
	return nil
}

func writeImageInformation(writer func([]string, []utils.Row), image *model.ImageV4Response) {
	var tableRows []utils.Row
	tableRows = append(tableRows, &imageDetailsOut{image.Date, image.Description, image.Version, image.UUID, image.Os, image.OsType, image.Images, image.PackageVersions})
	writer(imageDetailsHeader, tableRows)
}

func writeImageListInformation(writer func([]string, []utils.Row), images []*model.ImageV4Response) {
	var tableRows []utils.Row
	for _, i := range images {
		tableRows = append(tableRows, &imageDetailsOut{i.Date, i.Description, i.Version, i.UUID, i.Os, i.OsType, i.Images, i.PackageVersions})
	}
	writer(imageDetailsHeader, tableRows)
}

type getPublicImagesClient interface {
	GetImagesByNameInWorkspace(*v4img.GetImagesByNameInWorkspaceParams) (*v4img.GetImagesByNameInWorkspaceOK, error)
}

func listImagesImpl(client getPublicImagesClient, writer func([]string, []utils.Row), workspaceID int64, imagecatalog string) {
	log.Infof("[listImagesImpl] sending list images request")
	provider := cloud.GetProvider().GetName()
	imageResp, err := client.GetImagesByNameInWorkspace(v4img.NewGetImagesByNameInWorkspaceParams().WithWorkspaceID(workspaceID).WithName(imagecatalog).WithPlatform(provider))
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	var tableRows []utils.Row
	for _, i := range imageResp.Payload.BaseImages {
		tableRows = append(tableRows, &imageOut{i.Date, i.Description, i.Version, i.UUID})
	}
	for _, i := range imageResp.Payload.CdhImages {
		tableRows = append(tableRows, &imageOut{i.Date, i.Description, i.Version, i.UUID})
	}
	for _, i := range imageResp.Payload.CdhImages {
		tableRows = append(tableRows, &imageOut{i.Date, i.Description, i.Version, i.UUID})
	}
	for _, i := range imageResp.Payload.CdhImages {
		tableRows = append(tableRows, &imageOut{i.Date, i.Description, i.Version, i.UUID})
	}

	writer(imageHeader, tableRows)

}
