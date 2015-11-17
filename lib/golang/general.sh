# -*- mode: bash; tab-width: 2; -*-
# vim: ts=2 sw=2 ft=bash noet

golang_create_boxfile() {
  nos_template \
    "boxfile.mustache" \
    "-" \
    "$(golang_boxfile_payload)"
}

golang_boxfile_payload() {
    cat <<-END
{
  "code_dir": $(nos_live_dir)
}
END
}

golang_runtime() {
  echo $(nos_validate "$(nos_payload "boxfile_runtime")" "string" "go-1.4")
}

golang_install_runtime() {
  nos_install "$(golang_runtime)"
  nos_install 'mercurial-3'
}

golang_before() {
  if (nos_validate_presence 'boxfile_before_exec' &> /dev/null) ; then
    nos_run_hooks "before"
  else
    gom_install   && return 0
    godep_restore && return 0
  fi
}

golang_exec() {
  if (nos_validate_presence 'boxfile_exec' &> /dev/null) ; then
    nos_run_hooks "exec"
  else
    gom_build   && return 0
    godep_build && return 0
  fi
}

golang_after() {
  if (nos_validate_presence 'boxfile_after_exec' &> /dev/null) ; then
    nos_run_hooks "after"
  fi
}

golang_prep_env() {
  nos_set_evar 'GOPATH' "$(nos_code_dir)/_vendor"
  mkdir -p $GOPATH/bin
  nos_set_evar 'PATH' "$GOPATH/bin:$PATH"
}
