// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package constants

import "os"

// DefaultBaseImage for non-scratch builds.
// renovate: datasource=docker versioning=docker depName=alpine
const DefaultBaseImage = "docker.io/alpine:3.21@sha256:a8560b36e8b8210634f77d9f7f9efd7ffa463e380b75e2e74aff4511df3ef88c"

// DefaultDirMode is UNIX file mode for mkdir.
const DefaultDirMode os.FileMode = 0o755

// DefaultPath is default value for PATH environment variable.
const DefaultPath = "/bin:/usr/bin:/sbin:/usr/sbin"

// PkgYaml is the filename of 'pkg.yaml'.
const PkgYaml = "pkg.yaml"

// VarsYaml is the filename of 'vars.yaml'.
const VarsYaml = "vars.yaml"

// Pkgfile is the filename of 'Pkgfile'.
const Pkgfile = "Pkgfile"

// Template extension.
const TemplateExt = ".tmpl"

// StageXBusyboxImage is the image name for busybox from stageX.
// renovate: datasource=docker versioning=docker depName=siderolabs/stagex/core-busybox
const StageXBusyboxImage = "ghcr.io/pl4nty/stagex/core-busybox:1.36.1@sha256:553659c54726133733c9df7e4cacaa98c10b38d799d71f35f4eba5de0e87e5de"
