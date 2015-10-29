# -*- mode: bash; tab-width: 2; -*-
# vim: ts=2 sw=2 ft=bash noet

create_boxfile() {
  template \
    "boxfile.mustache" \
    "-" \
    "$(boxfile_payload)"
}

boxfile_payload() {
    cat <<-END
{
  "has_gomfile": $(has_gomfile)
}
END
}

app_name() {
  # payload app
  echo "$(payload app)"
}

live_dir() {
  # payload live_dir
  echo $(payload "live_dir")
}

deploy_dir() {
  # payload deploy_dir
  echo $(payload "deploy_dir")
}

etc_dir() {
  echo $(payload "etc_dir")
}

code_dir() {
  echo $(payload "code_dir")
}

environment() {
  if [[ -n "$(payload 'env_ENVIRONMENT')" ]]; then
    echo "$(payload 'env_ENVIRONMENT')"
  else
    if [[ "$(payload 'platform')" = 'local' ]]; then
      echo "development"
    else
      echo "production"
    fi
  fi
}

runtime() {
  echo $(validate "$(payload "boxfile_runtime")" "string" "go-1.4")
}

install_runtime() {
  install "$(runtime)"
}
