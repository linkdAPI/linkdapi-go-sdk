// Package linkdapi provides a Go SDK for the LinkdAPI service.
//
// LinkdAPI is the most advanced unofficial LinkedIn API that connects
// directly to LinkedIn's endpoints for reliable, scalable data access.
//
// Basic usage:
//
//	client := linkdapi.NewClient("your_api_key")
//	defer client.Close()
//
//	profile, err := client.GetProfileOverview("ryanroslansky")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
package linkdapi

import (
    "github.com/linkdAPI/linkdapi-go-sdk/linkdapi"
)

type Client = linkdapi.Client
type Config = linkdapi.Config
type JobSearchParams = linkdapi.JobSearchParams
type JobSearchV2Params = linkdapi.JobSearchV2Params
type PeopleSearchParams = linkdapi.PeopleSearchParams
type CompanySearchParams = linkdapi.CompanySearchParams
type ServiceSearchParams = linkdapi.ServiceSearchParams
type PostSearchParams = linkdapi.PostSearchParams

var (
    NewClient = linkdapi.NewClient
    NewClientWithConfig = linkdapi.NewClientWithConfig
    DefaultConfig = linkdapi.DefaultConfig
)
