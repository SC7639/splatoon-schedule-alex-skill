sam local invoke \
    --template ./template.yaml \
    --container-host host.docker.internal \
    --container-host-interface "0.0.0.0" \
    --docker-volume-basedir ${LOCAL_WORKSPACE_FOLDER}/.aws-sam/build/ScheduleFunction \
    --event ./events/event.json #\
    # --debug