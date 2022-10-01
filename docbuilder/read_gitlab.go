package docbuilder

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type GitlabData struct {
	ProjectIds []string
	Token string
	UserName string
}


type ProjDatas []GitlabProjectData

type GitlabProjectData struct {
	Id                int         `json:"id"`
	Description       string      `json:"description"`
	Name              string      `json:"name"`
	NameWithNamespace string      `json:"name_with_namespace"`
	Path              string      `json:"path"`
	PathWithNamespace string      `json:"path_with_namespace"`
	CreatedAt         time.Time   `json:"created_at"`
	DefaultBranch     string      `json:"default_branch"`
	TagList           []string    `json:"tag_list"`
	Topics            []string    `json:"topics"`
	SshUrlToRepo      string      `json:"ssh_url_to_repo"`
	HttpUrlToRepo     string      `json:"http_url_to_repo"`
	WebUrl            string      `json:"web_url"`
	ReadmeUrl         string      `json:"readme_url"`
	AvatarUrl         interface{} `json:"avatar_url"`
	ForksCount        int         `json:"forks_count"`
	StarCount         int         `json:"star_count"`
	LastActivityAt    time.Time   `json:"last_activity_at"`
	Namespace         struct {
		Id        int         `json:"id"`
		Name      string      `json:"name"`
		Path      string      `json:"path"`
		Kind      string      `json:"kind"`
		FullPath  string      `json:"full_path"`
		ParentId  int         `json:"parent_id"`
		AvatarUrl interface{} `json:"avatar_url"`
		WebUrl    string      `json:"web_url"`
	} `json:"namespace"`
	ContainerRegistryImagePrefix string `json:"container_registry_image_prefix"`
	Links                        struct {
		Self          string `json:"self"`
		Issues        string `json:"issues"`
		MergeRequests string `json:"merge_requests"`
		RepoBranches  string `json:"repo_branches"`
		Labels        string `json:"labels"`
		Events        string `json:"events"`
		Members       string `json:"members"`
	} `json:"_links"`
	PackagesEnabled                bool   `json:"packages_enabled"`
	EmptyRepo                      bool   `json:"empty_repo"`
	Archived                       bool   `json:"archived"`
	Visibility                     string `json:"visibility"`
	ResolveOutdatedDiffDiscussions bool   `json:"resolve_outdated_diff_discussions"`
	ContainerExpirationPolicy      struct {
		Cadence       string      `json:"cadence"`
		Enabled       bool        `json:"enabled"`
		KeepN         int         `json:"keep_n"`
		OlderThan     string      `json:"older_than"`
		NameRegex     string      `json:"name_regex"`
		NameRegexKeep interface{} `json:"name_regex_keep"`
		NextRunAt     time.Time   `json:"next_run_at"`
	} `json:"container_expiration_policy"`
	IssuesEnabled                             bool          `json:"issues_enabled"`
	MergeRequestsEnabled                      bool          `json:"merge_requests_enabled"`
	WikiEnabled                               bool          `json:"wiki_enabled"`
	JobsEnabled                               bool          `json:"jobs_enabled"`
	SnippetsEnabled                           bool          `json:"snippets_enabled"`
	ContainerRegistryEnabled                  bool          `json:"container_registry_enabled"`
	ServiceDeskEnabled                        bool          `json:"service_desk_enabled"`
	ServiceDeskAddress                        string        `json:"service_desk_address"`
	CanCreateMergeRequestIn                   bool          `json:"can_create_merge_request_in"`
	IssuesAccessLevel                         string        `json:"issues_access_level"`
	RepositoryAccessLevel                     string        `json:"repository_access_level"`
	MergeRequestsAccessLevel                  string        `json:"merge_requests_access_level"`
	ForkingAccessLevel                        string        `json:"forking_access_level"`
	WikiAccessLevel                           string        `json:"wiki_access_level"`
	BuildsAccessLevel                         string        `json:"builds_access_level"`
	SnippetsAccessLevel                       string        `json:"snippets_access_level"`
	PagesAccessLevel                          string        `json:"pages_access_level"`
	OperationsAccessLevel                     string        `json:"operations_access_level"`
	AnalyticsAccessLevel                      string        `json:"analytics_access_level"`
	ContainerRegistryAccessLevel              string        `json:"container_registry_access_level"`
	EmailsDisabled                            interface{}   `json:"emails_disabled"`
	SharedRunnersEnabled                      bool          `json:"shared_runners_enabled"`
	LfsEnabled                                bool          `json:"lfs_enabled"`
	CreatorId                                 int           `json:"creator_id"`
	ImportStatus                              string        `json:"import_status"`
	ImportError                               interface{}   `json:"import_error"`
	OpenIssuesCount                           int           `json:"open_issues_count"`
	RunnersToken                              string        `json:"runners_token"`
	CiDefaultGitDepth                         int           `json:"ci_default_git_depth"`
	CiForwardDeploymentEnabled                bool          `json:"ci_forward_deployment_enabled"`
	CiJobTokenScopeEnabled                    bool          `json:"ci_job_token_scope_enabled"`
	PublicJobs                                bool          `json:"public_jobs"`
	BuildGitStrategy                          string        `json:"build_git_strategy"`
	BuildTimeout                              int           `json:"build_timeout"`
	AutoCancelPendingPipelines                string        `json:"auto_cancel_pending_pipelines"`
	BuildCoverageRegex                        interface{}   `json:"build_coverage_regex"`
	CiConfigPath                              string        `json:"ci_config_path"`
	SharedWithGroups                          []interface{} `json:"shared_with_groups"`
	OnlyAllowMergeIfPipelineSucceeds          bool          `json:"only_allow_merge_if_pipeline_succeeds"`
	AllowMergeOnSkippedPipeline               interface{}   `json:"allow_merge_on_skipped_pipeline"`
	RestrictUserDefinedVariables              bool          `json:"restrict_user_defined_variables"`
	RequestAccessEnabled                      bool          `json:"request_access_enabled"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool          `json:"only_allow_merge_if_all_discussions_are_resolved"`
	RemoveSourceBranchAfterMerge              bool          `json:"remove_source_branch_after_merge"`
	PrintingMergeRequestLinkEnabled           bool          `json:"printing_merge_request_link_enabled"`
	MergeMethod                               string        `json:"merge_method"`
	SquashOption                              string        `json:"squash_option"`
	SuggestionCommitMessage                   interface{}   `json:"suggestion_commit_message"`
	MergeCommitTemplate                       interface{}   `json:"merge_commit_template"`
	SquashCommitTemplate                      interface{}   `json:"squash_commit_template"`
	AutoDevopsEnabled                         bool          `json:"auto_devops_enabled"`
	AutoDevopsDeployStrategy                  string        `json:"auto_devops_deploy_strategy"`
	AutocloseReferencedIssues                 bool          `json:"autoclose_referenced_issues"`
	KeepLatestArtifact                        bool          `json:"keep_latest_artifact"`
	RunnerTokenExpirationInterval             interface{}   `json:"runner_token_expiration_interval"`
	ApprovalsBeforeMerge                      int           `json:"approvals_before_merge"`
	Mirror                                    bool          `json:"mirror"`
	ExternalAuthorizationClassificationLabel  string        `json:"external_authorization_classification_label"`
	RequirementsEnabled                       bool          `json:"requirements_enabled"`
	SecurityAndComplianceEnabled              bool          `json:"security_and_compliance_enabled"`
	ComplianceFrameworks                      []interface{} `json:"compliance_frameworks"`
	IssuesTemplate                            interface{}   `json:"issues_template"`
	MergeRequestsTemplate                     interface{}   `json:"merge_requests_template"`
	Permissions                               struct {
		ProjectAccess interface{} `json:"project_access"`
		GroupAccess   struct {
			AccessLevel       int `json:"access_level"`
			NotificationLevel int `json:"notification_level"`
		} `json:"group_access"`
	} `json:"permissions"`
}

func GetGitLabProjectData(p string, t string)(GitlabProjectData, error){
		var s GitlabProjectData
		url := fmt.Sprintf("https://gitlab.com/api/v4/projects/%s", p)
		method := "GET"

		client := &http.Client {
		}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			return s, fmt.Errorf("%v",err.Error())
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t))

		res, err := client.Do(req)
		if err != nil {
			return s, fmt.Errorf("%v",err.Error())
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
			}
		}(res.Body)

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return s, fmt.Errorf("%v",err.Error())
		}
		err = json.Unmarshal(body, &s)
		if err != nil {
			return s, fmt.Errorf("%v",err.Error())
		}
	return s, nil
}

// GetGitlabRepos returns data of all the gitlab repos and underlying objects
func GetGitlabRepos(p ProjDatas)([]string, error){
	var repos []string
for _, v := range p {
	repos = append(repos, v.HttpUrlToRepo)
}
if len(repos) == 0 {
	return repos, fmt.Errorf("no http urls found for git repos, PAT may be expired...")
}
return repos, nil
}