/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by counterfeiter. DO NOT EDIT.
package anagofakes

import (
	"sync"

	"k8s.io/release/pkg/release"
)

type FakeStageClient struct {
	BuildStub        func([]string) error
	buildMutex       sync.RWMutex
	buildArgsForCall []struct {
		arg1 []string
	}
	buildReturns struct {
		result1 error
	}
	buildReturnsOnCall map[int]struct {
		result1 error
	}
	CheckPrerequisitesStub        func() error
	checkPrerequisitesMutex       sync.RWMutex
	checkPrerequisitesArgsForCall []struct {
	}
	checkPrerequisitesReturns struct {
		result1 error
	}
	checkPrerequisitesReturnsOnCall map[int]struct {
		result1 error
	}
	GenerateReleaseNotesStub        func() error
	generateReleaseNotesMutex       sync.RWMutex
	generateReleaseNotesArgsForCall []struct {
	}
	generateReleaseNotesReturns struct {
		result1 error
	}
	generateReleaseNotesReturnsOnCall map[int]struct {
		result1 error
	}
	GenerateReleaseVersionStub        func(string) (*release.Versions, error)
	generateReleaseVersionMutex       sync.RWMutex
	generateReleaseVersionArgsForCall []struct {
		arg1 string
	}
	generateReleaseVersionReturns struct {
		result1 *release.Versions
		result2 error
	}
	generateReleaseVersionReturnsOnCall map[int]struct {
		result1 *release.Versions
		result2 error
	}
	PrepareWorkspaceStub        func() error
	prepareWorkspaceMutex       sync.RWMutex
	prepareWorkspaceArgsForCall []struct {
	}
	prepareWorkspaceReturns struct {
		result1 error
	}
	prepareWorkspaceReturnsOnCall map[int]struct {
		result1 error
	}
	SetBuildCandidateStub        func() error
	setBuildCandidateMutex       sync.RWMutex
	setBuildCandidateArgsForCall []struct {
	}
	setBuildCandidateReturns struct {
		result1 error
	}
	setBuildCandidateReturnsOnCall map[int]struct {
		result1 error
	}
	StageArtifactsStub        func([]string) error
	stageArtifactsMutex       sync.RWMutex
	stageArtifactsArgsForCall []struct {
		arg1 []string
	}
	stageArtifactsReturns struct {
		result1 error
	}
	stageArtifactsReturnsOnCall map[int]struct {
		result1 error
	}
	TagRepositoryStub        func([]string) error
	tagRepositoryMutex       sync.RWMutex
	tagRepositoryArgsForCall []struct {
		arg1 []string
	}
	tagRepositoryReturns struct {
		result1 error
	}
	tagRepositoryReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateOptionsStub        func() error
	validateOptionsMutex       sync.RWMutex
	validateOptionsArgsForCall []struct {
	}
	validateOptionsReturns struct {
		result1 error
	}
	validateOptionsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStageClient) Build(arg1 []string) error {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.buildMutex.Lock()
	ret, specificReturn := fake.buildReturnsOnCall[len(fake.buildArgsForCall)]
	fake.buildArgsForCall = append(fake.buildArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	stub := fake.BuildStub
	fakeReturns := fake.buildReturns
	fake.recordInvocation("Build", []interface{}{arg1Copy})
	fake.buildMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStageClient) BuildCallCount() int {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return len(fake.buildArgsForCall)
}

func (fake *FakeStageClient) BuildCalls(stub func([]string) error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = stub
}

func (fake *FakeStageClient) BuildArgsForCall(i int) []string {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	argsForCall := fake.buildArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStageClient) BuildReturns(result1 error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = nil
	fake.buildReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) BuildReturnsOnCall(i int, result1 error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = nil
	if fake.buildReturnsOnCall == nil {
		fake.buildReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.buildReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) CheckPrerequisites() error {
	fake.checkPrerequisitesMutex.Lock()
	ret, specificReturn := fake.checkPrerequisitesReturnsOnCall[len(fake.checkPrerequisitesArgsForCall)]
	fake.checkPrerequisitesArgsForCall = append(fake.checkPrerequisitesArgsForCall, struct {
	}{})
	stub := fake.CheckPrerequisitesStub
	fakeReturns := fake.checkPrerequisitesReturns
	fake.recordInvocation("CheckPrerequisites", []interface{}{})
	fake.checkPrerequisitesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStageClient) CheckPrerequisitesCallCount() int {
	fake.checkPrerequisitesMutex.RLock()
	defer fake.checkPrerequisitesMutex.RUnlock()
	return len(fake.checkPrerequisitesArgsForCall)
}

func (fake *FakeStageClient) CheckPrerequisitesCalls(stub func() error) {
	fake.checkPrerequisitesMutex.Lock()
	defer fake.checkPrerequisitesMutex.Unlock()
	fake.CheckPrerequisitesStub = stub
}

func (fake *FakeStageClient) CheckPrerequisitesReturns(result1 error) {
	fake.checkPrerequisitesMutex.Lock()
	defer fake.checkPrerequisitesMutex.Unlock()
	fake.CheckPrerequisitesStub = nil
	fake.checkPrerequisitesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) CheckPrerequisitesReturnsOnCall(i int, result1 error) {
	fake.checkPrerequisitesMutex.Lock()
	defer fake.checkPrerequisitesMutex.Unlock()
	fake.CheckPrerequisitesStub = nil
	if fake.checkPrerequisitesReturnsOnCall == nil {
		fake.checkPrerequisitesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkPrerequisitesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) GenerateReleaseNotes() error {
	fake.generateReleaseNotesMutex.Lock()
	ret, specificReturn := fake.generateReleaseNotesReturnsOnCall[len(fake.generateReleaseNotesArgsForCall)]
	fake.generateReleaseNotesArgsForCall = append(fake.generateReleaseNotesArgsForCall, struct {
	}{})
	stub := fake.GenerateReleaseNotesStub
	fakeReturns := fake.generateReleaseNotesReturns
	fake.recordInvocation("GenerateReleaseNotes", []interface{}{})
	fake.generateReleaseNotesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStageClient) GenerateReleaseNotesCallCount() int {
	fake.generateReleaseNotesMutex.RLock()
	defer fake.generateReleaseNotesMutex.RUnlock()
	return len(fake.generateReleaseNotesArgsForCall)
}

func (fake *FakeStageClient) GenerateReleaseNotesCalls(stub func() error) {
	fake.generateReleaseNotesMutex.Lock()
	defer fake.generateReleaseNotesMutex.Unlock()
	fake.GenerateReleaseNotesStub = stub
}

func (fake *FakeStageClient) GenerateReleaseNotesReturns(result1 error) {
	fake.generateReleaseNotesMutex.Lock()
	defer fake.generateReleaseNotesMutex.Unlock()
	fake.GenerateReleaseNotesStub = nil
	fake.generateReleaseNotesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) GenerateReleaseNotesReturnsOnCall(i int, result1 error) {
	fake.generateReleaseNotesMutex.Lock()
	defer fake.generateReleaseNotesMutex.Unlock()
	fake.GenerateReleaseNotesStub = nil
	if fake.generateReleaseNotesReturnsOnCall == nil {
		fake.generateReleaseNotesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.generateReleaseNotesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) GenerateReleaseVersion(arg1 string) (*release.Versions, error) {
	fake.generateReleaseVersionMutex.Lock()
	ret, specificReturn := fake.generateReleaseVersionReturnsOnCall[len(fake.generateReleaseVersionArgsForCall)]
	fake.generateReleaseVersionArgsForCall = append(fake.generateReleaseVersionArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GenerateReleaseVersionStub
	fakeReturns := fake.generateReleaseVersionReturns
	fake.recordInvocation("GenerateReleaseVersion", []interface{}{arg1})
	fake.generateReleaseVersionMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStageClient) GenerateReleaseVersionCallCount() int {
	fake.generateReleaseVersionMutex.RLock()
	defer fake.generateReleaseVersionMutex.RUnlock()
	return len(fake.generateReleaseVersionArgsForCall)
}

func (fake *FakeStageClient) GenerateReleaseVersionCalls(stub func(string) (*release.Versions, error)) {
	fake.generateReleaseVersionMutex.Lock()
	defer fake.generateReleaseVersionMutex.Unlock()
	fake.GenerateReleaseVersionStub = stub
}

func (fake *FakeStageClient) GenerateReleaseVersionArgsForCall(i int) string {
	fake.generateReleaseVersionMutex.RLock()
	defer fake.generateReleaseVersionMutex.RUnlock()
	argsForCall := fake.generateReleaseVersionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStageClient) GenerateReleaseVersionReturns(result1 *release.Versions, result2 error) {
	fake.generateReleaseVersionMutex.Lock()
	defer fake.generateReleaseVersionMutex.Unlock()
	fake.GenerateReleaseVersionStub = nil
	fake.generateReleaseVersionReturns = struct {
		result1 *release.Versions
		result2 error
	}{result1, result2}
}

func (fake *FakeStageClient) GenerateReleaseVersionReturnsOnCall(i int, result1 *release.Versions, result2 error) {
	fake.generateReleaseVersionMutex.Lock()
	defer fake.generateReleaseVersionMutex.Unlock()
	fake.GenerateReleaseVersionStub = nil
	if fake.generateReleaseVersionReturnsOnCall == nil {
		fake.generateReleaseVersionReturnsOnCall = make(map[int]struct {
			result1 *release.Versions
			result2 error
		})
	}
	fake.generateReleaseVersionReturnsOnCall[i] = struct {
		result1 *release.Versions
		result2 error
	}{result1, result2}
}

func (fake *FakeStageClient) PrepareWorkspace() error {
	fake.prepareWorkspaceMutex.Lock()
	ret, specificReturn := fake.prepareWorkspaceReturnsOnCall[len(fake.prepareWorkspaceArgsForCall)]
	fake.prepareWorkspaceArgsForCall = append(fake.prepareWorkspaceArgsForCall, struct {
	}{})
	stub := fake.PrepareWorkspaceStub
	fakeReturns := fake.prepareWorkspaceReturns
	fake.recordInvocation("PrepareWorkspace", []interface{}{})
	fake.prepareWorkspaceMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStageClient) PrepareWorkspaceCallCount() int {
	fake.prepareWorkspaceMutex.RLock()
	defer fake.prepareWorkspaceMutex.RUnlock()
	return len(fake.prepareWorkspaceArgsForCall)
}

func (fake *FakeStageClient) PrepareWorkspaceCalls(stub func() error) {
	fake.prepareWorkspaceMutex.Lock()
	defer fake.prepareWorkspaceMutex.Unlock()
	fake.PrepareWorkspaceStub = stub
}

func (fake *FakeStageClient) PrepareWorkspaceReturns(result1 error) {
	fake.prepareWorkspaceMutex.Lock()
	defer fake.prepareWorkspaceMutex.Unlock()
	fake.PrepareWorkspaceStub = nil
	fake.prepareWorkspaceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) PrepareWorkspaceReturnsOnCall(i int, result1 error) {
	fake.prepareWorkspaceMutex.Lock()
	defer fake.prepareWorkspaceMutex.Unlock()
	fake.PrepareWorkspaceStub = nil
	if fake.prepareWorkspaceReturnsOnCall == nil {
		fake.prepareWorkspaceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.prepareWorkspaceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) SetBuildCandidate() error {
	fake.setBuildCandidateMutex.Lock()
	ret, specificReturn := fake.setBuildCandidateReturnsOnCall[len(fake.setBuildCandidateArgsForCall)]
	fake.setBuildCandidateArgsForCall = append(fake.setBuildCandidateArgsForCall, struct {
	}{})
	stub := fake.SetBuildCandidateStub
	fakeReturns := fake.setBuildCandidateReturns
	fake.recordInvocation("SetBuildCandidate", []interface{}{})
	fake.setBuildCandidateMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStageClient) SetBuildCandidateCallCount() int {
	fake.setBuildCandidateMutex.RLock()
	defer fake.setBuildCandidateMutex.RUnlock()
	return len(fake.setBuildCandidateArgsForCall)
}

func (fake *FakeStageClient) SetBuildCandidateCalls(stub func() error) {
	fake.setBuildCandidateMutex.Lock()
	defer fake.setBuildCandidateMutex.Unlock()
	fake.SetBuildCandidateStub = stub
}

func (fake *FakeStageClient) SetBuildCandidateReturns(result1 error) {
	fake.setBuildCandidateMutex.Lock()
	defer fake.setBuildCandidateMutex.Unlock()
	fake.SetBuildCandidateStub = nil
	fake.setBuildCandidateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) SetBuildCandidateReturnsOnCall(i int, result1 error) {
	fake.setBuildCandidateMutex.Lock()
	defer fake.setBuildCandidateMutex.Unlock()
	fake.SetBuildCandidateStub = nil
	if fake.setBuildCandidateReturnsOnCall == nil {
		fake.setBuildCandidateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setBuildCandidateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) StageArtifacts(arg1 []string) error {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.stageArtifactsMutex.Lock()
	ret, specificReturn := fake.stageArtifactsReturnsOnCall[len(fake.stageArtifactsArgsForCall)]
	fake.stageArtifactsArgsForCall = append(fake.stageArtifactsArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	stub := fake.StageArtifactsStub
	fakeReturns := fake.stageArtifactsReturns
	fake.recordInvocation("StageArtifacts", []interface{}{arg1Copy})
	fake.stageArtifactsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStageClient) StageArtifactsCallCount() int {
	fake.stageArtifactsMutex.RLock()
	defer fake.stageArtifactsMutex.RUnlock()
	return len(fake.stageArtifactsArgsForCall)
}

func (fake *FakeStageClient) StageArtifactsCalls(stub func([]string) error) {
	fake.stageArtifactsMutex.Lock()
	defer fake.stageArtifactsMutex.Unlock()
	fake.StageArtifactsStub = stub
}

func (fake *FakeStageClient) StageArtifactsArgsForCall(i int) []string {
	fake.stageArtifactsMutex.RLock()
	defer fake.stageArtifactsMutex.RUnlock()
	argsForCall := fake.stageArtifactsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStageClient) StageArtifactsReturns(result1 error) {
	fake.stageArtifactsMutex.Lock()
	defer fake.stageArtifactsMutex.Unlock()
	fake.StageArtifactsStub = nil
	fake.stageArtifactsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) StageArtifactsReturnsOnCall(i int, result1 error) {
	fake.stageArtifactsMutex.Lock()
	defer fake.stageArtifactsMutex.Unlock()
	fake.StageArtifactsStub = nil
	if fake.stageArtifactsReturnsOnCall == nil {
		fake.stageArtifactsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stageArtifactsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) TagRepository(arg1 []string) error {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.tagRepositoryMutex.Lock()
	ret, specificReturn := fake.tagRepositoryReturnsOnCall[len(fake.tagRepositoryArgsForCall)]
	fake.tagRepositoryArgsForCall = append(fake.tagRepositoryArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	stub := fake.TagRepositoryStub
	fakeReturns := fake.tagRepositoryReturns
	fake.recordInvocation("TagRepository", []interface{}{arg1Copy})
	fake.tagRepositoryMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStageClient) TagRepositoryCallCount() int {
	fake.tagRepositoryMutex.RLock()
	defer fake.tagRepositoryMutex.RUnlock()
	return len(fake.tagRepositoryArgsForCall)
}

func (fake *FakeStageClient) TagRepositoryCalls(stub func([]string) error) {
	fake.tagRepositoryMutex.Lock()
	defer fake.tagRepositoryMutex.Unlock()
	fake.TagRepositoryStub = stub
}

func (fake *FakeStageClient) TagRepositoryArgsForCall(i int) []string {
	fake.tagRepositoryMutex.RLock()
	defer fake.tagRepositoryMutex.RUnlock()
	argsForCall := fake.tagRepositoryArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStageClient) TagRepositoryReturns(result1 error) {
	fake.tagRepositoryMutex.Lock()
	defer fake.tagRepositoryMutex.Unlock()
	fake.TagRepositoryStub = nil
	fake.tagRepositoryReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) TagRepositoryReturnsOnCall(i int, result1 error) {
	fake.tagRepositoryMutex.Lock()
	defer fake.tagRepositoryMutex.Unlock()
	fake.TagRepositoryStub = nil
	if fake.tagRepositoryReturnsOnCall == nil {
		fake.tagRepositoryReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.tagRepositoryReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) ValidateOptions() error {
	fake.validateOptionsMutex.Lock()
	ret, specificReturn := fake.validateOptionsReturnsOnCall[len(fake.validateOptionsArgsForCall)]
	fake.validateOptionsArgsForCall = append(fake.validateOptionsArgsForCall, struct {
	}{})
	stub := fake.ValidateOptionsStub
	fakeReturns := fake.validateOptionsReturns
	fake.recordInvocation("ValidateOptions", []interface{}{})
	fake.validateOptionsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStageClient) ValidateOptionsCallCount() int {
	fake.validateOptionsMutex.RLock()
	defer fake.validateOptionsMutex.RUnlock()
	return len(fake.validateOptionsArgsForCall)
}

func (fake *FakeStageClient) ValidateOptionsCalls(stub func() error) {
	fake.validateOptionsMutex.Lock()
	defer fake.validateOptionsMutex.Unlock()
	fake.ValidateOptionsStub = stub
}

func (fake *FakeStageClient) ValidateOptionsReturns(result1 error) {
	fake.validateOptionsMutex.Lock()
	defer fake.validateOptionsMutex.Unlock()
	fake.ValidateOptionsStub = nil
	fake.validateOptionsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) ValidateOptionsReturnsOnCall(i int, result1 error) {
	fake.validateOptionsMutex.Lock()
	defer fake.validateOptionsMutex.Unlock()
	fake.ValidateOptionsStub = nil
	if fake.validateOptionsReturnsOnCall == nil {
		fake.validateOptionsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateOptionsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStageClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	fake.checkPrerequisitesMutex.RLock()
	defer fake.checkPrerequisitesMutex.RUnlock()
	fake.generateReleaseNotesMutex.RLock()
	defer fake.generateReleaseNotesMutex.RUnlock()
	fake.generateReleaseVersionMutex.RLock()
	defer fake.generateReleaseVersionMutex.RUnlock()
	fake.prepareWorkspaceMutex.RLock()
	defer fake.prepareWorkspaceMutex.RUnlock()
	fake.setBuildCandidateMutex.RLock()
	defer fake.setBuildCandidateMutex.RUnlock()
	fake.stageArtifactsMutex.RLock()
	defer fake.stageArtifactsMutex.RUnlock()
	fake.tagRepositoryMutex.RLock()
	defer fake.tagRepositoryMutex.RUnlock()
	fake.validateOptionsMutex.RLock()
	defer fake.validateOptionsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStageClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
