# -*- mode: bash; tab-width: 2; -*-
# vim: ts=2 sw=2 ft=bash noet

has_gomfile() {
  [[ -f $(nos_code_dir)/Gomfile ]] && echo "true" || echo "false"
}

install_gom() {
  [[ "$(has_gomfile)" = "true" ]] && (cd $(nos_code_dir); nos_run_subprocess "get gom" "go get github.com/mattn/gom")
}

gom_install() {
  golang_prep_env
  [[ "$(has_gomfile)" = "true" ]] && (cd $(nos_code_dir); nos_run_subprocess "gom install" "gom install")
}

gom_build() {
  golang_prep_env
  [[ "$(has_gomfile)" = "true" ]] && (cd $(nos_code_dir); nos_run_subprocess "gom build" "gom build")
}
