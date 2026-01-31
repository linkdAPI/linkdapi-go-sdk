# LinkdAPI Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/linkdapi/linkdapi-go-sdk.svg)](https://pkg.go.dev/github.com/linkdapi/linkdapi-go-sdk)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A lightweight Python wrapper for [LinkdAPI](https://linkdapi.com) — the most advanced API for accessing professional profile and company data. With unmatched **reliability**, **stability**, and **scalability**, it’s perfect for developers, analysts, and anyone building tools that work with professional networking data at scale.

---

## 📑 Table of Contents

- [Why LinkdAPI?](#why-linkdapi)
- [Installation](#-installation)
- [Quick Start](#-quick-start)
- [API Reference](#-api-reference)
- [Examples](#-examples)
- [Configuration](#-configuration)
- [Error Handling](#-error-handling)
- [Concurrency](#-concurrency)
- [Resources](#-resources)
- [License](#-license)

---

## Why LinkdAPI?

- We **do not rely on search engines** or SERP scraping – all data is retrieved **directly from LinkedIn.**
- Built for **scale, stability, and accuracy** using direct endpoints.
- Ideal for **automation**, **data extraction**, **reverse lookup**, and **lead generation**.

---

## 📦 Installation

```bash
go get github.com/linkdapi/linkdapi-go-sdk
```

---

## 🚀 Quick Start

### Basic Usage (Simple & Clean!)

No need to pass `context.Context` to every method! The SDK handles it internally.

```go
package main

import (
    "fmt"
    "log"

    "github.com/linkdapi/linkdapi-go-sdk/linkdapi"
)

func main() {
    // Initialize the client - simple!
    client := linkdapi.NewClient("your_api_key")
    defer client.Close()

    // Get profile overview - no ctx needed!
    profile, err := client.GetProfileOverview("ryanroslansky")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Profile: %+v\n", profile)

    // Get company information
    company, err := client.GetCompanyInfo("", "google")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Company: %+v\n", company)
}
```

### Get Full Profile (Everything in 1 Request)

```go
package main

import (
    "fmt"
    "log"

    "github.com/linkdapi/linkdapi-go-sdk/linkdapi"
)

func main() {
    client := linkdapi.NewClient("your_api_key")
    defer client.Close()

    // Get complete profile data in one request
    // You can use either username or URN
    fullProfile, err := client.GetFullProfile("ryanroslansky", "")
    if err != nil {
        log.Fatal(err)
    }

    if success, ok := fullProfile["success"].(bool); ok && success {
        if data, ok := fullProfile["data"].(map[string]interface{}); ok {
            fmt.Printf("Full Name: %v\n", data["fullName"])
            fmt.Printf("Headline: %v\n", data["headline"])
            fmt.Printf("Location: %v\n", data["location"])
            fmt.Printf("About: %v\n", data["about"])

            // Access nested data
            if experiences, ok := data["experience"].([]interface{}); ok {
                fmt.Printf("\nWork Experience (%d positions):\n", len(experiences))
                for i, exp := range experiences {
                    if expMap, ok := exp.(map[string]interface{}); ok {
                        fmt.Printf("  %d. %v at %v\n", i+1, expMap["title"], expMap["companyName"])
                    }
                }
            }

            if education, ok := data["education"].([]interface{}); ok {
                fmt.Printf("\nEducation (%d schools):\n", len(education))
            }

            if skills, ok := data["skills"].([]interface{}); ok {
                fmt.Printf("\nSkills: %d total\n", len(skills))
            }
        }
    }
}
```

### Get Company Details V2 (Extended Information)

```go
package main

import (
    "fmt"
    "log"

    "github.com/linkdapi/linkdapi-go-sdk/linkdapi"
)

func main() {
    client := linkdapi.NewClient("your_api_key")
    defer client.Close()

    // First, get the company ID from name
    company, err := client.GetCompanyInfo("", "google")
    if err != nil {
        log.Fatal(err)
    }

    var companyID string
    if success, ok := company["success"].(bool); ok && success {
        if data, ok := company["data"].(map[string]interface{}); ok {
            companyID = data["id"].(string)
        }
    }

    // Get extended company details with V2
    companyV2, err := client.GetCompanyDetailsV2(companyID)
    if err != nil {
        log.Fatal(err)
    }

    if success, ok := companyV2["success"].(bool); ok && success {
        if data, ok := companyV2["data"].(map[string]interface{}); ok {
            fmt.Printf("Company: %v\n", data["name"])
            fmt.Printf("Description: %v\n", data["description"])
            fmt.Printf("Industry: %v\n", data["industry"])
            fmt.Printf("Headquarters: %v\n", data["headquarters"])
            fmt.Printf("Website: %v\n", data["website"])
            fmt.Printf("Company Size: %v\n", data["companySize"])
            fmt.Printf("Founded: %v\n", data["foundedYear"])

            // V2 specific fields
            if peopleAlsoFollow, ok := data["peopleAlsoFollow"].([]interface{}); ok {
                fmt.Printf("\nPeople Also Follow: %d companies\n", len(peopleAlsoFollow))
            }

            if affiliatedByJobs, ok := data["affiliatedByJobs"].([]interface{}); ok {
                fmt.Printf("Affiliated Companies: %d\n", len(affiliatedByJobs))
            }
        }
    }
}
```

### With Custom Configuration

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/linkdapi/linkdapi-go-sdk/linkdapi"
)

func main() {
    // Custom timeout context (optional)
    ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
    defer cancel()

    // Create custom configuration
    config := &linkdapi.Config{
        BaseURL:    "https://linkdapi.com",
        Timeout:    60 * time.Second,
        MaxRetries: 5,
        RetryDelay: 2 * time.Second,
        Context:    ctx, // Optional: set custom context once
    }

    // Initialize client with custom config
    client := linkdapi.NewClientWithConfig("your_api_key", config)
    defer client.Close()

    // Use normally - no ctx in method calls!
    profile, err := client.GetProfileOverview("ryanroslansky")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Profile: %+v\n", profile)
}
```

---

## 📚 API Reference

All methods are simple and clean - **no `context.Context` parameter needed!**

### 🔹 Profile Endpoints

```go
// Profile Information
GetProfileOverview(username)                    // Basic profile info
GetProfileDetails(urn)                          // Detailed profile data
GetContactInfo(username)                        // Email, phone, websites
GetProfileAbout(urn)                            // About section & verification
GetFullProfile(username, urn)                   // Complete profile data in 1 request ⭐

// Work & Education
GetFullExperience(urn)                          // Complete work history
GetCertifications(urn)                          // Professional certifications
GetEducation(urn)                               // Education history
GetSkills(urn)                                  // Skills & endorsements

// Social & Engagement
GetSocialMatrix(username)                       // Connections & followers count
GetRecommendations(urn)                         // Given & received recommendations
GetSimilarProfiles(urn)                         // Similar profile suggestions
GetProfileReactions(urn, cursor)                // All profile reactions
GetProfileInterests(urn)                        // Profile interests
GetProfileServices(urn)                         // Profile services
GetProfileURN(username)                         // Get URN from username
```

### 🔹 Company Endpoints

```go
// Company Search & Info
CompanyNameLookup(query)                        // Search companies by name
GetCompanyInfo(companyID, name)                 // Get company details
GetSimilarCompanies(companyID)                  // Similar company suggestions
GetCompanyEmployeesData(companyID)              // Employee statistics
GetCompanyJobs(companyIDs, start)               // Active job listings
GetCompanyAffiliatedPages(companyID)            // Subsidiaries & affiliates
GetCompanyPosts(companyID, start)               // Company posts
GetCompanyID(universalName)                     // Get ID from universal name
GetCompanyDetailsV2(companyID)                  // Extended company details ⭐
```

### 🔹 Job Endpoints

```go
// Job Search
SearchJobs(JobSearchParams{
    Keyword:         "Software Engineer",
    Location:        "San Francisco, CA",
    JobTypes:        []string{"full_time"},
    Experience:      []string{"mid_senior"},
    TimePosted:      "1week",
    WorkArrangement: []string{"remote"},
    Start:           0,
})

// Job Details
GetJobDetails(jobID)                            // Detailed job information
GetSimilarJobs(jobID)                           // Similar job postings
GetPeopleAlsoViewedJobs(jobID)                  // Related jobs
GetJobDetailsV2(jobID)                          // Extended job details (all statuses)

// Advanced Job Search V2
SearchJobsV2(JobSearchV2Params{
    Keyword:            "Software Engineer",
    DatePosted:         "1week",
    Experience:         []string{"mid_senior"},
    WorkplaceTypes:     []string{"remote"},
    EasyApply:          &trueValue,
    Under10Applicants:  &trueValue,
    Start:              0,
})
```

### 🔹 Post Endpoints

```go
GetFeaturedPosts(urn)                           // Featured posts
GetAllPosts(urn, cursor, start)                 // All posts with pagination
GetPostInfo(urn)                                // Single post details
GetPostComments(urn, start, count, cursor)      // Post comments
GetPostLikes(urn, start)                        // Post likes/reactions
```

### 🔹 Comment Endpoints

```go
GetAllComments(urn, cursor)                     // All comments by profile
GetCommentLikes(urns, start)                    // Likes on specific comments
```

### 🔹 Search Endpoints

```go
// People Search
SearchPeople(PeopleSearchParams{
    Keyword:         "software engineer",
    CurrentCompany:  []string{"1337"},
    GeoURN:          []string{"103644278"},
    Title:           "founder",
    Start:           0,
})

// Company Search
SearchCompanies(CompanySearchParams{
    Keyword:     "software",
    GeoURN:      []string{"103644278"},
    CompanySize: []string{"51-200"},
    Industry:    []string{"6"},
    Start:       0,
})

// Post Search
SearchPosts(PostSearchParams{
    Keyword:     "google",
    DatePosted:  "past-week",
    ContentType: "videos",
    SortBy:      "relevance",
    Start:       10,
})

// Other Search
SearchServices(ServiceSearchParams{...})
SearchSchools(keyword, start)
```

### 🔹 Article Endpoints

```go
GetAllArticles(urn, start)                      // All articles by profile
GetArticleInfo(url)                             // Article details from URL
GetArticleReactions(urn, start)                 // Article likes/reactions
```

### 🔹 Services Endpoints

```go
GetServiceDetails(vanityname)                   // Get service by VanityName
GetSimilarServices(vanityname)                  // Get similar services
```

### 🔹 Lookup Endpoints

```go
GeoNameLookup(query)                            // Search locations & get geo IDs
TitleSkillsLookup(query)                        // Search skills & job titles
ServicesLookup(query)                           // Search service categories
```

> 📖 **Full documentation for all endpoints:** [linkdapi.com/docs](https://linkdapi.com/docs/intro)

---

## 💡 Examples

### Example 1: Single-Threaded Profile Enrichment

```go
package main

import (
    "fmt"
    "log"

    "github.com/linkdapi/linkdapi-go-sdk/linkdapi"
)

func main() {
    client := linkdapi.NewClient("your_api_key")
    defer client.Close()

    usernames := []string{"ryanroslansky", "satyanadella", "jeffweiner08"}

    fmt.Println("=== Single-Threaded Profile Enrichment ===\n")

    for _, username := range usernames {
        profile, err := client.GetProfileOverview(username)
        if err != nil {
            log.Printf("Error fetching %s: %v", username, err)
            continue
        }

        if success, ok := profile["success"].(bool); ok && success {
            if data, ok := profile["data"].(map[string]interface{}); ok {
                fmt.Printf("✓ %s - %v (%v)\n",
                    username,
                    data["fullName"],
                    data["headline"])
            }
        }
    }
}
```

### Example 2: Multi-Threaded (Concurrent) Profile Enrichment

```go
package main

import (
    "fmt"
    "log"
    "sync"
    "time"

    "github.com/linkdapi/linkdapi-go-sdk/linkdapi"
)

func main() {
    client := linkdapi.NewClient("your_api_key")
    defer client.Close()

    usernames := []string{
        "ryanroslansky",
        "satyanadella",
        "jeffweiner08",
        "billgates",
        "sundarpichai",
        "timcook",
        "elonmusk",
        "sherylsandberg",
    }

    fmt.Println("=== Multi-Threaded Profile Enrichment ===\n")
    start := time.Now()

    var wg sync.WaitGroup
    results := make(chan ProfileResult, len(usernames))

    // Launch goroutines for concurrent requests
    for _, username := range usernames {
        wg.Add(1)
        go func(u string) {
            defer wg.Done()

            profile, err := client.GetProfileOverview(u)
            if err != nil {
                results <- ProfileResult{Username: u, Error: err}
                return
            }

            results <- ProfileResult{Username: u, Data: profile}
        }(username)
    }

    // Wait for all goroutines to complete
    go func() {
        wg.Wait()
        close(results)
    }()

    // Process results
    successCount := 0
    errorCount := 0

    for result := range results {
        if result.Error != nil {
            errorCount++
            log.Printf("✗ Error fetching %s: %v", result.Username, result.Error)
            continue
        }

        successCount++
        if success, ok := result.Data["success"].(bool); ok && success {
            if data, ok := result.Data["data"].(map[string]interface{}); ok {
                fmt.Printf("✓ %s - %v (%v)\n",
                    result.Username,
                    data["fullName"],
                    data["headline"])
            }
        }
    }

    elapsed := time.Since(start)

    fmt.Printf("\n=== Results ===\n")
    fmt.Printf("Success: %d, Errors: %d\n", successCount, errorCount)
    fmt.Printf("Completed in %v\n", elapsed)
}

type ProfileResult struct {
    Username string
    Data     map[string]interface{}
    Error    error
}
```

### Example 3: Company Intelligence Dashboard

```go
package main

import (
    "fmt"
    "log"

    "github.com/linkdapi/linkdapi-go-sdk/linkdapi"
)

func main() {
    client := linkdapi.NewClient("your_api_key")
    defer client.Close()

    companyName := "google"

    fmt.Printf("=== Company Intelligence for %s ===\n\n", companyName)

    // Step 1: Get basic company info
    company, err := client.GetCompanyInfo("", companyName)
    if err != nil {
        log.Fatal(err)
    }

    var companyID string
    if success, ok := company["success"].(bool); ok && success {
        if data, ok := company["data"].(map[string]interface{}); ok {
            companyID = data["id"].(string)
            fmt.Printf("Company: %v\n", data["name"])
            fmt.Printf("Industry: %v\n", data["industry"])
            fmt.Printf("Company ID: %s\n\n", companyID)
        }
    }

    // Step 2: Get extended details with V2
    fmt.Println("Fetching extended company details...")
    companyV2, _ := client.GetCompanyDetailsV2(companyID)
    if data, ok := companyV2["data"].(map[string]interface{}); ok {
        fmt.Printf("Headquarters: %v\n", data["headquarters"])
        fmt.Printf("Founded: %v\n", data["foundedYear"])
        fmt.Printf("Specialties: %v\n", data["specialties"])
    }

    // Step 3: Get employee data
    fmt.Println("\nFetching employee data...")
    employees, _ := client.GetCompanyEmployeesData(companyID)

    // Step 4: Get jobs
    fmt.Println("\nFetching company jobs...")
    jobs, _ := client.GetCompanyJobs([]string{companyID}, 0)
    if data, ok := jobs["data"].(map[string]interface{}); ok {
        if jobList, ok := data["jobs"].([]interface{}); ok {
            fmt.Printf("Found %d active jobs\n", len(jobList))
        }
    }

    // Step 5: Get similar companies
    fmt.Println("\nFetching similar companies...")
    similar, _ := client.GetSimilarCompanies(companyID)
    if data, ok := similar["data"].(map[string]interface{}); ok {
        if companies, ok := data["companies"].([]interface{}); ok {
            fmt.Printf("Found %d similar companies\n", len(companies))
        }
    }

    fmt.Println("\n=== Analysis Complete ===")
}
```

### Example 4: Job Market Analysis

```go
package main

import (
    "fmt"
    "log"
    "sync"

    "github.com/linkdapi/linkdapi-go-sdk/linkdapi"
)

func main() {
    client := linkdapi.NewClient("your_api_key")
    defer client.Close()

    locations := []string{
        "San Francisco, CA",
        "New York, NY",
        "Austin, TX",
        "Seattle, WA",
        "Boston, MA",
    }

    fmt.Println("=== Job Market Analysis ===\n")

    var wg sync.WaitGroup

    for _, location := range locations {
        wg.Add(1)
        go func(loc string) {
            defer wg.Done()

            jobs, err := client.SearchJobs(linkdapi.JobSearchParams{
                Keyword:    "Software Engineer",
                Location:   loc,
                TimePosted: "1week",
                JobTypes:   []string{"full_time"},
            })

            if err != nil {
                log.Printf("Error searching %s: %v", loc, err)
                return
            }

            if data, ok := jobs["data"].(map[string]interface{}); ok {
                if jobList, ok := data["jobs"].([]interface{}); ok {
                    fmt.Printf("%-20s : %d jobs\n", loc, len(jobList))
                }
            }
        }(location)
    }

    wg.Wait()
    fmt.Println("\n=== Analysis Complete ===")
}
```

---

## 🔧 Configuration

### Default Configuration

```go
config := linkdapi.DefaultConfig()
// BaseURL:    "https://linkdapi.com"
// Timeout:    30 seconds
// MaxRetries: 3
// RetryDelay: 1 second (exponential backoff)
// Context:    nil (uses context.Background())
```

### Custom Configuration

```go
ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
defer cancel()

config := &linkdapi.Config{
    BaseURL:    "https://linkdapi.com",
    Timeout:    60 * time.Second,
    MaxRetries: 5,
    RetryDelay: 2 * time.Second,
    Context:    ctx, // Optional: custom context for all requests
}

client := linkdapi.NewClientWithConfig("your_api_key", config)
```

---

## 🔧 Error Handling

```go
package main

import (
    "fmt"
    "log"

    "github.com/linkdapi/linkdapi-go-sdk/linkdapi"
)

func main() {
    client := linkdapi.NewClient("your_api_key")
    defer client.Close()

    profile, err := client.GetProfileOverview("username")
    if err != nil {
        log.Printf("Error: %v", err)
        return
    }

    // Check success field
    if success, ok := profile["success"].(bool); ok && success {
        if data, ok := profile["data"].(map[string]interface{}); ok {
            fmt.Printf("Profile: %+v\n", data)
        }
    } else {
        if message, ok := profile["message"].(string); ok {
            fmt.Printf("API Error: %s\n", message)
        }
    }
}
```

---

## 🚀 Concurrency

The Go SDK is designed to be **safe for concurrent use**. You can make multiple requests concurrently using goroutines:

```go
package main

import (
    "fmt"
    "sync"

    "github.com/linkdapi/linkdapi-go-sdk/linkdapi"
)

func main() {
    client := linkdapi.NewClient("your_api_key")
    defer client.Close()

    usernames := []string{"user1", "user2", "user3", "user4", "user5"}

    var wg sync.WaitGroup

    for _, username := range usernames {
        wg.Add(1)
        go func(u string) {
            defer wg.Done()

            profile, err := client.GetProfileOverview(u)
            if err != nil {
                fmt.Printf("Error: %v\n", err)
                return
            }

            fmt.Printf("✓ Fetched profile for %s\n", u)
        }(username)
    }

    wg.Wait()
    fmt.Println("All profiles fetched!")
}
```

**Key Benefits:**
- Single client can handle multiple concurrent requests
- Connection pooling (100 max connections, 10 per host)
- Thread-safe design
- Automatic retry on failures
- No need to create multiple clients

---

## 🔗 Resources

### 📚 Documentation & Learning
- [🎓 Getting Started Guide](https://linkdapi.com/docs/intro)
- [📖 API Documentation](https://linkdapi.com/docs)

### 🛠️ Tools & Support
- [🔑 Get API Key](https://linkdapi.com/?p=signup)
- [💬 Help Center](https://linkdapi.com/help-center)
- [🗺️ Roadmap](https://linkdapi.com/roadmap)
- [🐦 Twitter/X](https://x.com/l1nkdapi)

---

## 📜 License

**MIT License** – Free to use for personal and commercial projects.

---

## 🌟 Support the Project

If you find LinkdAPI useful, consider:
- ⭐ **Starring the project** on GitHub
- 🐦 **Following us** on [Twitter/X](https://x.com/l1nkdapi)
- 📢 **Sharing** with your network
- 💡 **Contributing** ideas and feedback

---

<div align="center">

**Built with ❤️ for developers who need reliable access to professional data**

[Website](https://linkdapi.com) • [Documentation](https://linkdapi.com/docs) • [Twitter](https://x.com/l1nkdapi) • [Support](https://linkdapi.com/help-center)

</div>
