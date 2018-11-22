
package snapshots

import (
	"context"
	"strings"
	"log"
	"waze/terraform/gcp_terraforming/gcp_generator"
	"waze/terraform/terraform_utils"

	"golang.org/x/oauth2/google"

	"google.golang.org/api/compute/v1"
)

var ignoreKey = map[string]bool{
	//"url":                true,
	"id":                 	true,
	"self_link":          	true,
	"fingerprint": 			true,
	"label_fingerprint": 	true,
	"creation_timestamp": 	true,
	
	"snapshot_encryption_key_sha256":			true,
	"source_disk_encryption_key_sha256":			true,
	"source_disk_link":			true,
}

var allowEmptyValues = map[string]bool{}

var additionalFields = map[string]string{
	"project": "waze-development",
}

type SnapshotsGenerator struct {
	gcp_generator.BasicGenerator
}

func (SnapshotsGenerator) createResources(SnapshotsList *compute.SnapshotsListCall, ctx context.Context, region string) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	if err := SnapshotsList.Pages(ctx, func(page *compute.SnapshotList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewTerraformResource(
				obj.Name,
				obj.Name,
				"google_compute_snapshot",
				"google",
				nil,
				map[string]string{"name": obj.Name, "project": "waze-development", "region": region},
			))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

func (g SnapshotsGenerator) Generate(zone string) error {
	region := strings.Join(strings.Split(zone, "-")[:len(strings.Split(zone, "-"))-1], "-")
	project := "waze-development" //os.Getenv("GOOGLE_CLOUD_PROJECT")
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	SnapshotsList := computeService.Snapshots.List(project)

	resources := g.createResources(SnapshotsList, ctx, region)
	err = terraform_utils.GenerateTfState(resources)
	if err != nil {
		return err
	}
	converter := terraform_utils.TfstateConverter{}
	metadata := terraform_utils.NewResourcesMetaData(resources, ignoreKey, allowEmptyValues, additionalFields)
	resources, err = converter.Convert("terraform.tfstate", metadata)
	if err != nil {
		return err
	}
	err = terraform_utils.GenerateTf(resources, "Snapshots", region, "google")
	if err != nil {
		return err
	}
	return nil

}

