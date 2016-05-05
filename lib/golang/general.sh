# -*- mode: bash; tab-width: 2; -*-
# vim: ts=2 sw=2 ft=bash noet

# Copy the code into the live directory which will be used to run the app
publish_release() {
  nos_print_bullet "Moving code into app directory..."
  rsync -a $(nos_code_dir)/ $(nos_app_dir)
}


# Determine the golang runtime to install. This will first check
# within the Boxfile, then will rely on default_runtime to
# provide a sensible default
runtime() {
  echo $(nos_validate "$(nos_payload "config_runtime")" "string" "go-1.6")
}

# Install the golang runtime.
install_runtime() {
  nos_install "$(runtime)" 'mercurial-3'
}

# Uninstall build dependencies
uninstall_build_dependencies() {
  nos_uninstall "$(runtime)" 'mercurial-3'
}

# 
before() {
  if (nos_validate_presence 'config_before_exec' &> /dev/null) ; then
    prep_env
    nos_run_hooks "before"
  else
    gom_install   && return 0
    godep_restore && return 0
    go_get        && return 0
  fi
}

compile() {
  if (nos_validate_presence 'config_exec' &> /dev/null) ; then
    prep_env
    nos_run_hooks "exec"
  else
    gom_build   && return 0
    godep_build && return 0
    go_build    && return 0
  fi
}

after() {
  if (nos_validate_presence 'config_after_exec' &> /dev/null) ; then
    prep_env
    nos_run_hooks "after"
  fi
}

# Prepare the environment for golang
prep_env() {
  nos_set_evar 'GOPATH' "$(nos_code_dir)/.gopath"
  mkdir -p $GOPATH/bin
  nos_set_evar 'PATH' "$GOPATH/bin:$PATH"
}
