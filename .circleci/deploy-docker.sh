# push to gcloud docker registry
if [[ ! -z "${CIRCLE_TAG}" ]]; then
    TAG = echo "$CIRCLE_TAG" | sed 's/^v//g'
    echo "push $CIRCLE_TAG $TAG"
fi
if [[ "${CIRCLE_BRANCH}" == "master" && -z "${CIRCLE_TAG}" ]]; then
    echo "push latest"
fi