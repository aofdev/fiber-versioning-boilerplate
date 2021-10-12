#!/usr/bin/env bash

# Load file env
eval $(grep -v -e '^#' .env | xargs -I {} echo export \'{}\')
URI=$(echo $MONGO_URI | sed "s/localhost/mongo/g" | sed "s/27022/27017/g")

echo "Dumping database..."
docker run --rm -it -v $(pwd):/tmp --network fiber_versioning_boilerplate mongo:4.4 bash -c "cd /tmp/seed/mongodb/ && mongodump --uri $URI --out dump"
echo "Done."
