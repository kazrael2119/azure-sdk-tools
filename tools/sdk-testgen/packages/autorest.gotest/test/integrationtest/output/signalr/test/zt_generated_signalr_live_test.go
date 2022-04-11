//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package test_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/internal/testutil"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/stretchr/testify/suite"
)

type SignalrTestSuite struct {
	suite.Suite

	ctx               context.Context
	cred              azcore.TokenCredential
	options           *arm.ClientOptions
	location          string
	resourceGroupName string
	subscriptionId    string
}

func (testsuite *SignalrTestSuite) SetupSuite() {
	testsuite.ctx = context.Background()
	testsuite.cred, testsuite.options = testutil.GetCredAndClientOptions(testsuite.T())
	testsuite.location = testutil.GetEnv("LOCATION", "westus")
	testsuite.resourceGroupName = testutil.GetEnv("RESOURCE_GROUP_NAME", "scenarioTestTempGroup")
	testsuite.subscriptionId = testutil.GetEnv("AZURE_SUBSCRIPTION_ID", "")

	testutil.StartRecording(testsuite.T(), "sdk/resourcemanager//test/testdata")
	resourceGroup, _, err := testutil.CreateResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.location)
	testsuite.Require().NoError(err)
	testsuite.resourceGroupName = *resourceGroup.Name
}

func (testsuite *SignalrTestSuite) TearDownSuite() {
	_, err := testutil.DeleteResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.resourceGroupName)
	testsuite.Require().NoError(err)
	testutil.StopRecording(testsuite.T())
}

func TestSignalrTestSuite(t *testing.T) {
	suite.Run(t, new(SignalrTestSuite))
}

func (testsuite *SignalrTestSuite) TestSignalr() {
	var resourceName string
	var err error
	// From step Generate_Unique_Name
	template := map[string]interface{}{
		"$schema":        "https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#",
		"contentVersion": "1.0.0.0",
		"outputs": map[string]interface{}{
			"resourceName": map[string]interface{}{
				"type":  "string",
				"value": "[variables('name').value]",
			},
		},
		"resources": []interface{}{},
		"variables": map[string]interface{}{
			"name": map[string]interface{}{
				"type": "string",
				"metadata": map[string]interface{}{
					"description": "Name of the SignalR service.",
				},
				"value": "[concat('sw',uniqueString(resourceGroup().id))]",
			},
		},
	}
	params := map[string]interface{}{}
	deployment := armresources.Deployment{
		Properties: &armresources.DeploymentProperties{
			Template:   template,
			Parameters: params,
			Mode:       to.Ptr(armresources.DeploymentModeIncremental),
		},
	}
	deploymentExtend, err := testutil.CreateDeployment(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.resourceGroupName, "Generate_Unique_Name", &deployment)
	testsuite.Require().NoError(err)
	resourceName = deploymentExtend.Properties.Outputs.(map[string]interface{})["resourceName"].(map[string]interface{})["value"].(string)

	// From step SignalR_CheckNameAvailability
	signalRClient, err := test.NewSignalRClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	_, err = signalRClient.CheckNameAvailability(testsuite.ctx,
		testsuite.location,
		test.NameAvailabilityParameters{
			Name: to.Ptr(resourceName),
			Type: to.Ptr("Microsoft.SignalRService/SignalR"),
		},
		nil)
	testsuite.Require().NoError(err)

	// From step SignalR_CreateOrUpdate
	signalRClientCreateOrUpdateResponse, err := signalRClient.BeginCreateOrUpdate(testsuite.ctx,
		testsuite.resourceGroupName,
		resourceName,
		test.ResourceInfo{
			Location: to.Ptr(testsuite.location),
			Tags: map[string]*string{
				"key1": to.Ptr("value1"),
			},
			Identity: &test.ManagedIdentity{
				Type: to.Ptr(test.ManagedIdentityTypeSystemAssigned),
			},
			Kind: to.Ptr(test.ServiceKindSignalR),
			Properties: &test.SignalRProperties{
				Cors: &test.SignalRCorsSettings{
					AllowedOrigins: []*string{
						to.Ptr("https://foo.com"),
						to.Ptr("https://bar.com")},
				},
				DisableAADAuth:   to.Ptr(false),
				DisableLocalAuth: to.Ptr(false),
				Features: []*test.SignalRFeature{
					{
						Flag:       to.Ptr(test.FeatureFlagsServiceMode),
						Properties: map[string]*string{},
						Value:      to.Ptr("Serverless"),
					},
					{
						Flag:       to.Ptr(test.FeatureFlagsEnableConnectivityLogs),
						Properties: map[string]*string{},
						Value:      to.Ptr("True"),
					},
					{
						Flag:       to.Ptr(test.FeatureFlagsEnableMessagingLogs),
						Properties: map[string]*string{},
						Value:      to.Ptr("False"),
					},
					{
						Flag:       to.Ptr(test.FeatureFlagsEnableLiveTrace),
						Properties: map[string]*string{},
						Value:      to.Ptr("False"),
					}},
				NetworkACLs: &test.SignalRNetworkACLs{
					DefaultAction: to.Ptr(test.ACLActionDeny),
					PrivateEndpoints: []*test.PrivateEndpointACL{
						{
							Allow: []*test.SignalRRequestType{
								to.Ptr(test.SignalRRequestTypeServerConnection)},
							Name: to.Ptr(resourceName + ".1fa229cd-bf3f-47f0-8c49-afb36723997e"),
						}},
					PublicNetwork: &test.NetworkACL{
						Allow: []*test.SignalRRequestType{
							to.Ptr(test.SignalRRequestTypeClientConnection)},
					},
				},
				PublicNetworkAccess: to.Ptr("Enabled"),
				TLS: &test.SignalRTLSSettings{
					ClientCertEnabled: to.Ptr(false),
				},
				Upstream: &test.ServerlessUpstreamSettings{
					Templates: []*test.UpstreamTemplate{
						{
							Auth: &test.UpstreamAuthSettings{
								Type: to.Ptr(test.UpstreamAuthTypeManagedIdentity),
								ManagedIdentity: &test.ManagedIdentitySettings{
									Resource: to.Ptr("api://example"),
								},
							},
							CategoryPattern: to.Ptr("*"),
							EventPattern:    to.Ptr("connect,disconnect"),
							HubPattern:      to.Ptr("*"),
							URLTemplate:     to.Ptr("https://example.com/chat/api/connect"),
						}},
				},
			},
			SKU: &test.ResourceSKU{
				Name:     to.Ptr("Standard_S1"),
				Capacity: to.Ptr[int32](1),
				Tier:     to.Ptr(test.SignalRSKUTierStandard),
			},
		},
		&test.SignalRClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	testsuite.Require().NoError(err)
	_, err = testutil.PullResultForTest(ctx, signalRClientCreateOrUpdateResponse)
	testsuite.Require().NoError(err)

	// From step SignalR_Get
	_, err = signalRClient.Get(testsuite.ctx,
		testsuite.resourceGroupName,
		resourceName,
		nil)
	testsuite.Require().NoError(err)

	// From step SignalR_Update
	signalRClientUpdateResponse, err := signalRClient.BeginUpdate(testsuite.ctx,
		testsuite.resourceGroupName,
		resourceName,
		test.ResourceInfo{
			Location: to.Ptr(testsuite.location),
			Tags: map[string]*string{
				"key1": to.Ptr("value1"),
			},
			Identity: &test.ManagedIdentity{
				Type: to.Ptr(test.ManagedIdentityTypeSystemAssigned),
			},
			Kind: to.Ptr(test.ServiceKindSignalR),
			Properties: &test.SignalRProperties{
				Cors: &test.SignalRCorsSettings{
					AllowedOrigins: []*string{
						to.Ptr("https://foo.com"),
						to.Ptr("https://bar.com")},
				},
				DisableAADAuth:   to.Ptr(false),
				DisableLocalAuth: to.Ptr(false),
				Features: []*test.SignalRFeature{
					{
						Flag:       to.Ptr(test.FeatureFlagsServiceMode),
						Properties: map[string]*string{},
						Value:      to.Ptr("Serverless"),
					},
					{
						Flag:       to.Ptr(test.FeatureFlagsEnableConnectivityLogs),
						Properties: map[string]*string{},
						Value:      to.Ptr("True"),
					},
					{
						Flag:       to.Ptr(test.FeatureFlagsEnableMessagingLogs),
						Properties: map[string]*string{},
						Value:      to.Ptr("False"),
					},
					{
						Flag:       to.Ptr(test.FeatureFlagsEnableLiveTrace),
						Properties: map[string]*string{},
						Value:      to.Ptr("False"),
					}},
				NetworkACLs: &test.SignalRNetworkACLs{
					DefaultAction: to.Ptr(test.ACLActionDeny),
					PrivateEndpoints: []*test.PrivateEndpointACL{
						{
							Allow: []*test.SignalRRequestType{
								to.Ptr(test.SignalRRequestTypeServerConnection)},
							Name: to.Ptr(resourceName + ".1fa229cd-bf3f-47f0-8c49-afb36723997e"),
						}},
					PublicNetwork: &test.NetworkACL{
						Allow: []*test.SignalRRequestType{
							to.Ptr(test.SignalRRequestTypeClientConnection)},
					},
				},
				PublicNetworkAccess: to.Ptr("Enabled"),
				TLS: &test.SignalRTLSSettings{
					ClientCertEnabled: to.Ptr(false),
				},
				Upstream: &test.ServerlessUpstreamSettings{
					Templates: []*test.UpstreamTemplate{
						{
							Auth: &test.UpstreamAuthSettings{
								Type: to.Ptr(test.UpstreamAuthTypeManagedIdentity),
								ManagedIdentity: &test.ManagedIdentitySettings{
									Resource: to.Ptr("api://example"),
								},
							},
							CategoryPattern: to.Ptr("*"),
							EventPattern:    to.Ptr("connect,disconnect"),
							HubPattern:      to.Ptr("*"),
							URLTemplate:     to.Ptr("https://example.com/chat/api/connect"),
						}},
				},
			},
			SKU: &test.ResourceSKU{
				Name:     to.Ptr("Standard_S1"),
				Capacity: to.Ptr[int32](1),
				Tier:     to.Ptr(test.SignalRSKUTierStandard),
			},
		},
		&test.SignalRClientBeginUpdateOptions{ResumeToken: ""})
	testsuite.Require().NoError(err)
	_, err = testutil.PullResultForTest(ctx, signalRClientUpdateResponse)
	testsuite.Require().NoError(err)

	// From step SignalR_ListKeys
	_, err = signalRClient.ListKeys(testsuite.ctx,
		testsuite.resourceGroupName,
		resourceName,
		nil)
	testsuite.Require().NoError(err)

	// From step SignalR_RegenerateKey
	signalRClientRegenerateKeyResponse, err := signalRClient.BeginRegenerateKey(testsuite.ctx,
		testsuite.resourceGroupName,
		resourceName,
		test.RegenerateKeyParameters{
			KeyType: to.Ptr(test.KeyTypePrimary),
		},
		&test.SignalRClientBeginRegenerateKeyOptions{ResumeToken: ""})
	testsuite.Require().NoError(err)
	_, err = testutil.PullResultForTest(ctx, signalRClientRegenerateKeyResponse)
	testsuite.Require().NoError(err)

	// From step SignalR_Restart
	signalRClientRestartResponse, err := signalRClient.BeginRestart(testsuite.ctx,
		testsuite.resourceGroupName,
		resourceName,
		&test.SignalRClientBeginRestartOptions{ResumeToken: ""})
	testsuite.Require().NoError(err)
	_, err = testutil.PullResultForTest(ctx, signalRClientRestartResponse)
	testsuite.Require().NoError(err)

	// From step Usages_List
	usagesClient, err := test.NewUsagesClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	usagesClientListPager := usagesClient.List(testsuite.location,
		nil)
	for usagesClientListPager.More() {
	}

	// From step SignalR_ListByResourceGroup
	signalRClientListByResourceGroupPager := signalRClient.ListByResourceGroup(testsuite.resourceGroupName,
		nil)
	for signalRClientListByResourceGroupPager.More() {
	}

	// From step SignalR_ListBySubscription
	signalRClientListBySubscriptionPager := signalRClient.ListBySubscription(nil)
	for signalRClientListBySubscriptionPager.More() {
	}

	// From step Operations_List
	operationsClient, err := test.NewOperationsClient(testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	operationsClientListPager := operationsClient.List(nil)
	for operationsClientListPager.More() {
	}

	// From step SignalR_Delete
	signalRClientDeleteResponse, err := signalRClient.BeginDelete(testsuite.ctx,
		testsuite.resourceGroupName,
		resourceName,
		&test.SignalRClientBeginDeleteOptions{ResumeToken: ""})
	testsuite.Require().NoError(err)
	_, err = testutil.PullResultForTest(ctx, signalRClientDeleteResponse)
	testsuite.Require().NoError(err)
}