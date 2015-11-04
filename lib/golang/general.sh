# -*- mode: bash; tab-width: 2; -*-
# vim: ts=2 sw=2 ft=bash noet

golang_create_boxfile() {
  template \
    "boxfile.mustache" \
    "-" \
    "$(boxfile_payload)"
}

golang_boxfile_payload() {
    cat <<-END
{
}
END
}

golang_runtime() {
  echo $(nos_validate "$(nos_payload "boxfile_runtime")" "string" "go-1.4")
}

golang_install_runtime() {
  nos_install "$(golang_runtime)"
}

golang_before() {
  if nos_validate_presence 'boxfile_before_exec' ; then
    nos_run_hooks "before"
  fi
}

golang_exec() {
  if nos_validate_presence 'boxfile_exec' ; then
    nos_run_hooks "exec"
  else
    gom_build && return 0
  fi
}

golang_after() {
  if nos_validate_presence 'boxfile_after_exec' ; then
    nos_run_hooks "after"
  fi
}

golang_prep_env() {
  nos_set_evar 'GOROOT' "$(nos_cache_dir)/go"
  mkdir -p $GOROOT
  nos_set_evar 'GOPATH' "$(nos_code_dir)/go"
  mkdir -p $GOPATH/bin
  nos_set_evar 'PATH'   "$GOROOT/bin:$PATH"
}
