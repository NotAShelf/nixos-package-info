package internal

import (
	"encoding/json"
	"errors"
	"os"
)

type License struct {
	URL      string `json:"url,omitempty"`
	FullName string `json:"fullName,omitempty"`
}

type Maintainer struct {
	Name   *string `json:"name,omitempty"`
	Github string  `json:"github,omitempty"`
	Email  *string `json:"email,omitempty"`
}

type PackageInput struct {
	FlakeDescription       string       `json:"flake_description,omitempty"`
	FlakeResolved          Resolved     `json:"flake_resolved,omitempty"`
	FlakeName              string       `json:"flake_name,omitempty"`
	Revision               string       `json:"revision,omitempty"`
	FlakeSource            Source       `json:"flake_source,omitempty"`
	Type                   string       `json:"type,omitempty"`
	PackageAttrName        string       `json:"package_attr_name,omitempty"`
	PackageAttrSet         string       `json:"package_attr_set,omitempty"`
	PackagePName           string       `json:"package_pname,omitempty"`
	PackagePVersion        string       `json:"package_pversion,omitempty"`
	PackagePlatforms       []string     `json:"package_platforms,omitempty"`
	PackageOutputs         []string     `json:"package_outputs,omitempty"`
	PackageDefaultOutput   string       `json:"package_default_output,omitempty"`
	PackagePrograms        []string     `json:"package_programs,omitempty"`
	PackageLicense         []License    `json:"package_license,omitempty"`
	PackageLicenseSet      []string     `json:"package_license_set,omitempty"`
	PackageMaintainers     []Maintainer `json:"package_maintainers,omitempty"`
	PackageMaintainersSet  []string     `json:"package_maintainers_set,omitempty"`
	PackageDescription     *string      `json:"package_description,omitempty"`
	PackageLongDescription *string      `json:"package_longDescription,omitempty"`
	PackageHydra           *string      `json:"package_hydra,omitempty"`
	PackageSystem          string       `json:"package_system,omitempty"`
	PackageHomepage        []string     `json:"package_homepage,omitempty"`
	PackagePosition        *string      `json:"package_position,omitempty"`
}

type Resolved struct {
	Type  string `json:"type,omitempty"`
	Owner string `json:"owner,omitempty"`
	Repo  string `json:"repo,omitempty"`
}

type Source struct {
	Type string `json:"type,omitempty"`
	URL  string `json:"url,omitempty"`
}

// PackageOutput is the output JSON format
// this is not necessary, but it lets us rename the fields to manipulate
// the output to our liking
type PackageOutput struct {
	PackageAttrName        string       `json:"package_name,omitempty"`
	PackageAttrSet         string       `json:"package_attr_set,omitempty"`
	PackagePName           string       `json:"package_pname,omitempty"`
	PackagePVersion        string       `json:"package_version,omitempty"`
	PackagePlatforms       []string     `json:"package_platforms,omitempty"`
	PackageOutputs         []string     `json:"package_outputs,omitempty"`
	PackageDefaultOutput   string       `json:"package_default_output,omitempty"`
	PackagePrograms        []string     `json:"package_programs,omitempty"`
	PackageLicense         []License    `json:"package_license,omitempty"`
	PackageLicenseSet      []string     `json:"package_license_set,omitempty"`
	PackageMaintainers     []Maintainer `json:"package_maintainers,omitempty"`
	PackageMaintainersSet  []string     `json:"package_maintainers_set,omitempty"`
	PackageDescription     *string      `json:"package_description,omitempty"`
	PackageLongDescription *string      `json:"package_longDescription,omitempty"`
	PackageHydra           *string      `json:"package_hydra,omitempty"`
	PackageSystem          string       `json:"package_system,omitempty"`
	PackageHomepage        []string     `json:"package_homepage,omitempty"`
	PackagePosition        *string      `json:"package_position,omitempty"`
}

func ReadFile(filename string, fullFlag bool) ([]PackageInput, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var packages []PackageInput
	if fullFlag {
		err = json.Unmarshal(file, &packages)
	} else {
		var packagesWithoutExtra []PackageInput
		err = json.Unmarshal(file, &packagesWithoutExtra)
		if err != nil {
			return nil, err
		}
		for _, pkg := range packagesWithoutExtra {
			packages = append(packages, PackageInput{
				PackageAttrName:    pkg.PackageAttrName,
				PackagePVersion:    pkg.PackagePVersion,
				PackageDescription: pkg.PackageDescription,
				PackageHomepage:    pkg.PackageHomepage,
			})
		}
	}

	if err != nil {
		return nil, err
	}

	return packages, nil
}

func OutputJSON(packages []PackageInput, fullFlag bool) ([]byte, error) {
	if len(packages) == 0 {
		return nil, errors.New("no packages to output")
	}

	var outputPackages []map[string]interface{}
	for _, pkg := range packages {
		outputPackage := map[string]interface{}{
			"package_name":        pkg.PackageAttrName,
			"package_version":     pkg.PackagePVersion,
			"package_description": pkg.PackageDescription,
		}
		if fullFlag {
			outputPackage["package_longDescription"] = pkg.PackageLongDescription
			outputPackage["package_homepage"] = pkg.PackageHomepage
		}
		outputPackages = append(outputPackages, outputPackage)
	}

	jsonData, err := json.Marshal(outputPackages)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
