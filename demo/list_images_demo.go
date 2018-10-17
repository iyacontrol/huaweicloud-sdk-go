package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack/ims/v2/cloudimages"
)

func main() {

	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.cn-north-1.myhuaweicloud.com/v3",
		ProjectID:        "f9b60643bb8e44349b75da40923cbcd3",
		AccessKey:        "HYO2CHUIHR5SBMLJQVXK",
		SecretKey:        "y5e0TNThIzb0TbsgWAcYFVcK4ejjBGZecCutoZbw",
		Domain:           "myhuaweicloud.com",
		Region:           "cn-north-1",
        DomainID:         "0986aafba48049a6b9457b89968eeabf",
	}


	provider, err_auth := openstack.AuthenticatedClient(opts)
	if err_auth != nil {
		fmt.Println("Failed to get the provider: ", err_auth)
		return
	}

	client, err_client := openstack.NewIMSV2(provider, gophercloud.EndpointOpts{})

	if err_client != nil {
		fmt.Println("Failed to get the NewIMSV2 client: ", err_client)
		return
	}

	listOpts := cloudimages.ListOpts{
		Isregistered: "true",
	}

	allPages, err_list := cloudimages.List(client,listOpts ).AllPages()

	if err_list != nil {
		if ue, ok := err_list.(*gophercloud.UnifiedError); ok {
			fmt.Println("Failed to list images.")
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	allImages, err_extract := cloudimages.ExtractImages(allPages)

	if err_extract != nil {
		fmt.Println("Unable to extract images: ",err_extract)
	}

	fmt.Println("Succeed to list images!")
	fmt.Println("First image ID is:",allImages[0].ID)
}


