# -*- mode: bash; tab-width: 2; -*-
# vim: ts=2 sw=2 ft=bash noet

for i in `find ${engine_lib_dir} ! -name golang.sh -type f`; do
  . $i
done
