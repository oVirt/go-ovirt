//
// Copyright (c) 2020 Douglas Schilling Landgraf <dougsland@redhat.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package examples

import (
	"fmt"
	"strconv"
	"time"

	ovirtsdk4 "github.com/ovirt/go-ovirt/v4"
	"github.com/pkg/errors"
	"os"
	"strings"
)

type engineVersion struct {
	major       int // Major Version
	minor       int // Minor Version
	maintenance int // Maintenance Version
	build       int // Build Version
}

// parseRelease parse the Engine release string from argument.
// It also removes "-el7" from the release and returns the
// a engineVersion struct type. In case the error, it will return
// an error message.
func parseRelease(version string) (engineVersion, error) {
	e := engineVersion{}

	// Removes "-el7", "el-8" etc.
	if strings.Contains(version, "-") {
		version = strings.Split(version, "-")[0]
	}

	// Split the release by .
	for k, v := range strings.Split(version, ".") {
		intNumber, err := strconv.Atoi(v)
		if err != nil {
			return e, errors.New("cannot convert release to int()")
		}
		switch k {
		case 0:
			e.major = intNumber
		case 1:
			e.minor = intNumber
		case 2:
			e.maintenance = intNumber
		case 3:
			e.build = intNumber
		default:
			return e, errors.New("cannot parse Engine version")
		}
	}
	return e, nil
}

// checkReleaseSupport compare two engineVersion type, the first param is the current
// engine version and the second is the required version to run OCP on oVirt.
// Returns true in case the version is support or false and error message.
func checkReleaseSupport(current engineVersion, required engineVersion) (bool, error) {
	if current.major < required.major {
		return false, errors.New("error: engine MAJOR version is not supported")
	}
	if current.minor < required.minor {
		return false, errors.New("error: engine MINOR version is not supported")
	}
	if current.maintenance < required.maintenance {
		return false, errors.New("error: engine MAINTENANCE version is not supported")
	}
	if current.build < required.build {
		return false, errors.New("error: engine BUILD version is not supported")
	}
	return true, nil
}

func checkOCPSupport() {
	fmt.Print("Collecting inforation about engine versioning... ")
	inputRawURL := "https://my.domain.home/ovirt-engine/api"

	conn, err := ovirtsdk4.NewConnectionBuilder().
		URL(inputRawURL).
		Username("admin@internal").
		Password("mysuperpass").
		Insecure(true).
		Compress(true).
		Timeout(time.Second * 10).
		Build()
	if err != nil {
		fmt.Printf("Make connection failed, reason: %v\n", err)
		return
	}
	defer conn.Close()

	// To use `Must` methods, you should recover it if panics
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panics occurs, try the non-Must methods to find the reason")
		}
	}()

	// Get API information from the root service:
	api := conn.SystemService().Get().MustSend().MustApi()
	engineVer := api.MustProductInfo().MustVersion().MustFullVersion()
	fmt.Println("detected:", engineVer)

	currentVersion, err := parseRelease(engineVer)
	fmt.Println("\nParsed: Engine current version")
	fmt.Println("-------------------------------------")
	fmt.Println("\tMajor", currentVersion.major)
	fmt.Println("\tMinor", currentVersion.minor)
	fmt.Println("\tMinor", currentVersion.maintenance)
	fmt.Println("\tMinor", currentVersion.build)

	const engineMinimumVersionRequired = "4.3.9.4"

	requiredVersion, err := parseRelease(engineMinimumVersionRequired)
	fmt.Println("\nParsed: Engine required version:")
	fmt.Println("-------------------------------------")
	fmt.Println("\tMajor", requiredVersion.major)
	fmt.Println("\tMinor", requiredVersion.minor)
	fmt.Println("\tMinor", requiredVersion.maintenance)
	fmt.Println("\tMinor", requiredVersion.build)
	fmt.Print("\n")

	engineSupport, errorSup := checkReleaseSupport(currentVersion, requiredVersion)
	if engineSupport {
		fmt.Println("Congrats! The current Engine version can run OpenShift VMs! \\o/")
		os.Exit(0)
	}

	fmt.Println(errorSup)
	os.Exit(1)
}
