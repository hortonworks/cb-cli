package imagecatalog

import (
	"strings"
	"testing"

	v4img "github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_imagecatalogs"
	"github.com/hortonworks/cb-cli/dataplane/api/model"
	"github.com/hortonworks/cb-cli/dataplane/cloud"
	_ "github.com/hortonworks/cb-cli/dataplane/cloud/aws"
	"github.com/hortonworks/cb-cli/dataplane/types"
	"github.com/hortonworks/dp-cli-common/utils"
)

type mockListImageCatalogsByWorkspaceClient struct {
}

func (*mockListImageCatalogsByWorkspaceClient) ListImageCatalogsByWorkspace(params *v4img.ListImageCatalogsByWorkspaceParams) (*v4img.ListImageCatalogsByWorkspaceOK, error) {
	resp := []*model.ImageCatalogV4Response{
		{
			Name:          &(&types.S{S: "test"}).S,
			Description:   &(&types.S{S: "description"}).S,
			UsedAsDefault: &(&types.B{B: true}).B,
			URL:           &(&types.S{S: "testurl"}).S,
		},
	}
	return &v4img.ListImageCatalogsByWorkspaceOK{Payload: &model.ImageCatalogV4Responses{Responses: resp}}, nil
}

func TestListImagecatalogsImpl(t *testing.T) {
	t.Parallel()

	var rows []utils.Row
	listImagecatalogsImpl(new(mockListImageCatalogsByWorkspaceClient), int64(2), func(h []string, r []utils.Row) { rows = r })
	if len(rows) != 1 {
		t.Fatalf("row number doesn't match 1 == %d", len(rows))
	}
	for _, r := range rows {
		expected := "test description true testurl"
		if strings.Join(r.DataAsStringArray(), " ") != expected {
			t.Errorf("row data not match %s == %s", expected, strings.Join(r.DataAsStringArray(), " "))
		}
	}
}

type mockGetPublicImagesClient struct {
}

func (*mockGetPublicImagesClient) GetImagesByNameInWorkspace(params *v4img.GetImagesByNameInWorkspaceParams) (*v4img.GetImagesByNameInWorkspaceOK, error) {
	resp := &model.ImagesV4Response{
		BaseImages: []*model.BaseImageV4Response{
			{
				Date:        "1111-11-11",
				Version:     "1.1.1",
				Description: "images",
				UUID:        "uuid",
			},
		},
	}
	return &v4img.GetImagesByNameInWorkspaceOK{Payload: resp}, nil
}

func TestListImagesImpl(t *testing.T) {
	var rows []utils.Row
	cloud.SetProviderType(cloud.AWS)
	listImagesImpl(new(mockGetPublicImagesClient), func(h []string, r []utils.Row) { rows = r }, int64(2), "catalog")
	if len(rows) != 1 {
		t.Fatalf("row number doesn't match 1 == %d", len(rows))
	}
	for _, r := range rows {
		expected := "1111-11-11 images 1.1.1 uuid"
		if strings.Join(r.DataAsStringArray(), " ") != expected {
			t.Errorf("row data not match %s == %s", expected, strings.Join(r.DataAsStringArray(), " "))
		}
	}
}

func TestFindBaseImageByUUID(t *testing.T) {
	baseImages := []*model.BaseImageV4Response{
		{
			Date:        "1111-11-11",
			Version:     "1.1.1",
			Description: "images",
			UUID:        "uuid",
		},
		{
			Date:        "1111-11-11",
			Version:     "1.1.1",
			Description: "images",
			UUID:        "uuid2",
		},
		{
			Date:        "1111-11-11",
			Version:     "1.1.1",
			Description: "images",
			UUID:        "uuid3",
		},
	}

	imageResponse := &model.ImagesV4Response{
		BaseImages: baseImages,
	}
	expectedUUID := "uuid3"

	foundImage := findImageByUUID(imageResponse, expectedUUID)

	if foundImage.UUID != expectedUUID {
		t.Errorf("UUID not match %s == %s", expectedUUID, foundImage.UUID)
	}
}

var detailedImage = &model.ImageV4Response{
	Date:        "1111-11-11",
	Version:     "1.1.1",
	Description: "images",
	UUID:        "uuid",
	Os:          "ubuntu",
	OsType:      "ubuntu",
	Images: map[string]map[string]string{
		"gcp": {
			"default": "22aa22-44bb44",
		},
	},
	PackageVersions: map[string]string{
		"kernel":         "3.10.0-123.el7,3.10.0-327.36.1.el7,3.10.0-862.6.3.el7",
		"python":         "2.7.5-69.el7_5",
		"salt":           "2017.7.5-1.el7",
		"salt-bootstrap": "0.13.0-2018-05-03T07:39:07",
	},
}

func TestListImageInformation(t *testing.T) {
	var rows []utils.Row

	writeImageInformation(func(h []string, r []utils.Row) { rows = r }, detailedImage)
	if len(rows) != 1 {
		t.Fatalf("row number doesn't match %d == 1", len(rows))
	}
	for _, r := range rows {
		assertEquals(t, "1111-11-11", r.DataAsStringArray()[0])
		assertEquals(t, "images", r.DataAsStringArray()[1])
		assertEquals(t, "1.1.1", r.DataAsStringArray()[2])
		assertEquals(t, "uuid", r.DataAsStringArray()[3])
		assertEquals(t, "ubuntu", r.DataAsStringArray()[4])
		assertEquals(t, "ubuntu", r.DataAsStringArray()[5])
		assertEquals(t, "gcp:\n  default: 22aa22-44bb44\n", r.DataAsStringArray()[6])

		assertContains(t, r.DataAsStringArray()[7], "kernel: 3.10.0-123.el7,3.10.0-327.36.1.el7,3.10.0-862.6.3.el7")
		assertContains(t, r.DataAsStringArray()[7], "python: 2.7.5-69.el7_5")
		assertContains(t, r.DataAsStringArray()[7], "salt: 2017.7.5-1.el7")
		assertContains(t, r.DataAsStringArray()[7], "salt-bootstrap: 0.13.0-2018-05-03T07:39:07")
	}
}

func assertEquals(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf("data not match [%s] == [%s]", expected, actual)
	}
}

func assertContains(t *testing.T, actual string, substring string) {
	if !strings.Contains(actual, substring) {
		t.Errorf("data not contains substring [%s] not in [%s]", substring, actual)
	}
}
