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

import "github.com/blackducksoftware/perceptor-skyfire/pkg/dump"

type PerceptorHubReport struct {
	JustPerceptorImages []string
	JustHubImages       []string
}

func NewPerceptorHubReport(dump *dump.Dump) *PerceptorHubReport {
	return &PerceptorHubReport{
		JustPerceptorImages: PerceptorNotHubImages(dump),
		JustHubImages:       HubNotPerceptorImages(dump),
	}
}

func PerceptorNotHubImages(dump *dump.Dump) []string {
	images := []string{}
	for sha := range dump.Perceptor.ImagesBySha {
		_, ok := dump.Hub.ProjectsBySha[sha]
		if !ok {
			images = append(images, sha)
		}
	}
	return images
}

func HubNotPerceptorImages(dump *dump.Dump) []string {
	images := []string{}
	for sha := range dump.Hub.ProjectsBySha {
		_, ok := dump.Perceptor.ImagesBySha[sha]
		if !ok {
			images = append(images, sha)
		}
	}
	return images
}
