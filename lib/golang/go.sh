# -*- mode: bash; tab-width: 2; -*-
# vim: ts=2 sw=2 ft=bash noet

go_get() {
  prep_env
  (cd $(nos_code_dir); nos_run_subprocess "go get" "go get")
}

go_build() {
  prep_env
  (cd $(nos_code_dir); nos_run_subprocess "go build" "go build")
}

go_install() {
  prep_env
  (cd $(nos_code_dir); nos_run_subprocess "go install" "go install")
}
