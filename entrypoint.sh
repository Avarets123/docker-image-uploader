#!/bin/bash

echo "----------------------------- STARTING BUILD AND UPLOAD DOCKER IMAGE ----------------------------- "

./opt/dimage_uploader/main
RET=$?

echo "----------------------------- END ----------------------------- "

exit $RET