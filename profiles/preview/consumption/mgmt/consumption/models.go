// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package consumption

import original "github.com/Azure/azure-sdk-for-go/services/consumption/mgmt/2018-03-31/consumption"

type BudgetsClient = original.BudgetsClient

func NewBudgetsClient(subscriptionID string) BudgetsClient {
	return original.NewBudgetsClient(subscriptionID)
}
func NewBudgetsClientWithBaseURI(baseURI string, subscriptionID string) BudgetsClient {
	return original.NewBudgetsClientWithBaseURI(baseURI, subscriptionID)
}

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}

type CostAllocationTagsClient = original.CostAllocationTagsClient

func NewCostAllocationTagsClient(subscriptionID string) CostAllocationTagsClient {
	return original.NewCostAllocationTagsClient(subscriptionID)
}
func NewCostAllocationTagsClientWithBaseURI(baseURI string, subscriptionID string) CostAllocationTagsClient {
	return original.NewCostAllocationTagsClientWithBaseURI(baseURI, subscriptionID)
}

type MarketplacesClient = original.MarketplacesClient

func NewMarketplacesClient(subscriptionID string) MarketplacesClient {
	return original.NewMarketplacesClient(subscriptionID)
}
func NewMarketplacesClientWithBaseURI(baseURI string, subscriptionID string) MarketplacesClient {
	return original.NewMarketplacesClientWithBaseURI(baseURI, subscriptionID)
}

type CategoryType = original.CategoryType

const (
	Cost  CategoryType = original.Cost
	Usage CategoryType = original.Usage
)

func PossibleCategoryTypeValues() []CategoryType {
	return original.PossibleCategoryTypeValues()
}

type Datagrain = original.Datagrain

const (
	DailyGrain   Datagrain = original.DailyGrain
	MonthlyGrain Datagrain = original.MonthlyGrain
)

func PossibleDatagrainValues() []Datagrain {
	return original.PossibleDatagrainValues()
}

type OperatorType = original.OperatorType

const (
	EqualTo              OperatorType = original.EqualTo
	GreaterThan          OperatorType = original.GreaterThan
	GreaterThanOrEqualTo OperatorType = original.GreaterThanOrEqualTo
)

func PossibleOperatorTypeValues() []OperatorType {
	return original.PossibleOperatorTypeValues()
}

type TimeGrainType = original.TimeGrainType

const (
	Annually  TimeGrainType = original.Annually
	Monthly   TimeGrainType = original.Monthly
	Quarterly TimeGrainType = original.Quarterly
)

func PossibleTimeGrainTypeValues() []TimeGrainType {
	return original.PossibleTimeGrainTypeValues()
}

type Budget = original.Budget
type BudgetProperties = original.BudgetProperties
type BudgetsListResult = original.BudgetsListResult
type BudgetsListResultIterator = original.BudgetsListResultIterator
type BudgetsListResultPage = original.BudgetsListResultPage
type BudgetTimePeriod = original.BudgetTimePeriod
type CostAllocationTag = original.CostAllocationTag
type CostAllocationTagProperties = original.CostAllocationTagProperties
type CostAllocationTags = original.CostAllocationTags
type CurrentSpend = original.CurrentSpend
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type Filters = original.Filters
type Marketplace = original.Marketplace
type MarketplaceProperties = original.MarketplaceProperties
type MarketplacesListResult = original.MarketplacesListResult
type MarketplacesListResultIterator = original.MarketplacesListResultIterator
type MarketplacesListResultPage = original.MarketplacesListResultPage
type MeterDetails = original.MeterDetails
type Notification = original.Notification
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type PriceSheetModel = original.PriceSheetModel
type PriceSheetProperties = original.PriceSheetProperties
type PriceSheetResult = original.PriceSheetResult
type ProxyResource = original.ProxyResource
type ReservationDetails = original.ReservationDetails
type ReservationDetailsListResult = original.ReservationDetailsListResult
type ReservationDetailsListResultIterator = original.ReservationDetailsListResultIterator
type ReservationDetailsListResultPage = original.ReservationDetailsListResultPage
type ReservationDetailsProperties = original.ReservationDetailsProperties
type ReservationRecommendations = original.ReservationRecommendations
type ReservationRecommendationsListResult = original.ReservationRecommendationsListResult
type ReservationRecommendationsListResultIterator = original.ReservationRecommendationsListResultIterator
type ReservationRecommendationsListResultPage = original.ReservationRecommendationsListResultPage
type ReservationRecommendationsProperties = original.ReservationRecommendationsProperties
type ReservationSummaries = original.ReservationSummaries
type ReservationSummariesListResult = original.ReservationSummariesListResult
type ReservationSummariesListResultIterator = original.ReservationSummariesListResultIterator
type ReservationSummariesListResultPage = original.ReservationSummariesListResultPage
type ReservationSummariesProperties = original.ReservationSummariesProperties
type Resource = original.Resource
type ResourceAttributes = original.ResourceAttributes
type UsageDetail = original.UsageDetail
type UsageDetailProperties = original.UsageDetailProperties
type UsageDetailsListResult = original.UsageDetailsListResult
type UsageDetailsListResultIterator = original.UsageDetailsListResultIterator
type UsageDetailsListResultPage = original.UsageDetailsListResultPage
type OperationsClient = original.OperationsClient

func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}

type PriceSheetClient = original.PriceSheetClient

func NewPriceSheetClient(subscriptionID string) PriceSheetClient {
	return original.NewPriceSheetClient(subscriptionID)
}
func NewPriceSheetClientWithBaseURI(baseURI string, subscriptionID string) PriceSheetClient {
	return original.NewPriceSheetClientWithBaseURI(baseURI, subscriptionID)
}

type ReservationRecommendationsClient = original.ReservationRecommendationsClient

func NewReservationRecommendationsClient(subscriptionID string) ReservationRecommendationsClient {
	return original.NewReservationRecommendationsClient(subscriptionID)
}
func NewReservationRecommendationsClientWithBaseURI(baseURI string, subscriptionID string) ReservationRecommendationsClient {
	return original.NewReservationRecommendationsClientWithBaseURI(baseURI, subscriptionID)
}

type ReservationsDetailsClient = original.ReservationsDetailsClient

func NewReservationsDetailsClient(subscriptionID string) ReservationsDetailsClient {
	return original.NewReservationsDetailsClient(subscriptionID)
}
func NewReservationsDetailsClientWithBaseURI(baseURI string, subscriptionID string) ReservationsDetailsClient {
	return original.NewReservationsDetailsClientWithBaseURI(baseURI, subscriptionID)
}

type ReservationsSummariesClient = original.ReservationsSummariesClient

func NewReservationsSummariesClient(subscriptionID string) ReservationsSummariesClient {
	return original.NewReservationsSummariesClient(subscriptionID)
}
func NewReservationsSummariesClientWithBaseURI(baseURI string, subscriptionID string) ReservationsSummariesClient {
	return original.NewReservationsSummariesClientWithBaseURI(baseURI, subscriptionID)
}

type UsageDetailsClient = original.UsageDetailsClient

func NewUsageDetailsClient(subscriptionID string) UsageDetailsClient {
	return original.NewUsageDetailsClient(subscriptionID)
}
func NewUsageDetailsClientWithBaseURI(baseURI string, subscriptionID string) UsageDetailsClient {
	return original.NewUsageDetailsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
