# -*- mode: bash; tab-width: 2; -*-
# vim: ts=2 sw=2 ft=bash noet

has_godeps() {
  [[ -d $(nos_code_dir)/Godeps ]] && echo "true" || echo "false"
}

install_godep() {
  [[ "$(has_godeps)" = "true" ]] && (cd $(nos_code_dir); nos_run_subprocess "get godep" "go get github.com/tools/godep")
}

godep_build() {
  golang_prep_env
  [[ "$(has_godeps)" = "true" ]] && (cd $(nos_code_dir); nos_run_subprocess "godep go build" "godep go build")
}

godep_restore() {
  golang_prep_env
  [[ "$(has_godeps)" = "true" ]] && (cd $(nos_code_dir); nos_run_subprocess "godep restore" "godep restore")
}
