package linkdapi

// JobSearchParams holds parameters for job search.
type JobSearchParams struct {
	Keyword         string   // Job title, skills, or keywords
	Location        string   // City, state, or region
	GeoID           string   // LinkedIn's internal geographic identifier
	CompanyIDs      []string // Specific company LinkedIn IDs
	JobTypes        []string // Employment types (full_time, part_time, contract, temporary, internship, volunteer)
	Experience      []string // Experience levels (internship, entry_level, associate, mid_senior, director)
	Regions         []string // Specific region codes
	TimePosted      string   // How recently posted (any, 24h, 1week, 1month)
	Salary          string   // Minimum salary (any, 40k, 60k, 80k, 100k, 120k)
	WorkArrangement []string // Work arrangement (onsite, remote, hybrid)
	Start           int      // Pagination start index
}

// JobSearchV2Params holds parameters for job search V2.
type JobSearchV2Params struct {
	Keyword            string   // Search keyword
	Start              int      // Pagination offset (default: 0, increment by 25)
	Count              int      // Number of results per page (default: 25, max: 50)
	SortBy             string   // Sort by "relevance" (default) or "date_posted"
	DatePosted         string   // Filter by "24h", "1week", or "1month"
	Experience         []string // Experience levels (internship, entry_level, associate, mid_senior, director, executive)
	JobTypes           []string // Employment types (full_time, part_time, contract, temporary, internship, volunteer, other)
	WorkplaceTypes     []string // Work arrangement (onsite, remote, hybrid)
	Salary             string   // Minimum annual salary (20k, 30k, 40k, 50k, 60k, 70k, 80k, 90k, 100k)
	Companies          []string // Company IDs
	Industries         []string // Industry IDs
	Locations          []string // LinkedIn's internal geographic identifiers
	Functions          []string // Job function codes (e.g., "it,sales,eng")
	Titles             []string // Job title IDs
	Benefits           []string // Benefits offered (medical_ins, dental_ins, vision_ins, 401k, pension, paid_maternity, paid_paternity, commuter, student_loan, tuition, disability_ins)
	Commitments        []string // Company values (dei, environmental, work_life, social_impact, career_growth)
	EasyApply          *bool    // Show only LinkedIn Easy Apply jobs
	VerifiedJob        *bool    // Show only verified job postings
	Under10Applicants  *bool    // Show jobs with fewer than 10 applicants
	FairChance         *bool    // Show jobs from fair chance employers
}

// PeopleSearchParams holds parameters for people search.
type PeopleSearchParams struct {
	Keyword         string   // Search keyword (e.g., "software engineer") - optional
	Start           int      // Pagination start index (default is 0)
	Count           int      // Number of results per page (default: 20, max: 50)
	CurrentCompany  []string // Current company IDs
	FirstName       string   // First name filter
	GeoURN          []string // Geographic URNs
	Industry        []string // Industry IDs
	LastName        string   // Last name filter
	ProfileLanguage string   // Profile language (e.g., "en" for English)
	PastCompany     []string // Past company IDs
	School          []string // School IDs
	ServiceCategory string   // Service category ID
	Title           string   // Job title (e.g., "founder")
}

// CompanySearchParams holds parameters for company search.
type CompanySearchParams struct {
	Keyword     string   // Search keyword (e.g., "software")
	Start       int      // Pagination start index (default is 0)
	Count       int      // Number of results per page (default: 25, max: 50)
	GeoURN      []string // Geographic URNs
	CompanySize []string // Company sizes (e.g., "1-10", "11-50", "51-200")
	HasJobs     *bool    // Filter companies with job listings
	Industry    []string // Industry IDs
}

// ServiceSearchParams holds parameters for service search.
type ServiceSearchParams struct {
	Keyword         string   // Search keyword (e.g., "software")
	Start           int      // Pagination start index (default is 0)
	Count           int      // Number of results per page (default: 25, max: 50)
	GeoURN          []string // Geographic URNs
	ProfileLanguage string   // Profile language (e.g., "en,ch")
	ServiceCategory []string // Service category IDs
}

// PostSearchParams holds parameters for post search.
type PostSearchParams struct {
	Keyword              string   // Search keyword (e.g., "google")
	Start                int      // Pagination start index (default is 10)
	AuthorCompany        string   // Company ID of the post author
	AuthorIndustry       string   // Industry ID of the post author
	AuthorJobTitle       string   // Job title of the post author (e.g., "founder")
	ContentType          string   // Content type (videos, photos, jobs, liveVideos, documents, collaborativeArticles)
	DatePosted           string   // Date filter (past-24h, past-week, past-month, past-year)
	FromMember           string   // Profile URN of the post author
	FromOrganization     []string // Company IDs
	MentionsMember       string   // Profile URN mentioned in posts
	MentionsOrganization []string // Company IDs mentioned
	SortBy               string   // Sort order (relevance, date_posted) - default is "relevance"
}
