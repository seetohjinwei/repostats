package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"github.com/seetohjinwei/repostats/models"
)

// TODO: cache results in database

const GITHUB_API_VERSION = "2022-11-28"

func getClient(url string) (http.Client, *http.Request, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return client, req, err
	}

	req.Header = http.Header{
		"Accept":               {"application/vnd.github+json"},
		"X-GitHub-Api-Version": {GITHUB_API_VERSION},
	}

	return client, req, nil
}

// owner/organisation, repo, branch
const GITHUB_REPO_FILES = "https://api.github.com/repos/%s/%s/git/trees/%s?recursive=1"

// Works for the object type returned by the REST API.
func toFile(tf map[string]interface{}) models.File {
	path := tf["path"].(string)

	extension := filepath.Ext(path)
	if len(extension) > 0 {
		// trim the "."
		extension = extension[1:]
	}
	typeData := models.TypeData{
		Type:      models.GetFileType(extension),
		FileCount: 1,
		Bytes:     int64(tf["size"].(float64)),
	}

	return models.File{
		Path:      path,
		Name:      path,
		Extension: extension,
		TypeData:  typeData,
	}
}

// Returns the Repository object for a repository at "owner/name/branch".
func GetRepositoryWithData(owner, name, branch string) (models.Repository, error) {
	repo := models.Repository{}
	repo.FileTypes = map[string]models.TypeData{}

	url := fmt.Sprintf(GITHUB_REPO_FILES, owner, name, branch)
	client, req, err := getClient(url)
	if err != nil {
		return repo, err
	}

	res, err := client.Do(req)
	if err != nil {
		return repo, err
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return repo, err
	}

	var response map[string]interface{}
	json.Unmarshal(bytes, &response)

	for _, o := range response["tree"].([]interface{}) {
		object := o.(map[string]interface{})
		if object["size"] == nil {
			continue
		}

		file := toFile(object)
		repo.Files = append(repo.Files, file)

		_, ok := repo.FileTypes[file.TypeData.Type]
		if !ok {
			repo.FileTypes[file.TypeData.Type] = models.TypeData{
				Type:      file.TypeData.Type,
				FileCount: 0,
				Bytes:     0,
			}
		}

		entry := repo.FileTypes[file.TypeData.Type]
		entry.FileCount += file.TypeData.FileCount
		entry.Bytes += file.TypeData.Bytes
		repo.FileTypes[file.TypeData.Type] = entry
	}

	return repo, nil
}

// use to get default branch
// owner/organisation, repo
const GITHUB_REPO = "https://api.github.com/repos/%s/%s"

// Returns the default branch for a repository at "owner/name".
func GetDefaultBranch(owner, name string) (string, error) {
	url := fmt.Sprintf(GITHUB_REPO, owner, name)

	client, req, err := getClient(url)
	if err != nil {
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var response map[string]interface{}
	json.Unmarshal(bytes, &response)

	branch := response["default_branch"]
	if branch == nil {
		return "", errors.New("no such repo")
	}

	return branch.(string), nil
}

const GITHUB_USERNAME_REPOS = "https://api.github.com/users/%s/repos"

// Returns a list of repositories, with only the Name and DefaultBranch fields populated.
func getRepos(username string) ([]models.Repository, error) {
	var result []models.Repository

	url := fmt.Sprintf(GITHUB_USERNAME_REPOS, username)

	client, req, err := getClient(url)
	if err != nil {
		return result, err
	}

	res, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return result, err
	}

	var response []map[string]interface{}
	json.Unmarshal(bytes, &response)

	if len(response) == 0 {
		return result, errors.New("invalid user or no public repos found")
	}

	for _, repo := range response {
		name := repo["name"].(string)
		branch := repo["default_branch"].(string)

		repo := models.Repository{
			Name:          name,
			DefaultBranch: branch,
		}
		result = append(result, repo)
	}

	return result, nil
}

// TODO: takes around 10 seconds because of rate-limiting (i suspect)
func GetReposWithData(username string) ([]models.Repository, error) {
	repos, err := getRepos(username)
	if err != nil {
		return repos, err
	}

	result := []models.Repository{}

	for _, r := range repos {
		repo, err := GetRepositoryWithData(username, r.Name, r.DefaultBranch)
		if err != nil {
			return result, err
		}
		result = append(result, repo)
	}

	return result, nil
}
