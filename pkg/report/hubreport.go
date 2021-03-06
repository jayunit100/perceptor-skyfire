/*
Copyright (C) 2018 Black Duck Software, Inc.

Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements. See the NOTICE file
distributed with this work for additional information
regarding copyright ownership. The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied. See the License for the
specific language governing permissions and limitations
under the License.
*/

package report

import (
	"github.com/blackducksoftware/perceptor-skyfire/pkg/dump"
)

type HubReport struct {
	ProjectsMultipleVersions           []string
	VersionsMultipleCodeLocations      []string
	CodeLocationsMultipleScanSummaries []string
}

func NewHubReport(dump *dump.Dump) *HubReport {
	return &HubReport{
		ProjectsMultipleVersions:           HubProjectsWrongNumberOfVersions(dump),
		VersionsMultipleCodeLocations:      HubVersionsWrongNumberOfCodeLocations(dump),
		CodeLocationsMultipleScanSummaries: HubCodeLocationsWrongNumberOfScans(dump),
	}
}

func HubProjectsWrongNumberOfVersions(d *dump.Dump) []string {
	projectNames := []string{}
	for _, project := range d.Hub.Projects {
		if len(project.Versions) != 1 {
			projectNames = append(projectNames, project.Name)
		}
	}
	return projectNames
}

func HubVersionsWrongNumberOfCodeLocations(d *dump.Dump) []string {
	versionNames := []string{}
	for _, project := range d.Hub.Projects {
		for _, version := range project.Versions {
			if len(version.CodeLocations) != 1 {
				versionNames = append(versionNames, version.Name)
			}
		}
	}
	return versionNames
}

func HubCodeLocationsWrongNumberOfScans(d *dump.Dump) []string {
	codeLocationNames := []string{}
	for _, project := range d.Hub.Projects {
		for _, version := range project.Versions {
			for _, codeLocation := range version.CodeLocations {
				if len(codeLocation.ScanSummaries) != 1 {
					codeLocationNames = append(codeLocationNames, codeLocation.Name)
				}
			}
		}
	}
	return codeLocationNames
}
