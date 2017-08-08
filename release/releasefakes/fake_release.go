// Code generated by counterfeiter. DO NOT EDIT.
package releasefakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/release"
	boshjob "github.com/cloudfoundry/bosh-cli/release/job"
	boshlic "github.com/cloudfoundry/bosh-cli/release/license"
	boshman "github.com/cloudfoundry/bosh-cli/release/manifest"
	boshpkg "github.com/cloudfoundry/bosh-cli/release/pkg"
)

type FakeRelease struct {
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	SetNameStub        func(string)
	setNameMutex       sync.RWMutex
	setNameArgsForCall []struct {
		arg1 string
	}
	VersionStub        func() string
	versionMutex       sync.RWMutex
	versionArgsForCall []struct{}
	versionReturns     struct {
		result1 string
	}
	versionReturnsOnCall map[int]struct {
		result1 string
	}
	SetVersionStub        func(string)
	setVersionMutex       sync.RWMutex
	setVersionArgsForCall []struct {
		arg1 string
	}
	DescriptionStub        func() string
	descriptionMutex       sync.RWMutex
	descriptionArgsForCall []struct{}
	descriptionReturns     struct {
		result1 string
	}
	descriptionReturnsOnCall map[int]struct {
		result1 string
	}
	SetRepositoryStub        func(string)
	setRepositoryMutex       sync.RWMutex
	setRepositoryArgsForCall []struct {
		arg1 string
	}
	RepositoryStub        func() string
	repositoryMutex       sync.RWMutex
	repositoryArgsForCall []struct{}
	repositoryReturns     struct {
		result1 string
	}
	repositoryReturnsOnCall map[int]struct {
		result1 string
	}
	CommitHashWithMarkStub        func(string) string
	commitHashWithMarkMutex       sync.RWMutex
	commitHashWithMarkArgsForCall []struct {
		arg1 string
	}
	commitHashWithMarkReturns struct {
		result1 string
	}
	commitHashWithMarkReturnsOnCall map[int]struct {
		result1 string
	}
	SetCommitHashStub        func(string)
	setCommitHashMutex       sync.RWMutex
	setCommitHashArgsForCall []struct {
		arg1 string
	}
	SetUncommittedChangesStub        func(bool)
	setUncommittedChangesMutex       sync.RWMutex
	setUncommittedChangesArgsForCall []struct {
		arg1 bool
	}
	JobsStub        func() []*boshjob.Job
	jobsMutex       sync.RWMutex
	jobsArgsForCall []struct{}
	jobsReturns     struct {
		result1 []*boshjob.Job
	}
	jobsReturnsOnCall map[int]struct {
		result1 []*boshjob.Job
	}
	PackagesStub        func() []*boshpkg.Package
	packagesMutex       sync.RWMutex
	packagesArgsForCall []struct{}
	packagesReturns     struct {
		result1 []*boshpkg.Package
	}
	packagesReturnsOnCall map[int]struct {
		result1 []*boshpkg.Package
	}
	CompiledPackagesStub        func() []*boshpkg.CompiledPackage
	compiledPackagesMutex       sync.RWMutex
	compiledPackagesArgsForCall []struct{}
	compiledPackagesReturns     struct {
		result1 []*boshpkg.CompiledPackage
	}
	compiledPackagesReturnsOnCall map[int]struct {
		result1 []*boshpkg.CompiledPackage
	}
	LicenseStub        func() *boshlic.License
	licenseMutex       sync.RWMutex
	licenseArgsForCall []struct{}
	licenseReturns     struct {
		result1 *boshlic.License
	}
	licenseReturnsOnCall map[int]struct {
		result1 *boshlic.License
	}
	LicenseNameStub        func() string
	licenseNameMutex       sync.RWMutex
	licenseNameArgsForCall []struct{}
	licenseNameReturns     struct {
		result1 string
	}
	licenseNameReturnsOnCall map[int]struct {
		result1 string
	}
	IsCompiledStub        func() bool
	isCompiledMutex       sync.RWMutex
	isCompiledArgsForCall []struct{}
	isCompiledReturns     struct {
		result1 bool
	}
	isCompiledReturnsOnCall map[int]struct {
		result1 bool
	}
	FindJobByNameStub        func(string) (boshjob.Job, bool)
	findJobByNameMutex       sync.RWMutex
	findJobByNameArgsForCall []struct {
		arg1 string
	}
	findJobByNameReturns struct {
		result1 boshjob.Job
		result2 bool
	}
	findJobByNameReturnsOnCall map[int]struct {
		result1 boshjob.Job
		result2 bool
	}
	ManifestStub        func() boshman.Manifest
	manifestMutex       sync.RWMutex
	manifestArgsForCall []struct{}
	manifestReturns     struct {
		result1 boshman.Manifest
	}
	manifestReturnsOnCall map[int]struct {
		result1 boshman.Manifest
	}
	BuildStub        func(dev, final release.ArchiveIndicies) error
	buildMutex       sync.RWMutex
	buildArgsForCall []struct {
		dev   release.ArchiveIndicies
		final release.ArchiveIndicies
	}
	buildReturns struct {
		result1 error
	}
	buildReturnsOnCall map[int]struct {
		result1 error
	}
	FinalizeStub        func(final release.ArchiveIndicies) error
	finalizeMutex       sync.RWMutex
	finalizeArgsForCall []struct {
		final release.ArchiveIndicies
	}
	finalizeReturns struct {
		result1 error
	}
	finalizeReturnsOnCall map[int]struct {
		result1 error
	}
	CopyWithStub        func(jobs []*boshjob.Job, packages []*boshpkg.Package, lic *boshlic.License, compiledPackages []*boshpkg.CompiledPackage) release.Release
	copyWithMutex       sync.RWMutex
	copyWithArgsForCall []struct {
		jobs             []*boshjob.Job
		packages         []*boshpkg.Package
		lic              *boshlic.License
		compiledPackages []*boshpkg.CompiledPackage
	}
	copyWithReturns struct {
		result1 release.Release
	}
	copyWithReturnsOnCall map[int]struct {
		result1 release.Release
	}
	CleanUpStub        func() error
	cleanUpMutex       sync.RWMutex
	cleanUpArgsForCall []struct{}
	cleanUpReturns     struct {
		result1 error
	}
	cleanUpReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRelease) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.nameReturns.result1
}

func (fake *FakeRelease) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeRelease) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) NameReturnsOnCall(i int, result1 string) {
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) SetName(arg1 string) {
	fake.setNameMutex.Lock()
	fake.setNameArgsForCall = append(fake.setNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SetName", []interface{}{arg1})
	fake.setNameMutex.Unlock()
	if fake.SetNameStub != nil {
		fake.SetNameStub(arg1)
	}
}

func (fake *FakeRelease) SetNameCallCount() int {
	fake.setNameMutex.RLock()
	defer fake.setNameMutex.RUnlock()
	return len(fake.setNameArgsForCall)
}

func (fake *FakeRelease) SetNameArgsForCall(i int) string {
	fake.setNameMutex.RLock()
	defer fake.setNameMutex.RUnlock()
	return fake.setNameArgsForCall[i].arg1
}

func (fake *FakeRelease) Version() string {
	fake.versionMutex.Lock()
	ret, specificReturn := fake.versionReturnsOnCall[len(fake.versionArgsForCall)]
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct{}{})
	fake.recordInvocation("Version", []interface{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.versionReturns.result1
}

func (fake *FakeRelease) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeRelease) VersionReturns(result1 string) {
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) VersionReturnsOnCall(i int, result1 string) {
	fake.VersionStub = nil
	if fake.versionReturnsOnCall == nil {
		fake.versionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.versionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) SetVersion(arg1 string) {
	fake.setVersionMutex.Lock()
	fake.setVersionArgsForCall = append(fake.setVersionArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SetVersion", []interface{}{arg1})
	fake.setVersionMutex.Unlock()
	if fake.SetVersionStub != nil {
		fake.SetVersionStub(arg1)
	}
}

func (fake *FakeRelease) SetVersionCallCount() int {
	fake.setVersionMutex.RLock()
	defer fake.setVersionMutex.RUnlock()
	return len(fake.setVersionArgsForCall)
}

func (fake *FakeRelease) SetVersionArgsForCall(i int) string {
	fake.setVersionMutex.RLock()
	defer fake.setVersionMutex.RUnlock()
	return fake.setVersionArgsForCall[i].arg1
}

func (fake *FakeRelease) Description() string {
	fake.descriptionMutex.Lock()
	ret, specificReturn := fake.descriptionReturnsOnCall[len(fake.descriptionArgsForCall)]
	fake.descriptionArgsForCall = append(fake.descriptionArgsForCall, struct{}{})
	fake.recordInvocation("Description", []interface{}{})
	fake.descriptionMutex.Unlock()
	if fake.DescriptionStub != nil {
		return fake.DescriptionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.descriptionReturns.result1
}

func (fake *FakeRelease) DescriptionCallCount() int {
	fake.descriptionMutex.RLock()
	defer fake.descriptionMutex.RUnlock()
	return len(fake.descriptionArgsForCall)
}

func (fake *FakeRelease) DescriptionReturns(result1 string) {
	fake.DescriptionStub = nil
	fake.descriptionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) DescriptionReturnsOnCall(i int, result1 string) {
	fake.DescriptionStub = nil
	if fake.descriptionReturnsOnCall == nil {
		fake.descriptionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.descriptionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) SetRepository(arg1 string) {
	fake.setRepositoryMutex.Lock()
	fake.setRepositoryArgsForCall = append(fake.setRepositoryArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SetRepository", []interface{}{arg1})
	fake.setRepositoryMutex.Unlock()
	if fake.SetRepositoryStub != nil {
		fake.SetRepositoryStub(arg1)
	}
}

func (fake *FakeRelease) SetRepositoryCallCount() int {
	fake.setRepositoryMutex.RLock()
	defer fake.setRepositoryMutex.RUnlock()
	return len(fake.setRepositoryArgsForCall)
}

func (fake *FakeRelease) SetRepositoryArgsForCall(i int) string {
	fake.setRepositoryMutex.RLock()
	defer fake.setRepositoryMutex.RUnlock()
	return fake.setRepositoryArgsForCall[i].arg1
}

func (fake *FakeRelease) Repository() string {
	fake.repositoryMutex.Lock()
	ret, specificReturn := fake.repositoryReturnsOnCall[len(fake.repositoryArgsForCall)]
	fake.repositoryArgsForCall = append(fake.repositoryArgsForCall, struct{}{})
	fake.recordInvocation("Repository", []interface{}{})
	fake.repositoryMutex.Unlock()
	if fake.RepositoryStub != nil {
		return fake.RepositoryStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.repositoryReturns.result1
}

func (fake *FakeRelease) RepositoryCallCount() int {
	fake.repositoryMutex.RLock()
	defer fake.repositoryMutex.RUnlock()
	return len(fake.repositoryArgsForCall)
}

func (fake *FakeRelease) RepositoryReturns(result1 string) {
	fake.RepositoryStub = nil
	fake.repositoryReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) RepositoryReturnsOnCall(i int, result1 string) {
	fake.RepositoryStub = nil
	if fake.repositoryReturnsOnCall == nil {
		fake.repositoryReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.repositoryReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) CommitHashWithMark(arg1 string) string {
	fake.commitHashWithMarkMutex.Lock()
	ret, specificReturn := fake.commitHashWithMarkReturnsOnCall[len(fake.commitHashWithMarkArgsForCall)]
	fake.commitHashWithMarkArgsForCall = append(fake.commitHashWithMarkArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("CommitHashWithMark", []interface{}{arg1})
	fake.commitHashWithMarkMutex.Unlock()
	if fake.CommitHashWithMarkStub != nil {
		return fake.CommitHashWithMarkStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.commitHashWithMarkReturns.result1
}

func (fake *FakeRelease) CommitHashWithMarkCallCount() int {
	fake.commitHashWithMarkMutex.RLock()
	defer fake.commitHashWithMarkMutex.RUnlock()
	return len(fake.commitHashWithMarkArgsForCall)
}

func (fake *FakeRelease) CommitHashWithMarkArgsForCall(i int) string {
	fake.commitHashWithMarkMutex.RLock()
	defer fake.commitHashWithMarkMutex.RUnlock()
	return fake.commitHashWithMarkArgsForCall[i].arg1
}

func (fake *FakeRelease) CommitHashWithMarkReturns(result1 string) {
	fake.CommitHashWithMarkStub = nil
	fake.commitHashWithMarkReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) CommitHashWithMarkReturnsOnCall(i int, result1 string) {
	fake.CommitHashWithMarkStub = nil
	if fake.commitHashWithMarkReturnsOnCall == nil {
		fake.commitHashWithMarkReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.commitHashWithMarkReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) SetCommitHash(arg1 string) {
	fake.setCommitHashMutex.Lock()
	fake.setCommitHashArgsForCall = append(fake.setCommitHashArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SetCommitHash", []interface{}{arg1})
	fake.setCommitHashMutex.Unlock()
	if fake.SetCommitHashStub != nil {
		fake.SetCommitHashStub(arg1)
	}
}

func (fake *FakeRelease) SetCommitHashCallCount() int {
	fake.setCommitHashMutex.RLock()
	defer fake.setCommitHashMutex.RUnlock()
	return len(fake.setCommitHashArgsForCall)
}

func (fake *FakeRelease) SetCommitHashArgsForCall(i int) string {
	fake.setCommitHashMutex.RLock()
	defer fake.setCommitHashMutex.RUnlock()
	return fake.setCommitHashArgsForCall[i].arg1
}

func (fake *FakeRelease) SetUncommittedChanges(arg1 bool) {
	fake.setUncommittedChangesMutex.Lock()
	fake.setUncommittedChangesArgsForCall = append(fake.setUncommittedChangesArgsForCall, struct {
		arg1 bool
	}{arg1})
	fake.recordInvocation("SetUncommittedChanges", []interface{}{arg1})
	fake.setUncommittedChangesMutex.Unlock()
	if fake.SetUncommittedChangesStub != nil {
		fake.SetUncommittedChangesStub(arg1)
	}
}

func (fake *FakeRelease) SetUncommittedChangesCallCount() int {
	fake.setUncommittedChangesMutex.RLock()
	defer fake.setUncommittedChangesMutex.RUnlock()
	return len(fake.setUncommittedChangesArgsForCall)
}

func (fake *FakeRelease) SetUncommittedChangesArgsForCall(i int) bool {
	fake.setUncommittedChangesMutex.RLock()
	defer fake.setUncommittedChangesMutex.RUnlock()
	return fake.setUncommittedChangesArgsForCall[i].arg1
}

func (fake *FakeRelease) Jobs() []*boshjob.Job {
	fake.jobsMutex.Lock()
	ret, specificReturn := fake.jobsReturnsOnCall[len(fake.jobsArgsForCall)]
	fake.jobsArgsForCall = append(fake.jobsArgsForCall, struct{}{})
	fake.recordInvocation("Jobs", []interface{}{})
	fake.jobsMutex.Unlock()
	if fake.JobsStub != nil {
		return fake.JobsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.jobsReturns.result1
}

func (fake *FakeRelease) JobsCallCount() int {
	fake.jobsMutex.RLock()
	defer fake.jobsMutex.RUnlock()
	return len(fake.jobsArgsForCall)
}

func (fake *FakeRelease) JobsReturns(result1 []*boshjob.Job) {
	fake.JobsStub = nil
	fake.jobsReturns = struct {
		result1 []*boshjob.Job
	}{result1}
}

func (fake *FakeRelease) JobsReturnsOnCall(i int, result1 []*boshjob.Job) {
	fake.JobsStub = nil
	if fake.jobsReturnsOnCall == nil {
		fake.jobsReturnsOnCall = make(map[int]struct {
			result1 []*boshjob.Job
		})
	}
	fake.jobsReturnsOnCall[i] = struct {
		result1 []*boshjob.Job
	}{result1}
}

func (fake *FakeRelease) Packages() []*boshpkg.Package {
	fake.packagesMutex.Lock()
	ret, specificReturn := fake.packagesReturnsOnCall[len(fake.packagesArgsForCall)]
	fake.packagesArgsForCall = append(fake.packagesArgsForCall, struct{}{})
	fake.recordInvocation("Packages", []interface{}{})
	fake.packagesMutex.Unlock()
	if fake.PackagesStub != nil {
		return fake.PackagesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.packagesReturns.result1
}

func (fake *FakeRelease) PackagesCallCount() int {
	fake.packagesMutex.RLock()
	defer fake.packagesMutex.RUnlock()
	return len(fake.packagesArgsForCall)
}

func (fake *FakeRelease) PackagesReturns(result1 []*boshpkg.Package) {
	fake.PackagesStub = nil
	fake.packagesReturns = struct {
		result1 []*boshpkg.Package
	}{result1}
}

func (fake *FakeRelease) PackagesReturnsOnCall(i int, result1 []*boshpkg.Package) {
	fake.PackagesStub = nil
	if fake.packagesReturnsOnCall == nil {
		fake.packagesReturnsOnCall = make(map[int]struct {
			result1 []*boshpkg.Package
		})
	}
	fake.packagesReturnsOnCall[i] = struct {
		result1 []*boshpkg.Package
	}{result1}
}

func (fake *FakeRelease) CompiledPackages() []*boshpkg.CompiledPackage {
	fake.compiledPackagesMutex.Lock()
	ret, specificReturn := fake.compiledPackagesReturnsOnCall[len(fake.compiledPackagesArgsForCall)]
	fake.compiledPackagesArgsForCall = append(fake.compiledPackagesArgsForCall, struct{}{})
	fake.recordInvocation("CompiledPackages", []interface{}{})
	fake.compiledPackagesMutex.Unlock()
	if fake.CompiledPackagesStub != nil {
		return fake.CompiledPackagesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.compiledPackagesReturns.result1
}

func (fake *FakeRelease) CompiledPackagesCallCount() int {
	fake.compiledPackagesMutex.RLock()
	defer fake.compiledPackagesMutex.RUnlock()
	return len(fake.compiledPackagesArgsForCall)
}

func (fake *FakeRelease) CompiledPackagesReturns(result1 []*boshpkg.CompiledPackage) {
	fake.CompiledPackagesStub = nil
	fake.compiledPackagesReturns = struct {
		result1 []*boshpkg.CompiledPackage
	}{result1}
}

func (fake *FakeRelease) CompiledPackagesReturnsOnCall(i int, result1 []*boshpkg.CompiledPackage) {
	fake.CompiledPackagesStub = nil
	if fake.compiledPackagesReturnsOnCall == nil {
		fake.compiledPackagesReturnsOnCall = make(map[int]struct {
			result1 []*boshpkg.CompiledPackage
		})
	}
	fake.compiledPackagesReturnsOnCall[i] = struct {
		result1 []*boshpkg.CompiledPackage
	}{result1}
}

func (fake *FakeRelease) License() *boshlic.License {
	fake.licenseMutex.Lock()
	ret, specificReturn := fake.licenseReturnsOnCall[len(fake.licenseArgsForCall)]
	fake.licenseArgsForCall = append(fake.licenseArgsForCall, struct{}{})
	fake.recordInvocation("License", []interface{}{})
	fake.licenseMutex.Unlock()
	if fake.LicenseStub != nil {
		return fake.LicenseStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.licenseReturns.result1
}

func (fake *FakeRelease) LicenseCallCount() int {
	fake.licenseMutex.RLock()
	defer fake.licenseMutex.RUnlock()
	return len(fake.licenseArgsForCall)
}

func (fake *FakeRelease) LicenseReturns(result1 *boshlic.License) {
	fake.LicenseStub = nil
	fake.licenseReturns = struct {
		result1 *boshlic.License
	}{result1}
}

func (fake *FakeRelease) LicenseReturnsOnCall(i int, result1 *boshlic.License) {
	fake.LicenseStub = nil
	if fake.licenseReturnsOnCall == nil {
		fake.licenseReturnsOnCall = make(map[int]struct {
			result1 *boshlic.License
		})
	}
	fake.licenseReturnsOnCall[i] = struct {
		result1 *boshlic.License
	}{result1}
}

func (fake *FakeRelease) LicenseName() string {
	fake.licenseNameMutex.Lock()
	ret, specificReturn := fake.licenseNameReturnsOnCall[len(fake.licenseNameArgsForCall)]
	fake.licenseNameArgsForCall = append(fake.licenseNameArgsForCall, struct{}{})
	fake.recordInvocation("LicenseName", []interface{}{})
	fake.licenseNameMutex.Unlock()
	if fake.LicenseNameStub != nil {
		return fake.LicenseNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.licenseNameReturns.result1
}

func (fake *FakeRelease) LicenseNameCallCount() int {
	fake.licenseNameMutex.RLock()
	defer fake.licenseNameMutex.RUnlock()
	return len(fake.licenseNameArgsForCall)
}

func (fake *FakeRelease) LicenseNameReturns(result1 string) {
	fake.LicenseNameStub = nil
	fake.licenseNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) LicenseNameReturnsOnCall(i int, result1 string) {
	fake.LicenseNameStub = nil
	if fake.licenseNameReturnsOnCall == nil {
		fake.licenseNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.licenseNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) IsCompiled() bool {
	fake.isCompiledMutex.Lock()
	ret, specificReturn := fake.isCompiledReturnsOnCall[len(fake.isCompiledArgsForCall)]
	fake.isCompiledArgsForCall = append(fake.isCompiledArgsForCall, struct{}{})
	fake.recordInvocation("IsCompiled", []interface{}{})
	fake.isCompiledMutex.Unlock()
	if fake.IsCompiledStub != nil {
		return fake.IsCompiledStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isCompiledReturns.result1
}

func (fake *FakeRelease) IsCompiledCallCount() int {
	fake.isCompiledMutex.RLock()
	defer fake.isCompiledMutex.RUnlock()
	return len(fake.isCompiledArgsForCall)
}

func (fake *FakeRelease) IsCompiledReturns(result1 bool) {
	fake.IsCompiledStub = nil
	fake.isCompiledReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeRelease) IsCompiledReturnsOnCall(i int, result1 bool) {
	fake.IsCompiledStub = nil
	if fake.isCompiledReturnsOnCall == nil {
		fake.isCompiledReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isCompiledReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeRelease) FindJobByName(arg1 string) (boshjob.Job, bool) {
	fake.findJobByNameMutex.Lock()
	ret, specificReturn := fake.findJobByNameReturnsOnCall[len(fake.findJobByNameArgsForCall)]
	fake.findJobByNameArgsForCall = append(fake.findJobByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("FindJobByName", []interface{}{arg1})
	fake.findJobByNameMutex.Unlock()
	if fake.FindJobByNameStub != nil {
		return fake.FindJobByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findJobByNameReturns.result1, fake.findJobByNameReturns.result2
}

func (fake *FakeRelease) FindJobByNameCallCount() int {
	fake.findJobByNameMutex.RLock()
	defer fake.findJobByNameMutex.RUnlock()
	return len(fake.findJobByNameArgsForCall)
}

func (fake *FakeRelease) FindJobByNameArgsForCall(i int) string {
	fake.findJobByNameMutex.RLock()
	defer fake.findJobByNameMutex.RUnlock()
	return fake.findJobByNameArgsForCall[i].arg1
}

func (fake *FakeRelease) FindJobByNameReturns(result1 boshjob.Job, result2 bool) {
	fake.FindJobByNameStub = nil
	fake.findJobByNameReturns = struct {
		result1 boshjob.Job
		result2 bool
	}{result1, result2}
}

func (fake *FakeRelease) FindJobByNameReturnsOnCall(i int, result1 boshjob.Job, result2 bool) {
	fake.FindJobByNameStub = nil
	if fake.findJobByNameReturnsOnCall == nil {
		fake.findJobByNameReturnsOnCall = make(map[int]struct {
			result1 boshjob.Job
			result2 bool
		})
	}
	fake.findJobByNameReturnsOnCall[i] = struct {
		result1 boshjob.Job
		result2 bool
	}{result1, result2}
}

func (fake *FakeRelease) Manifest() boshman.Manifest {
	fake.manifestMutex.Lock()
	ret, specificReturn := fake.manifestReturnsOnCall[len(fake.manifestArgsForCall)]
	fake.manifestArgsForCall = append(fake.manifestArgsForCall, struct{}{})
	fake.recordInvocation("Manifest", []interface{}{})
	fake.manifestMutex.Unlock()
	if fake.ManifestStub != nil {
		return fake.ManifestStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.manifestReturns.result1
}

func (fake *FakeRelease) ManifestCallCount() int {
	fake.manifestMutex.RLock()
	defer fake.manifestMutex.RUnlock()
	return len(fake.manifestArgsForCall)
}

func (fake *FakeRelease) ManifestReturns(result1 boshman.Manifest) {
	fake.ManifestStub = nil
	fake.manifestReturns = struct {
		result1 boshman.Manifest
	}{result1}
}

func (fake *FakeRelease) ManifestReturnsOnCall(i int, result1 boshman.Manifest) {
	fake.ManifestStub = nil
	if fake.manifestReturnsOnCall == nil {
		fake.manifestReturnsOnCall = make(map[int]struct {
			result1 boshman.Manifest
		})
	}
	fake.manifestReturnsOnCall[i] = struct {
		result1 boshman.Manifest
	}{result1}
}

func (fake *FakeRelease) Build(dev release.ArchiveIndicies, final release.ArchiveIndicies) error {
	fake.buildMutex.Lock()
	ret, specificReturn := fake.buildReturnsOnCall[len(fake.buildArgsForCall)]
	fake.buildArgsForCall = append(fake.buildArgsForCall, struct {
		dev   release.ArchiveIndicies
		final release.ArchiveIndicies
	}{dev, final})
	fake.recordInvocation("Build", []interface{}{dev, final})
	fake.buildMutex.Unlock()
	if fake.BuildStub != nil {
		return fake.BuildStub(dev, final)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.buildReturns.result1
}

func (fake *FakeRelease) BuildCallCount() int {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return len(fake.buildArgsForCall)
}

func (fake *FakeRelease) BuildArgsForCall(i int) (release.ArchiveIndicies, release.ArchiveIndicies) {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return fake.buildArgsForCall[i].dev, fake.buildArgsForCall[i].final
}

func (fake *FakeRelease) BuildReturns(result1 error) {
	fake.BuildStub = nil
	fake.buildReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRelease) BuildReturnsOnCall(i int, result1 error) {
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

func (fake *FakeRelease) Finalize(final release.ArchiveIndicies) error {
	fake.finalizeMutex.Lock()
	ret, specificReturn := fake.finalizeReturnsOnCall[len(fake.finalizeArgsForCall)]
	fake.finalizeArgsForCall = append(fake.finalizeArgsForCall, struct {
		final release.ArchiveIndicies
	}{final})
	fake.recordInvocation("Finalize", []interface{}{final})
	fake.finalizeMutex.Unlock()
	if fake.FinalizeStub != nil {
		return fake.FinalizeStub(final)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.finalizeReturns.result1
}

func (fake *FakeRelease) FinalizeCallCount() int {
	fake.finalizeMutex.RLock()
	defer fake.finalizeMutex.RUnlock()
	return len(fake.finalizeArgsForCall)
}

func (fake *FakeRelease) FinalizeArgsForCall(i int) release.ArchiveIndicies {
	fake.finalizeMutex.RLock()
	defer fake.finalizeMutex.RUnlock()
	return fake.finalizeArgsForCall[i].final
}

func (fake *FakeRelease) FinalizeReturns(result1 error) {
	fake.FinalizeStub = nil
	fake.finalizeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRelease) FinalizeReturnsOnCall(i int, result1 error) {
	fake.FinalizeStub = nil
	if fake.finalizeReturnsOnCall == nil {
		fake.finalizeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.finalizeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRelease) CopyWith(jobs []*boshjob.Job, packages []*boshpkg.Package, lic *boshlic.License, compiledPackages []*boshpkg.CompiledPackage) release.Release {
	var jobsCopy []*boshjob.Job
	if jobs != nil {
		jobsCopy = make([]*boshjob.Job, len(jobs))
		copy(jobsCopy, jobs)
	}
	var packagesCopy []*boshpkg.Package
	if packages != nil {
		packagesCopy = make([]*boshpkg.Package, len(packages))
		copy(packagesCopy, packages)
	}
	var compiledPackagesCopy []*boshpkg.CompiledPackage
	if compiledPackages != nil {
		compiledPackagesCopy = make([]*boshpkg.CompiledPackage, len(compiledPackages))
		copy(compiledPackagesCopy, compiledPackages)
	}
	fake.copyWithMutex.Lock()
	ret, specificReturn := fake.copyWithReturnsOnCall[len(fake.copyWithArgsForCall)]
	fake.copyWithArgsForCall = append(fake.copyWithArgsForCall, struct {
		jobs             []*boshjob.Job
		packages         []*boshpkg.Package
		lic              *boshlic.License
		compiledPackages []*boshpkg.CompiledPackage
	}{jobsCopy, packagesCopy, lic, compiledPackagesCopy})
	fake.recordInvocation("CopyWith", []interface{}{jobsCopy, packagesCopy, lic, compiledPackagesCopy})
	fake.copyWithMutex.Unlock()
	if fake.CopyWithStub != nil {
		return fake.CopyWithStub(jobs, packages, lic, compiledPackages)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.copyWithReturns.result1
}

func (fake *FakeRelease) CopyWithCallCount() int {
	fake.copyWithMutex.RLock()
	defer fake.copyWithMutex.RUnlock()
	return len(fake.copyWithArgsForCall)
}

func (fake *FakeRelease) CopyWithArgsForCall(i int) ([]*boshjob.Job, []*boshpkg.Package, *boshlic.License, []*boshpkg.CompiledPackage) {
	fake.copyWithMutex.RLock()
	defer fake.copyWithMutex.RUnlock()
	return fake.copyWithArgsForCall[i].jobs, fake.copyWithArgsForCall[i].packages, fake.copyWithArgsForCall[i].lic, fake.copyWithArgsForCall[i].compiledPackages
}

func (fake *FakeRelease) CopyWithReturns(result1 release.Release) {
	fake.CopyWithStub = nil
	fake.copyWithReturns = struct {
		result1 release.Release
	}{result1}
}

func (fake *FakeRelease) CopyWithReturnsOnCall(i int, result1 release.Release) {
	fake.CopyWithStub = nil
	if fake.copyWithReturnsOnCall == nil {
		fake.copyWithReturnsOnCall = make(map[int]struct {
			result1 release.Release
		})
	}
	fake.copyWithReturnsOnCall[i] = struct {
		result1 release.Release
	}{result1}
}

func (fake *FakeRelease) CleanUp() error {
	fake.cleanUpMutex.Lock()
	ret, specificReturn := fake.cleanUpReturnsOnCall[len(fake.cleanUpArgsForCall)]
	fake.cleanUpArgsForCall = append(fake.cleanUpArgsForCall, struct{}{})
	fake.recordInvocation("CleanUp", []interface{}{})
	fake.cleanUpMutex.Unlock()
	if fake.CleanUpStub != nil {
		return fake.CleanUpStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanUpReturns.result1
}

func (fake *FakeRelease) CleanUpCallCount() int {
	fake.cleanUpMutex.RLock()
	defer fake.cleanUpMutex.RUnlock()
	return len(fake.cleanUpArgsForCall)
}

func (fake *FakeRelease) CleanUpReturns(result1 error) {
	fake.CleanUpStub = nil
	fake.cleanUpReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRelease) CleanUpReturnsOnCall(i int, result1 error) {
	fake.CleanUpStub = nil
	if fake.cleanUpReturnsOnCall == nil {
		fake.cleanUpReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanUpReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRelease) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.setNameMutex.RLock()
	defer fake.setNameMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	fake.setVersionMutex.RLock()
	defer fake.setVersionMutex.RUnlock()
	fake.descriptionMutex.RLock()
	defer fake.descriptionMutex.RUnlock()
	fake.setRepositoryMutex.RLock()
	defer fake.setRepositoryMutex.RUnlock()
	fake.repositoryMutex.RLock()
	defer fake.repositoryMutex.RUnlock()
	fake.commitHashWithMarkMutex.RLock()
	defer fake.commitHashWithMarkMutex.RUnlock()
	fake.setCommitHashMutex.RLock()
	defer fake.setCommitHashMutex.RUnlock()
	fake.setUncommittedChangesMutex.RLock()
	defer fake.setUncommittedChangesMutex.RUnlock()
	fake.jobsMutex.RLock()
	defer fake.jobsMutex.RUnlock()
	fake.packagesMutex.RLock()
	defer fake.packagesMutex.RUnlock()
	fake.compiledPackagesMutex.RLock()
	defer fake.compiledPackagesMutex.RUnlock()
	fake.licenseMutex.RLock()
	defer fake.licenseMutex.RUnlock()
	fake.licenseNameMutex.RLock()
	defer fake.licenseNameMutex.RUnlock()
	fake.isCompiledMutex.RLock()
	defer fake.isCompiledMutex.RUnlock()
	fake.findJobByNameMutex.RLock()
	defer fake.findJobByNameMutex.RUnlock()
	fake.manifestMutex.RLock()
	defer fake.manifestMutex.RUnlock()
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	fake.finalizeMutex.RLock()
	defer fake.finalizeMutex.RUnlock()
	fake.copyWithMutex.RLock()
	defer fake.copyWithMutex.RUnlock()
	fake.cleanUpMutex.RLock()
	defer fake.cleanUpMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRelease) recordInvocation(key string, args []interface{}) {
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

var _ release.Release = new(FakeRelease)
