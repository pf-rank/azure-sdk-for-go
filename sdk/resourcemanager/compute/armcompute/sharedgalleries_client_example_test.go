//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/41e4538ed7bb3ceac3c1322c9455a0812ed110ac/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-08-03/examples/sharedGalleryExamples/SharedGallery_List.json
func ExampleSharedGalleriesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSharedGalleriesClient().NewListPager("myLocation", &armcompute.SharedGalleriesClientListOptions{SharedTo: nil})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SharedGalleryList = armcompute.SharedGalleryList{
		// 	Value: []*armcompute.SharedGallery{
		// 		{
		// 			Name: to.Ptr("galleryUniqueName"),
		// 			Location: to.Ptr("myLocation"),
		// 			Identifier: &armcompute.SharedGalleryIdentifier{
		// 				UniqueID: to.Ptr("/SharedGalleries/galleryUniqueName"),
		// 			},
		// 			Properties: &armcompute.SharedGalleryProperties{
		// 				ArtifactTags: map[string]*string{
		// 					"ShareTag-Official1PGallery": to.Ptr("Official1PGallery"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/41e4538ed7bb3ceac3c1322c9455a0812ed110ac/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-08-03/examples/sharedGalleryExamples/SharedGallery_Get.json
func ExampleSharedGalleriesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSharedGalleriesClient().Get(ctx, "myLocation", "galleryUniqueName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SharedGallery = armcompute.SharedGallery{
	// 	Name: to.Ptr("myGalleryName"),
	// 	Location: to.Ptr("myLocation"),
	// 	Identifier: &armcompute.SharedGalleryIdentifier{
	// 		UniqueID: to.Ptr("/SharedGalleries/galleryUniqueName"),
	// 	},
	// 	Properties: &armcompute.SharedGalleryProperties{
	// 		ArtifactTags: map[string]*string{
	// 			"ShareTag-Official1PGallery": to.Ptr("Official1PGallery"),
	// 		},
	// 	},
	// }
}
