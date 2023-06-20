FROM golang:1.20 AS base

FROM base AS dev

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update \
    && apt-get -y install --no-install-recommends apt-utils 2>&1

# Verify git, process tools, lsb-release (common in install instructions for CLIs) installed.
RUN apt-get -y install git iproute2 procps lsb-release sudo

WORKDIR /go/src/app

# Install Go tools.
RUN apt-get update \
    # Install gocode-gomod.
    # Install other tools.
    # && go get -v golang.org/x/tools/gopls \
    # && go get github.com/google/wire/cmd/wire \
    && go install -v golang.org/x/tools/gopls@latest \
    && go install -v github.com/ramya-rao-a/go-outline@v0.0.0-20210608161538-9736a4bde949 \
    && go install -v github.com/stamblerre/gocode@v1.0.0 

  #  && go install gotests@latest \
  #  && go install gomodifytags \
  #  && go install impl \
#  RUN  go install -v github.com/go-delve/delve/cmd/dlv@latest \
#     go install -v honnef.co/go/tools/cmd/staticcheck@latest 
   # && go install gopls \
# Installs migrate tool
# RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2 \
#     && chmod -R a+w /go/pkg
# clean up
RUN apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

# Revert workaround at top layer.
ENV DEBIAN_FRONTEND=dialog

ARG USERNAME=runner
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Create the user
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    #
    # [Optional] Add sudo support. Omit if you don't need to install software after connecting.
    && apt-get update \
    && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

WORKDIR /go/src/app

COPY --chown="$SERNAME:$USERNAME" ./ /go/src/app

RUN chown -R $USERNAME:$USERNAME ./

USER $USERNAME

EXPOSE 8080
