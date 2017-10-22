package artifactorymc

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.resourceProvider
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ARTIFACTORYMC_USERNAME", nil),
				Description: "Username for authentication",
			},

			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("ARTIFACTORYMC_PASSWORD", nil),
				Description: "Password or API Key to use",
			},

			"url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ARTIFACTORYMC_URL", nil),
				Description: "The URL to your Artifactory instance ",
			},
		},
		ConfigureFunc: providerConfigure,
		ResourcesMap:  map[string]*schema.Resource{},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	user := d.Get("username").(string)
	pass := d.Get("password").(string)
	url := d.Get("url").(string)
	client := NewClient(user, pass, url)

	// fail early. validate the connection to Artifactory
	if err := client.Ping(); err != nil {
		return nil, fmt.Errorf("Error connecting to Artifactory: %s", err)
	}
	return client, nil
}
