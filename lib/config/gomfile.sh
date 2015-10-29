# -*- mode: bash; tab-width: 2; -*-
# vim: ts=2 sw=2 ft=bash noet

has_gomfile() {
  [[ -f $(code_dir)/Gomfile ]] && echo "true" || echo "false"
}

install_gom() {
  [[ "$(has_gomfile)" = "true" ]] && (cd $(code_dir); run_subprocess "get gom" "go get github.com/mattn/gom")
}
