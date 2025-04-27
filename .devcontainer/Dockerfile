FROM qmcgaw/godevcontainer:debian

RUN apt-get -qq update && \
    apt-get -qq install -y curl bash zip

# Install docker
# ENV DOCKER_VERSION=27.5.1
# RUN curl -fsSLO https://download.docker.com/linux/static/stable/x86_64/docker-${DOCKER_VERSION}.tgz  \
#     && tar xzvf docker-${DOCKER_VERSION}.tgz --strip 1 \
#         -C /usr/local/bin docker/docker \
#     && rm docker-${DOCKER_VERSION}.tgz

# Fix sam docker credential store issue

# Install aws sam cli
# RUN curl -fsSLO https://github.com/aws/aws-sam-cli/releases/latest/download/aws-sam-cli-linux-x86_64.zip \
#     && unzip aws-sam-cli-linux-x86_64.zip -d sam-installation \
#     && ./sam-installation/install \
#     && rm aws-sam-cli-linux-x86_64.zip

RUN rm -rf /root/.docker/config.json

# ARG USERNAME=devcontainer
# ARG USER_UID=1000
# ARG USER_GID=$USER_UID

# Create the user
# RUN groupadd --gid $USER_GID $USERNAME \
#     && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
#     #
#     # [Optional] Add sudo support. Omit if you don't need to install software after connecting.
#     && apt-get update \
#     && apt-get install -y sudo \
#     && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
#     && chmod 0440 /etc/sudoers.d/$USERNAME

# ********************************************************
# * Anything else you want to do like clean up goes here *
# ********************************************************

# [Optional] Set the default user. Omit if you want to keep the default as root.
# USER $USERNAME

# RUN go install github.com/cosmtrek/air@latest

# Install air
COPY --from=cosmtrek/air:latest /go/bin/air /go/bin/air

# Remove existing Go installation
RUN rm -rf /usr/local/go

# Clean up any cached files
RUN rm -rf /root/.cache/go-build

# Install lts go version
RUN curl -fsSL https://go.dev/dl/go1.24.2.linux-amd64.tar.gz | \
    tar -C /usr/local -xzvf -


VOLUME /workspace
WORKDIR /workspace
