#!/usr/bin/env bash

# Load file env
eval $(grep -v -e '^#' .env | xargs -I {} echo export \'{}\')
URI=$(echo $MONGO_URI | sed "s/localhost/mongo/g" | sed "s/27022/27017/g")

echo "Restoring database..."
docker run --rm -v $(pwd):/tmp --network fiber_versioning_boilerplate mongo:4.4 bash -c "cd /tmp/seed/mongodb/ && mongorestore --uri $URI"
echo "Done."
