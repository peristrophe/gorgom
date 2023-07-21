FROM golang:1.20.5-bullseye as base

FROM base as builder

COPY . /prj
WORKDIR /prj

RUN make build

# ============================== Production stage ==============================

FROM base as production

COPY --from=builder /prj/app /app

RUN useradd peristrophe
USER peristrophe

ENTRYPOINT ["/app"]

# ============================== Development stage ==============================

FROM base as development

ENV TOKEN_SECRET_KEY   2QqMv3rANCWH+NxVpGbesdVVIbJUeaT+a8K4lMucBVo=
ENV APP_HOST           localhost

RUN set -ex && \
    apt-get update && \
    apt-get install -y --no-install-recommends \
            vim \
            jq \
            postgresql-client \
            && \
    apt-get upgrade -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN set -ex && \
    curl -sSL https://raw.githubusercontent.com/peristrophe/dotfiles/master/.vimrc > ${HOME}/.vimrc && \
    git clone https://github.com/VundleVim/Vundle.vim.git ${HOME}/.vim/bundle/Vundle.vim && \
    vim +PluginInstall +qall

RUN go install -v golang.org/x/tools/gopls@latest && \
    go install -v github.com/ramya-rao-a/go-outline@latest && \
    go install -v github.com/golang/mock/mockgen@v1.6.0

RUN echo "PS1='\[\e[1;33m\]\u@\[\e[m\]\[\e[1;32m\]\h:\[\e[m\]\[\e[1;36m\]\w$ \[\e[m\]'" >> /root/.bashrc && \
    echo "alias la='ls -lA --color=auto'" >> /root/.bashrc && \
    echo "[dev]\nhost=172.56.56.100\nport=5432\nuser=test\ndbname=gorgom\n" > /root/.pg_service.conf && \
    echo "[init-dev]\nhost=172.56.56.100\nport=5432\nuser=test\n" >> /root/.pg_service.conf
