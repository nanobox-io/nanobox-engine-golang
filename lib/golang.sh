# -*- mode: bash; tab-width: 2; -*-
# vim: ts=2 sw=2 ft=bash noet

# Copy the compiled binary into the app directory to run live
publish_release() {
  nos_print_bullet "Moving compiled binary into app directory..."
  cp $(nos_code_dir)/$(package_name) $(nos_app_dir)
}

# Extract the package configuration from the boxfile
#
# Ex:
# code.build:
#   engine: golang
#   config:
#     package: 'github.com/tylerflint/purple'
package() {
  echo "$(nos_payload "config_package")"
}

# The full path to the package source
package_path() {
  echo "$(nos_code_dir)/.gopath/src/$(package)"
}

# This is the name of the project and also
# becomes the name of the compiled binary
package_name() {
  echo "$(basename $(package))"
}

# Determine the golang runtime to install. This will first check
# within the Boxfile, then will rely on default_runtime to
# provide a sensible default
runtime() {
  echo $(nos_validate "$(nos_payload "config_runtime")" "string" "go-1.6")
}

# Install the golang runtime.
install_runtime() {
  nos_install "$(runtime)" 'mercurial' 'git' 'bzr'
}

# Uninstall build dependencies
uninstall_build_dependencies() {
  nos_uninstall 'go' 'mercurial' 'git' 'bzr'
}

# Allow users to specify a custom fetch command to fetch dependencies
fetch_cmd() {
  echo $(nos_validate "$(nos_payload "config_fetch")" "string" "go get")
}

# Allow users to specify a custom build command to compile the app
build_cmd() {
  echo $(nos_validate "$(nos_payload "config_build")" "string" "go build")
}

# Prepare the environment for golang builds
#
# The strategy is to create a hidden directory (.gopath) in the app directory,
# which is also a lib_dir, and use it as the GOPATH. This allows the gopath
# to be re-used between builds and also shared in 'nanobox dev'. Then we
# will symlink the app source inside of the .gopath at the package location
# which is specified in the config section of the boxfile.yml.
prep_env() {
  # Set the GOPATH environment variable to the .gopath lib_dir
  nos_persist_evar 'GOPATH' "$(nos_code_dir)/.gopath"
  # Symlink the source code into the gopath at the specified package location
  if [ ! -L $(package_path) ]; then
    # create the full path structure
    mkdir -p $(package_path)
    # remove the child-most directory to convert to a symlink
    rmdir $(package_path)
    # link the source into the package path
    ln -sf $(nos_code_dir) $(package_path)
  fi
  # Ensure a bin directory exists within the GOPATH
  mkdir -p $(nos_code_dir)/.gopath/bin
  # Add the GOPATH/bin directory to the PATH
  nos_persist_evar 'PATH' "$(nos_code_dir)/.gopath/bin:$PATH"
}

# The GOPATH was manipulated in prep_env to accomodate building the app.
# This function will remove and clean the environment setup that will not
# be required to run the compiled application.
clean_env() {
  rm -f $(nos_payload 'env_dir')/GOPATH
  # NOTE: figure out how to remove gopath from the PATH and re-persist the evar
}

# Compile the go application
compile() {
  cd $(package_path)
  # fetch dependencies
  nos_run_process "go get" "$(fetch_cmd)"
  # build
  nos_run_process "go build" "$(build_cmd)"
  cd - >/dev/null
}
