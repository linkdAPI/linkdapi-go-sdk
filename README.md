# LinkdAPI Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/linkdapi/linkdapi-go-sdk.svg)](https://pkg.go.dev/github.com/linkdapi/linkdapi-go-sdk)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A lightweight Go SDK for [LinkdAPI](https://linkdapi.com) — the most advanced **unofficial LinkedIn API** you'll ever find. Instead of relying on brittle scrapers or search engine hacks, **LinkdAPI** connects straight to LinkedIn's own mobile and web endpoints. That means you get access to real-time data with unmatched **reliability**, **stability**, and **scalability** — perfect for developers, analysts, and anyone building tools that tap into LinkedIn at scale.

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

### Basic Usage

```go
package main

import (
    "context"
    "fmt"
    "log"

    linkdapi "github.com/linkdapi/linkdapi-go-sdk"
)

func main() {
    // Initialize the client
    client := linkdapi.NewClient("your_api_key")
    defer client.Close()

    ctx := context.Background()

    // Get profile overview
    profile, err := client.GetProfileOverview(ctx, "ryanroslansky")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Profile: %+v\n", profile)

    // Get company information
    company, err := client.GetCompanyInfo(ctx, "", "linkedin")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Company: %+v\n", company)
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

    linkdapi "github.com/linkdapi/linkdapi-go-sdk"
)

func main() {
    // Create custom configuration
    config := &linkdapi.Config{
        BaseURL:    "https://linkdapi.com",
        Timeout:    60 * time.Second,
        MaxRetries: 5,
        RetryDelay: 2 * time.Second,
    }

    // Initialize client with custom config
    client := linkdapi.NewClientWithConfig("your_api_key", config)
    defer client.Close()

    ctx := context.Background()

    profile, err := client.GetProfileOverview(ctx, "ryanroslansky")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Profile: %+v\n", profile)
}
```

---

## 📚 API Reference

All methods accept a `context.Context` as the first parameter for cancellation and timeout support.

### 🔹 Profile Endpoints

```go
// Profile Information
GetProfileOverview(ctx, username)          // Basic profile info
GetProfileDetails(ctx, urn)                // Detailed profile data
GetContactInfo(ctx, username)              // Email, phone, websites
GetProfileAbout(ctx, urn)                  // About section & verification
GetFullProfile(ctx, username, urn)         // Complete profile data in 1 request

// Work & Education
GetFullExperience(ctx, urn)                // Complete work history
GetCertifications(ctx, urn)                // Professional certifications
GetEducation(ctx, urn)                     // Education history
GetSkills(ctx, urn)                        // Skills & endorsements

// Social & Engagement
GetSocialMatrix(ctx, username)             // Connections & followers count
GetRecommendations(ctx, urn)               // Given & received recommendations
GetSimilarProfiles(ctx, urn)               // Similar profile suggestions
GetProfileReactions(ctx, urn, cursor)      // All profile reactions
GetProfileInterests(ctx, urn)              // Profile interests
GetProfileServices(ctx, urn)               // Profile services
GetProfileURN(ctx, username)               // Get URN from username
```

### 🔹 Company Endpoints

```go
// Company Search & Info
CompanyNameLookup(ctx, query)                    // Search companies by name
GetCompanyInfo(ctx, companyID, name)             // Get company details
GetSimilarCompanies(ctx, companyID)              // Similar company suggestions
GetCompanyEmployeesData(ctx, companyID)          // Employee statistics
GetCompanyJobs(ctx, companyIDs, start)           // Active job listings
GetCompanyAffiliatedPages(ctx, companyID)        // Subsidiaries & affiliates
GetCompanyPosts(ctx, companyID, start)           // Company posts
GetCompanyID(ctx, universalName)                 // Get ID from universal name
GetCompanyDetailsV2(ctx, companyID)              // Extended company details
```

### 🔹 Job Endpoints

```go
// Job Search
SearchJobs(ctx, JobSearchParams{
    Keyword:         "Software Engineer",
    Location:        "San Francisco, CA",
    JobTypes:        []string{"full_time"},
    Experience:      []string{"mid_senior"},
    TimePosted:      "1week",
    WorkArrangement: []string{"remote"},
    Start:           0,
})

// Job Details
GetJobDetails(ctx, jobID)                // Detailed job information
GetSimilarJobs(ctx, jobID)               // Similar job postings
GetPeopleAlsoViewedJobs(ctx, jobID)      // Related jobs
GetJobDetailsV2(ctx, jobID)              // Extended job details (all statuses)

// Advanced Job Search V2
SearchJobsV2(ctx, JobSearchV2Params{
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
// Posts
GetFeaturedPosts(ctx, urn)                        // Featured posts
GetAllPosts(ctx, urn, cursor, start)              // All posts with pagination
GetPostInfo(ctx, urn)                             // Single post details
GetPostComments(ctx, urn, start, count, cursor)   // Post comments
GetPostLikes(ctx, urn, start)                     // Post likes/reactions
```

### 🔹 Comment Endpoints

```go
GetAllComments(ctx, urn, cursor)       // All comments by profile
GetCommentLikes(ctx, urns, start)      // Likes on specific comments
```

### 🔹 Search Endpoints

```go
// People Search
SearchPeople(ctx, PeopleSearchParams{
    Keyword:         "software engineer",
    CurrentCompany:  []string{"1337"},
    GeoURN:          []string{"103644278"},
    Title:           "founder",
    Start:           0,
})

// Company Search
SearchCompanies(ctx, CompanySearchParams{
    Keyword:     "software",
    GeoURN:      []string{"103644278"},
    CompanySize: []string{"51-200"},
    Industry:    []string{"6"},
    Start:       0,
})

// Post Search
SearchPosts(ctx, PostSearchParams{
    Keyword:     "google",
    DatePosted:  "past-week",
    ContentType: "videos",
    SortBy:      "relevance",
    Start:       10,
})

// Other Search
SearchServices(ctx, ServiceSearchParams{...})
SearchSchools(ctx, keyword, start)
```

### 🔹 Article Endpoints

```go
GetAllArticles(ctx, urn, start)         // All articles by profile
GetArticleInfo(ctx, url)                // Article details from URL
GetArticleReactions(ctx, urn, start)    // Article likes/reactions
```

### 🔹 Services Endpoints

```go
GetServiceDetails(ctx, vanityname)    // Get service by VanityName
GetSimilarServices(ctx, vanityname)   // Get similar services
```

### 🔹 Lookup Endpoints

```go
GeoNameLookup(ctx, query)          // Search locations & get geo IDs
TitleSkillsLookup(ctx, query)      // Search skills & job titles
ServicesLookup(ctx, query)         // Search service categories
```

### 🔹 System

```go
GetServiceStatus(ctx)              // Check API service status
```

> 📖 **Full documentation for all endpoints:** [linkdapi.com/docs](https://linkdapi.com/docs/intro)

---

## 💡 Examples

### Example 1: Bulk Profile Enrichment

```go
package main

import (
    "context"
    "fmt"
    "log"
    "sync"

    linkdapi "github.com/linkdapi/linkdapi-go-sdk"
)

func main() {
    client := linkdapi.NewClient("your_api_key")
    defer client.Close()

    ctx := context.Background()

    usernames := []string{"ryanroslansky", "satyanadella", "jeffweiner08"}

    var wg sync.WaitGroup
    results := make(chan map[string]interface{}, len(usernames))

    for _, username := range usernames {
        wg.Add(1)
        go func(u string) {
            defer wg.Done()
            profile, err := client.GetProfileOverview(ctx, u)
            if err != nil {
                log.Printf("Error fetching %s: %v", u, err)
                return
            }
            results <- profile
        }(username)
    }

    go func() {
        wg.Wait()
        close(results)
    }()

    for profile := range results {
        if data, ok := profile["data"].(map[string]interface{}); ok {
            fmt.Printf("Name: %v, Headline: %v\n",
                data["fullName"], data["headline"])
        }
    }
}
```

### Example 2: Company Intelligence Dashboard

```go
package main

import (
    "context"
    "fmt"
    "log"

    linkdapi "github.com/linkdapi/linkdapi-go-sdk"
)

func main() {
    client := linkdapi.NewClient("your_api_key")
    defer client.Close()

    ctx := context.Background()

    // Get company info
    company, err := client.GetCompanyInfo(ctx, "", "linkedin")
    if err != nil {
        log.Fatal(err)
    }

    companyData := company["data"].(map[string]interface{})
    companyID := companyData["id"].(string)

    // Get company details
    employees, _ := client.GetCompanyEmployeesData(ctx, companyID)
    similar, _ := client.GetSimilarCompanies(ctx, companyID)
    jobs, _ := client.GetCompanyJobs(ctx, []string{companyID}, 0)

    fmt.Printf("Company: %v\n", companyData["name"])
    fmt.Printf("Employees: %+v\n", employees)
    fmt.Printf("Similar: %+v\n", similar)
    fmt.Printf("Jobs: %+v\n", jobs)
}
```

### Example 3: Job Market Analysis

```go
package main

import (
    "context"
    "fmt"
    "log"
    "sync"

    linkdapi "github.com/linkdapi/linkdapi-go-sdk"
)

func main() {
    client := linkdapi.NewClient("your_api_key")
    defer client.Close()

    ctx := context.Background()

    locations := []string{
        "San Francisco, CA",
        "New York, NY",
        "Austin, TX",
        "Seattle, WA",
        "Boston, MA",
    }

    var wg sync.WaitGroup

    for _, location := range locations {
        wg.Add(1)
        go func(loc string) {
            defer wg.Done()

            jobs, err := client.SearchJobs(ctx, linkdapi.JobSearchParams{
                Keyword:    "Software Engineer",
                Location:   loc,
                TimePosted: "1week",
            })

            if err != nil {
                log.Printf("Error searching %s: %v", loc, err)
                return
            }

            if data, ok := jobs["data"].(map[string]interface{}); ok {
                if jobList, ok := data["jobs"].([]interface{}); ok {
                    fmt.Printf("%s: %d jobs found\n", loc, len(jobList))
                }
            }
        }(location)
    }

    wg.Wait()
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
```

### Custom Configuration

```go
config := &linkdapi.Config{
    BaseURL:    "https://linkdapi.com",
    Timeout:    60 * time.Second,
    MaxRetries: 5,
    RetryDelay: 2 * time.Second,
}

client := linkdapi.NewClientWithConfig("your_api_key", config)
```

---

## 🔧 Error Handling

```go
package main

import (
    "context"
    "fmt"
    "log"

    linkdapi "github.com/linkdapi/linkdapi-go-sdk"
)

func main() {
    client := linkdapi.NewClient("your_api_key")
    defer client.Close()

    ctx := context.Background()

    profile, err := client.GetProfileOverview(ctx, "username")
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

### Context Timeout

```go
import "time"

// Create context with timeout
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

profile, err := client.GetProfileOverview(ctx, "username")
if err != nil {
    if ctx.Err() == context.DeadlineExceeded {
        log.Println("Request timed out")
    } else {
        log.Printf("Error: %v", err)
    }
}
```

---

## 🚀 Concurrency

The Go SDK is designed to be safe for concurrent use. You can make multiple requests concurrently using goroutines:

```go
package main

import (
    "context"
    "fmt"
    "sync"

    linkdapi "github.com/linkdapi/linkdapi-go-sdk"
)

func main() {
    client := linkdapi.NewClient("your_api_key")
    defer client.Close()

    ctx := context.Background()

    usernames := []string{"user1", "user2", "user3", "user4", "user5"}

    var wg sync.WaitGroup

    for _, username := range usernames {
        wg.Add(1)
        go func(u string) {
            defer wg.Done()

            profile, err := client.GetProfileOverview(ctx, u)
            if err != nil {
                fmt.Printf("Error: %v\n", err)
                return
            }

            fmt.Printf("Fetched profile for %s\n", u)
        }(username)
    }

    wg.Wait()
    fmt.Println("All profiles fetched!")
}
```

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

⚠️ **Important:** This SDK is intended for **research and educational purposes**. Always respect LinkedIn's Terms of Service and rate limits. Use responsibly and ethically.

---

## 🌟 Support the Project

If you find LinkdAPI useful, consider:
- ⭐ **Starring the project** on GitHub
- 🐦 **Following us** on [Twitter/X](https://x.com/l1nkdapi)
- 📢 **Sharing** with your network
- 💡 **Contributing** ideas and feedback

---

<div align="center">

**Built with ❤️ for developers who need reliable LinkedIn data**

[Website](https://linkdapi.com) • [Documentation](https://linkdapi.com/docs) • [Twitter](https://x.com/l1nkdapi) • [Support](https://linkdapi.com/help-center)

</div>
