#!/bin/bash
zip=$(which zip)

for pdir in bin/*; do
  platform=${pdir//*\/}
  for adir in ${pdir}/*; do
    arch=${adir//*\/}
    for bin in ${adir}/*; do
      [[ ! -d dist/ ]] && mkdir dist
      echo "Building dist/${platform}/${arch}"
      ${zip} -v -j dist/${platform}_${arch}.zip ${bin}
    done
  done
done
